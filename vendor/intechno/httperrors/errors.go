package httperrors

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var NotImplemented = errors.New("not implemented") // используется только при прототипировании
const StatusRendererError = 422

type Response struct {
	Err            error             `json:"-"` // низкоуровневая ошибка исполнения
	HTTPStatusCode int               `json:"-"` // HTTP статус код
	ErrorMessage   *Details          `json:"error"`
	Validation     map[string]string `json:"validation,omitempty"` // ошибки валидации
}

type Details struct {
	StatusText  string `json:"status"`            // сообщение пользовательского уровня
	AppCode     int64  `json:"code,omitempty"`    // application-определенный код ошибки
	MessageText string `json:"message,omitempty"` // application-level сообщение, для дебага
}

func (e *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func PreconditionFailed(err error, validatorErrs map[string]string) render.Renderer {
	return &Response{
		Err:            errors.New("validation errors"),
		HTTPStatusCode: http.StatusPreconditionFailed,
		ErrorMessage: &Details{
			AppCode:     http.StatusPreconditionFailed,
			StatusText:  http.StatusText(http.StatusPreconditionFailed),
			MessageText: err.Error(),
		},
		Validation: validatorErrs,
	}
}

func UnprocessableEntity(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		ErrorMessage: &Details{
			AppCode:     http.StatusUnprocessableEntity,
			StatusText:  http.StatusText(http.StatusUnprocessableEntity),
			MessageText: err.Error(),
		},
	}
}

// Неправильный запрос.
// Возникает тогда, когда к запросу переданы неверные параметры.
func InvalidRequest(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage: &Details{
			AppCode:     http.StatusBadRequest,
			StatusText:  http.StatusText(http.StatusBadRequest),
			MessageText: err.Error(),
		},
	}
}

// Неправильный ответ от Renderer.
// Возникает тогда, когда рендереру не удается отрисовать ответ.
func ErrRender(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: StatusRendererError,
		ErrorMessage: &Details{
			StatusText:  "Error rendering response",
			MessageText: err.Error(),
		},
	}
}

// Не найден какой-то ресурс.
func ResourceNotFound(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		ErrorMessage: &Details{
			AppCode:     http.StatusNotFound,
			StatusText:  "Resource not found",
			MessageText: err.Error(),
		},
	}
}

func TooManyRequests(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusTooManyRequests,
		ErrorMessage: &Details{
			AppCode:     http.StatusTooManyRequests,
			StatusText:  "Too many requests",
			MessageText: err.Error(),
		},
	}
}

// Внутренняя ошибка сервера.
func Internal(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		ErrorMessage: &Details{
			AppCode:     http.StatusInternalServerError,
			StatusText:  "Internal Server Error",
			MessageText: err.Error(),
		},
	}
}

// Неправильный логин и пароль.
func InvalidCredentials(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessage: &Details{
			AppCode:     http.StatusUnauthorized,
			StatusText:  "Invalid Credentials",
			MessageText: err.Error(),
		},
	}
}

// Недостаточно прав.
func AccessDenied(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusForbidden,
		ErrorMessage: &Details{
			AppCode:     http.StatusForbidden,
			StatusText:  "Access Denied/Forbidden",
			MessageText: err.Error(),
		},
	}
}

// Нет такого токена.
func TokenNotFound(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		ErrorMessage: &Details{
			AppCode:     http.StatusNotFound,
			StatusText:  "Token not found",
			MessageText: err.Error(),
		},
	}
}

func Conflict(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		ErrorMessage: &Details{
			AppCode:     http.StatusConflict,
			StatusText:  http.StatusText(http.StatusConflict),
			MessageText: err.Error(),
		},
	}
}

func ConflictPhone(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		ErrorMessage: &Details{
			AppCode:     101,
			StatusText:  http.StatusText(http.StatusConflict),
			MessageText: err.Error(),
		},
	}
}

func ConflictEmail(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		ErrorMessage: &Details{
			AppCode:     100,
			StatusText:  http.StatusText(http.StatusConflict),
			MessageText: err.Error(),
		},
	}
}

func Unauthorized(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		ErrorMessage: &Details{
			AppCode:     http.StatusUnauthorized,
			StatusText:  http.StatusText(http.StatusUnauthorized),
			MessageText: err.Error(),
		},
	}
}

func BadRequest(err error) render.Renderer {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		ErrorMessage: &Details{
			AppCode:     http.StatusBadRequest,
			StatusText:  http.StatusText(http.StatusBadRequest),
			MessageText: err.Error(),
		},
	}
}

func Custom(err error, code int, opts ...ResponseOption) render.Renderer {
	renderer := &Response{
		Err:            err,
		HTTPStatusCode: code,
		ErrorMessage: &Details{
			AppCode:     int64(code),
			StatusText:  http.StatusText(code),
			MessageText: err.Error(),
		},
	}

	for _, opt := range opts {
		opt(renderer)
	}

	return renderer
}
