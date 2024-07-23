// Copyright 2017 Baliance. All rights reserved.

package main

import (
	"log"

	"github.com/Snowmanazz/gooxml/document"
)

func main() {
	//FormatWordBaseTemplate()

	var pathPre = "F:/OpenSource/gooxml/_examples/document/use-template/"

	doc, err := document.Open(pathPre + "红头文件精简版本测试使用.docx")

	if err != nil {
		log.Fatalf("error opening Windows Word 2016 document: %s", err)
	}

}
