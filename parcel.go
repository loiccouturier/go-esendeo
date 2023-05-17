package esendeo

type Parcel struct {
	Weight              string  `json:"weight"`
	InsuredValue        int     `json:"insuredValue"`
	Reference           string  `json:"reference"`
	PersonalizedBarcode string  `json:"personalizedBarcode"`
	DeclaredWeight      string  `json:"declaredWeight"`
	VolumeWeight        string  `json:"volumeWeight"`
	Length              float32 `json:"length"`
	Width               float32 `json:"width"`
	Height              float32 `json:"height"`
	Items               []Item  `json:"items"`
}

type Item struct {
	HsCode        string  `json:"hsCode"`
	OriginCountry string  `json:"originCountry"`
	Weight        float32 `json:"weight"`
	Quantity      float32 `json:"quantity"`
	Value         float32 `json:"value"`
	Description   string  `json:"description"`
}
