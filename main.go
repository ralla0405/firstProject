package main

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/github"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

const (
	dsn = "host=localhost user=postgres password=@Demian2020 dbname=simple_bank port=5432 sslmode=disable TimeZone=Asia/Seoul"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB 연결에 실패하였습니다.")
	}

	// 테이블 자동 생성
	err = db.AutoMigrate(&Product{}, &User{})
	if err != nil {
		fmt.Println("Error occur")
	}

	// 생성
	db.Create(&Product{Code: "D123", Price: 5400})

	//var product Product
	/*db.First(&product, 1)
	db.First(&product, "Code = ?", "D123")
	tx := db.Find(&product, "ID = ?", 2)

	fmt.Println(tx, "123456")

	fmt.Println(product.Price)

	// 수정 하나의 필드만
	db.Model(&product).Update("Price", 200)
	fmt.Println(product.Price)

	// 수정 여러 필드
	db.Model(&product).Updates(Product{Price: 600, Code: "F42"})
	fmt.Println(product.Code)
	fmt.Println(product.Price)

	db.Model(&product).Updates(map[string]interface{}{"Price": 300, "Code": "F43"})
	fmt.Println(product.Code)
	fmt.Println(product.Price)

	// 삭제
	db.Delete(&product, 1)*/

	// 레코드 생성
	user := User{Name: "Logan", Age: 33, Birthday: time.Now()}

	result := db.Create(&user)

	fmt.Println(user.ID)             // 입려된 데이터의 primary key를 반환합니다.
	fmt.Println(result.Error)        // 에러를 반환합니다.
	fmt.Println(result.RowsAffected) // 입려된 레코드의 개수를 반환합니다.

	//r := gin.Default() // defalut settings
	//
	//Print("hello", "hello")
	//Print("hello", 4)
	//
	//Print2("hello", "hello")
	//Print2("hello", 4)
	//// default string
	//r.GET("/default/string", func(c *gin.Context) { // handler function
	//	c.String(http.StatusOK, "Hello world!!!")
	//})
	//
	//// default json
	//r.GET("/defalut/json", func(c *gin.Context) {
	//	c.JSONP(http.StatusOK, gin.H{
	//		"reponse": "Hello world!!!",
	//	})
	//})
	//
	//// HTTP 파라미터 받기
	//r.GET("/:name", func(c *gin.Context) {
	//	var val = c.Param("name")
	//	c.JSON(http.StatusOK, gin.H{
	//		"value": val,
	//	})
	//})
	//
	//// Body로 들어오는 POST 요청처리
	//r.POST("/add", func(c *gin.Context) {
	//	//var req := &Bind{}
	//	var data TestModel
	//	if err := c.ShouldBind(&data); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": fmt.Sprintf("%v", err),
	//		})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{
	//			"data": data,
	//		})
	//	}
	//})
	//
	//printArray()
	//r.Run("localhost:8080") // api를 호스트할 url과 포트번호
}

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
