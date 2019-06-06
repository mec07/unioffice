// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/mec07/unioffice/schema/soo/sml"
)

func TestCT_CalculatedItemConstructor(t *testing.T) {
	v := sml.NewCT_CalculatedItem()
	if v == nil {
		t.Errorf("sml.NewCT_CalculatedItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_CalculatedItem should validate: %s", err)
	}
}

func TestCT_CalculatedItemMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_CalculatedItem()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_CalculatedItem()
	xml.Unmarshal(buf, v2)
}
