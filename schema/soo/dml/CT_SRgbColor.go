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
	"fmt"

	"baliance.com/gooxml"
)

type CT_SRgbColor struct {
	ValAttr           string
	EG_ColorTransform []*EG_ColorTransform
}

func NewCT_SRgbColor() *CT_SRgbColor {
	ret := &CT_SRgbColor{}
	return ret
}

func (m *CT_SRgbColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	if m.EG_ColorTransform != nil {
		for _, c := range m.EG_ColorTransform {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SRgbColor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
		}
	}
lCT_SRgbColor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tint"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Tint = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Tint, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "shade"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Shade = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Shade, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "comp"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Comp = NewCT_ComplementTransform()
				if err := d.DecodeElement(tmpcolortransform.Comp, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "inv"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Inv = NewCT_InverseTransform()
				if err := d.DecodeElement(tmpcolortransform.Inv, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gray"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Gray = NewCT_GrayscaleTransform()
				if err := d.DecodeElement(tmpcolortransform.Gray, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alpha"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Alpha = NewCT_PositiveFixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.Alpha, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.AlphaOff = NewCT_FixedPercentage()
				if err := d.DecodeElement(tmpcolortransform.AlphaOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "alphaMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.AlphaMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(tmpcolortransform.AlphaMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hue"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Hue = NewCT_PositiveFixedAngle()
				if err := d.DecodeElement(tmpcolortransform.Hue, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hueOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.HueOff = NewCT_Angle()
				if err := d.DecodeElement(tmpcolortransform.HueOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hueMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.HueMod = NewCT_PositivePercentage()
				if err := d.DecodeElement(tmpcolortransform.HueMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sat"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Sat = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Sat, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "satOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.SatOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.SatOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "satMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.SatMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.SatMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lum"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Lum = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Lum, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lumOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.LumOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.LumOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "lumMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.LumMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.LumMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "red"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Red = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Red, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "redOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.RedOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.RedOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "redMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.RedMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.RedMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "green"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Green = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Green, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "greenOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.GreenOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.GreenOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "greenMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.GreenMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.GreenMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blue"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Blue = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.Blue, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blueOff"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.BlueOff = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.BlueOff, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blueMod"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.BlueMod = NewCT_Percentage()
				if err := d.DecodeElement(tmpcolortransform.BlueMod, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gamma"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.Gamma = NewCT_GammaTransform()
				if err := d.DecodeElement(tmpcolortransform.Gamma, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "invGamma"}:
				tmpcolortransform := NewEG_ColorTransform()
				tmpcolortransform.InvGamma = NewCT_InverseGammaTransform()
				if err := d.DecodeElement(tmpcolortransform.InvGamma, &el); err != nil {
					return err
				}
				m.EG_ColorTransform = append(m.EG_ColorTransform, tmpcolortransform)
			default:
				gooxml.Log("skipping unsupported element on CT_SRgbColor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SRgbColor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SRgbColor and its children
func (m *CT_SRgbColor) Validate() error {
	return m.ValidateWithPath("CT_SRgbColor")
}

// ValidateWithPath validates the CT_SRgbColor and its children, prefixing error messages with path
func (m *CT_SRgbColor) ValidateWithPath(path string) error {
	for i, v := range m.EG_ColorTransform {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ColorTransform[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
