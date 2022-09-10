package main

import (
	"context"
	"fmt"
	"log"
	"majoo/soal1/configs"
	"majoo/soal1/infrastructure"
	"net/http"

	"github.com/gorilla/mux"

	outletRepository "majoo/soal1/repositories/outlet"
	transactionSummaryRepository "majoo/soal1/repositories/transaction_summary"
	userRepository "majoo/soal1/repositories/user"

	authUsecase "majoo/soal1/usecases/auth"
	transactionSummaryUsecase "majoo/soal1/usecases/transaction_summary"

	authController "majoo/soal1/controllers/auth"
	"majoo/soal1/controllers/middlewares"
	transactionSummaryController "majoo/soal1/controllers/transaction_summary"
)

func main() {
	// setup config
	conf := configs.New()
	ctx := context.Background()

	// setup db
	postgresDB := infrastructure.SetupPostgresDB(conf.Postgres)

	// setup middleware
	jwt := middlewares.SetupJWT(conf.JWT)

	// setup repositories
	userRepo := userRepository.NewPostgreRepository(postgresDB)
	transactionSummaryRepo := transactionSummaryRepository.NewPostgreRepository(postgresDB)
	outletRepo := outletRepository.NewPostgreRepository(postgresDB)

	// setup usecases
	authUC := authUsecase.NewAuthUsecase(jwt, userRepo)
	transactionSummaryUC := transactionSummaryUsecase.NewTransactionSummaryUsecase(transactionSummaryRepo, outletRepo)

	// setup controllers
	authCtrl := authController.NewAuthHTTPController(authUC)
	transactionSummaryCtrl := transactionSummaryController.NewTransactionSummaryHTTPController(transactionSummaryUC)

	// Setup http server
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	v1 := router.PathPrefix("/v1").Subrouter()
	// auth controller
	apiAuth := v1.PathPrefix("/auth").Subrouter()
	apiAuth.HandleFunc("/login", authCtrl.Login).Methods(http.MethodPost)

	// transaction summary controller
	apiSummaryTransaction := v1.PathPrefix("/transactionSummary").Subrouter()
	apiSummaryTransaction.Use(jwt.MuxMiddleware)
	apiSummaryTransaction.HandleFunc("/month", transactionSummaryCtrl.GetByMonth).Methods(http.MethodGet)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.App.Port),
		Handler: router,
	}
	log.Printf("Starting server on port %d \n", conf.App.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(ctx, "%v", err)
	}
}
