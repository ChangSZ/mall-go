package copy

import (
	"fmt"
	"reflect"
	"time"
)

// AssignStruct 将src中有值的字段赋值到dst中
//
// - 是将相同字段名中src值赋给dst中对应字段
// - 入参必须是结构体对象引用
// - 若结构体中存在切片, 请先初始化至src\dst一致
// - 如果存在内联, 保证内联结构体名称一致
func AssignStruct(src, dst interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	if src == nil || reflect.ValueOf(src).IsNil() ||
		dst == nil || reflect.ValueOf(dst).IsNil() {
		fmt.Println("src or dst is nil")
		return
	}
	assignStructFields(reflect.ValueOf(src).Elem(), reflect.ValueOf(dst).Elem())
}

// assignStructFields 将src中有值的字段赋值到dst中, 递归至成员变量最小类型
func assignStructFields(src, dst reflect.Value) {
	srcType := src.Type()
	for i := 0; i < srcType.NumField(); i++ {
		field := srcType.Field(i)
		fieldName := field.Name

		srcFieldValue := src.FieldByName(fieldName)
		dstFieldValue := dst.FieldByName(fieldName)

		// 检查字段是否有效
		if srcFieldValue.IsValid() && dstFieldValue.IsValid() {
			// 检查 srcFieldValue 是否为 nil，且类型为指针
			if srcFieldValue.Kind() == reflect.Ptr && srcFieldValue.IsNil() {
				// 如果 srcFieldValue 是 nil 指针，则跳过它
				continue
			}
			if srcFieldValue.IsZero() {
				continue
			}

			if field.Type == reflect.TypeOf(time.Time{}) {
				dstFieldValue.Set(srcFieldValue)
				continue
			}

			if srcFieldValue.Kind() == reflect.Struct {
				assignStructFields(srcFieldValue, dstFieldValue)
				continue
			}

			if srcFieldValue.Kind() == reflect.Slice {
				assignSliceFields(srcFieldValue, dstFieldValue)
				continue
			}

			// 判断类型是否一样
			if srcFieldValue.Kind() == dstFieldValue.Kind() {
				// 设置 dstFieldValue 的值
				dstFieldValue.Set(srcFieldValue)
			}
		}
	}
}

// assignSliceFields 复制切片
func assignSliceFields(src, dst reflect.Value) {
	elemType := src.Type().Elem()
	// 若元素类型是结构体且源切片元素个数等于目标切片元素个数时, 依次递归复制
	if elemType.Kind() == reflect.Struct && src.Len() == dst.Len() {
		// 依次处理每个元素
		for j := 0; j < src.Len(); j++ {
			assignStructFields(src.Index(j), dst.Index(j))
		}
	} else {
		if src.Kind() == dst.Kind() {
			dst.Set(src)
		}
	}
}

// DeepCopy 深拷贝
func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}
	return value
}
