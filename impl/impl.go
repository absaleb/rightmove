package impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	RightmoveClientTimeout = 3000
	RightmoveListingUrl    = "https://adfapi.rightmove.co.uk/v1/property/"
)

type RightmoveMethod int

const (
	SendProperty          RightmoveMethod = 0
	RemoveProperty        RightmoveMethod = 1
	GetBranchPropertyList RightmoveMethod = 2
)

func (z RightmoveMethod) String() string {
	names := [...]string{
		"sendpropertydetails",
		"removeproperty",
		"getbranchpropertylist"}

	if z < 0 || z > 3 {
		return ""
	}

	return names[z]
}

// type Property struct {
// 	agent_ref    string `json:"agent_ref"`
// 	rightmove_id string `json:"rightmove_id"`
// 	update_date  string `json:"update_date"`
// }
//
// type RightmoveSendCallRequest struct {
// 	Network struct {
// 		NetworkID int `json:"network_id"`
// 	} `json:"network"`
// 	Branch struct {
// 		BranchID int `json:"branch_id"`
// 		Channel  int `json:"channel"`
// 		// 	Overseas bool `json:"overseas"`
// 	} `json:"branch"`
// 	Property struct {
// 		AgentRef     string `json:"agent_ref"`
// 		Published    bool   `json:"published"`
// 		PropertyType int    `json:"property_type"`
// 		Status       int    `json:"status"`
// 		// NewHome         bool   `json:"new_home"`
// 		// StudentProperty bool   `json:"student_property"`
// 		// CreateDate      string `json:"create_date"`
// 		// UpdateDate      string `json:"update_date"`
// 		// DateAvailable   string `json:"date_available"`
// 		Address struct {
// 			HouseNameNumber string  `json:"house_name_number"`
// 			Town            string  `json:"town"`
// 			Postcode1       string  `json:"postcode_1"`
// 			Postcode2       string  `json:"postcode_2"`
// 			DisplayAddress  string  `json:"display_address"`
// 			Latitude        float64 `json:"latitude"`
// 			Longitude       float64 `json:"longitude"`
// 			// PovLatitude     float64 `json:"pov_latitude"`
// 			// PovLongitude    float64 `json:"pov_longitude"`
// 			// PovPitch        float64 `json:"pov_pitch"`
// 			// PovHeading      float64 `json:"pov_heading"`
// 			// PovZoom         int     `json:"pov_zoom"`
// 		} `json:"address"`
// 		PriceInformation struct {
// 			Price float64 `json:"price"`
// 			// PriceQualifier    int    `json:"price_qualifier"`
// 			// Deposit           int    `json:"deposit"`
// 			// AdministrationFee string `json:"administration_fee"`
// 		} `json:"price_information"`
// 		Details struct {
// 			Summary     string `json:"summary"`
// 			Description string `json:"description"`
// 			// Features                 []string `json:"features"`
// 			Bedrooms int `json:"bedrooms"`
// 			// Bathrooms                int      `json:"bathrooms"`
// 			// ReceptionRooms           int      `json:"reception_rooms"`
// 			// Parking                  []int    `json:"parking"`
// 			// OutsideSpace             []int    `json:"outside_space"`
// 			// YearBuilt                int      `json:"year_built"`
// 			// EntranceFloor            int      `json:"entrance_floor"`
// 			// Condition                int      `json:"condition"`
// 			// Accessibility            []int    `json:"accessibility"`
// 			// Heating                  []int    `json:"heating"`
// 			// FurnishedType            int      `json:"furnished_type"`
// 			// PetsAllowed              bool     `json:"pets_allowed"`
// 			// SmokersConsidered        bool     `json:"smokers_considered"`
// 			// HousingBenefitConsidered bool     `json:"housing_benefit_considered"`
// 			// SharersConsidered        bool     `json:"sharers_considered"`
// 			// AllBillsInc              bool     `json:"all_bills_inc"`
// 			// CouncilTaxInc            bool     `json:"council_tax_inc"`
// 			// Rooms                    []struct {
// 			// 	RoomName          string   `json:"room_name"`
// 			// 	RoomDescription   string   `json:"room_description"`
// 			// 	RoomLength        float64  `json:"room_length"`
// 			// 	RoomWidth         float64  `json:"room_width"`
// 			// 	RoomDimensionUnit int      `json:"room_dimension_unit"`
// 			// 	RoomDimensionText string   `json:"room_dimension_text"`
// 			// 	RoomPhotoUrls     []string `json:"room_photo_urls"`
// 			// } `json:"rooms"`
// 		} `json:"details"`
// 		Media []struct {
// 			MediaType int    `json:"media_type"`
// 			MediaURL  string `json:"media_url"`
// 			// Caption         string `json:"caption"`
// 			// SortOrder       int    `json:"sort_order"`
// 			// MediaUpdateDate string `json:"media_update_date"`
// 		} `json:"media"`
// 		// Principal struct {
// 		// 	PrincipalEmailAddress string `json:"principal_email_address"`
// 		// 	AutoEmailWhenLive     bool   `json:"auto_email_when_live"`
// 		// 	AutoEmailUpdates      bool   `json:"auto_email_updates"`
// 		// } `json:"principal"`
// 	} `json:"property"`
// }

type RemovePropertyRequest struct {
	Network struct {
		NetworkID int `json:"network_id"`
	} `json:"network"`
	Branch struct {
		BranchID int    `json:"branch_id"`
		Channel  string `json:"channel"`
	} `json:"branch"`
	Property struct {
		AgentRef string `json:"agent_ref"`
		// RemovalReason   int    `json:"removal_reason"`
		// TransactionDate string `json:"transaction_date"`
	} `json:"property"`
}

type GetListingsRequest struct {
	Network struct {
		NetworkID int `json:"network_id"`
	} `json:"network"`
	Branch struct {
		BranchID int `json:"branch_id"`
		// Channel  int `json:"channel"`
	} `json:"branch"`
}

// type RightmoveSendCallResponse struct {
// 	request_id         string `json:"request_id"`
// 	message            string `json:"message"`
// 	success            bool   `json:"success"`
// 	request_timestamp  string `json:"request_timestamp"`
// 	response_timestamp string `json:"response_timestamp"`
// }

type GetListingsResponse struct {
	request_id         string      `json:"request_id"`
	message            string      `json:"message"`
	success            bool        `json:"success"`
	request_timestamp  string      `json:"request_timestamp"`
	response_timestamp string      `json:"response_timestamp"`
	properies          []*Property `json:"property"`
}

func SendPropertyImpl(request *RightmoveSendCallRequest) (*RightmoveSendCallResponse, error) {
	method := SendProperty

	respBytes, err := getBytesByMethod(request, method)
	if err != nil {
		return nil, err
	}

	var result RightmoveSendCallResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


func getBytesByMethod(request interface{}, method RightmoveMethod) ([]byte, error) {
	jsn, err := getJSON(request)
	if err != nil {
		return nil, err
	}

	data := []byte(*jsn)
	addr := fmt.Sprintf("%s%s", RightmoveListingUrl, method)

	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	// headerValue := fmt.Sprintf("application/json; profile=%s%s.json", ZooplaListingHeaderUrl, method)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	respBytes, err := getBytes(req, RightmoveClientTimeout)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func getJSON(request interface{}) (*string, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	result := string(b)
	return &result, nil
}

func getBytes(request *http.Request, webClientTimeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(webClientTimeout) * time.Millisecond}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("provider return error code " + resp.Status)
	}

	return []byte(resp_body), nil
}
