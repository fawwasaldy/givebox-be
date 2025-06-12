package response_profile

type (
	User struct {
		ID          string `json:"id"`
		FullName    string `json:"name"`
		Username    string `json:"username"`
		PhoneNumber string `json:"phone_number"`
	}

	UserCreate struct {
		ID       string `json:"id"`
		FullName string `json:"full_name"`
		Username string `json:"username"`
	}

	UserUpdate struct {
		ID          string `json:"id"`
		FullName    string `json:"name,omitempty"`
		Username    string `json:"username,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
	}
)
