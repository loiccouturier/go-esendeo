package esendeo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var ErrBadRequest = errors.New("ErrBadRequest")
var ErrForbidden = errors.New("ErrForbidden")
var ErrNotFound = errors.New("ErrNotFound")
var ErrUndefined = errors.New("ErrUndefined")

type Client interface {
	CreateShipment(shipment Shipment) (*[]ShipmentResponse, *ApiError)
}

type client struct {
	host   string
	apiKey string
	debug  bool
}

func NewClient(host, apiKey string, debug bool) Client {
	return &client{
		host:   host,
		apiKey: apiKey,
		debug:  debug,
	}
}

func (c *client) CreateShipment(shipment Shipment) (*[]ShipmentResponse, *ApiError) {
	if c.debug {
		log.Println("esendeo.client", "CreateShipment")
	}

	var result []ShipmentResponse

	apiError := c.call(http.MethodPost, fmt.Sprintf("%s/api/parcel", c.host), shipment, &result, true)
	if apiError != nil {
		return nil, apiError
	}

	return &result, nil
}

func (c *client) call(method, url string, body, result interface{}, isLabel bool) *ApiError {
	// Json encode body
	if body == nil {
		body = ""
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return &ApiError{Err: err}
	}
	if c.debug {
		log.Println("esendeo.client", "payload: ", string(jsonBody))
	}

	httpClient := http.DefaultClient
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		return &ApiError{Err: err}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return &ApiError{Err: err}
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return &ApiError{Err: err}
	}
	if c.debug {
		log.Println("esendeo.client", "Response: ", string(responseBody))
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			panic(fmt.Sprintf("can not close response body: %v\n", err))
		}
	}()

	err = nil
	if response.StatusCode == http.StatusBadRequest {
		if c.debug {
			log.Println("esendeo.client", "Http status 400 received for url ", url, "response", string(responseBody))
		}
		err = ErrBadRequest
	} else if response.StatusCode == http.StatusForbidden {
		if c.debug {
			log.Println("esendeo.client", "Http status 403 received for url ", url)
		}
		err = ErrForbidden
	} else if response.StatusCode == http.StatusNotFound {
		if c.debug {
			log.Println("esendeo.client", "Http status 404 received for url ", url)
		}
		err = ErrNotFound
	} else if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusNoContent {
		if c.debug {
			log.Println("esendeo.client", fmt.Sprintf("Http status %d received for url %s", response.StatusCode, url))
		}
		err = ErrUndefined
	}

	if err == nil {
		if c.debug {
			log.Println("esendeo.client", "Success for url ", url)
		}
		if result != nil {
			if c.debug {
				log.Println("esendeo.client", "Unmarshal the result")
			}
			err = json.Unmarshal(responseBody, result)
			if err != nil {
				return &ApiError{Err: err}
			}
		}
	} else {
		if c.debug {
			log.Println("esendeo.client", "error received for url ", url)
		}
		var apiError ApiError
		err2 := json.Unmarshal(responseBody, &apiError)
		if err2 != nil {
			return &ApiError{Err: err2}
		}
		apiError.Err = err
		return &apiError
	}

	return nil
}
