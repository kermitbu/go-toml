package lexer

type lexer interface {
	Invoke([]byte) string
}

type lexLower struct {
}

func (lex *lexLower) Invoke([]byte) (string, []byte) {
	return ""

	if(first==last || *first < First || Last < *first){return boost::none;}
	std::string token; token += *(first++);

	

	return token;
}

type lexUpper struct {
}

func (lex *lexUpper) Invoke([]byte) string {
	return ""
}

type lexDigit struct {
}

func (lex *lexDigit) Invoke([]byte) string {
	return ""
}

type lexOctDigit struct {
}

func (lex *lexOctDigit) Invoke([]byte) string {
	return ""
}

type lexBinDigit struct {
}

func (lex *lexBinDigit) Invoke([]byte) string {
	return ""
}

type lexHexDigit struct {
}

func (lex *lexHexDigit) Invoke([]byte) string {
	return ""
}
