package tomlx

// func is_digit(ch) {
//     return /[0-9]/i.test(ch);
// }
// func is_id_start(ch) {
//     return /[a-zλ_]/i.test(ch);
// }
// func is_id(ch) {
// 	return is_id_start(ch) || "?!-<>=0123456789".indexOf(ch) >= 0
// }
// func is_op_char(ch) {
// 	return "+-*/%=&|<>!".indexOf(ch) >= 0
// }

// 是否是符号
func isPunc(ch byte) bool {
	all_punc := []byte("[]")

	for punc := range all_punc {
		if punc == ch {
			return true
		}
	}
	return false
}
func isWhitespace(ch byte) bool {
	return ch == '\t' || ch == ' ' || ch == '\n'
}

type Predicate func(ch byte) bool

func readWhile(predicate Predicate) string {
	var str = ""
	while(!input.eof() && predicate(input.peek()))
	str += input.next()
	return str
}

// func read_number() {
//     var has_dot = false;
//     var number = read_while(func(ch){
//         if (ch == ".") {
//             if (has_dot) return false;
//             has_dot = true;
//             return true;
//         }
//         return is_digit(ch);
//     });
//     return { type: "num", value: parseFloat(number) };
// }
// func read_ident() {
//     var id = read_while(is_id);
//     return {
//         type  : is_keyword(id) ? "kw" : "var",
//         value : id
//     };
// }
// func read_escaped(end) {
//     var escaped = false, str = "";
//     input.next();
//     while (!input.eof()) {
//         var ch = input.next();
//         if (escaped) {
//             str += ch;
//             escaped = false;
//         } else if (ch == "\\") {
//             escaped = true;
//         } else if (ch == end) {
//             break;
//         } else {
//             str += ch;
//         }
//     }
//     return str;
// }
// func read_string() {
//     return { type: "str", value: read_escaped('"') };
// }
// func skip_comment() {
//     read_while(func(ch){ return ch != "\n" });
//     input.next();
// }
// func read_next() {
//     read_while(is_whitespace);
//     if (input.eof()) return null;
//     var ch = input.peek();
//     if (ch == "#") {
//         skip_comment();
//         return read_next();
//     }
//     if (ch == '"') return read_string();
//     if (is_digit(ch)) return read_number();
//     if (is_id_start(ch)) return read_ident();
//     if (is_punc(ch)) return {
//         type  : "punc",
//         value : input.next()
//     };
//     if (is_op_char(ch)) return {
//         type  : "op",
//         value : read_while(is_op_char)
//     };
//     input.croak("Can't handle character: " + ch);
// }
// func peek() {
//     return current || (current = read_next());
// }
// func next() {
//     var tok = current;
//     current = null;
//     return tok || read_next();
// }
// func eof() {
//     return peek() == null;
// }
