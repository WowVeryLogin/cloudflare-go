// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

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

// ClipService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewClipService] method instead.
type ClipService struct {
	Options []option.RequestOption
}

// NewClipService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewClipService(opts ...option.RequestOption) (r *ClipService) {
	r = &ClipService{}
	r.Options = opts
	return
}

// Clips a video based on the specified start and end times provided in seconds.
func (r *ClipService) New(ctx context.Context, params ClipNewParams, opts ...option.RequestOption) (res *StreamClipping, err error) {
	opts = append(r.Options[:], opts...)
	var env ClipNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/clip", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type StreamClipping struct {
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins []string `json:"allowedOrigins"`
	// The unique video identifier (UID).
	ClippedFromVideoUid string `json:"clippedFromVideoUID"`
	// The date and time the clip was created.
	Created time.Time `json:"created" format:"date-time"`
	// A user-defined identifier for the media creator.
	Creator string `json:"creator"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds int64 `json:"endTimeSeconds"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds int64 `json:"maxDurationSeconds"`
	// A user modifiable key-value store used to reference other systems of record for
	// managing videos.
	Meta interface{} `json:"meta"`
	// The date and time the live input was last modified.
	Modified time.Time              `json:"modified" format:"date-time"`
	Playback StreamClippingPlayback `json:"playback"`
	// The video's preview page URI. This field is omitted until encoding is complete.
	Preview string `json:"preview" format:"uri"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds int64 `json:"startTimeSeconds"`
	// Specifies the processing status for all quality levels for a video.
	Status StreamClippingStatus `json:"status"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct float64                 `json:"thumbnailTimestampPct"`
	Watermark             StreamClippingWatermark `json:"watermark"`
	JSON                  streamClippingJSON      `json:"-"`
}

// streamClippingJSON contains the JSON metadata for the struct [StreamClipping]
type streamClippingJSON struct {
	AllowedOrigins        apijson.Field
	ClippedFromVideoUid   apijson.Field
	Created               apijson.Field
	Creator               apijson.Field
	EndTimeSeconds        apijson.Field
	MaxDurationSeconds    apijson.Field
	Meta                  apijson.Field
	Modified              apijson.Field
	Playback              apijson.Field
	Preview               apijson.Field
	RequireSignedURLs     apijson.Field
	StartTimeSeconds      apijson.Field
	Status                apijson.Field
	ThumbnailTimestampPct apijson.Field
	Watermark             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *StreamClipping) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamClippingJSON) RawJSON() string {
	return r.raw
}

type StreamClippingPlayback struct {
	// DASH Media Presentation Description for the video.
	Dash string `json:"dash"`
	// The HLS manifest for the video.
	Hls  string                     `json:"hls"`
	JSON streamClippingPlaybackJSON `json:"-"`
}

// streamClippingPlaybackJSON contains the JSON metadata for the struct
// [StreamClippingPlayback]
type streamClippingPlaybackJSON struct {
	Dash        apijson.Field
	Hls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClippingPlayback) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamClippingPlaybackJSON) RawJSON() string {
	return r.raw
}

// Specifies the processing status for all quality levels for a video.
type StreamClippingStatus string

const (
	StreamClippingStatusPendingupload StreamClippingStatus = "pendingupload"
	StreamClippingStatusDownloading   StreamClippingStatus = "downloading"
	StreamClippingStatusQueued        StreamClippingStatus = "queued"
	StreamClippingStatusInprogress    StreamClippingStatus = "inprogress"
	StreamClippingStatusReady         StreamClippingStatus = "ready"
	StreamClippingStatusError         StreamClippingStatus = "error"
)

func (r StreamClippingStatus) IsKnown() bool {
	switch r {
	case StreamClippingStatusPendingupload, StreamClippingStatusDownloading, StreamClippingStatusQueued, StreamClippingStatusInprogress, StreamClippingStatusReady, StreamClippingStatusError:
		return true
	}
	return false
}

type StreamClippingWatermark struct {
	// The unique identifier for the watermark profile.
	Uid  string                      `json:"uid"`
	JSON streamClippingWatermarkJSON `json:"-"`
}

// streamClippingWatermarkJSON contains the JSON metadata for the struct
// [StreamClippingWatermark]
type streamClippingWatermarkJSON struct {
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamClippingWatermark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamClippingWatermarkJSON) RawJSON() string {
	return r.raw
}

type ClipNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The unique video identifier (UID).
	ClippedFromVideoUid param.Field[string] `json:"clippedFromVideoUID,required"`
	// Specifies the end time for the video clip in seconds.
	EndTimeSeconds param.Field[int64] `json:"endTimeSeconds,required"`
	// Specifies the start time for the video clip in seconds.
	StartTimeSeconds param.Field[int64] `json:"startTimeSeconds,required"`
	// Lists the origins allowed to display the video. Enter allowed origin domains in
	// an array and use `*` for wildcard subdomains. Empty arrays allow the video to be
	// viewed on any origin.
	AllowedOrigins param.Field[[]string] `json:"allowedOrigins"`
	// A user-defined identifier for the media creator.
	Creator param.Field[string] `json:"creator"`
	// The maximum duration in seconds for a video upload. Can be set for a video that
	// is not yet uploaded to limit its duration. Uploads that exceed the specified
	// duration will fail during processing. A value of `-1` means the value is
	// unknown.
	MaxDurationSeconds param.Field[int64] `json:"maxDurationSeconds"`
	// Indicates whether the video can be a accessed using the UID. When set to `true`,
	// a signed token must be generated with a signing key to view the video.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
	// The timestamp for a thumbnail image calculated as a percentage value of the
	// video's duration. To convert from a second-wise timestamp to a percentage,
	// divide the desired timestamp by the total duration of the video. If this value
	// is not set, the default thumbnail image is taken from 0s of the video.
	ThumbnailTimestampPct param.Field[float64]                `json:"thumbnailTimestampPct"`
	Watermark             param.Field[ClipNewParamsWatermark] `json:"watermark"`
}

func (r ClipNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClipNewParamsWatermark struct {
	// The unique identifier for the watermark profile.
	Uid param.Field[string] `json:"uid"`
}

func (r ClipNewParamsWatermark) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ClipNewResponseEnvelope struct {
	Errors   []ClipNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ClipNewResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamClipping                    `json:"result,required"`
	// Whether the API call was successful
	Success ClipNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    clipNewResponseEnvelopeJSON    `json:"-"`
}

// clipNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ClipNewResponseEnvelope]
type clipNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clipNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ClipNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    clipNewResponseEnvelopeErrorsJSON `json:"-"`
}

// clipNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ClipNewResponseEnvelopeErrors]
type clipNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clipNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ClipNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    clipNewResponseEnvelopeMessagesJSON `json:"-"`
}

// clipNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [ClipNewResponseEnvelopeMessages]
type clipNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClipNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clipNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ClipNewResponseEnvelopeSuccess bool

const (
	ClipNewResponseEnvelopeSuccessTrue ClipNewResponseEnvelopeSuccess = true
)

func (r ClipNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ClipNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}