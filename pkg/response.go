package server

type Response struct {
	Status     string
	StatusCode uint
	Headers    map[string]string
	Data       string
}
