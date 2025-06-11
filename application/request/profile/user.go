package profile

type (
	UserRegister struct {
		FullName    string `json:"full_name" form:"full_name" binding:"required,min=2,max=100"`
		Username    string `json:"username" form:"username" binding:"required,min=2,max=50"`
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
		Password    string `json:"password" form:"password" binding:"required,min=8"`
	}

	UserUpdate struct {
		Name        string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
	}

	UserLogin struct {
		Username string `json:"username" form:"username" binding:"required,min=2,max=50"`
		Password string `json:"password" form:"password" binding:"required,min=8"`
	}
)
