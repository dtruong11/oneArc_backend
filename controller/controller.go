package controller

import (
	"fmt"
	"net/http"

	"github.com/oneArc_backend/model"
)

func errorsandResponse(res []byte, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		switch err {
		case model.ErrorNotFound:
			w.WriteHeader(http.StatusNotFound)
			fmt.Println("Not found")
		case model.ErrorBadRequest:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Bad request")
		case model.ErrorInternalServer:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Internal Server Error")
		case model.ErrorNotAllowed:
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Println("Failed to authenticate")
		default:
			fmt.Println("Everything goes through.")
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


