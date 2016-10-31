package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateTopic create a topic
func CreateTopic(topic string) (err error) {
	return nil
}

// DeleteTopic delete topic
func DeleteTopic(topic string) (err error) {
	addr := "127.0.0.1:4161"
	url := fmt.Sprintf("http://%s/topic/delete?topic=%s", addr, topic)
	req, err := http.NewRequest("POST", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		GetLog().Error("utils.DeleteTopic error: %+v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return
}

// DeleteChannel delete channel
func DeleteChannel(topic, channel string) (err error) {
	Dump(topic, channel)
	addr := "127.0.0.1:4161"
	url := fmt.Sprintf("http://%s/channel/delete?topic=%s&channel=%s", addr, topic, channel)
	log.Println("url:", url)
	req, err := http.NewRequest("POST", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		GetLog().Error("utils.DeleteChannel error: %+v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return
}
