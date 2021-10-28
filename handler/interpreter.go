package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/andresbelo12/GuiOS/model"
	//kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
)

type Interpreter struct {
	Connection *kernelModel.ClientConnection
	ListenerChannel chan kernelModel.Message
	WaitingResponse bool
}

func CreateInterpreter(connection *kernelModel.ClientConnection, channel chan kernelModel.Message) (interpreter Interpreter) {
	interpreter.Connection = connection
	interpreter.ListenerChannel = channel
	return
}

func (interpreter *Interpreter) ProcessCommand(input string) (response model.OperationResponse) {
	
	switch commandData := strings.Split(input, " "); commandData[0] {
	case "dc":
		response = interpreter.DirectoryCreate(commandData)
	case "dd":
		response = interpreter.DirectoryDelete(commandData)
	case "log":
		response = interpreter.ReadLogs(commandData)
	case "stop":
		if response = Stop(); response.Success{
			os.Exit(0)
		}
	default:
		response.Message = "Operation " + commandData[0] + " with arguments: " + strings.Join(commandData[1:], " ") + " not recognized. "
		return
	}
	return
}

func (interpreter *Interpreter) ProcessResponse() (success bool, message string) {
	fmt.Println("Waiting response...")
	interpreter.WaitingResponse = true
	serverResponse := <-interpreter.ListenerChannel
	interpreter.WaitingResponse = false

	if serverResponse.Command == kernelModel.CMD_STOP {
		return false, serverResponse.Message
	}

	messageBody := strings.Split(serverResponse.Message, ";")
	message = strings.Split(messageBody[2], ":")[1]
	if messageBody[0] == "response:true" {
		success = true
	}

	return
}


func Stop()(response model.OperationResponse){
	response.Success = true
	return
}