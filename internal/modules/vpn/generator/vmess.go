package generator

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type vmessConfig struct {
	V    string `json:"v"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	ID   string `json:"id"`
	Aid  string `json:"aid"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	TLS  string `json:"tls"`
}

func GenerateVMess(
	identity *Identity,
	server ServerEndpoint,
) string {
	cfg := vmessConfig{
		V:    "2",
		Ps:   "Nebula",
		Add:  server.Host,
		Port: strconv.Itoa(server.Port),
		ID:   identity.UserUUID,
		Aid:  "0",
		Net:  "tcp",
		Type: "none",
		Host: "",
		Path: "",
		TLS:  "tls",
	}

	data, err := json.Marshal(cfg)

	if err != nil {
		return ""
	}

	return "vmess://" + base64.StdEncoding.EncodeToString(data)
}
