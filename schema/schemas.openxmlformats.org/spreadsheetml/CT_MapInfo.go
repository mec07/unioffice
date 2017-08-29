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

type CT_MapInfo struct {
	// Prefix Mappings for XPath Expressions
	SelectionNamespacesAttr string
	// XML Schema
	Schema []*CT_Schema
	// XML Mapping Properties
	Map []*CT_Map
}

func NewCT_MapInfo() *CT_MapInfo {
	ret := &CT_MapInfo{}
	return ret
}
func (m *CT_MapInfo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "SelectionNamespaces"},
		Value: fmt.Sprintf("%v", m.SelectionNamespacesAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	seSchema := xml.StartElement{Name: xml.Name{Local: "x:Schema"}}
	e.EncodeElement(m.Schema, seSchema)
	seMap := xml.StartElement{Name: xml.Name{Local: "x:Map"}}
	e.EncodeElement(m.Map, seMap)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_MapInfo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "SelectionNamespaces" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SelectionNamespacesAttr = parsed
		}
	}
lCT_MapInfo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "Schema":
				tmp := NewCT_Schema()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Schema = append(m.Schema, tmp)
			case "Map":
				tmp := NewCT_Map()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Map = append(m.Map, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MapInfo
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_MapInfo) Validate() error {
	return m.ValidateWithPath("CT_MapInfo")
}
func (m *CT_MapInfo) ValidateWithPath(path string) error {
	for i, v := range m.Schema {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Schema[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Map {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Map[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}