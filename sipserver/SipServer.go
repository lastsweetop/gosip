package sipserver

import (
	log "github.com/gogap/logrus"
	"net"
)

const (
	STATE_ON  = 1
	STATE_OFF = 2
)

type SipServer struct {
	port      int
	closeChan chan string
	state     byte
}

func NewSipServer(port int) *SipServer {
	return &SipServer{port: port, state: STATE_OFF, closeChan: make(chan string)}
}

func (this *SipServer) Start() {
	// 创建监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: this.port,
	})
	if err != nil {
		log.Error("端口被占用!", err)
		return
	}
	defer socket.Close()
	this.state = STATE_ON
	select {
	case msg := <-this.closeChan:
		log.Debug(msg)
	}
}

func (this *SipServer) Close() {
	if this.state == STATE_ON {
		this.closeChan <- "用户主动关闭sip服务，正常退出"
	}
}
