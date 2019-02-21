package web

import (
	"fmt"
	"net/http"

	"github.com/carbocation/interpose"
	gorilla_mux "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/zloy2005/webshop/internal/service/payment"
	"github.com/zloy2005/webshop/internal/storage"
	"github.com/zloy2005/webshop/web/handler"
)

type Service struct {
	storage *storage.Storage
	payment *payment.Payment
	server  *http.Server
	config  Config
}

func New(storage *storage.Storage, payment *payment.Payment, config Config) *Service {
	serverAddress := fmt.Sprintf(":%s", config.Port)
	server := &http.Server{Addr: serverAddress, Handler: handlers(storage, payment)}
	return &Service{
		storage: storage,
		payment: payment,
		config:  config,
		server:  server,
	}
}

func (s *Service) ListenAndServe() error {
	logrus.Infoln("running HTTP server on " + s.server.Addr)
	return s.server.ListenAndServe()
}

func handlers(storage *storage.Storage, payment *payment.Payment) *interpose.Middleware {
	middle := interpose.New()
	homeHandler := handler.NewHomeHandler(storage)
	paymentHandler := handler.NewPaymentHandler(storage, payment)
	ordersHandler := handler.NewOrderHandler(storage, payment)

	router := gorilla_mux.NewRouter()
	router.HandleFunc("/", homeHandler.GetHome).Methods("GET")
	router.HandleFunc("/payment", paymentHandler.NewPayment).Methods("GET")
	router.HandleFunc("/payment", paymentHandler.NewPaymentOrder).Methods("POST")
	router.HandleFunc("/orders", ordersHandler.Orders).Methods("GET")
	router.HandleFunc("/orders/refund", ordersHandler.OrderRefund).Methods("POST")
	// Path of static files must be last!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	middle.UseHandler(router)
	return middle
}
