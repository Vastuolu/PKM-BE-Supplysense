package helper

import (
	"net/http"
)

type responseHelp struct{
	Total	int `json:"total"`
	Status  int `json:"status"`
    Message string `json:"message"`
	Error   string `json:"error,omitempty"`
    Data    interface{} `json:"data,omitempty"` 
}

func JsonResponse(code int, data interface{},totalData int, err error) interface{} {
	responseMap := responseHelp{
		Status: code,
		Message: http.StatusText(code),
	}
	if data != nil{
		responseMap.Data = data
		responseMap.Total = totalData
	}else{
		responseMap.Error = err.Error()
	}
	return responseMap
}

func InterfaceMaker(key string, value interface{}) interface{} {
	result := map[string]interface{}{
		key:value,
	}
    return result
}