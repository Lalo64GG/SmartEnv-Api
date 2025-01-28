package application

import (
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/records/domain/ports"
)

type GetRecordByIDUseCase struct {
	RecordRepository ports.IRecordRepository
}

func NewGetRecordByIDUseCase(recordRepository ports.IRecordRepository) *GetRecordByIDUseCase{
	return &GetRecordByIDUseCase{RecordRepository: recordRepository} 
}


func (r *GetRecordByIDUseCase) Run(id int64) (entities.Record, error) {
	records, err := r.RecordRepository.GetRecordByID(id)

	if err != nil {
		return entities.Record{}, err
	}

	return records, nil 
}