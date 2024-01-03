package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	mqttBroker   = "tcp://test.mosquitto.org:1883" //URL do broker MQTT usado para comunicação MQTT na porta 1883
	mqttTopic    = "iot/data"                      //Tópico MQTT ao qual as mensagens são publicadas
	httpEndpoint = "http://localhost:8080"         //Endpoint HTTP para o servidor na porta 8080
)

type Message struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
}

var receivedMessages []Message

func main() {
	//cria um cliente MQTT para se conectar ao broker
	opts := mqtt.NewClientOptions().AddBroker(mqttBroker)
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	defer mqttClient.Disconnect(250)
	//
	go simulateMQTTData(mqttClient)
	http.HandleFunc("/receive", receiveHTTPData)
	http.HandleFunc("/messages", getMessages)
	//configura o servidor HTTP para servir arquivos estaticos na pasta public
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Println("Servidor HTTP iniciado em", httpEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// funcao para simulacao de dados MQTT, onde a funcao eh gerad dados de ID e temperatura a cada segundo
// e publicar no topico iot/data definido em mqttTopic
func simulateMQTTData(client mqtt.Client) {
	for {
		deviceID := "device123"
		temperature := rand.Float64() * 100

		message := Message{
			DeviceID:    deviceID,
			Temperature: temperature,
		}

		payload, err := json.Marshal(message)
		if err != nil {
			log.Println("Erro ao converter para JSON:", err)
			continue
		}

		token := client.Publish(mqttTopic, 1, false, payload)
		token.Wait()
		if token.Error() != nil {
			log.Println("Erro ao publicar mensagem MQTT:", token.Error())
		} else {
			log.Printf("Enviada mensagem via MQTT: %+v\n", message)
			receivedMessages = append(receivedMessages, message)
		}

		time.Sleep(1 * time.Second)
	}
}

// funcao de envio de mensagem HTTP, onde envia uma mensagem HTTP POST para o endpoint /receive
// registra a resposta do servidor no log.
func sendHTTPMessage(message Message) {
	payload, err := json.Marshal(message)
	if err != nil {
		log.Println("Erro ao converter para JSON:", err)
		return
	}

	resp, err := http.Post(httpEndpoint+"/receive", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Println("Erro ao enviar mensagem via HTTP:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao ler corpo da resposta HTTP:", err)
		return
	}

	log.Printf("Resposta do servidor HTTP: %s\n", body)
}

// funcao de recebimento de mensagem HTTP, onde decodifica uma mensagem JSON recebida via HTTP
// registra a mensagem no log e adiciona a variavel global receiveMessahes
// retorna uma resposta ao cliente HTTP.
func receiveHTTPData(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	log.Printf("Mensagem recebida via HTTP: %+v\n", message)
	receivedMessages = append(receivedMessages, message)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mensagem recebida com sucesso"))
}

// funcao de obter mensagens, retorna as mensagens recebidas em formato JSON
func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receivedMessages)
}
