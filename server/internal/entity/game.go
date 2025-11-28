package entity

type Game struct {
	Id              string `json:"id"`
	Quiz            Quiz   `json:"quiz"`
	CurrentQuestion int    `json:"current_question"`
	Code            string `json:"code"`
}
