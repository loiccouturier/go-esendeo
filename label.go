package esendeo

type LabelResponses struct {
	ShippingNumber string `json:"shippingNumber"`
	TrackingNumber string `json:"trackingNumber"`
	Label          string `json:"label"`
	LabelFormat    string `json:"labelFormat"`
}
