// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSTsigService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSTsigService] method
// instead.
type SecondaryDNSTsigService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSTsigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSTsigService(opts ...option.RequestOption) (r *SecondaryDNSTsigService) {
	r = &SecondaryDNSTsigService{}
	r.Options = opts
	return
}