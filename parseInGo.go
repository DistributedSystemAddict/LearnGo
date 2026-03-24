// =============================================================
// TỔNG HỢP CÁC KIỂU PARSE TRONG GO
// =============================================================

package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"math/big"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main5() {
	parse01_strconv()
	parse02_json()
	parse03_datetime()
	parse04_sscanf()
	parse05_url()
	parse06_csv()
	parse07_bytes()
	parse08_binary()
	parse09_hex()
	parse10_base64()
	parse11_xml()
	parse12_bufio()
	parse13_numbersystem()
	parse14_regexp()
	parse15_strings()
	parse16_unicode()
	parse17_network()
	parse18_env()
	parse19_flags()
	parse20_bigint()
}

// =============================================================
// 1. STRCONV — Parse String ↔ Number/Bool
// =============================================================
func parse01_strconv() {
	fmt.Println("\n===== 1. STRCONV =====")

	// String → int
	n, err := strconv.Atoi("123")
	fmt.Println(n, err) // 123 <nil>

	// String → int64 (base 10, 64-bit)
	n64, err := strconv.ParseInt("123", 10, 64)
	fmt.Println(n64, err) // 123 <nil>

	// String → uint64
	u64, err := strconv.ParseUint("123", 10, 64)
	fmt.Println(u64, err) // 123 <nil>

	// String → float64
	f, err := strconv.ParseFloat("3.14", 64)
	fmt.Println(f, err) // 3.14 <nil>

	// String → bool
	b, err := strconv.ParseBool("true")
	fmt.Println(b, err) // true <nil>
	// Nhận: "1","t","T","true","TRUE","0","f","false","FALSE"

	// int → String
	s := strconv.Itoa(123)
	fmt.Println(s) // "123"

	// float → String
	s = strconv.FormatFloat(3.14159, 'f', 2, 64) // 'f'=decimal, 2=số chữ số thập phân
	fmt.Println(s)                               // "3.14"

	// bool → String
	s = strconv.FormatBool(true)
	fmt.Println(s) // "true"

	// int64 → String theo base
	s = strconv.FormatInt(255, 16) // hex
	fmt.Println(s)                 // "ff"

	// Xử lý lỗi khi parse thất bại
	_, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Lỗi:", err) // strconv.Atoi: parsing "abc": invalid syntax
	}
}

// =============================================================
// 2. JSON — Parse JSON ↔ Struct/Map
// =============================================================
func parse02_json() {
	fmt.Println("\n===== 2. JSON =====")

	type Address struct {
		City string `json:"city"`
	}
	type User struct {
		Name    string  `json:"name"`
		Age     int     `json:"age"`
		Score   float64 `json:"score,omitempty"` // bỏ qua nếu = 0
		Address Address `json:"address"`
	}

	// JSON string → Struct (Unmarshal)
	data := `{"name":"An","age":25,"address":{"city":"Hanoi"}}`
	var user User
	err := json.Unmarshal([]byte(data), &user)
	fmt.Println(user, err) // {An 25 0 {Hanoi}} <nil>

	// Struct → JSON string (Marshal)
	user2 := User{Name: "Binh", Age: 30, Score: 9.5}
	out, err := json.Marshal(user2)
	fmt.Println(string(out), err) // {"name":"Binh","age":30,"score":9.5,"address":{"city":""}}

	// JSON đẹp (Marshal với indent)
	out, _ = json.MarshalIndent(user2, "", "  ")
	fmt.Println(string(out))

	// Parse JSON không biết cấu trúc trước (dùng map)
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)
	fmt.Println(result["name"].(string)) // "An"
	fmt.Println(result["age"].(float64)) // 25 (JSON number → float64)

	// Parse JSON array
	arrData := `[{"name":"An"},{"name":"Binh"}]`
	var users []User
	json.Unmarshal([]byte(arrData), &users)
	fmt.Println(users[0].Name) // "An"
}

