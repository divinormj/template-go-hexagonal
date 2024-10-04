package entities

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Exporta o cliente MongoDB para ser usado em outros testes
var MongoClient *mongo.Client

func TestMain(m *testing.M) {
	// Inicia o contêiner do MongoDB uma vez antes de todos os testes
	cmd := exec.Command("docker-compose", "-f", "../../docker-compose.test.yml", "up", "-d")
	if err := cmd.Run(); err != nil {
		panic("Erro ao iniciar o contêiner MongoDB: " + err.Error())
	}

	// Conectar ao MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27071")
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("Erro ao conectar ao MongoDB: " + err.Error())
	}

	// Executa os testes
	code := m.Run()

	// Para e remove o contêiner após todos os testes
	if err := exec.Command("docker-compose", "-f", "../../docker-compose.test.yml", "down").Run(); err != nil {
		panic("Erro ao parar o contêiner MongoDB: " + err.Error())
	}

	os.Exit(code)
}

// CleanupDatabase limpa o banco de dados após os testes.
func CleanupDatabase(t *testing.T) {
	// Limpa o banco de dados
	if err := MongoClient.Database("test_db").Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
