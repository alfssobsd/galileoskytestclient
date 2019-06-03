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

func SendOnePack(host string, port int, pack []byte, expectedResponse []byte) error {
	//connect to server
	conn, err := Connect(host, port)
	if err != nil {
		return err
	}

	//send pack to server
	log.Println("sending pack ", hex.EncodeToString(pack))
	_, err = conn.Write(pack)
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

func Connect(host string, port int) (net.Conn, error) {
	log.Println("Try connect to server ", host, ":", port)
	//connect to server
	conn, err := net.Dial("tcp4", host+":"+strconv.Itoa(port))
	if err != nil {
		return conn, err
	}

	return conn, nil
}

func SendPack(conn net.Conn, pack []byte, expectedResponse []byte) error {
	//send pack to server
	log.Println("sending pack ", hex.EncodeToString(pack))
	_, err := conn.Write(pack)
	if err != nil {
		return err
	}

	log.Println("reading response")
	tcpReadBuf := make([]byte, 1000)
	n, err := conn.Read(tcpReadBuf)
	if err != nil {
		log.Println("error read response = ", err)
		return err
	}

	// return error if response doesn't match
	if !bytes.Equal(tcpReadBuf[:n], expectedResponse) {
		return fmt.Errorf("response doesn't match, expected = %v actual = %v", hex.EncodeToString(expectedResponse), hex.EncodeToString(tcpReadBuf))
	}

	log.Println("got response = ", hex.EncodeToString(tcpReadBuf[:n]))

	return nil
}
