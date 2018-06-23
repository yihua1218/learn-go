package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	cer, err := tls.LoadX509KeyPair("certs/cert.pem", "certs/private.key")
	check(err)

	cid := "ASUS"

	// AutoReconnect option is true by default
	// CleanSession option is true by default
	// KeepAlive option is 30 seconds by default
	connOpts := MQTT.NewClientOptions()
	connOpts.SetClientID(cid)
	connOpts.SetMaxReconnectInterval(1 * time.Second)
	connOpts.SetTLSConfig(&tls.Config{Certificates: []tls.Certificate{cer}})

	host := "avzi68bry4as0.iot.ap-northeast-1.amazonaws.com"
	port := 8883
	path := "/mqtt"

	brokerURL := fmt.Sprintf("tcps://%s:%d%s", host, port, path)
	connOpts.AddBroker(brokerURL)
	connOpts.SetOnConnectHandler(onConnect())

	mqttClient := MQTT.NewClient(connOpts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("[MQTT] Connected")

	quit := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		mqttClient.Disconnect(250)
		fmt.Println("[MQTT] Disconnected")

		quit <- struct{}{}
	}()
	<-quit
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Reported message
type Reported struct {
	InterfaceAddrs interface{}
}

// State state message
type State struct {
	reported Reported
}

// Message AWS IoT message
type Message struct {
	state State
}

func onConnect() MQTT.OnConnectHandler {

	return func(client MQTT.Client) {
		log.Println("Running MQTT.OnConnectHandler")
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println("addrs:", addrs)
		}
		message := Message{
			state: State{
				reported: Reported{
					InterfaceAddrs: addrs,
				},
			},
		}
		payload, err := json.Marshal(message)

		fmt.Println("payload:", payload)

		// client.Publish("$aws/things/asus/shadow/update", byte(1), true, payload)
	}

}
