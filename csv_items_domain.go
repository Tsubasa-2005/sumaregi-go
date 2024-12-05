package sumaregi

import (
	"strings"
)

const (
	TransactionHeadId            = "transactionHeadId"            // 取引ID：数値型。
	TransactionDateTime          = "transactionDateTime"          // 取引日時。[YYYY-MM-DDThh:mm:ssTZD]
	TransactionHeadDivision      = "transactionHeadDivision"      // 取引区分：1:通常, 2:入金, 3:出金, ...
	StoreId                      = "storeId"                      // 店舗ID：店舗毎に付与するID。
	TerminalId                   = "terminalId"                   // 端末ID：店舗の端末毎に付与するID。
	CustomerId                   = "customerId"                   // 会員ID：10桁以内。
	TerminalTranDateTime         = "terminalTranDateTime"         // 端末取引ID：端末で設定された取引ID。
	SumDate                      = "sumDate"                      // 締め日：[YYYY-MM-DD]
	UpdDateTime                  = "updDateTime"                  // 更新日時：[YYYY-MM-DDThh:mm:ssTZD]
	TransactionDetailId          = "transactionDetailId"          // 取引明細ID：数値型。
	ParentTransactionDetailId    = "parentTransactionDetailId"    // 親取引明細ID：紐付く親の取引明細ID。ない場合はNULL。
	TransactionDetailDivision    = "transactionDetailDivision"    // 取引明細区分：1:通常, 2:返品, 3:部門売り。
	ProductId                    = "productId"                    // 商品ID：数字15桁以内。
	ProductCode                  = "productCode"                  // 商品コード：最大64桁。
	ProductName                  = "productName"                  // 商品名：85文字以内。
	PrintReceiptProductName      = "printReceiptProductName"      // レシート印字商品名：64文字以内。
	Color                        = "color"                        // カラー：85文字以内。
	Size                         = "size"                         // サイズ：85文字以内。
	GroupCode                    = "groupCode"                    // グループコード：関連商品として紐付け。
	TaxDivision                  = "taxDivision"                  // 税区分：0:税込, 1:税抜, 2:非課税。
	Price                        = "price"                        // 商品単価：数値型。
	SalesPrice                   = "salesPrice"                   // 販売単価。
	UnitDiscountPrice            = "unitDiscountPrice"            // 単品値引：値引き金額。
	UnitDiscountRate             = "unitDiscountRate"             // 単品割引率：単位：％。
	UnitDiscountDivision         = "unitDiscountDivision"         // 単品値引き/割引区分。
	Cost                         = "cost"                         // 原価：少数5桁まで。
	Quantity                     = "quantity"                     // 数量：数値型。
	UnitNonDiscountSum           = "unitNonDiscountSum"           // 値引き前計：販売価格×数量。
	UnitDiscountSum              = "unitDiscountSum"              // 単品値引き計：値引き×数量。
	UnitDiscountedSum            = "unitDiscountedSum"            // 値引き後計：値引き前計-単品値引き。
	CostSum                      = "costSum"                      // 原価計：少数5桁まで。
	CategoryId                   = "categoryId"                   // 部門ID：数字9桁以内。
	CategoryName                 = "categoryName"                 // 部門名：85文字以内。
	DiscriminationNo             = "discriminationNo"             // 識別番号：20桁以内。
	SalesDivision                = "salesDivision"                // 売上区分：0:売上対象, 1:売上対象外。
	ProductDivision              = "productDivision"              // 商品区分：0:通常, 4:バンドル(親), 7:バンドル商品(子), ...
	InventoryReservationDivision = "inventoryReservationDivision" // 在庫引当区分。
	PointNotApplicable           = "pointNotApplicable"           // ポイント対象：0:対象, 1:対象外。
	CalcDiscount                 = "calcDiscount"                 // 値引割引計算対象。
	TaxFreeDivision              = "taxFreeDivision"              // 免税区分。
	TaxFreeCommodityPrice        = "taxFreeCommodityPrice"        // 免税対象額：数値型。
	TaxFree                      = "taxFree"                      // 免税額：数値型。
	ProductBundleGroupId         = "productBundleGroupId"         // 商品バンドルグループID。
	DiscountPriceProportional    = "discountPriceProportional"    // 小計値引き按分。
	DiscountPointProportional    = "discountPointProportional"    // ポイント値引き按分。
	DiscountCouponProportional   = "discountCouponProportional"   // クーポン値引き按分。
	TaxIncludeProportional       = "taxIncludeProportional"       // 内税按分。
	TaxExcludeProportional       = "taxExcludeProportional"       // 外税按分。
	ProductBundleProportional    = "productBundleProportional"    // 商品バンドル値引按分。
	StaffDiscountProportional    = "staffDiscountProportional"    // 社員値引き按分。
	BargainDiscountProportional  = "bargainDiscountProportional"  // セール値引き按分。
	ProductStaffDiscountRate     = "productStaffDiscountRate"     // 商品毎の社員割引率。
	StaffRank                    = "staffRank"                    // 社員ランクコード。
	StaffRankName                = "staffRankName"                // 社員ランク名。
	StaffDiscountRate            = "staffDiscountRate"            // 社員販売割引率。
	StaffDiscountDivision        = "staffDiscountDivision"        // 社員販売割引区分。
	BargainId                    = "bargainId"                    // セールID。
	BargainName                  = "bargainName"                  // セール名称。
	BargainDivision              = "bargainDivision"              // セール区分。
	BargainValue                 = "bargainValue"                 // セール値。
	ApplyBargainValue            = "applyBargainValue"            // 適用セール値。
	ApplyBargainDiscountPrice    = "applyBargainDiscountPrice"    // 適用セール値引き額。
	TaxRate                      = "taxRate"                      // 適用税率。
	StandardTaxRate              = "standardTaxRate"              // 標準税率。
	ModifiedTaxRate              = "modifiedTaxRate"              // 修正税率。
	ReduceTaxId                  = "reduceTaxId"                  // 軽減税率ID。
	ReduceTaxName                = "reduceTaxName"                // 軽減税率名。
	ReduceTaxRate                = "reduceTaxRate"                // 軽減税率。
	ReduceTaxPrice               = "reduceTaxPrice"               // 軽減税率用商品単価。
	ReduceTaxMemberPrice         = "reduceTaxMemberPrice"         // 軽減税率用会員商品単価。
)

func SelectFields(fields ...string) string {
	return strings.Join(fields, ",")
}
