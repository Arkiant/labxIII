package internal

import (
	"io"

	"github.com/Arkiant/labxIII/src/conversation"
	// TODO: Delete this dependency in internal package
)

// Intent represent a conversation intent
type Intent string

// We have 2 intents search and book
const (
	SEARCH Intent = "search"
	BOOK   Intent = "book"
)

// Dialog is a interface represent a dialog API NLP
type Dialog interface {
	Convert(io.ReadCloser) (*conversation.Criteria, error)
}
