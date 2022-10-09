package customerr

type errorCode string

const (
	defaultErrorMsg string = "internal server error"

	BadUserInput    errorCode = "BAD_USER_INPUT"
	badUserInputMsg string    = "invalid argument"

	InternalServerError    errorCode = "INTERNAL_SERVER_ERROR"
	internalServerErrorMsg string    = "internal server error"
)

func ErrorMessage(e errorCode) string {
	switch e {
	case BadUserInput:
		return badUserInputMsg
	case InternalServerError:
		return internalServerErrorMsg
	default:
		return defaultErrorMsg
	}
}

func ExtensionBadUserInput(invalidFields map[string]string) map[string]any {
	return map[string]any{
		"code":          BadUserInput,
		"invalidFields": invalidFields,
	}
}

func ExtensionInternalServerError() map[string]any {
	return map[string]any{
		"code": InternalServerError,
	}
}
