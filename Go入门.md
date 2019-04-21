# Go入门

从[Go 语言官网](https://golang.google.cn/)下载对应的二进制包，也就是可以拿来即用的安装包。

在这个过程中，我们还需要配置 3 个环境变量，也就是 GOROOT、GOPATH 和 GOBIN。这里我可以简单介绍一下。

- GOROOT：Go 语言安装根目录的路径，也就是 GO 语言的安装路径。
- GOPATH：若干工作区目录的路径。是我们自己定义的工作空间。
- GOBIN：GO 程序生成的可执行文件（executable file）的路径。



.a文件是编译过程中生成的，每个package都会生成对应的.a文件，Go在编译的时候先判断package的源码是否有改动，如果没有的话，就不再重新编译.a文件，这样可以加快速度。

## 一、Go基本语法

### 1.1结构体

Go语言是类似c语言的一门语言，不像java这样是纯面向对象的，所有方法都是在类里面的。针对类似java中的**CLASS**.在go中有类似的概念是叫做结构体。

#### 1.1.1 结构体的使用

```go
//申明一个结构体
type People struct{
	Name string
	age int
}
func main(){
  //申明一个People类型的变量p
	var p People
  p.Name="初始值"
  fmt.Println(p.Name)
  p.setNameByPoint()
  fmt.Println("指针传递之后的值：",p.Name)
  p.setNameByValue()
  fmt.Println("值传递之后的值：",p.Name)
}
//申明一个People类型的方法，是指针传递
func (p *People)setNameByPoint(){
  p.Name="指针传递"
}
//申明一个People类型的方法，是值传递
func (p People)setNameByValue(){
   p.Name="值传递"
}
func setNameByValue(p *People){
   p.Name="函数的指针传递"
}
```

#### 1.1.2 结构体中定义

在结构体中定义的属性名，首字母大写表示可导出，小写表示不可导出。结构体名称也是相同定义，首字母大写是可导出，小写就是不可导出。

### 1.2方法和函数

在go语言中是存在方法和函数的区别的，简单的来说就是

- **方法:**是指定了结构体类型的，只有对应的结构体变量才可以调用，比如上面的`func (p People)setNameByValue(){}`就是方法。只有People类型的变量才可以调用

- **函数：** 函数名前是没有定义结构体类型的，可以在任意地方比调用(**方法名大写的情况**)

  > 1. 函数和方法也是存在是否能被外部包调用的情况的，与结构体类型，方法名或者函数名是大写的就是可被外部包调用的，小写就是不能被外部调用
  > 2. 方法和函数都是可以定义返回值类型的，在Go语言中，可以一次返回多个变量,如下

```go
func main(){
  //1、可以直接在打印中传入函数
  fmt.Println(returnSingleValue()) //单个值
	fmt.Println(returnDounleValue()) //字符串，12
  //2、也可以获取变量，通过:=来将返回的值赋值给新变量，go使用:=这种方式来，不需要指定类型,默认生成的变量就是返回值的类型
  rsv := returnSingleValue
  fmt.Println(rsv) //单个值
  rdv1,rdv2 := returnDounleValue()
  fmt.Println(rdv1，rdv2) //字符串，12
}

func returnSingleValue() string{
  return "单个值"
}
func returnDounleValue()(string,int){
  return "字符串",12
}
```

> 简式变量
> 使用 := 定义的变量，如果新变量rsv与那个同名已定义变量 (这里就是那个全局变rsv)不在一个作用域中时，那么Go 语言会新定义这个变量rsv，遮盖住全局变量rsv。刚开始很容易在此犯错而茫然，解决方法是局部变量尽量不同名。

例子：

```go
var (
    rsv int = 99
)
func main() {
    rsv := "string"
    fmt.Println("main函数中：", rsv) //输出是string
}
```



## 二、数据类型

### 2.1基本数据类型

布尔型：bool

数字类型：

- 整形 int8 int16 int32 int 64和相应的无符号整形的uint   

- 浮点型 float32、float64

- 其他数字类型：byte类型类似uint8，rune类似int32，uintptr无符号整形，用于存放指针

  > 一个`rune`类型的值即可表示一个Unicode字符。`rune`类型的值需要由单引号“'”包裹。例如，`'A'`或`'啊'`。rune 英文占一个字节，中文占三个字节，

字符串类型：string

> string底层是用byte数组存的,，在字符串上调用 len 函数，取得的是字符串包含的 byte 的个数
>
> 例如 s:="你好"  fmt.Println(len(s))  输出结果应该是6，因为中文字符是用3个字节存的.如果有中文，按下标是访问不到，如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数

**注意点！！**

```go
package main

import "fmt"
func main() {
  s := "go你好" // UTF-8
	fmt.Println(len(s))    //结果：8
 	fmt.Println(s[0:4])  //中文乱码，因为截取的索引下标2,3的组合无法被转换
  fmt.Println(s[0:5] 
	//
	st := []rune(s)
	fmt.Println(len(st))  //结果：4
	fmt.Println(string(st[0:4]))  
}
```

派生类型：包括

```
(a) 指针类型（Pointer）
(b) 数组类型 [5]int
(c) 结构类型(struct)
(d) Channel 类型
(e) 函数类型 (func)
(f) 切片类型 (slice)[]int
(g) 接口类型（interface）
(h) Map 类型 map/sync.map
```

### 2.2 空值

go里每种类型都有属于自己的空值，只要申明了此变量，则此变量就已经有值

例如：

int空值是0，

string空值是””而不是null或者nil（很多人不理解，这与其他语言确实有一点区别），

Slice空值是长度为0的Slice而不是nil，

map，channel,interface空值是nil，

struct空值是一个“所有成员都是空值”的空Struct而不是nil。

```go
type People struct{
  Name     string
  Age      int
  CarsName []String
}

func main(){
  //默认会初始化
  var p People
  fmt.Printf("p name is %s and Age is %d and Cars length %d",p.Name,p.Age,len(p.CarsName))
}
```

## 三、 包的概念

Go语言使用包（package）的概念来组织管理代码，包是结构化代码的一种方式。和其他语言如JAVA类似，Go语言中包的主要作用是把功能相似或相关的代码组织在同一个包中，以方便查找和使用。在Go语言中，每个.go文件都必须归属于某一个包，每个文件都可有init()函数。包名在源文件中第一行通过关键字package指定，包名要小写。如下所示：

```go
package fmt
```

- 每个目录下面可以有多个.go文件，这些文件只能属于同一个包，否则编译时会报错。同一个包下的不同.go文件相互之间可以直接引用变量和函数，所以这些文件中定义的全局变量和函数不能重名。

- Go语言的可执行应用程序必须有main包，而且在main包中必须且只能有一个main()函数，main()函数是应用程序运行开始入口。在main包中也可以使用init()函数。

- Go语言**不强制**要求包的名称和文件所在目录名称相同，但是这**两者最好保持相同**，否则很容易引起歧义。因为导入包的时候，会使用目录名作为包的路径，而在代码中使用时，却要**使用包的名称**。

  > 比如nsq的我们导入的路径如下：
  >
  >    import “github.com/bitly/go-nsq”
  >
  > 但在使用其提供的export functions时，却用nsq做前缀包名：
  >
  >    q, _ := nsq.NewConsumer("write_test", "ch", config)

  一个Go程序通过import关键字将一组包链接在一起。import其实是导入目录，而不是定义的包名称，实际应用中我们一般都会保持一致。

> 一般会从三个路径去查找，
>
> - 一个是GOROOT/src目录下，这里面存放的是Go的官方库
> - 一个是GOPATH/src目录下,这里面存放的是开源代码库,go get 下来的代码都在这里
> - 还有自己项目中的包

![image-20190420153146725](/Users/zhaohq/Library/Application Support/typora-user-images/image-20190420153146725.png)

## 四 、接口

Go 语言中的接口概念和Java概念类似，都是定义了一组方法集合，但是这些方法集合仅仅只是被定义，它们没有在接口中实现。**接口(interface)类型**是Go语言的一种数据类型(可以理解成java中的object类型)。而因为所有的类型包括自定义类型都实现了空接口interface{}，所以空接口interface{}可以被当做任意类型的数值。

例如：

```go
//定义一组接口
type Reader interface {
    Read(p []byte) (n int, err error)
}
type Writer interface {
    Write(p []byte) (n int, err error)
}
//定义i为借口类型
var i interface{} = 99 // i可以是任何类型
i = 44.09
i = "All"  // i 可接受任意类型的赋值
```



**Go语言借口的特点：**

1. Go语言中接口是不需要显示实现的，这与java中实现接口需要使用implements不同，在go语言中只要对应的类型中所包含的方法有实现对应的接口中的方法，那么这个类型就是实现了这个接口。
2. Go中的接口类型的空值是nil，但是要注意，**在底层，interface作为两个成员来实现，一个类型和一个值**

```go
type File struct { // ...
}
//File类型就实现了Read和Write两个接口,
func (f *File) Read(buf []byte) (n int, err error){
	return 10,nil
}
func (f *File) Write(buf []byte) (n int, err error){
	 return 10,nil
}
func main(){
    //将interface的类型设置为*interface{}
     var val interface{} = (map[string]string)(nil)
    if val==nil{
        fmt.Println("val is nil")  
    } else {  
        fmt.Println("val is not nil")  
    }  
}
```

