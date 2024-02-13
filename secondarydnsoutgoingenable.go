// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSOutgoingEnableService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSecondaryDNSOutgoingEnableService] method instead.
type SecondaryDNSOutgoingEnableService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSOutgoingEnableService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingEnableService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingEnableService) {
	r = &SecondaryDNSOutgoingEnableService{}
	r.Options = opts
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *SecondaryDNSOutgoingEnableService) SecondaryDNSPrimaryZoneEnableOutgoingZoneTransfers(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing/enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessages `json:"messages,required"`
	// The zone transfer status of a primary zone
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelope]
type secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrors]
type secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessages]
type secondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeSuccessTrue SecondaryDNSOutgoingEnableSecondaryDNSPrimaryZoneEnableOutgoingZoneTransfersResponseEnvelopeSuccess = true
)