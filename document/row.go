// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"

type Row struct {
	d *Document
	x *wml.CT_Row
}

func (r Row) X() *wml.CT_Row {
	return r.x
}

func (r Row) AddCell() Cell {
	cc := wml.NewEG_ContentCellContent()
	r.x.EG_ContentCellContent = append(r.x.EG_ContentCellContent, cc)
	tc := wml.NewCT_Tc()
	cc.Tc = append(cc.Tc, tc)
	return Cell{r.d, tc}
}