// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// CompleteSelfServiceBrowserSettingsProfileStrategyFlowReader is a Reader for the CompleteSelfServiceBrowserSettingsProfileStrategyFlow structure.
type CompleteSelfServiceBrowserSettingsProfileStrategyFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowFound creates a CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound with default headers values
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowFound() *CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound {
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound{}
}

/*CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound struct {
}

func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/settings/strategies/profile][%d] completeSelfServiceBrowserSettingsProfileStrategyFlowFound ", 302)
}

func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError creates a CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError with default headers values
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError() *CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError {
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError{}
}

/*CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /self-service/browser/flows/settings/strategies/profile][%d] completeSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
