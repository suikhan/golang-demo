package main

import (
    "log"
    "github.com/signintech/gopdf"
)

func main(){
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //A4's Size
    pdf.AddPage()
    err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
    if err != nil {
        log.Print(err.Error())
        return
    }

    err = pdf.SetFont("wts11", "", 14)
    if err != nil {
        log.Print(err.Error())
        return
    }

    pdf.Cell(nil, "您好")
    pdf.WritePdf("hello.pdf")
}
