package services

import (
	"context"
	"fmt"
	"time"

	"github.com/kopicee/hdm-go/lib/model"
	"github.com/kopicee/hdm-go/lib/repository"
	"github.com/kopicee/hdm-go/lib/services/merge"
	"github.com/kopicee/hdm-go/lib/services/normalize"
	"github.com/kopicee/hdm-go/lib/suppliers"
	"golang.org/x/exp/slog"
)

type ingester struct{ repo repository.HotelsRepository }

func (i ingester) Ingest() error {
	ctx := context.Background()
	timestamp := time.Now()

	hotels, errs := suppliers.FetchAllSuppliers(ctx)
	if len(errs) != 0 {
		slog.ErrorContext(ctx, fmt.Sprintf("received errors while fetching from suppliers: %+v", errs))
		// Don't return here, because maybe some suppliers didn't error out
	}

	for id, hotels := range groupByID(hotels) {
		record, err := i.repo.FindOne(id)
		if err != nil {
			return err
		}

		if record == nil {
			record = &model.Hotel{
				CreatedAt: timestamp,
			}
		}

		if err := combine(record, hotels, timestamp); err != nil {
			return err
		}

		if err = i.repo.Save(record); err != nil {
			return err
		}
	}
	return nil
}

func groupByID(hotels []*model.Hotel) map[string][]*model.Hotel {
	grouped := make(map[string][]*model.Hotel)

	for _, hotel := range hotels {
		id := hotel.ID

		if _, ok := grouped[id]; ok {
			grouped[id] = make([]*model.Hotel, 0)
		}
		grouped[id] = append(grouped[id], hotel)
	}

	return grouped
}

func combine(record *model.Hotel, hotels []*model.Hotel, timestamp time.Time) error {
	if err := merge.Merge(record, hotels); err != nil {
		return err
	}
	if err := normalize.Normalize(record); err != nil {
		return err
	}

	record.UpdatedAt = timestamp
	return nil
}
