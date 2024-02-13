// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestHealthcheckUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Healthchecks.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.HealthcheckUpdateParams{
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]cloudflare.HealthcheckUpdateParamsCheckRegion{cloudflare.HealthcheckUpdateParamsCheckRegionWeu, cloudflare.HealthcheckUpdateParamsCheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(cloudflare.HealthcheckUpdateParamsHTTPConfig{
				AllowInsecure:   cloudflare.F(true),
				ExpectedBody:    cloudflare.F("success"),
				ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
				FollowRedirects: cloudflare.F(true),
				Header: cloudflare.F[any](map[string]interface{}{
					"Host": map[string]interface{}{
						"0": "example.com",
					},
					"X-App-ID": map[string]interface{}{
						"0": "abc123",
					},
				}),
				Method: cloudflare.F(cloudflare.HealthcheckUpdateParamsHTTPConfigMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TcpConfig: cloudflare.F(cloudflare.HealthcheckUpdateParamsTcpConfig{
				Method: cloudflare.F(cloudflare.HealthcheckUpdateParamsTcpConfigMethodConnectionEstablished),
				Port:   cloudflare.F(int64(0)),
			}),
			Timeout: cloudflare.F(int64(0)),
			Type:    cloudflare.F("HTTPS"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Healthchecks.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Healthchecks.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"023e105f4ecef8ad9ca31a8372d0c353",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckHealthChecksNewHealthCheckWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Healthchecks.HealthChecksNewHealthCheck(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.HealthcheckHealthChecksNewHealthCheckParams{
			Address:              cloudflare.F("www.example.com"),
			Name:                 cloudflare.F("server-1"),
			CheckRegions:         cloudflare.F([]cloudflare.HealthcheckHealthChecksNewHealthCheckParamsCheckRegion{cloudflare.HealthcheckHealthChecksNewHealthCheckParamsCheckRegionWeu, cloudflare.HealthcheckHealthChecksNewHealthCheckParamsCheckRegionEnam}),
			ConsecutiveFails:     cloudflare.F(int64(0)),
			ConsecutiveSuccesses: cloudflare.F(int64(0)),
			Description:          cloudflare.F("Health check for www.example.com"),
			HTTPConfig: cloudflare.F(cloudflare.HealthcheckHealthChecksNewHealthCheckParamsHTTPConfig{
				AllowInsecure:   cloudflare.F(true),
				ExpectedBody:    cloudflare.F("success"),
				ExpectedCodes:   cloudflare.F([]string{"2xx", "302"}),
				FollowRedirects: cloudflare.F(true),
				Header: cloudflare.F[any](map[string]interface{}{
					"Host": map[string]interface{}{
						"0": "example.com",
					},
					"X-App-ID": map[string]interface{}{
						"0": "abc123",
					},
				}),
				Method: cloudflare.F(cloudflare.HealthcheckHealthChecksNewHealthCheckParamsHTTPConfigMethodGet),
				Path:   cloudflare.F("/health"),
				Port:   cloudflare.F(int64(0)),
			}),
			Interval:  cloudflare.F(int64(0)),
			Retries:   cloudflare.F(int64(0)),
			Suspended: cloudflare.F(true),
			TcpConfig: cloudflare.F(cloudflare.HealthcheckHealthChecksNewHealthCheckParamsTcpConfig{
				Method: cloudflare.F(cloudflare.HealthcheckHealthChecksNewHealthCheckParamsTcpConfigMethodConnectionEstablished),
				Port:   cloudflare.F(int64(0)),
			}),
			Timeout: cloudflare.F(int64(0)),
			Type:    cloudflare.F("HTTPS"),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHealthcheckHealthChecksListHealthChecks(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Healthchecks.HealthChecksListHealthChecks(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}