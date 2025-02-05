package kafka

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/gin-gonic/gin"
)

var (
	records     []entities.Record
	recordsLock sync.Mutex
)

func Consumer() {
	// Crear consumidor Kafka
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "54.235.169.219:9092",
		"group.id":          "test-consumer-group",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Error al crear el consumidor: %s\n", err)
		return
	}
	defer c.Close()

	// Suscribirse al tópico
	err = c.SubscribeTopics([]string{"myTopic"}, nil)
	if err != nil {
		fmt.Printf("Error al suscribirse al tópico: %s\n", err)
		return
	}

	for {
		// Leer mensaje
		msg, err := c.ReadMessage(5 * time.Second)
		if err != nil {
			if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.IsTimeout() {
				continue
			}
			fmt.Printf("Error al leer mensaje: %v\n", err)
			continue
		}

		var record entities.Record
		err = json.Unmarshal(msg.Value, &record)
		if err != nil {
			fmt.Printf("Error al deserializar mensaje: %s\n", err)
			continue
		}

		// Guardar el registro en memoria
		recordsLock.Lock()
		records = append(records, record)
		recordsLock.Unlock()

		fmt.Printf("Mensaje recibido en %s: %+v\n", msg.TopicPartition, record)
	}
}

// GetRecords devuelve el último registro consumido y responde en JSON
func GetRecords(ctx *gin.Context) {
	recordsLock.Lock()
	defer recordsLock.Unlock()

	// Si no hay registros aún
	if len(records) == 0 {
		ctx.JSON(404, gin.H{
			"success": false,
			"message": "No records found",
			"data":    nil,
		})
		return
	}

	// Obtener solo el último registro
	lastRecord := records[len(records)-1]

	// Responder con el último registro en formato JSON
	ctx.JSON(200, gin.H{
		"success": true,
		"message": "Last record retrieved successfully",
		"data":    lastRecord,
	})
}
