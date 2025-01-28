package ports

import (
    "github.com/lalo64/SmartEnv-api/src/records/domain/entities"
)

type IRecordRepository interface {
	Create(record entities.Record) (entities.Record, error)
	GetAllRecords(limit, page int64, orderBy, orderDir string ) ([]entities.Record, error)
	GetRecordByID(id int64) (entities.Record, error)
}