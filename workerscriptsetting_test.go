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

func TestWorkerScriptSettingGet(t *testing.T) {
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
	_, err := client.WorkerScripts.Settings.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWorkerScriptSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.WorkerScripts.Settings.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is_my_script-01",
		cloudflare.WorkerScriptSettingUpdateParams{
			Settings: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettings{
				Errors: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsError{{
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}, {
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}, {
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}}),
				Messages: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsMessage{{
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}, {
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}, {
					Code:    cloudflare.F(int64(1000)),
					Message: cloudflare.F("string"),
				}}),
				Result: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResult{
					Bindings: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBinding{cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding{
						Type: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace),
					}), cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding{
						Type: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace),
					}), cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBinding{
						Type: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultBindingsWorkersKvNamespaceBindingTypeKvNamespace),
					})}),
					CompatibilityDate:  cloudflare.F("2022-04-05"),
					CompatibilityFlags: cloudflare.F([]string{"formdata_parser_supports_files", "formdata_parser_supports_files", "formdata_parser_supports_files"}),
					Logpush:            cloudflare.F(false),
					Migrations: cloudflare.F[cloudflare.WorkerScriptSettingUpdateParamsSettingsResultMigrations](cloudflare.WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrations{
						NewTag:         cloudflare.F("v2"),
						OldTag:         cloudflare.F("v1"),
						DeletedClasses: cloudflare.F([]string{"string", "string", "string"}),
						NewClasses:     cloudflare.F([]string{"string", "string", "string"}),
						RenamedClasses: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsRenamedClass{{
							From: cloudflare.F("string"),
							To:   cloudflare.F("string"),
						}, {
							From: cloudflare.F("string"),
							To:   cloudflare.F("string"),
						}, {
							From: cloudflare.F("string"),
							To:   cloudflare.F("string"),
						}}),
						TransferredClasses: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsResultMigrationsWorkersSingleStepMigrationsTransferredClass{{
							From:       cloudflare.F("string"),
							FromScript: cloudflare.F("string"),
							To:         cloudflare.F("string"),
						}, {
							From:       cloudflare.F("string"),
							FromScript: cloudflare.F("string"),
							To:         cloudflare.F("string"),
						}, {
							From:       cloudflare.F("string"),
							FromScript: cloudflare.F("string"),
							To:         cloudflare.F("string"),
						}}),
					})),
					Placement: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultPlacement{
						Mode: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsResultPlacementModeSmart),
					}),
					Tags: cloudflare.F([]string{"my-tag", "my-tag", "my-tag"}),
					TailConsumers: cloudflare.F([]cloudflare.WorkerScriptSettingUpdateParamsSettingsResultTailConsumer{{
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
						Service:     cloudflare.F("my-log-consumer"),
					}, {
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
						Service:     cloudflare.F("my-log-consumer"),
					}, {
						Environment: cloudflare.F("production"),
						Namespace:   cloudflare.F("my-namespace"),
						Service:     cloudflare.F("my-log-consumer"),
					}}),
					UsageModel: cloudflare.F("unbound"),
				}),
				Success: cloudflare.F(cloudflare.WorkerScriptSettingUpdateParamsSettingsSuccessTrue),
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