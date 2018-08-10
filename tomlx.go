package tomlx

import "fmt"

type GenericValue struct {
}

// ParseFlags 解析标记
type ParseFlags int

const (
	_ ParseFlags = iota
	ParseNoFlags
	ParseDefaultFlags
	ParseInsituFlag
	ParseValidateEncodingFlag
	ParseIterativeFlag
	ParseStopWhenDoneFlag
	ParseFullPrecisionFlag
	ParseCommentsFlag
	ParseNumbersAsStringsFlag
	ParseTrailingCommasFlag
	ParseNanAndInfFlag
)

// ParseError 解析错误
type ParseError int

const (
	_ ParseError = iota
	// ParseErrorNone 无错误
	ParseErrorNone
	// ParseErrorDocumentEmpty 文档是空的。
	ParseErrorDocumentEmpty
	// ParseErrorDocumentRootNotSingular 文档的根后面不能有其它值。
	ParseErrorDocumentRootNotSingular
	// ParseErrorValueInvalid 不合法的值。
	ParseErrorValueInvalid
	// ParseErrorObjectMissName Object 成员缺少名字。
	ParseErrorObjectMissName
	// ParseErrorObjectMissColon Object 成员名字后缺少冒号。
	ParseErrorObjectMissColon
	// ParseErrorObjectMissCommaOrCurlyBracket Object 成员后缺少逗号或 }。
	ParseErrorObjectMissCommaOrCurlyBracket
	//ParseErrorArrayMissCommaOrSquareBracket Array 元素后缺少逗号或 ] 。
	ParseErrorArrayMissCommaOrSquareBracket
	//ParseErrorStringUnicodeEscapeInvalidHex String 中的 \\u 转义符后含非十六进位数字。
	ParseErrorStringUnicodeEscapeInvalidHex
	//ParseErrorStringUnicodeSurrogateInvalid String 中的代理对（surrogate pair）不合法。
	ParseErrorStringUnicodeSurrogateInvalid
	//ParseErrorStringEscapeInvalid String 含非法转义字符。
	ParseErrorStringEscapeInvalid
	//ParseErrorStringMissQuotationMark String 缺少关闭引号。
	ParseErrorStringMissQuotationMark
	//ParseErrorStringInvalidEncoding String 含非法编码。
	ParseErrorStringInvalidEncoding
	//ParseErrorNumberTooBig Number 的值太大，不能存储于 double。
	ParseErrorNumberTooBig
	//ParseErrorNumberMissFraction Number 缺少了小数部分。
	ParseErrorNumberMissFraction
	// ParseErrorNumberMissExponent Number 缺少了指数。
	ParseErrorNumberMissExponent
)

func Parse() {
	fmt.Println("xxx")
}
