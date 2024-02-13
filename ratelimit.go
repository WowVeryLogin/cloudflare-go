// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RateLimitService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRateLimitService] method instead.
type RateLimitService struct {
	Options []option.RequestOption
}

// NewRateLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRateLimitService(opts ...option.RequestOption) (r *RateLimitService) {
	r = &RateLimitService{}
	r.Options = opts
	return
}

// Updates an existing rate limit.
func (r *RateLimitService) Update(ctx context.Context, zoneIdentifier string, id string, body RateLimitUpdateParams, opts ...option.RequestOption) (res *RateLimitUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the rate limits for a zone.
func (r *RateLimitService) List(ctx context.Context, zoneIdentifier string, query RateLimitListParams, opts ...option.RequestOption) (res *[]RateLimitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitListResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a rate limit.
func (r *RateLimitService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *RateLimitGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RateLimitGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [RateLimitUpdateResponseUnknown] or [shared.UnionString].
type RateLimitUpdateResponse interface {
	ImplementsRateLimitUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitListResponse struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RateLimitListResponseAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RateLimitListResponseBypass `json:"bypass"`
	// An informative summary of the rate limit. This value is sanitized and any tags
	// will be removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RateLimitListResponseMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                   `json:"threshold"`
	JSON      rateLimitListResponseJSON `json:"-"`
}

// rateLimitListResponseJSON contains the JSON metadata for the struct
// [RateLimitListResponse]
type rateLimitListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitListResponseAction struct {
	// The action to perform.
	Mode RateLimitListResponseActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RateLimitListResponseActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                         `json:"timeout"`
	JSON    rateLimitListResponseActionJSON `json:"-"`
}

// rateLimitListResponseActionJSON contains the JSON metadata for the struct
// [RateLimitListResponseAction]
type rateLimitListResponseActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to perform.
type RateLimitListResponseActionMode string

const (
	RateLimitListResponseActionModeSimulate         RateLimitListResponseActionMode = "simulate"
	RateLimitListResponseActionModeBan              RateLimitListResponseActionMode = "ban"
	RateLimitListResponseActionModeChallenge        RateLimitListResponseActionMode = "challenge"
	RateLimitListResponseActionModeJsChallenge      RateLimitListResponseActionMode = "js_challenge"
	RateLimitListResponseActionModeManagedChallenge RateLimitListResponseActionMode = "managed_challenge"
)

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitListResponseActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                  `json:"content_type"`
	JSON        rateLimitListResponseActionResponseJSON `json:"-"`
}

// rateLimitListResponseActionResponseJSON contains the JSON metadata for the
// struct [RateLimitListResponseActionResponse]
type rateLimitListResponseActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitListResponseBypass struct {
	Name RateLimitListResponseBypassName `json:"name"`
	// The URL to bypass.
	Value string                          `json:"value"`
	JSON  rateLimitListResponseBypassJSON `json:"-"`
}

// rateLimitListResponseBypassJSON contains the JSON metadata for the struct
// [RateLimitListResponseBypass]
type rateLimitListResponseBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitListResponseBypassName string

const (
	RateLimitListResponseBypassNameURL RateLimitListResponseBypassName = "url"
)

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitListResponseMatch struct {
	Headers  []RateLimitListResponseMatchHeader `json:"headers"`
	Request  RateLimitListResponseMatchRequest  `json:"request"`
	Response RateLimitListResponseMatchResponse `json:"response"`
	JSON     rateLimitListResponseMatchJSON     `json:"-"`
}

// rateLimitListResponseMatchJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatch]
type rateLimitListResponseMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitListResponseMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RateLimitListResponseMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                               `json:"value"`
	JSON  rateLimitListResponseMatchHeaderJSON `json:"-"`
}

// rateLimitListResponseMatchHeaderJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchHeader]
type rateLimitListResponseMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitListResponseMatchHeadersOp string

const (
	RateLimitListResponseMatchHeadersOpEq RateLimitListResponseMatchHeadersOp = "eq"
	RateLimitListResponseMatchHeadersOpNe RateLimitListResponseMatchHeadersOp = "ne"
)

type RateLimitListResponseMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []RateLimitListResponseMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                `json:"url"`
	JSON rateLimitListResponseMatchRequestJSON `json:"-"`
}

// rateLimitListResponseMatchRequestJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchRequest]
type rateLimitListResponseMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitListResponseMatchRequestMethod string

const (
	RateLimitListResponseMatchRequestMethodGet    RateLimitListResponseMatchRequestMethod = "GET"
	RateLimitListResponseMatchRequestMethodPost   RateLimitListResponseMatchRequestMethod = "POST"
	RateLimitListResponseMatchRequestMethodPut    RateLimitListResponseMatchRequestMethod = "PUT"
	RateLimitListResponseMatchRequestMethodDelete RateLimitListResponseMatchRequestMethod = "DELETE"
	RateLimitListResponseMatchRequestMethodPatch  RateLimitListResponseMatchRequestMethod = "PATCH"
	RateLimitListResponseMatchRequestMethodHead   RateLimitListResponseMatchRequestMethod = "HEAD"
	RateLimitListResponseMatchRequestMethod_All   RateLimitListResponseMatchRequestMethod = "_ALL_"
)

type RateLimitListResponseMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                   `json:"origin_traffic"`
	JSON          rateLimitListResponseMatchResponseJSON `json:"-"`
}

// rateLimitListResponseMatchResponseJSON contains the JSON metadata for the struct
// [RateLimitListResponseMatchResponse]
type rateLimitListResponseMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RateLimitListResponseMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [RateLimitGetResponseUnknown] or [shared.UnionString].
type RateLimitGetResponse interface {
	ImplementsRateLimitGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RateLimitGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type RateLimitUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RateLimitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RateLimitUpdateResponseEnvelope struct {
	Errors   []RateLimitUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitUpdateResponseEnvelopeJSON    `json:"-"`
}

// rateLimitUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitUpdateResponseEnvelope]
type rateLimitUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitUpdateResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rateLimitUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RateLimitUpdateResponseEnvelopeErrors]
type rateLimitUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitUpdateResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    rateLimitUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitUpdateResponseEnvelopeMessages]
type rateLimitUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RateLimitUpdateResponseEnvelopeSuccess bool

const (
	RateLimitUpdateResponseEnvelopeSuccessTrue RateLimitUpdateResponseEnvelopeSuccess = true
)

type RateLimitListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RateLimitListParams]'s query parameters as `url.Values`.
func (r RateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RateLimitListResponseEnvelope struct {
	Errors   []RateLimitListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitListResponseEnvelopeMessages `json:"messages,required"`
	Result   []RateLimitListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RateLimitListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RateLimitListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       rateLimitListResponseEnvelopeJSON       `json:"-"`
}

// rateLimitListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitListResponseEnvelope]
type rateLimitListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    rateLimitListResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RateLimitListResponseEnvelopeErrors]
type rateLimitListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    rateLimitListResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitListResponseEnvelopeMessages]
type rateLimitListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RateLimitListResponseEnvelopeSuccess bool

const (
	RateLimitListResponseEnvelopeSuccessTrue RateLimitListResponseEnvelopeSuccess = true
)

type RateLimitListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       rateLimitListResponseEnvelopeResultInfoJSON `json:"-"`
}

// rateLimitListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RateLimitListResponseEnvelopeResultInfo]
type rateLimitListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitGetResponseEnvelope struct {
	Errors   []RateLimitGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RateLimitGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RateLimitGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RateLimitGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitGetResponseEnvelopeJSON    `json:"-"`
}

// rateLimitGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelope]
type rateLimitGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    rateLimitGetResponseEnvelopeErrorsJSON `json:"-"`
}

// rateLimitGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelopeErrors]
type rateLimitGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RateLimitGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    rateLimitGetResponseEnvelopeMessagesJSON `json:"-"`
}

// rateLimitGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RateLimitGetResponseEnvelopeMessages]
type rateLimitGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RateLimitGetResponseEnvelopeSuccess bool

const (
	RateLimitGetResponseEnvelopeSuccessTrue RateLimitGetResponseEnvelopeSuccess = true
)