package tomlx

import (
	"bytes"
	"fmt"
)

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

// Kind 解析错误
type Kind int

const (
	_        Kind = iota
	emptyTag      = iota
	booleanTag
	integerTag
	floatTag
	stringTag
	dateTag
	timeTag
	localDatetimeTag
	offsetDatetimeTag
	arrayTag
	tableTtag
	undefinedTag
)

type TomlValue interface {
	Value() interface{}
}

type BooleanValue struct {
}

func (value *BooleanValue) Value() {

}

type Table map[string]TomlValue

func Parse(reader *bytes.Reader) Table{
	fmt.Println("xxx")


	reader.Read
}

inline table parse(std::istream& is)
{
    const std::ios::pos_type beg = is.tellg();
    is.seekg(0, std::ios::end);
    const std::ios::pos_type end = is.tellg();
    const std::size_t fsize = end - beg;
    is.seekg(beg);

    std::vector<char> letters(fsize);
    is.read(letters.data(), fsize);

    std::vector<char>::const_iterator first = letters.begin(),
                                       last = letters.end();
    const detail::result<table, std::string> res(
            detail::parse_toml_file(first, last));
    if(!res)
    {
        throw std::runtime_error(res.unwrap_err());
    }
    else
    {
        return res.unwrap();
    }
}



// parse table body (key-value pairs until the iter hits the next [tablekey])
template<typename InputIterator>
result<table, std::string>
parse_ml_table(InputIterator& iter, const InputIterator last)
{
    const InputIterator first(iter);
    if(first == last)
    {
        return err(std::string("toml::detail::parse_ml_table: input is empty"));
    }

    typedef repeat<sequence<maybe<lex_ws>, sequence<maybe<lex_comment>, lex_newline> >, unlimited> skip_line;
    skip_line::invoke(iter, last);

    table tab;
    while(iter != last)
    {
        lex_ws::invoke(iter, last);
        const InputIterator bfr(iter);
        {
            const result<std::vector<key>, std::string> tabkey(parse_array_table_key(iter, last));
            if(tabkey || iter != bfr)
            {
                iter = bfr;
                return ok(tab);
            }
        }
        {
            const result<std::vector<key>, std::string> tabkey(parse_table_key(iter, last));
            if(tabkey || iter != bfr)
            {
                iter = bfr;
                return ok(tab);
            }
        }

        const result<std::pair<std::vector<key>, value>, std::string>
            kv(parse_key_value_pair(iter, last));
        if(kv)
        {
            const std::vector<key>& keys = kv.unwrap().first;
            const value&            val  = kv.unwrap().second;
            const result<boost::blank, std::string> inserted =
                insert_nested_key(tab, val, keys.begin(), keys.end());
            if(!inserted)
            {
                return err(inserted.unwrap_err());
            }
        }
        else
        {
            return err("toml::detail::parse_ml_table: invalid line appeared -> "
                + kv.unwrap_err());
        }
        skip_line::invoke(iter, last);
        // comment lines are skipped by the above function call.
        // However, if the file ends with comment without newline,
        // it might cause parsing error because skip_line matches
        // `comment + newline`, not `comment` itself. to skip the
        // last comment, call this one more time.
        lex_comment::invoke(iter, last);
    }
    return ok(tab);
}

template<typename InputIterator>
result<table, std::string>
parse_toml_file(InputIterator& iter, const InputIterator last)
{
    const InputIterator first(iter);
    if(first == last)
    {
        return err(std::string("toml::detail::parse_toml_file: input is empty"));
    }

    table data;
    {
        const result<table, std::string> tab = parse_ml_table(iter, last);
        if(tab) {data = tab.unwrap();}
        else    {return err(tab.unwrap_err());}
    }
    while(iter != last)
    {
        const InputIterator bfr(iter);
        {
            const result<std::vector<key>, std::string> tabkey(parse_array_table_key(iter, last));
            if(tabkey)
            {
                const result<table, std::string> tab = parse_ml_table(iter, last);
                if(!tab){return err(tab.unwrap_err());}
                const result<boost::blank, std::string> inserted(insert_nested_key(data, tab.unwrap(),
                        tabkey.unwrap().begin(), tabkey.unwrap().end(), true));
                if(!inserted) {return err(inserted.unwrap_err());}
                continue;
            }
        }
        {
            const result<std::vector<key>, std::string> tabkey(parse_table_key(iter, last));
            if(tabkey)
            {
                const result<table, std::string> tab = parse_ml_table(iter, last);
                if(!tab){return err(tab.unwrap_err());}
                const result<boost::blank, std::string> inserted(
                    insert_nested_key(data, tab.unwrap(),
                        tabkey.unwrap().begin(), tabkey.unwrap().end()));
                if(!inserted) {return err(inserted.unwrap_err());}
                continue;
            }
        }
        return err("toml::detail::parse_toml_file: " + current_line(bfr, last));
    }
    return ok(data);
}