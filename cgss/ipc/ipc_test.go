package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, request string) *Response {
	var body = "Echo:" + request
	return &Response{Code: "200", Body: body}
}

func (server *EchoServer) Name() string {
	return "EechoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, err1 := client1.Call("echo", "From Client 1")
	resp2, err2 := client2.Call("echo", "From Client 2")
	if err1 != nil {
		return
	}
	if err2 != nil {
		return
	}
	if resp1.Body != "Echo:From Client 1" || resp2.Body != "Echo:From Client 2" {
		t.Error(resp1, resp2)
	}
	client1.Close()
	client2.Close()
}
