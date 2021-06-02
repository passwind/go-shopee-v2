package goshopee

import "strings"

type OrderService interface{
	GetOrderDetail(uint64, []string, []string, string) (*GetOrderDetailResponse, error)
}

type OrderServiceOp struct {
	client *Client
}

// https://open.shopee.com/documents?module=94&type=1&id=557&version=2
type GetOrderDetailRequest struct {
	OrderSNList string `url:"order_sn_list"`
	ResponseOptionalFields string `url:"response_optional_fields"`
}

type GetOrderDetailResponse struct {
	BaseResponse

	Response GetOrderDetailResponseData `json:"response"`
}

type GetOrderDetailResponseData struct {
	OrderList []Order `json:"order_list"`
}

type Order struct {
	OrderSN string `json:"order_sn"`
	Region string `json:"region"`
	Currency string `json:"currency"`
	COD bool `json:"cod"`
	TotalAmount float64 `json:"total_amount"`
	OrderStatus string `json:"order_status"`
	ShippingCarrier string `json:"shipping_carrier"`
	PaymentMethod string `json:"payment_method"`
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
	MessageToSeller string `json:"message_to_seller"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
	DaysToShip int `json:"days_to_ship"`
	ShipByDate int `json:"ship_by_date"`
	BuyerUserID uint64 `json:"buyer_user_id"`
	BuyerUsername string `json:"buyer_username"`
	RecipientAddress Address `json:"recipient_address"`
	ActualShippingFee float64 `json:"actual_shipping_fee"`
	GoodsToDeclare bool `json:"goods_to_declare"`
	Note string `json:"note"`
	NoteUpdateTime int64 `json:"note_update_time"`
	ItemList []OrderItem `json:"item_list"`
	PayTime int64 `json:"pay_time"`
	Dropshipper string `json:"dropshipper"`
	CreditCardNumber string `json:"credit_card_number"`
	DropshipperPhone string `json:"dropshipper_phone"`
	SplitUp bool `json:"split_up"`
	BuyerCancelReason string `json:"buyer_cancel_reason"`
	CancelBy string `json:"cancel_by"`
	CancelReason string `json:"cancel_reason"`
	ActualShippingFeeConfirmed bool `json:"actual_shipping_fee_confirmed"`
	BuyerCpfID string `json:"buyer_cpf_id"`
	FulfillmentFlag string `json:"fulfillment_flag"`
	PickupDoneTime int64 `json:"pickup_done_time"`
	PackageList []OrderPackage `json:"package_list"`
	InvoiceData Invoice `json:"invoice_data"`
	CheckoutShippingCarrier string `json:"checkout_shipping_carrier"`
}

type Invoice struct {
	Number string `json:"number"`
	SeriesNumber string `json:"series_number"`
	AccessKey string `json:"access_key"`
	IssueDate int64 `json:"issue_date"`
	TotalValue float64 `json:"total_value"`
	ProductsTotalValue float64 `json:"products_total_value"`
	TaxCode string `json:"tax_code"`
}

type OrderPackage struct {
	PackageNumber string `json:"package_number"`
	LogisticsStatus string `json:"logistics_status"`
	ShippingCarrier string `json:"shipping_carrier"`
	ItemList []OrderPackageItem `json:"item_list"`
}

type OrderPackageItem struct {
	ItemID uint64 `json:"item_id"`
	ModelID uint64 `json:"model_id"`
}

type OrderItem struct {
	ItemID uint64 `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemSKU string `json:"item_sku"`
	ModelID uint64 `json:"model_id"`
	ModelName string `json:"model_name"`
	ModelSKU string `json:"model_sku"`
	ModelQuantityPurchased int `json:"model_quantity_purchased"`
	ModelOriginalPrice float64 `json:"model_original_price"`
	ModelDiscountedPrice float64 `json:"model_discounted_price"`
	Wholesale bool `json:"wholesale"`
	Weight float64 `json:"weight"`
	AddOnDeal bool `json:"add_on_deal"`
	MainItem bool `json:"main_item"`
	AddOnDealID uint64 `json:"add_on_deal_id"`
	PromotionType string `json:"promotion_type"`
	PromotionID uint64 `json:"promotion_id"`
}

type Address struct {
	Name string `json:"name"`
	Phone string `json:"phone"`
	Town string `json:"town"`
	District string `json:"district"`
	City string `json:"city"`
	State string `json:"state"`
	Region string `json:"region"`
	Zipcode string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

func (s *OrderServiceOp)GetOrderDetail(sid uint64, snlist, fields []string, tok string) (*GetOrderDetailResponse, error) {
	path := "/order/get_order_detail"

	opt:=GetOrderDetailRequest{
		OrderSNList: strings.Join(snlist,","),
		ResponseOptionalFields: strings.Join(fields,","),
	}

	resp := new(GetOrderDetailResponse)
	err := s.client.WithShop(sid,tok).Get(path, resp, opt)
	return resp, err
}