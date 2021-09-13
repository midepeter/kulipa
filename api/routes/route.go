package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/midepeter/gopay-app/api/utils"
	"github.com/midepeter/gopay-app/store/models"
)

type Server struct{}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is the hello world page of this breakage in the rare of the space"))
}

func (s *Server) Proceed(w http.ResponseWriter, r *http.Request) {
	var student models.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &student)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_ = utils.Render(w, http.StatusOK, &student)
}

func (s *Server) ProceedToPay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is handler for you to proceed to pay through flutterwave"))
}
