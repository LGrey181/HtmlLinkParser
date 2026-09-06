package main

import (
	"fmt"
	"strings"

	HtmlLinkParser "github.com/peytonweber/HtmlLinkParser"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
  A link to another page
  <span>some span</span>
  </a>
  <a href="/page-two">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := HtmlLinkParser.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
