# cobra

cobra框架，是一个用于构建强大的现代CLI应用程序的库。它提供了一组简单的接口，可以帮助您构建CLI应用程序，而无需考虑太多的细节。

[cobra](https://github.com/spf13/cobra)

[cobra-cli](https://github.com/spf13/cobra-cli/blob/main/README.md)

## 安装

demo使用cobra-cli生成，所以需要安装cobra-cli

```shell
go install github.com/spf13/cobra-cli@latest
```

## 使用

首先需要创建一个go mod项目

```shell
go mod init cobra-demo # go: creating new go.mod: module cobra-demo
```

然后使用cobra-cli生成项目

```shell
cobra-cli init
# Your Cobra application is ready at
# /yourpath/learngo/go-lib-test/cobra
```

生成的项目结构如下：

```shell
cobra
├── LICENSE
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
```

添加一个hello命令

```shell
cobra-cli add hello
# hello created at /yourpath/learngo/go-lib-test/cobra
```

会在cmd目录下生成hello.go文件

```go
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "A brief description of your command",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello called")
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // helloCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

```

创建命令的子命令

```shell
cobra-cli add world -p helloCmd # -p表示父命令, 注意xxxCmd是命令的默认变量名
# world created at /yourpath/learngo/go-lib-test/cobra
```

修改hello.go和world.go文件

```go
// hello.go
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
    Use:   "hello",
    Short: "a demo command for cobra",
    Long:  `a demo command for cobra, print hello to console.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello")
    },
}

func init() {
    rootCmd.AddCommand(helloCmd)
}
```

```go
// world.go
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
    "strings"

    "github.com/spf13/cobra"
)

// worldCmd represents the world command
var worldCmd = &cobra.Command{
    Use:   "world",
    Args:  cobra.MaximumNArgs(1), // accept at most 1 argument
    Short: "a demo sub command for cobra",
    Long:  `a demo sub command for cobra, accept a string parameter, print hello xxx to console, default print hello world.`,
    Run: func(cmd *cobra.Command, args []string) {
        upper, _ := cmd.Flags().GetBool("upper")
        world := "world"
        if len(args) > 0 {
            world = args[0]
        }

        if upper {
            world = strings.ToUpper(world)
        }

        cmd.Printf("hello %s!\n", world)
    },
}

func init() {
    helloCmd.AddCommand(worldCmd)

    // -U, --upper, a flag to print world in upper case
    worldCmd.Flags().BoolP("upper", "U", false, "print world in upper case")
}
```

测试使用

```shell
go run main.go hello
# hello

go run main.go hello world
# hello world!

go run main.go hello world -U
# hello WORLD!

go run main.go hello world xxx -U
# hello XXX!
```
