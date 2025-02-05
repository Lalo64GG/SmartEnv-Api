package http

import (
	"log"

	"github.com/lalo64/SmartEnv-api/src/records/application"
	"github.com/lalo64/SmartEnv-api/src/records/application/services"
	"github.com/lalo64/SmartEnv-api/src/records/domain/ports"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/adapters"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/controllers"
	"github.com/lalo64/SmartEnv-api/src/records/infraestructure/http/controllers/helpers"

)

var (
	recordRepository ports.IRecordRepository
	KafkaService services.KafkaService
)

func init(){
	var err error
	
	recordRepository, err = adapters.NewRecordRepository()

	if err != nil {
		log.Fatalf("Error initializing record repository: %v", err)
	}

	// Inicializar el repositorio de kafka
	KafkaService, err = helpers.NewKafkaHelper()
	if err != nil {
		log.Fatalf("Error initializing kafka service: %v", err)
	}
}

func SetUpRegisterController() *controllers.CreateRecordController {
	createService := application.NewCreateRecordUseCase(recordRepository, KafkaService)
	return controllers.NewCreateRecordController(createService)
}

func GetAllRecordsController() *controllers.GetAllRecordsController{
	getAllService := application.NewGetAllRecordsUseCase(recordRepository)
    return controllers.NewGetAllRecordsController(getAllService)
}

func GetRecordByIDController() *controllers.GetByIdRecordController {
    getByIDService := application.NewGetRecordByIDUseCase(recordRepository)
    return controllers.NewGetByIdRecordController(getByIDService)
}