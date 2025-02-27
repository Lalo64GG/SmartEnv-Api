package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/lalo64/SmartEnv-api/src/records/application/services"
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
)

type KafkaHelper struct {}

func NewKafkaHelper() (services.KafkaService, error) {
	return &KafkaHelper{}, nil
}

func (k *KafkaHelper) Producer(record entities.Record) (bool, error) {

	recordJSON, err := json.Marshal(record)

	if err != nil {
		fmt.Printf("Error al serializar el registro: %v\n", err)
		return false, err
	}

	//* Creamos un nuevo productor
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "54.235.169.219:9092",
	})

	if err != nil {
		return false, err
	}
	defer p.Close()

	//* Canal para recibir el resultado de la entrega
	resultChan := make(chan bool)

	//* Goroutine para manejar los eventos de entrega
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					// Si la entrega falla
					fmt.Printf("Entrega Fallida: %v\n", ev.TopicPartition)
					resultChan <- false
				} else {
					// Si la entrega es exitosa
					fmt.Printf("Entrega Exitosa: %v\n", ev.TopicPartition)
					resultChan <- true
				}
			}
		}
	}()

	// Enviar el mensaje al tópico
	topic := "myTopic"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: recordJSON,
	}, nil)

	if err != nil {
		fmt.Printf("Error al enviar mensaje: %v\n", err)
		return false, err
	}

	// Asegurarse de que los mensajes se envíen antes de finalizar
	p.Flush(15 * 1000)

	// Esperar el resultado de la goroutine
	success := <-resultChan

	// Devolver el resultado (true o false)
	return success, nil
}
