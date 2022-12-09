package mist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"
)

type envelope struct {
	MessageId string `json:"messageId"`
	TraceId   string `json: traceId`
	Payload   string `json: payload`
}

type handlers map[string]func(string)
type iFunc func()

func MistService(hs handlers) {
	action := os.Args[len(os.Args)-2]
	handler := hs[action]
	if handler != nil {
		invokeHandler(action, handler)
	}
}

func MistServiceWithInit(hs handlers, init iFunc) {
	action := os.Args[len(os.Args)-2]
	handler := hs[action]
	if handler != nil {
		invokeHandler(action, handler)
	} else if init != nil {
		init()
	}
}

func invokeHandler(action string, handler func(string)) {
	var e envelope
	err := json.Unmarshal([]byte(os.Args[len(os.Args)-1]), &e)
	if err != nil {
		fmt.Println(err)
	}
	handler(e.Payload)
}

func PostToRapid[T interface{}](event string, reply T) {
	body, _ := json.Marshal(reply)
	PostBodyToRapid(event, bytes.NewBuffer(body))
}

func PostBodyToRapid(event string, body io.Reader) {
	resp, err := http.Post(fmt.Sprintf("%s/%s", os.Getenv("RAPID"), event), "application/json", body)
	if err != nil {
		fmt.Println("Get failed with error: ", err)
	}
	defer resp.Body.Close()
}