// =============================================================
// 3. DATE/TIME — Parse chuỗi thời gian
// =============================================================
func parse03_datetime() {
	fmt.Println("\n===== 3. DATE/TIME =====")

	// QUAN TRỌNG: Go dùng mốc thời gian cố định: 2006-01-02 15:04:05
	// Mon Jan 2 15:04:05 MST 2006  ←  đây là "thước đo" format

	// String → Time
	t, err := time.Parse("2006-01-02", "2024-03-15")
	fmt.Println(t, err)

	t, err = time.Parse("02/01/2006", "15/03/2024")
	fmt.Println(t, err)

	t, err = time.Parse("2006-01-02 15:04:05", "2024-03-15 10:30:00")
	fmt.Println(t, err)

	t, err = time.Parse(time.RFC3339, "2024-03-15T10:30:00Z")
	fmt.Println(t, err)

	// Time → String
	now := time.Now()
	fmt.Println(now.Format("2006-01-02"))          // "2024-03-15"
	fmt.Println(now.Format("02/01/2006 15:04:05")) // "15/03/2024 10:30:00"
	fmt.Println(now.Format(time.RFC3339))

	// Unix timestamp → Time
	unixTime := time.Unix(1710000000, 0)
	fmt.Println(unixTime)

	// Time → Unix timestamp
	fmt.Println(now.Unix())      // seconds
	fmt.Println(now.UnixMilli()) // milliseconds
}

// =============================================================
// 4. FMT.SSCANF — Parse nhiều giá trị từ 1 string
// =============================================================
func parse04_sscanf() {
	fmt.Println("\n===== 4. FMT.SSCANF =====")

	var name string
	var age int
	var score float64

	// Sscanf — parse theo format cụ thể
	n, err := fmt.Sscanf("An 25 9.5", "%s %d %f", &name, &age, &score)
	fmt.Println(n, name, age, score, err) // 3 An 25 9.5 <nil>

	// Sscan — tự động tách theo khoảng trắng
	fmt.Sscan("Binh 30", &name, &age)
	fmt.Println(name, age) // Binh 30

	// Sscanln — dừng khi gặp newline
	fmt.Sscanln("An 25\nBinh", &name, &age)
	fmt.Println(name, age) // An 25
}

