package main

import (
	"fmt"
	gin "github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/github"
	"net/http"
)

func Print[T1 any, T2 any](a T1, b T2) {
	fmt.Println(a, b)
}

func Print2(a, b interface{}) {
	fmt.Println(a, b)
}

func printArray() {
	/*var a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	// slice use make function
	s := make([]int, 5, 10)
	if s == nil {
		println("s is nill")
	}
	fmt.Println(s, len(s), cap(s))

	// sub-slice
	ss := []int{0, 1, 2, 3, 4, 5}
	ss = ss[2:5] // 2, 3, 4
	ss = ss[1:]  // 3, 4
	fmt.Println(ss)

	// slice append and copy
	sa := []int{0, 1}
	sa = append(sa, 2)
	fmt.Println(sa)

	sa = append(sa, 3, 4, 5)
	fmt.Println(sa)

	// map
	var idMap map[int]string
	fmt.Println(idMap)
	idMap = make(map[int]string)

	// 리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}
	fmt.Println(tickers)

	var m map[int]string

	m = make(map[int]string)
	m[901] = "Apple"

	str := m[901]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)

	delete(m, 901)
	fmt.Println(str)
	fmt.Println(m[901])*/

	// Check map key
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
		"AMZN": "Amazon",
	}

	val, exists := tickers["MSFT"] // val: value값, exitst: bool 존재 여부 값
	if !exists {
		fmt.Printf("No MSFT ticker %s\n", val)
	}

	// map array
	for key, val := range tickers {
		fmt.Println(key, val)
	}
}

// Integer 타입 제한 선언 (interface로 타입제한이 가능하다.)
type Integer interface {
	int8 | int16 | int32 | int64 | int
}

type Float interface {
	~float32 | ~float64 // ~ -> 해당 타입을 기본으로 하는 모든 별칭 타입들까지 포함
}

func add[T Integer](a, b T) T {
	return a + b
}

type TestModel struct {
	Id   int    `json:"id" binding:"required"` // binding required = not empty
	Name string `json:"name" binding:"required"`
}

func main() {
	r := gin.Default() // defalut settings

	Print("hello", "hello")
	Print("hello", 4)

	Print2("hello", "hello")
	Print2("hello", 4)
	// default string
	r.GET("/default/string", func(c *gin.Context) { // handler function
		c.String(http.StatusOK, "Hello world!!!")
	})

	// default json
	r.GET("/defalut/json", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"reponse": "Hello world!!!",
		})
	})

	// HTTP 파라미터 받기
	r.GET("/:name", func(c *gin.Context) {
		var val = c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"value": val,
		})
	})

	// Body로 들어오는 POST 요청처리
	r.POST("/add", func(c *gin.Context) {
		//var req := &Bind{}
		var data TestModel
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		}
	})

	printArray()
	r.Run("localhost:8080") // api를 호스트할 url과 포트번호
}
