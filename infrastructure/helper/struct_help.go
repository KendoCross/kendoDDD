package helper

import "reflect"

func deepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

//StructCopy 相同名称字段结构复制，必须指针类型；源在前，目标在后
func StructCopy(srcPtr interface{}, dstPtr interface{}) {
	canCopy(srcPtr, dstPtr)

	srcv := reflect.ValueOf(srcPtr)
	dstv := reflect.ValueOf(dstPtr)

	srcValue := srcv.Elem()
	dstValue := dstv.Elem()
	srcfields := deepFields(srcValue.Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstValue.FieldByName(v.Name)
		src := srcValue.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func canCopy(srcPtr interface{}, dstPtr interface{}) {
	srcv := reflect.ValueOf(srcPtr)
	dstv := reflect.ValueOf(dstPtr)
	srct := reflect.TypeOf(srcPtr)
	dstt := reflect.TypeOf(dstPtr)

	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr || srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
}
