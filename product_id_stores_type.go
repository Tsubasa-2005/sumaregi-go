package sumaregi

type GetProductIDStoresResponse []struct {
	ProductID      string `json:"productId"`
	StoreID        string `json:"storeId"`
	AssignDivision string `json:"assignDivision"`
}

type GetProductIDStoresOpts struct {
	Fields                      string `url:"fields,omitempty"`                         // 検索パラメータ (カンマ区切りで指定可)。レスポンス項目を指定可能。一部ネスト項目は指定不可。
	Sort                        string `url:"sort,omitempty"`                           // 並び順 (カンマ区切りで指定可)。指定可能な項目はドキュメント参照。
	Limit                       int    `url:"limit,omitempty"`                          // 最大取得件数。
	Page                        int    `url:"page,omitempty"`                           // ページ番号 (ページネーション用)。
	StoreID                     int    `url:"store_id,omitempty"`                       // 店舗IDでフィルタリング。
	AssignDivision              string `url:"assign_division,omitempty"`                // 取扱区分 (Enum: 0=販売する, 1=販売しない)。
	WithDetailProductAttributes bool   `url:"with_detail_product_attributes,omitempty"` // 商品属性詳細を含むかどうか (true=含む, false=含まない)。
}
