// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

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

// V1VariantService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewV1VariantService] method instead.
type V1VariantService struct {
	Options []option.RequestOption
}

// NewV1VariantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1VariantService(opts ...option.RequestOption) (r *V1VariantService) {
	r = &V1VariantService{}
	r.Options = opts
	return
}

// Specify variants that allow you to resize images for different use cases.
func (r *V1VariantService) New(ctx context.Context, params V1VariantNewParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing variants.
func (r *V1VariantService) List(ctx context.Context, query V1VariantListParams, opts ...option.RequestOption) (res *ImageVariants, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deleting a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Delete(ctx context.Context, variantID string, body V1VariantDeleteParams, opts ...option.RequestOption) (res *V1VariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", body.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updating a variant purges the cache for all images associated with the variant.
func (r *V1VariantService) Edit(ctx context.Context, variantID string, params V1VariantEditParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", params.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch details for a single variant.
func (r *V1VariantService) Get(ctx context.Context, variantID string, query V1VariantGetParams, opts ...option.RequestOption) (res *ImageVariant, err error) {
	opts = append(r.Options[:], opts...)
	var env V1VariantGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/variants/%s", query.AccountID, variantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageVariant struct {
	Variant ImageVariantVariant `json:"variant"`
	JSON    imageVariantJSON    `json:"-"`
}

// imageVariantJSON contains the JSON metadata for the struct [ImageVariant]
type imageVariantJSON struct {
	Variant     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantJSON) RawJSON() string {
	return r.raw
}

type ImageVariantVariant struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantVariantOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                    `json:"neverRequireSignedURLs"`
	JSON                   imageVariantVariantJSON `json:"-"`
}

// imageVariantVariantJSON contains the JSON metadata for the struct
// [ImageVariantVariant]
type imageVariantVariantJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageVariantVariant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantVariantJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantVariantOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageVariantVariantOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageVariantVariantOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                        `json:"width,required"`
	JSON  imageVariantVariantOptionsJSON `json:"-"`
}

// imageVariantVariantOptionsJSON contains the JSON metadata for the struct
// [ImageVariantVariantOptions]
type imageVariantVariantOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantVariantOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantVariantOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageVariantVariantOptionsFit string

const (
	ImageVariantVariantOptionsFitScaleDown ImageVariantVariantOptionsFit = "scale-down"
	ImageVariantVariantOptionsFitContain   ImageVariantVariantOptionsFit = "contain"
	ImageVariantVariantOptionsFitCover     ImageVariantVariantOptionsFit = "cover"
	ImageVariantVariantOptionsFitCrop      ImageVariantVariantOptionsFit = "crop"
	ImageVariantVariantOptionsFitPad       ImageVariantVariantOptionsFit = "pad"
)

func (r ImageVariantVariantOptionsFit) IsKnown() bool {
	switch r {
	case ImageVariantVariantOptionsFitScaleDown, ImageVariantVariantOptionsFitContain, ImageVariantVariantOptionsFitCover, ImageVariantVariantOptionsFitCrop, ImageVariantVariantOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type ImageVariantVariantOptionsMetadata string

const (
	ImageVariantVariantOptionsMetadataKeep      ImageVariantVariantOptionsMetadata = "keep"
	ImageVariantVariantOptionsMetadataCopyright ImageVariantVariantOptionsMetadata = "copyright"
	ImageVariantVariantOptionsMetadataNone      ImageVariantVariantOptionsMetadata = "none"
)

func (r ImageVariantVariantOptionsMetadata) IsKnown() bool {
	switch r {
	case ImageVariantVariantOptionsMetadataKeep, ImageVariantVariantOptionsMetadataCopyright, ImageVariantVariantOptionsMetadataNone:
		return true
	}
	return false
}

type ImageVariants struct {
	Variants ImageVariantsVariants `json:"variants"`
	JSON     imageVariantsJSON     `json:"-"`
}

// imageVariantsJSON contains the JSON metadata for the struct [ImageVariants]
type imageVariantsJSON struct {
	Variants    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantsJSON) RawJSON() string {
	return r.raw
}

type ImageVariantsVariants struct {
	Hero ImageVariantsVariantsHero `json:"hero"`
	JSON imageVariantsVariantsJSON `json:"-"`
}

// imageVariantsVariantsJSON contains the JSON metadata for the struct
// [ImageVariantsVariants]
type imageVariantsVariantsJSON struct {
	Hero        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantsVariants) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantsVariantsJSON) RawJSON() string {
	return r.raw
}

type ImageVariantsVariantsHero struct {
	ID string `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options ImageVariantsVariantsHeroOptions `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs bool                          `json:"neverRequireSignedURLs"`
	JSON                   imageVariantsVariantsHeroJSON `json:"-"`
}

// imageVariantsVariantsHeroJSON contains the JSON metadata for the struct
// [ImageVariantsVariantsHero]
type imageVariantsVariantsHeroJSON struct {
	ID                     apijson.Field
	Options                apijson.Field
	NeverRequireSignedURLs apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ImageVariantsVariantsHero) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantsVariantsHeroJSON) RawJSON() string {
	return r.raw
}

// Allows you to define image resizing sizes for different use cases.
type ImageVariantsVariantsHeroOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit ImageVariantsVariantsHeroOptionsFit `json:"fit,required"`
	// Maximum height in image pixels.
	Height float64 `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata ImageVariantsVariantsHeroOptionsMetadata `json:"metadata,required"`
	// Maximum width in image pixels.
	Width float64                              `json:"width,required"`
	JSON  imageVariantsVariantsHeroOptionsJSON `json:"-"`
}

// imageVariantsVariantsHeroOptionsJSON contains the JSON metadata for the struct
// [ImageVariantsVariantsHeroOptions]
type imageVariantsVariantsHeroOptionsJSON struct {
	Fit         apijson.Field
	Height      apijson.Field
	Metadata    apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageVariantsVariantsHeroOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageVariantsVariantsHeroOptionsJSON) RawJSON() string {
	return r.raw
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type ImageVariantsVariantsHeroOptionsFit string

const (
	ImageVariantsVariantsHeroOptionsFitScaleDown ImageVariantsVariantsHeroOptionsFit = "scale-down"
	ImageVariantsVariantsHeroOptionsFitContain   ImageVariantsVariantsHeroOptionsFit = "contain"
	ImageVariantsVariantsHeroOptionsFitCover     ImageVariantsVariantsHeroOptionsFit = "cover"
	ImageVariantsVariantsHeroOptionsFitCrop      ImageVariantsVariantsHeroOptionsFit = "crop"
	ImageVariantsVariantsHeroOptionsFitPad       ImageVariantsVariantsHeroOptionsFit = "pad"
)

func (r ImageVariantsVariantsHeroOptionsFit) IsKnown() bool {
	switch r {
	case ImageVariantsVariantsHeroOptionsFitScaleDown, ImageVariantsVariantsHeroOptionsFitContain, ImageVariantsVariantsHeroOptionsFitCover, ImageVariantsVariantsHeroOptionsFitCrop, ImageVariantsVariantsHeroOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type ImageVariantsVariantsHeroOptionsMetadata string

const (
	ImageVariantsVariantsHeroOptionsMetadataKeep      ImageVariantsVariantsHeroOptionsMetadata = "keep"
	ImageVariantsVariantsHeroOptionsMetadataCopyright ImageVariantsVariantsHeroOptionsMetadata = "copyright"
	ImageVariantsVariantsHeroOptionsMetadataNone      ImageVariantsVariantsHeroOptionsMetadata = "none"
)

func (r ImageVariantsVariantsHeroOptionsMetadata) IsKnown() bool {
	switch r {
	case ImageVariantsVariantsHeroOptionsMetadataKeep, ImageVariantsVariantsHeroOptionsMetadataCopyright, ImageVariantsVariantsHeroOptionsMetadataNone:
		return true
	}
	return false
}

// Union satisfied by [images.V1VariantDeleteResponseUnknown] or
// [shared.UnionString].
type V1VariantDeleteResponse interface {
	ImplementsImagesV1VariantDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V1VariantDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V1VariantNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	ID        param.Field[string] `json:"id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantNewParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantNewParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantNewParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantNewParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantNewParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantNewParamsOptionsFit string

const (
	V1VariantNewParamsOptionsFitScaleDown V1VariantNewParamsOptionsFit = "scale-down"
	V1VariantNewParamsOptionsFitContain   V1VariantNewParamsOptionsFit = "contain"
	V1VariantNewParamsOptionsFitCover     V1VariantNewParamsOptionsFit = "cover"
	V1VariantNewParamsOptionsFitCrop      V1VariantNewParamsOptionsFit = "crop"
	V1VariantNewParamsOptionsFitPad       V1VariantNewParamsOptionsFit = "pad"
)

func (r V1VariantNewParamsOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantNewParamsOptionsFitScaleDown, V1VariantNewParamsOptionsFitContain, V1VariantNewParamsOptionsFitCover, V1VariantNewParamsOptionsFitCrop, V1VariantNewParamsOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantNewParamsOptionsMetadata string

const (
	V1VariantNewParamsOptionsMetadataKeep      V1VariantNewParamsOptionsMetadata = "keep"
	V1VariantNewParamsOptionsMetadataCopyright V1VariantNewParamsOptionsMetadata = "copyright"
	V1VariantNewParamsOptionsMetadataNone      V1VariantNewParamsOptionsMetadata = "none"
)

func (r V1VariantNewParamsOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantNewParamsOptionsMetadataKeep, V1VariantNewParamsOptionsMetadataCopyright, V1VariantNewParamsOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantNewResponseEnvelope struct {
	Errors   []V1VariantNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                           `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantNewResponseEnvelopeJSON    `json:"-"`
}

// v1VariantNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantNewResponseEnvelope]
type v1VariantNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    v1VariantNewResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1VariantNewResponseEnvelopeErrors]
type v1VariantNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantNewResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v1VariantNewResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantNewResponseEnvelopeMessages]
type v1VariantNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantNewResponseEnvelopeSuccess bool

const (
	V1VariantNewResponseEnvelopeSuccessTrue V1VariantNewResponseEnvelopeSuccess = true
)

func (r V1VariantNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantListResponseEnvelope struct {
	Errors   []V1VariantListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantListResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariants                           `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantListResponseEnvelopeJSON    `json:"-"`
}

// v1VariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantListResponseEnvelope]
type v1VariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1VariantListResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantListResponseEnvelopeErrors]
type v1VariantListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantListResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantListResponseEnvelopeMessages]
type v1VariantListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantListResponseEnvelopeSuccess bool

const (
	V1VariantListResponseEnvelopeSuccessTrue V1VariantListResponseEnvelopeSuccess = true
)

func (r V1VariantListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantDeleteResponseEnvelope struct {
	Errors   []V1VariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V1VariantDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// v1VariantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantDeleteResponseEnvelope]
type v1VariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantDeleteResponseEnvelopeErrors]
type v1VariantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    v1VariantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantDeleteResponseEnvelopeMessages]
type v1VariantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantDeleteResponseEnvelopeSuccess bool

const (
	V1VariantDeleteResponseEnvelopeSuccessTrue V1VariantDeleteResponseEnvelopeSuccess = true
)

func (r V1VariantDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantEditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Allows you to define image resizing sizes for different use cases.
	Options param.Field[V1VariantEditParamsOptions] `json:"options,required"`
	// Indicates whether the variant can access an image without a signature,
	// regardless of image access control.
	NeverRequireSignedURLs param.Field[bool] `json:"neverRequireSignedURLs"`
}

func (r V1VariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allows you to define image resizing sizes for different use cases.
type V1VariantEditParamsOptions struct {
	// The fit property describes how the width and height dimensions should be
	// interpreted.
	Fit param.Field[V1VariantEditParamsOptionsFit] `json:"fit,required"`
	// Maximum height in image pixels.
	Height param.Field[float64] `json:"height,required"`
	// What EXIF data should be preserved in the output image.
	Metadata param.Field[V1VariantEditParamsOptionsMetadata] `json:"metadata,required"`
	// Maximum width in image pixels.
	Width param.Field[float64] `json:"width,required"`
}

func (r V1VariantEditParamsOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The fit property describes how the width and height dimensions should be
// interpreted.
type V1VariantEditParamsOptionsFit string

const (
	V1VariantEditParamsOptionsFitScaleDown V1VariantEditParamsOptionsFit = "scale-down"
	V1VariantEditParamsOptionsFitContain   V1VariantEditParamsOptionsFit = "contain"
	V1VariantEditParamsOptionsFitCover     V1VariantEditParamsOptionsFit = "cover"
	V1VariantEditParamsOptionsFitCrop      V1VariantEditParamsOptionsFit = "crop"
	V1VariantEditParamsOptionsFitPad       V1VariantEditParamsOptionsFit = "pad"
)

func (r V1VariantEditParamsOptionsFit) IsKnown() bool {
	switch r {
	case V1VariantEditParamsOptionsFitScaleDown, V1VariantEditParamsOptionsFitContain, V1VariantEditParamsOptionsFitCover, V1VariantEditParamsOptionsFitCrop, V1VariantEditParamsOptionsFitPad:
		return true
	}
	return false
}

// What EXIF data should be preserved in the output image.
type V1VariantEditParamsOptionsMetadata string

const (
	V1VariantEditParamsOptionsMetadataKeep      V1VariantEditParamsOptionsMetadata = "keep"
	V1VariantEditParamsOptionsMetadataCopyright V1VariantEditParamsOptionsMetadata = "copyright"
	V1VariantEditParamsOptionsMetadataNone      V1VariantEditParamsOptionsMetadata = "none"
)

func (r V1VariantEditParamsOptionsMetadata) IsKnown() bool {
	switch r {
	case V1VariantEditParamsOptionsMetadataKeep, V1VariantEditParamsOptionsMetadataCopyright, V1VariantEditParamsOptionsMetadataNone:
		return true
	}
	return false
}

type V1VariantEditResponseEnvelope struct {
	Errors   []V1VariantEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                            `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantEditResponseEnvelopeJSON    `json:"-"`
}

// v1VariantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantEditResponseEnvelope]
type v1VariantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    v1VariantEditResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [V1VariantEditResponseEnvelopeErrors]
type v1VariantEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantEditResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    v1VariantEditResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantEditResponseEnvelopeMessages]
type v1VariantEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantEditResponseEnvelopeSuccess bool

const (
	V1VariantEditResponseEnvelopeSuccessTrue V1VariantEditResponseEnvelopeSuccess = true
)

func (r V1VariantEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V1VariantGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type V1VariantGetResponseEnvelope struct {
	Errors   []V1VariantGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V1VariantGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageVariant                           `json:"result,required"`
	// Whether the API call was successful
	Success V1VariantGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v1VariantGetResponseEnvelopeJSON    `json:"-"`
}

// v1VariantGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [V1VariantGetResponseEnvelope]
type v1VariantGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    v1VariantGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v1VariantGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V1VariantGetResponseEnvelopeErrors]
type v1VariantGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V1VariantGetResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    v1VariantGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v1VariantGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [V1VariantGetResponseEnvelopeMessages]
type v1VariantGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1VariantGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1VariantGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V1VariantGetResponseEnvelopeSuccess bool

const (
	V1VariantGetResponseEnvelopeSuccessTrue V1VariantGetResponseEnvelopeSuccess = true
)

func (r V1VariantGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V1VariantGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}