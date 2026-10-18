package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"basic_chat/models"
)

func main(){

	serverConnection, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Erro ao se conectar no servidor")
		return
	}

	defer serverConnection.Close()

	fmt.Println("Escreva seu nome:")

	var username string
	userNameInput := bufio.NewScanner(os.Stdin)

	if userNameInput.Scan(){
		username = userNameInput.Text()
	}

	go func() {
		serverScanner := bufio.NewScanner(serverConnection)

		for serverScanner.Scan(){
			text := serverScanner.Text()
			msg, _ := models.FromJsonString(text)

			fmt.Printf("[%s]: %s", msg.SenderName, msg.Message)
		}
	}()
	
	scannerInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Sua mensagem:")

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
