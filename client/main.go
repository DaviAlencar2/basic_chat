package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){

	serverConnection, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Erro ao se conectar no servidor")
		return
	}

	defer serverConnection.Close()

	go func() {
		serverScanner := bufio.NewScanner(serverConnection)

		for serverScanner.Scan(){
			fmt.Println(serverScanner.Text())
		}
	}()
	
	scannerInput := bufio.NewScanner(os.Stdin)
	fmt.Println("Sua mensagem:")

	for {
		if !scannerInput.Scan(){
			break
		}

		msgText := scannerInput.Text()

		_, err := serverConnection.Write([]byte (msgText + "\n"))

		if err != nil {
			fmt.Println("erro ao enviar mensagem")
			break
		}
	}
}
