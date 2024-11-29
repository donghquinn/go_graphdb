package person

type Person struct {
	Name          string   `json:"name"`
	Age           string   `json:"age"`
	Languages     []string `json:"languages"`
	Nationality   string   `json:"nationality"`
	ContactNumber string   `json:"contactNumber,omitempty"`
}

// ============== REQUEST ==============
type CreatePersonRequest struct {
	Name          string   `json:"name"`
	Age           string   `json:"age"`
	Languages     []string `json:"languages"`
	Nationality   string   `json:"nationality"`
	ContactNumber string   `json:"contactNumber,omitempty"`
	TechName      string   `json:"techName"`
}

type GetPersonRequestByNameRequest struct {
	Name string `json:"name"`
}

// ============== RESPONSE ==============
type GetPersonResponse struct {
	Status  int      `json:"status"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Person  []Person `json:"person"`
}
