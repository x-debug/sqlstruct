package tests

import (
	"github.com/shopspring/decimal"
	"github.com/x-debug/sqlstruct/pkg"
)

type Brand struct {
	// id
	Id int32 `json:"id" db:"id"`
	// name
	Name pkg.String `json:"name" db:"name"`
}

func (model Brand) TableName() string {
	return "brand"
}

type Category struct {
	// id
	Id int32 `json:"id" db:"id"`
	// name
	Name pkg.String `json:"name" db:"name"`
}

func (model Category) TableName() string {
	return "category"
}

type Goods struct {
	// id
	Id int32 `json:"id" db:"id"`
	// name
	Name pkg.String `json:"name" db:"name"`
	// price
	Price decimal.Decimal `json:"price" db:"price"`
	// category
	Category pkg.String `json:"category" db:"category"`
	// brand
	Brand pkg.String `json:"brand" db:"brand"`
	// category_id
	CategoryId pkg.Int `json:"category_id" db:"category_id"`
	// brand_id
	BrandId pkg.Int `json:"brand_id" db:"brand_id"`
}

func (model Goods) TableName() string {
	return "goods"
}

type T struct {
	// id
	Id int32 `json:"id" db:"id"`
	// a
	A pkg.Int `json:"a" db:"a"`
	// b
	B pkg.Int `json:"b" db:"b"`
}

func (model T) TableName() string {
	return "t"
}

type VTest struct {
	// f_int
	FInt int32 `json:"f_int" db:"f_int"`
	// f_int_null
	FIntNull pkg.Int `json:"f_int_null" db:"f_int_null"`
	// f_tint
	FTint int8 `json:"f_tint" db:"f_tint"`
	// f_tint_null
	FTintNull pkg.Int `json:"f_tint_null" db:"f_tint_null"`
	// f_sint
	FSint int `json:"f_sint" db:"f_sint"`
	// f_sint_null
	FSintNull pkg.Int `json:"f_sint_null" db:"f_sint_null"`
	// f_mint
	FMint int64 `json:"f_mint" db:"f_mint"`
	// f_mint_null
	FMintNull pkg.Int `json:"f_mint_null" db:"f_mint_null"`
	// f_bint
	FBint int64 `json:"f_bint" db:"f_bint"`
	// f_bint_null
	FBintNull pkg.Int `json:"f_bint_null" db:"f_bint_null"`
	// f_bit
	FBit pkg.Bit `json:"f_bit" db:"f_bit"`
	// f_bit_null
	FBitNull pkg.NullBit `json:"f_bit_null" db:"f_bit_null"`
	// id
	Id int32 `json:"id" db:"id"`
}

func (model VTest) TableName() string {
	return "v_test"
}
