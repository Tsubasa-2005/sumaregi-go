package sumaregi

type GetProductIDStoresResponse []struct {
	ProductID      string `json:"productId"`
	StoreID        string `json:"storeId"`
	AssignDivision string `json:"assignDivision"`
}

type GetProductIDStoresOpts struct {
	Fields                      string `url:"fields,omitempty"`
	Sort                        string `url:"sort,omitempty"`
	Limit                       int    `url:"limit,omitempty"`
	Page                        int    `url:"page,omitempty"`
	StoreID                     int    `url:"store_id,omitempty"`
	AssignDivision              string `url:"assign_division,omitempty"`
	WithDetailProductAttributes bool   `url:"with_detail_product_attributes,omitempty"`
}
