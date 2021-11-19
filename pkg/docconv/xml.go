package docconv

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// ConvertXML converts an XML file to text.
func ConvertXML(r io.Reader) (string, map[string]string, error) {
	meta := make(map[string]string)
	cleanXML, err := Tidy(r, true)
	if err != nil {
		return "", nil, fmt.Errorf("tidy error: %v", err)
	}
	result, err := XMLToText(bytes.NewReader(cleanXML), []string{}, []string{}, true)
	if err != nil {
		return "", nil, fmt.Errorf("error from XMLToText: %v", err)
	}
	return result, meta, nil
}

// XMLToText converts XML to plain text given how to treat elements.
func XMLToText(r io.Reader, breaks []string, skip []string, strict bool) (string, error) {
	var result string

	dec := xml.NewDecoder(r)
	dec.Strict = strict
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		switch v := t.(type) {
		case xml.CharData:
			result += string(v)
		case xml.StartElement:
			for _, breakElement := range breaks {
				if v.Name.Local == breakElement {
					result += "\n"
				}
			}
			for _, skipElement := range skip {
				if v.Name.Local == skipElement {
					depth := 1
					for {
						t, err := dec.Token()
						if err != nil {
							// An io.EOF here is actually an error.
							return "", err
						}

						switch t.(type) {
						case xml.StartElement:
							depth++
						case xml.EndElement:
							depth--
						}

						if depth == 0 {
							break
						}
					}
				}
			}
		}
	}
	return result, nil
}

// XMLToMap converts XML to a nested string map.
func XMLToMap(r io.Reader) (map[string]string, error) {
	m := make(map[string]string)
	dec := xml.NewDecoder(r)
	var tagName string
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		switch v := t.(type) {
		case xml.StartElement:
			tagName = string(v.Name.Local)
		case xml.CharData:
			m[tagName] = string(v)
		}
	}
	return m, nil
}

func XMLFilter(r io.Reader) (string, error) {
	dec := xml.NewDecoder(r)
	var html string
	var body bool
	bodyDom := "body"
	for {
		t, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		switch v := t.(type) {
		case xml.StartElement:
			if v.Name.Local == bodyDom {
				body = true
			}
			//fmt.Println("dom:", v.Name.Local)
			switch v.Name.Local {
			case "p": // 段落
				html += "<" + v.Name.Local + ">"
			case "blip": // 图片
				html += "<file "
				for _, attr := range v.Attr {
					html += fmt.Sprintf("%s=\"%s\"", attr.Name.Local, attr.Value)
				}
				html += " />"

			}
		case xml.CharData:
			if body {
				html += string(v)
			}
		case xml.EndElement:
			if body {
				if v.Name.Local == "p" {
					html += "</" + v.Name.Local + ">"
				}
			}
		}

	}
	return html, nil
}
