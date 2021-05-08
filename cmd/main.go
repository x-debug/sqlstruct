package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	sql2struct "github.com/x-debug/sqlstruct/internal"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var packageName string
var fileName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		tables, err := dbModel.GetTables(dbName)
		if err != nil {
			log.Fatalf("dbModel.GetTables err: %v", err)
		}

		file, err := os.Create(fileName)
		defer file.Close()
		w := bufio.NewWriter(file)

		w.WriteString(fmt.Sprintf("package %s\n\nimport (\n  \"time\"\n  \"github.com/shopspring/decimal\"\n)\n\n", packageName))

		for _, table := range tables {
			columns, err := dbModel.GetColumns(dbName, table.Name)
			if err != nil {
				log.Fatalf("dbModel.GetColumns err: %v", err)
			}

			template := sql2struct.NewStructTemplate()
			templateColumns := template.AssemblyColumns(columns)
			err = template.Generate(w, table.Name, templateColumns)
			w.WriteString("\n\n")
			if err != nil {
				log.Fatalf("template.Generate err: %v", err)
			}
		}
		w.Flush()
	},
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&packageName, "package", "", "", "请输入包名称")
	sql2structCmd.Flags().StringVarP(&fileName, "file", "", "", "请输入文件名称")

	rootCmd.AddCommand(sqlCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
