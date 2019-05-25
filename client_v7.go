package main

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net"
)

type TestData struct {
	Messages []TestMessage `yaml:"messages"`
}

type TestMessage struct {
	SendMessage        string `yaml:"send"`
	AnswerMessage      string `yaml:"answer"`
	EnableSplitMessage bool   `yaml:"split"`
}

func handleServerAnwerOnMessage(testName string, element TestMessage, conn net.Conn) {
	tmp := make([]byte, 256)
	n, err := conn.Read(tmp)
	if err != nil {
		if err != io.EOF {
			log.Println("[", testName, "]Read error = ", err)
		}
		log.Fatalln("[", testName, "]got EOF", err)
	}
	log.Println("[", testName, "]Read answer bytes = ", n)
	if hex.EncodeToString(tmp[:n]) != element.AnswerMessage {
		log.Fatalln("[", testName, "]Error expected answer = ", element.AnswerMessage, " server answer = ", hex.EncodeToString(tmp[:n]))
	}
	log.Println("[", testName, "]Server answer = ", hex.EncodeToString(tmp[:n]))
}

func runTest(testName string, messages []TestMessage) {
	conn, err := net.Dial("tcp4", "127.0.0.1:9998")
	if err != nil {
		log.Fatalln("[", testName, "]Error connect to server = ", err)
	}

	defer conn.Close()

	for _, element := range messages {
		log.Println("[", testName, "]Start send data", element)
		log.Println("[", testName, "]enable split message = ", element.EnableSplitMessage)
		if element.EnableSplitMessage {
			b, _ := hex.DecodeString(element.SendMessage)
			_, _ = conn.Write(b[:500])
			_, _ = conn.Write(b[500:])
			handleServerAnwerOnMessage(testName, element, conn)
		} else {
			b, _ := hex.DecodeString(element.SendMessage)
			_, _ = conn.Write(b)
			handleServerAnwerOnMessage(testName, element, conn)
		}
	}
}

func main() {
	fmt.Println("Test Client GalileoSky 7.0")
	var testData map[string]TestData
	yamlFile, err := ioutil.ReadFile("client_v7_test_data.yml")
	log.Println("Read test data")
	if err != nil {
		log.Fatalf("Error = %v", err)
	}
	log.Println("Unmarshal test data")
	err = yaml.Unmarshal(yamlFile, &testData)
	if err != nil {
		log.Fatalf("Unmarshal =  %v", err)
	}
	log.Printf("Unmarshal =  %v \n", testData)

	for key, value := range testData {
		log.Println("[START] Test = ", key)
		runTest(key, value.Messages)
		log.Println("[DONE] Test = ", key)
	}
}
