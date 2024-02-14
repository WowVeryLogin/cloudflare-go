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

// AlertingV3DestinationPagerdutyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAlertingV3DestinationPagerdutyService] method instead.
type AlertingV3DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewAlertingV3DestinationPagerdutyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationPagerdutyService(opts ...option.RequestOption) (r *AlertingV3DestinationPagerdutyService) {
	r = &AlertingV3DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Get a list of all configured PagerDuty services.
func (r *AlertingV3DestinationPagerdutyService) NotificationDestinationsWithPagerDutyListPagerDutyServices(ctx context.Context, accountID string, opts ...option.RequestOption) (res *[]AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponse struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string                                                                                               `json:"name"`
	JSON alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponse]
type alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelope struct {
	Errors   []AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessages `json:"messages,required"`
	Result   []AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelope]
type alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrors struct {
	Code    int64                                                                                                              `json:"code,required"`
	Message string                                                                                                             `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrors]
type alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessages struct {
	Code    int64                                                                                                                `json:"code,required"`
	Message string                                                                                                               `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessages]
type alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeSuccessTrue AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeSuccess = true
)

type AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                                `json:"total_count"`
	JSON       alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfo]
type alertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNotificationDestinationsWithPagerDutyListPagerDutyServicesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}