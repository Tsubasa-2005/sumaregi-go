package sumaregi

// GetTransactionsOpts represents options for GetTransactions.
type GetTransactionsOpts struct {
	Fields                      []string `url:"fields,omitempty"`                         // 検索パラメータ (カンマ区切り指定可)。レスポンス項目を指定可能。一部ネスト項目は指定不可。
	Sort                        string   `url:"sort,omitempty"`                           // 並び順 (カンマ区切り指定可)。
	Limit                       int      `url:"limit,omitempty"`                          // 最大取得件数。
	Page                        int      `url:"page,omitempty"`                           // ページ番号 (ページネーション用)。
	TransactionHeadIDFrom       string   `url:"transaction_head_id-from,omitempty"`       // 取引ID (From)。
	TransactionHeadIDTo         string   `url:"transaction_head_id-to,omitempty"`         // 取引ID (To)。
	TransactionDateTimeFrom     string   `url:"transaction_date_time-from,omitempty"`     // 取引日時 (From)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。最大31日。
	TransactionDateTimeTo       string   `url:"transaction_date_time-to,omitempty"`       // 取引日時 (To)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。最大31日。
	TransactionHeadDivision     string   `url:"transaction_head_division,omitempty"`      // 取引区分 (Enum: 1=通常, 2=入金, 3=出金, 4=預かり金, ... 16=領収証)。
	StoreID                     int      `url:"store_id,omitempty"`                       // 店舗ID。
	TerminalTranDateTimeFrom    string   `url:"terminal_tran_date_time-from,omitempty"`   // 端末取引日時 (From)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。
	TerminalTranDateTimeTo      string   `url:"terminal_tran_date_time-to,omitempty"`     // 端末取引日時 (To)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。
	AdjustmentDateTime          string   `url:"adjustment_date_time,omitempty"`           // 精算日時。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。
	SumDate                     string   `url:"sum_date,omitempty"`                       // 締め日。フォーマット: [YYYY-MM-DD]。
	SumDateFrom                 string   `url:"sum_date-from,omitempty"`                  // 締め日 (From)。フォーマット: [YYYY-MM-DD]。
	SumDateTo                   string   `url:"sum_date-to,omitempty"`                    // 締め日 (To)。フォーマット: [YYYY-MM-DD]。
	CustomerCode                string   `url:"customer_code,omitempty"`                  // 会員コード。
	TransactionUUID             string   `url:"transaction_uuid,omitempty"`               // レシート番号 (完全一致)。
	Barcode                     string   `url:"barcode,omitempty"`                        // バーコードでフィルタリング。
	UpdDateTimeFrom             string   `url:"upd_date_time-from,omitempty"`             // 更新日時 (From)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。最大31日。
	UpdDateTimeTo               string   `url:"upd_date_time-to,omitempty"`               // 更新日時 (To)。フォーマット: [YYYY-MM-DDThh:mm:ssTZD]。最大31日。
	WithDetails                 string   `url:"with_details,omitempty"`                   // 取引明細情報の付加設定 (Enum: "all"=全項目, "summary"=一部項目, "none"=付加しない)。
	WithDepositOthers           string   `url:"with_deposit_others,omitempty"`            // その他支払い情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithLayaway                 string   `url:"with_layaway,omitempty"`                   // (非推奨) 取置き情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithLayaways                string   `url:"with_layaways,omitempty"`                  // 取置き情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithLayawayPickUp           string   `url:"with_layaway_pick_up,omitempty"`           // (非推奨) 取置き引取情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithLayawayPickUps          string   `url:"with_layaway_pick_ups,omitempty"`          // 取置き引取情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithMoneyControl            string   `url:"with_money_control,omitempty"`             // 取引金操作情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithDetailProductAttributes string   `url:"with_detail_product_attributes,omitempty"` // 販売時の商品属性情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
}

