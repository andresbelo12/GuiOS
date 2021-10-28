package handler

import (
	"log"
	"strings"

	kernelHandler "github.com/andresbelo12/KernelOS/handler"
	kernelModel "github.com/andresbelo12/KernelOS/model"
)

type ClientListener struct {
	Interpreter *Interpreter
}

func CreateListener(interpreter *Interpreter) kernelModel.CommunicationListener {
	return ClientListener{Interpreter: interpreter}
}

func (listener ClientListener) ProcessMessage(connection interface{}, message *kernelModel.Message) (err error) {
	conn := connection.(**kernelModel.ClientConnection)

	if listener.Interpreter.WaitingResponse {
		messageBody := strings.Split(message.Message, ";")
		errorMessage := kernelModel.Message{
			Command:     kernelModel.CMD_STOP,
			Source:      kernelModel.MD_GUI,
			Destination: kernelModel.MD_KERNEL,
			Message:     "error: response not expected",
		}

		if len(messageBody) != 3 {
			if err = kernelHandler.WriteServer(*conn, &errorMessage); err != nil {
				log.Fatal("Communication error")
			}
			listener.Interpreter.ListenerChannel <- errorMessage
			return
		}

		if !strings.Contains(messageBody[0], "response") {
			if err = kernelHandler.WriteServer(*conn, &errorMessage); err != nil {
				log.Fatal("Communication error")
			}
			listener.Interpreter.ListenerChannel <- errorMessage
			return
		}

		listener.Interpreter.ListenerChannel <- *message
		return
	}

	return
}
