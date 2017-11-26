// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "goa Practice": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/imyslx/go-practice/goa/design
// --out=$(GOPATH)/src/github.com/imyslx/go-practice/goa
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Practice for goa API. (default view)
//
// Identifier: application/vnd.result+json; view=default
type Result struct {
	// message
	Message string `form:"message" json:"message" xml:"message"`
	// status
	Status int `form:"status" json:"status" xml:"status"`
}

// Validate validates the Result media type instance.
func (mt *Result) Validate() (err error) {

	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// DecodeResult decodes the Result instance encoded in resp body.
func (c *Client) DecodeResult(resp *http.Response) (*Result, error) {
	var decoded Result
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}