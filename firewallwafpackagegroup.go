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

// FirewallWAFPackageGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewFirewallWAFPackageGroupService] method instead.
type FirewallWAFPackageGroupService struct {
	Options []option.RequestOption
}

// NewFirewallWAFPackageGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFPackageGroupService(opts ...option.RequestOption) (r *FirewallWAFPackageGroupService) {
	r = &FirewallWAFPackageGroupService{}
	r.Options = opts
	return
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Update(ctx context.Context, zoneID string, packageID string, groupID string, body FirewallWAFPackageGroupUpdateParams, opts ...option.RequestOption) (res *FirewallWAFPackageGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageGroupUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) Get(ctx context.Context, zoneID string, packageID string, groupID string, opts ...option.RequestOption) (res *FirewallWAFPackageGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageGroupGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFPackageGroupService) WAFRuleGroupsListWAFRuleGroups(ctx context.Context, zoneID string, packageID string, query FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParams, opts ...option.RequestOption) (res *[]FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallWAFPackageGroupUpdateResponseUnknown],
// [FirewallWAFPackageGroupUpdateResponseArray] or [shared.UnionString].
type FirewallWAFPackageGroupUpdateResponse interface {
	ImplementsFirewallWAFPackageGroupUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageGroupUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageGroupUpdateResponseArray []interface{}

func (r FirewallWAFPackageGroupUpdateResponseArray) ImplementsFirewallWAFPackageGroupUpdateResponse() {
}

// Union satisfied by [FirewallWAFPackageGroupGetResponseUnknown],
// [FirewallWAFPackageGroupGetResponseArray] or [shared.UnionString].
type FirewallWAFPackageGroupGetResponse interface {
	ImplementsFirewallWAFPackageGroupGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallWAFPackageGroupGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallWAFPackageGroupGetResponseArray []interface{}

func (r FirewallWAFPackageGroupGetResponseArray) ImplementsFirewallWAFPackageGroupGetResponse() {}

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponse struct {
	// The unique identifier of the rule group.
	ID string `json:"id,required"`
	// An informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseMode `json:"mode,required"`
	// The name of the rule group.
	Name string `json:"name,required"`
	// The number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// The available states for the rule group.
	AllowedModes []FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedMode `json:"allowed_modes"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The unique identifier of a WAF package.
	PackageID string                                                            `json:"package_id"`
	JSON      firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseJSON `json:"-"`
}

// firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseJSON contains the
// JSON metadata for the struct
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponse]
type firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseJSON struct {
	ID                 apijson.Field
	Description        apijson.Field
	Mode               apijson.Field
	Name               apijson.Field
	RulesCount         apijson.Field
	AllowedModes       apijson.Field
	ModifiedRulesCount apijson.Field
	PackageID          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseMode string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseModeOn  FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseMode = "on"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseModeOff FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseMode = "off"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedMode string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedModeOn  FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedMode = "on"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedModeOff FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseAllowedMode = "off"
)

type FirewallWAFPackageGroupUpdateParams struct {
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[FirewallWAFPackageGroupUpdateParamsMode] `json:"mode"`
}

func (r FirewallWAFPackageGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupUpdateParamsMode string

const (
	FirewallWAFPackageGroupUpdateParamsModeOn  FirewallWAFPackageGroupUpdateParamsMode = "on"
	FirewallWAFPackageGroupUpdateParamsModeOff FirewallWAFPackageGroupUpdateParamsMode = "off"
)

type FirewallWAFPackageGroupUpdateResponseEnvelope struct {
	Errors   []FirewallWAFPackageGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageGroupUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupUpdateResponseEnvelope]
type firewallWAFPackageGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    firewallWAFPackageGroupUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [FirewallWAFPackageGroupUpdateResponseEnvelopeErrors]
type firewallWAFPackageGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    firewallWAFPackageGroupUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageGroupUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageGroupUpdateResponseEnvelopeMessages]
type firewallWAFPackageGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupUpdateResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageGroupUpdateResponseEnvelopeSuccessTrue FirewallWAFPackageGroupUpdateResponseEnvelopeSuccess = true
)

type FirewallWAFPackageGroupGetResponseEnvelope struct {
	Errors   []FirewallWAFPackageGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageGroupGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallWAFPackageGroupGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallWAFPackageGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [FirewallWAFPackageGroupGetResponseEnvelope]
type firewallWAFPackageGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFPackageGroupGetResponseEnvelopeErrors]
type firewallWAFPackageGroupGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FirewallWAFPackageGroupGetResponseEnvelopeMessages]
type firewallWAFPackageGroupGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupGetResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageGroupGetResponseEnvelopeSuccessTrue FirewallWAFPackageGroupGetResponseEnvelopeSuccess = true
)

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParams struct {
	// The direction used to sort returned rule groups.
	Direction param.Field[FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatch] `query:"match"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMode] `query:"mode"`
	// The field used to sort returned rule groups.
	Order param.Field[FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rule groups per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParams]'s query parameters
// as `url.Values`.
func (r FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rule groups.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirection string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirectionAsc  FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirection = "asc"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirectionDesc FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatch string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatchAny FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatch = "any"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatchAll FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMatch = "all"
)

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMode string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsModeOn  FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMode = "on"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsModeOff FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsMode = "off"
)

// The field used to sort returned rule groups.
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrder string

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrderMode       FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrder = "mode"
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrderRulesCount FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsParamsOrder = "rules_count"
)

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelope struct {
	Errors   []FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeJSON       `json:"-"`
}

// firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelope]
type firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrors]
type firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessages]
type firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeSuccess bool

const (
	FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeSuccessTrue FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeSuccess = true
)

type FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                             `json:"total_count"`
	JSON       firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfo]
type firewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFPackageGroupWAFRuleGroupsListWAFRuleGroupsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}