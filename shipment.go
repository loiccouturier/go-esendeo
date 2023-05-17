package esendeo

import "time"

type Shipment struct {
	Parcels         []Parcel    `json:"parcels"`
	PickupAt        string      `json:"pickupAt"`
	ReceiverAddress Address     `json:"receiverAddress"`
	SenderAddress   Address     `json:"senderAddress"`
	Transporter     string      `json:"transporter"`
	Route           string      `json:"route"`
	LabelFormat     string      `json:"labelFormat"`
	FileOptions     FileOptions `json:"FileOptions"`
}

type ShipmentResponse struct {
	Weight               float32        `json:"weight"`
	InsuredValue         float32        `json:"insuredValue"`
	Reference            string         `json:"reference"`
	PersonalizedBarcode  string         `json:"personalizedBarcode"`
	MultiParcelRank      string         `json:"multiParcelRank"`
	Items                []Item         `json:"items"`
	DeclaredWeight       float32        `json:"declaredWeight"`
	ShippingFee          float32        `json:"shippingFee"`
	InsuranceFee         float32        `json:"insuranceFee"`
	RemoteAreaFee        float32        `json:"remoteAreaFee"`
	ShippingToPrivateFee float32        `json:"shippingToPrivateFee"`
	PickupFee            float32        `json:"pickupFee"`
	FuelTaxFee           float32        `json:"fuelTaxFee"`
	FuelTaxRate          float32        `json:"fuelTaxRate"`
	Vat                  float32        `json:"vat"`
	Et                   float32        `json:"et"`
	It                   float32        `json:"it"`
	MinWeight            float32        `json:"minWeight"`
	MaxWeight            float32        `json:"maxWeight"`
	MinQuantity          float32        `json:"minQuantity"`
	IsMulti              bool           `json:"isMulti"`
	PricingRangeId       int            `json:"pricingRangeId"`
	TrackingNumber       string         `json:"trackingNumber"`
	ShippingNumber       string         `json:"shippingNumber"`
	LabelFormat          string         `json:"labelFormat"`
	Transporter          string         `json:"transporter"`
	LabelUrl             string         `json:"labelUrl"`
	BarCode              string         `json:"barCode"`
	LabelResponses       LabelResponses `json:"labelResponses"`
	PickupRequestNumber  string         `json:"pickupRequestNumber"`
	PayByReceiver        bool           `json:"payByReceiver"`
	FileOptions          FileOptions    `json:"FileOptions"`
	PickupAt             string         `json:"pickupAt"`
	Route                string         `json:"route"`
	SenderClientId       int            `json:"senderClientId"`
	ClientId             int            `json:"clientId"`
	UserId               int            `json:"userId"`
	IsApiClient          bool           `json:"isApiClient"`
	RateShipmentType     string         `json:"rateShipmentType"`
	RouteReturnOnly      bool           `json:"routeReturnOnly"`
	ValidatedAt          *time.Time     `json:"validatedAt"`
	IsReturnPricingId    bool           `json:"isReturnPricingId"`
	Code                 string         `json:"code"`
	Application          string         `json:"application"`
	SenderAddressId      int            `json:"senderAddressId"`
	ReceiverAddressId    int            `json:"receiverAddressId"`
	TrackingNumber2      string         `json:"trackingNumber2"`
	Source               string         `json:"source"`
	AccountShippingFee   int            `json:"accountShippingFee"`
	Length               float32        `json:"length"`
	Width                float32        `json:"width"`
	Height               float32        `json:"height"`
	TransferredAt        *time.Time     `json:"transferredAt"`
	ArrivedAt            *time.Time     `json:"arrivedAt"`
	Status               string         `json:"status"`
	OriginTrackingNumber string         `json:"originTrackingNumber"`
	IsArrived            bool           `json:"isArrived"`
	IsReturned           bool           `json:"isReturned"`
	ReturnedAt           *time.Time     `json:"returnedAt"`
	PickedAt             *time.Time     `json:"pickedAt"`
	//PickAtHour           interface{} `json:"pickAtHour"`
	SendingReason string `json:"sendingReason"`
	//	Pallets              interface{} `json:"pallets"`
	DeletedAt *time.Time `json:"deletedAt"`
	PricingId int        `json:"pricingId"`
	//	ClientRouteId        interface{} `json:"clientRouteId"`
	RelayPointId float32 `json:"relayPointId"`
	//	CustomClearance      interface{} `json:"customClearance"`
	Id           int       `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	EtWithFuel   float32   `json:"etWithFuel"`
	CustomsLabel string    `json:"customsLabel"`
	// options
}
