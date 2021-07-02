package main


/*
func main() {
	var sum int
	sum = SumTest(2, 3)
	fmt.Printf("sum : %d, G_sum : %d\n", sum, GSum)
}
 */

/*
// 函数

// 方法
type Circle struct {
	radius float64
}

func main() {
	var c Circle
	c.radius = 10.00
	fmt.Println("圆的面积 = ", c.getArea())
}

// 该method属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
 */

/*
// 回调函数
type cb func(int)

func main() {
	testCallBack(1, callback)
	testCallBack(2, func(x int) {
		fmt.Printf("我是回调, x : %d\n", x)
		//return x
	})
}

func testCallBack(x int, fun cb) {
	fun(x)
}

func callback(x int) {
	fmt.Printf("我是回调，x : %d\n", x)
	//return x
}
 */

/*
func main() {
	getSquare := func (x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquare(16))
}
 */

/*
// go函数可以返回多个值
func swap(x, y string) (string,string) {
	return y, x
}

func main() {
	a := "warden"
	b := "even"
	fmt.Println(swap(a, b)) // even warden
}
 */

/*
func getMax(a, b int) int {
	var res int
	if a > b {
		res = a
	} else {
		res = b
	}
	return res
}

func main() {
	x := 10
	y := 20
	res := getMax(x, y)
	fmt.Println("较大的数为", res)
}
 */

/*
func main() {
	// 获取当前系统时间
	fmt.Println("The time is", time.Now())
}
 */


// 循环与分支
/*
func main() {
	var score int = 0
	var evaluate string = "A"
	fmt.Printf("欢迎进入成绩查询系统\n")
ZHU:
	for true {
		var choose int = 0
		fmt.Println("1.进入成绩查询 2.退出程序")
		fmt.Printf("请输入序号选择: ")
		fmt.Scanln(&choose)
		fmt.Printf("\n")
		if choose == 1 {
			goto CHA
		}
		if choose == 2 {
			os.Exit(1)
		}

	}

CHA:
	for true {
		fmt.Printf("请输入一个学生的成绩: ")
		fmt.Scanln(&score)

		switch {
		case score >= 90:
			evaluate = "A"

		case score >= 80 && score < 90:
			evaluate = "B"

		case score >= 60 && score < 80:
			evaluate = "C"

		default:
			evaluate = "D"
		}

		var c string = strconv.Itoa(score)
		switch {
		case evaluate == "A":
			fmt.Printf("你考了%s分,评价为%s,成绩优秀\n", c, evaluate)
		case evaluate == "B" || evaluate == "C":
			fmt.Printf("你考了%s分,评价为%s,成绩良好\n", c, evaluate)
		case evaluate == "D":
			fmt.Printf("你考了%s分,评价为%s,成绩不合格\n", c, evaluate)
		}
		fmt.Printf("\n")
		goto ZHU
	}
	//fmt.Println(score)
}
 */

/*
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i * j)
		}
		fmt.Println()
	}
}
*/

/*
func main() {
	var n, i int
	n = 1
A:
	for n < 100 {
		n++
		for i = 2; i < n; i++ {
			if n % i == 0 {
				goto A
			}
		}
		//fmt.Printf("%d是素数\n", i)
		fmt.Println(i, "是素数")
	}
}
*/

/*
func main() {
	strings := []string {"wardell", "even"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int {1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位x的值 = %d\n", i , x)
	}
	//sum := 1
	//for ; sum <= 10; {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//for sum <= 10 {
	//	sum += sum
	//}
	//fmt.Println(sum)
}
*/

/*
func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 : %T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
*/

/*
func isPrime(n int) bool {
	if n < 2 {
		return false
    }
    for i := 2; i <= n / 2; i++ {
    	if n % i == 0 {
    		return false
		}
	}
	return true
}

func main() {
	for num := 1; num <= 100; num++ {
		if isPrime(num) {
			fmt.Print(num, " ")
		}
	}
}
*/

/*
func main() {
	var a int = 10
	var ptr *int = &a
	fmt.Println(ptr)
	fmt.Println(&a)
	fmt.Println(*ptr)
	a = 20
	fmt.Println(*ptr)
	a++
	fmt.Println(a)
}
*/

/*
const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a) // 16
	// 字符串在go里是个结构，包含指向底层数组的指针和长度，这两部分每部分都是 8 个字节，所以字符串类型大小为 16 个字节。
)

//iota : 可被编译器修改的常量
func main() {
	const (
		a = iota
		b
		c
		d = "xjw"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	//println(a, b, c)
}
*/

/*
var x, y int

var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func main() {
	var a string = "abc"
	fmt.Println("Hello world", a)
	//g, h := 123, "hello"
	//println(x, y, a, b, c, d, e, f, g, h)
	//var name string
	//fmt.Printf("%q\n", name)
	//
	//var a, b int
	//var c bool
	//fmt.Println(a, b)
	//fmt.Println(c)
}
*/

/*
//连接数据库并简单查询
import (
	//go get -u github.com/go-sql-driver/mysql
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接本地mytest数据库
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mytest?charset=utf8mb4")
	if err != nil {
		fmt.Println("数据库连接失败：" + err.Error())
		return
	}

	//查询bookinfo表
	res, err := db.Query("select * from bookinfo")
	if err != nil {
		fmt.Println("查询错误：" + err.Error())
		return
	}

	//打印列名
	fmt.Println(res.Columns())

	//循环扫描数据表取出数据
	for res.Next() {
		var book_id int
		var book_name string
		var price float32
		var public_date string
		var store int

		//扫描行并把扫描到的数据赋值
		res.Scan(&book_id, &book_name, &price, &public_date, &store)
		fmt.Println(book_id, book_name, price, public_date, store)

	}
}
*/

/*
func main() {
	str := "这里是 www\n.baidu\n .com"
	fmt.Println("原字符串如下：")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去吹换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("去除空格和换行符的字符串如下：")
	fmt.Println(str)
}
*/

/*
var isActive bool //全局变量声明
var enabled, disabled = true, false //忽略类型的声明

func test() {
	var available bool //一般声明
	vaild := false //简短声明
	available = true //赋值
	if vaild || available {
		fmt.Println("bool!!!")
	}
}

func main() {
	var a float32 = 1.5
	var b int = 2
	fmt.Println(a, b)
	test()
}
*/

/*
func main() {

	fmt.Print("a", "b", 1, 2, 3, "c", "d", "\n")
	fmt.Println("a", "b", 1, 2, 3, "c", "d")
	fmt.Printf("ab %d %d %d cd\n", 1, 2, 3)
	//ab1 2 3cd
	//a b 1 2 3 c d
	//ab 1 2 3 cd

	if err := percent(30, 70, 90, 160); err != nil {
		fmt.Println(err)
	}
	//30%
	//70%
	//90%
	//数值160超出范围（100）
}

func percent(i ...int) error {
	for _, n := range i {
		if n > 100 {
			return fmt.Errorf("数值%d超出范围（100）", n)
		}
		fmt.Print(n, "%\n")
	}
	return nil
}
*/
