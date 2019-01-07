# xml unmarshal to map for golang

>It's a Demo project, edit your __UnmarshalXML__ method to change behaviour

## usage

```golang
package main

import (
    xmltomap "github.com/nelsonken/xmltomap-go"
)

func main() {
    data := []byte(`<xml>
<a>1</a>
<b>hi hello world</b>
<c>xxxx</c>
<c>rrr</c>
<xml>`)
    strmap, err := xmltomap.UnmarshalXMLToMap(data)
    if err != nil {
        panic(err)
    }

    fmt.Println(strmap)
}
```
