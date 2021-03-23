package middlwares

import (
	"github.com/midepeter/gopay-app/store/models"
	"net/http"
)


func ReceiveCandidateDetails(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		
	})
}