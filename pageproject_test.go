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

func TestPageProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pages.Projects.New(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		cloudflare.PageProjectNewParams{
			BuildConfig: cloudflare.F(cloudflare.PageProjectNewParamsBuildConfig{
				BuildCaching:      cloudflare.F(true),
				BuildCommand:      cloudflare.F("npm run build"),
				DestinationDir:    cloudflare.F("build"),
				RootDir:           cloudflare.F("/"),
				WebAnalyticsTag:   cloudflare.F("cee1c73f6e4743d0b5e6bb1a0bcaabcc"),
				WebAnalyticsToken: cloudflare.F("021e1057c18547eca7b79f2516f06o7x"),
			}),
			CanonicalDeployment: cloudflare.F(cloudflare.PageProjectNewParamsCanonicalDeployment{}),
			DeploymentConfigs: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigs{
				Preview: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreview{
					AIBindings: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewAIBindings{
						AIBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewAIBindingsAIBinding{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewD1Databases{
						D1Binding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespaces{
						DoBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewEnvVars{
						EnvironmentVariable: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					KvNamespaces: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewKvNamespaces{
						KvBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewKvNamespacesKvBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					Placement: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewQueueProducers{
						QueueProducerBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewR2Buckets{
						R2Binding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewR2BucketsR2Binding{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					ServiceBindings: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewServiceBindings{
						ServiceBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsPreviewServiceBindingsServiceBinding{
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
				}),
				Production: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProduction{
					AIBindings: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionAIBindings{
						AIBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionAIBindingsAIBinding{
							ProjectID: cloudflare.F[any](map[string]interface{}{}),
						}),
					}),
					AnalyticsEngineDatasets: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasets{
						AnalyticsEngineBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionAnalyticsEngineDatasetsAnalyticsEngineBinding{
							Dataset: cloudflare.F("api_analytics"),
						}),
					}),
					CompatibilityDate:  cloudflare.F("2022-01-01"),
					CompatibilityFlags: cloudflare.F([]interface{}{"url_standard"}),
					D1Databases: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionD1Databases{
						D1Binding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionD1DatabasesD1Binding{
							ID: cloudflare.F("445e2955-951a-43f8-a35b-a4d0c8138f63"),
						}),
					}),
					DurableObjectNamespaces: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespaces{
						DoBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionDurableObjectNamespacesDoBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					EnvVars: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionEnvVars{
						EnvironmentVariable: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariable{
							Type:  cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionEnvVarsEnvironmentVariableTypePlainText),
							Value: cloudflare.F("hello world"),
						}),
					}),
					KvNamespaces: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionKvNamespaces{
						KvBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionKvNamespacesKvBinding{
							NamespaceID: cloudflare.F("5eb63bbbe01eeed093cb22bb8f5acdc3"),
						}),
					}),
					Placement: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionPlacement{
						Mode: cloudflare.F("smart"),
					}),
					QueueProducers: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionQueueProducers{
						QueueProducerBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionQueueProducersQueueProducerBinding{
							Name: cloudflare.F("some-queue"),
						}),
					}),
					R2Buckets: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionR2Buckets{
						R2Binding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionR2BucketsR2Binding{
							Name: cloudflare.F("some-bucket"),
						}),
					}),
					ServiceBindings: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionServiceBindings{
						ServiceBinding: cloudflare.F(cloudflare.PageProjectNewParamsDeploymentConfigsProductionServiceBindingsServiceBinding{
							Environment: cloudflare.F("production"),
							Service:     cloudflare.F("example-worker"),
						}),
					}),
				}),
			}),
			LatestDeployment: cloudflare.F(cloudflare.PageProjectNewParamsLatestDeployment{}),
			Name:             cloudflare.F("NextJS Blog"),
			ProductionBranch: cloudflare.F("main"),
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

func TestPageProjectUpdate(t *testing.T) {
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
	_, err := client.Pages.Projects.Update(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
		cloudflare.PageProjectUpdateParams{
			Body: cloudflare.F[any](map[string]interface{}{
				"deployment_configs": map[string]interface{}{
					"production": map[string]interface{}{
						"compatibility_date": "2022-01-01",
						"compatibility_flags": map[string]interface{}{
							"0": "url_standard",
						},
						"env_vars": map[string]interface{}{
							"BUILD_VERSION": map[string]interface{}{
								"value": "3.3",
							},
							"delete_this_env_var": nil,
							"secret_var": map[string]interface{}{
								"type":  "secret_text",
								"value": "A_CMS_API_TOKEN",
							},
						},
					},
				},
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

func TestPageProjectList(t *testing.T) {
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
	_, err := client.Pages.Projects.List(context.TODO(), "023e105f4ecef8ad9ca31a8372d0c353")
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageProjectDelete(t *testing.T) {
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
	_, err := client.Pages.Projects.Delete(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageProjectGet(t *testing.T) {
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
	_, err := client.Pages.Projects.Get(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPageProjectPurgeBuildCache(t *testing.T) {
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
	_, err := client.Pages.Projects.PurgeBuildCache(
		context.TODO(),
		"023e105f4ecef8ad9ca31a8372d0c353",
		"this-is-my-project-01",
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}