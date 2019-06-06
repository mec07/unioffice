// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/mec07/unioffice"
)

type CT_Path2DMoveTo struct {
	Pt *CT_AdjPoint2D
}

func NewCT_Path2DMoveTo() *CT_Path2DMoveTo {
	ret := &CT_Path2DMoveTo{}
	ret.Pt = NewCT_AdjPoint2D()
	return ret
}

func (m *CT_Path2DMoveTo) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sept := xml.StartElement{Name: xml.Name{Local: "a:pt"}}
	e.EncodeElement(m.Pt, sept)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Path2DMoveTo) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pt = NewCT_AdjPoint2D()
lCT_Path2DMoveTo:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pt"}:
				if err := d.DecodeElement(m.Pt, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Path2DMoveTo %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Path2DMoveTo
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Path2DMoveTo and its children
func (m *CT_Path2DMoveTo) Validate() error {
	return m.ValidateWithPath("CT_Path2DMoveTo")
}

// ValidateWithPath validates the CT_Path2DMoveTo and its children, prefixing error messages with path
func (m *CT_Path2DMoveTo) ValidateWithPath(path string) error {
	if err := m.Pt.ValidateWithPath(path + "/Pt"); err != nil {
		return err
	}
	return nil
}
