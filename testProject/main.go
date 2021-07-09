package main

import (
	"context"
	"fmt"
	"time"
)

// 并发
// Context
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

/*
// channel
func sum(arr []int, ch chan int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	ch <- sum
}

func fib(n int, ch chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x + y
	}
	close(ch) // 关闭通道
}

func main() {
	ch := make(chan int, 10)
	go fib(cap(ch), ch)

	// range 函数遍历每个从通道接收到的数据，因为 ch 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 ch 通道不关闭，那么 range 函数就不会结束
	// 从而在接收第 11 个数据的时候就阻塞了
	for i := range ch {
		fmt.Println(i)
	}

	//ch := make(chan int, 2) // 创建通道，并定义缓冲区，大小为2
	//
	//// ch是带缓冲区的通道，我们可以同时发送两个数据
	//// 而不用立刻去同步读数据
	//ch <- 1
	//ch <- 2
	//
	//// 获取这两个数字
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)

	//arr := []int{3, 5, -2, 1, 0, 8} // 定义切片
	//
	//ch := make(chan int) // 创建通道
	//go sum(arr[len(arr) / 2 :], ch)
	//go sum(arr[: len(arr) / 2], ch)
	//
	//x, y := <- ch, <- ch // 从通道ch中接收数据
	//fmt.Println(x, y, x + y)
}
 */

/*
// routine
func startRoutine(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}

func main() {
	go startRoutine("hello")
	startRoutine("world")
}
 */

/*
// 错误处理
// defer panic recover
func main(){
	defer func(){
		if err := recover() ; err != nil {
			fmt.Println(err)
		}
	}()
	defer func(){
		panic("three")
	}()
	defer func(){
		panic("two")
	}()
	panic("one")
}
 */

/* 错误处理案例
// 定义一个DivideError结构
type DivideError struct {
	dividend int // 被除数
	divisor int // 除数
}

// 实现`error`接口
func (de DivideError) Error() string {
	errorMsg := `
    Cannot proceed, the divisor is zero.
    dividend: %d
    divisor: 0
`
    return fmt.Sprintf(errorMsg, de.dividend)
}

// 定义`int`类型除法运算的函数
func Divide(dd int, dr int) (result int, errorMsg string) {
	if dr == 0 { // 除数为0
		de := DivideError {
			dividend: dd,
			divisor: dr,
		}
		errorMsg = de.Error()
		return
	} else {
		return dd / dr, "" // 正常返回，没有错误信息
	}
}

func main() {
	// 正常情况
	if res, errMsg := Divide(100, 10); errMsg == "" {
		fmt.Println("100 / 10 =", res)
	}

	// 当除数为0时返回错误信息
	if _, errMsg := Divide(100, 0); errMsg != "" {
		fmt.Print("errorMsg is :", errMsg)
	}
}
 */

/*
// 接口

// 将接口作为参数
type Phone interface {
	call() string
}

type Android struct {
	brand string
}

type IPhone struct {
	version string
}

func (android Android) call() string {
	return "I am Android " + android.brand
}

func (iPone IPhone) call() string {
	return "I am IPone " + iPone.version
}

func printCall(p Phone) {
	fmt.Println(p.call() + ", I can call you!")
}

func main() {
	var vivo = Android{"Vivo"}
	var huawei = Android{"Huawei"}

	i7 := IPhone{"7 Plus"}
	ix := IPhone{"X"}

	printCall(vivo)
	printCall(huawei)
	printCall(i7)
	printCall(ix)
}
 */

