package request_profile

type (
	UserRegister struct {
		FirstName   string `json:"first_name" form:"first_name" binding:"required,min=2,max=50"`
		LastName    string `json:"last_name" form:"last_name" binding:"required,min=2,max=50"`
		Email       string `json:"email" form:"email" binding:"required,email"`
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
		City        string `json:"city" form:"city" binding:"omitempty,min=2,max=100"`
		Password    string `json:"password" form:"password" binding:"required,min=8"`
	}

	UserUpdate struct {
		FirstName   string `json:"first_name" form:"first_name" binding:"omitempty,min=2,max=50"`
		LastName    string `json:"last_name" form:"last_name" binding:"omitempty,min=2,max=50"`
		Biography   string `json:"biography" form:"biography" binding:"omitempty,max=500"`
		Email       string `json:"email" form:"email" binding:"omitempty,email"`
		PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty,min=8,max=20"`
		City        string `json:"city" form:"city" binding:"omitempty,min=2,max=100"`
	}

	UserLogin struct {
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required,min=8"`
	}

	UserChangePassword struct {
		OldPassword string `json:"old_password" form:"old_password" binding:"required,min=8"`
		NewPassword string `json:"new_password" form:"new_password" binding:"required,min=8"`
	}
)