// PostTransactionCSVOpts represents the request structure for fetching transaction details.
type PostTransactionCSVOpts struct {
	Fields                   string `json:"fields,omitempty"`                   // 検索パラメータ (カンマ区切りで指定可)
	Sort                     string `json:"sort,omitempty"`                     // 並び順 (カンマ区切りで指定可)
	TransactionHeadIDFrom    string `json:"transactionHeadIdFrom,omitempty"`    // 取引ID(From)
	TransactionHeadIDTo      string `json:"transactionHeadIdTo,omitempty"`      // 取引ID(To)
	TransactionDateTimeFrom  string `json:"transactionDateTimeFrom"`            // 取引日時(From) (必須) [YYYY-MM-DDThh:mm:ssTZD]
	TransactionDateTimeTo    string `json:"transactionDateTimeTo"`              // 取引日時(To) (必須) [YYYY-MM-DDThh:mm:ssTZD]
	TransactionHeadDivision  int    `json:"transactionHeadDivision,omitempty"`  // 取引区分 (1:通常,10:取置き)
	StoreID                  string `json:"storeId,omitempty"`                  // 店舗ID
	TerminalTranDateTimeFrom string `json:"terminalTranDateTimeFrom,omitempty"` // 端末取引日時(From) [YYYY-MM-DDThh:mm:ssTZD]
	TerminalTranDateTimeTo   string `json:"terminalTranDateTimeTo,omitempty"`   // 端末取引日時(To) [YYYY-MM-DDThh:mm:ssTZD]
	SumDate                  string `json:"sumDate,omitempty"`                  // 締め日 [YYYY-MM-DD]
	SumDateFrom              string `json:"sumDateFrom,omitempty"`              // 締め日(From) [YYYY-MM-DD]
	SumDateTo                string `json:"sumDateTo,omitempty"`                // 締め日(To) [YYYY-MM-DD]
	CustomerSaleDivision     int    `json:"customerSaleDivision,omitempty"`     // 会員販売区分 (0:非会員販売, 1:会員販売)
	UpdDateTimeFrom          string `json:"updDateTimeFrom,omitempty"`          // 更新日時(From) [YYYY-MM-DDThh:mm:ssTZD]
	UpdDateTimeTo            string `json:"updDateTimeTo,omitempty"`            // 更新日時(To) [YYYY-MM-DDThh:mm:ssTZD]
	CallbackURL              string `json:"callbackUrl"`                        // 取引明細CSV作成完了通知URL (必須)
	State                    string `json:"state,omitempty"`                    // 完了通知ステートメント (20文字以内の任意文字列)
}

