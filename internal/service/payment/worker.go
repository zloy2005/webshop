package payment

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zloy2005/webshop/internal/storage"
)

type Worker struct {
	storage *storage.Storage
	payment *Payment
}

func NewWorker(storage *storage.Storage, payment *Payment) *Worker {
	return &Worker{storage: storage, payment: payment}
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				orders, _ := w.storage.PaymentOrdersInProccess()
				for _, o := range orders {
					if _, e := w.payment.Status(o); e != nil {
						logrus.Error(e)
					}

				}
			}
		}
	}()
}
