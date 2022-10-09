package customerr

type GQLErrorCode string

const (
	defaultErrorMessage = "internal server error"

	GQLErrorCodeBadUserInput        = "BAD_USER_INPUT"
	gqlErrorCodeBadUserInputMessage = "invalid argument"

	GQLErrorCodeInternalServerError        = "INTERNAL_SERVER_ERROR"
	gQLErrorCodeInternalServerErrorMessage = "internal server error"
)

func ErrorMessage(e GQLErrorCode) string {
	switch e {
	case GQLErrorCodeBadUserInput:
		return gqlErrorCodeBadUserInputMessage
	case GQLErrorCodeInternalServerError:
		return gQLErrorCodeInternalServerErrorMessage
	default:
		return defaultErrorMessage
	}
}

func ExtensionBadUserInput(invalidFields map[string]string) map[string]any {
	return map[string]any{
		"code":          GQLErrorCodeBadUserInput,
		"invalidFields": invalidFields,
	}
}

func ExtensionInternalServerError() map[string]any {
	return map[string]any{
		"code": GQLErrorCodeInternalServerError,
	}
}
