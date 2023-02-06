// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/muly/howto/golang/code-gen/swagger/hello-swagger/gen/models"
)

// GetStudentByIDOKCode is the HTTP code returned for type GetStudentByIDOK
const GetStudentByIDOKCode int = 200

/*
GetStudentByIDOK returns the details of a student

swagger:response getStudentByIdOK
*/
type GetStudentByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.StudentView `json:"body,omitempty"`
}

// NewGetStudentByIDOK creates GetStudentByIDOK with default headers values
func NewGetStudentByIDOK() *GetStudentByIDOK {

	return &GetStudentByIDOK{}
}

// WithPayload adds the payload to the get student by Id o k response
func (o *GetStudentByIDOK) WithPayload(payload *models.StudentView) *GetStudentByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get student by Id o k response
func (o *GetStudentByIDOK) SetPayload(payload *models.StudentView) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStudentByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
