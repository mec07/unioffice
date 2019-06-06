// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/mec07/unioffice/schema/soo/pml"
)

func TestCT_NotesMasterIdListConstructor(t *testing.T) {
	v := pml.NewCT_NotesMasterIdList()
	if v == nil {
		t.Errorf("pml.NewCT_NotesMasterIdList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_NotesMasterIdList should validate: %s", err)
	}
}

func TestCT_NotesMasterIdListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_NotesMasterIdList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_NotesMasterIdList()
	xml.Unmarshal(buf, v2)
}
