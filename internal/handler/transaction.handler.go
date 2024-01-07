package handler

import (
	"encoding/json"
	"io"
	"m2m/internal/dto"
	"m2m/internal/shared"
	"net/http"
	"strings"
)

type TransactionHandler struct {
	service shared.TransactionService
}

func NewTransactionHandler(service shared.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

func (th *TransactionHandler) GetRoutes() shared.TypeRouters {
	return shared.TypeRouters{
		"/transaction": th.handler,
	}
}

func (th *TransactionHandler) handler(w http.ResponseWriter, r *http.Request) {
	switch strings.ToUpper(r.Method) {
	case "POST":
		th.createTransaction(w, r)
		return
	case "GET":
		th.getTransactions(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (th *TransactionHandler) createTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payloadJson, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	var body dto.TransactionInput
	if err = json.Unmarshal(payloadJson, &body); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	th.service.CreateTransaction(body)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (th *TransactionHandler) getTransactions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
