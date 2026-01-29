package server

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	Rwrite  http.ResponseWriter
	Request *http.Request
	Ctx     context.Context
	userId  uint
}

// responder con archivo plano al cliente
func (c *Context) Send(text string) {
	c.Rwrite.Write([]byte(text))
}

// obtener el codigo que envia server
func (c *Context) Status(code int) {
	c.Rwrite.WriteHeader(code)
}

// responder con json al cliente
func (c *Context) JSON(code int, data interface{}) error {
	c.Rwrite.Header().Set("Content-Type", "application/json")
	c.Rwrite.WriteHeader(code)

	return json.NewEncoder(c.Rwrite).Encode(data)
}
func (c *Context) BindJSON(dest interface{}) error {

	return json.NewDecoder(c.Request.Body).Decode(dest)
}

// obtener userId del contexto
func (c *Context) UserID(id uint) {
	c.userId = id
}
func (c *Context) GetUserID() uint {
	return c.userId
}

func (c *Context) Context() context.Context {
	return c.Ctx
}