// =============================================================
// 5. URL — Parse URL và query params
// =============================================================
func parse05_url() {
	fmt.Println("\n===== 5. URL =====")

	raw := "https://example.com:8080/path/to/page?name=An&age=25&tags=go&tags=parse#section1"
	u, err := url.Parse(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(u.Scheme)     // "https"
	fmt.Println(u.Host)       // "example.com:8080"
	fmt.Println(u.Hostname()) // "example.com"
	fmt.Println(u.Port())     // "8080"
	fmt.Println(u.Path)       // "/path/to/page"
	fmt.Println(u.Fragment)   // "section1"
	fmt.Println(u.RawQuery)   // "name=An&age=25&tags=go&tags=parse"

	// Query params
	params := u.Query()
	fmt.Println(params.Get("name")) // "An"
	fmt.Println(params.Get("age"))  // "25"
	fmt.Println(params["tags"])     // ["go" "parse"] — nhiều giá trị

	// Encode/Decode URL
	encoded := url.QueryEscape("hello world & more")
	fmt.Println(encoded) // "hello+world+%26+more"

	decoded, _ := url.QueryUnescape(encoded)
	fmt.Println(decoded) // "hello world & more"
}

// =============================================================
// 6. CSV — Parse file CSV
// =============================================================
func parse06_csv() {
	fmt.Println("\n===== 6. CSV =====")

	// Parse CSV từ string
	data := "Tên,Tuổi,Thành phố\nAn,25,Hà Nội\nBinh,30,HCM"
	r := csv.NewReader(strings.NewReader(data))

	// Đọc tất cả records
	records, err := r.ReadAll()
	fmt.Println(err) // <nil>
	for _, row := range records {
		fmt.Println(row) // [Tên Tuổi Thành phố], [An 25 Hà Nội], ...
	}

	// Đọc từng dòng
	r2 := csv.NewReader(strings.NewReader(data))
	for {
		record, err := r2.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
	}

	// Ghi CSV ra string
	var sb strings.Builder
	w := csv.NewWriter(&sb)
	w.WriteAll([][]string{
		{"Tên", "Tuổi"},
		{"An", "25"},
	})
	w.Flush()
	fmt.Println(sb.String())

	// CSV với delimiter khác (ví dụ: dấu chấm phẩy)
	r3 := csv.NewReader(strings.NewReader("An;25;Hà Nội"))
	r3.Comma = ';'
	record, _ := r3.Read()
	fmt.Println(record) // [An 25 Hà Nội]
}

// =============================================================
// 7. BYTES — Parse []byte ↔ string và thao tác bytes
// =============================================================
func parse07_bytes() {
	fmt.Println("\n===== 7. BYTES =====")

	// String → []byte
	s := "Hello World"
	b := []byte(s)
	fmt.Println(b) // [72 101 108 108 111 32 87 111 114 108 100]

	// []byte → String
	s2 := string(b)
	fmt.Println(s2) // "Hello World"

	// Các thao tác với bytes package
	fmt.Println(bytes.Contains(b, []byte("World")))                         // true
	fmt.Println(string(bytes.ToUpper(b)))                                   // "HELLO WORLD"
	fmt.Println(string(bytes.ToLower(b)))                                   // "hello world"
	fmt.Println(string(bytes.TrimSpace([]byte("  hello  "))))               // "hello"
	fmt.Println(string(bytes.Replace(b, []byte("World"), []byte("Go"), 1))) // "Hello Go"

	parts := bytes.Split(b, []byte(" "))
	for _, p := range parts {
		fmt.Println(string(p)) // "Hello", "World"
	}

	fmt.Println(bytes.HasPrefix(b, []byte("Hello"))) // true
	fmt.Println(bytes.HasSuffix(b, []byte("World"))) // true
	fmt.Println(bytes.Count(b, []byte("l")))         // 3
	fmt.Println(bytes.Index(b, []byte("World")))     // 6

	// bytes.Buffer — ghi dần rồi lấy kết quả
	var buf bytes.Buffer
	buf.WriteString("Hello")
	buf.WriteByte(' ')
	buf.WriteString("World")
	fmt.Println(buf.String()) // "Hello World"
}

// =============================================================
// 8. BINARY — Parse bytes có cấu trúc (BigEndian/LittleEndian)
// =============================================================
func parse08_binary() {
	fmt.Println("\n===== 8. BINARY =====")

	// int32 → []byte (LittleEndian)
	buf := new(bytes.Buffer)
	n := int32(1234)
	err := binary.Write(buf, binary.LittleEndian, n)
	fmt.Println(buf.Bytes(), err) // [210 4 0 0] <nil>

	// []byte → int32 (LittleEndian)
	var result int32
	err = binary.Read(bytes.NewReader(buf.Bytes()), binary.LittleEndian, &result)
	fmt.Println(result, err) // 1234 <nil>

	// BigEndian
	buf2 := new(bytes.Buffer)
	binary.Write(buf2, binary.BigEndian, int32(1234))
	fmt.Println(buf2.Bytes()) // [0 0 4 210]

	// Đọc nhiều giá trị liên tiếp
	buf3 := new(bytes.Buffer)
	binary.Write(buf3, binary.LittleEndian, int16(100))
	binary.Write(buf3, binary.LittleEndian, int16(200))

	var a, b int16
	binary.Read(buf3, binary.LittleEndian, &a)
	binary.Read(buf3, binary.LittleEndian, &b)
	fmt.Println(a, b) // 100 200

	// Dùng binary.Size để biết kích thước
	fmt.Println(binary.Size(int32(0))) // 4 bytes
	fmt.Println(binary.Size(int64(0))) // 8 bytes
}

// =============================================================
// 9. HEX — Parse Hex string ↔ []byte
// =============================================================
func parse09_hex() {
	fmt.Println("\n===== 9. HEX =====")

	// []byte → Hex string
	b := []byte("Hello")
	s := hex.EncodeToString(b)
	fmt.Println(s) // "48656c6c6f"

	// Hex string → []byte
	decoded, err := hex.DecodeString("48656c6c6f")
	fmt.Println(string(decoded), err) // "Hello" <nil>

	// Hex dump (xem dữ liệu binary dạng hex đẹp)
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	fmt.Println(hex.Dump(data))
	// 00000000  48 65 6c 6c 6f                                    |Hello|

	// Encode trực tiếp vào dst
	dst := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(dst, b)
	fmt.Println(string(dst)) // "48656c6c6f"
}

// =============================================================
// 10. BASE64 — Parse Base64 ↔ []byte
// =============================================================
func parse10_base64() {
	fmt.Println("\n===== 10. BASE64 =====")

	original := []byte("Hello, World! This is a test.")

	// Encode — Standard (dùng + và /)
	encoded := base64.StdEncoding.EncodeToString(original)
	fmt.Println(encoded) // "SGVsbG8sIFdvcmxkISBUaGlzIGlzIGEgdGVzdC4="

	// Decode — Standard
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded), err) // "Hello, World! This is a test." <nil>

	// URL-safe Encoding (dùng - và _ thay vì + và /)
	// Hay dùng trong JWT, URL params
	urlEncoded := base64.URLEncoding.EncodeToString(original)
	fmt.Println(urlEncoded)

	urlDecoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println(string(urlDecoded))

	// RawStdEncoding — không có padding (=)
	rawEncoded := base64.RawStdEncoding.EncodeToString(original)
	fmt.Println(rawEncoded) // không có dấu = ở cuối
}

