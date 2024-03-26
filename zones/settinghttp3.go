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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingHTTP3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSettingHTTP3Service] method
// instead.
type SettingHTTP3Service struct {
	Options []option.RequestOption
}

// NewSettingHTTP3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingHTTP3Service(opts ...option.RequestOption) (r *SettingHTTP3Service) {
	r = &SettingHTTP3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *SettingHTTP3Service) Edit(ctx context.Context, params SettingHTTP3EditParams, opts ...option.RequestOption) (res *ZonesHTTP3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3EditResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Value of the HTTP3 setting.
func (r *SettingHTTP3Service) Get(ctx context.Context, query SettingHTTP3GetParams, opts ...option.RequestOption) (res *ZonesHTTP3, err error) {
	opts = append(r.Options[:], opts...)
	var env SettingHTTP3GetResponseEnvelope
	path := fmt.Sprintf("zones/%s/settings/http3", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// HTTP3 enabled for this zone.
type ZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZonesHTTP3ID `json:"id,required"`
	// Current value of the zone setting.
	Value ZonesHTTP3Value `json:"value,required"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time      `json:"modified_on,nullable" format:"date-time"`
	JSON       zonesHTTP3JSON `json:"-"`
}

// zonesHTTP3JSON contains the JSON metadata for the struct [ZonesHTTP3]
type zonesHTTP3JSON struct {
	ID          apijson.Field
	Value       apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesHTTP3JSON) RawJSON() string {
	return r.raw
}

func (r ZonesHTTP3) implementsZonesSettingEditResponse() {}

func (r ZonesHTTP3) implementsZonesSettingGetResponse() {}

// ID of the zone setting.
type ZonesHTTP3ID string

const (
	ZonesHTTP3IDHTTP3 ZonesHTTP3ID = "http3"
)

func (r ZonesHTTP3ID) IsKnown() bool {
	switch r {
	case ZonesHTTP3IDHTTP3:
		return true
	}
	return false
}

// Current value of the zone setting.
type ZonesHTTP3Value string

const (
	ZonesHTTP3ValueOn  ZonesHTTP3Value = "on"
	ZonesHTTP3ValueOff ZonesHTTP3Value = "off"
)

func (r ZonesHTTP3Value) IsKnown() bool {
	switch r {
	case ZonesHTTP3ValueOn, ZonesHTTP3ValueOff:
		return true
	}
	return false
}

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZonesHTTP3Editable bool

const (
	ZonesHTTP3EditableTrue  ZonesHTTP3Editable = true
	ZonesHTTP3EditableFalse ZonesHTTP3Editable = false
)

func (r ZonesHTTP3Editable) IsKnown() bool {
	switch r {
	case ZonesHTTP3EditableTrue, ZonesHTTP3EditableFalse:
		return true
	}
	return false
}

// HTTP3 enabled for this zone.
type ZonesHTTP3Param struct {
	// ID of the zone setting.
	ID param.Field[ZonesHTTP3ID] `json:"id,required"`
	// Current value of the zone setting.
	Value param.Field[ZonesHTTP3Value] `json:"value,required"`
}

func (r ZonesHTTP3Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonesHTTP3Param) implementsZonesSettingEditParamsItem() {}

type SettingHTTP3EditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the HTTP3 setting.
	Value param.Field[SettingHTTP3EditParamsValue] `json:"value,required"`
}

func (r SettingHTTP3EditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the HTTP3 setting.
type SettingHTTP3EditParamsValue string

const (
	SettingHTTP3EditParamsValueOn  SettingHTTP3EditParamsValue = "on"
	SettingHTTP3EditParamsValueOff SettingHTTP3EditParamsValue = "off"
)

func (r SettingHTTP3EditParamsValue) IsKnown() bool {
	switch r {
	case SettingHTTP3EditParamsValueOn, SettingHTTP3EditParamsValueOff:
		return true
	}
	return false
}

type SettingHTTP3EditResponseEnvelope struct {
	Errors   []SettingHTTP3EditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP3EditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZonesHTTP3                           `json:"result"`
	JSON   settingHTTP3EditResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3EditResponseEnvelope]
type settingHTTP3EditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3EditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3EditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    settingHTTP3EditResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP3EditResponseEnvelopeErrors]
type settingHTTP3EditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3EditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3EditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    settingHTTP3EditResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP3EditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP3EditResponseEnvelopeMessages]
type settingHTTP3EditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3EditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3EditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3GetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingHTTP3GetResponseEnvelope struct {
	Errors   []SettingHTTP3GetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingHTTP3GetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success bool `json:"success,required"`
	// HTTP3 enabled for this zone.
	Result ZonesHTTP3                          `json:"result"`
	JSON   settingHTTP3GetResponseEnvelopeJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingHTTP3GetResponseEnvelope]
type settingHTTP3GetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3GetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3GetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    settingHTTP3GetResponseEnvelopeErrorsJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingHTTP3GetResponseEnvelopeErrors]
type settingHTTP3GetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3GetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingHTTP3GetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    settingHTTP3GetResponseEnvelopeMessagesJSON `json:"-"`
}

// settingHTTP3GetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingHTTP3GetResponseEnvelopeMessages]
type settingHTTP3GetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingHTTP3GetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingHTTP3GetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}