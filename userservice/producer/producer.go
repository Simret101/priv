package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"practice/rabbitmq" // Update with the actual import path

	"github.com/charmbracelet/lipgloss"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Media structure to hold media details.
type Media struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Path         string             `json:"path" bson:"path"`
	UploadedDate time.Time          `json:"uploaded_date" bson:"uploaded_date"`
	FileData     string             `json:"file_data" bson:"file_data"`
}

var (
	mediaStore sync.Map // In-memory store for media content

	// Define styles using Lipgloss
	infoStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	successStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	errorStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))
	linkStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Underline(true)
	// Style for the header with blue border and beautiful font
	headerStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("15")).
			Foreground(lipgloss.Color("14")). // Light cyan color for text
			Bold(true).
			Italic(true).
			Underline(true).
			Padding(2, 4).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("12")) // Blue border color
)

func main() {
	// Beautiful Heading with Lipgloss styling (Big Font + Blue Border)
	welcomeMessage := "WELCOME TO BLOG SERVICE"
	fmt.Println(headerStyle.Render(welcomeMessage))

	// Set up RabbitMQ connection and channel.
	conn, ch, err := rabbitmq.NewRabbitMQConnection()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer rabbitmq.CloseConnection(conn, ch)

	// Declare the queue.
	q, err := ch.QueueDeclare(
		"mediaQueue", // Queue name
		true,         // Durable
		false,        // Delete when unused
		false,        // Exclusive
		false,        // No-wait
		nil,          // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}

	// Specify the file path.
	filePath := "C:\\Users\\semre\\Desktop\\FinalTrail\\internal\\services\\assets\\uploads\\RO-nJiOm38HtFuN7_g_G3lmCR0m897RJNDj8QLpPhb3VdhG3Gi2M-wBJVRvxYCM-5heUHFspphaEr3_LWF5eIAYHfVX0UUBkNKfEF_73ZMD1vUw0tuYvVzpF77cK0EyE6qxwYfGDOt0e9_7bGyNIG2WExgfUpRT5FnbcQhp2NMPSrHVJTGkU6ZQ47QrLKvxeyccwkiSVYOQvyV_.png" // Replace with your actual file path.

	// Read the file's binary data.
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	// Encode the file data to Base64.
	encodedFileData := base64.StdEncoding.EncodeToString(fileData)

	// Create a Media instance with file data.
	media := Media{
		ID:           primitive.NewObjectID(),
		Path:         filePath,
		UploadedDate: time.Now(),
		FileData:     encodedFileData,
	}

	// Marshal the Media struct to JSON.
	body, err := json.Marshal(media)
	if err != nil {
		log.Fatalf("Failed to marshal message: %s", err)
	}

	// Publish the message to the queue.
	err = publishMessage(ch, q.Name, body)
	if err != nil {
		log.Fatalf("Failed to publish message: %s", err)
	}

	// Notify success with Lipgloss styled success message.
	fmt.Printf(successStyle.Render("Sent message with file: %s\n"), filePath)
}

// publishMessage publishes a message to the specified queue.
func publishMessage(ch *amqp.Channel, queueName string, message []byte) error {
	return ch.Publish(
		"",        // Exchange (default)
		queueName, // Routing key
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
}
