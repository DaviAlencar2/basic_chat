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
		fmt.Println("Erro ao subir o servidor:" + err.Error())
		return
	}

	defer listener.Close()

	chatGroup := &models.ChatGroup{
		Connections:[]net.Conn{},
	}

	fmt.Println("Servidor rodando na 8080")

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Erro ao aceitar conexao")
			return
		}

		chatGroup.Connections = append(chatGroup.Connections, connection)
		go handleConnection(connection, chatGroup)

	}
}


func handleConnection(connection net.Conn, chat *models.ChatGroup){
	defer connection.Close()

	fmt.Println("Novo cliente conectado")

	scanner := bufio.NewScanner(connection)

	for scanner.Scan(){
		clientMsg := scanner.Text()
		message, err := models.FromJsonString(clientMsg)

		if err != nil{
			return
		}
	
		fmt.Println("Mensagem Recebida:" + clientMsg)

		for _, userConn := range chat.Connections {
			if userConn == connection {
				continue
			}
			
			_, err = userConn.Write([]byte(message.ToJsonString()))
		}
	}
}
