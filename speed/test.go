// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package speed

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// TestService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTestService] method instead.
type TestService struct {
	Options []option.RequestOption
}

// NewTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestService(opts ...option.RequestOption) (r *TestService) {
	r = &TestService{}
	r.Options = opts
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *TestService) New(ctx context.Context, url string, params TestNewParams, opts ...option.RequestOption) (res *Test, err error) {
	opts = append(r.Options[:], opts...)
	var env TestNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Test history (list of tests) for a specific webpage.
func (r *TestService) List(ctx context.Context, url string, params TestListParams, opts ...option.RequestOption) (res *TestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *TestService) Delete(ctx context.Context, url string, params TestDeleteParams, opts ...option.RequestOption) (res *TestDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TestDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", params.ZoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the result of a specific test.
func (r *TestService) Get(ctx context.Context, url string, testID string, query TestGetParams, opts ...option.RequestOption) (res *Test, err error) {
	opts = append(r.Options[:], opts...)
	var env TestGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", query.ZoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Test struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport LighthouseReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport LighthouseReport `json:"mobileReport"`
	// A test region with a label.
	Region LabeledRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency TestScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string   `json:"url"`
	JSON testJSON `json:"-"`
}

// testJSON contains the JSON metadata for the struct [Test]
type testJSON struct {
	ID                apijson.Field
	Date              apijson.Field
	DesktopReport     apijson.Field
	MobileReport      apijson.Field
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Test) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testJSON) RawJSON() string {
	return r.raw
}

// The frequency of the test.
type TestScheduleFrequency string

const (
	TestScheduleFrequencyDaily  TestScheduleFrequency = "DAILY"
	TestScheduleFrequencyWeekly TestScheduleFrequency = "WEEKLY"
)

func (r TestScheduleFrequency) IsKnown() bool {
	switch r {
	case TestScheduleFrequencyDaily, TestScheduleFrequencyWeekly:
		return true
	}
	return false
}

type TestListResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                       `json:"success,required"`
	ResultInfo TestListResponseResultInfo `json:"result_info"`
	JSON       testListResponseJSON       `json:"-"`
}

// testListResponseJSON contains the JSON metadata for the struct
// [TestListResponse]
type testListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseJSON) RawJSON() string {
	return r.raw
}

type TestListResponseResultInfo struct {
	Count      int64                          `json:"count"`
	Page       int64                          `json:"page"`
	PerPage    int64                          `json:"per_page"`
	TotalCount int64                          `json:"total_count"`
	JSON       testListResponseResultInfoJSON `json:"-"`
}

// testListResponseResultInfoJSON contains the JSON metadata for the struct
// [TestListResponseResultInfo]
type testListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type TestDeleteResponse struct {
	// Number of items affected.
	Count float64                `json:"count"`
	JSON  testDeleteResponseJSON `json:"-"`
}

// testDeleteResponseJSON contains the JSON metadata for the struct
// [TestDeleteResponse]
type testDeleteResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type TestNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[TestNewParamsRegion] `json:"region"`
}

