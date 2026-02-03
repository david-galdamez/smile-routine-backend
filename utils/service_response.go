package utils

type ServiceResponse[T any] struct {
	Success      bool
	ErrorMessage *string
	Data         *T
}

func Ok[T any](data T) ServiceResponse[T] {
	return ServiceResponse[T]{
		Success: true,
		Data:    &data,
	}
}

func Error[T any](msg string) ServiceResponse[T] {
	return ServiceResponse[T]{
		Success:      false,
		ErrorMessage: &msg,
	}
}
