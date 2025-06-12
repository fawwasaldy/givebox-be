package response_donation

type (
	DonationItem struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Condition   int    `json:"condition"`
		PickCity    string `json:"pick_city"`
	}

	DonationItemOpen struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id"`
		Status      string `json:"status"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Condition   int    `json:"condition"`
		PickCity    string `json:"pick_city"`
		PickAddress string `json:"pick_address"`
	}

	DonationItemRequest struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id,omitempty"`
		RecipientID string `json:"recipient_id,omitempty"`
		Status      string `json:"status,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Condition   int    `json:"condition,omitempty"`
		PickCity    string `json:"pick_city,omitempty"`
		PickAddress string `json:"pick_address,omitempty"`
	}

	DonationItemAccept struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id,omitempty"`
		RecipientID string `json:"recipient_id,omitempty"`
		Status      string `json:"status,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Condition   int    `json:"condition,omitempty"`
		PickCity    string `json:"pick_city,omitempty"`
		PickAddress string `json:"pick_address,omitempty"`
	}

	DonationItemReject struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id,omitempty"`
		RecipientID string `json:"recipient_id,omitempty"`
		Status      string `json:"status,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Condition   int    `json:"condition,omitempty"`
		PickCity    string `json:"pick_city,omitempty"`
		PickAddress string `json:"pick_address,omitempty"`
	}

	DonationItemTaken struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id,omitempty"`
		RecipientID string `json:"recipient_id,omitempty"`
		Status      string `json:"status,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Condition   int    `json:"condition,omitempty"`
		PickCity    string `json:"pick_city,omitempty"`
		PickAddress string `json:"pick_address,omitempty"`
	}

	DonationItemUpdate struct {
		ID          string `json:"id"`
		DonorID     string `json:"donor_id,omitempty"`
		RecipientID string `json:"recipient_id,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
		Condition   int    `json:"condition,omitempty"`
		PickCity    string `json:"pick_city,omitempty"`
		PickAddress string `json:"pick_address,omitempty"`
	}
)
