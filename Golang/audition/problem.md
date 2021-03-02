1. 关于接口和类的说法，下面说法正确的是（）  <br/>
   A. 一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口  
   B. 实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理  
   C. 类实现接口时，需要导入接口所在的包  
   D. 接口由使用方按自身需求来定义，使用方无需关心是否有其他模块定义过类似的接口

参考答案：ABD

2. 关于协程，下面说法正确是（）<br/>
   A. 协程和线程都可以实现程序的并发执行<br/>
   B. 线程比协程更轻量级<br/>
   C. 协程不存在死锁问题<br/>
   D. 通过channel来进行协程间的通信<br/>

参考答案： AD

3. 关于init函数，下面说法正确的是（）  
   A. 一个包中，可以包含多个init函数  
   B. 程序编译时，先执行导入包的init函数，再执行本包内的init函数  
   C. main包中，不能有init函数  
   D. init函数可以被其他函数调用  

参考答案：AB

4. 对于函数定义  
```go
func add(args ...int) int {
        sum := 0
        for _, arg := range args {
            sum += arg
        }
        return sum
}
```
下面对add函数调用正确的是（）  
A. add(1, 2)  
B. add(1, 3, 7)  
C. add([]int{1, 2})  
D. add([]int{1, 3, 7}...)  

参考答案： ABD

5. 关于类型转化，下面语法正确的是（）  
A. 
```go
   type MyInt int
   var i int = 1
   var j MyInt = i
   ```
B. 
```go
type MyInt int
var i int = 1
var j MyInt = (MyInt)i
```
C. 
```go
type MyInt int
var i int = 1
var j MyInt = MyInt(i)
```
D. 
```go
type MyInt int
var i int = 1
var j MyInt = i.(MyInt)
```

参考答案： C

6. 关于局部变量的初始化，下面正确的使用方式是（）  
   A. var i int = 10  
   B. var i = 10  
   C. i := 10  
   D. i = 10  

参考答案：ABC

7. 下面的程序的运行结果是（）: 考察defer 知识点
```go
func main() {  
        if (true) {
            defer fmt.Printf("1")
        } else {
            defer fmt.Printf("2")
        }
        fmt.Printf("3")
}
```
A. 321  
B. 32  
C. 31  
D. 13  

参考答案：C

8. golang中的引用类型包括（） ：基础知识点考察  
   A. 数组切片  
   B. map  
   C. channel  
   D. interface
   
参考答案：ABCD

9. golang中的指针运算包括（）：指针知识点考察  
   A. 可以对指针进行自增或自减运算  
   B. 可以通过“&”取指针的地址  
   C. 可以通过“*”取指针指向的数据  
   D. 可以对指针进行下标运算  
   

参考答案：BC

10. 下面赋值正确的是（）: 基础知识点考察  
    A. var x = nil  
    B. var x interface{} = nil  
    C. var x string = nil  
    D. var x error = nil  
    

参考答案：BD

11. 关于整型切片的初始化，下面正确的是（）：基础知识点考察  
    A. s := make([]int)  
    B. s := make([]int, 0)  
    C. s := make([]int, 5, 10)  
    D. s := []int{1, 2, 3, 4, 5}  
    

参考答案：BCD

12. 关于接口，下面说法正确的是（）:接口知识点考察  
    A. 只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等价的，可以相互赋值  
    B. 如果接口A的方法列表是接口B的方法列表的子集，那么接口B可以赋值给接口A  
    C. 接口查询是否成功，要在运行期才能够确定  
    D. 接口赋值是否可行，要在运行期才能够确定  
    
参考答案：ABC
 
13. 关于channel，下面语法正确的是（）：channel 知识点考察  
    A. var ch chan int  
    B. ch := make(chan int)  
    C. <- ch  
    D. ch <-  
    
参考答案：ABC

14. 关于无缓冲和有冲突的channel，下面说法正确的是（）:channel 知识点考察  
    A. 无缓冲的channel是默认的缓冲为1的channel  
    B. 无缓冲的channel和有缓冲的channel都是同步的  
    C. 无缓冲的channel和有缓冲的channel都是非同步的  
    D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的  
    
参考答案：D

15. 关于select机制，下面说法正确的是（）  ：select机制知识点考察  
    A. select机制用来处理异步IO问题  
    B. select机制最大的一条限制就是每个case语句里必须是一个IO操作  
    C. golang在语言级别支持select关键字  
    D. select关键字的用法与switch语句非常类似，后面要带判断条件  

参考答案：ABC

16. 