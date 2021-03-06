package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type eMap map[string]map[string]interface{}

var errorMap = &eMap{}

func addErrorMap(name string, message string, code string) {
	em := *errorMap
	em[name] = map[string]interface{}{
		"message": message,
		"code":    code,
	}
}

func getErrMessage(err validator.FieldError) (string, string, map[string]interface{}) {

	message := "Field validation for failed on the '{{ tag }}' tag"
	code := "a9fce4d8-7275-4727-956a-df4f21a10000"
	values := make(map[string]interface{})
	values["value"] = err.Value()

	em := *errorMap
	if e, ok := em[err.Tag()]; ok {
		code = (e["code"]).(string)
		message = (e["message"]).(string)
		values["value"] = err.Value()
	}

	//fmt.Println(err.Namespace())
	//fmt.Println(err.Field())
	//fmt.Println(err.StructNamespace())
	//fmt.Println(err.StructField())
	//fmt.Println(err.Tag())
	//fmt.Println(err.ActualTag())
	////fmt.Println(err.Kind())
	////fmt.Println(err.Type())
	//fmt.Println(err.Value())
	////fmt.Println(err.Param())

	switch err.Tag() {
	case "lte":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10001", "This value should be less than or equal to {{ compared_value }}."
		values["compared_value"] = err.Param()
	case "required":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10002", "This field is required"
	case "notblank":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10003", "This value should not be blank."
	case "email":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10004", "This value '{{ value }}' is not a valid email address."
	case "max":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10005", "This value is too long. It should have {{ limit }} characters or less."
		values["limit"] = err.Param()
	case "min":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10006", "This value is too short. It should have {{ limit }} characters or more."
		values["limit"] = err.Param()
	case "pass_regex":
		code, message = "a9fce4d8-7275-4727-956a-df4f21a10007", "The password should at least 1 number, at least 1 upper case, at least 1 special character."
	}

	message = strings.Replace(message, "{{ tag }}", err.Tag(), -1)
	for name, r := range values {
		message = strings.Replace(message, "{{ "+name+" }}", (r).(string), -1)
	}

	return code, message, values
}
