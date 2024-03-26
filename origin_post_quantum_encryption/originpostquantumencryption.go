// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_post_quantum_encryption

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// OriginPostQuantumEncryptionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewOriginPostQuantumEncryptionService] method instead.
type OriginPostQuantumEncryptionService struct {
	Options []option.RequestOption
}

// NewOriginPostQuantumEncryptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginPostQuantumEncryptionService(opts ...option.RequestOption) (r *OriginPostQuantumEncryptionService) {
	r = &OriginPostQuantumEncryptionService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Update(ctx context.Context, params OriginPostQuantumEncryptionUpdateParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginPostQuantumEncryptionUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Get(ctx context.Context, query OriginPostQuantumEncryptionGetParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env OriginPostQuantumEncryptionGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [origin_post_quantum_encryption.OriginPostQuantumEncryptionUpdateResponseUnknown]
// or [shared.UnionString].
type OriginPostQuantumEncryptionUpdateResponse interface {
	ImplementsOriginPostQuantumEncryptionOriginPostQuantumEncryptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [origin_post_quantum_encryption.OriginPostQuantumEncryptionGetResponseUnknown]
// or [shared.UnionString].
type OriginPostQuantumEncryptionGetResponse interface {
	ImplementsOriginPostQuantumEncryptionOriginPostQuantumEncryptionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPostQuantumEncryptionGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type OriginPostQuantumEncryptionUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r OriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionUpdateParamsValue string

const (
	OriginPostQuantumEncryptionUpdateParamsValuePreferred OriginPostQuantumEncryptionUpdateParamsValue = "preferred"
	OriginPostQuantumEncryptionUpdateParamsValueSupported OriginPostQuantumEncryptionUpdateParamsValue = "supported"
	OriginPostQuantumEncryptionUpdateParamsValueOff       OriginPostQuantumEncryptionUpdateParamsValue = "off"
)

func (r OriginPostQuantumEncryptionUpdateParamsValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateParamsValuePreferred, OriginPostQuantumEncryptionUpdateParamsValueSupported, OriginPostQuantumEncryptionUpdateParamsValueOff:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionUpdateResponseEnvelope struct {
	Errors   []OriginPostQuantumEncryptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginPostQuantumEncryptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginPostQuantumEncryptionUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionUpdateResponseEnvelope]
type originPostQuantumEncryptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginPostQuantumEncryptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    originPostQuantumEncryptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [OriginPostQuantumEncryptionUpdateResponseEnvelopeErrors]
type originPostQuantumEncryptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginPostQuantumEncryptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    originPostQuantumEncryptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [OriginPostQuantumEncryptionUpdateResponseEnvelopeMessages]
type originPostQuantumEncryptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginPostQuantumEncryptionGetResponseEnvelope struct {
	Errors   []OriginPostQuantumEncryptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginPostQuantumEncryptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   OriginPostQuantumEncryptionGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionGetResponseEnvelope]
type originPostQuantumEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginPostQuantumEncryptionGetResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [OriginPostQuantumEncryptionGetResponseEnvelopeErrors]
type originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginPostQuantumEncryptionGetResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [OriginPostQuantumEncryptionGetResponseEnvelopeMessages]
type originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionGetResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionGetResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}