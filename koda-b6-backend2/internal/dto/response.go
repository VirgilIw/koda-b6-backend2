package dto

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  T      `json:"result,omitempty"`
}
