package sumaregi

type GetStoresIDResponse struct {
	StoreID                    string `json:"storeId"`
	StoreCode                  string `json:"storeCode"`
	StoreName                  string `json:"storeName"`
	StoreAbbr                  string `json:"storeAbbr"`
	PrintReceiptStoreName      string `json:"printReceiptStoreName"`
	PrintStockReceiptStoreName string `json:"printStockReceiptStoreName"`
	Division                   string `json:"division"`
	PostCode                   string `json:"postCode"`
	Address                    string `json:"address"`
	PhoneNumber                string `json:"phoneNumber"`
	FaxNumber                  string `json:"faxNumber"`
	MailAddress                string `json:"mailAddress"`
	Homepage                   string `json:"homepage"`
	TempTranMailAddress        string `json:"tempTranMailAddress"`
	PriceChangeFlag            string `json:"priceChangeFlag"`
	SellDivision               string `json:"sellDivision"`
	SumProcDivision            string `json:"sumProcDivision"`
	SumDateChangeTime          string `json:"sumDateChangeTime"`
	SumRefColumn               string `json:"sumRefColumn"`
	PointNotApplicable         string `json:"pointNotApplicable"`
	TaxFreeDivision            string `json:"taxFreeDivision"`
	MaxBundleProductCount      string `json:"maxBundleProductCount"`
	MaxDiscountRate            string `json:"maxDiscountRate"`
	CarriageDisplayFlag        string `json:"carriageDisplayFlag"`
	TerminalAdjustmentCashFlag string `json:"terminalAdjustmentCashFlag"`
	TerminalCheckCashFlag      string `json:"terminalCheckCashFlag"`
	WaiterAdjustmentDivision   string `json:"waiterAdjustmentDivision"`
	SavingAutoDivision         string `json:"savingAutoDivision"`
	SavingAutoPrice            string `json:"savingAutoPrice"`
	CancelSettingDivision      string `json:"cancelSettingDivision"`
	RoundingDivision           string `json:"roundingDivision"`
	DiscountRoundingDivision   string `json:"discountRoundingDivision"`
	CardCompanySelectDivision  string `json:"cardCompanySelectDivision"`
	GiftReceiptValidDays       string `json:"giftReceiptValidDays"`
	TaxLabelNormal             string `json:"taxLabelNormal"`
	TaxLabelReduce             string `json:"taxLabelReduce"`
	PauseFlag                  string `json:"pauseFlag"`
	DisplaySequence            string `json:"displaySequence"`
	FacePaymentUseDivision     string `json:"facePaymentUseDivision"`
	InvoiceRegistrationNo      string `json:"invoiceRegistrationNo"`
	InsDateTime                string `json:"insDateTime"`
	UpdDateTime                string `json:"updDateTime"`
	PointCondition             struct {
		StoreID                 string `json:"storeId"`
		PointUseDivision        int    `json:"pointUseDivision"`
		SpendRate               string `json:"spendRate"`
		PointGivingUnitPrice    string `json:"pointGivingUnitPrice"`
		PointGivingUnit         string `json:"pointGivingUnit"`
		PointGivingDivision     int    `json:"pointGivingDivision"`
		PointUseUnit            string `json:"pointUseUnit"`
		PointSpendDivision      int    `json:"pointSpendDivision"`
		PointSpendLimitDivision int    `json:"pointSpendLimitDivision"`
		ExpireDivision          int    `json:"expireDivision"`
		ExpireLimit             string `json:"expireLimit"`
		MileageDivision         int    `json:"mileageDivision"`
	} `json:"pointCondition"`
	ReceiptPrintInfo struct {
		StoreID                      string `json:"storeId"`
		Header                       string `json:"header"`
		Footer                       string `json:"footer"`
		ReceiptTaxOfficeStampComment string `json:"receiptTaxOfficeStampComment"`
		TaxOfficeName                string `json:"taxOfficeName"`
		AirPrintLogo                 string `json:"airPrintLogo"`
		AdvertisementImage           string `json:"advertisementImage"`
		GiftReceiptImage             string `json:"giftReceiptImage"`
		GiftReceiptNote              string `json:"giftReceiptNote"`
		DiscountReceiptHeader        string `json:"discountReceiptHeader"`
		DiscountReceiptFooter        string `json:"discountReceiptFooter"`
	} `json:"receiptPrintInfo"`
}

type GetStoresIDOpts struct {
	Fields               []string `url:"fields,omitempty"`
	WithPointCondition   string   `url:"with_point_condition,omitempty"`
	WithReceiptPrintInfo string   `url:"with_receipt_print_info,omitempty"`
}
