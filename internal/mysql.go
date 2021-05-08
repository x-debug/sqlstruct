package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

type Table struct {
	Name string
}

var DBNullTypeToStructType = map[string]string{
	"int":        "pkg.Int",
	"tinyint":    "pkg.Int",
	"smallint":   "pkg.Int",
	"mediumint":  "pkg.Int",
	"bigint":     "pkg.Int",
	"bit":        "pkg.NullBit",
	"bool":       "pkg.Bool",
	"enum":       "pkg.String",
	"set":        "pkg.String",
	"varchar":    "pkg.String",
	"char":       "pkg.String",
	"tinytext":   "pkg.String",
	"mediumtext": "pkg.String",
	"text":       "pkg.String",
	"longtext":   "pkg.String",
	"blob":       "pkg.String",
	"tinyblob":   "pkg.String",
	"mediumblob": "pkg.String",
	"longblob":   "pkg.String",
	"date":       "pkg.Time",
	"datetime":   "pkg.Time",
	"timestamp":  "pkg.Time",
	"time":       "pkg.Time",
	"float":      "pkg.Float",
	"double":     "pkg.Float",
	"decimal":    "decimal.NullDecimal",
}

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "pkg.Bit",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
	"decimal":    "decimal.Decimal",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}

func (m *DBModel) GetTables(dbName string) ([]*Table, error) {
	sql := "select table_name from information_schema.tables where table_schema=? and table_type='BASE TABLE'"

	rows, err := m.DBEngine.Query(sql, dbName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}

	defer rows.Close()

	var tables []*Table
	for rows.Next() {
		var table Table
		err := rows.Scan(&table.Name)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}

	return tables, nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		column.ColumnComment = strings.ReplaceAll(column.ColumnComment, "\r\n", "")
		column.ColumnComment = strings.ReplaceAll(column.ColumnComment, "\n", "")
		columns = append(columns, &column)
	}

	return columns, nil
}
