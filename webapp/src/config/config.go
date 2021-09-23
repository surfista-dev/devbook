package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//APIURL representa a URL de comunicação com API
	APIURL = ""
	// Porta aonde a aplicaão web esta rodando
	Porta = 0
	// Hashkey e utilizada para autenticar o cookie
	Hashkey []byte
	// Blockkey e utilizado para criptografar o cookie
	Blockkey []byte
)

// Carregar inicializa as variveis de anbiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	Hashkey = []byte(os.Getenv("HASH_KEY"))
	Blockkey = []byte(os.Getenv("BLOCK_KEY"))
}
