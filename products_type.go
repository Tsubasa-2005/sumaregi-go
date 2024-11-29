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

// GetProductsOpts は、商品データを取得する際のクエリパラメータを表します。
type GetProductsOpts struct {
	Fields               []string `url:"fields,omitempty"`                 // レスポンスに含めるフィールドを指定 (カンマ区切りで複数指定可。一部のネスト項目は指定不可)。
	Sort                 string   `url:"sort,omitempty"`                   // レスポンスの並び順を指定 (カンマ区切りで複数条件指定可)。
	Limit                int      `url:"limit,omitempty"`                  // 取得する商品の最大件数。
	Page                 int      `url:"page,omitempty"`                   // ページ番号を指定 (ページネーション用)。
	CategoryID           int      `url:"category_id,omitempty"`            // 部門IDで商品をフィルタリング。
	ProductCode          string   `url:"product_code,omitempty"`           // 商品コードで商品をフィルタリング。
	GroupCode            string   `url:"group_code,omitempty"`             // グループコードで商品をフィルタリング。
	DisplayFlag          string   `url:"display_flag,omitempty"`           // 端末表示フラグ (Enum: 0=表示しない, 1=表示する)。
	Division             string   `url:"division,omitempty"`               // 商品区分 (Enum: 0=通常商品, 1=回数券, 2=オプション商品)。
	SalesDivision        string   `url:"sales_division,omitempty"`         // 売上区分 (Enum: 0=売上対象, 1=売上対象外)。
	StockControlDivision string   `url:"stock_control_division,omitempty"` // 在庫管理区分 (Enum: 0=在庫管理対象, 1=在庫管理対象外)。
	SupplierProductNo    string   `url:"supplier_product_no,omitempty"`    // メーカー品番で商品をフィルタリング。
	UpdDateTimeFrom      string   `url:"upd_date_time-from,omitempty"`     // 更新日時の開始を指定 (フォーマット: [YYYY-MM-DDThh:mm:ssTZD], 最大31日)。
	UpdDateTimeTo        string   `url:"upd_date_time-to,omitempty"`       // 更新日時の終了を指定 (フォーマット: [YYYY-MM-DDThh:mm:ssTZD], 最大31日)。
}

