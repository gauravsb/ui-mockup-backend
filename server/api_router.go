package server

import (
	"errors"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"ui-mockup-backend"
)

type standardRouter struct {
	standardService root.StandardService
	auth *authHelper
}

func NewStandardRouter(u root.StandardService, router *mux.Router, a *authHelper) *mux.Router {
	standardRouter := standardRouter{u,a}
	router.HandleFunc("/load_standards", a.validate(standardRouter.loadStandardHandler)).Methods("GET")
	router.HandleFunc("/get_standard/{standardName}", standardRouter.getStandardHandler).Methods("GET")
	return router
}

func(ur *standardRouter) loadStandardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	standardName := vars["standardName"]

	err, std := ur.standardService.CreateStandard(standardName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}

func(ur *standardRouter) getStandardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	standardName := vars["standardName"]

	err, std := ur.standardService.GetStandardsInfo(standardName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}

