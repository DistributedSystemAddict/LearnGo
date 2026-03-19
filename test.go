package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"
)

// 1. Cấu trúc Request
type Request struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Body    io.Reader // Cho phép JSON Decoder đọc thẳng từ TCP socket
}

// 2. Cấu trúc ResponseWriter
type ResponseWriter struct {
	conn        net.Conn
	Headers     map[string]string
	StatusCode  int
	wroteHeader bool
}

func NewResponseWriter(conn net.Conn) *ResponseWriter {
	return &ResponseWriter{
		conn:       conn,
		Headers:    make(map[string]string),
		StatusCode: 200, // Mặc định luôn là 200 OK
	}
}

// Hàm Write biến ResponseWriter thành chuẩn io.Writer
func (w *ResponseWriter) Write(p []byte) (n int, err error) {
	if !w.wroteHeader {
		// Tự động gán Content-Length dựa vào dữ liệu truyền vào
		if w.Headers["Content-Length"] == "" {
			w.Headers["Content-Length"] = fmt.Sprintf("%d", len(p))
		}
		if w.Headers["Content-Type"] == "" {
			w.Headers["Content-Type"] = "text/plain; charset=utf-8"
		}

		// Xây dựng chuỗi Header chuẩn HTTP/1.1
		headerStr := fmt.Sprintf("HTTP/1.1 %d \r\n", w.StatusCode)
		for k, v := range w.Headers {
			headerStr += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		headerStr += "\r\n" // Dòng trống quan trọng ngăn cách Header và Body

		w.conn.Write([]byte(headerStr))
		w.wroteHeader = true
	}
	// Ghi Body thẳng xuống mạng
	return w.conn.Write(p)
}

// 3. TẦNG ROUTER (BỘ ĐIỀU PHỐI)
type HandlerFunc func(w *ResponseWriter, r *Request)

type ServeMux struct {
	// Nested Map: Method -> Path -> Handler (Kiến trúc chuẩn framework)
	routes map[string]map[string]HandlerFunc
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

// Hàm đăng ký Route yêu cầu truyền rõ Method
func (m *ServeMux) HandleFunc(method string, path string, handler HandlerFunc) {
	if m.routes[method] == nil {
		m.routes[method] = make(map[string]HandlerFunc)
	}
	m.routes[method][path] = handler
}

// Cảnh sát giao thông: Kiểm tra Method và Path
func (m *ServeMux) ServeHTTP(w *ResponseWriter, r *Request) {
	methodRoutes, methodExists := m.routes[r.Method]
	if !methodExists {
		w.StatusCode = 405 // Method Not Allowed
		w.Write([]byte("405 - Method Not Allowed\n"))
		return
	}

	handler, pathExists := methodRoutes[r.Path]
	if !pathExists {
		w.StatusCode = 404 // Not Found
		w.Write([]byte("404 - Not Found\n"))
		return
	}

	// Hợp lệ 100% -> Chuyển cho Handler xử lý
	handler(w, r)
}

// 4. PARSER: Bóc tách gói tin TCP
func parseHeaders(reader *bufio.Reader) map[string]string {
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == "\r\n" {
			break
		}
		parts := strings.SplitN(strings.TrimSpace(line), ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	return headers
}

func handleConnection(conn net.Conn, mux *ServeMux) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	// Đọc Request Line (Ví dụ: "POST /create HTTP/1.1")
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) < 3 {
		return
	}

	req := &Request{
		Method:  parts[0],
		Path:    parts[1],
		Proto:   parts[2],
		Headers: parseHeaders(reader),
	}

	// Đọc Body dựa vào Content-Length
	var bodyBytes []byte
	if clStr, ok := req.Headers["Content-Length"]; ok {
		var cl int
		fmt.Sscanf(clStr, "%d", &cl)
		if cl > 0 {
			bodyBytes = make([]byte, cl)
			io.ReadFull(reader, bodyBytes)
		}
	}
	req.Body = bytes.NewReader(bodyBytes)

	w := NewResponseWriter(conn)

	// Gọi bộ định tuyến
	mux.ServeHTTP(w, req)
}

// =====================================================================
// TẦNG APPLICATION: USER CODE (Chỉ tập trung xử lý Business Logic)
// =====================================================================

type ShoppingList struct {
	Item string `json:"item"`
	Qty  int    `json:"qty"`
}

// "Database" giả lập
var allData []ShoppingList

// API Tạo List (Chỉ chạy khi có request POST vào /create)
func handleCreateList(w *ResponseWriter, r *Request) {
	// KHÔNG CẦN check Method ở đây nữa! Framework đã lo việc đó.

	var list ShoppingList
	err := json.NewDecoder(r.Body).Decode(&list)
	if err != nil {
		w.StatusCode = 400
		w.Write([]byte("Bad Request: Sai format JSON\n"))
		return
	}

	allData = append(allData, list)
	fmt.Printf("[LOG] Đã thêm item mới: %+v\n", list)

	// Trả về dữ liệu JSON
	w.StatusCode = 201 // 201 Created
	w.Headers["Content-Type"] = "application/json"
	json.NewEncoder(w).Encode(list)
}

// API Xem danh sách (Chỉ chạy khi có request GET vào /list)
func handleGetList(w *ResponseWriter, r *Request) {
	w.Headers["Content-Type"] = "application/json"
	json.NewEncoder(w).Encode(allData)
}

// =====================================================================
// KHỞI ĐỘNG SERVER
// =====================================================================
func main2() {
	// 1. Khởi tạo Framework Mux
	mux := NewServeMux()

	// 2. Đăng ký rành mạch: METHOD + PATH + HANDLER
	mux.HandleFunc("POST", "/create", handleCreateList)
	mux.HandleFunc("GET", "/list", handleGetList)

	// 3. Mở TCP Server
	port := ":8001"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("🚀 Custom Framework đang chạy tại http://localhost%s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// Xử lý đa luồng (Mỗi request một Goroutine)
		go handleConnection(conn, mux)
	}
}
