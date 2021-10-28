package handler

import (
	"fmt"

	kernelModel "github.com/andresbelo12/KernelOS/model"
)

type ClientListener struct{}

func CreateListener() kernelModel.CommunicationListener {
	return ClientListener{}
}

func (listener ClientListener) ProcessMessage(processorTools interface{}, connection interface{}, message *kernelModel.Message) (err error) {
	fmt.Println(message)
	return
}
