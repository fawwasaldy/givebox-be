package response_donation

type Image struct {
	ID            string `json:"id"`
	DonatedItemID string `json:"donated_item_id"`
	ImageURL      string `json:"image_url"`
}