// =============================================================
// 11. XML — Parse XML ↔ Struct
// =============================================================
func parse11_xml() {
	fmt.Println("\n===== 11. XML =====")

	type Address struct {
		City string `xml:"city"`
	}
	type User struct {
		XMLName xml.Name `xml:"user"`
		Name    string   `xml:"name"`
		Age     int      `xml:"age,attr"` // age là attribute
		Address Address  `xml:"address"`
	}

	// XML string → Struct (Unmarshal)
	data := `<user age="25"><name>An</name><address><city>Hanoi</city></address></user>`
	var user User
	err := xml.Unmarshal([]byte(data), &user)
	fmt.Println(user, err) // {{ user} An 25 {Hanoi}} <nil>

	// Struct → XML string (Marshal)
	user2 := User{Name: "Binh", Age: 30, Address: Address{City: "HCM"}}
	out, err := xml.Marshal(user2)
	fmt.Println(string(out), err)

	// XML đẹp (với indent)
	out, _ = xml.MarshalIndent(user2, "", "  ")
	fmt.Println(string(out))

	// Parse XML không biết cấu trúc (dùng Decoder)
	decoder := xml.NewDecoder(strings.NewReader(data))
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			fmt.Println("Start:", t.Name.Local)
		case xml.CharData:
			fmt.Println("Text:", string(t))
		case xml.EndElement:
			fmt.Println("End:", t.Name.Local)
		}
	}
}

// =============================================================
// 12. BUFIO — Parse đọc từng dòng / từng word
// =============================================================
func parse12_bufio() {
	fmt.Println("\n===== 12. BUFIO =====")

	data := "dòng 1\ndòng 2\ndòng 3"

	// Đọc từng dòng (ScanLines — mặc định)
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Đọc từng word (ScanWords)
	scanner2 := bufio.NewScanner(strings.NewReader("hello world foo bar"))
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}

	// Đọc từng byte (ScanBytes)
	scanner3 := bufio.NewScanner(strings.NewReader("abc"))
	scanner3.Split(bufio.ScanBytes)
	for scanner3.Scan() {
		fmt.Printf("%x ", scanner3.Bytes())
	}
	fmt.Println()

	// Đọc từng rune (ScanRunes)
	scanner4 := bufio.NewScanner(strings.NewReader("Xin chào"))
	scanner4.Split(bufio.ScanRunes)
	for scanner4.Scan() {
		fmt.Printf("%s ", scanner4.Text())
	}
	fmt.Println()

	// bufio.Reader — đọc đến delimiter tùy chỉnh
	reader := bufio.NewReader(strings.NewReader("hello;world;foo"))
	for {
		s, err := reader.ReadString(';')
		fmt.Print(s) // in cả delimiter
		if err != nil {
			break
		}
	}
	fmt.Println()
}

// =============================================================
// 13. NUMBER SYSTEM — Parse số hệ cơ số 2, 8, 16
// =============================================================
func parse13_numbersystem() {
	fmt.Println("\n===== 13. HỆ CƠ SỐ =====")

	// Binary (hệ 2) → int
	n, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println(n) // 10

	// Octal (hệ 8) → int
	n, _ = strconv.ParseInt("17", 8, 64)
	fmt.Println(n) // 15

	// Hex (hệ 16) → int
	n, _ = strconv.ParseInt("FF", 16, 64)
	fmt.Println(n) // 255

	n, _ = strconv.ParseInt("ff", 16, 64) // chữ thường cũng được
	fmt.Println(n)                        // 255

	// int → String theo hệ cơ số
	fmt.Println(strconv.FormatInt(10, 2))   // "1010"  (binary)
	fmt.Println(strconv.FormatInt(15, 8))   // "17"    (octal)
	fmt.Println(strconv.FormatInt(255, 16)) // "ff"   (hex)

	// Go literals hệ cơ số (trong code)
	bin := 0b1010               // binary
	oct := 0o17                 // octal
	hexa := 0xFF                // hex
	fmt.Println(bin, oct, hexa) // 10 15 255
}

