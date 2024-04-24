package utils

type ParserState struct {
	Input   string
	Pos     int
	Success bool
	Res     interface{}
}

func NewParserState(input string) *ParserState {
	return &ParserState{
		Input:   input,
		Pos:     0,
		Success: true,
		Res:     "",
	}
}

type StateTransformer func(*ParserState) *ParserState

type Parser struct {
	Transformer StateTransformer
}

func NewParser(transformer StateTransformer) *Parser {
	return &Parser{
		Transformer: transformer,
	}
}

func (p *Parser) Run(input string) *ParserState {
	state := NewParserState(input)
	state = p.Transformer(state)
	return state
}
