package utils

import (
	"api_ticket/models"
	"reflect"
	"regexp"
	"strings"
)

func ResponseError(status string, err interface{}, code int) models.MsgRes {
	return models.MsgRes{
		StatusCode: code,
		Status:     status,
		Error:      err,
		Data:       nil,
	}
}

func ResponseSuccess(status string, err interface{}, data interface{}, code int) models.MsgRes {
	return models.MsgRes{
		StatusCode: code,
		Status:     status,
		Error:      err,
		Data:       data,
	}
}

func ValidatorJsonName(data map[string]interface{}, inputModels interface{}) models.MsgRes {
	rt := reflect.TypeOf(inputModels)
	keyku := reflect.ValueOf(data).MapKeys()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		v := strings.Split(f.Tag.Get("json"), ",")[0] // use split to ignore tag "options"
		if !Search(v, keyku) {
			return ResponseError("wrong json key", nil, 404)
		}
	}
	return ResponseSuccess("ok", nil, nil, 200)
}

func Search(str string, key []reflect.Value) bool {
	for i := 0; i < len(key); i++ {
		if str == key[i].String() {
			return true
		}
	}
	return false
}

func ValidateVal(value models.Ticket) models.MsgRes {
	if value.Name == "" || value.Seat == "" || value.Studio == "" || value.Title_movie == "" {
		return ResponseError("there's empty value of request body json", nil, 404)
	}
	RegexSeat := "^([0-2][0-9]|[1-9])[A-Z]$"
	RegexStudio := "^[1-5]$"
	if b := regexp.MustCompile(RegexSeat).MatchString(value.Seat); b != true {
		return ResponseError("seat number must be around [1-29][A-Z]", nil, 404)
	}
	if c := regexp.MustCompile(RegexStudio).MatchString(value.Studio); c != true {
		return ResponseError("Studio number must be around [1-5]", nil, 404)
	}
	return ResponseSuccess("ok", nil, nil, 200)
}
