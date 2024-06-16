package suppliers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kopicee/hdm-go/lib/model"
	"golang.org/x/exp/slog"
)

type Supplier interface {
	Name() string
	Fetch(context.Context) ([]*model.Hotel, error)
}

var transport = http.DefaultTransport

var suppliers = []Supplier{
	acmeSupplier{transport},
	patagoniaSupplier{transport},
	paperfliesSupplier{transport},
}

func FetchAllSuppliers(ctx context.Context) ([]*model.Hotel, []error) {
	hotels := make([]*model.Hotel, 0)
	errs := make([]error, 0)

	for _, supplier := range suppliers {
		logger := slog.With("supplier", supplier.Name())

		results, err := supplier.Fetch(ctx)
		hotels = append(hotels, results...)
		if err != nil {
			errs = append(errs, err)
			logger.ErrorContext(ctx, "Got errors while fetching data.", "err", err)
		}

		logger.InfoContext(ctx, fmt.Sprintf("Fetched %d hotels.", len(results)))
	}

	return hotels, errs
}
