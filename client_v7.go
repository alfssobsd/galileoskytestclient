package main

import (
	"encoding/hex"
	"github.com/alfssobsd/galileoskytestclient/usecases"
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

type MovementRoute struct {
	Route             []string `yaml:"route"`
	IMEI              string   `yaml:"imei"`
	HwVersion         string   `yaml:"hw_version"`
	FwVersion         string   `yaml:"fw_version"`
	DeviceModel       string   `yaml:"device_model"`
	SpeedKm           string   `yaml:"speed_km"`
	HightMeters       string   `yaml:"hight_meters"`
	HDOP              string   `yaml:"hdop"`
	HWStatus          string   `yaml:"hw_status"`
	IntervalInSeconds int      `yaml:"interval_send_signals_seconds"`
}

func handleServerAnwerOnMessage(testName string, element TestMessage, conn net.Conn) {
	tcpReadBuf := make([]byte, 1000)

	n, err := conn.Read(tcpReadBuf)
	if err != nil {

		if err != io.EOF {
			log.Println("[", testName, "]Read error = ", err)
		}

		log.Fatalln("[", testName, "]got EOF", err)
	}

	log.Println("[", testName, "]Read answer bytes = ", n)
	if hex.EncodeToString(tcpReadBuf[:n]) != element.AnswerMessage {
		log.Fatalln("[", testName, "]Error expected answer = ", element.AnswerMessage, " server answer = ", hex.EncodeToString(tcpReadBuf[:n]))
	}

	log.Println("[", testName, "]Server answer = ", hex.EncodeToString(tcpReadBuf[:n]))
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
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
	log.Println("Test Client GalileoSky 7.0")

	usecases.SendOneMessageUseCase("127.0.0.1", 9998,
		"011780011a02e703383634343935303330383631333033043200a3af", "02a3af")
	//hexSender()
	//emulateMovmentUseCase()
}

//Send hex from file
func hexSender() {
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

//Emulate movement
func emulateMovmentUseCase() {

	var testData MovementRoute

	yamlFile, err := ioutil.ReadFile("client_v7_test_movment.yml")

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
	for _, element := range testData.Route {
		println(element)
	}
}
