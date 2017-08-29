// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"log"
)

type CT_CellWatches struct {
	// Cell Watch Item
	CellWatch []*CT_CellWatch
}

func NewCT_CellWatches() *CT_CellWatches {
	ret := &CT_CellWatches{}
	return ret
}
func (m *CT_CellWatches) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secellWatch := xml.StartElement{Name: xml.Name{Local: "x:cellWatch"}}
	e.EncodeElement(m.CellWatch, secellWatch)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_CellWatches) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CellWatches:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cellWatch":
				tmp := NewCT_CellWatch()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellWatch = append(m.CellWatch, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellWatches
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_CellWatches) Validate() error {
	return m.ValidateWithPath("CT_CellWatches")
}
func (m *CT_CellWatches) ValidateWithPath(path string) error {
	for i, v := range m.CellWatch {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellWatch[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}