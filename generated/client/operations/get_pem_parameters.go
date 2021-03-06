// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPEMParams creates a new GetPEMParams object
// with the default values initialized.
func NewGetPEMParams() *GetPEMParams {

	return &GetPEMParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPEMParamsWithTimeout creates a new GetPEMParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPEMParamsWithTimeout(timeout time.Duration) *GetPEMParams {

	return &GetPEMParams{

		timeout: timeout,
	}
}

// NewGetPEMParamsWithContext creates a new GetPEMParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPEMParamsWithContext(ctx context.Context) *GetPEMParams {

	return &GetPEMParams{

		Context: ctx,
	}
}

// NewGetPEMParamsWithHTTPClient creates a new GetPEMParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPEMParamsWithHTTPClient(client *http.Client) *GetPEMParams {

	return &GetPEMParams{
		HTTPClient: client,
	}
}

/*GetPEMParams contains all the parameters to send to the API endpoint
for the get p e m operation typically these are written to a http.Request
*/
type GetPEMParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get p e m params
func (o *GetPEMParams) WithTimeout(timeout time.Duration) *GetPEMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get p e m params
func (o *GetPEMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get p e m params
func (o *GetPEMParams) WithContext(ctx context.Context) *GetPEMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get p e m params
func (o *GetPEMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get p e m params
func (o *GetPEMParams) WithHTTPClient(client *http.Client) *GetPEMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get p e m params
func (o *GetPEMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPEMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
