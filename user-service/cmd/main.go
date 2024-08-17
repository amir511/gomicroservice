package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func main() {
	// Load environment variables
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	if cassandraHost == "" {
		log.Fatal("CASSANDRA_HOST environment variable not set")
	}

	// Set up Cassandra cluster configuration
	cluster := gocql.NewCluster(cassandraHost)
	cluster.Keyspace = "system" // For health check; switch to actual keyspace for operations
	cluster.Consistency = gocql.Quorum

	// Health check logic
	var session *gocql.Session
	var err error
	for {
		session, err = cluster.CreateSession()
		if err != nil {
			log.Printf("Unable to connect to Cassandra: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	defer session.Close()
	log.Println("Connected to Cassandra successfully")

	// Set up Gin router
	r := gin.Default()

	// Define a simple hello world endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
