// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// IPService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IPs used on the Cloudflare/JD Cloud network, see
// https://www.cloudflare.com/ips for Cloudflare IPs or
// https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD
// Cloud IPs.
func (r *IPService) List(ctx context.Context, query IPListParams, opts ...option.RequestOption) (res *IPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IPListResponseEnvelope
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [IPListResponseAddressingIPs] or
// [IPListResponseAddressingIPsJdcloud].
type IPListResponse interface {
	implementsIPListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*IPListResponse)(nil)).Elem(), "")
}

type IPListResponseAddressingIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6Cidrs []string                        `json:"ipv6_cidrs"`
	JSON      ipListResponseAddressingIPsJSON `json:"-"`
}

// ipListResponseAddressingIPsJSON contains the JSON metadata for the struct
// [IPListResponseAddressingIPs]
type ipListResponseAddressingIPsJSON struct {
	Etag        apijson.Field
	Ipv4Cidrs   apijson.Field
	IPV6Cidrs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseAddressingIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IPListResponseAddressingIPs) implementsIPListResponse() {}

type IPListResponseAddressingIPsJdcloud struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	Ipv4Cidrs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6Cidrs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JdcloudCidrs []string                               `json:"jdcloud_cidrs"`
	JSON         ipListResponseAddressingIPsJdcloudJSON `json:"-"`
}

// ipListResponseAddressingIPsJdcloudJSON contains the JSON metadata for the struct
// [IPListResponseAddressingIPsJdcloud]
type ipListResponseAddressingIPsJdcloudJSON struct {
	Etag         apijson.Field
	Ipv4Cidrs    apijson.Field
	IPV6Cidrs    apijson.Field
	JdcloudCidrs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPListResponseAddressingIPsJdcloud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IPListResponseAddressingIPsJdcloud) implementsIPListResponse() {}

type IPListParams struct {
	// Specified as `jdcloud` to list IPs used by JD Cloud data centers.
	Networks param.Field[string] `query:"networks"`
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IPListResponseEnvelope struct {
	Errors   []IPListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPListResponseEnvelopeMessages `json:"messages,required"`
	Result   IPListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IPListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipListResponseEnvelopeJSON    `json:"-"`
}

// ipListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPListResponseEnvelope]
type ipListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    ipListResponseEnvelopeErrorsJSON `json:"-"`
}

// ipListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeErrors]
type ipListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IPListResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    ipListResponseEnvelopeMessagesJSON `json:"-"`
}

// ipListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeMessages]
type ipListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IPListResponseEnvelopeSuccess bool

const (
	IPListResponseEnvelopeSuccessTrue IPListResponseEnvelopeSuccess = true
)