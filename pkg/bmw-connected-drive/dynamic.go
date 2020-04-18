package bmw

type DynamicResponse struct {
	AttributesMap DynamicAttributes `json:"attributesMap"`
}

type DynamicAttributes struct {
	UnitOfLength                               string  `json:"unitOfLength"`
	ChargingLogicCurrentlyActive               string  `json:"chargingLogicCurrentlyActive"`
	VehicleTracking                            int     `json:"vehicle_tracking,string"`
	ChargeNowAllowed                           string  `json:"chargeNowAllowed"`
	UpdateTimeConverted                        string  `json:"updateTime_converted"`
	DoorDriverRear                             string  `json:"door_driver_rear"`
	HeadUnitPuSoftware                         string  `json:"head_unit_pu_software"`
	BeMaxRangeElectricKm                       float64 `json:"beMaxRangeElectricKm,string"`
	DoorPassengerRear                          string  `json:"door_passenger_rear"`
	BeRemainingRangeFuelKm                     float64 `json:"beRemainingRangeFuelKm,string"`
	SegmentLastTripTimeSegmentEndFormattedDate string  `json:"Segment_LastTrip_time_segment_end_formatted_date"`
	DoorDriverFront                            string  `json:"door_driver_front"`
	HoodState                                  string  `json:"hood_state"`
	ChargingStatus                             string  `json:"charging_status"`
	KombiCurrentRemainingRangeFuel             float64 `json:"kombi_current_remaining_range_fuel,string"`
	BeMaxRangeElectric                         float64 `json:"beMaxRangeElectric,string"`
	WindowDriverRear                           string  `json:"window_driver_rear"`
	BeRemainingRangeElectricKm                 float64 `json:"beRemainingRangeElectricKm,string"`
	Mileage                                    int     `json:"mileage,string"`
	SegmentLastTripTimeSegmentEndFormattedTime string  `json:"Segment_LastTrip_time_segment_end_formatted_time"`
	BeMaxRangeElectricMile                     float64 `json:"beMaxRangeElectricMile,string"`
	SegmentLastTripTimeSegmentEndFormatted     string  `json:"Segment_LastTrip_time_segment_end_formatted"`
	LastChargingEndResult                      string  `json:"lastChargingEndResult"`
	UnitOfEnergy                               string  `json:"unitOfEnergy"`
	BeRemainingRangeElectric                   float64 `json:"beRemainingRangeElectric,string"`
	SocHvPercent                               float64 `json:"soc_hv_percent,string"`
	SingleImmediateCharging                    string  `json:"single_immediate_charging"`
	UpdateTimeConvertedTime                    string  `json:"updateTime_converted_time"`
	ChargingHVStatus                           string  `json:"chargingHVStatus"`
	ConnectorStatus                            string  `json:"connectorStatus"`
	ChargingLevelHv                            float64 `json:"chargingLevelHv,string"`
	ChargingSystemStatus                       string  `json:"chargingSystemStatus"`
	FuelPercent                                int     `json:"fuelPercent,string"`
	UnitOfCombustionConsumption                string  `json:"unitOfCombustionConsumption"`
	GpsLat                                     float64 `json:"gps_lat,string"`
	WindowDriverFront                          string  `json:"window_driver_front"`
	SegmentLastTripRatioElectricDrivenDistance int     `json:"Segment_LastTrip_ratio_electric_driven_distance,string"`
	GpsLng                                     float64 `json:"gps_lng,string"`
	ConditionBasedServices                     string  `json:"condition_based_services"`
	WindowPassengerFront                       string  `json:"window_passenger_front"`
	WindowPassengerRear                        string  `json:"window_passenger_rear"`
	LastChargingEndReason                      string  `json:"lastChargingEndReason"`
	UpdateTimeConvertedDate                    string  `json:"updateTime_converted_date"`
	BeRemainingRangeFuelMile                   float64 `json:"beRemainingRangeFuelMile,string"`
	BeRemainingRangeFuel                       float64 `json:"beRemainingRangeFuel,string"`
	DoorPassengerFront                         string  `json:"door_passenger_front"`
	UpdateTimeConvertedTimestamp               int64   `json:"updateTime_converted_timestamp,string"`
	RemainingFuel                              float64 `json:"remaining_fuel,string"`
	ChargingInductivePositioning               string  `json:"charging_inductive_positioning"`
	Heading                                    int     `json:"heading,string"`
	LscTrigger                                 string  `json:"lsc_trigger"`
	LightsParking                              string  `json:"lights_parking"`
	DoorLockState                              string  `json:"door_lock_state"`
	UpdateTime                                 string  `json:"updateTime"`
	PrognosisWhileChargingStatus               string  `json:"prognosisWhileChargingStatus"`
	HeadUnit                                   string  `json:"head_unit"`
	TrunkState                                 string  `json:"trunk_state"`
	BatterySizeMax                             int     `json:"battery_size_max,string"`
	ChargingConnectionType                     string  `json:"charging_connection_type"`
	BeRemainingRangeElectricMile               float64 `json:"beRemainingRangeElectricMile,string"`
	UnitOfElectricConsumption                  string  `json:"unitOfElectricConsumption"`
	SegmentLastTripTimeSegmentEnd              string  `json:"Segment_LastTrip_time_segment_end"`
	LastUpdateReason                           string  `json:"lastUpdateReason"`
	VehicleMessages                            struct {
		CcmMessages []DynamicVehicleMessage `json:"ccmMessages"`
		CbsMessages []DynamicVehicleMessage `json:"cbsMessages"`
	} `json:"vehicleMessages"`
}

type DynamicVehicleMessage struct {
	Description string `json:"description"`
	Text        string `json:"text"`
	ID          int    `json:"id"`
	Status      string `json:"status"`
	MessageType string `json:"messageType"`
	Date        string `json:"date"`
}
