package request_donation

type (
	DonationItemOpen struct {
		Name                string   `json:"name" form:"name" binding:"required,min=2,max=100"`
		Description         string   `json:"description" form:"description" binding:"required,max=500"`
		CategoryID          string   `json:"category_id" form:"category_id" binding:"required,uuid"`
		Condition           int      `json:"condition" form:"condition" binding:"required,oneof=1 2 3 4 5"`
		QuantityDescription string   `json:"quantity_description" form:"quantity_description" binding:"required,min=2,max=100"`
		Images              []string `json:"images" form:"images" binding:"omitempty,dive,uri,min=5,max=255"`
		PickCity            string   `json:"pick_city" form:"pick_city" binding:"required,min=2,max=50"`
		PickAddress         string   `json:"pick_address" form:"pick_address" binding:"required,min=5,max=255"`
		PickingStatus       string   `json:"picking_status" form:"picking_status" binding:"required"`
		DeliveryTime        string   `json:"delivery_time" form:"delivery_time" binding:"required,min=2,max=50"`
		IsUrgent            bool     `json:"is_urgent" form:"is_urgent"`
		AdditionalNote      string   `json:"additional_note" form:"additional_note" binding:"omitempty,max=500"`
	}

	DonationItemAccept struct {
		ID          string `json:"id" form:"id" binding:"required,uuid"`
		RecipientID string `json:"recipient_id" form:"recipient_id" binding:"required,uuid"`
	}
)
