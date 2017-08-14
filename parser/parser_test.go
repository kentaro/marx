package parser

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	src string
	ast Document
}{
	{
		`@ Heading1
		
		paragraph1

		@@ Heading2

		paragraph2

		paragraph3`,
		Document{
			BlockNodes: []BlockNode{
				Heading{
					Level: 1,
					InlineNodes: []InlineNode{
						Text{
							Literal: "Heading1",
						},
					},
				},
				Paragraph{
					InlineNodes: []InlineNode{
						Text{
							Literal: "paragraph1",
						},
					},
				},
				Heading{
					Level: 2,
					InlineNodes: []InlineNode{
						Text{
							Literal: "Heading2",
						},
					},
				},
				Paragraph{
					InlineNodes: []InlineNode{
						Text{
							Literal: "paragraph2",
						},
					},
				},
				Paragraph{
					InlineNodes: []InlineNode{
						Text{
							Literal: "paragraph3",
						},
					},
				},
			},
		},
	},
}

func TestParse(t *testing.T) {
	for i, test := range tests {
		doc, _ := Parse(strings.NewReader(test.src))
		if !reflect.DeepEqual(doc, test.ast) {
			t.Errorf("case %d:\n\n%s\n\nactual:\t\t%#v\nexpected:\t%#v", i, test.src, doc, test.ast)
		}
	}
}
