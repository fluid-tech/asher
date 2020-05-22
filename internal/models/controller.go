package models

type Controller struct {
	Rest        bool     `json:"rest"`
	Mvc         bool     `json:"mvc"`
	HttpMethods []string `json:"httpMethods"`
	Type string 		 `json:"type"`
}
