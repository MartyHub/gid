package gid

import (
	"strings"
	"unicode"
)

type (
	Option    func(Tokenizer)
	Tokenizer struct {
		replacements map[string]string
		reserved     map[string]struct{}
	}
)

func Default() Tokenizer {
	return New(WithDefaultReplacements(), WithDefaultReserved())
}

func New(opts ...Option) Tokenizer {
	result := Tokenizer{
		replacements: make(map[string]string),
		reserved:     make(map[string]struct{}),
	}

	for _, opt := range opts {
		opt(result)
	}

	return result
}

func WithDefaultReplacements() Option {
	return func(tokenizer Tokenizer) {
		for _, s := range []string{
			"Acl", "Api", "Ascii",
			"Cpu", "Css",
			"Dns",
			"Eof",
			"Guid",
			"Html", "Http", "Https",
			"Id", "Ip",
			"Json",
			"Lhs",
			"Qps",
			"Ram", "Rhs", "Rpc",
			"Sla", "Smtp", "Sql", "Ssh",
			"Tcp", "Tls", "Ttl",
			"Udp", "Ui", "Uid", "Uuid", "Uri", "Url", "Utf8",
			"Vm",
			"Xml", "Xmpp", "Xsrf", "Xss",
		} {
			tokenizer.addReplacement(s)
		}
	}
}

func WithDefaultReserved() Option {
	return func(tokenizer Tokenizer) {
		for _, s := range []string{
			"any", "append",
			"bool", "break", "byte",
			"cap", "case", "chan", "clear", "close", "comparable",
			"complex", "complex64", "complex128",
			"const", "continue", "copy",
			"default", "defer", "delete",
			"else", "error",
			"fallthrough", "false", "float32", "float64", "for", "func",
			"go", "goto",
			"if", "imag", "import", "int", "int8", "int16", "int32", "int64", "interface", "iota",
			"len",
			"make", "map", "max", "min",
			"new", "nil",
			"package", "panic", "print", "println",
			"range", "real", "recover", "return", "rune",
			"select", "string", "struct", "switch",
			"true", "type",
			"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
			"var",
		} {
			tokenizer.addReserved(s)
		}
	}
}

func WithReplacement(s string) Option {
	return func(tokenizer Tokenizer) {
		tokenizer.addReplacement(s)
	}
}

func WithReserved(s string) Option {
	return func(tokenizer Tokenizer) {
		tokenizer.addReserved(s)
	}
}

func (tokenizer Tokenizer) ExportID(s string) string {
	return tokenizer.ToCamel(s, true)
}

func (tokenizer Tokenizer) UnexportID(s string) string {
	return tokenizer.ToCamel(s, false)
}

func (tokenizer Tokenizer) ToCamel(s string, export bool) string {
	prevRep := false
	sb := new(strings.Builder)

	for _, tok := range tokenizer.Tokens(s) {
		if !tok.Valid() {
			continue
		}

		elem := string(tok.Runes)

		if sb.Len() == 0 && tok.Class == ClassDigit {
			if export {
				elem = "a" + elem
			} else {
				elem = "_" + elem
			}
		}

		if export || sb.Len() > 0 {
			elem = Capitalize(elem)

			if !prevRep {
				elem, prevRep = tokenizer.replace(elem)
			}
		}

		sb.WriteString(elem)
	}

	return tokenizer.toCamel(sb, export)
}

func (tokenizer Tokenizer) Tokens(s string) []*Token {
	tokens := make([]*Token, 0, 1)

	for _, r := range s {
		var class TokenClass

		switch {
		case unicode.IsLower(r):
			class = ClassLowerCase
		case unicode.IsUpper(r):
			class = ClassUpperCase
			r = unicode.ToLower(r)
		case unicode.IsDigit(r):
			class = ClassDigit
		}

		if len(tokens) == 0 {
			tokens = append(tokens, NewToken(class, r))

			continue
		}

		if tokens[len(tokens)-1].Class == class {
			tokens[len(tokens)-1].Append(r)

			continue
		}

		if class == ClassLowerCase && tokens[len(tokens)-1].Class == ClassUpperCase {
			tokens = append(tokens, NewToken(class, tokens[len(tokens)-1].Pop(), r))

			continue
		}

		tokens = append(tokens, NewToken(class, r))
	}

	return tokens
}

func (tokenizer Tokenizer) addReplacement(s string) {
	tokenizer.replacements[s] = strings.ToUpper(s)
}

func (tokenizer Tokenizer) addReserved(s string) {
	tokenizer.reserved[s] = struct{}{}
}

func (tokenizer Tokenizer) replace(s string) (string, bool) {
	if replacement, found := tokenizer.replacements[s]; found {
		return replacement, true
	}

	return s, false
}

func (tokenizer Tokenizer) isReserved(s string) bool {
	_, found := tokenizer.reserved[s]

	return found
}

func (tokenizer Tokenizer) toCamel(sb *strings.Builder, export bool) string {
	if sb.Len() == 0 {
		if export {
			return "A"
		}

		return "a"
	}

	result := sb.String()

	if !export && tokenizer.isReserved(result) {
		return "_" + result
	}

	return result
}
