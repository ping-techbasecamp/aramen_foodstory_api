package qrcodeModel

//    foodStoryOrder, err := UnmarshalFoodStoryOrder(bytes)
//    bytes, err = foodStoryOrder.Marshal()

import "encoding/json"

func UnmarshalFoodStoryOrder(data []byte) (FoodStoryOrder, error) {
	var r FoodStoryOrder
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FoodStoryOrder) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FoodStoryOrder struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Success int    `json:"success"`
	Data    Data   `json:"data"`
}

type Data struct {
	QrcodeID              string         `json:"qrcode_id"`
	UsedQrcode            bool           `json:"used_qrcode"`
	OrderID               string         `json:"order_id"`
	OrderStatus           *string        `json:"order_status"`
	BranchName            *string        `json:"branch_name"`
	OrderChannel          *string        `json:"order_channel"`
	DeliveryPartner       *interface{}   `json:"delivery_partner"`
	DeliveryOrderRef      *interface{}   `json:"delivery_order_ref"`
	OrderMenu             *[]interface{} `json:"order_menu"`
	DeliveryFee           *[]DeliveryFee `json:"delivery_fee"`
	Promotion             *Promotion     `json:"promotion"`
	Discount              *Discount      `json:"discount"`
	Voucher               *interface{}   `json:"voucher"`
	Tip                   *float64       `json:"tip"`
	SubTotal              *float64       `json:"sub_total"`
	VatInclude            *float64       `json:"vat_include"`
	TotalPrice            *float64       `json:"total_price"`
	TotalDiscount         *float64       `json:"total_discount"`
	TotalPaid             *float64       `json:"total_paid"`
	PaymentChannel        *string        `json:"payment_channel"`
	DrawerID              *string        `json:"drawer_id"`
	PosID                 *interface{}   `json:"pos_id"`
	Staff                 *string        `json:"staff"`
	OrderCreatedTimestamp *string        `json:"order_created_timestamp"`
	OrderClosedTimestamp  *string        `json:"order_closed_timestamp"`
}

type DeliveryFee struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
}

type Discount struct {
	Discounts     *[]interface{} `json:"discounts"`
	TotalDiscount *float64       `json:"total_discount"`
}

type Promotion struct {
	Promotions             *[]interface{} `json:"promotions"`
	PromotionTotalDiscount *float64       `json:"promotion_total_discount"`
}
