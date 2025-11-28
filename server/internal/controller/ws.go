package controller

import (
	"log"

	"github.com/asliddinberdiev/kahoot/internal/service"
	"github.com/gofiber/contrib/websocket"
)

type WebsocketControlller struct {
	netService *service.NetService
}

func Ws(netService *service.NetService) WebsocketControlller {
	return WebsocketControlller{
		netService: netService,
	}
}

func (c WebsocketControlller) Ws(con *websocket.Conn) {
	var (
		mt  int
		msg []byte
		err error
	)

	for {
		if mt, msg, err = con.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}

		c.netService.OnIncomingMessage(con, mt, msg)
	}
}
