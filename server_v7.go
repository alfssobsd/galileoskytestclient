package main

//This isn't complete code, just draft!!!
import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp4", ":9998")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("start listen")

	defer l.Close()

	for {
		fmt.Println("wait data")
		c, err := l.Accept()
		fmt.Println("got data")
		if err != nil {
			fmt.Println(err)
			return
		}

		handler(c)
	}
}

func handler(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	buf := make([]byte, 1000)
	for {
		n, err := bufio.NewReader(c).Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Read: Got error = ", err)
			}
			break
		}

		fmt.Printf("%s", hex.EncodeToString(buf[:n]))

		if hex.EncodeToString(buf[:n]) == "011780011a02e703383634343935303330383631333033043200a3af" {
			fmt.Println()
			answer, _ := hex.DecodeString("02a3af")
			_, _ = c.Write(answer)
			fmt.Println("send answer")
		}
	}
}
