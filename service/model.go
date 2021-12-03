package service

import "reflect"

var Model = model{}

type model struct{}

func (m model) GetTable(entity interface{}) string {
	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = reflect.TypeOf(entity).Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		table := t.Field(i).Tag.Get("table")
		if table != "" {
			return table
		}
	}

	return m.convert2SnakeCase(t.Name())
}

func (m model) GetPK(entity interface{}) string {
	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = reflect.TypeOf(entity).Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		pk := t.Field(i).Tag.Get("pk")
		if pk != "" {
			return pk
		}
	}

	if t.NumField() < 1 {
		return "id"
	}

	return m.convert2SnakeCase(t.Field(0).Name)
}

func (m model) GetColumn(column interface{}) interface{} {
	if reflect.TypeOf(column).Kind() != reflect.Ptr {
		return column
	}

	t := reflect.TypeOf(column).Elem()
	v := reflect.ValueOf(column).Elem()
	for i := 0; i < v.NumField(); i++ {
		c := t.Field(i).Tag.Get("column")
		if c != "" {
			v.Field(i).SetString(c)
		} else {
			v.Field(i).SetString(m.convert2SnakeCase(t.Field(i).Name))
		}
	}

	return v.Interface()
}

func (m model) convert2SnakeCase(s string) string {
	r := []rune("")
	for i, c := range s {
		if 65 <= c && c <= 90 {
			if i == 0 || 65 <= s[i-1] && s[i-1] <= 90 {
				r = append(r, c+32)
			} else {
				r = append(r, 95, c+32)
			}
		} else {
			r = append(r, c)
		}
	}

	return string(r)
}
