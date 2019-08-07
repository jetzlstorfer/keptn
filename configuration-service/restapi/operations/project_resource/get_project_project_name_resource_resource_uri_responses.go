// Code generated by go-swagger; DO NOT EDIT.

package project_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/keptn/keptn/configuration-service/models"
)

// GetProjectProjectNameResourceResourceURIOKCode is the HTTP code returned for type GetProjectProjectNameResourceResourceURIOK
const GetProjectProjectNameResourceResourceURIOKCode int = 200

/*GetProjectProjectNameResourceResourceURIOK Success

swagger:response getProjectProjectNameResourceResourceUriOK
*/
type GetProjectProjectNameResourceResourceURIOK struct {

	/*
	  In: Body
	*/
	Payload *models.Resource `json:"body,omitempty"`
}

// NewGetProjectProjectNameResourceResourceURIOK creates GetProjectProjectNameResourceResourceURIOK with default headers values
func NewGetProjectProjectNameResourceResourceURIOK() *GetProjectProjectNameResourceResourceURIOK {

	return &GetProjectProjectNameResourceResourceURIOK{}
}

// WithPayload adds the payload to the get project project name resource resource Uri o k response
func (o *GetProjectProjectNameResourceResourceURIOK) WithPayload(payload *models.Resource) *GetProjectProjectNameResourceResourceURIOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get project project name resource resource Uri o k response
func (o *GetProjectProjectNameResourceResourceURIOK) SetPayload(payload *models.Resource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProjectProjectNameResourceResourceURIOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
