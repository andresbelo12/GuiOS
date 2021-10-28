package handler

import (
	"fmt"
	"os"

	"github.com/andresbelo12/GuiOS/model"
	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
)

const DefaultWorkspacePath = "/jalopezb/logs"

func (interpreter *Interpreter) ReadLogs(commandData []string) (response model.OperationResponse) {

	logFile := kernelModel.MD_FILES
	response.Message = "bad argument. Example of usage \n>>(read all logs): log \n>>(read log of gui): log gui\nPossible log files: gui, kernel, files"

	if len(commandData) == 1 {
		fmt.Println("Read all logs ")
	} else if len(commandData) == 2 {
		if commandData[1] == "gui" {
			logFile = kernelModel.MD_GUI
		} else if commandData[1] == "kernel" {
			logFile = kernelModel.MD_KERNEL
		} else if commandData[1] != "files" {
			return
		}
		fmt.Println("Read logs of module " + commandData[1])
	} else {
		return
	}

	message2 := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_GUI,
		Destination: kernelModel.MD_KERNEL,
		Message:     "lsdasdasdad:" + logFile,
	}

	if err := kernelHandler.WriteServer(interpreter.Connection, &message2); err != nil {
		response.Message = err.Error()
		return
	}

	message := kernelModel.Message{
		Command:     kernelModel.CMD_SEND,
		Source:      kernelModel.MD_GUI,
		Destination: kernelModel.MD_FILES,
		Message:     "log:" + logFile,
	}

	if err := kernelHandler.WriteServer(interpreter.Connection, &message); err != nil {
		response.Message = err.Error()
		return
	}

	success, result := interpreter.ProcessResponse()
	if success{
		if logs, err := readLogFile(result); err != nil{
			response.Message = err.Error()
		}else{
			response.Success = true
			response.Message = "Log printed"
			fmt.Println(logs)
		}
	}else{
		response.Message = result
	}

	return
}

func readLogFile(moduleFilePath string) (logs string, err error) {
	buffer := make([]byte, 10000)

	moduleFile, err := os.OpenFile(moduleFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	_, err = moduleFile.Read(buffer)
	logs = string(buffer)
	return
}
