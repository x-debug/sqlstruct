package tests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dsn string = "root:root@tcp(127.0.0.1)/test?parseTime=true&charset=utf8"

type closeHandler func()

func initDB(dsn string) (db *sqlx.DB, close closeHandler, err error) {
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}

	return db, func() { db.Close() }, err
}

func TestNull(t *testing.T) {
}

func TestNormal(t *testing.T) {
	db, closef, err := initDB(dsn)

	if err != nil {
		t.Errorf("initDb error: %v", err)
	}

	var row1 VTest
	err = db.Get(&row1, "SELECT * FROM v_test WHERE id=1")
	if err != nil {
		t.Errorf("db.Get1 error: %v", err)
	}

	var row2 VTest
	err = db.Get(&row2, "SELECT * FROM v_test WHERE id=2")
	if err != nil {
		t.Errorf("db.Get2 error: %v", err)
	}

	//check init
	if row1.FInt != 1 {
		t.Errorf("FInt error, expected: %d, actual: %d", 1, row1.FInt)
	}

	if row1.FIntNull.Valid {
		t.Errorf("FInt vaild error")
	}

	if row2.FIntNull.Int64 != 1 || !row2.FIntNull.Valid {
		t.Errorf("FInt error, expected: %d, actual: %d", 1, row1.FIntNull.Int64)
	}

	//check tinyint
	if row1.FTint != 2 {
		t.Errorf("FTint error, expected: %d, actual: %d", 2, row1.FTint)
	}

	if row1.FTintNull.Valid {
		t.Errorf("FTint vaild error")
	}

	if row2.FTintNull.Int64 != 2 || !row2.FTintNull.Valid {
		t.Errorf("FTint error, expected: %d, actual: %d", 2, row1.FTintNull.Int64)
	}

	//check smallint
	if row1.FSint != 3 {
		t.Errorf("FSint error, expected: %d, actual: %d", 3, row1.FSint)
	}

	if row1.FSintNull.Valid {
		t.Errorf("FSint vaild error")
	}

	if row2.FSintNull.Int64 != 3 || !row2.FSintNull.Valid {
		t.Errorf("FSint error, expected: %d, actual: %d", 3, row1.FSintNull.Int64)
	}

	//check mediumint
	if row1.FMint != 4 {
		t.Errorf("FMint error, expected: %d, actual: %d", 4, row1.FMint)
	}

	if row1.FMintNull.Valid {
		t.Errorf("FMint vaild error")
	}

	if row2.FMintNull.Int64 != 4 || !row2.FMintNull.Valid {
		t.Errorf("FMint error, expected: %d, actual: %d", 4, row1.FMintNull.Int64)
	}

	//check bitint
	if row1.FBint != 5 {
		t.Errorf("FBint error, expected: %d, actual: %d", 5, row1.FBint)
	}

	if row1.FBintNull.Valid {
		t.Errorf("FBint vaild error")
	}

	if row2.FBintNull.Int64 != 5 || !row2.FBintNull.Valid {
		t.Errorf("FBint error, expected: %d, actual: %d", 5, row1.FBintNull.Int64)
	}

	//check bit
	if !row1.FBit {
		t.Errorf("FBit error, expected: %v, actual: %v", true, row1.FBit)
	}

	if row1.FBitNull.Valid {
		t.Errorf("FBit vaild error")
	}

	if bool(row2.FBitNull.Bit) || !row2.FBitNull.Valid {
		t.Errorf("FBit error, expected: %v, actual: %v", false, row1.FBitNull.Bit)
	}

	closef()
}
