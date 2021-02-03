package main

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"io/ioutil"

	"github.com/gorilla/websocket"
)

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

//func Dial(url_, protocol, origin string) (ws *Conn, err error)
func main() {
	url := flag.String("url", "ws://service-j3301z72-1255960573.shjr.apigw.tencentcs.com:80/ws", "url")
	msgLen := flag.Int("n", 10, "msg len")
	filePath := flag.String("f", "", "file path")
	flag.Parse()

	ws, _, err := websocket.DefaultDialer.Dial(*url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	msg := ""
	if *filePath == "" {
		msg = strings.Repeat("a", *msgLen)
	} else {
		msgBuff, err := ReadAll(*filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg = string(msgBuff)
	}

	err = ws.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}

	
	return
}

