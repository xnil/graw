// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package handlers

import "github.com/turnage/graw/internal/data"

// MessageHandler processes Reddit messages.
type MessageHandler interface {
	// HandleMessage processes a Reddit message.
	HandleMessage(p *data.Message) error
}

type messageHandler struct {
	handler func(p *data.Message) error
}

func (h *messageHandler) HandleMessage(message *data.Message) error {
	return h.handler(message)
}

// MessageHandlerFunc returns a MessageHandler using the given function.
func MessageHandlerFunc(f func(e *data.Message) error) MessageHandler {
	return &messageHandler{f}
}