// GetTransactionsResponse represents a response from GetTransactions.
type GetTransactionsResponse []struct {
	TransactionHeadID              string `json:"transactionHeadId"`
	TransactionDateTime            string `json:"transactionDateTime"`
	TransactionHeadDivision        string `json:"transactionHeadDivision"`
	CancelDivision                 string `json:"cancelDivision"`
	UnitNonDiscountsubtotal        string `json:"unitNonDiscountsubtotal"`
	UnitDiscountsubtotal           string `json:"unitDiscountsubtotal"`
	UnitStaffDiscountsubtotal      string `json:"unitStaffDiscountsubtotal"`
	UnitBargainDiscountsubtotal    string `json:"unitBargainDiscountsubtotal"`
	Subtotal                       string `json:"subtotal"`
	SubtotalForDiscount            string `json:"subtotalForDiscount"`
	SubtotalDiscountPrice          string `json:"subtotalDiscountPrice"`
	SubtotalDiscountRate           string `json:"subtotalDiscountRate"`
	SubtotalDiscountDivision       string `json:"subtotalDiscountDivision"`
	PointDiscount                  string `json:"pointDiscount"`
	CouponDiscount                 string `json:"couponDiscount"`
	Total                          string `json:"total"`
	TaxInclude                     string `json:"taxInclude"`
	TaxExclude                     string `json:"taxExclude"`
	RoundingDivision               string `json:"roundingDivision"`
	RoundingPrice                  string `json:"roundingPrice"`
	CashTotal                      string `json:"cashTotal"`
	CreditTotal                    string `json:"creditTotal"`
	Deposit                        string `json:"deposit"`
	DepositCash                    string `json:"depositCash"`
	DepositCredit                  string `json:"depositCredit"`
	Change                         string `json:"change"`
	TipCash                        string `json:"tipCash"`
	TipCredit                      string `json:"tipCredit"`
	Amount                         string `json:"amount"`
	ReturnAmount                   string `json:"returnAmount"`
	CostTotal                      string `json:"costTotal"`
	SalesHeadDivision              string `json:"salesHeadDivision"`
	InTaxSalesTotal                string `json:"inTaxSalesTotal"`
	OutTaxSalesTotal               string `json:"outTaxSalesTotal"`
	NonTaxSalesTotal               string `json:"nonTaxSalesTotal"`
	NonSalesTargetTotal            string `json:"nonSalesTargetTotal"`
	NonSalesTargetOutTaxTotal      string `json:"nonSalesTargetOutTaxTotal"`
	NonSalesTargetInTaxTotal       string `json:"nonSalesTargetInTaxTotal"`
	NonSalesTargetTaxFreeTotal     string `json:"nonSalesTargetTaxFreeTotal"`
	NonSalesTargetCostTotal        string `json:"nonSalesTargetCostTotal"`
	NonSalesTargetAmount           string `json:"nonSalesTargetAmount"`
	NonSalesTargetReturnAmount     string `json:"nonSalesTargetReturnAmount"`
	NewPoint                       string `json:"newPoint"`
	SpendPoint                     string `json:"spendPoint"`
	Point                          string `json:"point"`
	TotalPoint                     string `json:"totalPoint"`
	CurrentMile                    string `json:"currentMile"`
	EarnMile                       string `json:"earnMile"`
	TotalMile                      string `json:"totalMile"`
	AdjustmentMile                 string `json:"adjustmentMile"`
	AdjustmentMileDivision         string `json:"adjustmentMileDivision"`
	AdjustmentMileValue            string `json:"adjustmentMileValue"`
	StoreID                        string `json:"storeId"`
	StoreCode                      string `json:"storeCode"`
	TerminalID                     string `json:"terminalId"`
	CustomerID                     string `json:"customerId"`
	CustomerCode                   string `json:"customerCode"`
	TerminalTranID                 string `json:"terminalTranId"`
	TerminalTranDateTime           string `json:"terminalTranDateTime"`
	SumDivision                    string `json:"sumDivision"`
	AdjustmentDateTime             string `json:"adjustmentDateTime"`
	SumDate                        string `json:"sumDate"`
	CustomerRank                   string `json:"customerRank"`
	CustomerGroupID                string `json:"customerGroupId"`
	CustomerGroupID2               string `json:"customerGroupId2"`
	CustomerGroupID3               string `json:"customerGroupId3"`
	CustomerGroupID4               string `json:"customerGroupId4"`
	CustomerGroupID5               string `json:"customerGroupId5"`
	StaffID                        string `json:"staffId"`
	StaffCode                      string `json:"staffCode"`
	StaffName                      string `json:"staffName"`
	CreditDivision                 string `json:"creditDivision"`
	PaymentCount                   string `json:"paymentCount"`
	SlipNumber                     string `json:"slipNumber"`
	CancelSlipNumber               string `json:"cancelSlipNumber"`
	AuthNumber                     string `json:"authNumber"`
	AuthDate                       string `json:"authDate"`
	CardCompany                    string `json:"cardCompany"`
	Denomination                   string `json:"denomination"`
	Memo                           string `json:"memo"`
	ReceiptMemo                    string `json:"receiptMemo"`
	Carriage                       string `json:"carriage"`
	Commission                     string `json:"commission"`
	GuestNumbers                   string `json:"guestNumbers"`
	GuestNumbersMale               string `json:"guestNumbersMale"`
	GuestNumbersFemale             string `json:"guestNumbersFemale"`
	GuestNumbersUnknown            string `json:"guestNumbersUnknown"`
	EnterDateTime                  string `json:"enterDateTime"`
	TaxFreeSalesDivision           string `json:"taxFreeSalesDivision"`
	NetTaxFreeGeneralTaxInclude    string `json:"netTaxFreeGeneralTaxInclude"`
	NetTaxFreeGeneralTaxExclude    string `json:"netTaxFreeGeneralTaxExclude"`
	NetTaxFreeConsumableTaxInclude string `json:"netTaxFreeConsumableTaxInclude"`
	NetTaxFreeConsumableTaxExclude string `json:"netTaxFreeConsumableTaxExclude"`
	Tags                           string `json:"tags"`
	PointGivingDivision            string `json:"pointGivingDivision"`
	PointGivingUnitPrice           string `json:"pointGivingUnitPrice"`
	PointGivingUnit                string `json:"pointGivingUnit"`
	PointSpendDivision             string `json:"pointSpendDivision"`
	MileageDivision                string `json:"mileageDivision"`
	MileageLabel                   string `json:"mileageLabel"`
	CustomerPinCode                string `json:"customerPinCode"`
	ReturnSales                    string `json:"returnSales"`
	DisposeDivision                string `json:"disposeDivision"`
	DisposeServerTransactionHeadID string `json:"disposeServerTransactionHeadId"`
	CancelDateTime                 string `json:"cancelDateTime"`
	SellDivision                   string `json:"sellDivision"`
	TaxRate                        string `json:"taxRate"`
	TaxRounding                    string `json:"taxRounding"`
	DiscountRoundingDivision       string `json:"discountRoundingDivision"`
	TransactionUUID                string `json:"transactionUuid"`
	ExchangeTicketNo               string `json:"exchangeTicketNo"`
	GiftReceiptValidDays           string `json:"giftReceiptValidDays"`
	Barcode                        string `json:"barcode"`
	UpdDateTime                    string `json:"updDateTime"`
	Details                        []struct {
		TransactionHeadID            string `json:"transactionHeadId"`
		TransactionDetailID          string `json:"transactionDetailId"`
		ParentTransactionDetailID    string `json:"parentTransactionDetailId"`
		TransactionDetailDivision    string `json:"transactionDetailDivision"`
		ProductID                    string `json:"productId"`
		ProductCode                  string `json:"productCode"`
		ProductName                  string `json:"productName"`
		PrintReceiptProductName      string `json:"printReceiptProductName"`
		Color                        string `json:"color"`
		Size                         string `json:"size"`
		GroupCode                    string `json:"groupCode"`
		TaxDivision                  string `json:"taxDivision"`
		Price                        string `json:"price"`
		SalesPrice                   string `json:"salesPrice"`
		UnitDiscountPrice            string `json:"unitDiscountPrice"`
		UnitDiscountRate             string `json:"unitDiscountRate"`
		UnitDiscountDivision         string `json:"unitDiscountDivision"`
		Cost                         string `json:"cost"`
		Quantity                     string `json:"quantity"`
		UnitNonDiscountSum           string `json:"unitNonDiscountSum"`
		UnitDiscountSum              string `json:"unitDiscountSum"`
		UnitDiscountedSum            string `json:"unitDiscountedSum"`
		CostSum                      string `json:"costSum"`
		CategoryID                   string `json:"categoryId"`
		CategoryName                 string `json:"categoryName"`
		DiscriminationNo             string `json:"discriminationNo"`
		SalesDivision                string `json:"salesDivision"`
		ProductDivision              string `json:"productDivision"`
		InventoryReservationDivision string `json:"inventoryReservationDivision"`
		PointNotApplicable           string `json:"pointNotApplicable"`
		CalcDiscount                 string `json:"calcDiscount"`
		TaxFreeDivision              string `json:"taxFreeDivision"`
		TaxFreeCommodityPrice        string `json:"taxFreeCommodityPrice"`
		TaxFree                      string `json:"taxFree"`
		ProductBundleGroupID         string `json:"productBundleGroupId"`
		DiscountPriceProportional    string `json:"discountPriceProportional"`
		DiscountPointProportional    string `json:"discountPointProportional"`
		DiscountCouponProportional   string `json:"discountCouponProportional"`
		TaxIncludeProportional       string `json:"taxIncludeProportional"`
		TaxExcludeProportional       string `json:"taxExcludeProportional"`
		ProductBundleProportional    string `json:"productBundleProportional"`
		StaffDiscountProportional    string `json:"staffDiscountProportional"`
		BargainDiscountProportional  string `json:"bargainDiscountProportional"`
		RoundingPriceProportional    string `json:"roundingPriceProportional"`
		ProductStaffDiscountRate     string `json:"productStaffDiscountRate"`
		StaffRank                    string `json:"staffRank"`
		StaffRankName                string `json:"staffRankName"`
		StaffDiscountRate            string `json:"staffDiscountRate"`
		StaffDiscountDivision        string `json:"staffDiscountDivision"`
		ApplyStaffDiscountRate       string `json:"applyStaffDiscountRate"`
		ApplyStaffDiscountPrice      string `json:"applyStaffDiscountPrice"`
		BargainID                    string `json:"bargainId"`
		BargainName                  string `json:"bargainName"`
		BargainDivision              string `json:"bargainDivision"`
		BargainValue                 string `json:"bargainValue"`
		ApplyBargainValue            string `json:"applyBargainValue"`
		ApplyBargainDiscountPrice    string `json:"applyBargainDiscountPrice"`
		TaxRate                      string `json:"taxRate"`
		StandardTaxRate              string `json:"standardTaxRate"`
		ModifiedTaxRate              string `json:"modifiedTaxRate"`
		ReduceTaxID                  string `json:"reduceTaxId"`
		ReduceTaxName                string `json:"reduceTaxName"`
		ReduceTaxRate                string `json:"reduceTaxRate"`
		ReduceTaxPrice               string `json:"reduceTaxPrice"`
		ReduceTaxMemberPrice         string `json:"reduceTaxMemberPrice"`
		Memo                         string `json:"memo"`
		ProductAttributes            []struct {
			TransactionHeadID   string `json:"transactionHeadId"`
			TransactionDetailID string `json:"transactionDetailId"`
			Code                string `json:"code"`
			Name                string `json:"name"`
		} `json:"productAttributes"`
		RfidTags []string `json:"rfidTags"`
	} `json:"details"`
	DepositOthers []struct {
		TransactionHeadID     string `json:"transactionHeadId"`
		No                    string `json:"no"`
		PaymentMethodID       string `json:"paymentMethodId"`
		PaymentMethodCode     string `json:"paymentMethodCode"`
		PaymentMethodName     string `json:"paymentMethodName"`
		DepositOthers         string `json:"depositOthers"`
		PaymentUnitPrice      string `json:"paymentUnitPrice"`
		PaymentChangeFlag     int    `json:"paymentChangeFlag"`
		PaymentDivision       string `json:"paymentDivision"`
		PaymentSecuritiesFlag int    `json:"paymentSecuritiesFlag"`
		DenominationCode      string `json:"denominationCode"`
		DenominationName      string `json:"denominationName"`
		CardCompanyName       string `json:"cardCompanyName"`
		SlipNumber            string `json:"slipNumber"`
		CancelSlipNumber      string `json:"cancelSlipNumber"`
		PointGivingUnitPrice  string `json:"pointGivingUnitPrice"`
		PointGivingUnit       string `json:"pointGivingUnit"`
	} `json:"depositOthers"`
	Layaway struct {
		TransactionHeadID       string `json:"transactionHeadId"`
		PickUpDate              string `json:"pickUpDate"`
		Status                  string `json:"status"`
		PartPayment             string `json:"partPayment"`
		PartPaymentClass        string `json:"partPaymentClass"`
		PickUpTransactionHeadID string `json:"pickUpTransactionHeadId"`
		DisabledEdit            string `json:"disabledEdit"`
	} `json:"layaway"`
	Layaways []struct {
		TransactionHeadID       string `json:"transactionHeadId"`
		PickUpDate              string `json:"pickUpDate"`
		Status                  string `json:"status"`
		PartPayment             string `json:"partPayment"`
		PartPaymentClass        string `json:"partPaymentClass"`
		PickUpTransactionHeadID string `json:"pickUpTransactionHeadId"`
		DisabledEdit            string `json:"disabledEdit"`
	} `json:"layaways"`
	LayawayPickUp struct {
		TransactionHeadID              string `json:"transactionHeadId"`
		LayawayServerTransactionHeadID string `json:"layawayServerTransactionHeadId"`
		ReceivedDepositCash            string `json:"receivedDepositCash"`
		ReceivedDepositCredit          string `json:"receivedDepositCredit"`
	} `json:"layawayPickUp"`
	LayawayPickUps []struct {
		TransactionHeadID              string `json:"transactionHeadId"`
		LayawayServerTransactionHeadID string `json:"layawayServerTransactionHeadId"`
		ReceivedDepositCash            string `json:"receivedDepositCash"`
		ReceivedDepositCredit          string `json:"receivedDepositCredit"`
	} `json:"layawayPickUps"`
	MoneyControl struct {
		TransactionHeadID     string `json:"transactionHeadId"`
		AttributeDivision     string `json:"attributeDivision"`
		AttributeDivisionName string `json:"attributeDivisionName"`
	} `json:"moneyControl"`
}

