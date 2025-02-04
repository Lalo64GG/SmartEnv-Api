package application

import (
	"github.com/lalo64/SmartEnv-api/src/records/application/services"
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/records/domain/ports"
)

type CreateRecordUseCase struct {
	RecordRepository ports.IRecordRepository
	KafkaRepository services.KafkaService
}

func NewCreateRecordUseCase(recordRepositoty ports.IRecordRepository, kafkaRepository services.KafkaService)* CreateRecordUseCase{
	return &CreateRecordUseCase{RecordRepository: recordRepositoty, KafkaRepository: kafkaRepository}
}


func (r *CreateRecordUseCase) Run(Temperature, Humidity, Gas_level float64) (entities.Record, error){
	record := entities.Record{
		Temperature: Temperature,
		Humidity: Humidity,
		Gas_level: Gas_level,
	}

	status, err := r.KafkaRepository.Producer(record)

	if err != nil || !status {
		return entities.Record{}, err
	}

	newRecord, err := r.RecordRepository.Create(record)

	if err != nil {
		return entities.Record{}, err
	}

	return newRecord, nil
}