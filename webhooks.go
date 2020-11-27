package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type webhookMessage struct {
	Content    string `json:"content,omitempty"`
	Username   string `json:"username,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Tts        bool   `json:"tts,omitempty"`
}

func sendWebhook(m webhookMessage, w string) {
	msg, _ := json.Marshal(m)
	r := bytes.NewReader(msg)
	http.Post(w, "application/json", r)
}
