package assistance

import (
	"errors"
	"github.com/TechSir3n/analytics-platform/assistance/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GenerateTransaction() models.Transaction {
	types := []string{"покупка", "перевод", "списание"}
	randomType := types[rand.Intn(len(types))]

	amount := rand.Float64() * 100000000
	date := time.Now()
	userID := uuid.New()
	name := "Transaction_" + userID.String()

	return models.Transaction{ID: userID.String(), Name: name, Type: randomType, Amount: amount, Date: date}
}

func GenerateProduct() models.Product {
	productID := uuid.New()
	nameProduct := "Product_" + productID.String()
	price := rand.Float64() * 1000000
	quantity := rand.Intn(10)

	return models.Product{ID: productID.String(), Name: nameProduct, Price: price, Quantity: quantity}
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
