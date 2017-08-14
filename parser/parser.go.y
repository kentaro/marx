%{
package parser

%}

%union{
  token        Token
  document     Document
  block_nodes  []BlockNode
  block_node   BlockNode
  inline_nodes []InlineNode
  inline_node  InlineNode
  depth        int8
}

%token<token> TEXT
%token HEADING_MARKER CR

%type<document>     document
%type<block_nodes>  block_nodes
%type<block_node>   block_node heading paragraph
%type<inline_nodes> inline_nodes
%type<inline_node>  inline_node text
%type<depth>        heading_markers

%%

document:
    block_nodes
    {
        $$ = Document{BlockNodes: []BlockNode{$1}}
        yylex.(*Lexer).result = $$
    }

block_nodes:
    block_node
    {
        $$ = []BlockNode{$1}        
    }
    | block_nodes block_node
    {
        $$ = append($1, $2)
    }

block_node:
    heading
    {
        $$ = $1
    }
    | paragraph
    {
        $$ = $1        
    }

heading:
    heading_markers inline_nodes CR
    {
        $$ = Heading{Level: $1, InlineNodes: $2}        
    }

heading_markers:
    HEADING_MARKER
    {
        $$ = 1
    }
    | HEADING_MARKER heading_markers
    {
        $$ = $2 + 1
    }

paragraph:
    empty_line
    {
        $$ = Paragraph{InlineNodes: []InlineNode}        
    }
    |
    inline_nodes CR
    {
        $$ = Paragraph{InlineNodes: $1}
    }

empty_line: CR

inline_nodes:
    inline_node
    {
        $$ = []InlineNode{$1}
    }
    | inline_nodes inline_node
    {
        $$ = append($1, $2)
    }

inline_node:
    text

text:
    TEXT
    {
        $$ = Text{Literal: $1.literal}
    }

%%
