package Smuggling

import (
	"log"
	"strings"

	"github.com/w3security/reconbot/lib/socket"
)

type Base struct {
	Payload []string
	Type    string
}

func (r *Base) New() int {
	return 1
}

func (r *Base) GetTimes() int {
	return 1
}
func (r *Base) CheckResponse(body string, payload string) bool {
	if "" != body && (-1 < strings.Index("Unrecognized method GPOST", body)) {
		log.Println("Unrecognized method GPOST")
		return true
	}
	return false
}

func (r *Base) GetPayloads(t *socket.CheckTarget) *[]string {
	return &r.Payload
}

func (r *Base) GetVulType() string {
	return r.Type
}
