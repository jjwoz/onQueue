package util

import (
	"bytes"
	"html/template"
	"log"
)

// CheckErr checks for errors and logs it.
// Generic wrapper so code can be cleaner and not littered with if cases
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func SetSubject(subName, product string) string {
	sub := Subject{To: subName, Product: product}
	var data bytes.Buffer
	subTmpl, err := template.New("subject").Parse("Hey {{.To}}... Renee from the Frey team.")
	CheckErr(err)
	err = subTmpl.Execute(&data, sub)
	CheckErr(err)
	res := data.String()
	return res
}