package service

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/asliddinberdiev/kahoot/internal/entity"
	"github.com/gofiber/contrib/websocket"
)

type NetService struct {
	quizService *QuizService

	host *websocket.Conn
	tick int
}

func Net(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}

type ConnectPacket struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type HostGamePacket struct {
	QuizId string `json:"quiz_id"`
}

type QuestionShowPacket struct {
	Question entity.QuizQuestion `json:"question"`
}

func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	if len(msg) < 2 {
		return
	}

	packetId := msg[0]
	data := msg[1:]

	packet := c.packetIdToPacket(packetId)
	if packet == nil {
		log.Println("packet:", packet)
		return
	}

	if err := json.Unmarshal(data, packet); err != nil {
		log.Println(err)
		return
	}

	switch data := packet.(type) {
	case *ConnectPacket:
		{
			log.Println(data.Name, "wants to join game", data.Code)
			break
		}
	case *HostGamePacket:
		{
			log.Println("User wants to host quiz", data.QuizId)
			go func() {
				time.Sleep(time.Second * 5)
				c.SendPacket(con, QuestionShowPacket{
					Question: entity.QuizQuestion{
						Name: "What is 2+2?",
						Choices: []entity.QuizChoice{
							{
								Name: "4",
							},
							{
								Name: "9",
							},
							{
								Name: "11",
							},
							{
								Name: "elephant",
							},
						},
					},
				})
			}()
			break
		}
	}
}

func (c *NetService) PacketToBytes(packet any) ([]byte, error) {
	packetId, err := c.packetToPacketId(packet)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}

	final := append([]byte{packetId}, bytes...)
	return final, nil
}

func (c *NetService) SendPacket(con *websocket.Conn, packet any) error {
	bytes, err := c.PacketToBytes(packet)
	if err != nil {
		return err
	}

	return con.WriteMessage(websocket.BinaryMessage, bytes)
}

func (c *NetService) packetIdToPacket(packetId uint8) any {
	switch packetId {
	case 0:
		{
			return &ConnectPacket{}
		}
	case 1:
		{
			return &HostGamePacket{}
		}
	}

	return nil
}

func (c *NetService) packetToPacketId(packet any) (uint8, error) {
	switch packet.(type) {
	case QuestionShowPacket:
		{
			return 2, nil
		}
	}

	return 0, errors.New("invalid packet type")
}
