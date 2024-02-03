package assistance

import (
	"errors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Transaction struct {
	ID     string
	Name   string
	Type   string
	Amount float64
	Date   time.Time
}

func GenerateTransaction() Transaction {
	types := []string{"покупка", "перевод", "списание"}
	randomType := types[rand.Intn(len(types))]

	amount := rand.Float64() * 1000
	date := time.Now()
	userID := uuid.New()
	name := "Transaction_" + userID.String()

	return Transaction{ID: userID.String(), Name: name, Type: randomType, Amount: amount, Date: date}
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rootDir, err := findRootDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		log.Fatal("Error loading .env file: " + err.Error())
	}
}

func findRootDir(dir string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("go.mod not found")
		}
		dir = parent
	}
}
