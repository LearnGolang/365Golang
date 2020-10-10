# 《365天深入理解Golang》

## 0x01-项目说明 

这是一份准备用365天深入理解Golang之后送给[自己](https://github.com/0e0w)的礼物。本项目记录这365天内的学习收获。

本人编程零基础，只略懂一些Python。因此本项目部分章节可能很基础，请见谅。

本项目大量的参考借鉴甚至复制了其他类似项目。感谢这些作者们，感谢Gopher。

关于Go的其他资源，参考此项目：[https://github.com/0e0w/LearnGolang](https://github.com/0e0w/LearnGolang)

本项目创建于2020年9月1日，最近的一次更新日期为2020年10月10日。

项目处于未完成阶段。不定期推倒重来，暂时取消更新的最新说明。

~~已经更新至第一章Day017: 函数-Go语言函数~~

本项目暂时计划共九章：

- [第一章：Go语言基础](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%B8%80%E7%AB%A0go%E8%AF%AD%E8%A8%80%E5%9F%BA%E7%A1%80)
- [第二章：Go语言进阶](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%BA%8C%E7%AB%A0go%E8%AF%AD%E8%A8%80%E8%BF%9B%E9%98%B6)
- [第三章：Go标准库包](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%B8%89%E7%AB%A0go%E6%A0%87%E5%87%86%E5%BA%93%E5%8C%85)
- [第四章：Go语言算法](https://github.com/0e0w/365GoLang#%E7%AC%AC%E5%9B%9B%E7%AB%A0go%E8%AF%AD%E8%A8%80%E7%AE%97%E6%B3%95)
- [第五章：Go安全开发](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%BA%94%E7%AB%A0go%E5%AE%89%E5%85%A8%E5%BC%80%E5%8F%91)
- [第六章：Go网络爬虫](https://github.com/0e0w/365GoLang#%E7%AC%AC%E5%85%AD%E7%AB%A0go%E7%BD%91%E7%BB%9C%E7%88%AC%E8%99%AB)
- [第七章：GoWeb开发](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%B8%83%E7%AB%A0goweb%E5%BC%80%E5%8F%91)
- [第八章：Go语言源码](https://github.com/0e0w/365GoLang#%E7%AC%AC%E5%85%AB%E7%AB%A0go%E8%AF%AD%E8%A8%80%E6%BA%90%E7%A0%81)
- [第九章：Go逆向工程](https://github.com/0e0w/365GoLang#%E7%AC%AC%E4%B9%9D%E7%AB%A0go%E9%80%86%E5%90%91%E5%B7%A5%E7%A8%8B)

## 0x02-学习进度

### 第一章：Go语言基础

<details>
<summary>Day001: 基础-Go语言入门</summary>

- [x] 本节说明：本节介绍Go语言的历史与发展。

- [x] Go语言介绍：

  - Go语言是一门编译型的开源的程序设计语言。因为开源，所以编译器、库和工具的源代码可以免费获得。通过Go语言可以容易的构造出简单、可靠且高效的软件。 
  - Go语言是由[Robert Griesemer](https://github.com/griesemer)、[Rob Pike](https://github.com/robpike)、[Ken Thompson](https://baike.baidu.com/item/Ken%20Thompson/3441433)在2007年末主持开发，后来加入了[Ian Lance Taylor](https://github.com/ianlancetaylor/)、[Russ Cox](https://github.com/rsc)等。Go语言最终在2009年11月开源，并在2012年发布了Go1稳定版本。
  - Go语言没有构造或析构函数、没有运算符重载、没有形参默认值、没有异常、没有宏、没有函数注解、没有线程局部存储。暂时没有泛型。
  - Go 语言不是面向对象编程语言：没有类继承，甚至没有类。
  - Go语言以一种不同寻常的方式来诠释面向对象程序设计。
  - Go语言有垃圾回收、有包系统、有一等公民函数、有词法作用域、有系统调用接口等。
  
- [x] Go语言特点：Go语言和其他语言相比的优势是什么？

  - 快速编译，高效执行，易于开发。
  - 对于网络通信、并发和并行编程有很好的支持，可以更好地利用大量的分布式和多核的计算机。
  - Go语言通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信。他们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。
  - Go 语言从本质上（程序和结构方面）来实现并发编程。
  - 作为强类型语言，隐式的类型转换是不被允许的。一条原则：让所有的东西都是显式的。
  - Go语言本身是由C语言开发的，而不是Go语言（Go1.5开始自举）。
  - Go语言的二进制文件体积是最大的（每个可执行文件都包含 runtime）。
  - Go语言做为一门静态语言，但却和很多动态脚本语言一样得灵活。

- [x] Go语言官网：有大量的教程和代码想项目案例。

  - https://golang.org
  - https://github.com/golang
  - [LearnGolang](https://github.com/0e0w/LearnGolang)

- [x] Go语言安装：

  - [官网下载](https://golang.org/dl/)之后直接按照安装说明安装即可。在Ubuntu虚拟机里面开发使用Go语言：

    ```
    wget https://golang.google.cn/dl/go1.15.2.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.15.2.linux-amd64.tar.gz
    ```

- [x] Go环境变量：

  - 设置GOPATH：

    ```
    mkdir ~/go
    echo "GOPATH=$HOME/go" >> ~/.bashrc
    echo "export GOPATH" >> ~/.bashrc
    echo "PATH=\$PATH:\$GOPATH/bin # Add GOPATH/bin to PATH for scripting" >> ~/.bashrc
    source ~/.bashrc
    ```

- [x] Go语言编辑器：

  - [Goland](https://www.jetbrains.com/go)：JetBrains 公司的 Go 开发工具。
  - [LiteIDE](http://liteide.org)：一款开源跨平台轻量级的Go语言IDE。
  
- [x] Go语言命令：

  - 运行go程序：

    ```
    go run hello.go
    ```

  - 打包成可执行程序：
    
    ```
    go build hello.go
    go build
    ```
    
  - 生成不同平台下的可执行程序：需配置GOPATH。
    
    ```
    // Linux下编译Mac, Windows平台的64位可执行程序：
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 001.go
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 001.go
    ```
    
    ```
    // Windows下编译Mac, Linux平台的64位可执行程序：
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 001.go
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 001.go
    ```
    
    ```
    // Mac下编译Linux, Windows平台的64位可执行程序：
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 001.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 001.go
    ```
    
  - 安装依赖：
    
    ```
    go mod download
    go get -u github.com/0e0w/365GoLang
    ```
    
  - 格式化Go代码：
    
    ```
    gofmt
    ```
    
  - 查看Go环境配置：
    
    ```
    go env
    ```
    
  - 第三方包进行部署：

    ```
    go install
    ```

  - 其他的go命令：

    ```
    bug         start a bug report
    build       compile packages and dependencies
    clean       remove object files and cached files
    doc         show documentation for package or symbol
    env         print Go environment information
    fix         update packages to use new APIs
    fmt         gofmt (reformat) package sources
    generate    generate Go files by processing source
    get         add dependencies to current module and install them
    install     compile and install packages and dependencies
    list        list packages or modules
    mod         module maintenance
    run         compile and run Go program
    test        test packages
    tool        run specified go tool
    version     print Go version
    vet         report likely mistakes in packages
    ```
    
  - 从GOPATH/src目录开始引入包，需关闭go mod模式：

    ```
    export GO111MODULE=off
    ```

- [x] Go语言代理：

  - Go语言大量项目托管于Github，导致国内进行构建程序时会很慢。可使用下列的代理加快构建。

    ```
    https://mirrors.aliyun.com/goproxy/
    https://goproxy.io/zh/
    ```
    
    ```
    go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
    ```

- [x] Go语言未来：

  - Go语言拥有大量的优秀社区框架。
  - Go语言的未来发展前景是光明的。
  
  </details>
  
<details>
<summary>Day002: 基础-Go语言例子</summary>

- [x] 本节说明：通过一个简单例子来认识Go语言的基本程序结构。

- [x] 一个例子：Hello World！

  ```go
  package main
  
  import (
  	"fmt"
  )
  
  func main() {
  	fmt.Println("Hello World!")
  }
  ```
  
  
  - 包声明：package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
    - 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
    - 引入包：使用import圆括号进行引入包。
  
      - 引入标准包
      - 引入第三方包
    - 函数：使用func定义
    - 变量：
    - 常量：
    - 语句：
    - 语法：
  
- [x] Go语言名称：

  - 程序一般由关键字、常量、变量、运算符、类型和函数组成。 程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。 程序中可能会使用到这些标点符号：.、,、;、: 和 …。
  - 在变量与运算符间加入空格，程序看起来更加美观。
  - 标识符：用来命名变量、类型等程序实体。一个标识符就是一个或是多个字母( A ~ Z 和 a ~ z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

- [x] 关键字：关键字是一些特殊的用来帮助编译器理解和解析源代码的单词。Go语言中有25个关键字或保留字。只能用在语法允许的地方，不能作为名称使用。

    - 声明代码元素：const、func、import、package、type、var
    - 组合类型表示：chan、interface、map、struct
    - 流程控制语句：break、case、continue、default、else、fallthrough、for、goto、if、switch、select、range、return
    - 特殊的关键字：defer、go。也可以看作是流程控制关键字， 但它们有一些特殊的作用。

- [x] 预定义标识符：三十几个内置的预申明的常量、类型和函数。所有的类型名、变量名、常量名、跳转标签、包名和包的引入名都必须是标识符。

     - 常量：true、flase、iota、nil
     - 类型：int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、float32、float64、complex128、complex64、bool、byte、rune、string、error
     - 函数：make、len、cap、new、append、copy、close、delete、complex、real、imag、panic、recover

- [x] 声明：

   - 声明是给一个程序实体命名，并设定其部分或全部属性。
   - 有4个主要的声明：
     - 变量（var）
     - 常量（const）
     - 类型（type）
     - 函数（func）

   - 函数的声明包含一个名字、一个参数列表、一个可选的返回值列表以及函数体。

- [x] 本节案例：

   ```go
   package main
   
   import (
   	"fmt"
   )
   
   func main() {
   	fmt.Println("Hello World!")
   }
   ```

</details>

<details>
<summary>Day003: 基础-Go约定惯例</summary>

- [x] 本节说明：本节介绍Go语言中的一些约定和惯例。

- [x] Go编码规范：https://golang.org/ref/spec

- [x] 可见性规则：

  - 在Go语言中，标识符必须以一个大写字母开头，这样才可以被外部包的代码所使用，这被称为导出。标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的。但是包名不管在什么情况下都必须小写。
  - 在设计Go语言时，设计者们也希望确保它不是过于以ASCII为中心，这意味着需要从7位ASCII的范围来扩展标识符的空间。 所以Go语言标识符规定必须是Unicode定义的字母或数字，标识符是一个或多个Unicode字母和数字的序列， 标识符中的第一个字符必须是Unicode字母。
  - 总而言之，为了确保我们的标识符能正常导出，我们建议在开发中还是尽量使用ASCII 码来作为标识符，虽然设计者们在避免以ASCII 码为中心，但出于习惯我们还是服从于这个现实。

- [x] 命名规范：

  - 当某个函数需要被外部包调用的时候需要使用大写字母开头，并遵循 Pascal 命名法（“大驼峰式命名法”）；否则就遵循“小驼峰式命名法”，即第一个单词的首字母小写，其余单词的首字母大写。
  - 单词之间不以空格断开或连接号（-）、底线（_）连结，第一个单词首字母采用大写字母；后续单词的首字母亦用大写字母，例如：FirstName、LastName。每一个单词的首字母都采用大写字母的命名格式，被称为“Pascal命名法”，源自于Pascal语言的命名惯例，也有人称之为“大驼峰式命名法”（Upper Camel Case），为驼峰式大小写的子集。
  - Go 语言追求简洁的代码风格，并通过 gofmt 强制实现风格统一。
  
- [x] 语法惯例：

  - Go 语言也使用分号作为语句的结束，但一般会省略分号。像在标识符后面；整数、浮点、复数、Rune或字符串等字面量后面；关键字break、continue、fallthrough、或者return后面；操作符或标点符号++、--、)、]或}之后等等都可以使用分号，但是往往会省略掉，像LiteIDE编辑器会在保存.go文件时自动过滤掉这些分号，所以在Go语言开发中一般不用过多关注分号的使用。
  - 左大括号 { 不能单独一行，这是编译器的强制规定，否则你在使用 gofmt 时就会出现错误提示“ expected declaration, found '{' ”。右大括号 } 需要单独一行。
  - 在定义接口名时也有惯例，接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

- [x] 注释：

  - 尽管高级编程语言代码比底层机器指令友好和易懂，我们还是需要一些注释来帮助自己和其他程序员理解我们所写的代码。
  - 行注释：使用双斜线//开始，一般后面紧跟一个空格。行注释是Go语言中最常见的注释形式，在标准包中，一般都采用行注释，建议采用这种方式。
- 块注释：使用 /* */，块注释不能嵌套。块注释一般用于包描述或注释成块的代码片段。
  
  </details>

<details>
<summary>Day004: 数据-Go语言变量</summary>

- [x] 本节说明：本节介绍Go语言变量的相关内容。

- [x] Go变量基本描述：

  - 变量可以被看作是在运行时刻存储在内存中并且可以被更改的有名字的值。
  - 所有的变量值都是类型确定值。每个局部声明的变量至少要被有效使用一次。
  - Go 语言变量标识符由字母、数字、下划线组成，首字母不能是数字，区分大小写。
  - Go语言规范中，下划线“_”也被认为是字母。
  - Go语言声明变量时将变量的类型放在变量的名称后面。
  - Go语言变量有2种声明方式，var标准变量申明和短变量声明。

- [x] 标准变量声明：

  - 使用var声明，声明之后需要进行初始化。未初始化时是对应数据的零值。

  - Go语言变量的命名规则遵循骆驼命名法。

  - 下列的因式分解关键字的写法一般用于声明全局变量。

    ```go
    var (
        a int
        b bool
        str string
        浮点 float32    // 中文可以作为变量标识符
    )
    ```

    ```go
    // 变量a，b都是指针类型
    var a, b *int
    ```

  - 声明变量之后，变量会自动初始化。初始值对应类型的零值。当一个变量被var声明之后，系统自动赋予它该类型的零值。所有的内存在 Go 中都是经过初始化的。

    - 数字类型对应的是0
    - 布尔类型对应的是flase
    - 字符串类型对应的是""
    - 接口和引用类型对应的是nil

    ```go
    var name type = expression
    var i, j, k int
    var b, f, s, = true, 2.3, "four" 
    ```

- [x] 短变量声明：

  - 变量初始化时省略变量的类型会由系统自动推断。短变量声明时会同时进行初始化。

  - 变变量声明是变量声明的首选形式，但是只能用在函数体内，不可以用于全局变量的声明与赋值。

  - 使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。

  - 简式声明一般用在func内，要注意的是：全局变量和简式声明的变量尽量不要同名，否则很容易产生偶然的变量隐藏Accidental Variable Shadowing。

  - name := expression

    ```go
    a, b, c := 5, 7, "abc"  // 注意等号前的冒号
    ```

- [x] 变量赋值：

  - 多变量可以在同一行进行赋值，也称为并行或同时或平行赋值。

    ```go
    a, b, c = 5, 7, "abc"
    ```

  - 并行赋值也被用于当一个函数返回多个返回值时，比如这里的 val 和错误 err 是通过调用 Func1 函数同时得到：

    ```go
    val, err = Func1(var1)
    ```

- [x] 空白标识符 _ ：

  - 空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
  - _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
  - 空白标识符是一个特殊的标识符.它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。
  - 由于Go语言有个强制规定，在函数内一定要使用声明的变量，但未使用的全局变量是没问题的。为了避免有未使用的变量，代码将编译失败，我们可以将该未使用的变量改为 _。
  - 另外，在Go语言中，如果引入的包未使用，也不能通过编译。有时我们需要引入的包，比如需要init()，或者调试代码时我们可能去掉了某些包的功能使用，你可以添加一个下划线标记符，_，来作为这个包的名字，从而避免编译失败。下滑线标记符用于引入，但不使用。

- [x] 零值nil：

  - nil 标志符用于表示interface、函数、maps、slices、channels、error、指针等的“零值”。如果你不指定变量的类型，编译器将无法编译你的代码，因为它猜不出具体的类型。

- [ ] 本节案例：

  ```
  
  ```
  
  </details>
<details>
<summary>Day005: 数据-Go语言常量</summary>

- [x] 本节说明：本节介绍Go语言常量的内容。

- [x] 常量说明：

  - 常量使用关键字 const 定义，用来存储不会改变的数据。
  - 常量不能被重新赋予任何值。常量名建议使用大写字母。
  - 存储在常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
  - 常量声明中的等号=表示“绑定”而非“赋值”。每个常量描述将一个或多个字面量绑定到各自对应的有名常量上。每个有名常量其实代表着一个字面常量。
  - 常量可以直接声明在包中（全局常量），也可以声明在函数体中（局部常量）。

- [x] 常量定义：

  - 常量的定义格式：const identifier [type] = value，例如：

    ```go
    const Pi = 3.14159
    ```
    
  - Go语言常量定义可以指定常量类型，但不是必需的。如果定义常量时没有指定类型，那么它与字面常量一样，是无类型（untyped）的常量。
    
  - 一个未指定类型的常量被使用时，会根据其使用环境而推断出它所需要具备的类型。
    
    ```go
    显式类型定义：const b string = "abc"
    隐式类型定义：const b = "abc"
    ```
  - 常量也可以在单行进行多重赋值：
  
    ```go
    const a, b, c = 1, false, "str" //多重赋值
    ```
  
- [x] iota： 特殊常量

  - iota 在 const关键字出现时将被重置为 0，const 中每新增一行常量声明将使 iota 计数一次。

  - iota 可理解为 const 语句块中的行索引。

  - iota 可以被用作枚举值。第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1。

     ```go
     const (
         a = iota
         b = iota
         c = iota
     )
     ```

     第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

     ```go
     const (
         a = iota
         b
         c
      )
     ```

     如果对b重新赋值之后，a, b, c分别为0, 8, 8，新的常量b声明后，iota 不再向下赋值，后面常量如果没有赋值，则继承上一个常量值。

     ```go
     const (
         a = iota
         b = 8
         c
      )
     ```

     使用位左移与 iota 计数配合可优雅地实现存储单位的常量枚举：

     ```go
     type ByteSize float64
     const (
         _ = iota // 通过赋值给空白标识符来忽略值
         KB ByteSize = 1<<(10*iota)
         MB
         GB
         TB
     )
     ```

  </details>
<details>
<summary>Day006: 语法-Go语言语法</summary>

- [x] 本节说明：本节介绍表达式、语句和简单语句的相关内容。

- [x] Go语言语法介绍：

  - 一个表达式表示一个值，而一条语句表示一个操作。 
  - 和很多其它流行语言一样，Go也支持break、continue和goto等跳转语句。 另外，Go 还支持一个特有的fallthrough跳转语句。
  
- [x] 简单语句类型：

  - 变量短声明语句。
  - 纯赋值语句，包括x op= y这种运算形式。
  - 有返回结果的函数或方法调用，以及通道的接收数据操作。 
  - 通道的发送数据操作。
  - 空语句。
  - 自增（x++）和自减（x--）语句。
  
- [ ] 非简单语句类型：
  
  - 标准变量声明语句。
  - 常量声明语句。
  - 类型声明语句。
  - 包引入语句。
  - 显式代码块。
  - 函数声明。 
  - 流程控制跳转语句。
  - 函数返回（return）语句。
  - 延迟函数调用和协程创建语句。
  
- [ ] 三种基本的流程控制：
  
  - if-else条件分支代码块。
  - switch-case多条件分支代码块
  - for循环代码块。
  
- [ ] 特定类型相关的流程控制代码块：
  
  - 容器类型相关的for-range循环代码块。
  - 接口类型相关的type-switch多条件分支代码块。
  - 通道类型相关的select-case多分支代码块。
  
- [ ] 本节案例
  
  </details>
<details>
<summary>Day007: 语句-Go运算符号</summary>

- [x] 本节说明：本节介绍Go运算符相关内容。

- [x] Go运算符：

  - Go语言不需要在语句或声明后面是有分号结尾。
  
- [x] 算术运算符：

  | 算术运算符 | 描述 | 实例               |
  | :--------- | :--- | :----------------- |
  | +          | 相加 | A + B 输出结果 30  |
  | -          | 相减 | A - B 输出结果 -10 |
  | *          | 相乘 | A * B 输出结果 200 |
  | /          | 相除 | B / A 输出结果 2   |
  | %          | 求余 | B % A 输出结果 0   |
  | ++         | 自增 | A++ 输出结果 11    |
  | --         | 自减 | A-- 输出结果 9     |

- [x] 比较运算符：

  | 运算符 |                             描述                             | 实例              |
  | :----- | :----------------------------------------------------------: | :---------------- |
  | ==    |    检查两个值是否相等，如果相等返回 True 否则返回 False。    | (A == B) 为 False |
  | !=     |  检查两个值是否不相等，如果不相等返回 True 否则返回 False。  | (A != B) 为 True  |
  | >      |  检查左边值是否大于右边值，如果是返回 True 否则返回 False。  | (A > B) 为 False  |
  | <      |  检查左边值是否小于右边值，如果是返回 True 否则返回 False。  | (A < B) 为 True   |
  | >=    | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
  | <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True |
  
- [x] 逻辑运算符：

  - Go支持两种布尔二元运算符和一种布尔一元运算符。
  
  | 运算符 |  描述  |        实例        |
  | :----: | :----: | :----------------: |
  |   &&   | 逻辑与 | (A && B) 为 False  |
  |  \|\|  | 逻辑或 | (A \|\| B) 为 True |
  |   !    | 逻辑非 | !(A && B) 为 True  |
  
- [x] 位运算符：

  位运算符对整数在内存中的二进制位进行操作。 下表列出了位运算符 &，|，和 ^ 的计算：
  
  | p    | q    | p & q | p \| q | p ^ q |
  | :---  | :--- | :---- | :----- | :---: |
  | 0    | 0    | 0     | 0      |   0   |
  | 0    | 1    | 0     | 1      |   1   |
  | 1    | 1    | 1     | 1      |   0   |
  | 1    | 0    | 0     | 1      |   1   |
  
  | 运算符 |                    描述                    |                  实例                  |
  | :----: | :----------------------------------------: | :------------------------------------: |
  |   &    | 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
  |   \|   |  其功能是参与运算的两数各对应的二进位相或  | (A \| B) 结果为 61, 二进制为 0011 1101 |
  |   ^    | 二进位相异或，两对应的二进位相异时结果为1  | (A ^ B) 结果为 49, 二进制为 0011 0001  |
  |   <<   |                                            | A << 2 结果为 240 ，二进制为 1111 0000 |
  |   >>   |                                            |       A >> 2 结果为 15 ，二进制        |
  
- [x] 赋值运算符：

  | 运算符 |       描述       |                 实例                  |
  | :----: | :--------------: | :-----------------------------------: |
  |   =    | 简单的赋值运算符 | C = A + B 将 A + B 表达式结果赋值给 C |
  |   +=   |   相加后再赋值   |         C += A 等于 C = C + A         |
  |   -=   |   相减后再赋值   |         C -= A 等于 C = C - A         |
  |   *=   |   相乘后再赋值   |         C *= A 等于 C = C * A         |
  |   /=   |   相除后再赋值   |         C /= A 等于 C = C / A         |
  |   %=   |   求余后再赋值   |         C %= A 等于 C = C % A         |
  |  <<=   |    左移后赋值    |        C <<= 2 等于 C = C << 2        |
  |  >>=   |    右移后赋值    |        C >>= 2 等于 C = C >> 2        |
  |   &=   |   按位与后赋值   |         C &= 2 等于 C = C & 2         |
  |   ^=   |  按位异或后赋值  |         C ^= 2 等于 C = C ^ 2         |
  |  \|=   |   按位或后赋值   |        C \|= 2 等于 C = C \| 2        |
  
- [x] 其他运算符：

  | 运算符 | 描述             |            实例            |
  | :----- | :--------------- | :------------------------: |
  | &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
  | *      | 指针变量。       |     *a; 是一个指针变量     |
  
- [x] 运算符优先级：
  
  | 优先级 |      运算符      |
  | :----: | :--------------: |
  |   5    | * / % << >> & &^ |
  |   4    |     + - \| ^     |
  |   3    | == != < <= > >=  |
  |   2    |        &&        |
  |   1    |       \|\|       |
  
- [ ] 几个特殊运算符：

  - 位清除 &^：将指定位置上的值设置为 0。将运算符左边数据相异的位保留，相同位清零 ：

- [ ] 本节案例：

  
  </details>
<details>
<summary>Day008: 语句-Go条件判断</summary>

- [x] 本节说明：本节介绍Go语言条件判断语句的相关内容。

- [x] 条件判断语句介绍：

  - 条件语句需要指定一个或多个条件，通过测试条件是否为 true 来决定是否执行指定语句，当条件为 false 的情况在执行另外的语句。
  
- [x] if语句：

  - if 语句用于测试某个条件（布尔型或逻辑型）的语句。由一个布尔表达式或关系运算符后紧跟一个或多个语句组成。

    ```go
    if 布尔表达式 {
       /* 在布尔表达式为 true 时执行 */
    }
    ```

- [x] if...else 语句：

  - if 语句后可以使用可选的else语句, else语句中的表达式在布尔表达式为 false 时执行。

    ```go
    if 布尔表达式 {
       /* 在布尔表达式为 true 时执行 */
    } else {
      /* 在布尔表达式为 false 时执行 */
    }
    ```

- [x] if 嵌套语句：

  - 你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。

    ```go
    if 布尔表达式 1 {
       /* 在布尔表达式 1 为 true 时执行 */
       if 布尔表达式 2 {
          /* 在布尔表达式 2 为 true 时执行 */
       }
    }
    ```

  - 三层嵌套。

    ```go
    if condition1 {
    	// do something	
    } else if condition2 {
    	// do something else	
    } else {
    	// catch-all or default
    }
    ```

- [x] if条件判断语句案例：

  - 案例1：

    ```go
  if a := 10; a == 10 {
    	fmt.Println("10")
  } else {
    	fmt.Println("not 10")
    }
    ```
  
  - 案例2：
  
    ```go
    a := 10 // 初始化赋值语句
    if a == 10 {
    	fmt.Println("10")
  } else {
    	fmt.Println("not 10")
  }
    ```
  
  - 案例3：
  
    ```go
    if a := 10; a == 10 {
    	fmt.Println("=10")
    } else if a > 10 {
    	fmt.Println(">10")
    } else if a < 10 {
    	fmt.Println("<10")
    } else {
        fmt.Println("这是不可能的!")
    }
    ```
  
- [x] switch 语句：

  - switch语句是存在多个条件判断的情况下，分别执行其对应的语句。

  - switch语句默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case。

  - 如果我们需要执行后面的 case，可以使用 fallthrough 。

  - switch后面可以写变量本身。和case后面的变量进行比较之后进行执行对应语句。

    ```go
    switch var1 { 
        case val1:
            ...
        case val2:
            ...
        default:
            ...
    }
    ```

  - 任何支持进行相等判断的类型都可以作为测试表达式的条件，包括 int、string、指针等。

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	var num1 int = 7
    
    	switch { //无条件执行
    	    case num1 < 0:
    		    fmt.Println("Number is negative")
    	    case num1 > 0 && num1 < 10:
    		    fmt.Println("Number is between 0 and 10")
    	    default:
    		    fmt.Println("Number is 10 or greater")
    	}
    }
    ```

- [ ] [参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/05.3.md)

- [x] select 语句：

  - select 结构，用于 channel 的选择。

  - select 是 Go 中的一个控制结构，类似于 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

  - select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。

  - select没有条件表达式，一直在等待分支进入可运行状态。

    ```go
    select {
        case communication clause  :
           statement(s);      
        case communication clause  :
           statement(s);
        /* 你可以定义任意数量的 case */
        default : /* 可选 */
           statement(s);
    }
    ```

  注意：Go 没有三目运算符，所以不支持 ?: 形式的条件判断。

- [ ] [参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/05.1.md)
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day009: 语句-Go循环语句</summary>

- [x] 本节说明：本节介绍Go语言循环语句的相关内容。

- [x] Go循环语句：

  - 在实际问题中有大量的具有规律性的重复操作，在程序开发中便需要重复执行某些语句。
  - 在Go语言中只有for循环一种循环结构。

- [x] for循环：重复执行语句块

  - Go 语言中只有 for 结构可以重复执行某些语句。

  - for 循环是一个循环控制结构，可以执行指定次数的循环。

  - Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

  - **基于计数器的迭代**：

    ```go
    for  初始化条件; 判断条件; 条件变化 {}
    for init; condition; post { }
    // init： 一般为赋值表达式，给控制变量赋初值；
    // condition： 关系表达式或逻辑表达式，循环控制条件；
    // post： 一般为赋值表达式，给控制变量增量或减量。
    // 在循环中同时使用多个计数器：
    for i, j := 0, N; i < j; i, j = i+1, j-1 {}
    ```

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	for i := 0; i < 5; i++ {
    		fmt.Printf("This is the %d iteration\n", i)
    	}
    }
    ```

  - **基于条件判断的迭代**：

    ```go
    package main
    
    import "fmt"
    
    func main() {
    	var i int = 5
    
    	for i >= 0 {
    		i = i - 1
    		fmt.Printf("The variable i is now: %d\n", i)
    	}
    }
    ```

  - **无限循环：**

    条件语句是可以被省略的，如 i:=0; ; i++ 或 for { } 或 for ;; { }（;; 会在使用 gofmt 时被移除）：这些循环的本质就是无限循环。最后一个形式可以被改写为 for true { }，但一般情况下都会直接写

    ```go
    for { }
    ```

    ```go
    for t, err = p.Token(); err == nil; t, err = p.Token() {
    	...
    }
    ```

    ```go
    package main
    
    import "fmt"
    
    func main() {
            sum := 0
            for {
                sum++ // 无限循环下去
            }
            fmt.Println(sum) // 无法输出
    }
    ```

  - **for-range 结构：**

    for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。

    在循环中可以同时使用多个计数器：

    ```go
    for key, value := range oldMap {
        newMap[key] = value
    }
    ```

    ```go
    package main
    import "fmt"
    
    func main() {
            strings := []string{"google", "runoob"}
            for i, s := range strings {
                    fmt.Println(i, s)
            }
            numbers := [6]int{1, 2, 3, 5}
            for i,x:= range numbers {
                    fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
            }  
    }
    ```

- [x] 循环嵌套：在循环内使用循环。

  - 使用方法：

    ```go
    for [condition |  ( init; condition; increment ) | Range]
    {
       for [condition |  ( init; condition; increment ) | Range]
       {
          statement(s);
       }
       statement(s);
    }
    ```

  - 使用循环嵌套来输出 2 到 100 间的素数：

    ```go
    package main
    import "fmt"
    func main() {
       /* 定义局部变量 */
       var i, j int
       for i=2; i < 100; i++ {
          for j=2; j <= (i/j); j++ {
             if(i%j==0) {
                break; // 如果发现因子，则不是素数
             }
          }
          if(j > (i/j)) {
             fmt.Printf("%d  是素数\n", i);
          }
       }  
    }
    ```

- [x] 循环控制语句：

  - break 语句：

    break可用于for、switch、select语句中。

    用于循环语句中跳出循环，并开始执行循环之后的语句。

    break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。

    在多重循环中，可以用标号 label 标出想 break 的循环。

    ```go
    break
    ```

    ```go
    for {
    	i = i - 1
    	fmt.Printf("The variable i is now: %d\n", i)
    	if i < 0 {
    		break
    	}
    }
    ```

    ```go
    i := 0
    for { // for循环条件永真，会一直循环
    	if i >= 10 { // if条件判断i是否大于10
    	break // if大于10的话就break跳出for循环
    	}
    	i++ // i+1
    	g.Println(i) // 打印i
    }
    ```

  - continue语句：

    continue只能用于for循环。跳过当前循环执行下一次循环语句。

    for 循环中，执行 continue 语句会触发 for 增量语句的执行。

    在多重循环中，可以用标号 label 标出想 continue 的循环。

    ```go
    continue
    ```

    ```go
    package main
    
    func main() {
    	for i := 0; i < 10; i++ {
    		if i == 5 { // 判断i等于5时跳出for循环
    			continue
    		}
    		print(i)
    		print(" ")
    	}
    }
    ```

  - goto 语句：

    Go语言的 goto 语句可以无条件地转移到过程中指定的行。

    goto语句可以用在任何地方，但是不能跨函数使用。

    goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

    但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
    
    ```go
    goto End  //goto是关键字，End是用户起的名字，叫做标签
    .......
    End:
  	fmt.Print("aaa")
    ```
    
    ```go
    package main
    
    func main() {
    	i:=0
    	HERE:
    		print(i)
    		i++
    		if i==5 {
    			return // 跳出函数
    		}
    		goto HERE
    }
    ```

  - return语句：

    如果 return 语句使用在普通的 函数 中，则表示跳出该函数，不再执行函数中 return 后面的代码，可以理解成终止函数。

    如果 return 语句使用在 main 函数中，表示终止 main 函数，也就是终止程序的运行。

- [ ] 本节案例： 
  
  
  
  </details> 
<details>
<summary>Day010: 数据-Go基本数据</summary>

- [x] 本节说明：本节介绍Go语言的一些基本数据类型。

- [x] 基本数据类型：

  - 在 Go 编程语言中，数据类型用于声明函数和变量。
  - 数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。
  - 在大多数高级编程语言中，数据通常被抽象为各种类型（type）和值（value）。 
  - 一个类型可以看作是值的模板。一个值可以看作是某个类型的实例。
  - 大多数编程语言支持自定 义类型和若干预定义类型（即内置类型）。 
  - 一门语言的类型系统是这门语言的灵魂。

- [x] 布尔类型：

  - 布尔型的值只可以是常量 true 或者 false。布尔类型的初值是false。
  
  - 两个类型相同的值可以使用相等 == 或者不等 != 运算符来进行比较并获得一个布尔型的值。
  
  - 非运算符：
  
    ```go
    !T -> false
    !F -> true
    ```
  
  - 和运算符：
  
    ```go
    T && T -> true
    T && F -> false
    F && T -> false
    F && F -> false
    ```
  
  - 或运算符：
  
    ```go
    T || T -> true
    T || F -> true
    F || T -> true
    F || F -> false
    ```
  
- [x] 数字类型：

  - Go语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。整型 int 和浮点型 float32、float64。

  - Go语言中没有 float 类型。Go语言中只有 float32 和 float64。没有double类型。

  - 整型的零值为 0，浮点型的零值为 0.0。

  - 整数：rune是int32的内置别名。 

    ```
    int8（-128 -> 127）
    int16（-32768 -> 32767）
    int32（-2,147,483,648 -> 2,147,483,647）
    int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
    ```

  - 无符号整数：byte是uint8的内置别名。

    ```
    uint8（0 -> 255）
    uint16（0 -> 65,535）
    uint32（0 -> 4,294,967,295）
    uint64（0 -> 18,446,744,073,709,551,615）
    ```

  - 浮点型（IEEE-754 标准）：

    ```
    float32（+- 1e-45 -> +- 3.4 * 1e38）
    float64（+- 5 * 1e-324 -> 107 * 1e308）
    ```

  - 字节型：byte 型，也就是uint8类型。代表了ASCII码的一个字符。

    ```go
    vat bt byte
    bt = 'a'
    fmt.Println(bt)
    //打印97
    ```

- [x] 字符串类型：

  - 从逻辑上说，一个字符串值表示一段文本。 在内存中，一个字符串存储为一个字节 （byte）序列。

  - Go语言中可以使用反引号或者双引号来定义字符串。反引号表示原生的字符串，不进行转义。

  - Go语言中的string类型是一种值类型，存储的字符串是不可变的，如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的。

  - Go语言中的string类型是一种值类型，存储的字符串是不可变的，如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的。

  - 字符串是字节的定长数组。字符串的零值是为长度为零的字符串，即空字符串 ""。

  - 一般的比较运算符（==、!=、<、<=、>=、>）通过在内存中按字节比较来实现字符串的对比。可以通过函数 len() 来获取字符串所占的字节长度，例如：len(str)。

  - 在Go中，字符串值是UTF-8编码的， 甚至所有的Go源代码都必须是UTF-8编码的。

  - 字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 [] 内写入索引，索引从 0 开始计数。

  - 解释字符串：该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括：

    ```
    \n：换行符
    \r：回车符
    \t：tab 键
    \u 或 \U：Unicode 字符
    \\：反斜杠自身
    ```

  - 非解释字符串：
    
    ```
    `This is a raw string \n` 中的 `\n\` 会被原样输出。
    ```
    
  - 字符串拼接：
    
    - 直接使用运算符+
    - fmt.Sprintf()
    - strings.Join()
    - bytes.Buffer
    
  - strings.Builder
    
  - 标准库中有四个包对字符串处理尤为重要：bytes、strings、strconv和unicode包。

  - [字符串参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.7.md)

- [ ] 格式化输出：fmt.Printf()

  - %d  //整型格式
  - %c  //字符型
  - %s  //字符串格式
  - %v  //自动匹配类型
  - %T  //数据类型
  
- [ ] 数据类型转换：

  - int(a)

- [ ] 类型别名：

  - 给一个类型另取一个别名：

    ```go
    type bigint int64
    var a bigint
    ```

- [x] 本节案例：

  

  </details>
<details>
<summary>Day011: 数据-Go语言数组</summary>

- [x] 本节说明：本节介绍Go语言数组(array)的相关内容。

- [x] Array介绍：

  - 数组是具有相同类型的一组已知编号且长度固定的数据项序列。一个数组可以由零个或多个元素组成。
  - 数组类型可以是任意的原始类型例如整型、字符串或者自定义类型。
  - 数组长度必须是一个常量表达式，并且必须是一个非负整数。
  - 以 [] 符号标识的数组类型几乎在所有的编程语言中都是一个基本主力。因为数组的长度是固定的，所以在Go语言中很少直接使用数组。
  - 数组长度也是数组类型的一部分，所以[5]int和[10]int是属于不同类型的。
  
- [x] 声明数组：

  - Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

    ```go
    var 数组变量名 [元素数量]Type
    var a [3]int             // 定义三个整数的数组
    ```
  
  - Go 语言中的数组是一种值类型，所以可以通过 new() 来创建：
  
    ```go
    var arr1 = new([5]int)
    ```
  
- [x] 初始化数组：

  - 初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

- [x] 访问数组元素：

  - 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。

    ```go
    var team [3]string
    team[0] = "hammer"
    team[1] = "soldier"
    team[2] = "mum"
    for k, v := range team {
        fmt.Println(k, v)
    }
    ```

- [x] 多维数组：

  - 数数组通常是一维的，但是可以用来组装成多维数组例如：

    ```
    [3][5]int
    [2][2][2]float64
    ```

- [x] 将数组传递给函数：

  - 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：
    - 传递数组的指针
    - 使用数组的切片

- [x] 本节案例：

  

  </details>

<details>
<summary>Day012: 数据-Go语言切片</summary>

- [x] 本节说明：本节介绍Go语言切片(slice)的相关内容。

- [x] Slice介绍：

  - Go 语言切片是对数组的抽象。
  - 切片是对底层数组一个连续片段的引用，所以切片是一个引用类型。
  - 切片提供对该数组中编号的元素序列的访问。未初始化切片的值为nil。
  - Go 数组的长度不可改变，但切片好比动态数组，可以追加元素，在追加时可能使切片的容量增大。
  - 因为切片是引用，不需要使用额外的内存且比使用数组更有效率，所以在Go代码中切片比数组更常用。
  
- [x] 定义切片：

  - 切片有俩种定义方式

    ```go
    var 切片变量名 []type // 声明一个未指定大小的数组来定义切片
    ```

    ```go
    var slice1 []type = make([]type, len,cap)// 使用make()函数来创建切片
    ```

  - 使用make生成切片：

    ```go
    package main
    import "fmt"
    
    func main() {
    	var slice1 []int = make([]int, 10)
    	// load the array/slice:
    	for i := 0; i < len(slice1); i++ {
    		slice1[i] = 5 * i
    	}
    
    	// print the slice:
    	for i := 0; i < len(slice1); i++ {
    		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
    	}
    	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
    	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
    }
    ```

- [x] 将切片传递给函数：

  ```go
  func sum(a []int) int {
  	s := 0
  	for i := 0; i < len(a); i++ {
  		s += a[i]
  	}
  	return s
  }
  
  func main() {
  	var arr = [5]int{0, 1, 2, 3, 4}
  	sum(arr[:])
  }
  ```

- [x] 切片重组

  - 通过改变切片长度得到新切片的过程称之为切片重组 reslicing。
  - 在一个切片基础上重新划分一个切片时，新的切片会继续引用原有切片的数组。
  - 为了避免这个陷阱，我们需要从临时的切片中使用内置函数copy()，拷贝数据到新切片。

- [ ] 切片初始化：

- [ ] 空(nil)切片：

- [ ] [参考1：Go Slice全面指南](https://mp.weixin.qq.com/s/rYY6TnZcb0FIWjouD2cznQ)、[参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.2.md)

- [ ] 本节案例：

  

  </details>

<details>
<summary>Day013: 数据-Go语言集合</summary>

- [x] 本节说明：本节介绍集合Go语言集合(Map)的相关内容。

- [x] Map介绍：

  - Map 是一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，Map结构也称为关联数组或字典。
  - Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
  - Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。
  - 在声明的时候不需要知道 Map 的长度，Map 是可以动态增长的。
  - 不要使用 new，永远用 make 来构造 map。
  
- [x] 申明定义Map：

  - map 是引用类型，可以使用如下声明：

    ```go
    var map1 map[keytype]valuetype
    var map1 map[string]int
    ```
    
  - 可以使用内建函数 make 也可以使用 map 关键字来定义 Map。

    ```go
    // 声明变量，默认 map 是 nil
    var map_variable map[key_data_type]value_data_type
    
    // 使用 make 函数
    map_variable := make(map[key_data_type]value_data_type)
    
    var m map[string]int
    
    // 声明但未初始化map，此时是map的零值状态
    map1 := make(map[string]string, 5)
    
    map2 := make(map[string]string)
    
    // 创建了初始化了一个空的的map，这个时候没有任何元素
    map3 := map[string]string{}
    
    // map中有三个值
    map4 := map[string]string{"a": "1", "b": "2", "c": "3"}
    ```

  - 一个示例：

    ```go
    package main
    import "fmt"
    
    func main() {
    	var mapLit map[string]int
    	//var mapCreated map[string]float32
    	var mapAssigned map[string]int
    
    	mapLit = map[string]int{"one": 1, "two": 2}
    	mapCreated := make(map[string]float32)
    	mapAssigned = mapLit
    
    	mapCreated["key1"] = 4.5
    	mapCreated["key2"] = 3.14159
    	mapAssigned["two"] = 3
    
    	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
    	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
    	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
    	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
    }
    ```

  - for-range与map

    ```go
    // 使用 for 循环构造 map
    for key, value := range map1 {
    	...
    }
    ```

    ```go
    // 只获取值
    for _, value := range map1 {
    	...
    }
    ```

    ```go
    // 只获取 key
    for key := range map1 {
    	fmt.Printf("key is: %d\n", key)
    }
    ```

- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day014: 数据-Go语言结构</summary>

- [x] 本节说明：本节介绍Go语言结构体(struct)的相关内容。

- [x] struct介绍：

  - Go语言通过类型别名（alias types）和结构体的形式支持用户自定义类型，或者叫定制类型。
  - 一个带属性的结构体试图表示一个现实世界中的实体。
  - 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。结构体中可以定义不同类型的数据。
  - 结构体是复合类型（composite types），它由一系列属性组成，每个属性都有自己的类型和值的，结构体通过属性把数据聚集在一起。结构体也是值类型，因此可以通过new函数来创建。
  - 组成结构体类型的那些数据称为 字段（fields）。每个字段都有一个类型和一个名字；在一个结构体中，字段名字必须是唯一的。
  - 方法（Method）可以访问这些数据，就好像它们是这个独立实体的一部分。

- [x] 定义结构体：

  - 结构体定义需要使用 type 和 struct 语句。结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

    ```go
    type identifier struct {
        field1 type1
        field2 type2
        ...
    }
    ```
    
  - `type T struct {a, b int}` 也是合法的语法，它更适用于简单的结构体。

  - 一个空结构体：struct {}

  - 一旦定义了结构体类型，它就能用于变量的声明。

  - 结构体的字段可以是任何类型，甚至是结构体本身，也可以是函数或者接口。可以声明结构体类型的一个变量，然后像下面这样给它的字段赋值：

    ```go
    var s T
    s.a = 5
    s.b = 8
    ```

  - 一个例子：

    ```go
    package main
    import "fmt"
    
    type struct1 struct {
        i1  int
        f1  float32
        str string
    }
    
    func main() {
        ms := new(struct1)
        ms.i1 = 10
        ms.f1 = 15.5
        ms.str= "Chris"
    
        fmt.Printf("The int is: %d\n", ms.i1)
        fmt.Printf("The float is: %f\n", ms.f1)
        fmt.Printf("The string is: %s\n", ms.str)
        fmt.Println(ms)
    }
    ```

- [x] 访问结构体成员：

  - 如果要访问结构体成员，需要使用点号 . 操作符，格式为：

    ```go
    结构体.成员名
    ```

- [ ] 结构体特性：

  - 结构体的内存布局：Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

  - 递归结构体：递归结构体类型可以通过引用自身指针来定义。这在定义链表或二叉树的节点时特别有用，此时节点包含指向临近节点的链接。

  - 可见性：通过参考应用可见性规则，如果结构体名不能导出，可使用 new 函数使用工厂方法的方法达到同样的目的。

  - 带标签的结构体：结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）。它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容不可以在一般的编程中使用，只有 reflect 包能获取它。

- [ ] 本节参考：[参考1](https://github.com/ffhelicopter/Go42/blob/master/content/42_18_struct.md)、[参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.1.md)
  
- [x] 本节案例：
  
  ```go
  package main
  
  import (
  	"fmt"
  )
  
  type Human struct {
  	name   string // 姓名
  	Gender string // 性别
  	Age    int    // 年龄
  	string        // 匿名字段
  }
  
  type Student struct {
  	Human     // 匿名字段
  	Room  int // 教室
  	int       // 匿名字段
  }
  
  func main() {
  	//使用new方式
  	stu := new(Student)
  	stu.Room = 102
  	stu.Human.name = "Titan"
  	stu.Gender = "男"
  	stu.Human.Age = 14
  	stu.Human.string = "Student"
  
  	fmt.Println("stu is:", stu)
  	fmt.Printf("Student.Room is: %d\n", stu.Room)
  	fmt.Printf("Student.int is: %d\n", stu.int) // 初始化时已自动给予零值：0
  	fmt.Printf("Student.Human.name is: %s\n", stu.name) //  (*stu).name
  	fmt.Printf("Student.Human.Gender is: %s\n", stu.Gender)
  	fmt.Printf("Student.Human.Age is: %d\n", stu.Age)
  	fmt.Printf("Student.Human.string is: %s\n", stu.string)
  
  	// 使用结构体字面量赋值
  	stud := Student{Room: 102, Human: Human{"Hawking", "男", 14, "Monitor"}}
  
  	fmt.Println("stud is:", stud)
  	fmt.Printf("Student.Room is: %d\n", stud.Room)
  	fmt.Printf("Student.int is: %d\n", stud.int) // 初始化时已自动给予零值：0
  	fmt.Printf("Student.Human.name is: %s\n", stud.Human.name)
  	fmt.Printf("Student.Human.Gender is: %s\n", stud.Human.Gender)
  	fmt.Printf("Student.Human.Age is: %d\n", stud.Human.Age)
  	fmt.Printf("Student.Human.string is: %s\n", stud.Human.string)
  }
  ```
  
  </details> 

<details>
<summary>Day015: 数据-Go语言接口</summary>

- [x] 本节说明：本节介绍Go语言接口(interface)的相关内容。

- [x] 接口interface介绍：

  - Go语言有非常灵活的接口概念，通过它可以实现很多面向对象的特性。接口提供了一种方式来说明对象的行为：如果谁能搞定这件事，它就可以用在这儿。
  
  - Go语言接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。
  
  - Go 语言中的所有类型包括自定义类型都实现了interface{}接口，所有的类型如string、 int、 int64甚至是自定义的结构体类型都拥有interface{}空接口，这一点interface{}和Java中的Object类比较相似。
  
  - 空接口interface{}可以被当做任意类型的数值。
  
  - Go语言中的接口都很简短，通常它们会包含0个、最多3个方法。
  
  - 使用接口使代码更具有普适性。
  
  - 接口类型的未初始化变量的值为nil。
  
    ```go
    var i interface{} = 99 // i可以是任何类型
    i = 44.09
    i = "All"  // i 可接受任意类型的赋值
    ```
  
- [x] 接口interface定义：

  - 接口是一组抽象方法的集合，它必须由其他非接口类型实现，不能自我实现。Go 语言通过它可以实现很多面向对象的特性。通过如下格式定义接口：
  
    ```go
    type Namer interface {
        Method1(param_list) return_type
        Method2(param_list) return_type
        ...
    }
    ```
  
- [x] 一些例子：

  - 示例1：

    ```go
    package main
    
    import "fmt"
    
    type Shaper interface {
    	Area() float32
    }
    
    type Square struct {
    	side float32
    }
    
    func (sq *Square) Area() float32 {
    	return sq.side * sq.side
    }
    
    func main() {
    	sq1 := new(Square)
    	sq1.side = 5
    
    	var areaIntf Shaper
    	areaIntf = sq1
    	// shorter,without separate declaration:
    	// areaIntf := Shaper(sq1)
    	// or even:
    	// areaIntf := sq1
    	fmt.Printf("The square has area: %f\n", areaIntf.Area())
    }
    ```

  - 示例2：接口嵌套接口

    一个接口可以包含一个或多个其他的接口，但是在接口内不能嵌入结构体，也不能嵌入接口自身，否则编译会出错。

    接口File包含了ReadWrite和Lock的所有方法，它还额外有一个Close()方法。

    ```go
    type ReadWrite interface {
        Read(b Buffer) bool
        Write(b Buffer) bool
    }
    
    type Lock interface {
        Lock()
        Unlock()
    }
    
    type File interface {
        ReadWrite
        Lock
        Close()
    }
    ```

- [ ] 接口排序：
  
  - 使用 Sorter 接口排序
  
- [ ] 参考链接：[接口参考1](https://github.com/ffhelicopter/Go42/blob/master/content/42_19_interface.md)、[接口参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.1.md)
  
- [ ] 本节案例：
  
  
  
  </details>

<details>
<summary>Day016: 数据-Go语言反射</summary>

- [x] 本节说明：本节介绍Go语言反射(reflect)相关内容。

- [ ] 反射reflect介绍：

  - 反射是程序执行时检查其所拥有的结构。尤其是类型的一种能力。这是元编程的一种形式。
  
- [ ] 本节案例：

  
  </details>

<details>
<summary>Day017: 函数-Go语言函数</summary>

- [x] 本节说明：本节介绍Go语言函数相关内容。

- [x] Go函数介绍：

  - 函数是基本的代码块，用于执行一个任务。Go 语言最少有个 main() 函数。
  - Go语言是编译型语言，函数编写的顺序是无关紧要的；鉴于可读性的需求，最好把 main() 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。
  - Go语言标准库提供了多种内置函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。
  - main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。main 函数既没有参数，也没有返回类型（与 C 家族中的其它语言恰好相反）。如果为 main 函数添加了参数或者返回类型，将会引发构建错误。
  - 目前Go语言没有泛型（generic）的概念，也就是说它不支持那种支持多种类型的函数。
  
- [x] Go语言里面有三种类型的函数：

  - 普通的带有名字的函数。
  - 匿名函数或者lambda函数。
  
  - 方法（Methods）。
  
- [x] 函数定义：

  - 函数声明告诉了编译器函数的名称，返回类型，和参数。

  - Go语言函数不支持输入参数默认值。每个返回结果的默认值是它的类型的零值。

  - 任何一个函数都不能被声明在另一个函数体内。 虽然匿名函数可以定义在函数 体内，但匿名函数定义不属于函数声明。

  - Go语言函数基本组成：关键字func、函数名、参数列表、返回值、函数体和返回语句。语法如下：

    ```GO
    func 函数名(参数) 返回类型{
        函数体
    }
    ```

  - 无参数无返回的函数

    ```go
    func MyFunc(){
    	
    }
    ```

  - 有参数无返回的函数

    ```go
    func main() {
    	Myfunc("url", "t")
    }
    
    func Myfunc(aa, bb string) {
    	fmt.Println(aa)
    	fmt.Println(bb)
    }
    ```

    不定参数类型：不定参数只能是形参中的最好一个参数。

    ```go
    func Myfunc(args ...int) {
    	fmt.Println(aa)
    	fmt.Println(bb)
    }
    ```

  - 无参数有返回的函数：有返回值的函数需要通过return返回。

    ```go
    func MyFunc() int {
    
    }
    ```

  - 有参数有返回的函数

    ```go
    func main() {
    	max, min := MaxAndMin(10, 20)
    	fmt.Println(max, min)
    }
    
    func MaxAndMin(a, b int) (max, min int) {
    	if a > b {
    		max = a
    		min = b
    	} else {
    		max = b
    		min = a
    	}
    	return
    }
    ```

- [x] 函数调用：

  - 一个声明的函数可以通过它的名称和一个实参列表来调用。

  - 一个函数的声明可以出现在它的调用之前，也可以出现在它的调用之后。

  - 一个函数调用可以被延迟执行或者在另一个协程（goroutine）中执行。

  - 当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

  - 函数调用流程：先调用的函数后返回。先进后出。

  - 调用函数，向函数传递参数，并返回值。

    ```go
    pack1.Function(arg1, arg2, …, argn)
    ```
    
    上面的代码中Function 是 pack1 包里面的一个函数，括号里的是被调用函数的实参（argument）：这些值被传递给被调用函数的形参。
    
    ```go
    package main
    
    func main() {
        greeting()
    }
    
    func greeting() {
        println("aa")
    }
    ```
    
  - 函数可以将其他函数调用作为它的参数，只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的

  - 函数也可以以申明的方式被使用，作为一个函数类型，就像：

    ```go
  type binOp func(int, int) int
    ```

    这里不需要函数体 {}。

  - 函数是一等值（first-class value）：它们可以赋值给变量，就像 add := binOp 一样。

- [x] 函数参数：

  - 函数定义时，它的形参一般是有名字的，不过我们也可以定义没有形参名的函数，只有相应的形参类型，就像这样：func f(int, int, float64)。

  - 函数如果使用参数，该变量可称为函数的形参。

  - 形参就像定义在函数体内的局部变量。

  - 没有参数的函数通常被称为 niladic 函数（niladic function），就像 main.main()。

  - 调用函数，可以通过两种方式来传递参数：

      - 按值传递(call by value)：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。

      - 按引用传递(call by reference)：引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

    ```go
    package main
    
    import "fmt"
    
    func main() {
        fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
        // var i1 int = MultiPly3Nums(2, 5, 6)
        // fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
    }
    
    func MultiPly3Nums(a int, b int, c int) int {
        // var product int = a * b * c
        // return product
        return a * b * c
    }
    ```

- [x] 函数返回值：

  - 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂。

  - 一个例子：

    ```go
    package main
    import "fmt"
    func swap(x, y string) (string, string) {
       return y, x
    }
    func main() {
       a, b := swap("Google", "Runoob")
       fmt.Println(a, b)
    }
    ```

- [ ] 延迟函数调用：

  - 在Go语言中，一个函数调用可以跟在一个defer关键字后面，形成一个延迟函数调用。 
- 和协程调用类似，被延迟的函数调用的所有返回值必须全部被舍弃。
  - 当一个函数调用被延迟后，它不会立即被执行。它将被推入由当前协程维护的一个延迟调用堆栈。
  
  - 延迟调用例子：
  
    ```go
    defer fmt.Println("3")
    defer fmt.Println("2")
    fmt.Println("1")
    // 打印 1 2 3
    ```
  
- [x] 回调函数：

  - 函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。

    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	callback(1, Add)
    }
    
    func Add(a, b int) {
    	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
    }
    
    func callback(y int, f func(int, int)) {
    	f(y, 2) // this becomes Add(1, 2)
    }
    ```

- [ ] 递归函数：

  - 函数自己调用自己。

- [x] 匿名函数：

  - 当我们不希望给函数起名字的时候，可以使用匿名函数，例如：func(x, y int) int { return x + y }。

  - 关键字 defer经常配合匿名函数使用，它可以用于改变函数的命名返回值。

  - 匿名函数同样被称之为闭包（函数式语言的术语）：它们被允许调用定义在其它环境下的变量。

  - 包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。

  - Go中的所有的自定义函数（包括声明的函数和匿名函数）都可以被视为闭包。一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。

  - 一个匿名函数在定义后可以被立即调用。

  - 闭包经常被用作包装函数：它们会预先定义好 1 个或多个参数以用于包装。

    ```go
    func() {
    	sum := 0
    	for i := 1; i <= 1e6; i++ {
    		sum += i
    	}
    }()
    ```

- [x] 函数用法：

  - 函数作为另外一个函数的实参：函数定义后可作为另外一个函数的实参数传入。
  - 闭包：闭包是匿名函数，可在动态编程中使用。
  - 方法：方法就是一个包含了接受者的函数

- [ ] 本节案例：

  

  </details>

<details>
<summary>Day018: 函数-Go内置函数</summary>

- [x] 本节说明：本节介绍Go语言内置函数的相关内容。

- [x] Go内置函数介绍：

  - 不引入任何库包而调用一个内置函数。
  - 不需要进行导入操作就可以使用的内置函数。它们有时可以针对不同的类型进行操作，例如：len、cap 和 append，或必须用于系统级的操作，例如：panic。因此，它们需要直接获得编译器的支持。
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day019: 函数-Go语言方法</summary>

- [x] 本节说明：本节介绍Go语言方法相关内容。

- [x] Go语言方法介绍：

  - 在Go语言中，结构体就像是类的一种简化形式，Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
  - 接收者类型可以是（几乎）任何类型，不仅仅是结构体类型。任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。但是接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现。
  - 接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。
  - 一个类型加上它的方法等价于面向对象中的一个类。
  
- [ ] 参考链接：[参考1](https://github.com/ffhelicopter/Go42/blob/master/content/42_20_method.md)、[参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.6.md)
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day020: 错误-Go错误处理</summary>

- [x] 本节说明：本节介绍Go语言中的错误处理。

- [x] 错误处理介绍：

  - Go语言中没有像 Java等语言的 try/catch异常机制，不能执行抛异常操作。但是有defer-panic-and-recover 机制。这种机制就是错误处理。
  - Go语言的设计者觉得try/catch 机制的使用太泛滥了，而且从底层向更高的层级抛异常太耗费资源。他们给Go语言设计的机制也可以 “捕捉” 异常，但是更轻量，并且只应该作为（处理错误的）最后的手段。
  - Go语言是怎么处理普通错误的呢？通过在函数和方法中返回错误对象作为它们的唯一或最后一个返回值——如果返回 nil，则没有错误发生——并且主调（calling）函数总是应该检查收到的错误。
  - Go语言检查和报告错误条件的惯有方式：
    - 产生错误的函数会返回两个变量，一个值和一个错误码；如果后者是 nil 就是成功，非 nil 就是发生了错误。
    - 为了防止发生错误时正在执行的函数（如果有必要的话甚至会是整个程序）被中止，在调用函数后必须检查错误。
  
- [x] 错误处理：

  - Go 有一个预先定义的 error 接口类型：
  
    ```go
    type error interface {
    	Error() string
    }
    ```
  
- [ ] 定义错误：

- [ ] [错误参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/13.1.md)

- [ ] 本节案例：

  </details>

<details>
<summary>Day021: 并发-Go语言协程</summary>

- [x] 本节说明：本节介绍Go语言协程(goroutine)相关内容。

- [x] 协程goroutine介绍：

  - Go语言原生支持应用之间的通信（网络，客户端和服务端，分布式计算）和程序的并发。
    - 不要通过共享内存来通信，而通过通信来共享内存。 
    - Go语言协程比其他语言协程更强大，也很容易从协程的逻辑复用到Go协程。
    - 在Go语言中，协程是创建计算的唯一途径。
    - Go语言不支持创建系统线程，协程是一个Go程序内部唯一的并发实现方式。
  
- [ ] 协程goroutine用法：

  - 使用go关键字。
  - 启动一个新的协程时，协程的调用会立即返回。与函数不同，程序控制不会去等待 Go 协程执行完毕。在调用 Go 协程之后，程序控制会立即返回到代码的下一行，忽略该协程的任何返回值。
  - 如果希望运行其他 Go 协程，Go 主协程必须继续运行着。如果 Go 主协程终止，则程序终止，于是其他 Go 协程也不会继续运行。

- [x] 协程goroutine案例：

  - 案例一：使用延时来返回协程的返回值

    ```
    func main() {	
    	start := time.Now()
    	go tester()
    	time.Sleep(1 * time.Millisecond)
    	end := time.Now()
    	delta := end.Sub(start)
    	g.Println(delta)
    }
    
    func tester() {
    	i := 0
    HERE:
    	g.Println(i)
    	i++
    	if i == 10 {
    		return
    	}
    	goto HERE
    }
    ```

- [ ] 恐慌：一些致命性错误不属于恐慌。对于官方标准编译器来说，很多致命性错误（比如堆栈溢出和内存不足）不能被恢复。它们一旦产生，程序将崩溃。

  - 产生一个恐慌
  - 消除一个恐慌

- [ ] 恢复：

- [x] 参考链接：[协程参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.1.md)、[协程参考2](https://gfw.go101.org/article/control-flows-more.html)

- [ ] 本节案例：

  

  </details>
<details>
<summary>Day022: 并发-Go语言通道</summary>

- [x] 本节说明：本节介绍Go语言通道(channel)的相关内容。
- [x] 通道channel介绍：

  - Go语言奉行通过通信来共享内存，而不是共享内存来通信。所以，channel是协程之间互相通信的通道，协程之间可以通过它发送消息和接收消息。
  - 通道是进程内的通信方式，因此通过通道传递对象的行为与函数调用时参数传递行为比较一致，比如也可以传递指针等。
- [ ] 通道操作符：<-
  - 同一个操作符 <- 既用于发送也用于接收，但Go会根据操作对象弄明白该干什么
- [ ] 本节参考：[通道参考1](https://github.com/ffhelicopter/Go42/blob/master/content/42_22_channel.md)、[通道参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md)
- [ ] 本节案例：


  </details>
### 第二章：Go语言进阶

<details>
<summary>Day000: 数据-Go指针内存</summary>

- [x] 本节说明：本节介绍Go语言指针内存相关内容。

- [x] Go语言指针介绍：

  - Go 语言是一门类型安全和内存安全的编程语言。虽然 Go 语言中仍有指针的存在，但并不允许进行指针运算
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节参考：[参考1](https://github.com/ffhelicopter/Go42/blob/master/content/42_24_pointer.md)、[参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.9.md)
  
- [ ] 本节案例：
  
  
  </details>
<details>
<summary>Day000: 进阶-Go面向对象</summary>

- [x] 本节说明：本节介绍Go语言面向对象的相关内容。

- [ ] Go面向对象：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [ ] [面向对象参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.13.md)

- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 进阶-Go同步与锁</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 进阶-Go语言反射</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 进阶-Go反序列化</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 基础-Go垃圾回收</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] [垃圾回收参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.8.md)

- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 基础-Go语言泛型</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 基础-Go语言进阶</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>

### 第三章：Go标准库包

<details>
<summary>Day000: 库包-Go包的管理</summary>

- [x] 本节说明：本节介绍Go语言库包的相关内容。

- [x] Go语言标准库概述：

  - Go语言标准库就是Go包。需要import导入之后使用某些功能。
  - 像 fmt、os 等这样具有常用功能的内置包在 Go 语言中有 150 个以上，它们被称为标准库，大部分(一些底层的除外)内置于 Go 本身
  
- [ ] 本节案例：

  </details>
<details>
<summary>Day000: 库包-Go输入输出</summary>

- [x] 本节说明：本节介绍Go语言中输入输出的相关内容。

- [x] Go语言输入输出：

  - Go语言标准库就是Go包。需要import导入之后使用某些功能。
  
- [ ] 输入案例参考：

  - 案例1：
  
    ```
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	var a int
    	fmt.Printf("请输入a:")
    	fmt.Scanf("%d", &a)
    	//fmt.Scan(&a)
    	fmt.Println("a =", a)
    
    }
    ```
  
- [ ] 本节案例：

  
  </details>
<details>
<summary>Day000: 库包-Go时间日期</summary>

- [ ] 本节说明：本节介绍Go语言中时间和日期的相关内容。

- [ ] Go语言时间操作：

- [ ] Go语言日期操作：

- [x] 一些具体使用的例子。

  - 计算函数执行时间：
  
    ```go
    start := time.Now()
    longCalculation()
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("longCalculation took this amount of time: %s\n", delta)
    ```
  
  - 计算日期差值：
  
- [ ] [时间日期参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.8.md)、[时间日期参考2](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.11.md)

- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go文本处理</summary>

- [x] 本节说明：本节介绍Go语言处理文本格式的相关内容。

- [x] Go文本处理：也就是处理TXT格式的文件。

  - 文本处理在Go开发中用到的比较多，在安全开发中经常会将文本文件的内容作为参数传递给函数，根据响应进行漏洞扫描与漏洞验证。
  
- [x] 按行读取文本文件：

  - io/ioutil

  - 案例一：按行读取之按行打印内容

    ```go
    package main
    
    import (
        "fmt"
        "os"
    )
    
    func main() {
        userFile := "asatxie.txt"
        fl, err := os.Open(userFile)        
        if err != nil {
            fmt.Println(userFile, err)
            return
        }
        defer fl.Close()
        buf := make([]byte, 1024)
        for {
            n, _ := fl.Read(buf)
            if 0 == n {
                break
            }
            os.Stdout.Write(buf[:n])
        }
    }
    ```

  - 案例二：按行读取之按行打印内容

    ```go
    package main
    
    import (
    	"bufio"
    	"fmt"
    	"io"
    	"os"
    	"strings"
    )
    
    func main() {
    	fileName := "ip.txt"
    	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
    	if err != nil {
    		fmt.Println("Open file error!", err)
    		return
    	}
    	defer file.Close()
    
    	stat, err := file.Stat()
    	if err != nil {
    		panic(err)
    	}
    
    	var size = stat.Size()
    	fmt.Println("file size=", size)
    
    	buf := bufio.NewReader(file)
    	for {
    		line, err := buf.ReadString('\n')
    		line = strings.TrimSpace(line)
            // 实际编程中处理line即可
    		fmt.Println(line)
    		if err != nil {
    			if err == io.EOF {
    				fmt.Println("File read ok!")
    				break
    			} else {
    				fmt.Println("Read file error!", err)
    				return
    			}
    		}
    	}
    }
    
    ```

- [ ] 把文本文件作为参数：

- [ ] 把结果写入到文本文件中：

  - 案例一：

    ```go
    package main
    
    import (
    	"fmt"
    	"os"
    )
    
    func main() {
    	userFile := "astaxie.txt"
    	fout, err := os.Create(userFile)
    	if err != nil {
    		fmt.Println(userFile, err)
    		return
    	}
    	defer fout.Close()
    	for i := 0; i < 10; i++ {
    		fout.WriteString("Just a test!\r\n")
    		fout.Write([]byte("Just a test!\r\n"))
    	}
    }
    
    ```

  - 案例二：将值传递给domain即可

    ```go
    fileName := "is.txt"
    fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    s := strings.Join([]string{domain, "\t", "\n"}, "")
    buf := []byte(s)
    fd.Write(buf)
    fd.Close()
    ```

- [x] [读写数据参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/12.0.md)

- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go请求响应</summary>

- [x] 本节说明：本节介绍Go语言中请求响应的内容。

- [x] HTTP请求响应：

  - 在漏洞利用或是在漏洞验证过程中， 经常使用到的一个方法就是对一个URL发送一个请求，然后从响应中获取相关数据来判断漏洞是否存在。
  - 本小节内容是漏洞验证中较为重要的部分。
  
- [ ] 请求响应包：

  - net/http
  
- [x] 请求响应案例：

  - 访问并读取页面：
  
    ```go
    package main
    
    import (
    	"fmt"
    	"net/http"
    )
    
    var urls = []string{
    	"http://www.google.com",
    	"http://golang.org",
    	"http://blog.golang.org",
    }
    
    func main() {
    	for _, url := range urls {
    		resp, err := http.Head(url)
    		if err != nil {
    			fmt.Println("Error:", url, err)
    		}
    		fmt.Println(url, ": ", resp.Status)
    	}
    }
    ```
    
  - 发送get参数请求：
  - 发送post参数请求：
  - 发送cookie参数请求：
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go正则匹配</summary>

- [x] 本节说明：本节介绍Go语言中正则匹配的相关内容。

- [x] Go语言正则匹配：

  - 在请求响应中经常会根据响应的内容作出判断及操作。也经常需要对响应的内容进行信息提取之后进行下一步的操作，操作的过程中会经常使用正则匹配的方法进行。
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go精密计算</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go读写数据</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go发送邮件</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go语言库包</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go语言库包</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go语言库包</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
<details>
<summary>Day000: 库包-Go语言库包</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>
  
### 第四章：Go语言算法

<details>
<summary>Day000: 算法-Go排序算法</summary>

- [ ] 本节说明：

- [x] 排序算法介绍：

  - 递归函数、递归算法、
  
- [ ] Go排序算法：
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go递归算法</summary>

- [ ] 本节说明：

- [x] 递归算法介绍：

  - 递归函数、递归算法、
  
- [ ] Go递归算法：
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go分治算法</summary>

- [ ] 本节说明：

- [x] 分治算法介绍：

  - 递归函数、递归算法、
  
- [ ] Go分治算法：
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go动态规划</summary>

- [ ] 本节说明：

- [x] 动态规划算法介绍：

  - 递归函数、递归算法、
  
- [ ] Go动态规划算法：
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go贪心算法</summary>

- [ ] 本节说明：

- [x] 贪心算法介绍：

  - 递归函数、递归算法
  
- [ ] Go贪心算法：
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go回溯算法</summary>

- [ ] 本节说明：

- [ ] 回溯算法介绍：

- [x] Go回溯算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>
<details>
<summary>Day000: 算法-Go搜索算法</summary>

- [ ] 本节说明：

- [ ] 搜索算法介绍：

- [x] Go搜索算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go随机算法</summary>

- [ ] 本节说明：

- [ ] 随机算法：

- [x] Go随机算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go图论算法</summary>

- [ ] 本节说明：

- [ ] 图论算法介绍：

- [x] Go图论算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go数论算法</summary>

- [ ] 本节说明：

- [ ] 数论算法介绍：

- [x] Go数论算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go完全算法</summary>

- [ ] 本节说明：

- [ ] 完全算法介绍：

- [x] Go完全算法：

  - 递归函数、递归算法、
  
- [ ] 本节案例：
  
  </details>

<details>
<summary>Day000: 算法-Go漏桶算法</summary>

- [ ] 本节说明：本节介绍漏桶算法的相关内容。

- [ ] 漏桶算法介绍：

- [x] Go漏桶算法：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [ ] [漏桶算法参考1](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.15.md)

- [ ] 本节案例：

  

  </details>

### 第五章：Go安全开发

<details>
<summary>Day000:安全-Go域名扫描</summary>

- [x] 本节说明：本节介绍通过Go语言进行子域名发现的相关内容。

- [x] 子域名发现方法：二级三级域名的发现无外乎下面的这几种方法。

  - 搜索引擎搜索：通过搜索引擎搜索子域名是一种比较好的域名收集方法。
  - 子域名爆破法：通过子域名爆破收集子域名也是很好的方法。
  
- [ ] 子域名爆破原理：

  - 
  
- [ ] 子域名字典整理：
  
- [ ] 项目成品：
  
  - [SubDomainG](https://github.com/0e0w/SubDomainG)：未开源
  
  </details>

<details>
<summary>Day000: 安全-Go目录扫描</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 安全-Go端口扫描</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 安全-Go密码爆破</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 安全-Go漏洞扫描</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 安全-Go隧道代理</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 安全-Go病毒免杀</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>

<details>
<summary>Day000: 安全-Go代码审计</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
### 第六章：Go网络爬虫

<details>
<summary>Day000: 基础-Go语言并发</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>

### 第七章：GoWeb开发

<details>
<summary>Day000: 基础-Go语言并发</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：

  

  </details>

### 第八章：Go语言源码

<details>
<summary>Day000: 源码-Go漏洞扫描</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 源码-Go域名扫描</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  </details>

### 第九章：Go逆向工程

<details>
<summary>Day000: 逆向-Go逆向工程</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] [逆向工程参考1](https://www.anquanke.com/member/122079)

- [ ] 本节案例：
  
  
  
  </details>
<details>
<summary>Day000: 逆向-Go逆向工程</summary>

- [ ] 本节说明：

- [x] Go语言介绍：

  - Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。 
  
- [x] Go语言命令：

  - go run hello.go //编译运行hello.go
  
- [ ] 本节案例：
  
  
  </details>

## 0x03-参考资源

- [x] https://www.runoob.com/go/go-tutorial.html
- [x] https://github.com/ffhelicopter/Go42
- [x] https://github.com/unknwon/the-way-to-go_ZH_CN
- [ ] https://github.com/iswbm/GolangCodingTime
- [ ] https://github.com/golang101/golang101