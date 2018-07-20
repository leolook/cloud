package tools

import (
	"fmt"
	"strings"
)

type Struct struct {
	Field string
	Type  string
}

//produce struct
func ToStruct(table string, data []*Struct) string {
	name := Hump(table)
	ru := []rune(name)
	name = strings.ToUpper(string(ru[0])) + string(ru[1:])
	str := "\n"
	str += "type " + name + " struct {\n"

	num := MaxSpace(data) + 3
	for _, v := range data {
		name := Field(v.Field)
		space := Space(num - len(name))
		name = fmt.Sprintf("%s %s", name, space)
		typ := Typ(v.Type)
		typ = fmt.Sprintf("%s %s", typ, Space(8-len(typ)))
		tag := fmt.Sprintf("%s %s", JSONTag(v.Field, 0), GormTag(v.Field))

		str += fmt.Sprintf("   %s %s `%s`\n", name, typ, tag)
	}

	str += "}"
	str += "\n"
	return str
}

//max length field
func MaxSpace(data []*Struct) int {
	num := 1
	for _, v := range data {
		if len(Field(v.Field)) > num {
			num = len(Field(v.Field))
		}
	}
	return num
}

func Space(num int) string {
	space := ""
	for i := 0; i < num; i++ {
		space += " "
	}
	return space
}

//struct field
func Field(field string) string {
	ru := []rune(Hump(field))
	return strings.ToUpper(string(ru[0])) + string(ru[1:])
}

//struct type
func Typ(typ string) string {

	if strings.Contains(typ, "int") {
		return "int64"
	}

	if strings.Contains(typ, "varchar") {
		return "string"
	}

	if strings.Contains(typ, "double") {
		return "float64"
	}

	if strings.Contains(typ, "timestamp") {
		return "*time.Time"
	}

	return "interface{}"
}

//struct tag
//0-保持
//1-驼峰
func JSONTag(field string, typ int) string {
	tag := field
	if typ == 1 {
		tag = Hump(tag)
	}
	str := fmt.Sprintf(`json:"%s"`, tag)
	return str
}

//struct tag
//0-保持
//1-驼峰
func GormTag(field string) string {
	str := fmt.Sprintf(`gorm:"column:%s"`, field)
	return str
}

//驼峰
//body_age body + Age
//body_age_tmp body+Age+Tmp
func Hump(field string) string {
	if strings.Index(field, "_") == -1 {
		return field
	}
	arr := strings.Split(field, "_")

	if len(arr) == 1 {
		return arr[0]
	}

	field = arr[0]
	for i := 1; i < len(arr); i++ {
		ru := []rune(arr[i])
		field += strings.ToUpper(string(ru[0])) + string(ru[1:])
	}
	return field
}
