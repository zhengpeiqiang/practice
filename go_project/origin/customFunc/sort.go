package customFunc

import (
	"fmt"
	"reflect"
)

/*
* action 排序，数字和字符串排序
* param sortMode 排序方式 asc desc
* param sortMethod 排序方法 冒泡，快速等等; bubble quick
* param values 数据
* return interface 传入什么格式，返回什么格式
 */
func Sort(sortMode string, sortMethod string, values interface{}) {
	typeOf := reflect.TypeOf(values)
	switch sortMethod {
	case "bubble":
		if sortMode == "asc" {
			switch typeOf.String() {
			case "[]string":
				bubbleSortAscString(values.([]string))
			case "[]int":
				bubbleSortAscInt(values.([]int))
			default:
				bubbleSortAscInt(values.([]int))
			}
		}
		if sortMode == "desc" {
			switch typeOf.String() {
			case "[]string":
				bubbleSortDescString(values.([]string))
				break
			case "[]int":
				bubbleSortDescInt(values.([]int))
				break
			default:
				bubbleSortDescInt(values.([]int))
				break
			}
		}
		break
	case "quick":
		if sortMode == "asc" {
			switch typeOf.String() {
			case "[]string":
				fmt.Println("未完善")
				//quickSortAscString(values.([]string),0,len(values.([]string))-1)
			case "[]int":
				quickSortAscInt(values.([]int),0,len(values.([]int))-1)
			default:
				quickSortAscInt(values.([]int),0,len(values.([]int))-1)
			}
		}
		if sortMode == "desc" {
			switch typeOf.String() {
			case "[]string":
				fmt.Println("未完善")
				break
			case "[]int":
				fmt.Println("未完善")
				break
			default:
				fmt.Println("未完善")
				break
			}
		}
		break
	}
}

func bubbleSortAscString(values []string) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if len(values[i]) > len(values[j]) {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func bubbleSortAscInt(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func bubbleSortDescInt(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func bubbleSortDescString(values []string) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if len(values[i]) < len(values[j]) {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}

func quickSortAscString(a []string, left int, right int) {
	if left >= right { //一定是left >= right
		return
	}
	temp := len(a[left])
	start := left
	stop := right
	for right != left {
		for right > left && len(a[right]) >= temp {
			right--
		}
		for left < right && len(a[left]) <= temp {
			left++
		}
		if right > left {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[right], a[start] = a[left], a[right]
	quickSortAscString(a, start, left)
	quickSortAscString(a, right+1, stop)
}

func quickSortAscInt(a []int, left int, right int) {
	if left >= right { //一定是left >= right
		return
	}
	temp := a[left]
	start := left
	stop := right
	for right != left {
		for right > left && a[right] >= temp {
			right--
		}
		for left < right && a[left] <= temp {
			left++
		}
		if right > left {
			a[right], a[left] = a[left], a[right]
		}
	}
	a[right], a[start] = temp, a[right]
	quickSortAscInt(a, start, left)
	quickSortAscInt(a, right+1, stop)
}