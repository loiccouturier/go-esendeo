package esendeo

import (
	"testing"
	"time"
)

func TestClient_CreateShipment(t *testing.T) {

	fOptions := FileOptions{
		IsSeparated: true,
		Encoding:    "base64",
	}

	senderAddress := Address{
		Company:     "Cindy H",
		ShopName:    "TEST sender shopSignage",
		Street:      "8-10 RUE DE LA HAIE COQ,",
		CountryCode: "FR",
		City:        "AUBERVILLIERS",
		PostalCode:  "93300",
		PhoneNumber: "0148340955",
		Email:       "cindyhbyazalee@gmail.com",
		LastName:    "La plus belle",
		FirstName:   "Jessica",
		Complement:  "LOT N°96",
	}

	receiverAddress := Address{
		FirstName:    "yangyang",
		LastName:     "xiao",
		Company:      "TEST COMPANY",
		ShopName:     "TEST receiver shopName",
		Email:        "xiaoyangyang@ftl-express.com",
		Type:         "receiver",
		PhoneNumber:  "",
		MobileNumber: "004531905171",
		Street:       "OSTERBROGAGE 114 2ND",
		Complement:   "",
		City:         "COPENHAGEN",
		PostalCode:   "2100",
		CountryCode:  "DK",
		Instructions: "TESTAAA",
		Comment:      "",
	}

	parcels := []Parcel{
		{
			Weight:              "3",
			InsuredValue:        300,
			Reference:           "test",
			PersonalizedBarcode: "98479879287498374928789",
		},
	}

	now := time.Now()
	now = now.Add(time.Duration(24) * time.Hour)

	shipment := Shipment{
		Parcels:         parcels,
		PickupAt:        now.Format("2006-01-02"),
		ReceiverAddress: receiverAddress,
		SenderAddress:   senderAddress,
		Transporter:     "UPS",
		Route:           "UPS_EXPRESS",
		LabelFormat:     "A6",
		FileOptions:     fOptions,
	}

	c := NewClient("https://sandbox.esendeo.com", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjUxMTcsImlhdCI6MTY4NDE3OTM5OX0.g2GukdhjVeZQ58e9Fbx3K5KCaP068z7Gac9rEkC6f7o", true)
	r, err := c.CreateShipment(shipment)

	if err != nil {
		t.Error(err)
	}

	if r != nil {
		if len(*r) > 0 {
			for _, response := range *r {
				if response.TrackingNumber == "" {
					t.Errorf("No tracking id")
				} else {
					t.Log("TrackingId", response.TrackingNumber)
				}
			}
		}
	}

}
