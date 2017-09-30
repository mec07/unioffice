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

	"baliance.com/gooxml"
)

type CT_ShapeProperties struct {
	BwModeAttr ST_BlackWhiteMode
	Xfrm       *CT_Transform2D
	CustGeom   *CT_CustomGeometry2D
	PrstGeom   *CT_PresetGeometry2D
	NoFill     *CT_NoFillProperties
	SolidFill  *CT_SolidColorFillProperties
	GradFill   *CT_GradientFillProperties
	BlipFill   *CT_BlipFillProperties
	PattFill   *CT_PatternFillProperties
	GrpFill    *CT_GroupFillProperties
	Ln         *CT_LineProperties
	EffectLst  *CT_EffectList
	EffectDag  *CT_EffectContainer
	Scene3d    *CT_Scene3D
	Sp3d       *CT_Shape3D
	ExtLst     *CT_OfficeArtExtensionList
}

func NewCT_ShapeProperties() *CT_ShapeProperties {
	ret := &CT_ShapeProperties{}
	return ret
}

func (m *CT_ShapeProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BwModeAttr != ST_BlackWhiteModeUnset {
		attr, err := m.BwModeAttr.MarshalXMLAttr(xml.Name{Local: "bwMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Xfrm != nil {
		sexfrm := xml.StartElement{Name: xml.Name{Local: "a:xfrm"}}
		e.EncodeElement(m.Xfrm, sexfrm)
	}
	if m.CustGeom != nil {
		secustGeom := xml.StartElement{Name: xml.Name{Local: "a:custGeom"}}
		e.EncodeElement(m.CustGeom, secustGeom)
	}
	if m.PrstGeom != nil {
		seprstGeom := xml.StartElement{Name: xml.Name{Local: "a:prstGeom"}}
		e.EncodeElement(m.PrstGeom, seprstGeom)
	}
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.BlipFill != nil {
		seblipFill := xml.StartElement{Name: xml.Name{Local: "a:blipFill"}}
		e.EncodeElement(m.BlipFill, seblipFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.GrpFill != nil {
		segrpFill := xml.StartElement{Name: xml.Name{Local: "a:grpFill"}}
		e.EncodeElement(m.GrpFill, segrpFill)
	}
	if m.Ln != nil {
		seln := xml.StartElement{Name: xml.Name{Local: "a:ln"}}
		e.EncodeElement(m.Ln, seln)
	}
	if m.EffectLst != nil {
		seeffectLst := xml.StartElement{Name: xml.Name{Local: "a:effectLst"}}
		e.EncodeElement(m.EffectLst, seeffectLst)
	}
	if m.EffectDag != nil {
		seeffectDag := xml.StartElement{Name: xml.Name{Local: "a:effectDag"}}
		e.EncodeElement(m.EffectDag, seeffectDag)
	}
	if m.Scene3d != nil {
		sescene3d := xml.StartElement{Name: xml.Name{Local: "a:scene3d"}}
		e.EncodeElement(m.Scene3d, sescene3d)
	}
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "a:sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ShapeProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "bwMode" {
			m.BwModeAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_ShapeProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "xfrm"}:
				m.Xfrm = NewCT_Transform2D()
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custGeom"}:
				m.CustGeom = NewCT_CustomGeometry2D()
				if err := d.DecodeElement(m.CustGeom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstGeom"}:
				m.PrstGeom = NewCT_PresetGeometry2D()
				if err := d.DecodeElement(m.PrstGeom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "noFill"}:
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "solidFill"}:
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gradFill"}:
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blipFill"}:
				m.BlipFill = NewCT_BlipFillProperties()
				if err := d.DecodeElement(m.BlipFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pattFill"}:
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "grpFill"}:
				m.GrpFill = NewCT_GroupFillProperties()
				if err := d.DecodeElement(m.GrpFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ln"}:
				m.Ln = NewCT_LineProperties()
				if err := d.DecodeElement(m.Ln, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectLst"}:
				m.EffectLst = NewCT_EffectList()
				if err := d.DecodeElement(m.EffectLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "effectDag"}:
				m.EffectDag = NewCT_EffectContainer()
				if err := d.DecodeElement(m.EffectDag, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scene3d"}:
				m.Scene3d = NewCT_Scene3D()
				if err := d.DecodeElement(m.Scene3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sp3d"}:
				m.Sp3d = NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ShapeProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ShapeProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ShapeProperties and its children
func (m *CT_ShapeProperties) Validate() error {
	return m.ValidateWithPath("CT_ShapeProperties")
}

// ValidateWithPath validates the CT_ShapeProperties and its children, prefixing error messages with path
func (m *CT_ShapeProperties) ValidateWithPath(path string) error {
	if err := m.BwModeAttr.ValidateWithPath(path + "/BwModeAttr"); err != nil {
		return err
	}
	if m.Xfrm != nil {
		if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
			return err
		}
	}
	if m.CustGeom != nil {
		if err := m.CustGeom.ValidateWithPath(path + "/CustGeom"); err != nil {
			return err
		}
	}
	if m.PrstGeom != nil {
		if err := m.PrstGeom.ValidateWithPath(path + "/PrstGeom"); err != nil {
			return err
		}
	}
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.BlipFill != nil {
		if err := m.BlipFill.ValidateWithPath(path + "/BlipFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.GrpFill != nil {
		if err := m.GrpFill.ValidateWithPath(path + "/GrpFill"); err != nil {
			return err
		}
	}
	if m.Ln != nil {
		if err := m.Ln.ValidateWithPath(path + "/Ln"); err != nil {
			return err
		}
	}
	if m.EffectLst != nil {
		if err := m.EffectLst.ValidateWithPath(path + "/EffectLst"); err != nil {
			return err
		}
	}
	if m.EffectDag != nil {
		if err := m.EffectDag.ValidateWithPath(path + "/EffectDag"); err != nil {
			return err
		}
	}
	if m.Scene3d != nil {
		if err := m.Scene3d.ValidateWithPath(path + "/Scene3d"); err != nil {
			return err
		}
	}
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
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
