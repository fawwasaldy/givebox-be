package response_donation

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CategoryAnalytic struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
