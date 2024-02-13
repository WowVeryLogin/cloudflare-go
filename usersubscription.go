// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserSubscriptionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserSubscriptionService] method
// instead.
type UserSubscriptionService struct {
	Options []option.RequestOption
}

// NewUserSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSubscriptionService(opts ...option.RequestOption) (r *UserSubscriptionService) {
	r = &UserSubscriptionService{}
	r.Options = opts
	return
}

// Updates a user's subscriptions.
func (r *UserSubscriptionService) Update(ctx context.Context, identifier string, body UserSubscriptionUpdateParams, opts ...option.RequestOption) (res *UserSubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionUpdateResponseEnvelope
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a user's subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists all of a user's subscriptions.
func (r *UserSubscriptionService) UserSubscriptionGetUserSubscriptions(ctx context.Context, opts ...option.RequestOption) (res *[]UserSubscriptionUserSubscriptionGetUserSubscriptionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelope
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserSubscriptionUpdateResponseUnknown] or
// [shared.UnionString].
type UserSubscriptionUpdateResponse interface {
	ImplementsUserSubscriptionUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSubscriptionUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserSubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                             `json:"subscription_id"`
	JSON           userSubscriptionDeleteResponseJSON `json:"-"`
}

// userSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponse]
type userSubscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponse struct {
	// Subscription identifier tag.
	ID  string                                                          `json:"id"`
	App UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValue `json:"component_values"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState `json:"state"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseZone `json:"zone"`
	JSON userSubscriptionUserSubscriptionGetUserSubscriptionsResponseJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseJSON contains the
// JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponse]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseJSON struct {
	ID                 apijson.Field
	App                apijson.Field
	ComponentValues    apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	Zone               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseApp struct {
	// app install id.
	InstallID string                                                              `json:"install_id"`
	JSON      userSubscriptionUserSubscriptionGetUserSubscriptionsResponseAppJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseAppJSON contains the
// JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseApp]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A component value for a subscription.
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                                                                        `json:"value"`
	JSON  userSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValueJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValueJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValue]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency string

const (
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequencyWeekly    UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency = "weekly"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequencyMonthly   UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency = "monthly"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequencyQuarterly UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency = "quarterly"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequencyYearly    UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlan struct {
	// The ID of the rate plan.
	ID interface{} `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string                                                                 `json:"sets"`
	JSON userSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlanJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlanJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlan]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state that the subscription is in.
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState string

const (
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateTrial           UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Trial"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateProvisioned     UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Provisioned"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStatePaid            UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Paid"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateAwaitingPayment UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "AwaitingPayment"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateCancelled       UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Cancelled"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateFailed          UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Failed"
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseStateExpired         UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseState = "Expired"
)

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string                                                               `json:"name"`
	JSON userSubscriptionUserSubscriptionGetUserSubscriptionsResponseZoneJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseZoneJSON contains
// the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseZone]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateParams struct {
	App param.Field[UserSubscriptionUpdateParamsApp] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]UserSubscriptionUpdateParamsComponentValue] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[UserSubscriptionUpdateParamsFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[UserSubscriptionUpdateParamsRatePlan] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[UserSubscriptionUpdateParamsZone] `json:"zone"`
}

func (r UserSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionUpdateParamsApp struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r UserSubscriptionUpdateParamsApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type UserSubscriptionUpdateParamsComponentValue struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r UserSubscriptionUpdateParamsComponentValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type UserSubscriptionUpdateParamsFrequency string

const (
	UserSubscriptionUpdateParamsFrequencyWeekly    UserSubscriptionUpdateParamsFrequency = "weekly"
	UserSubscriptionUpdateParamsFrequencyMonthly   UserSubscriptionUpdateParamsFrequency = "monthly"
	UserSubscriptionUpdateParamsFrequencyQuarterly UserSubscriptionUpdateParamsFrequency = "quarterly"
	UserSubscriptionUpdateParamsFrequencyYearly    UserSubscriptionUpdateParamsFrequency = "yearly"
)

// The rate plan applied to the subscription.
type UserSubscriptionUpdateParamsRatePlan struct {
	// The ID of the rate plan.
	ID param.Field[interface{}] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r UserSubscriptionUpdateParamsRatePlan) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A simple zone object. May have null properties if not a zone subscription.
type UserSubscriptionUpdateParamsZone struct {
}

func (r UserSubscriptionUpdateParamsZone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserSubscriptionUpdateResponseEnvelope struct {
	Errors   []UserSubscriptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserSubscriptionUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserSubscriptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserSubscriptionUpdateResponseEnvelope]
type userSubscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserSubscriptionUpdateResponseEnvelopeErrors]
type userSubscriptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    userSubscriptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [UserSubscriptionUpdateResponseEnvelopeMessages]
type userSubscriptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionUpdateResponseEnvelopeSuccess bool

const (
	UserSubscriptionUpdateResponseEnvelopeSuccessTrue UserSubscriptionUpdateResponseEnvelopeSuccess = true
)

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelope struct {
	Errors   []UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserSubscriptionUserSubscriptionGetUserSubscriptionsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeJSON       `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelope]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrors struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrorsJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrors]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessages struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessagesJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessages]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeSuccess bool

const (
	UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeSuccessTrue UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeSuccess = true
)

type UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                            `json:"total_count"`
	JSON       userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfo]
type userSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUserSubscriptionGetUserSubscriptionsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}