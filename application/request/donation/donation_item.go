package request_donation

type (
	DonationItemOpen struct {
		Name        string   `json:"name" form:"name" binding:"required,min=2,max=100"`
		Description string   `json:"description" form:"description" binding:"required,max=500"`
		Condition   int      `json:"condition" form:"condition" binding:"required,oneof=0 1 2 3 4 5"`
		PickCity    string   `json:"pick_city" form:"pick_city" binding:"required,min=2,max=50"`
		PickAddress string   `json:"pick_address" form:"pick_address" binding:"required,min=5,max=255"`
		Images      []string `json:"images" form:"images" binding:"omitempty,dive,uri,min=5,max=255"`
	}

	DonationItemRequest struct {
		ID string `json:"id" form:"id" binding:"required,uuid"`
	}

	DonationItemAccept struct {
		ID string `json:"id" form:"id" binding:"required,uuid"`
	}

	DonationItemReject struct {
		ID string `json:"id" form:"id" binding:"required,uuid"`
	}

	DonationItemTaken struct {
		ID string `json:"id" form:"id" binding:"required,uuid"`
	}

	DonationItemUpdate struct {
		ID          string `json:"id" form:"id" binding:"required,uuid"`
		Name        string `json:"name,omitempty" form:"name,omitempty" binding:"omitempty,min=2,max=100"`
		Description string `json:"description,omitempty" form:"description,omitempty" binding:"omitempty,max=500"`
		Condition   int    `json:"condition,omitempty" form:"condition,omitempty" binding:"omitempty,oneof=0 1 2 3 4 5"`
		PickAddress string `json:"pick_address,omitempty" form:"pick_address,omitempty" binding:"omitempty,min=5,max=255"`
	}
)
