package errors

const (
	UNKNOWN_ERROR          ErrorCode = 0
	DATA_INVALID           ErrorCode = 1001
	FAILED_RETRIEVE_DATA   ErrorCode = 1002
	STATUS_PAGE_NOT_FOUND  ErrorCode = 1003
	UNAUTHORIZED           ErrorCode = 1004
	FAILED_CREATE_DATA     ErrorCode = 1005
	USER_ALREADY_EXIST     ErrorCode = 1006
	FAILED_SENDING_MESSAGE ErrorCode = 1007
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Data not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	UNAUTHORIZED: {
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
	USER_ALREADY_EXIST: {
		ClientMessage: "Email Already Exist.",
		SystemMessage: "Email Already Exist.",
		ErrorCode:     USER_ALREADY_EXIST,
	},
	FAILED_SENDING_MESSAGE: {
		ClientMessage: "message_cant_be_send",
		SystemMessage: "message_cant_be_send.",
		ErrorCode:     FAILED_SENDING_MESSAGE,
	},
}
