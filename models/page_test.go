package models_test

import (
	"testing"

	"github.com/pocketbase/pocketbase/models"
)

func TestPageTableName(t *testing.T) {
	t.Parallel()

	m := models.Page{}
	if m.TableName() != "_pages" {
		t.Fatalf("Unexpected table name, got %q", m.TableName())
	}
}
