// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkersForPlatformDispatchNamespaceScriptService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformDispatchNamespaceScriptService] method instead.
type WorkersForPlatformDispatchNamespaceScriptService struct {
	Options  []option.RequestOption
	Content  *WorkersForPlatformDispatchNamespaceScriptContentService
	Settings *WorkersForPlatformDispatchNamespaceScriptSettingService
}

// NewWorkersForPlatformDispatchNamespaceScriptService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWorkersForPlatformDispatchNamespaceScriptService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceScriptService) {
	r = &WorkersForPlatformDispatchNamespaceScriptService{}
	r.Options = opts
	r.Content = NewWorkersForPlatformDispatchNamespaceScriptContentService(opts...)
	r.Settings = NewWorkersForPlatformDispatchNamespaceScriptSettingService(opts...)
	return
}