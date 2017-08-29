// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_NotesViewProperties struct {
	// Common Slide View Properties
	CSldViewPr *CT_CommonSlideViewProperties
	ExtLst     *CT_ExtensionList
}

func NewCT_NotesViewProperties() *CT_NotesViewProperties {
	ret := &CT_NotesViewProperties{}
	ret.CSldViewPr = NewCT_CommonSlideViewProperties()
	return ret
}
func (m *CT_NotesViewProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	secSldViewPr := xml.StartElement{Name: xml.Name{Local: "p:cSldViewPr"}}
	e.EncodeElement(m.CSldViewPr, secSldViewPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_NotesViewProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CSldViewPr = NewCT_CommonSlideViewProperties()
lCT_NotesViewProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "cSldViewPr":
				if err := d.DecodeElement(m.CSldViewPr, &el); err != nil {
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
			break lCT_NotesViewProperties
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_NotesViewProperties) Validate() error {
	return m.ValidateWithPath("CT_NotesViewProperties")
}
func (m *CT_NotesViewProperties) ValidateWithPath(path string) error {
	if err := m.CSldViewPr.ValidateWithPath(path + "/CSldViewPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}