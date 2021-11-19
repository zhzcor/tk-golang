package htmltojson

import (
	"golang.org/x/net/html"
	"strings"
)

type Node struct {
	Type    string           `json:"type"`
	Attr    []html.Attribute `json:"attr"`
	Context string           `json:"context"`
	Src     string           `json:"src"`
}

func HtmlToJson(s string) ([]Node, error) {
	var Nodes []Node
	doc, err := parseHtml(s)
	if err != nil {
		return nil, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			Nodes = append(Nodes, Node{Context: n.Data, Type: "text"})
		}
		if n.Type == html.ElementNode {
			if n.Data == "img" {
				var src string
				for _, attribute := range n.Attr {
					if attribute.Key == "src" {
						src = attribute.Val
					}
				}
				Nodes = append(Nodes, Node{Context: n.Data, Type: "element", Attr: n.Attr, Src: src})
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return Nodes, nil

}

func parseHtml(s string) (*html.Node, error) {
	return html.Parse(strings.NewReader(s))
}
