// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingIPGeolocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingIPGeolocationService]
// method instead.
type SettingIPGeolocationService struct {
	Options []option.RequestOption
}

// NewSettingIPGeolocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingIPGeolocationService(opts ...option.RequestOption) (r *SettingIPGeolocationService) {
	r = &SettingIPGeolocationService{}
	r.Options = opts
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *SettingIPGeolocationService) Edit(ctx context.Context, params SettingIPGeolocationEditParams, opts ...option.RequestOption) (res *IPGeolocation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
func (r *SettingIPGeolocationService) Get(ctx context.Context, query SettingIPGeolocationGetParams, opts ...option.RequestOption) (res *IPGeolocation, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingIPGeolocationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/ip_geolocation", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type IPGeolocation struct {
	// ID of the zone setting.
	ID IPGeolocationID `json:"id,required"`
	// Current value of the zone setting.
	Value IPGeolocationValue `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable IPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time         `json:"modified_on,nullable" format:"date-time"`
	JSON       ipGeolocationJSON `json:"-"`
}

// ipGeolocationJSON contains the JSON metadata for the struct [IPGeolocation]
type ipGeolocationJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipGeolocationJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type IPGeolocationID string

const (
	IPGeolocationIDIPGeolocation IPGeolocationID = "ip_geolocation"
)

func (r IPGeolocationID) IsKnown() bool {
	switch r {
	case IPGeolocationIDIPGeolocation:
		return true
	}
	return false
}

// Current value of the zone setting.
type IPGeolocationValue string

const (
	IPGeolocationValueOn  IPGeolocationValue = "on"
	IPGeolocationValueOff IPGeolocationValue = "off"
)

func (r IPGeolocationValue) IsKnown() bool {
	switch r {
	case IPGeolocationValueOn, IPGeolocationValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type IPGeolocationEditable bool

const (
	IPGeolocationEditableTrue  IPGeolocationEditable = true
	IPGeolocationEditableFalse IPGeolocationEditable = false
)

func (r IPGeolocationEditable) IsKnown() bool {
	switch r {
	case IPGeolocationEditableTrue, IPGeolocationEditableFalse:
		return true
	}
	return false
}

type SettingIPGeolocationEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the zone setting.
	Value param.Field[SettingIPGeolocationEditParamsValue] `json:"value,required"`
}

func (r SettingIPGeolocationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type SettingIPGeolocationEditParamsValue string

const (
	SettingIPGeolocationEditParamsValueOn  SettingIPGeolocationEditParamsValue = "on"
	SettingIPGeolocationEditParamsValueOff SettingIPGeolocationEditParamsValue = "off"
)

func (r SettingIPGeolocationEditParamsValue) IsKnown() bool {
	switch r {
	case SettingIPGeolocationEditParamsValueOn, SettingIPGeolocationEditParamsValueOff:
		return true
	}
	return false
}

type SettingIPGeolocationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result IPGeolocation                                `json:"result"`
	JSON   settingIPGeolocationEditResponseEnvelopeJSON `json:"-"`
}

// settingIPGeolocationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingIPGeolocationEditResponseEnvelope]
type settingIPGeolocationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPGeolocationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingIPGeolocationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingIPGeolocationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
	// pass the country code to you.
	// (https://support.cloudflare.com/hc/en-us/articles/200168236).
	Result IPGeolocation                               `json:"result"`
	JSON   settingIPGeolocationGetResponseEnvelopeJSON `json:"-"`
}

// settingIPGeolocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingIPGeolocationGetResponseEnvelope]
type settingIPGeolocationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingIPGeolocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingIPGeolocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}