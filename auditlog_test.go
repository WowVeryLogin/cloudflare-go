// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestAuditLogAuditLogsGetAccountAuditLogsWithOptionalParams(t *testing.T) {
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
	_, err := client.AuditLogs.AuditLogsGetAccountAuditLogs(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.AuditLogAuditLogsGetAccountAuditLogsParams{
			ID: cloudflare.F("f174be97-19b1-40d6-954d-70cd5fbd52db"),
			Action: cloudflare.F(cloudflare.AuditLogAuditLogsGetAccountAuditLogsParamsAction{
				Type: cloudflare.F("add"),
			}),
			Actor: cloudflare.F(cloudflare.AuditLogAuditLogsGetAccountAuditLogsParamsActor{
				IP:    cloudflare.F("17.168.228.63"),
				Email: cloudflare.F("alice@example.com"),
			}),
			Before:       cloudflare.F(time.Now()),
			Direction:    cloudflare.F(cloudflare.AuditLogAuditLogsGetAccountAuditLogsParamsDirectionDesc),
			Export:       cloudflare.F(true),
			HideUserLogs: cloudflare.F(true),
			Page:         cloudflare.F(50.000000),
			PerPage:      cloudflare.F(25.000000),
			Since:        cloudflare.F(time.Now()),
			Zone: cloudflare.F(cloudflare.AuditLogAuditLogsGetAccountAuditLogsParamsZone{
				Name: cloudflare.F("example.com"),
			}),
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