// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEntityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEntityService] method
// instead.
type RadarEntityService struct {
	Options []option.RequestOption
	Asns    *RadarEntityAsnService
}

// NewRadarEntityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEntityService(opts ...option.RequestOption) (r *RadarEntityService) {
	r = &RadarEntityService{}
	r.Options = opts
	r.Asns = NewRadarEntityAsnService(opts...)
	return
}

// Get IP address information.
func (r *RadarEntityService) IPs(ctx context.Context, query RadarEntityIPsParams, opts ...option.RequestOption) (res *RadarEntityIPsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEntityIPsResponseEnvelope
	path := "radar/entities/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEntityIPsResponse struct {
	IP   RadarEntityIPsResponseIP   `json:"ip,required"`
	JSON radarEntityIPsResponseJSON `json:"-"`
}

// radarEntityIPsResponseJSON contains the JSON metadata for the struct
// [RadarEntityIPsResponse]
type radarEntityIPsResponseJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityIPsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityIPsResponseIP struct {
	Asn          string                       `json:"asn,required"`
	AsnLocation  string                       `json:"asnLocation,required"`
	AsnName      string                       `json:"asnName,required"`
	AsnOrgName   string                       `json:"asnOrgName,required"`
	IP           string                       `json:"ip,required"`
	IPVersion    string                       `json:"ipVersion,required"`
	Location     string                       `json:"location,required"`
	LocationName string                       `json:"locationName,required"`
	JSON         radarEntityIPsResponseIPJSON `json:"-"`
}

// radarEntityIPsResponseIPJSON contains the JSON metadata for the struct
// [RadarEntityIPsResponseIP]
type radarEntityIPsResponseIPJSON struct {
	Asn          apijson.Field
	AsnLocation  apijson.Field
	AsnName      apijson.Field
	AsnOrgName   apijson.Field
	IP           apijson.Field
	IPVersion    apijson.Field
	Location     apijson.Field
	LocationName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEntityIPsResponseIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEntityIPsParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required"`
	// Format results are returned in.
	Format param.Field[RadarEntityIPsParamsFormat] `query:"format"`
}

// URLQuery serializes [RadarEntityIPsParams]'s query parameters as `url.Values`.
func (r RadarEntityIPsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format results are returned in.
type RadarEntityIPsParamsFormat string

const (
	RadarEntityIPsParamsFormatJson RadarEntityIPsParamsFormat = "JSON"
	RadarEntityIPsParamsFormatCsv  RadarEntityIPsParamsFormat = "CSV"
)

type RadarEntityIPsResponseEnvelope struct {
	Result  RadarEntityIPsResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarEntityIPsResponseEnvelopeJSON `json:"-"`
}

// radarEntityIPsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarEntityIPsResponseEnvelope]
type radarEntityIPsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEntityIPsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}