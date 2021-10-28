package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/andresbelo12/GuiOS/model"
	kernelModel "github.com/andresbelo12/KernelOS/model"
	kernelHandler "github.com/andresbelo12/KernelOS/handler"
)

type Interpreter struct {
	Connection *kernelModel.ClientConnection
}

func CreateInterpreter(connection *kernelModel.ClientConnection) (interpreter Interpreter) {
	interpreter.Connection = connection
	return
}

func (interpreter Interpreter) ProcessCommand(input string) (response model.OperationResponse) {
	
	switch commandData := strings.Split(input, " "); commandData[0] {
	case "dc":
		response = interpreter.DirectoryCreate(commandData)
	case "dd":
		response = interpreter.DirectoryDelete(commandData)
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

func (interpreter Interpreter)DirectoryCreate(commandData []string) (response model.OperationResponse) {
	if len(commandData) < 2 {
		response.Message = "missing argument \"directory\"name. Example of usage (create logs directory): dc logs"
		return
	}

	message := kernelModel.Message{
		Command: kernelModel.CMD_SEND,
		Source: kernelModel.MD_GUI,
		Destination: kernelModel.MD_FILES,
		Message: "create:"+commandData[1],
	}

	if err := kernelHandler.WriteServer(interpreter.Connection, &message); err != nil{
		response.Message = err.Error()
		return
	}

	

	fmt.Println("Creating directory: " + commandData[1])
	response.Message = "directory " + commandData[1] + " created"
	response.Success = true
	return
}

func (interpreter Interpreter)DirectoryDelete(commandData []string) (response model.OperationResponse) {
	if len(commandData) < 2 {
		response.Message = "missing argument \"directory\"name. Example of usage (delete logs directory): dd logs"
		return
	}

	fmt.Println("Deleting directory: " + commandData[1])
	response.Message = "directory " + commandData[1] + " deleted"
	response.Success = true
	return
}

func Stop()(response model.OperationResponse){
	response.Success = true
	return
}