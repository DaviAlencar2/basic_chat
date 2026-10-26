package main

import (
	"basic_chat/models"
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Erro ao subir o servidor:" + err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Servidor rodando na 8080")

	salas := make(map[string]models.ChatGroup)
	salas["teste1"] = models.ChatGroup{Name: "teste1", Connections: []net.Conn{}}
	salas["teste2"] = models.ChatGroup{}

	fmt.Println(salas)
	
	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println("Erro ao aceitar conexao")
			return
		}

		go handleConnection(connection, salas)

	}
}

func handleConnection(connection net.Conn, salas map[string]models.ChatGroup) {
	defer connection.Close()
	fmt.Println("Novo cliente conectado")

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		clientMsg := scanner.Text()
		message, err := models.FromJsonString(clientMsg)

		if err != nil {
			return
		}

		fmt.Println("Mensagem Recebida:" + clientMsg)
		
		sala := salas[message.ChatGroupName]
		sala.Connections = append(sala.Connections, connection)
		salas[message.ChatGroupName] = sala

		for _, userConn := range sala.Connections {
			if userConn == connection{
				continue
			}
			_, err = userConn.Write([]byte(message.ToJsonString()))

		}

	}
}
