package inspect

import "reflect"

func ChildsFieldNames(obj interface{}) []string {
	var fields []string

	v := reflect.ValueOf(obj)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}

	return fields
}

func FieldsByType(obj interface{}, fieldType string) []string {
	var fields []string

	v := reflect.ValueOf(obj)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		// fmt.Printf("%d: %s %s %s = %v\n", i, t.Field(i).Name, f.Type(), f.Type().Kind(), f.Interface())
		if f.Type().Kind() == reflect.Struct {
			fields = append(fields, FieldsByType(f.Interface(), fieldType)...)
		}

		if f.Type().String() == fieldType {
			fields = append(fields, t.Field(i).Name)
		}
	}

	return fields
}

// type Attribute struct {
// 	Name string
// 	Type string
// 	Dice string
// 	Path string
// }

// var Attrs = []Attribute{
// 	{"x", "", "", ""},
// 	{"y", "", "", ""},
// 	{"z", "", "", ""},
// }

// func test() {
// 	st := struct {
// 		foo int
// 		bar string
// 		baz struct{}
// 	}{
// 		foo: 5, bar: "bar", baz: struct{}{},
// 	}

// 	allAttributes(st)
// }

// func allAttributes(s interface{}) {
// 	fields := inspect.FieldsByType(s, "string")
// 	fmt.Println(fields)

// 	for attr := range Attrs {
// 		inspect.FieldsByType(s, attr.Type)
// 	}
// }
