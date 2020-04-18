package bmw

type Vehicle struct {
	Series                 string   `json:"series"`
	BasicType              string   `json:"basicType"`
	BodyType               string   `json:"bodyType"`
	Brand                  string   `json:"brand"`
	ModelName              string   `json:"modelName"`
	VIN                    string   `json:"vin"`
	LicensePlate           string   `json:"licensePlate"`
	ModelYearNA            string   `json:"modelYearNA"`
	DcOnly                 bool     `json:"dcOnly"`
	HasNavi                bool     `json:"hasNavi"`
	HasSunRoof             bool     `json:"hasSunRoof"`
	DoorCount              int      `json:"doorCount"`
	MaxFuel                float64  `json:"maxFuel"`
	HasRex                 bool     `json:"hasRex"`
	Steering               string   `json:"steering"`
	DriveTrain             string   `json:"driveTrain"`
	SupportedChargingModes []string `json:"supportedChargingModes"`
}
