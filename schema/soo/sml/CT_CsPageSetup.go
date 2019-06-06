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

	"github.com/mec07/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_CsPageSetup struct {
	// Paper Size
	PaperSizeAttr *uint32
	// Paper Height
	PaperHeightAttr *string
	// Paper Width
	PaperWidthAttr *string
	// First Page Number
	FirstPageNumberAttr *uint32
	// Orientation
	OrientationAttr ST_Orientation
	// Use Printer Defaults
	UsePrinterDefaultsAttr *bool
	// Black And White
	BlackAndWhiteAttr *bool
	// Draft
	DraftAttr *bool
	// Use First Page Number
	UseFirstPageNumberAttr *bool
	// Horizontal DPI
	HorizontalDpiAttr *uint32
	// Vertical DPI
	VerticalDpiAttr *uint32
	// Number Of Copies
	CopiesAttr *uint32
	IdAttr     *string
}

func NewCT_CsPageSetup() *CT_CsPageSetup {
	ret := &CT_CsPageSetup{}
	return ret
}

func (m *CT_CsPageSetup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PaperSizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperSize"},
			Value: fmt.Sprintf("%v", *m.PaperSizeAttr)})
	}
	if m.PaperHeightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperHeight"},
			Value: fmt.Sprintf("%v", *m.PaperHeightAttr)})
	}
	if m.PaperWidthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "paperWidth"},
			Value: fmt.Sprintf("%v", *m.PaperWidthAttr)})
	}
	if m.FirstPageNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "firstPageNumber"},
			Value: fmt.Sprintf("%v", *m.FirstPageNumberAttr)})
	}
	if m.OrientationAttr != ST_OrientationUnset {
		attr, err := m.OrientationAttr.MarshalXMLAttr(xml.Name{Local: "orientation"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.UsePrinterDefaultsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "usePrinterDefaults"},
			Value: fmt.Sprintf("%d", b2i(*m.UsePrinterDefaultsAttr))})
	}
	if m.BlackAndWhiteAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "blackAndWhite"},
			Value: fmt.Sprintf("%d", b2i(*m.BlackAndWhiteAttr))})
	}
	if m.DraftAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "draft"},
			Value: fmt.Sprintf("%d", b2i(*m.DraftAttr))})
	}
	if m.UseFirstPageNumberAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "useFirstPageNumber"},
			Value: fmt.Sprintf("%d", b2i(*m.UseFirstPageNumberAttr))})
	}
	if m.HorizontalDpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "horizontalDpi"},
			Value: fmt.Sprintf("%v", *m.HorizontalDpiAttr)})
	}
	if m.VerticalDpiAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "verticalDpi"},
			Value: fmt.Sprintf("%v", *m.VerticalDpiAttr)})
	}
	if m.CopiesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "copies"},
			Value: fmt.Sprintf("%v", *m.CopiesAttr)})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CsPageSetup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "paperSize" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.PaperSizeAttr = &pt
			continue
		}
		if attr.Name.Local == "paperWidth" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PaperWidthAttr = &parsed
			continue
		}
		if attr.Name.Local == "firstPageNumber" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.FirstPageNumberAttr = &pt
			continue
		}
		if attr.Name.Local == "orientation" {
			m.OrientationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "usePrinterDefaults" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UsePrinterDefaultsAttr = &parsed
			continue
		}
		if attr.Name.Local == "paperHeight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PaperHeightAttr = &parsed
			continue
		}
		if attr.Name.Local == "draft" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DraftAttr = &parsed
			continue
		}
		if attr.Name.Local == "useFirstPageNumber" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UseFirstPageNumberAttr = &parsed
			continue
		}
		if attr.Name.Local == "horizontalDpi" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.HorizontalDpiAttr = &pt
			continue
		}
		if attr.Name.Local == "verticalDpi" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.VerticalDpiAttr = &pt
			continue
		}
		if attr.Name.Local == "copies" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CopiesAttr = &pt
			continue
		}
		if attr.Name.Local == "blackAndWhite" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BlackAndWhiteAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CsPageSetup: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CsPageSetup and its children
func (m *CT_CsPageSetup) Validate() error {
	return m.ValidateWithPath("CT_CsPageSetup")
}

// ValidateWithPath validates the CT_CsPageSetup and its children, prefixing error messages with path
func (m *CT_CsPageSetup) ValidateWithPath(path string) error {
	if m.PaperHeightAttr != nil {
		if !sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(*m.PaperHeightAttr) {
			return fmt.Errorf(`%s/m.PaperHeightAttr must match '%s' (have %v)`, path, sharedTypes.ST_PositiveUniversalMeasurePatternRe, *m.PaperHeightAttr)
		}
	}
	if m.PaperHeightAttr != nil {
		if !sharedTypes.ST_UniversalMeasurePatternRe.MatchString(*m.PaperHeightAttr) {
			return fmt.Errorf(`%s/m.PaperHeightAttr must match '%s' (have %v)`, path, sharedTypes.ST_UniversalMeasurePatternRe, *m.PaperHeightAttr)
		}
	}
	if m.PaperWidthAttr != nil {
		if !sharedTypes.ST_PositiveUniversalMeasurePatternRe.MatchString(*m.PaperWidthAttr) {
			return fmt.Errorf(`%s/m.PaperWidthAttr must match '%s' (have %v)`, path, sharedTypes.ST_PositiveUniversalMeasurePatternRe, *m.PaperWidthAttr)
		}
	}
	if m.PaperWidthAttr != nil {
		if !sharedTypes.ST_UniversalMeasurePatternRe.MatchString(*m.PaperWidthAttr) {
			return fmt.Errorf(`%s/m.PaperWidthAttr must match '%s' (have %v)`, path, sharedTypes.ST_UniversalMeasurePatternRe, *m.PaperWidthAttr)
		}
	}
	if err := m.OrientationAttr.ValidateWithPath(path + "/OrientationAttr"); err != nil {
		return err
	}
	return nil
}
