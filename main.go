package main

import (
	"fmt"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	"github.com/andresbelo12/GuiOS/handler"
)

const (
	LOCALHOST   = "127.0.0.1"
	SERVER_PORT = "8080"
)

func main() {

	channel := make(chan kernelModel.Message)
	connection := connectToServer()
	interpreter := handler.CreateInterpreter(&connection, channel)
	listener := handler.CreateListener(&interpreter)

	go handler.LaunchGUI(&interpreter)
	kernelHandler.ListenServer(listener, &connection)

}

func connectToServer() (connection kernelModel.ClientConnection) {
	connection = kernelModel.ClientConnection{ServerHost: LOCALHOST, ServerPort: SERVER_PORT}
	firstMessage := kernelModel.Message{
		Source: kernelModel.MD_GUI, 
		Destination: kernelModel.MD_KERNEL, 
		Command: kernelModel.CMD_START, 
		Message: "listening",
	}

	if err := kernelHandler.EstablishClient(&connection, firstMessage); err != nil {
		fmt.Println(err)
		return
	}

	return
}
