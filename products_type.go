package sumaregi

type GetProductsResponse []struct {
	ProductID               string `json:"productId"`
	CategoryID              string `json:"categoryId"`
	ProductCode             string `json:"productCode"`
	ProductName             string `json:"productName"`
	ProductKana             string `json:"productKana"`
	TaxDivision             string `json:"taxDivision"`
	ProductPriceDivision    string `json:"productPriceDivision"`
	Price                   string `json:"price"`
	CustomerPrice           string `json:"customerPrice"`
	Cost                    string `json:"cost"`
	Attribute               string `json:"attribute"`
	Description             string `json:"description"`
	CatchCopy               string `json:"catchCopy"`
	Size                    string `json:"size"`
	Color                   string `json:"color"`
	Tag                     string `json:"tag"`
	GroupCode               string `json:"groupCode"`
	URL                     string `json:"url"`
	PrintReceiptProductName string `json:"printReceiptProductName"`
	DisplaySequence         string `json:"displaySequence"`
	SalesDivision           string `json:"salesDivision"`
	StockControlDivision    string `json:"stockControlDivision"`
	DisplayFlag             string `json:"displayFlag"`
	Division                string `json:"division"`
	ProductOptionGroupID    string `json:"productOptionGroupId"`
	PointNotApplicable      string `json:"pointNotApplicable"`
	TaxFreeDivision         string `json:"taxFreeDivision"`
	SupplierProductNo       string `json:"supplierProductNo"`
	CalcDiscount            string `json:"calcDiscount"`
	StaffDiscountRate       string `json:"staffDiscountRate"`
	UseCategoryReduceTax    string `json:"useCategoryReduceTax"`
	ReduceTaxID             string `json:"reduceTaxId"`
	ReduceTaxPrice          string `json:"reduceTaxPrice"`
	ReduceTaxCustomerPrice  string `json:"reduceTaxCustomerPrice"`
	OrderPoint              string `json:"orderPoint"`
	PurchaseCost            string `json:"purchaseCost"`
	AppStartDateTime        string `json:"appStartDateTime"`
	InsDateTime             string `json:"insDateTime"`
	UpdDateTime             string `json:"updDateTime"`
}

type PostProductsResponse struct {
	ProductID               string `json:"productId"`
	CategoryID              string `json:"categoryId"`
	ProductCode             string `json:"productCode"`
	ProductName             string `json:"productName"`
	ProductKana             string `json:"productKana"`
	TaxDivision             int    `json:"taxDivision"`
	ProductPriceDivision    int    `json:"productPriceDivision"`
	Price                   string `json:"price"`
	CustomerPrice           string `json:"customerPrice"`
	Cost                    string `json:"cost"`
	Attribute               string `json:"attribute"`
	Description             string `json:"description"`
	CatchCopy               string `json:"catchCopy"`
	Size                    string `json:"size"`
	Color                   string `json:"color"`
	Tag                     string `json:"tag"`
	GroupCode               string `json:"groupCode"`
	URL                     string `json:"url"`
	PrintReceiptProductName string `json:"printReceiptProductName"`
	DisplaySequence         string `json:"displaySequence"`
	DisplayFlag             int    `json:"displayFlag"`
	Division                int    `json:"division"`
	ProductOptionGroupID    string `json:"productOptionGroupId"`
	SalesDivision           int    `json:"salesDivision"`
	StockControlDivision    int    `json:"stockControlDivision"`
	PointNotApplicable      int    `json:"pointNotApplicable"`
	TaxFreeDivision         int    `json:"taxFreeDivision"`
	CalcDiscount            int    `json:"calcDiscount"`
	StaffDiscountRate       string `json:"staffDiscountRate"`
	UseCategoryReduceTax    int    `json:"useCategoryReduceTax"`
	ReduceTaxID             string `json:"reduceTaxId"`
	ReduceTaxPrice          string `json:"reduceTaxPrice"`
	ReduceTaxCustomerPrice  string `json:"reduceTaxCustomerPrice"`
	OrderPoint              string `json:"orderPoint"`
	PurchaseCost            string `json:"purchaseCost"`
	SupplierProductNo       string `json:"supplierProductNo"`
	AppStartDateTime        string `json:"appStartDateTime"`
	InsDateTime             string `json:"insDateTime"`
	UpdDateTime             string `json:"updDateTime"`
	Prices                  []struct {
		ProductID     string      `json:"productId"`
		StoreID       interface{} `json:"storeId"`
		PriceDivision string      `json:"priceDivision"`
		StartDate     interface{} `json:"startDate"`
		EndDate       interface{} `json:"endDate"`
		Price         interface{} `json:"price"`
	} `json:"prices"`
	ReserveItems []struct {
		ProductID string      `json:"productId"`
		No        interface{} `json:"no"`
		Value     string      `json:"value"`
	} `json:"reserveItems"`
	Stores []struct {
		ProductID            string      `json:"productId"`
		StoreID              interface{} `json:"storeId"`
		ProductOptionGroupID string      `json:"productOptionGroupId"`
		AssignDivision       string      `json:"assignDivision"`
	} `json:"stores"`
	InventoryReservations []struct {
		ProductID            string      `json:"productId"`
		ReservationProductID interface{} `json:"reservationProductId"`
		ReservationAmount    interface{} `json:"reservationAmount"`
	} `json:"inventoryReservations"`
	AttributeItems []struct {
		ProductID string `json:"productId"`
		Code      string `json:"code"`
		No        string `json:"no"`
		Name      string `json:"name"`
	} `json:"attributeItems"`
	OrderSetting struct {
		ProductID             string `json:"productId"`
		ContinuationDivision  string `json:"continuationDivision"`
		OrderStatusDivision   int    `json:"orderStatusDivision"`
		OrderNoReasonDivision string `json:"orderNoReasonDivision"`
		OrderLimitAmount      string `json:"orderLimitAmount"`
		OrderSupplierEditable int    `json:"orderSupplierEditable"`
		PbDivision            int    `json:"pbDivision"`
		DisplayFlag           int    `json:"displayFlag"`
		OrderUnit             struct {
			Division int    `json:"division"`
			Num      string `json:"num"`
			Name     string `json:"name"`
		} `json:"orderUnit"`
		Stores []struct {
			StoreID          string `json:"storeId"`
			OrderLimitAmount string `json:"orderLimitAmount"`
			DisplayFlag      int    `json:"displayFlag"`
		} `json:"stores"`
	} `json:"orderSetting"`
}

