package main

import (
	"fmt"
	"flag"
	"strings"

	"github.com/gorilla/websocket"
)

//func Dial(url_, protocol, origin string) (ws *Conn, err error)
func main() {
	url := "ws://service-j3301z72-1255960573.shjr.apigw.tencentcs.com:80/ws"
	msgLen := flag.Int("n", 10, "msg len")
	flag.Parse()

	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	msg := strings.Repeat("a", *msgLen)

	err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	
	return
}