func (r TestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A test region.
type TestNewParamsRegion string

const (
	TestNewParamsRegionAsiaEast1           TestNewParamsRegion = "asia-east1"
	TestNewParamsRegionAsiaNortheast1      TestNewParamsRegion = "asia-northeast1"
	TestNewParamsRegionAsiaNortheast2      TestNewParamsRegion = "asia-northeast2"
	TestNewParamsRegionAsiaSouth1          TestNewParamsRegion = "asia-south1"
	TestNewParamsRegionAsiaSoutheast1      TestNewParamsRegion = "asia-southeast1"
	TestNewParamsRegionAustraliaSoutheast1 TestNewParamsRegion = "australia-southeast1"
	TestNewParamsRegionEuropeNorth1        TestNewParamsRegion = "europe-north1"
	TestNewParamsRegionEuropeSouthwest1    TestNewParamsRegion = "europe-southwest1"
	TestNewParamsRegionEuropeWest1         TestNewParamsRegion = "europe-west1"
	TestNewParamsRegionEuropeWest2         TestNewParamsRegion = "europe-west2"
	TestNewParamsRegionEuropeWest3         TestNewParamsRegion = "europe-west3"
	TestNewParamsRegionEuropeWest4         TestNewParamsRegion = "europe-west4"
	TestNewParamsRegionEuropeWest8         TestNewParamsRegion = "europe-west8"
	TestNewParamsRegionEuropeWest9         TestNewParamsRegion = "europe-west9"
	TestNewParamsRegionMeWest1             TestNewParamsRegion = "me-west1"
	TestNewParamsRegionSouthamericaEast1   TestNewParamsRegion = "southamerica-east1"
	TestNewParamsRegionUsCentral1          TestNewParamsRegion = "us-central1"
	TestNewParamsRegionUsEast1             TestNewParamsRegion = "us-east1"
	TestNewParamsRegionUsEast4             TestNewParamsRegion = "us-east4"
	TestNewParamsRegionUsSouth1            TestNewParamsRegion = "us-south1"
	TestNewParamsRegionUsWest1             TestNewParamsRegion = "us-west1"
)

func (r TestNewParamsRegion) IsKnown() bool {
	switch r {
	case TestNewParamsRegionAsiaEast1, TestNewParamsRegionAsiaNortheast1, TestNewParamsRegionAsiaNortheast2, TestNewParamsRegionAsiaSouth1, TestNewParamsRegionAsiaSoutheast1, TestNewParamsRegionAustraliaSoutheast1, TestNewParamsRegionEuropeNorth1, TestNewParamsRegionEuropeSouthwest1, TestNewParamsRegionEuropeWest1, TestNewParamsRegionEuropeWest2, TestNewParamsRegionEuropeWest3, TestNewParamsRegionEuropeWest4, TestNewParamsRegionEuropeWest8, TestNewParamsRegionEuropeWest9, TestNewParamsRegionMeWest1, TestNewParamsRegionSouthamericaEast1, TestNewParamsRegionUsCentral1, TestNewParamsRegionUsEast1, TestNewParamsRegionUsEast4, TestNewParamsRegionUsSouth1, TestNewParamsRegionUsWest1:
		return true
	}
	return false
}

type TestNewResponseEnvelope struct {
	Errors   interface{}                 `json:"errors,required"`
	Messages interface{}                 `json:"messages,required"`
	Success  interface{}                 `json:"success,required"`
	Result   Test                        `json:"result"`
	JSON     testNewResponseEnvelopeJSON `json:"-"`
}

// testNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestNewResponseEnvelope]
type testNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TestListParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Page    param.Field[int64]  `query:"page"`
	PerPage param.Field[int64]  `query:"per_page"`
	// A test region.
	Region param.Field[TestListParamsRegion] `query:"region"`
}

// URLQuery serializes [TestListParams]'s query parameters as `url.Values`.
func (r TestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type TestListParamsRegion string

const (
	TestListParamsRegionAsiaEast1           TestListParamsRegion = "asia-east1"
	TestListParamsRegionAsiaNortheast1      TestListParamsRegion = "asia-northeast1"
	TestListParamsRegionAsiaNortheast2      TestListParamsRegion = "asia-northeast2"
	TestListParamsRegionAsiaSouth1          TestListParamsRegion = "asia-south1"
	TestListParamsRegionAsiaSoutheast1      TestListParamsRegion = "asia-southeast1"
	TestListParamsRegionAustraliaSoutheast1 TestListParamsRegion = "australia-southeast1"
	TestListParamsRegionEuropeNorth1        TestListParamsRegion = "europe-north1"
	TestListParamsRegionEuropeSouthwest1    TestListParamsRegion = "europe-southwest1"
	TestListParamsRegionEuropeWest1         TestListParamsRegion = "europe-west1"
	TestListParamsRegionEuropeWest2         TestListParamsRegion = "europe-west2"
	TestListParamsRegionEuropeWest3         TestListParamsRegion = "europe-west3"
	TestListParamsRegionEuropeWest4         TestListParamsRegion = "europe-west4"
	TestListParamsRegionEuropeWest8         TestListParamsRegion = "europe-west8"
	TestListParamsRegionEuropeWest9         TestListParamsRegion = "europe-west9"
	TestListParamsRegionMeWest1             TestListParamsRegion = "me-west1"
	TestListParamsRegionSouthamericaEast1   TestListParamsRegion = "southamerica-east1"
	TestListParamsRegionUsCentral1          TestListParamsRegion = "us-central1"
	TestListParamsRegionUsEast1             TestListParamsRegion = "us-east1"
	TestListParamsRegionUsEast4             TestListParamsRegion = "us-east4"
	TestListParamsRegionUsSouth1            TestListParamsRegion = "us-south1"
	TestListParamsRegionUsWest1             TestListParamsRegion = "us-west1"
)

func (r TestListParamsRegion) IsKnown() bool {
	switch r {
	case TestListParamsRegionAsiaEast1, TestListParamsRegionAsiaNortheast1, TestListParamsRegionAsiaNortheast2, TestListParamsRegionAsiaSouth1, TestListParamsRegionAsiaSoutheast1, TestListParamsRegionAustraliaSoutheast1, TestListParamsRegionEuropeNorth1, TestListParamsRegionEuropeSouthwest1, TestListParamsRegionEuropeWest1, TestListParamsRegionEuropeWest2, TestListParamsRegionEuropeWest3, TestListParamsRegionEuropeWest4, TestListParamsRegionEuropeWest8, TestListParamsRegionEuropeWest9, TestListParamsRegionMeWest1, TestListParamsRegionSouthamericaEast1, TestListParamsRegionUsCentral1, TestListParamsRegionUsEast1, TestListParamsRegionUsEast4, TestListParamsRegionUsSouth1, TestListParamsRegionUsWest1:
		return true
	}
	return false
}

type TestDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A test region.
	Region param.Field[TestDeleteParamsRegion] `query:"region"`
}