/*
// 接口案例
// 定义接口
type Phone interface {
	call()
	call2()
}

type Phone1 struct {
	id           int
	name         string
	categoryId   int
	categoryName string
}

// 第一个类的第一个回调函数
func (test Phone1) call() {
	fmt.Println("这是第一个类的第一个接口回调函数call", Phone1{id: 1, name: "浅笑"})
}

// 第一个类的第二个回调函数
func (test Phone1) call2() {
	fmt.Println("这是第一个类的第二个接口回调函数call2", Phone1{id: 1, name: "浅笑", categoryId: 4, categoryName: "分类名称"})
}



//第二个结构体的数据类型
type Phone2 struct {
	memberId       int
	memberBalance  float32
	memberSex      bool
	memberNickname string
}

// 第二个类的第一个回调函数
func (test2 Phone2) call() {
	fmt.Println("这是第二个类的第一个接口回调函数call", Phone2{memberId: 22, memberBalance: 15.23, memberSex: false, memberNickname: "浅笑18"})
}


// 第二个类的第二个回调函数
func (test2 Phone2) call2() {
	fmt.Println("这是第二个类的第二个接口回调函数call2", Phone2{memberId: 44, memberBalance: 100, memberSex: true, memberNickname: "陈超"})
}

// 开始运行
func main() {
	var phone Phone

	// 先实例化第一个接口
	phone = new(Phone1)
	phone.call()
	phone.call2()

	// 实例化第二个接口
	phone = new(Phone2)
	phone.call()
	phone.call2()
}
 */

/*
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (np NokiaPhone) call() {
	fmt.Println("i am Nokia, I can call you!")
}

type IPhone struct {
}

func (ip IPhone) call() {
	fmt.Println("I am IPhone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
 */

/*
// go不支持隐式类型转换
func main() {
	var a uint64
	var b uint32 = 3
	// a = b // 会报错

	// 使用类型转换
	a = uint64(b)
	fmt.Printf("a : %d\n", a)
}
 */

// 递归
/*
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n - 1) + fib(n - 2)
}

func main() {
	var n int = 10
	for i := 0; i < n; i++ {
		fmt.Printf("%d\t", fib(i))
	}
}
 */

/*
func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}

func main() {
	var a int
	fmt.Print("请输入你要求的阶乘：")
	fmt.Scan(&a)
	res := factorial(uint64(a))
	fmt.Printf("%d的阶乘为%d\n", a, res)
}
*/

/*
// map

// delete : 参数为map和对应的key
func main() {
	infoMap := map[string]int{"小明" : 18, "小红" : 16, "小李" : 20}

	fmt.Println("原始信息：")
	// 打印信息
	for name, age := range infoMap {
		fmt.Printf("%s的年龄是%d\n", name, age)
	}

	delete(infoMap, "小李")
	fmt.Printf("小李的信息已被删除\n")

	fmt.Printf("删除信息后：\n")
	for name, age := range infoMap {
		fmt.Printf("%s的年龄是%d\n", name, age)
	}
}
 */

/*
func main() {
	// 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
	// var countryCapitalMap map[string]string // 创建集合
	// countryCapitalMap = make(map[string]string)

	// 或者
	countryCapitalMap := make(map[string]string)

	// map插入key - value对,各个国家对应的首都
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	// 使用键输出地图值
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap [ "American" ]
	//fmt.Println(capital) // ""
	//fmt.Println(ok) // false
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}
 */

/*
// range
// 通过range获取参数列表
func main() {
	fmt.Println(len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
 */

/*
func main() {
	// 使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 4, 6}
	var sum int
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum ==", sum)

	// range也可以用在map的键值对上
	kvs := map[string]string{"a" : "apple", "b" : "banana"}
	for k, v := range kvs {
		fmt.Printf("%s : %s\n", k, v)
	}

	// range也可以用来枚举Unicode字符串。第一个参数是字符的索引
	// 第二个是字符（Unicode的值）本身
	for i, c := range "abcABC" {
		fmt.Println(i, c)
	}
}
 */

/*
// 切片
// slice 的底层是数组指针，所以 slice a 和 s 指向的是同一个底层数组
// 所以当修改 s[0] 时，a 也会被修改
func main() {
	s := []int{1, 2, 3} // len=3, cap=3
	a := s
	s[0] = 888
	s = append(s, 4)
	fmt.Println(a, len(a), cap(a)) // 输出：[888 2 3] 3 3
	fmt.Println(s, len(s), cap(s)) // 输出：[888 2 3 4] 4 6
}
 */

