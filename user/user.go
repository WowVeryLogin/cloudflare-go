// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	AuditLogs     *AuditLogService
	Billing       *BillingService
	Invites       *InviteService
	Organizations *OrganizationService
	Subscriptions *SubscriptionService
	Tokens        *TokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.AuditLogs = NewAuditLogService(opts...)
	r.Billing = NewBillingService(opts...)
	r.Invites = NewInviteService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Tokens = NewTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Edit(ctx context.Context, body UserEditParams, opts ...option.RequestOption) (res *UserEditResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UserEditResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UserGetResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [user.UserEditResponseUnknown] or [shared.UnionString].
type UserEditResponseUnion interface {
	ImplementsUserUserEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [user.UserGetResponseUnknown] or [shared.UnionString].
type UserGetResponseUnion interface {
	ImplementsUserUserGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserEditParams struct {
	// The country in which the user lives.
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
	// User's telephone number
	Telephone param.Field[string] `json:"telephone"`
	// The zipcode or postal code where the user lives.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r UserEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UserEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success UserEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userEditResponseEnvelopeJSON    `json:"-"`
}

// userEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelope]
type userEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserEditResponseEnvelopeSuccess bool

const (
	UserEditResponseEnvelopeSuccessTrue UserEditResponseEnvelopeSuccess = true
)

func (r UserEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UserGetResponseUnion  `json:"result,required"`
	// Whether the API call was successful
	Success UserGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userGetResponseEnvelopeJSON    `json:"-"`
}

// userGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelope]
type userGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserGetResponseEnvelopeSuccess bool

const (
	UserGetResponseEnvelopeSuccessTrue UserGetResponseEnvelopeSuccess = true
)

func (r UserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}