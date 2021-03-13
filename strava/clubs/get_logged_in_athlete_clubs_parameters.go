// Code generated by go-swagger; DO NOT EDIT.

package clubs

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

// NewGetLoggedInAthleteClubsParams creates a new GetLoggedInAthleteClubsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoggedInAthleteClubsParams() *GetLoggedInAthleteClubsParams {
	return &GetLoggedInAthleteClubsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoggedInAthleteClubsParamsWithTimeout creates a new GetLoggedInAthleteClubsParams object
// with the ability to set a timeout on a request.
func NewGetLoggedInAthleteClubsParamsWithTimeout(timeout time.Duration) *GetLoggedInAthleteClubsParams {
	return &GetLoggedInAthleteClubsParams{
		timeout: timeout,
	}
}

// NewGetLoggedInAthleteClubsParamsWithContext creates a new GetLoggedInAthleteClubsParams object
// with the ability to set a context for a request.
func NewGetLoggedInAthleteClubsParamsWithContext(ctx context.Context) *GetLoggedInAthleteClubsParams {
	return &GetLoggedInAthleteClubsParams{
		Context: ctx,
	}
}

// NewGetLoggedInAthleteClubsParamsWithHTTPClient creates a new GetLoggedInAthleteClubsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoggedInAthleteClubsParamsWithHTTPClient(client *http.Client) *GetLoggedInAthleteClubsParams {
	return &GetLoggedInAthleteClubsParams{
		HTTPClient: client,
	}
}

/* GetLoggedInAthleteClubsParams contains all the parameters to send to the API endpoint
   for the get logged in athlete clubs operation.

   Typically these are written to a http.Request.
*/
type GetLoggedInAthleteClubsParams struct {

	/* Page.

	   Page number. Defaults to 1.
	*/
	Page *int64

	/* PerPage.

	   Number of items per page. Defaults to 30.

	   Default: 30
	*/
	PerPage *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logged in athlete clubs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteClubsParams) WithDefaults() *GetLoggedInAthleteClubsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logged in athlete clubs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoggedInAthleteClubsParams) SetDefaults() {
	var (
		perPageDefault = int64(30)
	)

	val := GetLoggedInAthleteClubsParams{
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) WithTimeout(timeout time.Duration) *GetLoggedInAthleteClubsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) WithContext(ctx context.Context) *GetLoggedInAthleteClubsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) WithHTTPClient(client *http.Client) *GetLoggedInAthleteClubsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) WithPage(page *int64) *GetLoggedInAthleteClubsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) WithPerPage(perPage *int64) *GetLoggedInAthleteClubsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get logged in athlete clubs params
func (o *GetLoggedInAthleteClubsParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInAthleteClubsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
