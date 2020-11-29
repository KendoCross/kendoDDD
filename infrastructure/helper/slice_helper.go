package helper

import (
	"reflect"
	"strconv"
	"strings"
)

//字符串`,`分隔，转为[]int
func ConvertStr2Ints(v string) []int {
	is := make([]int, 0, 6)

	ss := strings.Split(v, ",")
	for _, s := range ss {
		i, e := strconv.Atoi(s)
		if e == nil {
			is = append(is, i)
		}
	}

	return is
}

//[]Int转为`,`分割的字符串
func ConvertIntsToStr(is []int) string {

	ss := make([]string, 0, 6)
	for _, v := range is {
		ss = append(ss, strconv.Itoa(v))
	}
	return strings.Join(ss, ",")

	//僵尸代码：比较两种写法差异,居然上面的写法性能更好！
	// var buf bytes.Buffer
	// for i, v := range is {
	// 	if i > 0 {
	// 		buf.WriteString(",")
	// 	}
	// 	fmt.Fprintf(&buf, "%d", v)
	// }
	// return buf.String()
}

//[]Int64转为`,`分割的字符串
func ConvertInt64sToStr(is []int64) string {

	ss := make([]string, 0, 6)
	for _, v := range is {
		ss = append(ss, strconv.FormatInt(v, 10))
	}
	return strings.Join(ss, ",")
}

//RemoveRep 移除重复值
func RemoveRep(slc []int) []int {
	if len(slc) > 1024 {
		return removeRepByMap(slc)
	}
	return removeRepByLoop(slc)
}

// 通过两重循环过滤重复元素
func removeRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

// 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

//修改原切片的
func ConvertInt642Ints(slc []int64) []int {
	rst := make([]int, len(slc), len(slc))
	for i, v := range slc {
		rst[i] = int(v)
	}
	return rst
}

//判定某值是否在切片里
func IsValInArr(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}
