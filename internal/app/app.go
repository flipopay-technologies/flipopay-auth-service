package app

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func StartServer() {

	// err := godotenv.Load("/dev.env")
	// if err != nil {
	// 	log.Fatalf("Failed to load dev.env file: %v", err)
	// }

	url := "postgres://postgres:postgres@localhost:5432/auth-service-db?sslmode=disable"
	if url == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully...")

	// lis, err := net.Listen("tcp", ":9000")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// log.Println("Application started")

	// grpcServer := grpc.NewServer()

	// log.Println("Grpc server started")

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }

}
