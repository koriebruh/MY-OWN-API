package entity

type Pokemon struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Species string `json:"species"`
}
