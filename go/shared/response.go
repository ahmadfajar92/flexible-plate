package shared

type (

	// Error struct
	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	// Response struct
	Response struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    interface{} `json:"meta,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	// JSONMeta struct
	JSONMeta struct {
		Page         int `json:"page"`
		Limit        int `json:"limit"`
		TotalRecords int `json:"totalRecords"`
		TotalPages   int `json:"totalPages"`
	}
)

// JSONResponse func response
func JSONResponse(code int, message string, status bool, params ...interface{}) *Response {

	response := new(Response)

	for _, param := range params {
		switch param.(type) {
		case JSONMeta:
			response.Meta = param
		case []*Error, Error:
			response.Errors = param
		default:
			response.Data = param
		}
	}

	response.Success = status
	response.Code = code
	response.Message = message

	return response
}
