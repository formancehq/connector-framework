package registry

import (
	"errors"

	errorsutils "github.com/formancehq/connector-framework/pkg/errors"
	"github.com/formancehq/connector-framework/pkg/httpwrapper"
	"github.com/formancehq/connector-framework/pkg/models"
	"github.com/formancehq/connector-framework/pkg/plugins"
)

func translateError(err error) error {
	switch {
	case errors.Is(err, plugins.ErrNotImplemented):
		return err
	case errors.Is(err, models.ErrMissingFromPayloadInRequest),
		errors.Is(err, models.ErrMissingAccountInRequest),
		errors.Is(err, models.ErrInvalidRequest),
		errors.Is(err, plugins.ErrCurrencyNotSupported),
		errors.Is(err, httpwrapper.ErrStatusCodeClientError),
		errors.Is(err, models.ErrInvalidConfig):
		return errorsutils.NewWrappedError(
			err,
			plugins.ErrInvalidClientRequest,
		)
	case errors.Is(err, httpwrapper.ErrStatusCodeTooManyRequests):
		return errorsutils.NewWrappedError(
			err,
			plugins.ErrUpstreamRatelimit,
		)
	default:
		return err
	}
}
