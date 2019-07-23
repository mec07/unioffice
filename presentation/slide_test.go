package presentation_test

import (
	"testing"

	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/presentation"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestGetTables(t *testing.T) {
	ppt := presentation.New()
	slide := ppt.AddSlide()
	table := slide.AddTable()
	tables := slide.Tables()

	if len(tables) != 1 {
		t.Errorf("expected 1 table, got %d", len(tables))
		return
	}

	if table != tables[0] {
		t.Error("retrieved table != added table")
		return
	}

	tbl := document.New().AddTable().X()

	tc := table.AddRow().AddCell().X()
	elts := pml.NewEG_BlockLevelElts()
	tc.EG_BlockLevelElts = append(tc.EG_BlockLevelElts, elts)
	c := pml.NewEG_ContentBlockContent()
	elts.EG_ContentBlockContent = append(elts.EG_ContentBlockContent, c)
	c.Tbl = append(c.Tbl, tbl)

	tables = slide.Tables()
	if len(tables) < 2 {
		t.Errorf("nested table not enumerated. found %d, expected 2", len(tables))
	}
}
