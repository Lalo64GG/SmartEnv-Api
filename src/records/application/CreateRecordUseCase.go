package application

import (
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/records/domain/ports"
)

type CreateRecordUseCase struct {
	RecordRepository ports.IRecordRepository
}

func NewCreateRecordUseCase(recordRepositoty ports.IRecordRepository)* CreateRecordUseCase{
	return &CreateRecordUseCase{RecordRepository: recordRepositoty}
}


func (r *CreateRecordUseCase) Run(Temperature, Distance float64) (entities.Record, error){
	record := entities.Record{
		Temperature: Temperature,
		Distance: Distance,
	}

	newRecord, err := r.RecordRepository.Create(record)

	if err != nil {
		return entities.Record{}, err
	}

	return newRecord, nil
}