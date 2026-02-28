package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go-kafka-consumer/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func main() {
	// 1. Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	topic := os.Getenv("KAFKA_TOPIC_NAME")
	broker := "localhost:9092"

	app := fiber.New()

	// 2. Initialize Kafka Reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  "course-consumer-group", // Keeps track of where we stopped reading
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	// 3. Background Consumer Loop
	go func() {
		fmt.Printf("Kafka Consumer started on topic: %s\n", topic)
		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %v", err)
				break
			}

			// Unmarshal JSON into our Course struct
			var course models.Course
			if err := json.Unmarshal(m.Value, &course); err != nil {
				log.Printf("Failed to parse JSON: %v", err)
				continue
			}

			fmt.Printf("Received Course: %+v\n", course)
		}
	}()

	// 4. Fiber API Routes
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "running",
			"topic":   topic,
			"broker":  broker,
		})
	})

	// Start Fiber on port 8082
	log.Fatal(app.Listen(":8082"))
}
