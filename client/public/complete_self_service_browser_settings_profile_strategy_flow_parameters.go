// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParams creates a new CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams object
// with the default values initialized.
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParams() *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	var ()
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithTimeout creates a new CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	var ()
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams{

		timeout: timeout,
	}
}

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithContext creates a new CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithContext(ctx context.Context) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	var ()
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams{

		Context: ctx,
	}
}

// NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithHTTPClient creates a new CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteSelfServiceBrowserSettingsProfileStrategyFlowParamsWithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	var ()
	return &CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams{
		HTTPClient: client,
	}
}

/*CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams contains all the parameters to send to the API endpoint
for the complete self service browser settings profile strategy flow operation typically these are written to a http.Request
*/
type CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams struct {

	/*Body*/
	Body *models.CompleteSelfServiceBrowserSettingsStrategyProfileFlowPayload
	/*Request
	  Request is the request ID.

	*/
	Request string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WithTimeout(timeout time.Duration) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WithContext(ctx context.Context) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WithHTTPClient(client *http.Client) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WithBody(body *models.CompleteSelfServiceBrowserSettingsStrategyProfileFlowPayload) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) SetBody(body *models.CompleteSelfServiceBrowserSettingsStrategyProfileFlowPayload) {
	o.Body = body
}

// WithRequest adds the request to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WithRequest(request string) *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the complete self service browser settings profile strategy flow params
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) SetRequest(request string) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteSelfServiceBrowserSettingsProfileStrategyFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param request
	qrRequest := o.Request
	qRequest := qrRequest
	if qRequest != "" {
		if err := r.SetQueryParam("request", qRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
