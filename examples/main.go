package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tsubasa-2005/sumaregi-go"
)

func main() {
	development := false
	envVari := sumaregi.LoadEnv(development)
	config := sumaregi.NewConfig(envVari.SmaregiAPIHost, envVari.SmaregiContractID)
	scopes := []string{sumaregi.ProductsRead, sumaregi.ProductsWrite}
	client, err := sumaregi.NewClient(config, scopes, envVari)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	//formattedTime := domain.FormatToISO8601(time.Now())
	//created, err := client.PostProducts(ctx, sumaregi.PostProductsOpts{
	//	CategoryID:              "1",
	//	ProductCode:             "PROD-001",
	//	ProductName:             "Sample Product",
	//	ProductKana:             "サンプルプロダクト",
	//	TaxDivision:             "0", // Assuming 1 represents a standard tax division
	//	ProductPriceDivision:    "1", // Assuming 1 means a standard price division
	//	Price:                   "5000",
	//	CustomerPrice:           "5500",
	//	Cost:                    "3000",
	//	Attribute:               "Attribute Example",
	//	Description:             "This is a sample product description.",
	//	CatchCopy:               "Catchy Copy for Sample Product",
	//	Size:                    "Medium",
	//	Color:                   "Red",
	//	Tag:                     "Sample,Product",
	//	GroupCode:               "GRP-001",
	//	URL:                     "https://example.com/product/prod-001",
	//	PrintReceiptProductName: "Sample Product",
	//	DisplaySequence:         "1",
	//	DisplayFlag:             "1", // Assuming 1 means the product is displayed
	//	Division:                "1",
	//	ProductOptionGroupID:    1,
	//	SalesDivision:           "1",
	//	StockControlDivision:    "1",
	//	PointNotApplicable:      "0",
	//	TaxFreeDivision:         "0",
	//	CalcDiscount:            "1",
	//	StaffDiscountRate:       "10", // Assuming this means a 10% staff discount
	//	UseCategoryReduceTax:    "0",
	//	ReduceTaxID:             "1",
	//	ReduceTaxPrice:          "4500",
	//	ReduceTaxCustomerPrice:  "5000",
	//	OrderPoint:              "10",
	//	PurchaseCost:            "2800",
	//	SupplierProductNo:       "SUP-001",
	//	AppStartDateTime:        formattedTime,
	//	ReserveItems:            nil, // Assuming no reserve items
	//	Prices:                  nil, // Assuming no special prices
	//	Stores:                  nil, // Assuming available in all stores
	//	InventoryReservations:   nil, // Assuming no inventory reservations
	//	AttributeItems:          nil, // Assuming no specific attribute items
	//	OrderSetting: struct {
	//		ContinuationDivision  interface{} `json:"continuationDivision"`
	//		OrderStatusDivision   string      `json:"orderStatusDivision"`
	//		OrderNoReasonDivision string      `json:"orderNoReasonDivision"`
	//		OrderUnit             struct {
	//			Division string `json:"division"`
	//			Num      string `json:"num"`
	//			Name     string `json:"name"`
	//		} `json:"orderUnit"`
	//		OrderLimitAmount      string `json:"orderLimitAmount"`
	//		OrderSupplierEditable string `json:"orderSupplierEditable"`
	//		PbDivision            string `json:"pbDivision"`
	//		DisplayFlag           string `json:"displayFlag"`
	//		Stores                []struct {
	//			StoreID          string `json:"storeId"`
	//			OrderLimitAmount string `json:"orderLimitAmount"`
	//			DisplayFlag      string `json:"displayFlag"`
	//		} `json:"stores"`
	//	}{
	//		ContinuationDivision:  nil, // Assuming no continuation
	//		OrderStatusDivision:   "1", // Assuming 1 represents an active status
	//		OrderNoReasonDivision: "ORD-REASON-001",
	//		OrderUnit: struct {
	//			Division string `json:"division"`
	//			Num      string `json:"num"`
	//			Name     string `json:"name"`
	//		}{
	//			Division: "1", // Assuming 1 is a standard division
	//			Num:      "1",
	//			Name:     "Box",
	//		},
	//		OrderLimitAmount:      "100",
	//		OrderSupplierEditable: "1", // Assuming 1 means editable
	//		PbDivision:            "1",
	//		DisplayFlag:           "1", // Assuming 1 means display
	//		Stores: []struct {
	//			StoreID          string `json:"storeId"`
	//			OrderLimitAmount string `json:"orderLimitAmount"`
	//			DisplayFlag      string `json:"displayFlag"`
	//		}{
	//			{
	//				StoreID:          "1",
	//				OrderLimitAmount: "50",
	//				DisplayFlag:      "1",
	//			},
	//		},
	//	},
	//})
	//if err != nil {
	//	return
	//}
	//
	//fmt.Print(created)

	products, err := client.GetProducts(ctx, sumaregi.GetProductsOpts{})
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}

	fmt.Print(products)
}
