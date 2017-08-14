package parser

type Document struct {
	BlockNodes []BlockNode
}

type BlockNode interface{}

type Heading struct {
	Level       int8
	InlineNodes []InlineNode
}

type Paragraph struct {
	InlineNodes []InlineNode
}

type InlineNode interface{}

type Text struct {
	Literal string
}
