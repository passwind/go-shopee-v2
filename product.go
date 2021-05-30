package goshopee

type ProductService interface {
	GetCategory(uint64, string, string) (*GetCategoryResponse,error)
	GetBrandList(uint64, uint64, int, int, int, string) (*GetBrandListResponse, error)
	GetDTSLimit(uint64, uint64, string) (*GetDTSLimitResponse, error)
	GetAttributes(uint64, uint64, string, string) (*GetAttributesResponse, error)
	SupportSizeChart(uint64, uint64, string) (*SupportSizeChartResponse, error)
	UpdateSizeChart(uint64, uint64, string, string)(*UpdateSizeChartResponse, error)
	GetItemBaseInfo(uint64, []uint64, string) (*GetItemBaseInfoResponse, error)
	AddItem(uint64, AddItemRequest, string)(*AddItemResponse,error)
	DeleteItem(uint64, uint64, string) (*BaseResponse, error)
	UpdateItem(uint64, UpdateItemRequest, string) (*UpdateItemResponse, error)
	UnlistItem(uint64, UnlistItemRequest, string) (*UnlistItemResponse, error)
	InitTierVariation(uint64, InitTierVariationRequest, string) (*InitTierVariationResponse,error)
	GetModelList(uint64, uint64, string) (*GetModelListResponse, error)
	AddModel(uint64, AddModelRequest, string)(*AddModelResponse, error)
	DeleteModel(uint64, uint64, uint64, string) (*BaseResponse, error)
	UpdateModel(uint64, UpdateModelRequest, string) (*UpdateModelResponse, error)
	UpdatePrice(uint64, UpdatePriceRequest, string) (*UpdatePriceResponse, error)
	UpdateStock(uint64, UpdateStockRequest, string) (*UpdateStockResponse, error)
	CategoryRecommend(uint64, string, string) (*CategoryRecommendResponse, error)
}

type GetCategoryResponse struct {
	BaseResponse

	Response GetCategoryResponseData `json:"response"`
}

type GetCategoryResponseData struct {
	CategoryList []Category `json:"category_list"`
}