/*
// 合并多个数组
func main() {
	var arr1 = []int{1,2,3}
	var arr2 = []int{4,5,6}
	var arr3 = []int{7,8,9}
	var s1 = append(append(arr1, arr2...), arr3...)
	fmt.Printf("s1: %v\n", s1)
}
 */

/*
// append()和copy()函数
func main() {
	var numbers []int
	printSlice(numbers)

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)

	// 同时添加多个元素
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	// 创建切片 numbers1 是之前切片的两倍容量
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	// 拷贝 numbers 的内容到 numbers1
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int){
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
*/

/*
func main() {
	// 创建切片
	//var arr = [3]int{1, 2, 4}
	//numbers := arr[0:2]
	//var numbers = make([]int, 3, 5)
	//var numbers []int
	//if numbers == nil {
	//	fmt.Println("切片为空！")
	//}
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	// 打印原始切片
	fmt.Println("numbers ==", numbers)

	// 打印子切片从索引1（包含）到索引4（不包含）
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	// 默认下限为0
	fmt.Println("numbers[:3] ==", numbers[:3])

	// 默认上限为len(s)
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	// 打印子切片从索引0到索引2
	numbers2 := numbers[:2]
	printSlice(numbers2)

	// 打印子切片从索引2到索引5
	numbers3 := numbers[2:5]
	printSlice(numbers3)

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
 */

/*
// 结构体

//结构体中属性的首字母大小写问题
//首字母大写相当于 public
//首字母小写相当于 private
//注意: 这个 public 和 private 是相对于包（go 文件首行的 package 后面跟的包名）来说的。
//敲黑板，划重点
//当要将结构体对象转换为 JSON 时，对象中的属性首字母必须是大写，才能正常转换为 JSON
type Person struct {
	Name string // 大写
	Age int
}

func main() {
	p1 := Person{"小明", 18}
	if res, err := json.Marshal(&p1); err == nil {
		fmt.Println(string(res))
	}
}
 */

/*
type Books struct {
	title  string
	author string
	bookId int
}

func main() {
	var Book1 Books // 声明 Book1 为 Books 类型
	var Book2 Books // 声明 Book2 为 Books 类型

	// book 1 描述
	Book1.title = "Go 语言"
	Book1.author = "wardell"
	Book1.bookId = 6495407

	// book 2 描述
	Book2.title = "Python 教程"
	Book2.author = "even"
	Book2.bookId = 6495700

	// 打印 Book1 信息
	printBook(&Book1)

	// 打印 Book2 信息
	printBook(&Book2)
}

// 结构体传参
func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title) // 通过"."操作符访问结构体成员
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book book_id : %d\n", book.bookId)
}
*/

/*
// 指针

// 指针作为参数
func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a int = 10
	var b int = 20

	fmt.Printf("交换前 ： a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("交换后 ： a = %d, b = %d\n", a, b)
}
 */

/*
// 指针数组 VS 数组指针
const max = 3

func main() {
	var arr = [3]int{1, 2, 3}
	var parr [3]*int // 指针数组
	var ptr *[3]int = &arr // 数组指针

	for i := range arr {
		parr[i] = &arr[i]
	}

	for i := range arr {
		fmt.Printf("arr[%d] = %d, 地址为 ： %p %p %p\n", i, arr[i], &arr[i], parr[i], &(*ptr)[i])
	}
}
 */

/*
func main() {
	number := [max]int{5, 6, 7}
	var ptr [max]*int // 指针数组
	// 将number数组的值的地址赋给ptr
	// 错误代码，x为临时变量，无法得到正确地址
	//for i, x := range &number {
	//	ptr[i] = &x
	//}
	// 正确写法
	//for i := 0; i < max; i++ {
	//	ptr[i] = &number[i]
	//}
	// 或者
	for i := range number {
		ptr[i] = &number[i]
	}
	for i, x := range ptr {
		fmt.Printf("指针数组：索引:%d 值:%d 值的内存地址:%d\n", i, *x, x)
	}
}
 */

