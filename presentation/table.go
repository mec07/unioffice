// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentation

import (
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// Table is a table within a document.
type Table struct {
	s Slide
	x *pml.CT_Tbl
}

// X returns the inner wrapped XML type.
func (t Table) X() *pml.CT_Tbl {
	return t.x
}

// Properties returns the table properties.
func (t Table) Properties() TableProperties {
	if t.x.TblPr == nil {
		t.x.TblPr = pml.NewCT_TblPr()
	}
	return TableProperties{t.x.TblPr}
}

// AddRow adds a row to a table.
func (t Table) AddRow() Row {
	c := pml.NewEG_ContentRowContent()
	t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, c)
	tr := pml.NewCT_Row()
	c.Tr = append(c.Tr, tr)
	return Row{t.d, tr}
}

// InsertRowAfter inserts a row after another row
func (t Table) InsertRowAfter(r Row) Row {
	for i, rc := range t.x.EG_ContentRowContent {
		if len(rc.Tr) > 0 && r.X() == rc.Tr[0] {
			c := pml.NewEG_ContentRowContent()
			if len(t.x.EG_ContentRowContent) <= i+2 {
				return t.AddRow()
			}
			t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, nil)
			copy(t.x.EG_ContentRowContent[i+2:], t.x.EG_ContentRowContent[i+1:])
			t.x.EG_ContentRowContent[i+1] = c
			tr := pml.NewCT_Row()
			c.Tr = append(c.Tr, tr)
			return Row{t.d, tr}
		}
	}
	return t.AddRow()
}

// InsertRowBefore inserts a row before another row
func (t Table) InsertRowBefore(r Row) Row {
	for i, rc := range t.x.EG_ContentRowContent {
		if len(rc.Tr) > 0 && r.X() == rc.Tr[0] {
			c := pml.NewEG_ContentRowContent()
			t.x.EG_ContentRowContent = append(t.x.EG_ContentRowContent, nil)
			copy(t.x.EG_ContentRowContent[i+1:], t.x.EG_ContentRowContent[i:])
			t.x.EG_ContentRowContent[i] = c
			tr := pml.NewCT_Row()
			c.Tr = append(c.Tr, tr)
			return Row{t.d, tr}
		}
	}
	return t.AddRow()
}

// Rows returns the rows defined in the table.
func (t Table) Rows() []Row {
	ret := []Row{}
	for _, rc := range t.x.EG_ContentRowContent {
		for _, ctRow := range rc.Tr {
			ret = append(ret, Row{t.d, ctRow})
		}
		if rc.Sdt != nil && rc.Sdt.SdtContent != nil {
			for _, ctRow := range rc.Sdt.SdtContent.Tr {
				ret = append(ret, Row{t.d, ctRow})
			}
		}
	}
	return ret
}
