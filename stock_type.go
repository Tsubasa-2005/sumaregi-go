package sumaregi

type GetStockResponse []struct {
	StoreID            string `json:"storeId"`
	ProductID          string `json:"productId"`
	StockAmount        string `json:"stockAmount"`
	LayawayStockAmount string `json:"layawayStockAmount"`
	UpdDateTime        string `json:"updDateTime"`
}

type GetStockOpts struct {
	Fields          []string `url:"fields,omitempty"`             // 検索パラメータ (カンマ区切りで指定可)。レスポンス項目を指定可能。一部ネスト項目は指定不可。
	Sort            string   `url:"sort,omitempty"`               // 並び順 (カンマ区切りで指定可)。指定可能な項目はドキュメント参照。
	Limit           int      `url:"limit,omitempty"`              // 最大取得件数。
	Page            int      `url:"page,omitempty"`               // ページ番号 (ページネーション用)。
	StoreID         int      `url:"store_id,omitempty"`           // 店舗ID。ユーザーアクセストークン利用時はユーザーの所属店舗IDを指定。
	ProductID       string   `url:"product_id,omitempty"`         // 商品IDで在庫情報をフィルタリング。
	UpdDateTimeFrom string   `url:"upd_date_time-from,omitempty"` // 更新日時の開始を指定 (フォーマット: [YYYY-MM-DDThh:mm:ssTZD], 最大31日)。
	UpdDateTimeTo   string   `url:"upd_date_time-to,omitempty"`   // 更新日時の終了を指定 (フォーマット: [YYYY-MM-DDThh:mm:ssTZD], 最大31日)。
}
