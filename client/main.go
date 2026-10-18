package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"basic_chat/models"
	"github.com/chzyer/readline"
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


	terminal, err := readline.New(">")
	if err != nil {
		fmt.Println("Erro ao preparar o terminal:" + err.Error())
		return
	}
	defer terminal.Close()
	

	go func () {
		serverScanner := bufio.NewScanner(serverConnection)
		for serverScanner.Scan() {
			text := serverScanner.Text()
			message, _ := models.FromJsonString(text)
			fmt.Fprintf(
				terminal.Stdout(),
				"[%s]: %s\n",
				message.SenderName,
				message.Message,
			)
		}
	}()


	for {
	
		msgText, err  := terminal.Readline()

		if err != nil {
			break
		}

		message := models.Message{Message: msgText, SenderName: username}

		_, err = serverConnection.Write([]byte (message.ToJsonString()))

		if err != nil {
			fmt.Fprintln(terminal.Stdout(), "Erro ao enviar mensagem")
			break
		}
	}
}
