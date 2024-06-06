package tree_sitter_gowork_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-gowork"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_gowork.Language())
	if language == nil {
		t.Errorf("Error loading Gowork grammar")
	}
}
