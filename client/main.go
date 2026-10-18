package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"basic_chat/models"
)

func main(){
	fmt.Println("Bem vindo!")
	fmt.Print("Escreva seu nome:")
	var username string
	userNameInput := bufio.NewScanner(os.Stdin)
	if userNameInput.Scan(){
		username = userNameInput.Text()
	}

	serverConnection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao se conectar no servidor")
		return
	}
	

	defer serverConnection.Close()


	fmt.Print("Escreva algo:")
	scannerInput := bufio.NewScanner(os.Stdin)
	

	for {
		if !scannerInput.Scan(){
			break
		}

		fmt.Print("Escreva algo:")
		msgText := scannerInput.Text()

		message := models.Message{Message: msgText, SenderName: username}

		_, err := serverConnection.Write([]byte (message.ToJsonString()))

		if err != nil {
			fmt.Println("erro ao enviar mensagem")
			break
		}
	}
}
