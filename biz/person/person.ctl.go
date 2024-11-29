package person

import (
	"net/http"

	"org.donghyuns.com/graph/neo4j/response"
	"org.donghyuns.com/graph/neo4j/utils"
)

func CreatePersonController(res http.ResponseWriter, req *http.Request) {
	var request CreatePersonRequest

	parseErr := utils.DecodeBody(req, &request)

	if parseErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "CRP001",
			Message: "Request Body Parsing Error",
		})
		return
	}

	result := CreateSinglePerson(request.Name, request.Age, request.Languages, request.TechName)

	if result != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusInternalServerError,
			Code:    "CRP002",
			Message: "Create Person Error",
		})
		return
	}

	response.Response(res, response.CommonResponseWithMessage{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
	})
}

func GetPersonController(res http.ResponseWriter, req *http.Request) {
	var request GetPersonRequestByNameRequest

	parseErr := utils.DecodeBody(req, &request)

	if parseErr != nil {
		response.Response(res, response.CommonResponseWithMessage{
			Status:  http.StatusBadRequest,
			Code:    "GTP001",
			Message: "Request Body Parsing Error",
		})
		return
	}

	result, err := GetPerson(request.Name)

	if err != nil {
		response.Response(res, GetPersonResponse{
			Status:  http.StatusInternalServerError,
			Code:    "GTP002",
			Message: "Get Person Error",
		})
		return
	}

	response.Response(res, GetPersonResponse{
		Status:  http.StatusOK,
		Code:    "0000",
		Message: "Success",
		Person:  result,
	})
}