/*
// 数组
func main() {
	// 二维数组
	var value = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	// 遍历二维数组的其他方法，使用 range
	// 其实，这里的 i, j 表示行游标和列游标
	// v2 就是具体的每一个元素
	// v  就是每一行的所有元素
	for i, v := range value {
		for j, v2 := range v {
			fmt.Printf("value[%v][%v]=%v \t ", i, j, v2)
		}
		fmt.Print(v)
		fmt.Println()
	}
}
 */

/*
//使用数组打印杨辉三角，根据上一行的内容来获取下一行内容并打印
func printYang(inArr []int) []int {
	var outArr []int
	lenArr := len(inArr)
	outArr = append(outArr, 1)
	if lenArr == 0 {
		return outArr
	}
	for i := 0; i < lenArr - 1; i++ {
		outArr = append(outArr, inArr[i] + inArr[i + 1])
	}
	outArr = append(outArr, 1)
	return outArr
}

func main() {
	var yang []int

	for i := 0; i < 10; i++ {
		yang = printYang(yang)
		fmt.Println(yang)
	}

}
 */

/*
// Go 语言的数组是值，其长度是其类型的一部分，作为函数参数时，是值传递，函数中的修改对调用者不可见
func change1(nums [3]int) {
	nums[0] = 4
}
// 传递进来数组的内存地址，然后定义指针变量指向该地址，则会改变数组的值
func change2(nums *[3]int) {
	nums[0] = 5
}
// Go 语言中对数组的处理，一般采用 切片 的方式，切片包含对底层数组内容的引用
//作为函数参数时，类似于指针传递，函数中的修改对调用者可见
func change3(nums []int) {
	nums[0] = 6
}
func main() {
	var nums1 = [3]int{1, 2, 3}
	var nums2 = []int{1, 2, 3}
	change1(nums1)
	fmt.Println(nums1)  //  [1 2 3]
	change2(&nums1)
	fmt.Println(nums1)  //  [5 2 3]
	change3(nums2)
	fmt.Println(nums2)  //  [6 2 3]
}
 */

/*
func getAverage(score [6]int, size int) float32 {
	var avg float32
	var sum, i int
	for i = 0; i < size; i++ {
		sum += score[i]
	}

	avg = float32(sum) / float32(size)
	return avg
}

func main() {
	// 数组作为参数传参
	var res float32
	var score = [6]int{80, 95, 76, 88, 53, 90}
	res = getAverage(score, 6)

	fmt.Printf("平均成绩为 ： %f\n", res)

	//arr := [4][2]int{{0, 1}, {6, 3}, {3, 4}, {7, 8}}
	//var i, j int
	//for i = 0; i < 4; i++ {
	//	fmt.Printf("第%d行 ： ", i)
	//	for j = 0; j < 2 ;j++ {
	//		fmt.Printf("%d ", arr[i][j])
	//	}
	//	fmt.Println()
	//}

	//// 创建空的二维数组
	//var animals [][]string
	//
	//// 创建三个一维数组，各数组长度不同
	//row1 := []string{"fish", "shark", "eel"}
	//row2 := []string{"bird"}
	//row3 := []string{"lizard", "salamander"}
	//
	//// 使用append()函数将一维数组添加到二维数组中
	//animals = append(animals, row1)
	//animals = append(animals, row2)
	//animals = append(animals, row3)
	//
	////循环输出
	//for i := range animals {
	//	fmt.Printf("Row : %d\n", i)
	//	fmt.Println(animals[i])
	//}


	//var arr = [5] float32 {1.2, 2.3, 4.5, 10.0, 2.2}
	//balance := [5] int {1, 2, 3, 4}
	//
	//for i := 0; i < 5; i++ {
	//	fmt.Printf("arr[%d] = %f\n", i, arr[i])
	//	fmt.Printf("balance[%d] = %d\n", i, balance[i])
	//}
}
 */


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
