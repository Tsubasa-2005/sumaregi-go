package utils

import (
	"fmt"
	"log"
	"reflect"
)

// 単一の構造体の中身を表示する関数
func PrintSingleStruct(data interface{}) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		log.Fatalf("Error: provided variable is not a struct")
	}

	elemType := v.Type()

	// 横の区切りを追加
	fmt.Println("+----------------------+----------------------+")
	fmt.Println("| Field               | Value                |")
	fmt.Println("+----------------------+----------------------+")

	for i := 0; i < v.NumField(); i++ {
		fieldName := elemType.Field(i).Name
		fieldValue := fmt.Sprintf("%v", v.Field(i).Interface())

		// フィールド名とフィールド値の表示
		fmt.Printf("| %-20s | %-20s |\n", fieldName, fieldValue)
	}

	// エントリの区切り
	fmt.Println("+----------------------+----------------------+\n")
}

// ジェネリックな表示関数
func PrintResponse(response interface{}) {
	v := reflect.ValueOf(response)

	// スライスかどうかをチェック
	if v.Kind() == reflect.Slice {
		if v.Len() == 0 {
			fmt.Println("No data to display")
			return
		}

		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			PrintSingleStruct(elem.Interface())
		}
	} else if v.Kind() == reflect.Struct {
		// スライスでない場合、単一の構造体を表示
		PrintSingleStruct(response)
	} else {
		log.Fatalf("Error: provided variable is not a slice or struct")
	}
}
