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

	"github.com/mec07/unioffice"
)

type CT_ExternalLinkChoice struct {
	ExternalBook *CT_ExternalBook
	DdeLink      *CT_DdeLink
	OleLink      *CT_OleLink
}

func NewCT_ExternalLinkChoice() *CT_ExternalLinkChoice {
	ret := &CT_ExternalLinkChoice{}
	return ret
}

func (m *CT_ExternalLinkChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ExternalBook != nil {
		seexternalBook := xml.StartElement{Name: xml.Name{Local: "ma:externalBook"}}
		e.EncodeElement(m.ExternalBook, seexternalBook)
	}
	if m.DdeLink != nil {
		seddeLink := xml.StartElement{Name: xml.Name{Local: "ma:ddeLink"}}
		e.EncodeElement(m.DdeLink, seddeLink)
	}
	if m.OleLink != nil {
		seoleLink := xml.StartElement{Name: xml.Name{Local: "ma:oleLink"}}
		e.EncodeElement(m.OleLink, seoleLink)
	}
	return nil
}

func (m *CT_ExternalLinkChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ExternalLinkChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "externalBook"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "externalBook"}:
				m.ExternalBook = NewCT_ExternalBook()
				if err := d.DecodeElement(m.ExternalBook, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "ddeLink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "ddeLink"}:
				m.DdeLink = NewCT_DdeLink()
				if err := d.DecodeElement(m.DdeLink, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "oleLink"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "oleLink"}:
				m.OleLink = NewCT_OleLink()
				if err := d.DecodeElement(m.OleLink, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ExternalLinkChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalLinkChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalLinkChoice and its children
func (m *CT_ExternalLinkChoice) Validate() error {
	return m.ValidateWithPath("CT_ExternalLinkChoice")
}

// ValidateWithPath validates the CT_ExternalLinkChoice and its children, prefixing error messages with path
func (m *CT_ExternalLinkChoice) ValidateWithPath(path string) error {
	if m.ExternalBook != nil {
		if err := m.ExternalBook.ValidateWithPath(path + "/ExternalBook"); err != nil {
			return err
		}
	}
	if m.DdeLink != nil {
		if err := m.DdeLink.ValidateWithPath(path + "/DdeLink"); err != nil {
			return err
		}
	}
	if m.OleLink != nil {
		if err := m.OleLink.ValidateWithPath(path + "/OleLink"); err != nil {
			return err
		}
	}
	return nil
}
