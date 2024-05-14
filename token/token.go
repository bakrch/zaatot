package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	/* MISC */
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	/*  IDENTIFIERS, LITERALS  */
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"
	/*  OPERATORS  */
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="
	/*  DELIMITERS */
	COMMA     = ","
	COLON     = ":"
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACKET  = "["
	RBRACKET  = "]"
	LBRACE    = "{"
	RBRACE    = "}"
	/*  KEYWORDS  */
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	FOR      = "FOR"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

var keywords = map[string]TokenType{
	"if":     IF,
	"for":    FOR,
	"fn":     FUNCTION,
	"let":    LET,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
