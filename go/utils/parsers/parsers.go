package parsers

import . "github.com/shreyassanthu77/aoc/go/utils"

func Char(c byte) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		if state.Pos >= len(state.Input) {
			state.Success = false
			return state
		}
		if state.Input[state.Pos] == c {
			state.Res = string(c)
			state.Pos++
			return state
		}
		state.Success = false
		return state
	})
}

func Str(s string) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		for i := 0; i < len(s); i++ {
			if state.Pos >= len(state.Input) {
				state.Success = false
				return state
			}
			if state.Input[state.Pos] != s[i] {
				state.Success = false
				return state
			}
			state.Res = string(s[i])
			state.Pos++
		}
		return state
	})
}

func Many(p *Parser) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		if !state.Success {
			return state
		}
		res := []interface{}{}
		for state.Success {
			state = p.Transformer(state)
			if state.Success {
				res = append(res, state.Res)
			}
		}
		state.Success = true
		state.Res = res
		return state
	})
}

func Many1(p *Parser) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		if !state.Success {
			return state
		}

		res := Many(p).Transformer(state)

		if len(res.Res.([]interface{})) != 0 {
			return state
		}

		state.Success = false
		return state
	})
}

func Choice(parsers ...*Parser) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		for _, p := range parsers {
			res := p.Transformer(state)
			if res.Success {
				return res
			}
		}
		state.Success = false
		return state
	})
}

func Seq(parsers ...*Parser) *Parser {
	return NewParser(func(state *ParserState) *ParserState {
		res := []interface{}{}
		for _, p := range parsers {
			res = append(res, p.Transformer(state).Res)
			if !state.Success {
				return state
			}
		}
		state.Res = res
		return state
	})
}
