package main

import (
	"bufio"
	"fmt"
	"net"
	"basic_chat/models"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Erro ao subir o servidor!")
		return
	}

	defer listener.Close()

	fmt.Println("Servidor rodando na 8080")

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Erro ao aceitar conexao")
			return
		}

		go handleConnection(connection)

	}
}


func handleConnection(connection net.Conn){
	defer connection.Close()

	fmt.Println("Novo cliente conectado")

	scanner := bufio.NewScanner(connection)

	for scanner.Scan(){
		clientMsg := scanner.Text()
		msg, err :=  models.FromJsonString(clientMsg)

		fmt.Printf("Mensagem recebida do user [%s] - Mensagem [%s]", msg.SenderName, msg.Message)

		if err != nil {
			return
		}

		fmt.Println("Mensagem Recebida:" + clientMsg)

		response := fmt.Sprintf("Servidor: Mensagem %s recebida com sucesso", clientMsg)

		_, err = connection.Write([]byte(msg.ToJsonString()))

		if err != nil {
			fmt.Println("Erro ao enviar mensagem ao cliente")
			break
		}
	}
}