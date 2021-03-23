package routes

import (
	"net/http"
)

func Proceed(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write("this is a very nice handerler")
}

func ProceedToPay(w http.Response,r *http.Request){
	w.WriteHeader(http.StatusOK)
	W.Write("You can now proceed to pay")
}