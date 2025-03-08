package lexer

type TokenReader interface {
	Eat(in []byte)
}

type Token byte

const (
	INC Token = '+'
	DEC Token='-'
	IN Token=','
	OUT Token='.'
	LJUMP Token='['
	RJUMP Token=']'
	LMOVE Token='<'
	RMOVE Token='>'
)

type Binary struct {
	TK         Token
	Occurrence int
}

type Jump struct {
	TK    Token
	Start int
	End   int
}

type Move struct {
	TK Token
}

type Lexer struct{
	Input []byte
	ReadPos int
	Ch byte
}