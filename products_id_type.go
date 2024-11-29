package sumaregi

type GetProductResponse struct {
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
	ReserveItems            []struct {
		ProductID string `json:"productId"`
		No        string `json:"no"`
		Value     string `json:"value"`
	} `json:"reserveItems"`
	Prices []struct {
		ProductID     string `json:"productId"`
		StoreID       string `json:"storeId"`
		PriceDivision string `json:"priceDivision"`
		StartDate     string `json:"startDate"`
		EndDate       string `json:"endDate"`
		Price         string `json:"price"`
	} `json:"prices"`
	Stores []struct {
		ProductID            string `json:"productId"`
		StoreID              string `json:"storeId"`
		ProductOptionGroupID string `json:"productOptionGroupId"`
		AssignDivision       string `json:"assignDivision"`
	} `json:"stores"`
	InventoryReservations []struct {
		ProductID            string `json:"productId"`
		ReservationProductID string `json:"reservationProductId"`
		ReservationAmount    string `json:"reservationAmount"`
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

type GetProductOpt struct {
	Fields                    []string `url:"fields,omitempty"`
	WithPrices                string   `url:"with_prices,omitempty"`
	WithReserveItems          string   `url:"with_reserve_items,omitempty"`
	WithStores                string   `url:"with_stores,omitempty"`
	WithInventoryReservations string   `url:"with_inventory_reservations,omitempty"`
	WithAttributeItems        string   `url:"with_attribute_items,omitempty"`
	WithOrderSetting          string   `url:"with_order_setting,omitempty"`
}
