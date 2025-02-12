package main

import (
	"fmt"
	"visibility/models"
)

func main() {
	var u2 models.Userv2
	fmt.Printf("%T\n", u2)
	fmt.Printf("%#v\n", u2)
	// fmt.Println(u.id)
	u2.Name = "kk"
	fmt.Println(u2.Name)
	fmt.Printf("%#v\n", u2)

	u1 := models.NewUserv1(1, "村长") // 只能使用短声明
	fmt.Printf("%T, %#v\n", u1, u1)
	// fmt.Println(u1.id)
	fmt.Println(u1.Name)

	models.SetUserv1Id(u1, 1000) // 语法是否报错(无问题) 能否修改成功(不能)
	fmt.Printf("%#v\n", u1)

	models.SetUserv1IdP(&u1, 1000) // 语法是否报错(无问题) 能否修改成功(能)
	fmt.Printf("%#v\n", u1)

	var u4 models.Userv4
	fmt.Println(u4.Name)
	u4.Name = "xxxxx"
	fmt.Printf("%#v\n", u4)

}
