// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/mec07/unioffice"
)

type CT_PlotAreaChoice1 struct {
	ValAx  []*CT_ValAx
	CatAx  []*CT_CatAx
	DateAx []*CT_DateAx
	SerAx  []*CT_SerAx
}

func NewCT_PlotAreaChoice1() *CT_PlotAreaChoice1 {
	ret := &CT_PlotAreaChoice1{}
	return ret
}

func (m *CT_PlotAreaChoice1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ValAx != nil {
		sevalAx := xml.StartElement{Name: xml.Name{Local: "c:valAx"}}
		for _, c := range m.ValAx {
			e.EncodeElement(c, sevalAx)
		}
	}
	if m.CatAx != nil {
		secatAx := xml.StartElement{Name: xml.Name{Local: "c:catAx"}}
		for _, c := range m.CatAx {
			e.EncodeElement(c, secatAx)
		}
	}
	if m.DateAx != nil {
		sedateAx := xml.StartElement{Name: xml.Name{Local: "c:dateAx"}}
		for _, c := range m.DateAx {
			e.EncodeElement(c, sedateAx)
		}
	}
	if m.SerAx != nil {
		seserAx := xml.StartElement{Name: xml.Name{Local: "c:serAx"}}
		for _, c := range m.SerAx {
			e.EncodeElement(c, seserAx)
		}
	}
	return nil
}

func (m *CT_PlotAreaChoice1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PlotAreaChoice1:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "valAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "valAx"}:
				tmp := NewCT_ValAx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ValAx = append(m.ValAx, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "catAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "catAx"}:
				tmp := NewCT_CatAx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CatAx = append(m.CatAx, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dateAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dateAx"}:
				tmp := NewCT_DateAx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DateAx = append(m.DateAx, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "serAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "serAx"}:
				tmp := NewCT_SerAx()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SerAx = append(m.SerAx, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_PlotAreaChoice1 %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PlotAreaChoice1
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PlotAreaChoice1 and its children
func (m *CT_PlotAreaChoice1) Validate() error {
	return m.ValidateWithPath("CT_PlotAreaChoice1")
}

// ValidateWithPath validates the CT_PlotAreaChoice1 and its children, prefixing error messages with path
func (m *CT_PlotAreaChoice1) ValidateWithPath(path string) error {
	for i, v := range m.ValAx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ValAx[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CatAx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CatAx[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.DateAx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DateAx[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.SerAx {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SerAx[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
