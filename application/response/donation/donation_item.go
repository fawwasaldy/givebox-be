package response_donation

type (
	DonationItem struct {
		ID          string `json:"id"`
		DonorName   string `json:"donor_name"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Category    string `json:"category"`
		Condition   int    `json:"condition"`
		PickCity    string `json:"pick_city"`
		IsUrgent    bool   `json:"is_urgent"`
		CreatedAt   string `json:"created_at"`
	}

	DetailedDonationItem struct {
		ID                  string `json:"id"`
		DonorName           string `json:"donor_name"`
		Name                string `json:"name"`
		Description         string `json:"description"`
		Category            string `json:"category"`
		Condition           int    `json:"condition"`
		QuantityDescription string `json:"quantity_description"`
		PickCity            string `json:"pick_city"`
		PickAddress         string `json:"pick_address"`
		PickingStatus       string `json:"picking_status"`
		DeliveryTime        string `json:"delivery_time"`
		IsUrgent            bool   `json:"is_urgent"`
		AdditionalNote      string `json:"additional_note,omitempty"`
	}

	DonationItemOpen struct {
		ID                  string `json:"id"`
		DonorName           string `json:"donor_name"`
		Status              string `json:"status"`
		Name                string `json:"name"`
		Description         string `json:"description"`
		Category            string `json:"category"`
		Condition           int    `json:"condition"`
		QuantityDescription string `json:"quantity_description"`
		PickCity            string `json:"pick_city"`
		PickAddress         string `json:"pick_address"`
		PickingStatus       string `json:"picking_status"`
		DeliveryTime        string `json:"delivery_time"`
		IsUrgent            bool   `json:"is_urgent"`
		AdditionalNote      string `json:"additional_note,omitempty"`
	}

	DonationItemAccept struct {
		ID                  string `json:"id"`
		DonorName           string `json:"donor_name,omitempty"`
		RecipientName       string `json:"recipient_name,omitempty"`
		Status              string `json:"status,omitempty"`
		Name                string `json:"name,omitempty"`
		Description         string `json:"description,omitempty"`
		Category            string `json:"category,omitempty"`
		Condition           int    `json:"condition,omitempty"`
		QuantityDescription string `json:"quantity_description,omitempty"`
		PickCity            string `json:"pick_city,omitempty"`
		PickAddress         string `json:"pick_address,omitempty"`
		PickingStatus       string `json:"picking_status,omitempty"`
		DeliveryTime        string `json:"delivery_time,omitempty"`
		IsUrgent            bool   `json:"is_urgent,omitempty"`
		AdditionalNote      string `json:"additional_note,omitempty"`
	}
)
