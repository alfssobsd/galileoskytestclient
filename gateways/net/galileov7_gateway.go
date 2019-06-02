package net

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
)

type GalileoError struct {
	ErrorMessage string
}

func (e GalileoError) Error() string {
	return fmt.Sprintf("%v", e.ErrorMessage)
}

func SendOneMessage(host string, port int, message []byte, expectedResponse []byte) error {
	log.Println("Try connect to server ", host, ":", port)
	//connect to server
	conn, err := net.Dial("tcp4", host+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	//send message to server
	log.Println("sending message ", hex.EncodeToString(message))
	_, err = conn.Write(message)
	if err != nil {
		return err
	}

	log.Println("reading response")
	tcpReadBuf := make([]byte, 1000)
	n, err := conn.Read(tcpReadBuf)
	if err != nil {
		return err
	}

	// return error if response doesn't match
	if !bytes.Equal(tcpReadBuf[:n], expectedResponse) {
		return fmt.Errorf("response doesn't match, expected = %v actual = %v", hex.EncodeToString(expectedResponse), hex.EncodeToString(tcpReadBuf))
	}

	return nil
}

func SendData() error {

	return nil
}
