// Code generated by go-swagger; DO NOT EDIT.

package message

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/szemskov/herald/models"
)

// CreateCreatedCode is the HTTP code returned for type CreateCreated
const CreateCreatedCode int = 201

/*CreateCreated Message Created

swagger:response createCreated
*/
type CreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewCreateCreated creates CreateCreated with default headers values
func NewCreateCreated() *CreateCreated {

	return &CreateCreated{}
}

// WithPayload adds the payload to the create created response
func (o *CreateCreated) WithPayload(payload *models.Message) *CreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create created response
func (o *CreateCreated) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateBadRequestCode is the HTTP code returned for type CreateBadRequest
const CreateBadRequestCode int = 400

/*CreateBadRequest Bad Request

swagger:response createBadRequest
*/
type CreateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateBadRequest creates CreateBadRequest with default headers values
func NewCreateBadRequest() *CreateBadRequest {

	return &CreateBadRequest{}
}

// WithPayload adds the payload to the create bad request response
func (o *CreateBadRequest) WithPayload(payload *models.Error) *CreateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create bad request response
func (o *CreateBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateInternalServerErrorCode is the HTTP code returned for type CreateInternalServerError
const CreateInternalServerErrorCode int = 500

/*CreateInternalServerError Something went wrang

swagger:response createInternalServerError
*/
type CreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateInternalServerError creates CreateInternalServerError with default headers values
func NewCreateInternalServerError() *CreateInternalServerError {

	return &CreateInternalServerError{}
}

// WithPayload adds the payload to the create internal server error response
func (o *CreateInternalServerError) WithPayload(payload *models.Error) *CreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create internal server error response
func (o *CreateInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateDefault Unxpected Error

swagger:response createDefault
*/
type CreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDefault creates CreateDefault with default headers values
func NewCreateDefault(code int) *CreateDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create default response
func (o *CreateDefault) WithStatusCode(code int) *CreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create default response
func (o *CreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create default response
func (o *CreateDefault) WithPayload(payload *models.Error) *CreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create default response
func (o *CreateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
