package application

import (
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
	"github.com/lalo64/SmartEnv-api/src/records/domain/ports"
)

type GetAllRecordsUseCase struct {
	RecordRepository ports.IRecordRepository
}


func NewGetAllRecordsUseCase(recordRepository ports.IRecordRepository) *GetAllRecordsUseCase {
	return &GetAllRecordsUseCase{ RecordRepository: recordRepository }
}


func (r *GetAllRecordsUseCase) Run(limit, page int64, orderBy, orderDir string) (entities.Record, error) {
	records, err := r.RecordRepository.GetAllRecords(limit, page, orderBy, orderDir)

    if err != nil {
        return entities.Record{}, err
    }

    return records, nil
}