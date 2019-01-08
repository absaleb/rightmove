package impl

const (
	sendCallUrl       = "https://adfapi.rightmove.co.uk/v1/property/sendpropertydetails"
	removePropertyUrl = "https://adfapi.rightmove.co.uk/v1/property/removeproperty"
	getListingUrl     = "https://adfapi.rightmove.co.uk/v1/property/getbranchpropertylist"
)

type Property struct {
	agent_ref string `json:"agent_ref"`
	rightmove_id string `json:"rightmove_id"`
	update_date string `json:"update_date"`
}

type SendCallRequest struct {
	Network struct {
		NetworkID int `json:"network_id"`
	} `json:"network"`
	Branch struct {
		BranchID int `json:"branch_id"`
		Channel  int `json:"channel"`
		//	Overseas bool `json:"overseas"`
	} `json:"branch"`
	Property struct {
		AgentRef     string `json:"agent_ref"`
		Published    bool   `json:"published"`
		PropertyType int    `json:"property_type"`
		Status       int    `json:"status"`
		//NewHome         bool   `json:"new_home"`
		//StudentProperty bool   `json:"student_property"`
		//CreateDate      string `json:"create_date"`
		//UpdateDate      string `json:"update_date"`
		//DateAvailable   string `json:"date_available"`
		Address struct {
			HouseNameNumber string  `json:"house_name_number"`
			Town            string  `json:"town"`
			Postcode1       string  `json:"postcode_1"`
			Postcode2       string  `json:"postcode_2"`
			DisplayAddress  string  `json:"display_address"`
			Latitude        float64 `json:"latitude"`
			Longitude       float64 `json:"longitude"`
			//PovLatitude     float64 `json:"pov_latitude"`
			//PovLongitude    float64 `json:"pov_longitude"`
			//PovPitch        float64 `json:"pov_pitch"`
			//PovHeading      float64 `json:"pov_heading"`
			//PovZoom         int     `json:"pov_zoom"`
		} `json:"address"`
		PriceInformation struct {
			Price float64 `json:"price"`
			//PriceQualifier    int    `json:"price_qualifier"`
			//Deposit           int    `json:"deposit"`
			//AdministrationFee string `json:"administration_fee"`
		} `json:"price_information"`
		Details struct {
			Summary     string `json:"summary"`
			Description string `json:"description"`
			//Features                 []string `json:"features"`
			Bedrooms int `json:"bedrooms"`
			//Bathrooms                int      `json:"bathrooms"`
			//ReceptionRooms           int      `json:"reception_rooms"`
			//Parking                  []int    `json:"parking"`
			//OutsideSpace             []int    `json:"outside_space"`
			//YearBuilt                int      `json:"year_built"`
			//EntranceFloor            int      `json:"entrance_floor"`
			//Condition                int      `json:"condition"`
			//Accessibility            []int    `json:"accessibility"`
			//Heating                  []int    `json:"heating"`
			//FurnishedType            int      `json:"furnished_type"`
			//PetsAllowed              bool     `json:"pets_allowed"`
			//SmokersConsidered        bool     `json:"smokers_considered"`
			//HousingBenefitConsidered bool     `json:"housing_benefit_considered"`
			//SharersConsidered        bool     `json:"sharers_considered"`
			//AllBillsInc              bool     `json:"all_bills_inc"`
			//CouncilTaxInc            bool     `json:"council_tax_inc"`
			//Rooms                    []struct {
			//	RoomName          string   `json:"room_name"`
			//	RoomDescription   string   `json:"room_description"`
			//	RoomLength        float64  `json:"room_length"`
			//	RoomWidth         float64  `json:"room_width"`
			//	RoomDimensionUnit int      `json:"room_dimension_unit"`
			//	RoomDimensionText string   `json:"room_dimension_text"`
			//	RoomPhotoUrls     []string `json:"room_photo_urls"`
			//} `json:"rooms"`
		} `json:"details"`
		Media []struct {
			MediaType int    `json:"media_type"`
			MediaURL  string `json:"media_url"`
			//Caption         string `json:"caption"`
			//SortOrder       int    `json:"sort_order"`
			//MediaUpdateDate string `json:"media_update_date"`
		} `json:"media"`
		//Principal struct {
		//	PrincipalEmailAddress string `json:"principal_email_address"`
		//	AutoEmailWhenLive     bool   `json:"auto_email_when_live"`
		//	AutoEmailUpdates      bool   `json:"auto_email_updates"`
		//} `json:"principal"`
	} `json:"property"`
}

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
		//RemovalReason   int    `json:"removal_reason"`
		//TransactionDate string `json:"transaction_date"`
	} `json:"property"`
}

type GetListingsRequest struct {
	Network struct {
		NetworkID int `json:"network_id"`
	} `json:"network"`
	Branch struct {
		BranchID int `json:"branch_id"`
		//Channel  int `json:"channel"`
	} `json:"branch"`
}

type SendCallResponse struct {
	request_id         string `json:"request_id"`
	message            string `json:"message"`
	success            bool   `json:"success"`
	request_timestamp  string `json:"request_timestamp"`
	response_timestamp string `json:"response_timestamp"`
}

type GetListingsResponse struct {
	request_id         string `json:"request_id"`
	message            string `json:"message"`
	success            bool   `json:"success"`
	request_timestamp  string `json:"request_timestamp"`
	response_timestamp string `json:"response_timestamp"`
	 properies []*Property `json:"property"`
}


