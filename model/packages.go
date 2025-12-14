package model

type CommonList[T any] struct {
	Success bool   `json:"success"`
	Data    []T    `json:"data"`
	Message string `json:"message"`
}

type Package struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Version string `json:"version"`
}
