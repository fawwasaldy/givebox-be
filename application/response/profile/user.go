package response_profile

type (
	User struct {
		ID          string `json:"id"`
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		City        string `json:"city"`
	}

	UserCreate struct {
		ID       string `json:"id"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}

	UserUpdate struct {
		ID          string `json:"id"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Biography   string `json:"biography,omitempty"`
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		City        string `json:"city,omitempty"`
	}

	UserChangePassword struct {
		ID string `json:"id"`
	}
)