// PostProductsParams 、APIで商品を作成または更新するためのパラメータを表します。
type PostProductsParams struct {
	CategoryID              string     `json:"categoryId"`              // 部門ID (必須) - 部門ごとに付与されるID
	ProductCode             string     `json:"productCode"`             // 商品コード (必須) - ユニーク値, 最大20文字, 正規表現: ^[ -~]+$
	ProductName             string     `json:"productName"`             // 商品名 (必須) - 最大85文字
	ProductKana             string     `json:"productKana"`             // 商品名カナ - 最大85文字
	TaxDivision             string     `json:"taxDivision"`             // 税区分 (デフォルト: 1) - Enum: 0: 税込, 1: 税抜, 2: 非課税
	ProductPriceDivision    string     `json:"productPriceDivision"`    // 商品価格区分 (デフォルト: 1) - Enum: 1: 通常価格, 2: オープン価格
	Price                   string     `json:"price"`                   // 商品単価 (必須) - 整数, 範囲: -99999999 ~ 99999999
	CustomerPrice           string     `json:"customerPrice"`           // 会員価格 - 整数, 範囲: -99999999 ~ 99999999
	Cost                    string     `json:"cost"`                    // 原価 - 小数点付き, 範囲: -99999999.99999 ~ 99999999.99999
	Attribute               string     `json:"attribute"`               // 商品属性 - 最大1000文字
	Description             string     `json:"description"`             // 商品説明 - 最大1000文字
	CatchCopy               string     `json:"catchCopy"`               // キャッチコピー - 最大1000文字
	Size                    string     `json:"size"`                    // サイズ - 最大85文字
	Color                   string     `json:"color"`                   // カラー - 最大85文字
	Tag                     string     `json:"tag"`                     // タグ - 最大85文字
	GroupCode               string     `json:"groupCode"`               // グループコード - 最大85文字 (親の商品ID)
	URL                     string     `json:"url"`                     // URL - 最大255文字
	PrintReceiptProductName string     `json:"printReceiptProductName"` // レシート印字商品名 - 最大64文字
	DisplaySequence         string     `json:"displaySequence"`         // 表示順 - 整数, 範囲: 0 ~ 999999999
	DisplayFlag             string     `json:"displayFlag"`             // 端末表示 (デフォルト: 1) - Enum: 0: 表示しない, 1: 表示する
	Division                string     `json:"division"`                // 商品区分 (デフォルト: 0) - Enum: 0: 通常, 1: 回数券, 2: オプション
	ProductOptionGroupID    int        `json:"productOptionGroupId"`    // オプショングループID - 範囲: 1 ~ 999999999
	SalesDivision           string     `json:"salesDivision"`           // 売上区分 (デフォルト: 0) - Enum: 0: 売上対象, 1: 売上対象外
	StockControlDivision    string     `json:"stockControlDivision"`    // 在庫管理区分 (デフォルト: 0) - Enum: 0: 在庫管理対象, 1: 在庫管理対象外
	PointNotApplicable      string     `json:"pointNotApplicable"`      // ポイント対象区分 (デフォルト: 0) - Enum: 0: 対象, 1: 対象外
	TaxFreeDivision         string     `json:"taxFreeDivision"`         // 免税区分 (デフォルト: 0) - Enum: 0: 対象外, 1: 一般品, 2: 消耗品
	CalcDiscount            string     `json:"calcDiscount"`            // 値引割引計算対象 (デフォルト: 1) - Enum: 0: 対象外, 1: 対象
	StaffDiscountRate       string     `json:"staffDiscountRate"`       // 社員販売割引率 - 整数, 範囲: 0 ~ 100
	UseCategoryReduceTax    string     `json:"useCategoryReduceTax"`    // 部門税設定参照フラグ (デフォルト: 1) - Enum: 0: reduceTaxIdを使用, 1: 部門の税設定を使用
	ReduceTaxID             string     `json:"reduceTaxId"`             // 軽減税率ID - 範囲: 0 ~ 999999999
	ReduceTaxPrice          string     `json:"reduceTaxPrice"`          // 軽減税率用商品単価 - 整数, 範囲: -99999999 ~ 99999999
	ReduceTaxCustomerPrice  string     `json:"reduceTaxCustomerPrice"`  // 軽減税率用会員単価 - 整数, 範囲: -99999999 ~ 99999999
	OrderPoint              string     `json:"orderPoint"`              // 発注点 - 整数, 範囲: 0 ~ 999999999
	PurchaseCost            string     `json:"purchaseCost"`            // 仕入原価 - 小数点付き, 範囲: -99999999.99999 ~ 99999999.99999
	SupplierProductNo       string     `json:"supplierProductNo"`       // メーカー品番 - 最大85文字
	AppStartDateTime        string     `json:"appStartDateTime"`        // 適用開始日時 - ISO8601形式 (YYYY-MM-DDThh:mm:ssTZD)
	ReserveItems            []struct { // 商品自由項目
		No    string `json:"no"`    // 自由項目番号
		Value string `json:"value"` // 自由項目値
	} `json:"reserveItems"`
	Prices []struct { // 商品価格
		StoreID       string `json:"storeId"`       // 店舗ID
		PriceDivision int    `json:"priceDivision"` // 価格区分
		StartDate     string `json:"startDate"`     // 開始日 - ISO8601形式
		EndDate       string `json:"endDate"`       // 終了日 - ISO8601形式
		Price         string `json:"price"`         // 価格
	} `json:"prices"`
	Stores []struct { // 商品取扱店舗
		StoreID              string `json:"storeId"`              // 店舗ID
		ProductOptionGroupID string `json:"productOptionGroupId"` // 商品オプショングループID
		AssignDivision       string `json:"assignDivision"`       // 割り当て区分
	} `json:"stores"`
	InventoryReservations []struct { // 在庫引当商品
		ReservationProductID string `json:"reservationProductId"` // 引当商品ID
		ReservationAmount    string `json:"reservationAmount"`    // 引当数量
	} `json:"inventoryReservations"`
	AttributeItems []struct { // 商品属性項目
		Code string `json:"code"` // 属性コード
		No   string `json:"no"`   // 属性番号
	} `json:"attributeItems"`
	OrderSetting struct { // 発注設定
		ContinuationDivision  interface{} `json:"continuationDivision"`  // 継続区分
		OrderStatusDivision   string      `json:"orderStatusDivision"`   // 発注状態区分
		OrderNoReasonDivision string      `json:"orderNoReasonDivision"` // 発注理由区分
		OrderUnit             struct {    // 発注単位
			Division string `json:"division"` // 単位区分
			Num      string `json:"num"`      // 単位数
			Name     string `json:"name"`     // 単位名
		} `json:"orderUnit"`
		OrderLimitAmount      string     `json:"orderLimitAmount"`      // 発注上限額
		OrderSupplierEditable string     `json:"orderSupplierEditable"` // 発注先編集可能
		PbDivision            string     `json:"pbDivision"`            // PB区分
		DisplayFlag           string     `json:"displayFlag"`           // 表示フラグ
		Stores                []struct { // 発注設定店舗情報
			StoreID          string `json:"storeId"`          // 店舗ID
			OrderLimitAmount string `json:"orderLimitAmount"` // 発注上限額
			DisplayFlag      string `json:"displayFlag"`      // 表示フラグ
		} `json:"stores"`
	} `json:"orderSetting"`
}
