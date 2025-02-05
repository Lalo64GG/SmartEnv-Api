package adapters

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lalo64/SmartEnv-api/src/config"
	"github.com/lalo64/SmartEnv-api/src/records/domain/entities"
)

type RecordRepository struct {
	DB *sql.DB
}

func NewRecordRepository() (*RecordRepository, error) {
	db, err := config.Connect()
	if err != nil {
		return nil, err
	}
    return &RecordRepository{DB: db}, nil
}



func (r *RecordRepository) Create(record entities.Record) (entities.Record, error) {
	query := `INSERT INTO records (temperature, humidity, gas_level) VALUES (?, ?, ?)`

	stmt, err := r.DB.Prepare(query)

	if err!= nil {
        log.Fatal(err, 1)
	}
	defer stmt.Close()

	result, err := stmt.Exec(record.Temperature, record.Humidity, record.Gas_level)

	if err != nil {
		return entities.Record{}, err
	}

	id, err := result.LastInsertId()
	if err!= nil {
        return entities.Record{}, err
    }

	record.ID = id

	return record, nil

}


func (r *RecordRepository) GetAllRecords(limit, page int64, orderBy, orderDir string ) ([]entities.Record, error) {
	offset := limit * (page -1)

	query := fmt.Sprintf("SELECT id, temperature, humidity, gas_level FROM records ORDER BY %s %s LIMIT ? OFFSET ?", orderBy, orderDir)
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var records []entities.Record
	for rows.Next() {
		var record entities.Record
		err := rows.Scan(&record.ID, &record.Temperature, &record.Humidity, &record.Gas_level)
		if err != nil {
			return nil, err
		}

		records = append(records, record)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows error: %s", err)
		return nil, err
	}


	return records, nil
} 

func (r *RecordRepository) GetRecordByID(id int64) (entities.Record, error) {
	query := `SELECT id, temperature, humidity, gas_level FROM records WHERE id =?`

    stmt, err := r.DB.Prepare(query)

    if err!= nil {
        log.Fatal(err, 1)
    }
    defer stmt.Close()

    row := stmt.QueryRow(id)

    var record entities.Record
    err = row.Scan(&record.ID, &record.Temperature, &record.Humidity, &record.Gas_level)

    if err == sql.ErrNoRows {
        return entities.Record{}, fmt.Errorf("record not found")
    } else if err!= nil {
        return entities.Record{}, err
    }

    return record, nil
}
  