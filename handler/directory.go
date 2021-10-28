package handler

import (
	"fmt"

	"github.com/andresbelo12/GuiOS/model"
	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
)

func (interpreter *Interpreter) DirectoryCreate(commandData []string) (response model.OperationResponse) {
	if len(commandData) < 2 {
		response.Message = "missing argument \"directory\"name. Example of usage (create sample directory): dc sample"
		return
	}
	fmt.Println("Create directory: " + commandData[1])

	message := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_GUI,
		Destination: kernelModel.MD_FILES,
		Message:     "create:" + commandData[1],
	}

	if err := kernelHandler.WriteServer(interpreter.Connection, &message); err != nil {
		response.Message = err.Error()
		return
	}

	response.Success, response.Message = interpreter.ProcessResponse()

	return
}

func (interpreter *Interpreter) DirectoryDelete(commandData []string) (response model.OperationResponse) {	
	if len(commandData) < 2 {
		response.Message = "missing argument \"directory\"name. Example of usage (delete sample directory): dd sample"
		return
	}
	fmt.Println("Delete directory: " + commandData[1])

	message := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_GUI,
		Destination: kernelModel.MD_FILES,
		Message:     "delete:" + commandData[1],
	}

	if err := kernelHandler.WriteServer(interpreter.Connection, &message); err != nil {
		response.Message = err.Error()
		return
	}

	response.Success, response.Message = interpreter.ProcessResponse()

	return
}
