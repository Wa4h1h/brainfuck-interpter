package lexer


func New(input []byte)*Lexer{
	return &Lexer{
		Input: input,
	}
}

func (l *Lexer) NextToken() TokenReader {
	switch l.Ch{
	case INC:
	case DEC:
	case IN:
	case OUT:
	case LJUMP:
	case RJUMP:
	case LMOVE:
	case RMOVE:
	default:
		return nil
	}
}
