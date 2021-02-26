package httperrors

type ResponseOption func(r *Response)

func WithValidationErrors(validationErrors map[string]string) ResponseOption {
	return func(r *Response) {
		r.Validation = validationErrors
	}
}