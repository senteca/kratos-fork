package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/ory/kratos/schema"

	"github.com/ory/kratos/text"
	"github.com/ory/x/stringslice"
)

// swagger:model uiNodeType
type Type string

// swagger:model uiNodeGroup
type Group string

func (g Group) String() string {
	return string(g)
}

const (
	DefaultGroup          Group = "default"
	PasswordGroup         Group = "password"
	OpenIDConnectGroup    Group = "oidc"
	ProfileGroup          Group = "profile"
	RecoveryLinkGroup     Group = "link"
	VerificationLinkGroup Group = "link"

	Text   Type = "text"
	Input  Type = "input"
	Image  Type = "img"
	Anchor Type = "a"
)

// swagger:model uiNodes
type Nodes []*Node

// Node represents a flow's nodes
//
// Nodes are represented as HTML elements or their native UI equivalents. For example,
// a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.
//
// swagger:model uiNode
type Node struct {
	// The node's type
	//
	// Can be one of: text, input, img, a
	//
	// required: true
	Type Type `json:"type" faker:"-"`

	// Group specifies which group (e.g. password authenticator) this node belongs to.
	//
	// required: true
	Group Group `json:"group"`

	// The node's attributes.
	//
	// required: true
	// swagger:type uiNodeAttributes
	Attributes Attributes `json:"attributes" faker:"ui_node_attributes"`

	// The node's messages
	//
	// Contains error, validation, or other messages relevant to this node.
	//
	// required: true
	Messages text.Messages `json:"messages"`
}

// Used for en/decoding the Attributes field.
type jsonRawNode struct {
	Type       Type          `json:"type"`
	Group      Group         `json:"group"`
	Attributes Attributes    `json:"attributes"`
	Messages   text.Messages `json:"messages"`
}

func (n *Node) ID() string {
	return n.Attributes.ID()
}

func (n *Node) Reset() {
	n.Messages = nil
	n.Attributes.Reset()
}

func (n *Node) GetValue() interface{} {
	return n.Attributes.GetValue()
}

func (n Nodes) Find(id string) *Node {
	for _, nn := range n {
		if nn.ID() == id {
			return nn
		}
	}

	return nil
}

func (n Nodes) Reset(exclude ...string) {
	for k, nn := range n {
		nn.Messages = nil
		if !stringslice.Has(exclude, nn.ID()) {
			nn.Reset()
		}
		n[k] = nn
	}
}

func (n Nodes) ResetNodes(reset ...string) {
	for k, nn := range n {
		if stringslice.Has(reset, nn.ID()) {
			nn.Reset()
		}
		n[k] = nn
	}
}

func getStringSliceIndexOf(needle []string, haystack string) int {
	for k := range needle {
		if needle[k] == haystack {
			return k
		}
	}
	return -1
}

type sortOptions struct {
	orderByGroups   []string
	schemaRef       string
	keysInOrder     []string
	keysInOrderPost func([]string) []string
}

type SortOption func(*sortOptions)

func SortByGroups(orderByGroups []Group) func(*sortOptions) {
	return func(options *sortOptions) {
		options.orderByGroups = make([]string, len(orderByGroups))
		for k := range orderByGroups {
			options.orderByGroups[k] = string(orderByGroups[k])
		}
	}
}

func SortBySchema(schemaRef string) func(*sortOptions) {
	return func(options *sortOptions) {
		options.schemaRef = schemaRef
	}
}

func SortUseOrder(keysInOrder []string) func(*sortOptions) {
	return func(options *sortOptions) {
		options.keysInOrder = keysInOrder
	}
}

func SortUpdateOrder(f func([]string) []string) func(*sortOptions) {
	return func(options *sortOptions) {
		options.keysInOrderPost = f
	}
}

