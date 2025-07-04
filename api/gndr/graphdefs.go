// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package gndr

// schema: gndr.app.graph.defs

import (
	comatprototypes "github.com/gander-social/gander-indigo-sovereign/api/atproto"
	"github.com/gander-social/gander-indigo-sovereign/lex/util"
)

// GraphDefs_ListItemView is a "listItemView" in the gndr.app.graph.defs schema.
type GraphDefs_ListItemView struct {
	Subject *ActorDefs_ProfileView `json:"subject" cborgen:"subject"`
	Uri     string                 `json:"uri" cborgen:"uri"`
}

// GraphDefs_ListView is a "listView" in the gndr.app.graph.defs schema.
//
// RECORDTYPE: GraphDefs_ListView
type GraphDefs_ListView struct {
	LexiconTypeID     string                             `json:"$type,const=gndr.app.graph.defs#listView" cborgen:"$type,const=gndr.app.graph.defs#listView"`
	Avatar            *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Cid               string                             `json:"cid" cborgen:"cid"`
	Creator           *ActorDefs_ProfileView             `json:"creator" cborgen:"creator"`
	Description       *string                            `json:"description,omitempty" cborgen:"description,omitempty"`
	DescriptionFacets []*RichtextFacet                   `json:"descriptionFacets,omitempty" cborgen:"descriptionFacets,omitempty"`
	IndexedAt         string                             `json:"indexedAt" cborgen:"indexedAt"`
	Labels            []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	ListItemCount     *int64                             `json:"listItemCount,omitempty" cborgen:"listItemCount,omitempty"`
	Name              string                             `json:"name" cborgen:"name"`
	Purpose           *string                            `json:"purpose" cborgen:"purpose"`
	Uri               string                             `json:"uri" cborgen:"uri"`
	Viewer            *GraphDefs_ListViewerState         `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// GraphDefs_ListViewBasic is a "listViewBasic" in the gndr.app.graph.defs schema.
type GraphDefs_ListViewBasic struct {
	Avatar        *string                            `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Cid           string                             `json:"cid" cborgen:"cid"`
	IndexedAt     *string                            `json:"indexedAt,omitempty" cborgen:"indexedAt,omitempty"`
	Labels        []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	ListItemCount *int64                             `json:"listItemCount,omitempty" cborgen:"listItemCount,omitempty"`
	Name          string                             `json:"name" cborgen:"name"`
	Purpose       *string                            `json:"purpose" cborgen:"purpose"`
	Uri           string                             `json:"uri" cborgen:"uri"`
	Viewer        *GraphDefs_ListViewerState         `json:"viewer,omitempty" cborgen:"viewer,omitempty"`
}

// GraphDefs_ListViewerState is a "listViewerState" in the gndr.app.graph.defs schema.
type GraphDefs_ListViewerState struct {
	Blocked *string `json:"blocked,omitempty" cborgen:"blocked,omitempty"`
	Muted   *bool   `json:"muted,omitempty" cborgen:"muted,omitempty"`
}

// GraphDefs_NotFoundActor is a "notFoundActor" in the gndr.app.graph.defs schema.
//
// indicates that a handle or DID could not be resolved
//
// RECORDTYPE: GraphDefs_NotFoundActor
type GraphDefs_NotFoundActor struct {
	LexiconTypeID string `json:"$type,const=gndr.app.graph.defs#notFoundActor" cborgen:"$type,const=gndr.app.graph.defs#notFoundActor"`
	Actor         string `json:"actor" cborgen:"actor"`
	NotFound      bool   `json:"notFound" cborgen:"notFound"`
}

// GraphDefs_Relationship is a "relationship" in the gndr.app.graph.defs schema.
//
// lists the bi-directional graph relationships between one actor (not indicated in the object), and the target actors (the DID included in the object)
//
// RECORDTYPE: GraphDefs_Relationship
type GraphDefs_Relationship struct {
	LexiconTypeID string `json:"$type,const=gndr.app.graph.defs#relationship" cborgen:"$type,const=gndr.app.graph.defs#relationship"`
	Did           string `json:"did" cborgen:"did"`
	// followedBy: if the actor is followed by this DID, contains the AT-URI of the follow record
	FollowedBy *string `json:"followedBy,omitempty" cborgen:"followedBy,omitempty"`
	// following: if the actor follows this DID, this is the AT-URI of the follow record
	Following *string `json:"following,omitempty" cborgen:"following,omitempty"`
}

// GraphDefs_StarterPackView is a "starterPackView" in the gndr.app.graph.defs schema.
type GraphDefs_StarterPackView struct {
	Cid                string                             `json:"cid" cborgen:"cid"`
	Creator            *ActorDefs_ProfileViewBasic        `json:"creator" cborgen:"creator"`
	Feeds              []*FeedDefs_GeneratorView          `json:"feeds,omitempty" cborgen:"feeds,omitempty"`
	IndexedAt          string                             `json:"indexedAt" cborgen:"indexedAt"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty" cborgen:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty" cborgen:"joinedWeekCount,omitempty"`
	Labels             []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	List               *GraphDefs_ListViewBasic           `json:"list,omitempty" cborgen:"list,omitempty"`
	ListItemsSample    []*GraphDefs_ListItemView          `json:"listItemsSample,omitempty" cborgen:"listItemsSample,omitempty"`
	Record             *util.LexiconTypeDecoder           `json:"record" cborgen:"record"`
	Uri                string                             `json:"uri" cborgen:"uri"`
}

// GraphDefs_StarterPackViewBasic is a "starterPackViewBasic" in the gndr.app.graph.defs schema.
//
// RECORDTYPE: GraphDefs_StarterPackViewBasic
type GraphDefs_StarterPackViewBasic struct {
	LexiconTypeID      string                             `json:"$type,const=gndr.app.graph.defs#starterPackViewBasic" cborgen:"$type,const=gndr.app.graph.defs#starterPackViewBasic"`
	Cid                string                             `json:"cid" cborgen:"cid"`
	Creator            *ActorDefs_ProfileViewBasic        `json:"creator" cborgen:"creator"`
	IndexedAt          string                             `json:"indexedAt" cborgen:"indexedAt"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty" cborgen:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty" cborgen:"joinedWeekCount,omitempty"`
	Labels             []*comatprototypes.LabelDefs_Label `json:"labels,omitempty" cborgen:"labels,omitempty"`
	ListItemCount      *int64                             `json:"listItemCount,omitempty" cborgen:"listItemCount,omitempty"`
	Record             *util.LexiconTypeDecoder           `json:"record" cborgen:"record"`
	Uri                string                             `json:"uri" cborgen:"uri"`
}
