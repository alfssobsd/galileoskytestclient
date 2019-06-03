package file

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

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

func ReadConfigEmulateMovment(pathToConfig string) (*MovementRoute, error) {
	var data = new(MovementRoute)

	yamlFile, err := ioutil.ReadFile(pathToConfig)

	log.Println("Read config")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, data)
	if err != nil {
		return nil, err
	}
	//log.Printf("Unmarshal =  %v \n", data)

	return data, nil
}
