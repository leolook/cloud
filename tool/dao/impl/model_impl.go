package impl

import (
	"cloud/tool/mysql"
	pb "cloud/tool/protocol"
	"cloud/tool/tools"
	"database/sql"
	"fmt"
	. "github.com/gitbubhwt/baseserver/util"
	"log"
)

type Model struct{}

func (m Model) CreateModel(url, table string) (string, error) {
	db := mysql.GetClient(url)
	if db == nil {
		return "", Error(pb.ERR_FAIL_CONNECT_DB, "Fail connect db")
	}
	rows, err := db.Query(fmt.Sprintf("DESCRIBE %s ", table))
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return "", Error(pb.ERR_DB_QUERY, fmt.Sprintf("%v", err))
	}
	data := make([]*tools.Struct, 0)
	for rows.Next() {
		var field, typ, nul, key, defau, extra sql.NullString
		err = rows.Scan(&field, &typ, &nul, &key, &defau, &extra)
		if err != nil {
			log.Println(fmt.Sprintf("field=%v", err))
			return "", err
		}
		tmp := &tools.Struct{
			Field: field.String,
			Type:  typ.String,
		}
		data = append(data, tmp)
	}
	str := tools.ToStruct(table, data)
	return str, nil
}

func (m Model) ShowTables(url string) ([]string, error) {
	db := mysql.GetClient(url)
	if db == nil {
		return nil, Error(pb.ERR_FAIL_CONNECT_DB, "Fail connect db")
	}
	row, err := db.Query(fmt.Sprintf("show tables"))
	if err != nil {
		return nil, Error(pb.ERR_DB_QUERY, err.Error())
	}
	list := make([]string, 0, 10)
	for row.Next() {
		var name string
		err = row.Scan(&name)
		if err != nil {
			return nil, Error(pb.ERR_DB_QUERY, err.Error())
		}
		list = append(list, name)
	}
	return list, nil
}
