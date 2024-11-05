package models

type WebResponse struct {
	Code    int
	Status  bool
	Message string
	Data    interface{}
}
