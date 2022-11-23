package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"starfission_go_api/cmd/mysqlmd/mysql"
	"starfission_go_api/pkg/convert"

	"gorm.io/gorm"
)

type tableInfo struct {
	Name    string         `db:"table_name"`    // name
	Comment sql.NullString `db:"table_comment"` // comment
}

type tableColumn struct {
	OrdinalPosition uint16         `db:"ORDINAL_POSITION"` // position
	ColumnName      string         `db:"COLUMN_NAME"`      // name
	ColumnType      string         `db:"COLUMN_TYPE"`      // column_type
	DataType        string         `db:"DATA_TYPE"`        // data_type
	ColumnKey       sql.NullString `db:"COLUMN_KEY"`       // key
	IsNullable      string         `db:"IS_NULLABLE"`      // nullable
	Extra           sql.NullString `db:"EXTRA"`            // extra
	ColumnComment   sql.NullString `db:"COLUMN_COMMENT"`   // comment
	ColumnDefault   sql.NullString `db:"COLUMN_DEFAULT"`   // default value
}

var (
	dbAddr      string
	dbUser      string
	dbPass      string
	dbName      string
	genTables   string
	tablePrefix string
	repoSaveDir string
)

func init() {
	addr := flag.String("addr", "", "请输入 db 地址，例如：127.0.0.1:3306\n")
	user := flag.String("user", "", "请输入 db 用户名\n")
	pass := flag.String("pass", "", "请输入 db 密码\n")
	name := flag.String("name", "", "请输入 db 名称\n")
	table := flag.String("tables", "*", "请输入 table 名称，默认为“*”，多个可用“,”分割\n")
	prefix := flag.String("prefix", "*", "请输入 prefix 名称，默认为“*”，多个可用“,”分割\n")
	saveDir := flag.String("savedir", "", "请输入 保存目录\n")

	flag.Parse()

	dbAddr = *addr
	dbUser = *user
	dbPass = *pass
	dbName = strings.ToLower(*name)
	genTables = strings.ToLower(*table)
	tablePrefix = *prefix
	repoSaveDir = *saveDir

}

func main() {

	// 初始化 DB
	db, err := mysql.New(dbAddr, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal("new db err", err)
	}

	defer func() {
		if err := db.DbClose(); err != nil {
			log.Println("db close err", err)
		}
	}()

	tables, err := queryTables(db.GetDb(), dbName, genTables)
	if err != nil {
		log.Println("query tables of database err", err)
		return
	}

	var tableSaveName string

	for _, table := range tables {

		if tablePrefix != "*" {
			tableSaveName = strings.Replace(table.Name, tablePrefix, "", 1)
		} else {
			tableSaveName = table.Name
		}

		var dbPath string
		if repoSaveDir != "" {
			dbPath = "./internal/model/" + repoSaveDir
			_ = os.Mkdir(dbPath, 0766)
		} else {
			dbPath = "./internal/model/"
		}

		filepath := dbPath + "/" + tableSaveName + "_repo"
		_ = os.Mkdir(filepath, 0766)
		fmt.Println("create dir : ", filepath)

		mdName := fmt.Sprintf("%s/gen_table.md", filepath)
		mdFile, err := os.OpenFile(mdName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("markdown file error %v\n", err.Error())
			return
		}
		fmt.Println("  └── file : ", dbPath+"/"+tableSaveName+"_repo/gen_table.md")

		modelName := fmt.Sprintf("%s/gen_model.go", filepath)
		modelFile, err := os.OpenFile(modelName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0766)
		if err != nil {
			fmt.Printf("create and open model file error %v\n", err.Error())
			return
		}
		fmt.Println("  └── file : ", dbPath+"/"+tableSaveName+"_repo/gen_model.go")

		modelContent := fmt.Sprintf("package %s%s\n", tableSaveName, "_repo")
		modelContent += fmt.Sprintf(`import "starfission_go_api/pkg/timef"`)
		modelContent += fmt.Sprintf("\n\n// %s \n", table.Comment.String)
		modelContent += fmt.Sprintf("//go:generate gormgen -structs %s -input . -pre %s \n", capitalize(tableSaveName), tablePrefix)
		modelContent += fmt.Sprintf("type %s struct {\n", capitalize(tableSaveName))

		tableContent := fmt.Sprintf("#### %s.%s \n", dbName, tableSaveName)
		if table.Comment.String != "" {
			tableContent += table.Comment.String + "\n"
		}
		tableContent += "\n" +
			"| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |\n" +
			"| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |\n"

		columnInfo, columnInfoErr := queryTableColumn(db.GetDb(), dbName, table.Name)
		if columnInfoErr != nil {
			continue
		}
		for _, info := range columnInfo {
			tableContent += fmt.Sprintf(
				"| %d | %s | %s | %s | %s | %s | %s | %s |\n",
				info.OrdinalPosition,
				info.ColumnName,
				strings.ReplaceAll(strings.ReplaceAll(info.ColumnComment.String, "|", "\\|"), "\n", ""),
				info.ColumnType,
				info.ColumnKey.String,
				info.IsNullable,
				info.Extra.String,
				info.ColumnDefault.String,
			)

			gormAdd := []string{}

			gormAdd = append(gormAdd, "column:"+info.ColumnName)

			if info.ColumnKey.Valid && info.ColumnKey.String == "PRI" {
				gormAdd = append(gormAdd, "primary_key")
			}
			if info.Extra.Valid && info.Extra.String == "auto_increment" {
				gormAdd = append(gormAdd, "AUTO_INCREMENT")
			}
			if textType(info.DataType) == "timef.Time" {
				gormAdd = append(gormAdd, "time")
			}

			if info.ColumnDefault.Valid {
				gormAdd = append(gormAdd, "default:"+info.ColumnDefault.String)
			}

			jsonAdd := convert.Case2LowerCamel(info.ColumnName)
			formAdd := convert.Case2LowerCamel(info.ColumnName)

			if len(gormAdd) > 0 {
				modelContent += fmt.Sprintf("%s %s `%s %s %s` // %s\n", capitalize(info.ColumnName), textType(info.DataType), "gorm:\""+strings.Join(gormAdd, ";")+"\"", "json:\""+jsonAdd+"\"", "form:\""+formAdd+"\"", info.ColumnComment.String)
			} else {
				modelContent += fmt.Sprintf("%s %s // %s\n", capitalize(info.ColumnName), textType(info.DataType), info.ColumnComment.String)
			}

		}

		mdFile.WriteString(tableContent)
		mdFile.Close()

		modelContent += "}\n"
		modelFile.WriteString(modelContent)
		modelFile.Close()

	}

}

