// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/jgensler8/go-auth0/generated/models"
)

// GetUserinfoReader is a Reader for the GetUserinfo structure.
type GetUserinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserinfoOK creates a GetUserinfoOK with default headers values
func NewGetUserinfoOK() *GetUserinfoOK {
	return &GetUserinfoOK{}
}

/*GetUserinfoOK handles this case with default header values.

Consumer was created
*/
type GetUserinfoOK struct {
	Payload *models.UserInfo
}

func (o *GetUserinfoOK) Error() string {
	return fmt.Sprintf("[GET /userinfo][%d] getUserinfoOK  %+v", 200, o.Payload)
}

func (o *GetUserinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
