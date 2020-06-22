package shared

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
)

// JSONValidate func
func JSONValidate(schema string, data []byte) []*Error {

	errors := make([]*Error, 0)

	rules := gojsonschema.NewStringLoader(schema)
	payload := gojsonschema.NewBytesLoader(data)

	res, err := gojsonschema.Validate(rules, payload)
	if err != nil {
		errors = append(errors, &Error{
			Field:   "keyError",
			Message: err.Error(),
		})

		return errors
	}

	if len(res.Errors()) != 0 {
		for _, d := range res.Errors() {
			errors = append(errors, &Error{
				Field:   d.Details()["field"].(string),
				Message: fmt.Sprintf("%s", d),
			})
		}

		return errors
	}

	return errors
}
