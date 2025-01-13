package response

import (
	"net/http"

	"gg/ptr"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Error   *string     `json:"error"`
	Message *string     `json:"message"`
	Status  int         `json:"status"`
}

var (
	DefaultResponse = &Response{
		Data:    map[string]interface{}{},
		Error:   ptr.Ref(""),
		Message: ptr.Ref(""),
		Status:  http.StatusOK,
	}
)

func With(c *gin.Context, status int, options ...Option) {
	response := &Response{
		Status: status}

	for _, option := range options {
		option(response)
	}

	c.JSON(status, response)
}

func WithData(c *gin.Context, status int, data interface{}, options ...Option) {
	With(c, status, append(options, Data(data))...)
}

func WithError(c *gin.Context, status int, err string, options ...Option) {
	With(c, status, append(options, Error(err))...)
}

type Option func(*Response) error

func Data(v interface{}) Option {
	return func(r *Response) error {
		r.Data = v
		return nil
	}
}

func Err(err error) Option {
	return func(r *Response) error {
		err_string := err.Error()
		r.Error = &err_string
		return nil
	}
}

func Error(err string) Option {
	return func(r *Response) error {
		r.Error = &err
		return nil
	}
}

func Message(message string) Option {
	return func(r *Response) error {
		r.Message = &message
		return nil
	}
}
