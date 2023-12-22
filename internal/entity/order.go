package entity

import "time"

type CreateOrderReq struct {
	UserId int    `json:"user_id"`
	SkuId  int    `json:"sku_id"`
	Serial string `json:"serial"`
	Count  int    `json:"count"`
}

type CreateOrderResp struct {
	OrderId string `json:"order_id"`
}

type OrderListReq struct {
	UserId int `json:"user_id"`
}
type OrderListResp struct {
	OrderList []OrderInfo `json:"order_list"`
}

type GetOrderDetailReq struct {
	OrderId string `json:"order_id" binding:"required"`
}
type GetOrderDetailResp struct {
	OrderInfo
}

type OrderInfo struct {
	Id          int               `json:"id"`
	OrderId     string            `json:"order_id"`
	UserId      int               `json:"user_id"`
	Status      int               `json:"status"`
	Serial      string            `json:"serial"`
	TotalAmount int               `json:"total_amount"`
	GoodsNum    int               `json:"goods_num"`
	PayId       string            `json:"pay_id"`
	PayStatus   int               `json:"pay_status"`
	PayType     int               `json:"pay_type"`
	PayTime     time.Time         `json:"pay_time"`
	CancelTime  time.Time         `json:"cancel_time"`
	CreateAt    time.Time         `json:"create_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	OrderDetail []OrderDetailInfo `json:"order_detail"`
}

type OrderDetailInfo struct {
	OrderId           string    `json:"order_id"`
	Id                int       `json:"order_detail_id"`
	SkuId             int       `json:"sku_id"`
	Price             int       `json:"price"`
	Num               int       `json:"num"`
	Status            int       `json:"status"`
	CreateAt          time.Time `json:"create_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	OrderDetailExpand string    `json:"order_detail_expand"`
}
