// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/mec07/unioffice/schema/soo/wml"
)

func TestEG_SectPrContentsConstructor(t *testing.T) {
	v := wml.NewEG_SectPrContents()
	if v == nil {
		t.Errorf("wml.NewEG_SectPrContents must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_SectPrContents should validate: %s", err)
	}
}

func TestEG_SectPrContentsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_SectPrContents()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_SectPrContents()
	xml.Unmarshal(buf, v2)
}
