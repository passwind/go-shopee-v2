package goshopee

type ProductService interface {
	GetCategory(int64, string, string) (*GetCategoryResponse,error)
	GetBrandList(int64, int64, int, int, int, string) (*GetBrandListResponse, error)
	GetDTSLimit(int64, int64, string) (*GetDTSLimitResponse, error)
	GetAttributes(int64, int64, string, string) (*GetAttributesResponse, error)
	SupportSizeChart(int64, int64, string) (*SupportSizeChartResponse, error)
	UpdateSizeChart(int64, int64, string, string)(*UpdateSizeChartResponse, error)
	AddItem(int64, AddItemRequest, string)(*AddItemResponse,error)
	InitTierVariation(int64, InitTierVariationRequest, string) (*InitTierVariationResponse,error)
	AddModel(int64, AddModelRequest, string)(*AddModelResponse, error)
	GetModelList(int64, int64, string) (*GetModelListResponse, error)
}

type GetCategoryResponse struct {
	BaseResponse

	Response GetCategoryResponseData `json:"response"`
}

type GetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"`
}

type Category struct {
	CategoryID int64 `json:"category_id"`
	ParentCategoryID int64 `json:"parent_category_id"`
	OriginalCategoryName string `json:"original_category_name"`
	DisplayCategoryName string `json:"display_category_name"`
	HasChildren bool `json:"has_children"`
}

type ProductServiceOp struct {
	client *Client
}

type GetCategoryRequest struct {
	Language string `url:"language"`
}

