package main

import (
	"fmt"
	"golang.org/x/image/colornames"
	"study-server/app/libs/qrcode"
)

func main() {
	q, e := qrcode.New("xxx")

	if e != nil {
		fmt.Println(e)
		return
	}

	q.SetSize(100).
		SetBackgroundColor(colornames.Red)

	pic, e := q.CreateBase64QrCode()

	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(pic)

}
