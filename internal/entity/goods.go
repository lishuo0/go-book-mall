package entity

type CreateGoodsReq struct {
	GoodsInfo
}

type CreateGoodsResp struct {
	GoodsId int `json:"goods_id"`
}

type UpdateGoodsReq struct {
	GoodsInfo
}

type UpdateGoodsResp struct {
}

type DeleteGoodsReq struct {
	GoodsInfo
}

type DeleteGoodsResp struct {
	GoodsId int `json:"goods_id"`
}

type GetGoodsDetailReq struct {
	GoodsId int `json:"goods_id" binding:"required"`
}

type GetGoodsDetailResp struct {
	GoodsInfo
}

type GoodsListReq struct {
	CategoryId int `json:"category_id" binding:"required"`
}

type GoodsListResp struct {
	GoodsList []GoodsInfo `json:"goods_list"`
}

type GoodsInfo struct {
	Id          int            `json:"id"`
	Name        string         `json:"name" binding:"required"`
	SmallImage  string         `json:"small_image,omitempty" binding:"required"`
	DetailImage string         `json:"detail_image,omitempty" binding:"required"`
	CategoryId  int            `json:"category_id" binding:"required"`
	Description string         `json:"description" binding:"required"`
	Status      int            `json:"status"`
	SkuInfo     []GoodsSkuInfo `json:"sku_infos,omitempty"`
}

type GoodsSkuInfo struct {
	Id            int   `json:"id"`
	GoodsId       int   `json:"-"`
	AttIds        []int `json:"att_ids" binding:"required"`
	SpendPrice    int   `json:"spend_price" binding:"required"`
	Price         int   `json:"price" binding:"required"`
	DiscountPrice int   `json:"discount_price" binding:"required"`
	Allstore      int   `json:"all_store" binding:"required"`
	Leftstore     int   `json:"left_store" binding:"required"`
}
