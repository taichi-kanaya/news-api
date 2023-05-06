/*
カスタムバリデーションのルールを定義する
*/
package validation

import (
	"strconv"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Ginにカスタムバリデーションのルールを登録する
func RegisterValidation() {
	if v, isValidator := binding.Validator.Engine().(*validator.Validate); isValidator {
		v.RegisterValidation("string-min-value", stringMinValue)
	}
}

// 文字列型の数字が指定された値以上かを判定するバリデーション
//
// Parameters:
//   - fl: validator.FieldLevel (バリデーション対象のフィールド情報)
//
// Returns:
//   - bool: バリデーション結果(true: OK, false: NG)
func stringMinValue(fl validator.FieldLevel) bool {
	// 許容する最小値を取得
	minValue, err := strconv.ParseInt(fl.Param(), 10, 64)
	if err != nil {
		return false
	}

	// フィールドの値を数値に変換
	fieldValue, err := strconv.ParseInt(fl.Field().String(), 10, 64)
	if err != nil {
		return false
	}

	// 最小値以上かどうかを判定
	return fieldValue >= minValue
}
