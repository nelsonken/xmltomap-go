package xmltomap

import (
	"encoding/xml"
	"io"
)

func Unmarshal(xmlData []byte) (map[string]string, error) {
	u := AccessesMap{}
	err := xml.Unmarshal(xmlData, &u)
	if err == io.EOF {
		err = nil
	}

	return u.M, err
}

type AccessesMap struct {
	M map[string]string
}

func (am *AccessesMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	am.M = map[string]string{}

	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch tt := t.(type) {
		case xml.StartElement:
			if tt.Name.Local == "" {
				continue
			}

			v, err := nextTokenValue(d)
			if err != nil {
				return err
			}

			if mv, ok := am.M[tt.Name.Local]; ok {
				v = mv + "," + v
			}

			am.M[tt.Name.Local] = v
		}
	}

	return nil
}

func nextTokenValue(d *xml.Decoder) (string, error) {
	val := ""
	b := false
	for !b {
		token, err := d.Token()
		if err != nil {
			break
		}
		switch x := token.(type) {
		case xml.EndElement:
			b = true
		case xml.CharData:
			val += string(x)
		}
	}
	return val, nil
}
