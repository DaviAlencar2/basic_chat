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


	scannerInput := bufio.NewScanner(os.Stdin)
	

	go func () {
		serverScanner := bufio.NewScanner(serverConnection)
		for serverScanner.Scan() {
			text := serverScanner.Text()
			message, _ := models.FromJsonString(text)
			fmt.Println("\n" + "[" + message.SenderName + "]: " + message.Message)
		}
	}()


	for {
		if !scannerInput.Scan(){
			break
		}

		msgText := scannerInput.Text()

		message := models.Message{Message: msgText, SenderName: username}

		_, err := serverConnection.Write([]byte (message.ToJsonString()))

		if err != nil {
			fmt.Println("erro ao enviar mensagem")
			break
		}
	}
}
