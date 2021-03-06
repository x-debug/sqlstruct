package internal

import (
	"fmt"
	"io"
	"text/template"
)

const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	strcutTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
	IsNull  bool
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\" db:\"%s\""+"`", column.ColumnName, column.ColumnName)
		var typ string
		if column.IsNullable != "YES" {
			typ = DBTypeToStructType[column.DataType]
		} else {
			typ = DBNullTypeToStructType[column.DataType]
		}
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    typ,
			Tag:     tag,
			Comment: column.ColumnComment,
			IsNull:  column.IsNullable == "YES",
		})
	}

	return tplColumns
}

func (t *StructTemplate) Generate(buffer io.Writer, tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": UnderscoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(buffer, tplDB)
	if err != nil {
		return err
	}

	return nil
}