func queryTables(db *gorm.DB, dbName string, tableName string) ([]tableInfo, error) {
	var tableCollect []tableInfo
	var tableArray []string
	var commentArray []sql.NullString

	sqlTables := fmt.Sprintf("SELECT `table_name`,`table_comment` FROM `information_schema`.`tables` WHERE `table_schema`= '%s'", dbName)
	rows, err := db.Raw(sqlTables).Rows()
	if err != nil {
		return tableCollect, err
	}
	defer rows.Close()

	for rows.Next() {
		var info tableInfo
		err = rows.Scan(&info.Name, &info.Comment)
		if err != nil {
			fmt.Printf("execute query tables action error,had ignored, detail is [%v]\n", err.Error())
			continue
		}

		tableCollect = append(tableCollect, info)
		tableArray = append(tableArray, info.Name)
		commentArray = append(commentArray, info.Comment)
	}

	// filter tables when specified tables params
	if tableName != "*" {
		tableCollect = nil
		chooseTables := strings.Split(tableName, ",")
		indexMap := make(map[int]int)
		for _, item := range chooseTables {
			subIndexMap := getTargetIndexMap(tableArray, item)
			for k, v := range subIndexMap {
				if _, ok := indexMap[k]; ok {
					continue
				}
				indexMap[k] = v
			}
		}

		if len(indexMap) != 0 {
			for _, v := range indexMap {
				var info tableInfo
				info.Name = tableArray[v]
				info.Comment = commentArray[v]
				tableCollect = append(tableCollect, info)
			}
		}
	}

	//fmt.Println(tableCollect)

	return tableCollect, err
}

func queryTableColumn(db *gorm.DB, dbName string, tableName string) ([]tableColumn, error) {
	// 定义承载列信息的切片
	var columns []tableColumn

	sqlTableColumn := fmt.Sprintf("SELECT `ORDINAL_POSITION`,`COLUMN_NAME`,`COLUMN_TYPE`,`DATA_TYPE`,`COLUMN_KEY`,`IS_NULLABLE`,`EXTRA`,`COLUMN_COMMENT`,`COLUMN_DEFAULT` FROM `information_schema`.`columns` WHERE `table_schema`= '%s' AND `table_name`= '%s' ORDER BY `ORDINAL_POSITION` ASC",
		dbName, tableName)

	rows, err := db.Raw(sqlTableColumn).Rows()
	if err != nil {
		fmt.Printf("execute query table column action error, detail is [%v]\n", err.Error())
		return columns, err
	}
	defer rows.Close()

	for rows.Next() {
		var column tableColumn
		err = rows.Scan(
			&column.OrdinalPosition,
			&column.ColumnName,
			&column.ColumnType,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.Extra,
			&column.ColumnComment,
			&column.ColumnDefault)
		if err != nil {
			fmt.Printf("query table column scan error, detail is [%v]\n", err.Error())
			return columns, err
		}
		columns = append(columns, column)
	}

	return columns, err
}

func getTargetIndexMap(tableNameArr []string, item string) map[int]int {
	indexMap := make(map[int]int)
	for i := 0; i < len(tableNameArr); i++ {
		if match, _ := regexp.MatchString(item, tableNameArr[i]); match {
			if _, ok := indexMap[i]; ok {
				continue
			}
			indexMap[i] = i
		}
	}
	return indexMap
}

func capitalize(s string) string {
	var upperStr string
	chars := strings.Split(s, "_")
	for _, val := range chars {
		vv := []rune(val)
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				if vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					upperStr += string(vv[i])
				}
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	//upperStr = strings.Replace(upperStr, "Id", "ID", 1)
	//upperStr = strings.Replace(upperStr, "Url", "URL", 1)
	//upperStr = strings.Replace(upperStr, "Sn", "SN", 1)
	return upperStr
}

func textType(s string) string {
	var mysqlTypeToGoType = map[string]string{
		"tinyint":    "int32",
		"smallint":   "int32",
		"mediumint":  "int32",
		"int":        "int32",
		"integer":    "int64",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"date":       "string",
		"time":       "string",
		"year":       "string",
		"datetime":   "timef.Time",
		"timestamp":  "timef.Time",
		"char":       "string",
		"varchar":    "string",
		"tinyblob":   "string",
		"tinytext":   "string",
		"blob":       "string",
		"text":       "string",
		"mediumblob": "string",
		"mediumtext": "string",
		"longblob":   "string",
		"longtext":   "string",
	}
	return mysqlTypeToGoType[s]
}
