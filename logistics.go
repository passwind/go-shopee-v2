package goshopee

type LogisticsService interface {
	GetChannelList(uint64, string) (*GetChannelListResponse, error)
	GetShippingParameter(uint64, string,string) (*GetShippingParameterResponse, error)
	ShipOrder(uint64, ShipOrderRequest, string) (*ShipOrderResponse, error)
}

type LogisticsServiceOp struct{
	client *Client
}

// https://open.shopee.com/documents?module=95&type=1&id=559&version=2
type GetChannelListResponse struct {
	BaseResponse

	Response GetChannelListResponseData `json:"response"`
}

type GetChannelListResponseData struct {
	LogisticsChannelList []LogisticsChannel `json:"logistics_channel_list"` // TODO: infact is logistics_channel_list ?
}

type LogisticsChannel struct {
	LogisticsChannelID uint64 `json:"logistics_channel_id"`
	LogisticsChannelName string `json:"logistics_channel_name"`
	CODEnabled bool `json:"cod_enabled"`
	Enabled bool `json:"enabled"`
	FeeType string `json:"fee_type"`
	SizeList []Size `json:"size_list"`
	WeightLimit WeightLimit `json:"weight_limit"`
	ItemMaxDimension ItemMaxDimension `json:"item_max_dimension"`
	Preferred bool `json:"preferred"`
	ForceEnabled bool `json:"force_enabled"` // TODO: infact in list
	MaskChannelID uint64 `json:"mask_channel_id"` // TODO: infact no?
	LogisticsDescription string `json:"logistics_description"` // TODO: infact no?
	VolumeLimit VolumeLimit `json:"volume_limit"`
}

type Size struct {
	SizeID string `json:"size_id"`
	Name string `json:"name"`
	DefaultPrice float64 `json:"default_price"`
}

type WeightLimit struct {
	ItemMaxWeight float64 `json:"item_max_weight"`
	ItemMinWeight float64 `json:"item_min_weight"`
}

type ItemMaxDimension struct {
	Height float64 `json:"height"`
	Width float64 `json:"width"`
	Length float64 `json:"length"`
	Unit string `json:"unit"`
}

type VolumeLimit struct {
	ItemMaxVolume float64 `json:"item_max_volume"`
	ItemMinVolume float64 `json:"item_min_volume"`
}

func (s *LogisticsServiceOp)GetChannelList(sid uint64, tok string) (*GetChannelListResponse, error){
	path := "/logistics/get_channel_list"

	resp := new(GetChannelListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, nil)
	return resp, err
}

type GetShippingParameterRequest struct {
	OrderSN string `url:"order_sn"`
}

type GetShippingParameterResponse struct {
	BaseResponse

	Response GetShippingParameterResponseData `json:"response"`
}

type GetShippingParameterResponseData struct {
	InfoNeeded GetShippingParameterResponseDataInfo `json:"info_needed"`
	Dropoff Dropoff `json:"dropoff"`
	Pickup Pickup `json:"pickup"`
}

type Pickup struct {
	AddressList []LogisticsAddress `json:"address_list"`
}

type LogisticsAddress struct {
	AddressID uint64 `json:"address_id"`
	Region string `json:"region"`
	State string `json:"state"`
	City string `json:"city"`
	Address string `json:"address"`
	Zipcode string `json:"zipcode"`
	District string `json:"district"`
	Town string `json:"town"`
	AddressFlag []string `json:"address_flag"`
	TimeSlotList []TimeSlot `json:"time_slot_list"`
}

type TimeSlot struct {
	Date int64 `json:"date"`
	TimeText string `json:"time_text"`
	PickupTimeID string `json:"pickup_time_id"`
}

type Dropoff struct {
	BranchList []Branch `json:"branch_list"`
}

type Branch struct {
	BranchID uint64 `json:"branch_id"`
	Region string `json:"region"`
	State string `json:"state"`
	City string `json:"city"`
	Address string `json:"address"`
	Zipcode string `json:"zipcode"`
	District string `json:"district"`
	Town string `json:"town"`
}

type GetShippingParameterResponseDataInfo struct {
	Dropoff []string `json:"dropoff"`
	Pickup []string `json:"pickup"`
	NonIntegrated []string `json:"non_integrated"`
}

func (s *LogisticsServiceOp)GetShippingParameter(sid uint64, ordersn, tok string) (*GetShippingParameterResponse, error){
	path := "/logistics/get_shipping_parameter"
	opt:=GetShippingParameterRequest{
		OrderSN: ordersn,
	}

	resp := new(GetShippingParameterResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type ShipOrderRequest struct {
	OrderSN string `json:"order_sn"`
	PackageNumber string `json:"package_number"`
	Pickup *ShipOrderRequestPickup `json:"pickup,omitempty"`
	Dropoff *ShipOrderRequestDropoff `json:"dropoff,omitempty"`
	NonIntegrated *ShipOrderRequestNonIntegrated `json:"non_integrated,omitempty"`
}

type ShipOrderRequestNonIntegrated struct {
	TrackingNumber string `json:"tracking_number"`
}

type ShipOrderRequestDropoff struct {
	BranchID uint64 `json:"branch_id"`
	SenderRealName string `json:"sender_real_name"`
	TrackingNumber string `json:"tracking_number"`
}

type ShipOrderRequestPickup struct {
	AddressID uint64 `json:"address_id"`
	PickupTimeID string `json:"pickup_time_id"`
	TrackingNumber string `json:"tracking_number"`
}

type ShipOrderResponse struct {
	BaseResponse
}

func (s *LogisticsServiceOp)ShipOrder(sid uint64, data ShipOrderRequest, tok string)(*ShipOrderResponse,error){
	path := "/logistics/ship_order"
	resp := new(ShipOrderResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}