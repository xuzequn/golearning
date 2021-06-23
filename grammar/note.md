
# Go初级工程师

## Golang基础语法 和 RPC 基础

一个客户端 

一个服务端

注册中心

配置中心

网关

服务治理

序列化 反序列化

golang 返回值可以命名

在函数的返回值列表里定义的变量 在函数中定义赋值 在函数最后写return 返回

stirng [start:end]

swtich case 不需要加break ， 基本类型 字符串

切片 数组
[]byte slice
[10]byte array

data[i] 通过索引查数据

len(data) 长度 cap (data) 容量

append(data, 123)

for i =0 ; i <= n ; i ++

for i, v := range list:

for {
    break
    continue
}

切片是ArrayList， 给予数组的 list实现， 切片不支持随机增删 只有append 操作

切片支持子切片 操作， 子切片与原切片共享底层数组

只有发生扩容子切片与原切片 才会发生变化不共享底层数据

string与 byte 数组相互转化 string()

string 是一个8字节byte 是一个utf-8 数组

string 拼接 使用+

rune

就是字符

不是byte 

是int32 数字的别名 一个rune 四个字节

utf8.RuneCountInString(str)

将 string 统计成真正的字符大小

string 辅助方法 在strings


## RPC

### interface

接口是契约， 实现形式不同， 还有行为

代表抽象。将一致性的抽象定义在一起就是契约、行为

方法接收器

### 反射

运行时获得一些运行程序本身的机制

首先获得方法原本的信息
其次 讲方法的内容改为HTTP调用的内容

	t.NumMethod()
    t.Method(i)
	t.NumField()
    t.Field(i)

TypeOf ValueOf

指针

接口前面不加星号

结构体可以

* 取值， & 取地址

new() 创建对象分配内存， 变量都是对应类型的零值 对比java 不执行构造函数

结构体自引用 使用指针

指针接收器可以修改字段值 *

方法接收器

核心原则：遇事不决用指针

次级原则：不可变对象用结构体

出事换结构体 取地址

ele ：= valueof(val).Elem() 拿到指针指向的结构体

t := ele.Type()

fieldValue := ele.Field(i)

TypeOf 与 ValueOf 不能共存

reflect.MakeFunc() 反射创建方法 

第一个参数是反射类型，第二个参数是反射方法

反射类型通过通过拿到的结构体Type()获得

反射方法通过 func(args []reflect.Value) (results []reflect.Value) 创建方法


[]reflect.Value 创建返回值列表

reflect.ValueOf()创建返回值

reflect.Zero(reflect.TypeOf(new(error)).Elem()) 创建空的error 返回值

fieldValue.Set() 反射篡改方法

fieldValue.Set(reflect.MakeFunc(reflect.Type, func (args []reflect.Value) (results []reflect.Value){})) 

类型断言 

t, ok := x.(T) 或 t := x.(T)

T 可以是结构体也可以是指针

instanceof + 类型强转换

value, ok := x.(T)
x表示要断言的接口变量；
T表示要断言的目标类型；
value表示断言成功之后目标类型变量；
ok表示断言的结果，是一个bool型变量，true表示断言成功，false表示失败，如果失败value的值为nil。


复杂调用支持 

反射分析 FieldValue 约定 func(in *InputType) (* Output, error) 调用规约


变量声明用 首字符大写 公有访问


client := http.Client{}

httpreq := client.Do()

req := http.NewRequest()

### go 单元测试

"testing"

golang 单元测试规范： 

1. 文件用xxx_test.go结尾 

2. 方法形式 TestXXXX(t *testing.T)

t.FailNow() 报错

assert 断言库

"github.com/stretchr/testify"

Ginkgo




### go mod

go mod init xxxxx

go mod tidy 



### api 设计

type  Api interface {
    Action1()
    Action2()
}

type defaultapi struct {

}

func (d *defaultapi) Action() string {
    // 业务逻辑
    return string
}

func (d * defaultapi) Action2() string {
    // 业务逻辑
    return string
}

type MyCustiomApi struct {
    Api
    // 需要的字段
}

func (d *MyCustomApi) Action() string {
    // 自定义逻辑 前置逻辑
    result := d.Api.Action1()

    // 自定义逻辑 后置逻辑

    return result
}

### init 
无参数

无返回值

包被引入的时候执行，只执行一次

执行顺序不定

可以有多个，一个包内每个文件各一个，官方说明随机执行，实际安文件名字母序执行

通过指定一个init 方法， 把几个init方法按顺序写在指定init

不推荐init 有顺序

### 控制字段

"``"
//可以选择的控制字段有三种：

// -：不要解析这个字段

// omitempty：当字段为空（默认值）时，不要解析这个字段。比如 false、0、nil、长度为 0 的 array，map，slice，string

// FieldName：当解析 json 的时候，使用这个名字

### defer

先进后出

相当于 finally


### 文件操作

s.OpenFile("./demo.txt", os.O_CREATE|os.O_APPEND, 6) // 读写方式打开
```
	/*
	  os.O_CREATE|os.O_APPEND
	  或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	  os.O_RDONLY // 只读
	  os.O_WRONLY // 只写
	  os.O_RDWR // 读写
	  os.O_APPEND // 追加（Append）
	  os.O_CREATE // 如果文件不存在则先创建
	  os.O_TRUNC // 文件打开时裁剪文件
	  os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	  os.O_SYNC // 以同步I/O的方式打开
	第三个参数：权限(rwx:0-7)
	  0：没有任何权限
	  1：执行权限
	  2：写权限
	  3：写权限和执行权限
	  4：读权限
	  5：读权限和执行权限
	  6：读权限和写权限
	  7：读权限，写权限，执行权限
	*/
``` 

### 并发变成 sync

sync 包提供了基本的并发工具 

• sync.Map：并发安全 map 

• sync.Mutex：锁 

• sync.RWMutex：读写锁 

• sync.Once：只执行一次



###  invoker
 将和HTTP无关的部分抽出来作为一个INvoker的抽象
 
 1、 考虑扩展非HTTP协议，例如直接使用TCP协议传输
 2、引入Filter等机制
 



 // 初始化注册中心

addservice
addservice

 // 启动服务
 // 判断服务启动成功
 // 注册