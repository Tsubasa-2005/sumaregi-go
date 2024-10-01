package sumaregi

type GetStockResponse []struct {
	StoreID            string `json:"storeId"`
	ProductID          string `json:"productId"`
	StockAmount        string `json:"stockAmount"`
	LayawayStockAmount string `json:"layawayStockAmount"`
	UpdDateTime        string `json:"updDateTime"`
}

type GetStockOpts struct {
	Fields          []string `url:"fields,omitempty"`
	Sort            string   `url:"sort,omitempty"`
	Limit           int      `url:"limit,omitempty"`
	Page            int      `url:"page,omitempty"`
	StoreID         int      `url:"store_id,omitempty"`
	ProductID       string   `url:"product_id,omitempty"`
	UpdDateTimeFrom string   `url:"upd_date_time-from,omitempty"`
	UpdDateTimeTo   string   `url:"upd_date_time-to,omitempty"`
}