type GetProductsOpts struct {
	Fields               []string `url:"fields,omitempty"`
	Sort                 string   `url:"sort,omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Page                 int      `url:"page,omitempty"`
	CategoryID           int      `url:"category_id,omitempty"`
	ProductCode          string   `url:"product_code,omitempty"`
	GroupCode            string   `url:"group_code,omitempty"`
	DisplayFlag          string   `url:"display_flag,omitempty"`           // Enum: 0 1 (0: 表示しない, 1: 表示する)
	Division             string   `url:"division,omitempty"`               // Enum: 0 1 2 (0: 通常商品, 1: 回数券, 2: オプション商品)
	SalesDivision        string   `url:"sales_division,omitempty"`         // Enum: 0 1 (0: 売上対象, 1: 売上対象外)
	StockControlDivision string   `url:"stock_control_division,omitempty"` // Enum: 0 1 (0: 在庫管理対象, 1: 在庫管理対象外)
	SupplierProductNo    string   `url:"supplier_product_no,omitempty"`
	UpdDateTimeFrom      string   `url:"upd_date_time-from,omitempty"`
	UpdDateTimeTo        string   `url:"upd_date_time-to,omitempty"`
}

type PostProductsParams struct {
	CategoryID              string `json:"categoryId"`
	ProductCode             string `json:"productCode"`
	ProductName             string `json:"productName"`
	ProductKana             string `json:"productKana"`
	TaxDivision             string `json:"taxDivision"`
	ProductPriceDivision    string `json:"productPriceDivision"`
	Price                   string `json:"price"`
	CustomerPrice           string `json:"customerPrice"`
	Cost                    string `json:"cost"`
	Attribute               string `json:"attribute"`
	Description             string `json:"description"`
	CatchCopy               string `json:"catchCopy"`
	Size                    string `json:"size"`
	Color                   string `json:"color"`
	Tag                     string `json:"tag"`
	GroupCode               string `json:"groupCode"`
	URL                     string `json:"url"`
	PrintReceiptProductName string `json:"printReceiptProductName"`
	DisplaySequence         string `json:"displaySequence"`
	DisplayFlag             string `json:"displayFlag"`
	Division                string `json:"division"`
	ProductOptionGroupID    int    `json:"productOptionGroupId"`
	SalesDivision           string `json:"salesDivision"`
	StockControlDivision    string `json:"stockControlDivision"`
	PointNotApplicable      string `json:"pointNotApplicable"`
	TaxFreeDivision         string `json:"taxFreeDivision"`
	CalcDiscount            string `json:"calcDiscount"`
	StaffDiscountRate       string `json:"staffDiscountRate"`
	UseCategoryReduceTax    string `json:"useCategoryReduceTax"`
	ReduceTaxID             string `json:"reduceTaxId"`
	ReduceTaxPrice          string `json:"reduceTaxPrice"`
	ReduceTaxCustomerPrice  string `json:"reduceTaxCustomerPrice"`
	OrderPoint              string `json:"orderPoint"`
	PurchaseCost            string `json:"purchaseCost"`
	SupplierProductNo       string `json:"supplierProductNo"`
	AppStartDateTime        string `json:"appStartDateTime"`
	ReserveItems            []struct {
		No    string `json:"no"`
		Value string `json:"value"`
	} `json:"reserveItems"`
	Prices []struct {
		StoreID       string `json:"storeId"`
		PriceDivision int    `json:"priceDivision"`
		StartDate     string `json:"startDate"`
		EndDate       string `json:"endDate"`
		Price         string `json:"price"`
	} `json:"prices"`
	Stores []struct {
		StoreID              string `json:"storeId"`
		ProductOptionGroupID string `json:"productOptionGroupId"`
		AssignDivision       string `json:"assignDivision"`
	} `json:"stores"`
	InventoryReservations []struct {
		ReservationProductID string `json:"reservationProductId"`
		ReservationAmount    string `json:"reservationAmount"`
	} `json:"inventoryReservations"`
	AttributeItems []struct {
		Code string `json:"code"`
		No   string `json:"no"`
	} `json:"attributeItems"`
	OrderSetting struct {
		ContinuationDivision  interface{} `json:"continuationDivision"`
		OrderStatusDivision   string      `json:"orderStatusDivision"`
		OrderNoReasonDivision string      `json:"orderNoReasonDivision"`
		OrderUnit             struct {
			Division string `json:"division"`
			Num      string `json:"num"`
			Name     string `json:"name"`
		} `json:"orderUnit"`
		OrderLimitAmount      string `json:"orderLimitAmount"`
		OrderSupplierEditable string `json:"orderSupplierEditable"`
		PbDivision            string `json:"pbDivision"`
		DisplayFlag           string `json:"displayFlag"`
		Stores                []struct {
			StoreID          string `json:"storeId"`
			OrderLimitAmount string `json:"orderLimitAmount"`
			DisplayFlag      string `json:"displayFlag"`
		} `json:"stores"`
	} `json:"orderSetting"`
}
