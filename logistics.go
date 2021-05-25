package goshopee

type LogisticsService interface {
	GetChannelList(int64, string) (*GetChannelListResponse, error)
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
	LogisticsChannelList []LogisticsChannel `json:"logistics"` // TODO: or logistics_channel_list ?
	LogisticsDescription string `json:"logistics_description"`
	ForceEnabled bool `json:"force_enabled"`
	MaskChannelID int64 `json:"mask_channel_id"`
}

type LogisticsChannel struct {
	LogisticsChannelID int64 `json:"logistics_channel_id"`
	Preferred bool `json:"preferred"`
	LogisticsChannelName string `json:"logistics_channel_name"`
	CODEnabled bool `json:"cod_enabled"`
	Enabled bool `json:"enabled"`
	FeeType string `json:"fee_type"`
	SizeList []Size `json:"size_list"`
	WeightLimit WeightLimit `json:"weight_limit"`
	ItemMaxDimension ItemMaxDimension `json:"item_max_dimension"`
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

func (s *LogisticsServiceOp)GetChannelList(sid int64, tok string) (*GetChannelListResponse, error){
	path := "/logistics/get_channel_list"

	resp := new(GetChannelListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, nil)
	return resp, err
}