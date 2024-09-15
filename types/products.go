package types

import (
	"time"
)

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
	Fields                      []string   `url:"fields,omitempty"`
	Sort                        string     `url:"sort,omitempty"`
	Limit                       int        `url:"limit,omitempty"`
	Page                        int        `url:"page,omitempty"`
	TransactionHeadIDFrom       string     `url:"transaction_head_id-from,omitempty"`
	TransactionHeadIDTo         string     `url:"transaction_head_id-to,omitempty"`
	TransactionDateTimeFrom     *time.Time `url:"transaction_date_time-from,omitempty"`
	TransactionDateTimeTo       *time.Time `url:"transaction_date_time-to,omitempty"`
	TransactionHeadDivision     int        `url:"transaction_head_division,omitempty"`
	StoreID                     int        `url:"store_id,omitempty"`
	TerminalTranDateTimeFrom    *time.Time `url:"terminal_tran_date_time-from,omitempty"`
	TerminalTranDateTimeTo      *time.Time `url:"terminal_tran_date_time-to,omitempty"`
	AdjustmentDateTime          *time.Time `url:"adjustment_date_time,omitempty"`
	SumDate                     *time.Time `url:"sum_date,omitempty"`
	SumDateFrom                 *time.Time `url:"sum_date-from,omitempty"`
	SumDateTo                   *time.Time `url:"sum_date-to,omitempty"`
	CustomerCode                string     `url:"customer_code,omitempty"`
	TransactionUUID             string     `url:"transaction_uuid,omitempty"`
	Barcode                     string     `url:"barcode,omitempty"`
	UpdDateTimeFrom             *time.Time `url:"upd_date_time-from,omitempty"`
	UpdDateTimeTo               *time.Time `url:"upd_date_time-to,omitempty"`
	WithDetails                 string     `url:"with_details,omitempty"`
	WithDepositOthers           string     `url:"with_deposit_others,omitempty"`
	WithLayaway                 string     `url:"with_layaway,omitempty"`
	WithLayaways                string     `url:"with_layaways,omitempty"`
	WithLayawayPickUp           string     `url:"with_layaway_pick_up,omitempty"`
	WithLayawayPickUps          string     `url:"with_layaway_pick_ups,omitempty"`
	WithMoneyControl            string     `url:"with_money_control,omitempty"`
	WithDetailProductAttributes string     `url:"with_detail_product_attributes,omitempty"`
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
