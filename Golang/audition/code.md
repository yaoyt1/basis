1. 从切片中删除一个元素，下面的算法实现正确的是（）  
A.
```go
func (s *Slice)Remove(value interface{}) error {
        for i, v := range *s {
            if isEqual(value, v) {
                if i== len(*s) - 1 {
                    *s = (*s)[:i]
                }else {
                    *s = append((*s)[:i],(*s)[i + 2:]...)
                }
                return nil
            }
        }
        return ERR_ELEM_NT_EXIST
}
```
B.
```go
func (s *Slice)Remove(value interface{}) error {
        for i, v := range *s {
            if isEqual(value, v) {
                *s = append((*s)[:i],(*s)[i + 1:])
                return nil
            }
        }
        return ERR_ELEM_NT_EXIST
}
```
C. 
```go
func (s *Slice)Remove(value interface{}) error {
        for i, v := range *s {
            if isEqual(value, v) {
                delete(*s, v)
                return nil
            }
        }
        return ERR_ELEM_NT_EXIST
}
```
D.
```go
func (s *Slice)Remove(value interface{}) error {
        for i, v := range *s {
            if isEqual(value, v) {
                *s = append((*s)[:i],(*s)[i + 1:]...)
                return nil
            }
        }
        return ERR_ELEM_NT_EXIST
}
```

参考答案： D

2. 对于局部变量整型切片x的赋值，下面定义正确的是（）  
A. 
```go
x := []int{
1, 2, 3,
4, 5, 6,
}
```
B. 
```go
x := []int{
        1, 2, 3,
        4, 5, 6
}
```
C.
```go
x := []int{
        1, 2, 3,
        4, 5, 6}
```
D.
```go
x := []int{1, 2, 3, 4, 5, 6,}
```

参考答案：ACD

3. 如果Add函数的调用代码为:
```go
func main() {
        var a Integer = 1
        var b Integer = 2
        var i interface{} = &a
        sum := i.(*Integer).Add(b)
        fmt.Println(sum)
}
```
则Add函数定义正确的是（）: 中级知识点  
A. 
```go
type Integer int
func (a Integer) Add(b Integer) Integer {
        return a + b
}
```
B.
```go
type Integer int
func (a Integer) Add(b *Integer) Integer {
        return a + *b
}
```
C.
```go
type Integer int
func (a *Integer) Add(b Integer) Integer {
        return *a + b
}
```
D.
```go
type Integer int
func (a *Integer) Add(b *Integer) Integer {
        return *a + *b
}
```

参考答案：AC

4. 关于slice或map操作，下面正确的是（） ：slice,map知识点考察  
A.
```go
var s []int
s = append(s,1)
```
B.
```go
var m map[string]int
m["one"] = 1 
```
C.
```go
var s []int
s = make([]int, 0)
s = append(s,1)
```
D.
```go
var m map[string]int
m = make(map[string]int)
m["one"] = 1 
```

参考答案：ACD

5. 下面的程序的运行结果是__________
```go
for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
}
```

参考答案：4 3 2 1 0

6. 下面的程序的运行结果是__________
```go
func main() {
        x := 1
        {
            x := 2
            fmt.Print(x)
        }
        fmt.Println(x)
}
```

参考答案： 21

7. 下面的程序的运行结果是__________
```go
func main() {
        strs := []string{"one", "two", "three"}

        for _, s := range strs {
            go func() {
                time.Sleep(1 * time.Second)
                fmt.Printf("%s ", s)
            }()
        }
        time.Sleep(3 * time.Second)
}
```

参考答案：three three three

8. 下面的程序的运行结果是__________
```go
func main() {
        strs := []string{"one", "two", "three"}

        for _, s := range strs {
            go func(s string) {
                time.Sleep(1 * time.Second)
                fmt.Printf("%s ", s)
            }(s)
        }
        time.Sleep(3 * time.Second)
}
```

参考答案：one three two (随机打印)

9.下面的程序的运行结果是__________
```go
func main() {
    x := []string{"a", "b", "c"}
    for v := range x {
    	fmt.Print(v)
    }
}
```

参考答案：012

10.下面的程序的运行结果是__________
```go
func main() {  
        x := []string{"a", "b", "c"}
        for _, v := range x {
            fmt.Print(v)
        }
}
```
参考答案：abc

11. 下面的程序的运行结果是__________
```go
type Slice []int
func NewSlice() Slice {
         return make(Slice, 0)
}
func (s* Slice) Add(elem int) *Slice {
         *s = append(*s, elem)
         fmt.Print(elem)
         return s
}
func main() {  
         s := NewSlice()
         defer s.Add(1).Add(2)
         s.Add(3)
}
```

参考答案：132

12. 下面这段代码输出的内容
```go
func main() {
    defer_call()
}

func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}
```

参考答案：
打印后  
打印中  
打印前  
panic: 触发异常  

13. 下面这段代码输出什么，说明原因。
```go
func main() {

	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		m[key] = &val
	}

	for k,v := range m {
		fmt.Println(k,"->",*v)
	}
}
```

参考答案：  
0 -> 3  
1 -> 3  
2 -> 3  
3 -> 3  

14. 下面这段代码输出什么，说明原因。
```go
type Test struct {
	name string
}

func (this *Test) Point(){
	fmt.Println(this.name)
}

func main() {
	ts := []Test{
		{"a"},
		{"b"},
		{"c"},
	}

	for _,t := range ts {
		//fmt.Println(reflect.TypeOf(t))
		defer t.Point()
	}
	
}
```

参考答案：  
c  
c  
c  

15. 下面这段代码能否通过编译，如果可以，输出什么？
```go
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```

参考答案： 编译不过  
解析： append()第二个参数不能直接使用slice 可以改成append(s1,s2...)或者append(s1,4,5)

16. 下面这段代码能否通过编译，如果可以，输出什么？
```go
var(
	size := 1024
	max_size = size*2
)

func main() {
	fmt.Println(size,max_size)
}
```

参考答案： 编译不过
解析： 简短声明只能使用在函数内部   

17. 下面这段代码能否通过编译？如果通过，输出什么？
```go
type MyInt1 int
type MyInt2 = int

func main() {
    var i int =0
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1,i2)
}
```

参考答案:编译不过
解析： golang是强类型语言， 需要使用类型转化int()

18. 下面这段代码能否编译通过？如果可以，输出什么？
```go
const (
	x = iota
	_
	y
	z = "zz"
	k 
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
```

参考答案： 02zzzz5

19.下面这段代码输出什么以及原因？
```go
func hello() []string {  
    return nil
}

func main() {  
    h := hello
    if h == nil {
        fmt.Println("nil")
    } else {
        fmt.Println("not nil")
    }
}
```
A. nil
B. not nil
C. error

参考答案： B

20. 

