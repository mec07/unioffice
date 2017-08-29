// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml"
)

type CT_MultiLvlStrRef struct {
	F                string
	MultiLvlStrCache *CT_MultiLvlStrData
	ExtLst           *CT_ExtensionList
}

func NewCT_MultiLvlStrRef() *CT_MultiLvlStrRef {
	ret := &CT_MultiLvlStrRef{}
	return ret
}
func (m *CT_MultiLvlStrRef) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	sef := xml.StartElement{Name: xml.Name{Local: "f"}}
	gooxml.AddPreserveSpaceAttr(&sef, m.F)
	e.EncodeElement(m.F, sef)
	if m.MultiLvlStrCache != nil {
		semultiLvlStrCache := xml.StartElement{Name: xml.Name{Local: "multiLvlStrCache"}}
		e.EncodeElement(m.MultiLvlStrCache, semultiLvlStrCache)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MultiLvlStrRef) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MultiLvlStrRef:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "f":
				if err := d.DecodeElement(m.F, &el); err != nil {
					return err
				}
			case "multiLvlStrCache":
				m.MultiLvlStrCache = NewCT_MultiLvlStrData()
				if err := d.DecodeElement(m.MultiLvlStrCache, &el); err != nil {
					return err
				}
			case "extLst":
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MultiLvlStrRef
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MultiLvlStrRef) Validate() error {
	return m.ValidateWithPath("CT_MultiLvlStrRef")
}
func (m *CT_MultiLvlStrRef) ValidateWithPath(path string) error {
	if m.MultiLvlStrCache != nil {
		if err := m.MultiLvlStrCache.ValidateWithPath(path + "/MultiLvlStrCache"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}