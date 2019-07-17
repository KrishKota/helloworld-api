package routes

import (
	"fmt"
	"helloworld-api/internal/model"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

func PrintMessage(res http.ResponseWriter, req *http.Request) {

	var requestParams model.Request
	err := schema.NewDecoder().Decode(&requestParams, req.URL.Query())
	if err != nil {
		errMsg := fmt.Sprintf("failed to decode request, %v", err)
		log.Println(errMsg)
		res.Write([]byte(errMsg))
		res.WriteHeader(400)
	}
	response := fmt.Sprintf("Hi, %s %s!", requestParams.FirstName, requestParams.LastName)
	res.Write([]byte(response))
	res.WriteHeader(200)
}
