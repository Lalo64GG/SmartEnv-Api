package services

import "github.com/lalo64/SmartEnv-api/src/records/domain/entities"

type Ikafka interface {
	Producer(record entities.Record) (bool, error)
}