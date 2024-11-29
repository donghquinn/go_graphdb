package person

type Person struct {
	Name      string   `json:"name"`
	Age       string   `json:"age"`
	Languages []string `json:"languages"`
}
