// Code generated by go-swagger; DO NOT EDIT.

package running_races

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetRunningRacesParams creates a new GetRunningRacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRunningRacesParams() *GetRunningRacesParams {
	return &GetRunningRacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunningRacesParamsWithTimeout creates a new GetRunningRacesParams object
// with the ability to set a timeout on a request.
func NewGetRunningRacesParamsWithTimeout(timeout time.Duration) *GetRunningRacesParams {
	return &GetRunningRacesParams{
		timeout: timeout,
	}
}

// NewGetRunningRacesParamsWithContext creates a new GetRunningRacesParams object
// with the ability to set a context for a request.
func NewGetRunningRacesParamsWithContext(ctx context.Context) *GetRunningRacesParams {
	return &GetRunningRacesParams{
		Context: ctx,
	}
}

// NewGetRunningRacesParamsWithHTTPClient creates a new GetRunningRacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRunningRacesParamsWithHTTPClient(client *http.Client) *GetRunningRacesParams {
	return &GetRunningRacesParams{
		HTTPClient: client,
	}
}

/* GetRunningRacesParams contains all the parameters to send to the API endpoint
   for the get running races operation.

   Typically these are written to a http.Request.
*/
type GetRunningRacesParams struct {

	/* Year.

	   Filters the list by a given year.
	*/
	Year *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get running races params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunningRacesParams) WithDefaults() *GetRunningRacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get running races params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRunningRacesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get running races params
func (o *GetRunningRacesParams) WithTimeout(timeout time.Duration) *GetRunningRacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get running races params
func (o *GetRunningRacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get running races params
func (o *GetRunningRacesParams) WithContext(ctx context.Context) *GetRunningRacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get running races params
func (o *GetRunningRacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get running races params
func (o *GetRunningRacesParams) WithHTTPClient(client *http.Client) *GetRunningRacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get running races params
func (o *GetRunningRacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithYear adds the year to the get running races params
func (o *GetRunningRacesParams) WithYear(year *int64) *GetRunningRacesParams {
	o.SetYear(year)
	return o
}

// SetYear adds the year to the get running races params
func (o *GetRunningRacesParams) SetYear(year *int64) {
	o.Year = year
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunningRacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Year != nil {

		// query param year
		var qrYear int64

		if o.Year != nil {
			qrYear = *o.Year
		}
		qYear := swag.FormatInt64(qrYear)
		if qYear != "" {

			if err := r.SetQueryParam("year", qYear); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
