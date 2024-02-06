# generate

## go generate相关学习

go generate是go 1.4引入的一个工具(命令)，会读取go源文件中的`//go:generate`注释，并执行后面的命令。

`//go:generate`注释后面的命令会被shell执行，所以可以使用shell的语法。

[参考文章](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)

## samplegentool的例子

`go build`编译samplegentool，生成的二进制放入`$GOPATH/bin`(或其他在`$PATH`中的目录)。

在mymod目录下运行`go generate ./...`。

可看到如下输出：

```shell
# Running samplegentool go on anotherfile.go
#   cwd = /yourpath/learngo/demo/generate/mymod
#   os.Args = []string{"samplegentool", "arg1", "arg2", "arg3", "arg4"}
#    GOARCH = amd64
#    GOOS = darwin
#    GOFILE = anotherfile.go
#    GOLINE = 1
#    GOPACKAGE = mymod
#    DOLLAR = $
# Running samplegentool go on mymod.go
#   cwd = /yourpath/learngo/demo/generate/mymod
#   os.Args = []string{"samplegentool", "arg1", "arg2", "-flag"}
#    GOARCH = amd64
#    GOOS = darwin
#    GOFILE = mymod.go
#    GOLINE = 3
#    GOPACKAGE = mymod
#    DOLLAR = $
# Running samplegentool go on mypack.go
#   cwd = /yourpath/learngo/demo/generate/mymod/mypack
#   os.Args = []string{"samplegentool", "arg1", "multiword arg"}
#    GOARCH = amd64
#    GOOS = darwin
#    GOFILE = mypack.go
#    GOLINE = 3
#    GOPACKAGE = mypack
#    DOLLAR = $
```

可以看到在执行过程中可以获取到运行目录(包含`//go:generate`的文件)、命令行参数、文件名、包名等信息。

## stringer的例子

确保stringer已经安装。

```shell
go install golang.org/x/tools/cmd/stringer@latest
```

在stringer目录中我们添加了main.go，其中包含类型`type RoundingMode byte`和他的枚举，并在main中`fmt.Println(ToPositiveInf)`

我们运行`go run .`会看到输出`5`。

我们再运行`go generate ./...`。

会生成`roundingmode_string.go`的文件，包含对类型添加了String方法。

再次运行`go run .`会看到输出`ToPositiveInf`。

对于stringer的使用和生成的内容这里不做细究。