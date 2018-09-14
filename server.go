package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	// "reflect"
	"time"
)

type Room struct {
	User map[string]net.Conn
}

var globalRoom = Room{User: make(map[string]net.Conn)}
var leaveName = make(chan string)

func main() {
	addr := "127.0.0.1:8000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()
	for {
		ms, _ := lis.Accept()
		// fmt.Println(reflect.TypeOf(ms))
		uname := globalRoom.join(ms)
		go handleConnect(ms, uname)

	}

}

func (r *Room) join(u net.Conn) string {
	nickName := []string{"阳刚", "桃子", "北京二环8套房", "三爷", "月圆", "乡土", "瑶瑶"}
	rand.Seed(time.Now().Unix())
	uname := nickName[rand.Intn(len(nickName))]
	r.User[uname] = u
	r.broadCost(uname+"加入群聊\r", uname)
	return uname
}

// 群播
func (r *Room) broadCost(cont string, uname string) {
	for _, cl := range r.User {
		_, err := cl.Write([]byte(cont))
		if err != nil {
			globalRoom.User[uname].Close()
			delete(globalRoom.User, uname)
		}
	}
}

func handleConnect(ms net.Conn, uname string) {
	cli := ms.RemoteAddr()
	fmt.Println(cli.String() + " connected!")
	ms.Write([]byte("欢迎加入直播室，请文明用语！\r"))
	go listenCli(ms, uname)
}

// 监听客户端，并广播
func listenCli(ms net.Conn, uname string) bool {
	r := bufio.NewReader(ms)
Loop:
	for {
		// 当前监听用户不在聊天室时，退出该循环
		con := globalRoom.User[uname]
		if con == nil {
			break Loop
		}
		str, err := r.ReadString('\r')
		if err == io.EOF {
			ms.Close()
		}
		globalRoom.broadCost(uname+":"+str, uname)
	}
	globalRoom.broadCost(uname+"离开房间！", uname)
	return true
}
