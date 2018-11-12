package server

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"ui-mockup-backend"
	"ui-mockup-backend/mongo"
)

type standardRouter struct {
	standardService root.StandardService
	auth *authHelper
}

func NewStandardRouter(u root.StandardService, router *mux.Router, a *authHelper) *mux.Router {
	standardRouter := standardRouter{u,a}
	//router.HandleFunc("/load_standards", a.validate(standardRouter.loadStandardHandler)).Methods("GET")
	router.HandleFunc("/load_standards", standardRouter.loadStandardHandler).Methods("GET")
	router.HandleFunc("/get_standard/{standardName}", standardRouter.getStandardHandler).Methods("GET")
	return router
}

func(ur *standardRouter) loadStandardHandler(w http.ResponseWriter, r *http.Request) {	
	err, std := LoadStandards()
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}

func(ur *standardRouter) getStandardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	standardName := vars["standardName"]

	err, std := ur.standardService.GetStandardInfo(standardName)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	Json(w, http.StatusOK, std)
}


func LoadStandards() (error, string){

	//print("LOADING STANDARDS")

	standardsYamlFile, err := ioutil.ReadFile("/home/mukul/git/standards/nist-800-53-latest.yaml")
	if err != nil {
		log.Printf("standardsYamlFile.Get err   #%v ", err)
	}
	standardsJson, err := yaml.YAMLToJSON(standardsYamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err, "nist-800-53-latest"
	}

	//print(standardsJson)

	var standardsResult map[string]interface{}
	json.Unmarshal([]byte(standardsJson), &standardsResult)

	//var controls[] root.Controls
	controls := []root.Controls{}
	i := 0
	for key, value := range standardsResult {
		// Each value is an interface{} type, that is type asserted as a string

		var desc, family, name string
		for k, v := range value.(map[string]interface{}) {
			if k == "family" {
				family = v.(string)
			}
			if k == "name" {
				name = v.(string)
			}
			if k == "description" {
				desc = v.(string)
			}
		}

		//controlInfo := ControlInfo{ Family:family, Name:name, Description:desc }
		//standard := Standards{ControlInfo: controlInfo, ControlName:key}
		//controlInfo := root.Controls{ Family:family, Name:name, Description:desc }
		controlInfo := root.ControlInfo{ Family:family, Name:name, Description:desc }
		//print(controlInfo)
		//controls[i] = root.Controls{ ControlName: key , ControlInfo: controlInfo }
		controls = append(controls, root.Controls{ ControlName: key , ControlInfo: controlInfo })
		i += 1
		// todo: Replace with standard name from file name
		standard := root.Standard{StandardName:"nist-800-53-latest", Controls: controls}
		//fmt.Print(standard)
		// TODO: insert every standard into DB
		standardService := new(mongo.StandardsService)
		standardService.CreateStandard(&standard)
		//fmt.Println(standard)
		break // TODO: remove after test
	}
	// todo: Replace with standard name from file name
	return err, "nist-800-53-latest"

}


