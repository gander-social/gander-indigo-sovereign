// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.graph.listitem

import (
	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

func init() {
	util.RegisterType("gndr.app.graph.listitem", &GraphListitem{})
} //
// RECORDTYPE: GraphListitem
type GraphListitem struct {
	LexiconTypeID string `json:"$type,const=gndr.app.graph.listitem" cborgen:"$type,const=gndr.app.graph.listitem"`
	CreatedAt     string `json:"createdAt" cborgen:"createdAt"`
	// list: Reference (AT-URI) to the list record (gndr.app.graph.list).
	List string `json:"list" cborgen:"list"`
	// subject: The account which is included on the list.
	Subject string `json:"subject" cborgen:"subject"`
}
