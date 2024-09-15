package utils

import (
	"fmt"
	"log"
	"reflect"

	"github.com/Tsubasa-2005/sumaregi-go/types"
)

func PrintResponse(response types.GetProductsResponse) {
	v := reflect.ValueOf(response)
	if v.Kind() != reflect.Slice {
		log.Fatalf("Error: provided variable is not a slice")
	}

	if v.Len() == 0 {
		fmt.Println("No data to display")
		return
	}

	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i)
		elemType := elem.Type()

		// 横の区切りを追加
		fmt.Println("+----------------------+----------------------+")
		fmt.Println("| Field               | Value                |")
		fmt.Println("+----------------------+----------------------+")

		for j := 0; j < elem.NumField(); j++ {
			fieldName := elemType.Field(j).Name
			fieldValue := fmt.Sprintf("%v", elem.Field(j).Interface())

			// フィールド名とフィールド値の表示
			fmt.Printf("| %-20s | %-20s |\n", fieldName, fieldValue)
		}

		// エントリの区切り
		fmt.Println("+----------------------+----------------------+\n")
	}
}
