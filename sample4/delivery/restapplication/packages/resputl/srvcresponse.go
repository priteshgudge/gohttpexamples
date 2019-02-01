package resputl

import (
	"encoding/json"
	"fmt"
	"net/http"

	customerrors "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/errors"
)

//EmptyStruct represents empty strcuture
type EmptyStruct struct {
}

//SrvcRes service response struct
type SrvcRes struct {
	Code     int
	Response interface{}
	Message  string
	Headers  map[string]string
}

func marshalResponse(r interface{}) ([]byte, error) {
	return json.MarshalIndent(r, "", "")
}

//RenderResponse renders json reponse
func (s *SrvcRes) RenderResponse(w http.ResponseWriter) {
	if s.Headers == nil {
		s.Headers = map[string]string{"Content-Type": "application/json",
			"Access-Control-Allow-Headers": "source,x-authorization-token,Content-Type",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*"}
	}
	for h, val := range s.Headers {
		w.Header().Set(h, val)
	}
	var statusBool bool
	switch s.Code {
	case http.StatusOK:
		statusBool = true
	default:
		statusBool = false
	}

	formatted := map[string]interface{}{
		"responseData": s.Response,
		"message":      s.Message,
		"status":       statusBool}

	data, _ := marshalResponse(formatted)
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.WriteHeader(s.Code)
	fmt.Fprint(w, string(data))
}

//Simple200OK ok response with message
func Simple200OK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

//Simple404Response is given if a requested object is not found
func Simple404Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusNotFound, inf, message, nil}
}

//Simple422Response is given if a requested object is not processable
func Simple422Response(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusUnprocessableEntity, inf, message, nil}
}

//PreconditionFailed gives 412 response
func PreconditionFailed(message string) SrvcRes {
	var inf EmptyStruct
	return SrvcRes{http.StatusPreconditionFailed, inf, message, nil}
}

//OptionsResponseOK accepts message
func OptionsResponseOK(message string) SrvcRes {

	var inf EmptyStruct
	return SrvcRes{http.StatusOK, inf, message, nil}
}

//SimpleBadRequest bad request template response
func SimpleBadRequest(message string) SrvcRes {
	return SrvcRes{http.StatusBadRequest, "{}", message, nil}
}

//Response200OK returns json data
func Response200OK(response interface{}) SrvcRes {
	return SrvcRes{http.StatusOK, response, "OK", nil}
}

//ResponseNotImplemented templte for not implemeted response
func ResponseNotImplemented(response interface{}) SrvcRes {
	return SrvcRes{http.StatusNotImplemented, "{}", "Method not implementd", nil}
}

//ReponseCustomError custom error which requires error as input
func ReponseCustomError(err error) SrvcRes {
	//err := errors.
	//return SrvcRes{}
	//var cusErr *customerrors.CustomError
	var inf EmptyStruct
	cusErr := err.(*customerrors.CustomError)
	return SrvcRes{cusErr.GetStatusCode(), inf, cusErr.GetMessage(), nil}
}

//ReponseInternalError template for internal error response
func ReponseInternalError() SrvcRes {
	//err := errors.
	//return SrvcRes{}
	var inf EmptyStruct
	return SrvcRes{http.StatusInternalServerError, inf, "Internal Server Error", nil}
}

//ProcessError template for process error response
func ProcessError(err error, inf interface{}) SrvcRes {

	switch err.(type) {
	case *customerrors.CustomError:
		return ReponseCustomError(err)
	default:
		//go internalErrorEmail(err, inf)
		return ReponseInternalError()

	}
}
