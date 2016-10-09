package main

import (
	"log"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

const (
	myDB     = "square_holes"
	username = "bubba"
	password = "bubba"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "htp://localhost:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal("Error: ", err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	tags := map[string]string{
		"cpu": "cpu-total",
	}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	bp.AddPoint(pt)
	c.Write(bp)
}
