package ch03

import (
	"bufio"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }()

	t.Logf("bound to %q", listener.Addr())

	done := make(chan struct{})
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			go func(c net.Conn) {
				defer c.Close()
				scanner := bufio.NewScanner(c)
				if scanner.Scan() {
					t.Logf("[Server] Đã nhận được tin nhắn: %s", scanner.Text())
				}
				done <- struct{}{}
			}(conn)
		}
	}()

	clientConn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = clientConn.Close() }()

	message := "Xin chào Server, tôi là Client!\n"
	_, err = clientConn.Write([]byte(message))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("[Client] Đã gửi tin nhắn thành công tới %q", listener.Addr())

	<-done
	t.Log("Test hoàn tất thành công!")
}
