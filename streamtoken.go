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

// StreamTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamTokenService] method
// instead.
type StreamTokenService struct {
	Options []option.RequestOption
}

// NewStreamTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamTokenService(opts ...option.RequestOption) (r *StreamTokenService) {
	r = &StreamTokenService{}
	r.Options = opts
	return
}

// Creates a signed URL token for a video. If a body is not provided in the
// request, a token is created with default values.
func (r *StreamTokenService) StreamVideosNewSignedURLTokensForVideos(ctx context.Context, accountID string, identifier string, body StreamTokenStreamVideosNewSignedURLTokensForVideosParams, opts ...option.RequestOption) (res *StreamTokenStreamVideosNewSignedURLTokensForVideosResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/%s/token", accountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamTokenStreamVideosNewSignedURLTokensForVideosResponse struct {
	// The signed token used with the signed URLs feature.
	Token string                                                         `json:"token"`
	JSON  streamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON `json:"-"`
}

// streamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON contains the JSON
// metadata for the struct
// [StreamTokenStreamVideosNewSignedURLTokensForVideosResponse]
type streamTokenStreamVideosNewSignedURLTokensForVideosResponseJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenStreamVideosNewSignedURLTokensForVideosResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamTokenStreamVideosNewSignedURLTokensForVideosParams struct {
	// The optional ID of a Stream signing key. If present, the `pem` field is also
	// required.
	ID param.Field[string] `json:"id"`
	// The optional list of access rule constraints on the token. Access can be blocked
	// or allowed based on an IP, IP range, or by country. Access rules are evaluated
	// from first to last. If a rule matches, the associated action is applied and no
	// further rules are evaluated.
	AccessRules param.Field[[]StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule] `json:"accessRules"`
	// The optional boolean value that enables using signed tokens to access MP4
	// download links for a video.
	Downloadable param.Field[bool] `json:"downloadable"`
	// The optional unix epoch timestamp that specficies the time after a token is not
	// accepted. The maximum time specification is 24 hours from issuing time. If this
	// field is not set, the default is one hour after issuing.
	Exp param.Field[int64] `json:"exp"`
	// The optional unix epoch timestamp that specifies the time before a the token is
	// not accepted. If this field is not set, the default is one hour before issuing.
	Nbf param.Field[int64] `json:"nbf"`
	// The optional base64 encoded private key in PEM format associated with a Stream
	// signing key. If present, the `id` field is also required.
	Pem param.Field[string] `json:"pem"`
}

func (r StreamTokenStreamVideosNewSignedURLTokensForVideosParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines rules for fine-grained control over content than signed URL tokens
// alone. Access rules primarily make tokens conditionally valid based on user
// information. Access Rules are specified on token payloads as the `accessRules`
// property containing an array of Rule objects.
type StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule struct {
	// The action to take when a request matches a rule. If the action is `block`, the
	// signed token blocks views for viewers matching the rule.
	Action param.Field[StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction] `json:"action"`
	// An array of 2-letter country codes in ISO 3166-1 Alpha-2 format used to match
	// requests.
	Country param.Field[[]string] `json:"country"`
	// An array of IPv4 or IPV6 addresses or CIDRs used to match requests.
	IP param.Field[[]string] `json:"ip"`
	// Lists available rule types to match for requests. An `any` type matches all
	// requests and can be used as a wildcard to apply default actions after other
	// rules.
	Type param.Field[StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType] `json:"type"`
}

func (r StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when a request matches a rule. If the action is `block`, the
// signed token blocks views for viewers matching the rule.
type StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction string

const (
	StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesActionAllow StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction = "allow"
	StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesActionBlock StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesAction = "block"
)

// Lists available rule types to match for requests. An `any` type matches all
// requests and can be used as a wildcard to apply default actions after other
// rules.
type StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType string

const (
	StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeAny            StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "any"
	StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeIPSrc          StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "ip.src"
	StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesTypeIPGeoipCountry StreamTokenStreamVideosNewSignedURLTokensForVideosParamsAccessRulesType = "ip.geoip.country"
)

type StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelope struct {
	Errors   []StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamTokenStreamVideosNewSignedURLTokensForVideosResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeJSON    `json:"-"`
}

// streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelope]
type streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrors struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrorsJSON `json:"-"`
}

// streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrors]
type streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessages struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessagesJSON `json:"-"`
}

// streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessages]
type streamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeSuccess bool

const (
	StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeSuccessTrue StreamTokenStreamVideosNewSignedURLTokensForVideosResponseEnvelopeSuccess = true
)