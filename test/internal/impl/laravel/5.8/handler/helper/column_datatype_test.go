package helper

import (
	"asher/internal/impl/laravel/5.8/handler"
	"fmt"
	"testing"
)

func Test_Columns(t *testing.T) {
	colHandler := new(handler.ColumnHandler)
	fmt.Println(colHandler.ColTypeSwitcher("string|255,43,32", "desc", nil))
	fmt.Println(colHandler.ColTypeSwitcher("integer", "user_id", nil))
	fmt.Println(colHandler.ColTypeSwitcher("enum", "user_id", []string{"a", "b"}))

}