// =============================================================
// 14. REGEXP — Parse bằng biểu thức chính quy
// =============================================================
func parse14_regexp() {
	fmt.Println("\n===== 14. REGEXP =====")

	// Tìm kiếm
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.FindString("abc 123 def"))     // "123"
	fmt.Println(re.FindAllString("1 22 333", -1)) // [1 22 333]
	fmt.Println(re.FindAllString("1 22 333", 2))  // [1 22] — tối đa 2 kết quả

	// Match (kiểm tra có khớp không)
	matched, _ := regexp.MatchString(`^\d+$`, "12345")
	fmt.Println(matched) // true

	// Bắt nhóm (capture groups)
	re2 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	match := re2.FindStringSubmatch("an@gmail.com")
	fmt.Println(match[0]) // "an@gmail.com"  — toàn bộ match
	fmt.Println(match[1]) // "an"            — group 1
	fmt.Println(match[2]) // "gmail"         — group 2
	fmt.Println(match[3]) // "com"           — group 3

	// Named groups
	re3 := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+)\.(?P<tld>\w+)`)
	match3 := re3.FindStringSubmatch("an@gmail.com")
	names := re3.SubexpNames()
	for i, name := range names {
		if i != 0 && name != "" {
			fmt.Printf("%s: %s\n", name, match3[i])
		}
	}

	// Thay thế
	re4 := regexp.MustCompile(`\d+`)
	result := re4.ReplaceAllString("abc 123 def 456", "NUM")
	fmt.Println(result) // "abc NUM def NUM"

	// Thay thế với hàm tùy chỉnh
	result2 := re4.ReplaceAllStringFunc("price: 100, qty: 5", func(s string) string {
		n, _ := strconv.Atoi(s)
		return strconv.Itoa(n * 2) // nhân đôi tất cả số
	})
	fmt.Println(result2) // "price: 200, qty: 10"

	// Split bằng regex
	re5 := regexp.MustCompile(`\s+`)
	parts := re5.Split("hello   world\t\tfoo", -1)
	fmt.Println(parts) // [hello world foo]
}

// =============================================================
// 15. STRINGS — Thao tác và parse string
// =============================================================
func parse15_strings() {
	fmt.Println("\n===== 15. STRINGS =====")

	s := "  Hello, World!  "

	// Trim
	fmt.Println(strings.TrimSpace(s))                        // "Hello, World!"
	fmt.Println(strings.Trim("***hello***", "*"))            // "hello"
	fmt.Println(strings.TrimLeft("***hello***", "*"))        // "hello***"
	fmt.Println(strings.TrimRight("***hello***", "*"))       // "***hello"
	fmt.Println(strings.TrimPrefix("hello_world", "hello_")) // "world"
	fmt.Println(strings.TrimSuffix("hello_world", "_world")) // "hello"

	// Split
	fmt.Println(strings.Split("a,b,c", ","))      // [a b c]
	fmt.Println(strings.SplitN("a,b,c", ",", 2))  // [a b,c] — tối đa 2 phần
	fmt.Println(strings.SplitAfter("a,b,c", ",")) // [a, b, c]
	fmt.Println(strings.Fields("a  b  c"))        // [a b c] — tách theo whitespace

	// Tách theo nhiều delimiter
	fmt.Println(strings.FieldsFunc("a1b2c3d", func(r rune) bool {
		return r >= '0' && r <= '9' // tách theo số
	})) // [a b c d]

	// Search
	fmt.Println(strings.Contains("hello world", "world")) // true
	fmt.Println(strings.HasPrefix("hello", "he"))         // true
	fmt.Println(strings.HasSuffix("hello", "lo"))         // true
	fmt.Println(strings.Index("hello", "ll"))             // 2
	fmt.Println(strings.LastIndex("hello", "l"))          // 3
	fmt.Println(strings.Count("hello", "l"))              // 2

	// Transform
	fmt.Println(strings.ToLower("HELLO"))            // "hello"
	fmt.Println(strings.ToUpper("hello"))            // "HELLO"
	fmt.Println(strings.Title("hello world"))        // "Hello World"
	fmt.Println(strings.Replace("aaa", "a", "b", 2)) // "bba" — thay 2 lần
	fmt.Println(strings.ReplaceAll("aaa", "a", "b")) // "bbb"
	fmt.Println(strings.Repeat("ab", 3))             // "ababab"

	// Join
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // "a-b-c"

	// Builder — ghép chuỗi hiệu quả (thay cho + nhiều lần)
	var builder strings.Builder
	for i := 0; i < 5; i++ {
		builder.WriteString(strconv.Itoa(i))
		builder.WriteString(" ")
	}
	fmt.Println(builder.String()) // "0 1 2 3 4 "

	// Reader — đọc string như io.Reader
	r := strings.NewReader("hello")
	buf := make([]byte, 3)
	r.Read(buf)
	fmt.Println(string(buf)) // "hel"
}

// =============================================================
// 16. UNICODE/UTF8 — Parse chuỗi Unicode
// =============================================================
func parse16_unicode() {
	fmt.Println("\n===== 16. UNICODE/UTF8 =====")

	s := "Xin chào Go"

	// Đếm ký tự Unicode (không phải bytes)
	fmt.Println(utf8.RuneCountInString(s)) // 11 ký tự
	fmt.Println(len(s))                    // 13 bytes (vì "à" và "à" chiếm 2 bytes)

	// Kiểm tra UTF-8 hợp lệ
	fmt.Println(utf8.ValidString(s)) // true

	// Lặp ĐÚNG qua từng ký tự Unicode
	for i, r := range s {
		fmt.Printf("byte[%d] = %c (U+%04X)\n", i, r, r)
	}

	// String → []rune (để xử lý từng ký tự)
	runes := []rune(s)
	fmt.Println(string(runes[4])) // lấy ký tự thứ 5

	// []rune → String
	s2 := string([]rune{'H', 'e', 'l', 'l', 'o'})
	fmt.Println(s2) // "Hello"

	// Kiểm tra loại ký tự
	import_unicode_example()
}

func import_unicode_example() {
	// Cần import "unicode"
	// unicode.IsLetter('A')  → true
	// unicode.IsDigit('5')   → true
	// unicode.IsSpace(' ')   → true
	// unicode.IsUpper('A')   → true
	// unicode.IsLower('a')   → true
	// unicode.ToUpper('a')   → 'A'
	// unicode.ToLower('A')   → 'a'
	fmt.Println("Xem comment trong code cho unicode package")
}

// =============================================================
// 17. NETWORK — Parse IP, địa chỉ mạng
// =============================================================
func parse17_network() {
	fmt.Println("\n===== 17. NETWORK =====")

	// Parse IP
	ip := net.ParseIP("192.168.1.1")
	fmt.Println(ip)              // 192.168.1.1
	fmt.Println(ip.IsPrivate())  // true
	fmt.Println(ip.IsLoopback()) // false

	ip2 := net.ParseIP("127.0.0.1")
	fmt.Println(ip2.IsLoopback()) // true

	ip3 := net.ParseIP("::1")     // IPv6 loopback
	fmt.Println(ip3.IsLoopback()) // true

	// Parse CIDR (dải địa chỉ mạng)
	_, network, err := net.ParseCIDR("192.168.1.0/24")
	fmt.Println(network, err)                                   // 192.168.1.0/24 <nil>
	fmt.Println(network.Contains(net.ParseIP("192.168.1.100"))) // true
	fmt.Println(network.Contains(net.ParseIP("192.168.2.1")))   // false

	// Parse host:port
	host, port, err := net.SplitHostPort("localhost:8080")
	fmt.Println(host, port, err) // localhost 8080 <nil>

	host, port, err = net.SplitHostPort("[::1]:9090") // IPv6
	fmt.Println(host, port, err)                      // ::1 9090 <nil>

	// Ghép host:port
	addr := net.JoinHostPort("localhost", "8080")
	fmt.Println(addr) // "localhost:8080"

	// Parse MAC address
	mac, err := net.ParseMAC("00:1A:2B:3C:4D:5E")
	fmt.Println(mac, err) // 00:1a:2b:3c:4d:5e <nil>
}

// =============================================================
// 18. ENV — Parse biến môi trường
// =============================================================
func parse18_env() {
	fmt.Println("\n===== 18. ENV =====")

	// Set và Get biến môi trường
	os.Setenv("APP_HOST", "localhost")
	os.Setenv("APP_PORT", "8080")

	host := os.Getenv("APP_HOST")
	fmt.Println(host) // "localhost"

	// LookupEnv — phân biệt "không có" và "rỗng"
	port, exists := os.LookupEnv("APP_PORT")
	fmt.Println(port, exists) // "8080" true

	_, exists2 := os.LookupEnv("KHÔNG_TỒN_TẠI")
	fmt.Println(exists2) // false

	// Lấy tất cả biến môi trường
	// envs := os.Environ() // []string dạng "KEY=VALUE"

	// Parse .env file thủ công
	envContent := "DB_HOST=localhost\nDB_PORT=5432\nDB_NAME=mydb"
	scanner := bufio.NewScanner(strings.NewReader(envContent))
	envMap := make(map[string]string)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue // bỏ qua dòng trống và comment
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	fmt.Println(envMap) // map[DB_HOST:localhost DB_NAME:mydb DB_PORT:5432]
}

// =============================================================
// 19. FLAG — Parse command-line arguments
// =============================================================
func parse19_flags() {
	fmt.Println("\n===== 19. FLAG =====")

	// Định nghĩa flags
	// name := flag.String("name", "World", "tên cần chào")
	// age := flag.Int("age", 0, "tuổi")
	// verbose := flag.Bool("v", false, "verbose mode")
	// flag.Parse()
	//
	// Chạy: go run main.go -name=An -age=25 -v
	// fmt.Printf("Hello %s, %d tuổi\n", *name, *age)

	// Dùng os.Args trực tiếp
	args := os.Args
	fmt.Println("Program:", args[0])
	if len(args) > 1 {
		fmt.Println("Args:", args[1:])
	}

	// FlagSet — tạo bộ flag riêng (không dùng global)
	fs := flag.NewFlagSet("myapp", flag.ContinueOnError)
	name := fs.String("name", "World", "tên")
	err := fs.Parse([]string{"-name=An"})
	fmt.Println(*name, err) // An <nil>
}

// =============================================================
// 20. BIG INT/FLOAT — Parse số rất lớn
// =============================================================
func parse20_bigint() {
	fmt.Println("\n===== 20. BIG INT/FLOAT =====")

	// big.Int — số nguyên không giới hạn
	n := new(big.Int)
	n.SetString("123456789012345678901234567890", 10)
	fmt.Println(n) // 123456789012345678901234567890

	// Tính toán với big.Int
	a, _ := new(big.Int).SetString("999999999999999999999", 10)
	b, _ := new(big.Int).SetString("1", 10)
	result := new(big.Int).Add(a, b)
	fmt.Println(result) // 1000000000000000000000

	// big.Float — số thực độ chính xác cao
	f := new(big.Float).SetPrec(256) // 256 bits precision
	f.SetString("3.14159265358979323846264338327950288")
	fmt.Println(f)

	// big.Rat — phân số chính xác
	r := new(big.Rat).SetFrac(
		big.NewInt(1),
		big.NewInt(3),
	)
	fmt.Println(r.FloatString(10)) // "0.3333333333"

	// Hex → big.Int
	hexN := new(big.Int)
	hexN.SetString("DEADBEEF", 16)
	fmt.Println(hexN) // 3735928559
}

// =============================================================
// BẢNG TỔNG HỢP
// =============================================================
//
// Parse gì?                      Package / Cách
// ──────────────────────────────────────────────────────────
// String ↔ int/float/bool        strconv
// String → nhiều giá trị         fmt.Sscanf / strings.Split
// []byte ↔ string                ép kiểu trực tiếp []byte(s)
// Bytes manipulation             bytes
// String manipulation            strings
// Bytes có cấu trúc              encoding/binary
// Hex string ↔ []byte            encoding/hex
// Base64 ↔ []byte                encoding/base64
// JSON ↔ Struct/Map              encoding/json
// XML ↔ Struct                   encoding/xml
// HTML (external)                golang.org/x/net/html
// YAML (external)                gopkg.in/yaml.v3
// TOML (external)                github.com/BurntSushi/toml
// Date/Time                      time.Parse / Format
// URL + query params             net/url
// CSV                            encoding/csv
// ENV file                       os + bufio
// Arguments                      flag / os.Args
// IP / Network / MAC             net
// Đọc từng dòng/word/byte/rune   bufio.Scanner
// Hệ cơ số 2/8/16                strconv.ParseInt với base
// Format phức tạp                regexp
// Unicode / Rune                 unicode/utf8 + unicode
// Số rất lớn                     math/big
// =============================================================