func (n Nodes) SortBySchema(opts ...SortOption) error {
	var o sortOptions
	for _, f := range opts {
		f(&o)
	}

	if o.schemaRef != "" {
		schemaKeys, err := schema.GetKeysInOrder(o.schemaRef)
		if err != nil {
			return err
		}

		for _, k := range schemaKeys {
			o.keysInOrder = append(o.keysInOrder, k)
		}
	}

	if o.keysInOrderPost != nil {
		o.keysInOrder = o.keysInOrderPost(o.keysInOrder)
	}

	getKeyPosition := func(node *Node) int {
		lastPrefix := len(o.keysInOrder)

		// Method should always be the last element in the list
		if node.Attributes.ID() == "method" {
			return len(n) + 1
		}

		for i, n := range o.keysInOrder {
			if strings.HasPrefix(node.ID(), n) {
				return i
			}
		}

		return lastPrefix
	}

	if len(o.orderByGroups) > 0 {
		// Sort by groups so that default is in front, then oidc, password, ...
		sort.Slice(n, func(i, j int) bool {
			a := string(n[i].Group)
			b := string(n[j].Group)
			return getStringSliceIndexOf(o.orderByGroups, a) < getStringSliceIndexOf(o.orderByGroups, b)
		})
	}

	sort.SliceStable(n, func(i, j int) bool {
		a := n[i]
		b := n[j]

		if a.Group == b.Group {
			pa, pb := getKeyPosition(a), getKeyPosition(b)
			if pa < pb {
				return true
			} else if pa > pb {
				return false
			}

			return fmt.Sprintf("%v", a.GetValue()) < fmt.Sprintf("%v", b.GetValue())
		}

		return false
	})

	return nil
}

// Remove removes one or more nodes by their IDs.
func (n *Nodes) Remove(ids ...string) {
	if n == nil {
		return
	}

	var r Nodes
	for k, v := range *n {
		var found bool
		for _, needle := range ids {
			if (*n)[k].ID() == needle {
				found = true
				break
			}
		}
		if !found {
			r = append(r, v)
		}
	}
	*n = r
}

// Upsert updates or appends a node.
func (n *Nodes) Upsert(node *Node) {
	if n == nil {
		*n = append(*n, node)
		return
	}

	for i := range *n {
		if (*n)[i].ID() == node.ID() {
			(*n)[i] = node
			return
		}
	}

	*n = append(*n, node)
}

// SetValueAttribute sets a node's attribute's value or returns false if no node is found.
func (n *Nodes) SetValueAttribute(id string, value interface{}) bool {
	for i := range *n {
		if (*n)[i].ID() == id {
			(*n)[i].Attributes.SetValue(value)
			return true
		}
	}
	return false
}

// Append appends a node.
func (n *Nodes) Append(node *Node) {
	*n = append(*n, node)
}

func (n *Node) UnmarshalJSON(data []byte) error {
	var attr Attributes
	switch t := gjson.GetBytes(data, "type").String(); Type(t) {
	case Text:
		attr = new(TextAttributes)
	case Input:
		attr = new(InputAttributes)
	case Anchor:
		attr = new(AnchorAttributes)
	case Image:
		attr = new(ImageAttributes)
	default:
		return fmt.Errorf("unexpected node type: %s", t)
	}

	var d jsonRawNode
	d.Attributes = attr
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&d); err != nil {
		return err
	}

	*n = Node(d)
	return nil
}

func (n *Node) MarshalJSON() ([]byte, error) {
	var t Type
	if n.Attributes != nil {
		switch n.Attributes.(type) {
		case *TextAttributes:
			t = Text
		case *InputAttributes:
			t = Input
		case *AnchorAttributes:
			t = Anchor
		case *ImageAttributes:
			t = Image
		default:
			return nil, errors.WithStack(fmt.Errorf("unknown node type: %T", n.Attributes))
		}
	}

	if n.Type == "" {
		n.Type = t
	} else if n.Type != t {
		return nil, errors.WithStack(fmt.Errorf("node type and node attributes mismatch: %T != %s", n.Attributes, n.Type))
	}

	return json.Marshal((*jsonRawNode)(n))
}
