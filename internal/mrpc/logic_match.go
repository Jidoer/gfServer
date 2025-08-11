package mrpc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	// "time"
)

type ReqPlayers struct {
	Passport   string `json:"passport"`
	UID        int64  `json:"uid"`
	Key        string `json:"key"`
	Nickname   string `json:"nickname"`
	PlayerType string `json:"player_type"`
	CreatedAt  int64  `json:"created_at,omitempty"`
	Version    int    `json:"version,omitempty"`
}

type EBody struct {
	GameType int         `json:"game_type"`
	Players  []ReqPlayers `json:"players"`
}

func EncodePlayerToBase64(p EBody) (string, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func CreateMatchEnv(p []ReqPlayers) (string,error){
	ebody := EBody{
		GameType: 1,
		Players: p,
	}

	enc, err := EncodePlayerToBase64(ebody)
	if err != nil {
		log.Println(err)
		return "",err
	}
	fmt.Println("Base64 编码字符串:")
	fmt.Println(enc)
	// 下面演示如何快速解回 JSON（仅作演示）
	jsonBytes, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println("\n解码得到的 JSON:")
	fmt.Println(string(jsonBytes))
	return enc,nil
}