type Category struct {
	CategoryID uint64 `json:"category_id"`
	ParentCategoryID uint64 `json:"parent_category_id"`
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

func (s *ProductServiceOp)	GetCategory(sid uint64, lang, tok string) (*GetCategoryResponse,error){
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
	CategoryID uint64 `url:"category_id"`
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
	BrandID uint64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
	DisplayBrandName string `json:"display_brand_name"`
}

func (s *ProductServiceOp)	GetBrandList(sid, cid uint64, status, offset, pageSize int, tok string) (*GetBrandListResponse, error){
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
	CategoryID uint64 `url:"category_id"`
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

func (s *ProductServiceOp)	GetDTSLimit(sid, cid uint64, tok string) (*GetDTSLimitResponse, error){
	path := "/product/get_dts_limit"

	opt:=GetDTSLimitRequest{
		CategoryID: cid,
	}

	resp := new(GetDTSLimitResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type GetAttibutesRequest struct {
	CategoryID uint64 `url:"category_id"`
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
	AttributeID uint64 `json:"attribute_id"`
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
	ValueID uint64 `json:"value_id"`
	OriginalValueName string `json:"original_value_name"`
	DisplayValueName string `json:"display_value_name"`
	ValueUnit string `json:"value_unit"`
	ParentAttributeList []ParentAttribute `json:"parent_attribute_list"`
	ParentBrandList []ParentBrand `json:"parent_brand_list"`
}

type ParentAttribute struct {
	ParentAttributeID uint64 `json:"parent_attribute_id"`
	ParentValueID uint64 `json:"parent_value_id"`
} 

type ParentBrand struct {
	ParentBrandID uint64 `json:"parent_brand_id"`
}

func (s *ProductServiceOp)	GetAttributes(sid, cid uint64, lang, tok string) (*GetAttributesResponse, error){
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
	CategoryID uint64 `url:"category_id"`
}

type SupportSizeChartResponse struct {
	BaseResponse

	Response SupportSizeChartResponseData `json:"response"`
}

type SupportSizeChartResponseData struct {
	SupportSizeChart bool `json:"support_size_chart"`
}

func (s *ProductServiceOp)SupportSizeChart(sid, cid uint64, tok string) (*SupportSizeChartResponse, error){
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

func (s *ProductServiceOp)UpdateSizeChart(sid, itemID uint64, sizeChart,tok string)(*UpdateSizeChartResponse, error) {
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

	OriginalPrice float64 `json:"original_price"`
	NormalStock int `json:"normal_stock"`
	VideoUploadID []string `json:"video_upload_id"`
}

type ItemBase struct {
	CategoryID uint64 `json:"category_id"`
	ItemName string `json:"item_name"`
	Description string `json:"description"`
	ItemSKU string `json:"item_sku"`
	CreateTime uint64 `json:"create_time"`
	UpdateTime uint64 `json:"update_time"`
	AttributeList []ItemAttribute `json:"attribute_list"`
	Image ItemImage `json:"image"`
	Weight float64 `json:"weight"` // int or float? sample is "1.000"? or string? https://open.shopee.com/documents?module=89&type=1&id=616&version=2
	Dimension Dimension `json:"dimension"`
	LogisticInfo []LogisticInfo `json:"logistic_info"`
	PreOrder ItemPreOrder `json:"pre_order"`
	Wholesale []ItemWholesale `json:"wholesale"`
	Condition string `json:"condition"`
	SizeChart string `json:"size_chart"`
	ItemStatus string `json:"item_status"`
	HasModel bool `json:"has_model"`
	PromotionID uint64 `json:"promotion_id"`
	VideoInfo []ItemVideo `json:"video_info"`
	Brand ItemBrand `json:"brand"`
	ItemDangerous int `json:"item_dangerous"`
}

// https://open.shopee.com/documents?module=89&type=1&id=612&version=2
type ItemVideo struct {
	VideoURL string `json:"video_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Duration int `json:"duration"`
}

type Dimension struct {
	PackageHeight int `json:"package_height"`
	PackageLength int `json:"package_length"`
	PackageWidth int `json:"package_width"`
}

type LogisticInfo struct {
	LogisticID uint64 `json:"logistic_id"`
	LogisticName string `json:"logistic_name"`
	Enabled bool `json:"enabled"`
	ShippingFee float64 `json:"shipping_fee"`
	SizeID uint64 `json:"size_id"`
	IsFree bool `json:"is_free"`
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"` // TODO: boolean ? https://open.shopee.com/documents?module=89&type=1&id=612&version=2
}

type ItemAttribute struct {
	AttributeID uint64 `json:"attribute_id"`
	AttributeValueList []ItemAttributeValue `json:"attribute_value_list"`
}

type ItemAttributeValue struct {
	ValueId uint64 `json:"value_id"`
	OriginalValueName string `json:"original_value_name"`
	ValueUnit string `json:"value_unit"`
}

type ItemImage struct {
	ImageIDList []string `json:"image_id_list"`
	ImageURLList []string `json:"image_url_list"`
}

type ItemPreOrder struct {
	IsPreOrder bool `json:"is_pre_order"`
	DaysToShip int `json:"days_to_ship"`
}

type ItemWholesale struct {
	MinCount int `json:"min_count"`
	MaxCount int `json:"max_count"`
	UnitPrice float64 `json:"unit_price"`
	InflatedPriceOfUnitPrice float64 `json:"inflated_price_of_unit_price"`
}

type ItemBrand struct {
	BrandID uint64 `json:"brand_id"`
	OriginalBrandName string `json:"original_brand_name"`
}

type Item struct {
	ItemBase

	ItemID uint64 `json:"item_id"`
	PriceInfo []PriceInfo `json:"price_info"`
	StockInfo []StockInfo `json:"stock_info"`
}

// https://open.shopee.com/documents?module=89&type=1&id=616&version=2
type AddItemResponse struct {
	BaseResponse

	Response AddItemResponseData `json:"response"`
	ItemDangerous int `json:"item_dangerous"` // TODO: why here again? error https://open.shopee.com/documents?module=89&type=1&id=616&version=2
}

type AddItemResponseData struct {
	ItemBase

	ItemID uint64 `json:"item_id"`
	PriceInfo PriceInfo `json:"price_info"`
	StockInfo StockInfo `json:"stock_info"`
}

func (s *ProductServiceOp)AddItem(sid uint64,item AddItemRequest, tok string)(*AddItemResponse, error) {
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
	ItemID uint64 `json:"item_id"`
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
	ItemID uint64 `json:"item_id"`
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
	ModelID uint64 `json:"model_id"`
	ModelSKU string `json:"model_sku"`
	StockInfo []StockInfo `json:"stock_info"`
	PriceInfo []PriceInfo `json:"price_info"`
	PromotionID uint64 `json:"promotion_id"`
}

type StockInfo struct {
	StockType int `json:"stock_type"`
	StockLocationID string `json:"stock_location_id"`
	NormalStock int `json:"normal_stock"`
	CurrentStock int `json:"current_stock"`
	ReservedStock int `json:"reserved_stock"`
}

type PriceInfo struct {
	Currency string `json:"currency"`
	OriginalPrice float64 `json:"original_price"`
	CurrentPrice float64 `json:"current_price"`
	InflatedPriceOfOriginalPrice float64 `json:"inflated_price_of_original_price"`
	InflatedPriceOfCurrentPrice float64 `json:"inflated_price_of_current_price"`
	SipItemPrice float64 `json:"sip_item_price"`
	SipItemPriceSource string `json:"sip_item_price_source"`
}

func (s *ProductServiceOp)InitTierVariation(sid uint64,vars InitTierVariationRequest, tok string)(*InitTierVariationResponse, error) {
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
	ItemID uint64 `json:"item_id"`
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

func (s *ProductServiceOp)AddModel(sid uint64,vars AddModelRequest, tok string)(*AddModelResponse, error) {
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
	ItemID uint64 `url:"item_id"`
}

type GetModelListResponse struct {
	BaseResponse

	Response GetModelListResponseData `json:"response"`
}

type GetModelListResponseData struct {
	TierVariation []TierVariation `json:"tier_variation"`
	Model []Model `json:"model"`
}

func (s *ProductServiceOp)	GetModelList(sid, itemID uint64, tok string) (*GetModelListResponse,error){
	path := "/product/get_model_list"

	opt:=GetModelListRequest{
		ItemID: itemID,
	}

	resp := new(GetModelListResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

type GetItemBaseInfoRequest struct {
	ItemIDList []uint64 `url:"item_id_list"`
}

type GetItemBaseInfoResponse struct {
	BaseResponse

	Response GetItemBaseInfoResponseData `json:"response"`
}

type GetItemBaseInfoResponseData struct {
	ItemList []Item `json:"item_list"`
}

func (s *ProductServiceOp)	GetItemBaseInfo(sid uint64, itemIDs []uint64, tok string) (*GetItemBaseInfoResponse,error){
	path := "/product/get_item_base_info"

	opt:=GetItemBaseInfoRequest{
		ItemIDList: itemIDs,
	}

	resp := new(GetItemBaseInfoResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}

func (s *ProductServiceOp)DeleteItem(sid, itemID uint64, tok string)(*BaseResponse, error) {
	path := "/product/delete_item"
	resp := new(BaseResponse)
	req:=map[string]interface{}{
		"item_id": itemID,
	}
	err := s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type UpdateItemRequest struct {
	ItemBase

	ItemID uint64 `json:"item_id"`
	OriginalPrice float64 `json:"original_price"`
	NormalStock int `json:"normal_stock"`
	VideoUploadID []string `json:"video_upload_id"`
}

type UpdateItemResponse struct {
	BaseResponse

	Response UpdateItemResponseData `json:"response"`
}

type UpdateItemResponseData struct {
	ItemBase

	ItemID uint64 `json:"item_id"`
	PriceInfo PriceInfo `json:"price_info"`
	StockInfo StockInfo `json:"stock_info"`
}

func (s *ProductServiceOp)UpdateItem(sid uint64,item UpdateItemRequest, tok string)(*UpdateItemResponse, error) {
	path := "/product/update_item"
	resp := new(UpdateItemResponse)
	req,err:=StructToMap(item)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type UnlistItemRequest struct {
	ItemList []UnlistItemReqData `json:"item_list"`
}

type UnlistItemReqData struct {
	ItemID uint64 `json:"item_id"`
	Unlist bool `json:"unlist"`
}

type UnlistItemResponse struct {
	BaseResponse

	Response UnlistItemResponseData `json:"response"`
}

type UnlistItemResponseData struct {
	FailureList []UnlistItemResponseDataFail `json:"failure_list"`
	SuccessList []UnlistItemResponseDataSuccess `json:"success_list"`
}

type UnlistItemResponseDataFail struct {
	ItemID uint64 `json:"item_id"`
	FailedReason string `json:"failed_reason"`
}

type UnlistItemResponseDataSuccess struct {
	ItemID uint64 `json:"item_id"`
	Unlist bool `json:"unlist"` 
}

func (s *ProductServiceOp)UnlistItem(sid uint64, data UnlistItemRequest, tok string)(*UnlistItemResponse, error) {
	path := "/product/unlist_item"
	resp := new(UnlistItemResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

func (s *ProductServiceOp)DeleteModel(sid, itemID, modelID uint64, tok string) (*BaseResponse, error) {
	path := "/product/delete_model"
	resp := new(BaseResponse)
	req:=map[string]interface{}{
		"item_id": itemID,
		"model_id": modelID,
	}

	err := s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type UpdateModelRequest struct {
	ItemID uint64 `json:"item_id"`
	Model []Model `json:"model"`
}

// TODO: response is too simple? https://open.shopee.com/documents?module=89&type=1&id=648&version=2
type UpdateModelResponse struct {
	BaseResponse
}

func (s *ProductServiceOp)UpdateModel(sid uint64, data UpdateModelRequest, tok string) (*UpdateModelResponse, error){
	path := "/product/update_model"
	resp := new(UpdateModelResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type UpdatePriceRequest struct {
	ItemID uint64 `json:"item_id"`
	PriceList []UpdatePriceRequestData `json:"price_list"`
}

type UpdatePriceRequestData struct {
	ModelID uint64 `json:"model_id"`
	OriginalPrice float64 `json:"original_price"`
}

type UpdatePriceResponse struct {
	BaseResponse

	Response UpdatePriceResponseData `json:"response"`
}

type UpdatePriceResponseData struct {
	FailureList []UpdatePriceResponseDataFail `json:"failure_list"`
	SuccessList []UpdatePriceResponseDataSuccess `json:"success_list"`
}

type UpdatePriceResponseDataFail struct {
	ModelID uint64 `json:"model_id"`
	FailedReason string `json:"failed_reason"`
}

type UpdatePriceResponseDataSuccess struct {
	ModelID uint64 `json:"model_id"`
	OriginalPrice float64 `json:"original_price"`
}

func (s *ProductServiceOp)UpdatePrice(sid uint64, data UpdatePriceRequest, tok string) (*UpdatePriceResponse, error) {
	path := "/product/update_price"
	resp := new(UpdatePriceResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type UpdateStockRequest struct {
	ItemID uint64 `json:"item_id"`
	StockList []UpdateStockRequestData `json:"stock_list"`
}

type UpdateStockRequestData struct {
	ModelID uint64 `json:"model_id"`
	NormalStock int `json:"normal_stock"`
}

type UpdateStockResponse struct {
	BaseResponse

	Response UpdateStockResponseData `json:"response"`
}

type UpdateStockResponseData struct {
	FailureList []UpdateStockResponseDataFail `json:"failure_list"`
	SuccessList []UpdateStockResponseDataSuccess `json:"success_list"`
}

type UpdateStockResponseDataFail struct {
	ModelID uint64 `json:"model_id"`
	FailedReason string `json:"failed_reason"`
}

type UpdateStockResponseDataSuccess struct {
	ModelID uint64 `json:"model_id"`
	NormalStock int `json:"normal_stock"`
}

func (s *ProductServiceOp)UpdateStock(sid uint64, data UpdateStockRequest, tok string) (*UpdateStockResponse, error) {
	path := "/product/update_stock"
	resp := new(UpdateStockResponse)
	req,err:=StructToMap(data)
	if err!=nil {
		return nil,err
	}
	err = s.client.WithShop(sid,tok).Post(path, req, resp)
	return resp, err
}

type CategoryRecommendRequest struct {
	ItemName string `json:"item_name"`
}

type CategoryRecommendResponse struct {
	BaseResponse

	Response CategoryRecommendResponseData `json:"response"`
}

type CategoryRecommendResponseData struct {
	CategoryID []uint64 `json:"category_id"` // TODO: sample is item_list? error
}

// https://open.shopee.com/documents?module=89&type=1&id=702&version=2
func (s *ProductServiceOp)	CategoryRecommend(sid uint64, itemName, tok string) (*CategoryRecommendResponse,error){
	path := "/product/category_recommand" // TODO: recommand or recommend?

	opt:=CategoryRecommendRequest{
		ItemName: itemName,
	}

	resp := new(CategoryRecommendResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}