func (s *ProductServiceOp)	GetCategory(sid int64, lang, tok string) (*GetCategoryResponse,error){
	path := "/product/get_category"

	opt:=GetCategoryRequest{
		Language:lang,
	}

	resp := new(GetCategoryResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type GetBrandListRequest struct {
	Offset int `url:"offset"`
	PageSize int `url:"page_size"`
	CategoryID int64 `url:"category_id"`
	Status int `url:"status"`
}

type GetBrandListResponse struct {
	BaseResponse

	Response GetBrandListResponseData `json:"response"`
}

type GetBrandListResponseData struct {
	BrandList []Brand `json:"brand_list"`
	HasNextPage bool `json:"has_next_page"`
	NextOffset int `json:"next_offset"`
	IsMandatory bool `json:"is_mandatory"`
	InputType string `json:"input_type"`
}

type Brand struct {
	BrandID int64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
	DisplayBrandName string `json:"display_brand_name"`
}

func (s *ProductServiceOp)	GetBrandList(sid, cid int64, status, offset, pageSize int, tok string) (*GetBrandListResponse, error){
	path := "/product/get_brand_list"

	opt:=GetBrandListRequest{
		CategoryID: cid,
		Offset: offset,
		PageSize: pageSize,
		Status: status,
	}

	resp := new(GetBrandListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type GetDTSLimitRequest struct {
	CategoryID int64 `url:"category_id"`
}

type GetDTSLimitResponse struct {
	BaseResponse

	Response GetDTSLimitResponseData `json:"response"`
}

type GetDTSLimitResponseData struct {
	DaysToShipLimit DaysToShipLimit `json:"days_to_ship_limit"`
	NonPreOrderDaysToShip int `json:"non_pre_order_days_to_ship"`
}

type DaysToShipLimit struct {
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

func (s *ProductServiceOp)	GetDTSLimit(sid, cid int64, tok string) (*GetDTSLimitResponse, error){
	path := "/product/get_dts_limit"

	opt:=GetDTSLimitRequest{
		CategoryID: cid,
	}

	resp := new(GetDTSLimitResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type GetAttibutesRequest struct {
	CategoryID int64 `url:"category_id"`
	Language string `url:"language"`
}

type GetAttributesResponse struct {
	BaseResponse

	Response GetAttributesResponseData `json:"response"`
}

type GetAttributesResponseData struct {
	AttributeList []Attribute `json:"attribute_list"`
}

type Attribute struct {
	AttributeID int64 `json:"attribute_id"`
	OriginalAttributeName string `json:"original_attribute_name"`
	DisplayAttributeName string `json:"display_attribute_name"`
	IsMandatory bool `json:"is_mandatory"`
	InputValidationType string `json:"input_validation_type"`
	FormatType string `json:"format_type"`
	DateFormatType string `json:"date_format_type"`
	InputType string `json:"input_type"`
	AttributeUnit []string `json:"attribute_unit"`
	AttributeValueList []AttributeValue `json:"attribute_value_list"`
}

type AttributeValue struct {
	ValueID int64 `json:"value_id"`
	OriginalValueName string `json:"original_value_name"`
	DisplayValueName string `json:"display_value_name"`
	ValueUnit string `json:"value_unit"`
	ParentAttributeList []ParentAttribute `json:"parent_attribute_list"`
	ParentBrandList []ParentBrand `json:"parent_brand_list"`
}

type ParentAttribute struct {
	ParentAttributeID int64 `json:"parent_attribute_id"`
	ParentValueID int64 `json:"parent_value_id"`
} 

type ParentBrand struct {
	ParentBrandID int64 `json:"parent_brand_id"`
}

func (s *ProductServiceOp)	GetAttributes(sid, cid int64, lang, tok string) (*GetAttributesResponse, error){
	path := "/product/get_attributes"

	opt:=GetAttibutesRequest{
		CategoryID: cid,
		Language: lang,
	}

	resp := new(GetAttributesResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type SupportSizeChartRequest struct {
	CategoryID int64 `url:"category_id"`
}

type SupportSizeChartResponse struct {
	BaseResponse

	Response SupportSizeChartResponseData `json:"response"`
}

type SupportSizeChartResponseData struct {
	SupportSizeChart bool `json:"support_size_chart"`
}

func (s *ProductServiceOp)SupportSizeChart(sid, cid int64, tok string) (*SupportSizeChartResponse, error){
	path := "/product/support_size_chart"

	opt:=SupportSizeChartRequest{
		CategoryID: cid,
	}

	resp := new(SupportSizeChartResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type UpdateSizeChartResponse struct {
	BaseResponse
}

func (s *ProductServiceOp)UpdateSizeChart(sid, itemID int64, sizeChart,tok string)(*UpdateSizeChartResponse, error) {
	path := "/product/update_size_chart"
	wrappedData := map[string]interface{}{
		"item_id": itemID,
		"size_chart":  sizeChart,
	}
	resp := new(UpdateSizeChartResponse)
	err := s.client.WithShop(sid,tok).Post(path, wrappedData, resp)
	return resp, err
}

type AddItemRequest struct {
	ItemBase
}

type ItemBase struct {
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	OriginalPrice float64 `json:"original_price"`
	Weight float64 `json:"weight"`
	ItemStatus string `json:"item_status"`
	Dimension Dimension `json:"dimension"`
	NormalStock int `json:"normal_stock"`
	LogisticInfo []LogisticInfo `json:"logistic_info"`
	AttributeList []ItemAttribute `json:"attribute_list"`
	CategoryID int64 `json:"category_id"`
	Image ItemImage `json:"image"`
	PreOrder ItemPreOrder `json:"pre_order"`
	ItemSKU string `json:"item_sku"`
	Condition string `json:"condition"`
	Wholesale []ItemWholesale `json:"wholesale"`
	VideoUploadID []string `json:"video_upload_id"`
	Brand ItemBrand `json:"brand"`
	ItemDangerous int `json:"item_dangerous"`
}

type Dimension struct {
	PackageHeight int `json:"package_height"`
	PackageLength int `json:"package_length"`
	PackageWidth int `json:"package_width"`
}

type LogisticInfo struct {
	SizeID int64 `json:"size_id"`
	ShippingFee float64 `json:"shipping_fee"`
	Enabled bool `json:"enabled"`
	LogisticID int64 `json:"logistic_id"`
	IsFree bool `json:"is_free"`
}

type ItemAttribute struct {
	AttributeID int64 `json:"attribute_id"`
	AttributeValueList []ItemAttributeValue `json:"attribute_value_list"`
}

type ItemAttributeValue struct {
	ValueId int64 `json:"value_id"`
	OriginalValueName string `json:"original_value_name"`
	ValueUnit string `json:"value_unit"`
}

type ItemImage struct {
	ImageIDList []string `json:"image_id_list"`
}

type ItemPreOrder struct {
	IsPreOrder bool `json:"is_pre_order"`
	DaysToShip int `json:"days_to_ship"`
}

type ItemWholesale struct {
	MinCount int `json:"min_count"`
	MaxCount int `json:"max_count"`
	UnitPrice float64 `json:"unit_price"`
}

type ItemBrand struct {
	BrandID int64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
}

type Item struct {
	ItemBase

	ItemID int64 `json:"item_id"`
}

// https://open.shopee.com/documents?module=89&type=1&id=616&version=2
type AddItemResponse struct {
	BaseResponse

	Response Item `json:"response"`
	ItemDangerous int `json:"item_dangerous"` // TODO: why here again?
}

func (s *ProductServiceOp)AddItem(sid int64,item AddItemRequest, tok string)(*AddItemResponse, error) {
	path := "/product/add_item"
	resp := new(AddItemResponse)
	req,err:=StructToMap(item)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type InitTierVariationRequest struct {
	ItemID int64 `json:"item_id"`
	TierVariation []TierVariation `json:"tier_variation"`
	Model []InitTierVariationRequestModel `json:"model"`
}

type TierVariation struct {
	Name string `json:"name"`
	OptionList []TierVariationOption `json:"option_list"`
}

type TierVariationOption struct {
	Option string `json:"option"`
	Image TierVariationOptionImage `json:"image"`
}

type TierVariationOptionImage struct {
	ImageID string `json:"image_id"`
	ImageURL string `json:"image_url"`
}

type InitTierVariationRequestModel struct {
	TierIndex []int `json:"tier_index"`
	NormalStock int `json:"normal_stock"`
	OriginalPrice float64 `json:"original_price"`
	ModelSKU string `json:"model_sku"`
}

type InitTierVariationResponse struct {
	BaseResponse

	Response InitTierVariationResponseData `json:"response"`
}

type InitTierVariationResponseData struct {
	ItemID int64 `json:"item_id"`
	TierVariation []InitTierVariationResponseDataTierVariation `json:"tier_variation"`
	Model []Model `json:"model"`
}

type InitTierVariationResponseDataTierVariation struct {
	Name string `json:"name"`
	OptionList []InitTierVariationResponseDataTierVariationTierVariationOption `json:"option_list"`
}

type InitTierVariationResponseDataTierVariationTierVariationOption struct {
	Option string `json:"option"`
	Image InitTierVariationResponseDataTierVariationTierVariationOptionImage `json:"image"`
}

type InitTierVariationResponseDataTierVariationTierVariationOptionImage struct {
	ImageURL string `json:"image_url"`
}

type Model struct {
	TierIndex []int `json:"tier_index"`
	ModelID int64 `json:"model_id"`
	ModelSKU string `json:"model_sku"`
	StockInfo []StockInfo `json:"stock_info"`
	PriceInfo []PriceInfo `json:"price_info"`
	PromotionID int64 `json:"promotion_id"`
}

type StockInfo struct {
	StockType int `json:"stock_type"`
	NormalStock int `json:"normal_stock"`
	CurrentStock int `json:"current_stock"`
	ReservedStock int `json:"reserved_stock"`
}

type PriceInfo struct {
	OriginalPrice float64 `json:"original_price"`
	CurrentPrice float64 `json:"current_price"`
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"`
	InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price"`
	SipItemPrice float64 `json:"sip_item_price"`
	SipItemPriceSource string `json:"sip_item_price_source"`
}

func (s *ProductServiceOp)InitTierVariation(sid int64,vars InitTierVariationRequest, tok string)(*InitTierVariationResponse, error) {
	path := "/product/init_tier_variation"
	resp := new(InitTierVariationResponse)
	req,err:=StructToMap(vars)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

// https://open.shopee.com/documents?module=89&type=1&id=649&version=2
type AddModelRequest struct {
	ItemID int64 `json:"item_id"`
	ModelList []AddModelRequestModel `json:"model_list"`
}

type AddModelRequestModel struct {
	TierIndex []int `json:"tier_index"` // TODO: doc error?
	NormalStock int `json:"normal_stock"`
	OriginalPrice float64 `json:"original_price"`
	ModelSku string `json:"model_sku"`
}

type AddModelResponse struct {
	BaseResponse

	Response AddModelResponseData `json:"response"`
}

type AddModelResponseData struct {
	Model []Model `json:"model"`
}

func (s *ProductServiceOp)AddModel(sid int64,vars AddModelRequest, tok string)(*AddModelResponse, error) {
	path := "/product/add_model"
	resp := new(AddModelResponse)
	req,err:=StructToMap(vars)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type GetModelListRequest struct {
	ItemID int64 `url:"item_id"`
}

type GetModelListResponse struct {
	BaseResponse

	Response GetModelListResponseData `json:"response"`
}

type GetModelListResponseData struct {
	TierVariation []TierVariation `json:"tier_variation"`
	Model []Model `json:"model"`
}

func (s *ProductServiceOp)	GetModelList(sid, itemID int64, tok string) (*GetModelListResponse,error){
	path := "/product/get_model_list"

	opt:=GetModelListRequest{
		ItemID: itemID,
	}

	resp := new(GetModelListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}