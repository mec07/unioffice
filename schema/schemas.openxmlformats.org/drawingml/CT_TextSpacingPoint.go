// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TextSpacingPoint struct {
	ValAttr int32
}

func NewCT_TextSpacingPoint() *CT_TextSpacingPoint {
	ret := &CT_TextSpacingPoint{}
	return ret
}
func (m *CT_TextSpacingPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TextSpacingPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ValAttr = int32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TextSpacingPoint: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_TextSpacingPoint) Validate() error {
	return m.ValidateWithPath("CT_TextSpacingPoint")
}
func (m *CT_TextSpacingPoint) ValidateWithPath(path string) error {
	if m.ValAttr < 0 {
		return fmt.Errorf("%s/m.ValAttr must be >= 0 (have %v)", path, m.ValAttr)
	}
	if m.ValAttr > 158400 {
		return fmt.Errorf("%s/m.ValAttr must be <= 158400 (have %v)", path, m.ValAttr)
	}
	return nil
}