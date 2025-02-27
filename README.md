# SmartEnv-API

SmartEnv-API is a Go-based backend API designed for managing IoT devices like LEDs and temperature sensors. It features an event-driven architecture powered by Kafka and uses a hexagonal architecture to ensure scalability and maintainability. This API handles user authentication, registration, and smart environment data management.

## Features

- **User Management**: Create, register, and authenticate users.
- **IoT Device Management**: Interact with devices like LEDs and temperature sensors.
- **Event-Driven Architecture**: Utilizes Kafka for real-time event processing.
- **Hexagonal Architecture**: Ensures modularity, scalability, and maintainability.
- **Database Integration**: Supports MySQL for persistent data storage.

## Setup

1. Clone the repository:
   git clone https://github.com/Lalo64GG/SmartEnv-Api.git

2. Install dependencies:
   Ensure you have Go installed and configured. Then run:
   go mod tidy

3. Configure environment variables:
   Create a .env file and define the following variables:
   DB_USER=your_database_user
   DB_PASSWORD=your_database_password
   DB_HOST=your_database_host
   DB_PORT=your_database_port
   DB_NAME=your_database_name
   KAFKA_BROKER=your_kafka_broker

4. Run the application:
   go run main.go

5. To run migrations, execute:
   go run ./config/migrations/Migrations.go

## API Endpoints

- POST /register: Register a new user.
- POST /login: Authenticate and log in a user.
- GET /home: Retrieve user-specific data and IoT device status.

## Technologies

- **Go**: The backend is developed in Go.
- **Kafka**: For event-driven communication.
- **MySQL**: For database management.
- **Gin**: Go web framework for API routes.
- **Bcrypt**: For secure password hashing.

## Contributing

1. Fork the repository.
2. Create a new branch (git checkout -b feature-branch).
3. Make your changes and commit them (git commit -am 'Add new feature').
4. Push to the branch (git push origin feature-branch).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
