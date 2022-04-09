/*package main

import (
	"fmt"
	"reflect"
)

type rpcError struct {
	code    int
	message string
}

//rpcError结构体实现了Error接口
func (r *rpcError) Error() string {
	return fmt.Sprintln(r.message, r.code)
}

//这个返回的是结构体但是这个结构体指针实现了error接口，就可以当做是这个接口的一个实例
func newrpcError(code int, msg string) error {
	return &rpcError{
		code:    code,
		message: msg,
	}
}
func asErr(err error) error {
	return err
}
func p(v interface{}) {
	print(v)
}

func main() {

	name := "umbrella"
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.ValueOf(name))
	v := reflect.ValueOf(1)
	fmt.Println(v)
	i := v.Interface().(int)
	fmt.Println(i)
	t := 1
	y := reflect.ValueOf(&t)
	y.Elem().SetInt(12)
	fmt.Println(t)

	var r error = newrpcError(404, "not found")
	fmt.Println(r)
	err := asErr(r)
	fmt.Println(err)
}
*/
package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var track []int //遍历路径
	var result [][]int
	var backTree func(startIndex, sum int)
	backTree = func(startIndex, sum int) {
		if sum > n {
			return
		}
		if len(track) == k {
			temp := make([]int, k)
			copy(temp, track)
			//相加之和为n
			if sum == n {
				result = append(result, temp)
				//result = append(result, track)
			}
			return
		}
		//剪枝操作，k-len(track)表示还剩下多少个可以填充的元素
		//初始的循环次数是9-k+1，1-9，选择7个数，那么循环次数是从1-3，就是9-7+1
		//初始情况下最外层循环i可以是1-7，里面那一层就是2-8（i只能选到8，再下一层就是3-9，i最多到9了），再往下那一层3-9
		num := 9 - k + 1 + len(track)
		for i := startIndex; i <= num; i++ {
			sum += i
			track = append(track, i)
			backTree(i+1, sum)           //递归
			track = track[:len(track)-1] //回溯
			sum -= i
		}
	}
	backTree(1, 0)
	return result
}

func main() {
	sum3 := combinationSum3(3, 28)
	fmt.Println(sum3)
}