// GetTransactionDetailOpts represents options for GetTransactionDetail.
type GetTransactionDetailOpts struct {
	Fields                      []string `url:"fields,omitempty"`                         // 検索パラメータ (カンマ区切り指定可)。レスポンス項目を指定可能。一部ネスト項目は指定不可。
	Sort                        string   `url:"sort,omitempty"`                           // 並び順 (カンマ区切り指定可)。
	Limit                       int      `url:"limit,omitempty"`                          // 最大取得件数。
	Page                        int      `url:"page,omitempty"`                           // ページ番号 (ページネーション用)。
	WithDiscounts               string   `url:"with_discounts,omitempty"`                 // 単品値引／割引情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
	WithDetailProductAttributes string   `url:"with_detail_product_attributes,omitempty"` // 商品属性情報の付加設定 (Enum: "all"=付加する, "none"=付加しない)。
}

// GetTransactionDetailResponse represents a response from GetTransactionDetail.
type GetTransactionDetailResponse []struct {
	TransactionHeadID            string `json:"transactionHeadId"`
	TransactionDetailID          string `json:"transactionDetailId"`
	ParentTransactionDetailID    string `json:"parentTransactionDetailId"`
	TransactionDetailDivision    string `json:"transactionDetailDivision"`
	ProductID                    string `json:"productId"`
	ProductCode                  string `json:"productCode"`
	ProductName                  string `json:"productName"`
	PrintReceiptProductName      string `json:"printReceiptProductName"`
	Color                        string `json:"color"`
	Size                         string `json:"size"`
	GroupCode                    string `json:"groupCode"`
	TaxDivision                  string `json:"taxDivision"`
	Price                        string `json:"price"`
	SalesPrice                   string `json:"salesPrice"`
	UnitDiscountPrice            string `json:"unitDiscountPrice"`
	UnitDiscountRate             string `json:"unitDiscountRate"`
	UnitDiscountDivision         string `json:"unitDiscountDivision"`
	Cost                         string `json:"cost"`
	Quantity                     string `json:"quantity"`
	UnitNonDiscountSum           string `json:"unitNonDiscountSum"`
	UnitDiscountSum              string `json:"unitDiscountSum"`
	UnitDiscountedSum            string `json:"unitDiscountedSum"`
	CostSum                      string `json:"costSum"`
	CategoryID                   string `json:"categoryId"`
	CategoryName                 string `json:"categoryName"`
	DiscriminationNo             string `json:"discriminationNo"`
	SalesDivision                string `json:"salesDivision"`
	ProductDivision              string `json:"productDivision"`
	InventoryReservationDivision string `json:"inventoryReservationDivision"`
	PointNotApplicable           string `json:"pointNotApplicable"`
	CalcDiscount                 string `json:"calcDiscount"`
	TaxFreeDivision              string `json:"taxFreeDivision"`
	TaxFreeCommodityPrice        string `json:"taxFreeCommodityPrice"`
	TaxFree                      string `json:"taxFree"`
	ProductBundleGroupID         string `json:"productBundleGroupId"`
	DiscountPriceProportional    string `json:"discountPriceProportional"`
	DiscountPointProportional    string `json:"discountPointProportional"`
	DiscountCouponProportional   string `json:"discountCouponProportional"`
	TaxIncludeProportional       string `json:"taxIncludeProportional"`
	TaxExcludeProportional       string `json:"taxExcludeProportional"`
	ProductBundleProportional    string `json:"productBundleProportional"`
	StaffDiscountProportional    string `json:"staffDiscountProportional"`
	BargainDiscountProportional  string `json:"bargainDiscountProportional"`
	RoundingPriceProportional    string `json:"roundingPriceProportional"`
	ProductStaffDiscountRate     string `json:"productStaffDiscountRate"`
	StaffRank                    string `json:"staffRank"`
	StaffRankName                string `json:"staffRankName"`
	StaffDiscountRate            string `json:"staffDiscountRate"`
	StaffDiscountDivision        string `json:"staffDiscountDivision"`
	ApplyStaffDiscountRate       string `json:"applyStaffDiscountRate"`
	ApplyStaffDiscountPrice      string `json:"applyStaffDiscountPrice"`
	BargainID                    string `json:"bargainId"`
	BargainName                  string `json:"bargainName"`
	BargainDivision              string `json:"bargainDivision"`
	BargainValue                 string `json:"bargainValue"`
	ApplyBargainValue            string `json:"applyBargainValue"`
	ApplyBargainDiscountPrice    string `json:"applyBargainDiscountPrice"`
	TaxRate                      string `json:"taxRate"`
	StandardTaxRate              string `json:"standardTaxRate"`
	ModifiedTaxRate              string `json:"modifiedTaxRate"`
	ReduceTaxID                  string `json:"reduceTaxId"`
	ReduceTaxName                string `json:"reduceTaxName"`
	ReduceTaxRate                string `json:"reduceTaxRate"`
	ReduceTaxPrice               string `json:"reduceTaxPrice"`
	ReduceTaxMemberPrice         string `json:"reduceTaxMemberPrice"`
	Memo                         string `json:"memo"`
	Discounts                    []struct {
		TransactionHeadID           string `json:"transactionHeadId"`
		TransactionDetailID         string `json:"transactionDetailId"`
		TransactionDiscountDivision string `json:"transactionDiscountDivision"`
		DiscountDivision            string `json:"discountDivision"`
		DiscountDivisionName        string `json:"discountDivisionName"`
		AwardType                   string `json:"awardType"`
		AwardValue                  string `json:"awardValue"`
		DiscountPrice               string `json:"discountPrice"`
	} `json:"discounts"`
	ProductAttributes []struct {
		TransactionHeadID   string `json:"transactionHeadId"`
		TransactionDetailID string `json:"transactionDetailId"`
		Code                string `json:"code"`
		Name                string `json:"name"`
	} `json:"productAttributes"`
	RfidTags []string `json:"rfidTags"`
}

type PostTransactionCSVResponse struct {
	CallbackURL string `json:"callbackUrl"`
	State       string `json:"state"`
	RequestCode string `json:"requestCode"`
}
