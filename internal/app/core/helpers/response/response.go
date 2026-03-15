package response

const OK = "OK"
const Created = "Created"
const Updated = "Updated"
const Deleted = "Deleted"
const Unauthorized = "Unauthorized"
const ServerError = "Server Error"
const BadRequest = "Bad Request"

func SuccessResponse(data interface{}, message string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": message,
		"data":    data,
	}
}

func ErrorResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
		"data":    nil,
	}
}
