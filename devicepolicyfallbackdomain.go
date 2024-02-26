// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DevicePolicyFallbackDomainService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDevicePolicyFallbackDomainService] method instead.
type DevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyFallbackDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyFallbackDomainService) {
	r = &DevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyFallbackDomainService) Update(ctx context.Context, accountID interface{}, policyID string, body DevicePolicyFallbackDomainUpdateParams, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyFallbackDomainService) List(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/fallback_domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyFallbackDomainService) Get(ctx context.Context, accountID interface{}, policyID string, opts ...option.RequestOption) (res *[]DevicePolicyFallbackDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/fallback_domains", accountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyFallbackDomainUpdateResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                                `json:"dns_server"`
	JSON      devicePolicyFallbackDomainUpdateResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseJSON contains the JSON metadata for the
// struct [DevicePolicyFallbackDomainUpdateResponse]
type devicePolicyFallbackDomainUpdateResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                              `json:"dns_server"`
	JSON      devicePolicyFallbackDomainListResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseJSON contains the JSON metadata for the
// struct [DevicePolicyFallbackDomainListResponse]
type devicePolicyFallbackDomainListResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainGetResponse struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}                             `json:"dns_server"`
	JSON      devicePolicyFallbackDomainGetResponseJSON `json:"-"`
}

// devicePolicyFallbackDomainGetResponseJSON contains the JSON metadata for the
// struct [DevicePolicyFallbackDomainGetResponse]
type devicePolicyFallbackDomainGetResponseJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainUpdateParams struct {
	Body param.Field[[]DevicePolicyFallbackDomainUpdateParamsBody] `json:"body,required"`
}

func (r DevicePolicyFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyFallbackDomainUpdateParamsBody struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r DevicePolicyFallbackDomainUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyFallbackDomainUpdateResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyFallbackDomainUpdateResponseEnvelope]
type devicePolicyFallbackDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    devicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainUpdateResponseEnvelopeErrors]
type devicePolicyFallbackDomainUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    devicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainUpdateResponseEnvelopeMessages]
type devicePolicyFallbackDomainUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainListResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyFallbackDomainListResponseEnvelope]
type devicePolicyFallbackDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainListResponseEnvelopeErrors]
type devicePolicyFallbackDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainListResponseEnvelopeMessages]
type devicePolicyFallbackDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainListResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainListResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainListResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainListResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainGetResponseEnvelope struct {
	Errors   []DevicePolicyFallbackDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DevicePolicyFallbackDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []DevicePolicyFallbackDomainGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyFallbackDomainGetResponseEnvelope]
type devicePolicyFallbackDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    devicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainGetResponseEnvelopeErrors]
type devicePolicyFallbackDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DevicePolicyFallbackDomainGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    devicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DevicePolicyFallbackDomainGetResponseEnvelopeMessages]
type devicePolicyFallbackDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainGetResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainGetResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainGetResponseEnvelopeSuccess = true
)

type DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