// URLQuery serializes [TestDeleteParams]'s query parameters as `url.Values`.
func (r TestDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type TestDeleteParamsRegion string

const (
	TestDeleteParamsRegionAsiaEast1           TestDeleteParamsRegion = "asia-east1"
	TestDeleteParamsRegionAsiaNortheast1      TestDeleteParamsRegion = "asia-northeast1"
	TestDeleteParamsRegionAsiaNortheast2      TestDeleteParamsRegion = "asia-northeast2"
	TestDeleteParamsRegionAsiaSouth1          TestDeleteParamsRegion = "asia-south1"
	TestDeleteParamsRegionAsiaSoutheast1      TestDeleteParamsRegion = "asia-southeast1"
	TestDeleteParamsRegionAustraliaSoutheast1 TestDeleteParamsRegion = "australia-southeast1"
	TestDeleteParamsRegionEuropeNorth1        TestDeleteParamsRegion = "europe-north1"
	TestDeleteParamsRegionEuropeSouthwest1    TestDeleteParamsRegion = "europe-southwest1"
	TestDeleteParamsRegionEuropeWest1         TestDeleteParamsRegion = "europe-west1"
	TestDeleteParamsRegionEuropeWest2         TestDeleteParamsRegion = "europe-west2"
	TestDeleteParamsRegionEuropeWest3         TestDeleteParamsRegion = "europe-west3"
	TestDeleteParamsRegionEuropeWest4         TestDeleteParamsRegion = "europe-west4"
	TestDeleteParamsRegionEuropeWest8         TestDeleteParamsRegion = "europe-west8"
	TestDeleteParamsRegionEuropeWest9         TestDeleteParamsRegion = "europe-west9"
	TestDeleteParamsRegionMeWest1             TestDeleteParamsRegion = "me-west1"
	TestDeleteParamsRegionSouthamericaEast1   TestDeleteParamsRegion = "southamerica-east1"
	TestDeleteParamsRegionUsCentral1          TestDeleteParamsRegion = "us-central1"
	TestDeleteParamsRegionUsEast1             TestDeleteParamsRegion = "us-east1"
	TestDeleteParamsRegionUsEast4             TestDeleteParamsRegion = "us-east4"
	TestDeleteParamsRegionUsSouth1            TestDeleteParamsRegion = "us-south1"
	TestDeleteParamsRegionUsWest1             TestDeleteParamsRegion = "us-west1"
)

func (r TestDeleteParamsRegion) IsKnown() bool {
	switch r {
	case TestDeleteParamsRegionAsiaEast1, TestDeleteParamsRegionAsiaNortheast1, TestDeleteParamsRegionAsiaNortheast2, TestDeleteParamsRegionAsiaSouth1, TestDeleteParamsRegionAsiaSoutheast1, TestDeleteParamsRegionAustraliaSoutheast1, TestDeleteParamsRegionEuropeNorth1, TestDeleteParamsRegionEuropeSouthwest1, TestDeleteParamsRegionEuropeWest1, TestDeleteParamsRegionEuropeWest2, TestDeleteParamsRegionEuropeWest3, TestDeleteParamsRegionEuropeWest4, TestDeleteParamsRegionEuropeWest8, TestDeleteParamsRegionEuropeWest9, TestDeleteParamsRegionMeWest1, TestDeleteParamsRegionSouthamericaEast1, TestDeleteParamsRegionUsCentral1, TestDeleteParamsRegionUsEast1, TestDeleteParamsRegionUsEast4, TestDeleteParamsRegionUsSouth1, TestDeleteParamsRegionUsWest1:
		return true
	}
	return false
}

type TestDeleteResponseEnvelope struct {
	Errors   interface{}                    `json:"errors,required"`
	Messages interface{}                    `json:"messages,required"`
	Success  interface{}                    `json:"success,required"`
	Result   TestDeleteResponse             `json:"result"`
	JSON     testDeleteResponseEnvelopeJSON `json:"-"`
}

// testDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestDeleteResponseEnvelope]
type testDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TestGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TestGetResponseEnvelope struct {
	Errors   interface{}                 `json:"errors,required"`
	Messages interface{}                 `json:"messages,required"`
	Success  interface{}                 `json:"success,required"`
	Result   Test                        `json:"result"`
	JSON     testGetResponseEnvelopeJSON `json:"-"`
}

// testGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TestGetResponseEnvelope]
type testGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}