package mock

import (
	"encoding/json"
	"example/vtr-mailer-service/structs"
	"example/vtr-mailer-service/tools"
	"fmt"
	"os"
)

func main() {
	// Caminho para o arquivo exemplo.json no diretório mock
	filepath := "./mock/create-email-request.json"

	// Lendo o arquivo JSON
	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parseando o conteúdo do JSON para a estrutura Email
	var email structs.TransactionCreateRequest
	err = json.Unmarshal(byteValue, &email)
	if err != nil {
		fmt.Println("Error parsing JSON to struct:", err)
		return
	}

	// Imprimindo o resultado para verificar
	tools.PrintStruct(email)
}
