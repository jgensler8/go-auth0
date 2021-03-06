// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetUserinfo Create a consumer
*/
func (a *Client) GetUserinfo(params *GetUserinfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetUserinfo",
		Method:             "GET",
		PathPattern:        "/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserinfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserinfoOK), nil

}

/*
GetPEM get p e m API
*/
func (a *Client) GetPEM(params *GetPEMParams, writer io.Writer) (*GetPEMOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPEMParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPEM",
		Method:             "GET",
		PathPattern:        "/pem",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPEMReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPEMOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
