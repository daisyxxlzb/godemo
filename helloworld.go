package main

import "fmt"

func init() {
    fmt.Println("Hello, This is init function!")
}
func main() {
    //{ 不能在单独的行上
    /* 在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数） */
    fmt.Println("Hello, World!")
    //fmt.Print("hello, world\n") //和上面一样的结果
    s1 := `第一行
    第二行
    第三行
    `
    fmt.Println(s1)
}
/*
Hello, This is init function!
Hello, World!
第一行
    第二行
    第三行
*/