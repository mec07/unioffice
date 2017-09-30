// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"

	"baliance.com/gooxml"
)

type CT_SharedItems struct {
	// Contains Semi Mixed Data Types
	ContainsSemiMixedTypesAttr *bool
	// Contains Non Date
	ContainsNonDateAttr *bool
	// Contains Date
	ContainsDateAttr *bool
	// Contains String
	ContainsStringAttr *bool
	// Contains Blank
	ContainsBlankAttr *bool
	// Contains Mixed Data Types
	ContainsMixedTypesAttr *bool
	// Contains Numbers
	ContainsNumberAttr *bool
	// Contains Integer
	ContainsIntegerAttr *bool
	// Minimum Numeric Value
	MinValueAttr *float64
	// Maximum Numeric Value
	MaxValueAttr *float64
	// Minimum Date Time
	MinDateAttr *time.Time
	// Maximum Date Time Value
	MaxDateAttr *time.Time
	// Shared Items Count
	CountAttr *uint32
	// Long Text
	LongTextAttr *bool
	// No Value
	M []*CT_Missing
	// Numeric
	N []*CT_Number
	// Boolean
	B []*CT_Boolean
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
	// Date Time
	D []*CT_DateTime
}

func NewCT_SharedItems() *CT_SharedItems {
	ret := &CT_SharedItems{}
	return ret
}

func (m *CT_SharedItems) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ContainsSemiMixedTypesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsSemiMixedTypes"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsSemiMixedTypesAttr))})
	}
	if m.ContainsNonDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsNonDate"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsNonDateAttr))})
	}
	if m.ContainsDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsDate"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsDateAttr))})
	}
	if m.ContainsStringAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsString"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsStringAttr))})
	}
	if m.ContainsBlankAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsBlank"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsBlankAttr))})
	}
	if m.ContainsMixedTypesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsMixedTypes"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsMixedTypesAttr))})
	}
	if m.ContainsNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsNumber"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsNumberAttr))})
	}
	if m.ContainsIntegerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "containsInteger"},
			Value: fmt.Sprintf("%d", b2i(*m.ContainsIntegerAttr))})
	}
	if m.MinValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minValue"},
			Value: fmt.Sprintf("%v", *m.MinValueAttr)})
	}
	if m.MaxValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxValue"},
			Value: fmt.Sprintf("%v", *m.MaxValueAttr)})
	}
	if m.MinDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minDate"},
			Value: fmt.Sprintf("%v", *m.MinDateAttr)})
	}
	if m.MaxDateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "maxDate"},
			Value: fmt.Sprintf("%v", *m.MaxDateAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.LongTextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "longText"},
			Value: fmt.Sprintf("%d", b2i(*m.LongTextAttr))})
	}
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "ma:m"}}
		for _, c := range m.M {
			e.EncodeElement(c, sem)
		}
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "ma:n"}}
		for _, c := range m.N {
			e.EncodeElement(c, sen)
		}
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "ma:b"}}
		for _, c := range m.B {
			e.EncodeElement(c, seb)
		}
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "ma:e"}}
		for _, c := range m.E {
			e.EncodeElement(c, see)
		}
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "ma:s"}}
		for _, c := range m.S {
			e.EncodeElement(c, ses)
		}
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "ma:d"}}
		for _, c := range m.D {
			e.EncodeElement(c, sed)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SharedItems) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "containsSemiMixedTypes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsSemiMixedTypesAttr = &parsed
		}
		if attr.Name.Local == "containsNonDate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsNonDateAttr = &parsed
		}
		if attr.Name.Local == "containsDate" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsDateAttr = &parsed
		}
		if attr.Name.Local == "containsString" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsStringAttr = &parsed
		}
		if attr.Name.Local == "containsBlank" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsBlankAttr = &parsed
		}
		if attr.Name.Local == "containsMixedTypes" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsMixedTypesAttr = &parsed
		}
		if attr.Name.Local == "containsNumber" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsNumberAttr = &parsed
		}
		if attr.Name.Local == "containsInteger" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ContainsIntegerAttr = &parsed
		}
		if attr.Name.Local == "minValue" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MinValueAttr = &parsed
		}
		if attr.Name.Local == "maxValue" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MaxValueAttr = &parsed
		}
		if attr.Name.Local == "minDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MinDateAttr = &parsed
		}
		if attr.Name.Local == "maxDate" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.MaxDateAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
		}
		if attr.Name.Local == "longText" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LongTextAttr = &parsed
		}
	}
lCT_SharedItems:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "m"}:
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "n"}:
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "b"}:
				tmp := NewCT_Boolean()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "e"}:
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "s"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "d"}:
				tmp := NewCT_DateTime()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.D = append(m.D, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_SharedItems %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SharedItems
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SharedItems and its children
func (m *CT_SharedItems) Validate() error {
	return m.ValidateWithPath("CT_SharedItems")
}

// ValidateWithPath validates the CT_SharedItems and its children, prefixing error messages with path
func (m *CT_SharedItems) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.D {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/D[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
