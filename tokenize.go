package prose

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var notSplitSingleQuote = []string{"'ll", "'s", "'re", "'m", "'d", "'ve", "n't"}

var noNeedsplitJoiner = []string{"-", ".", "'"}

// iterTokenizer splits a sentence into words.
type iterTokenizer struct {
}

// newIterTokenizer is a iterTokenizer constructor.
func newIterTokenizer() *iterTokenizer {
	return new(iterTokenizer)
}

func addToken(s string, toks []*Token) []*Token {
	if strings.TrimSpace(s) != "" {
		toks = append(toks, &Token{Text: s})
	}
	return toks
}

// check
func checkIsNeedSplitMiddle(noneWord string) bool {
	//if hasAnyIndex(lower, noNeedsplitJoiner) != -1 {
	if noneWord == `'` || noneWord == `.` || noneWord == `-` {
		return false
	}
	return true
}
func isSpecial(token string) bool {
	_, found := emoticons[token]
	// return found || internalRE.MatchString(token)
	// 这里取消大写开头不拆分的限制
	return found
}

func doSplit(token string) []*Token {
	tokens := []*Token{}
	suffs := []*Token{}
	prefixHyphenReg := regexp.MustCompile(`^\W\S+$`)
	suffixHyphenReg := regexp.MustCompile(`^\S+\W$`)
	beginHyphenReg := regexp.MustCompile(`^\W+`)
	middleHyphenReg := regexp.MustCompile(`\W+`)
	endHyphenReg := regexp.MustCompile(`\W+$`)
	//beginSingleQuoteReg := regexp.MustCompile(`^'\S+$`)
	noneWordReg := regexp.MustCompile(`\W`)

	last := 0
	for token != "" && utf8.RuneCountInString(token) != last {
		if isSpecial(token) {
			// We've found a special case (e.g., an emoticon) -- so, we add it as a token without
			// any further processing.
			tokens = addToken(token, tokens)
			break
		}
		last = utf8.RuneCountInString(token)
		lower := strings.ToLower(token)
		if noneWordReg.MatchString(token) {
			if suffixHyphenReg.MatchString(token) { // 单纯后面有非字母的 big- -> ["big", "-"]
				iList := endHyphenReg.FindStringIndex(token)
				if len(iList) > 0 {
					i := iList[0]
					suffs = append([]*Token{
						{Text: string(token[i:])}},
						suffs...)
					token = token[0:i]
				}
			} else if prefixHyphenReg.MatchString(token) && !stringInSlice(lower, notSplitSingleQuote) {
				//前面有非字母的拆开, 满足单词缩写的这里不要拆, 由后面的去处理
				hyphen := beginHyphenReg.FindString(token)
				i := len(hyphen)
				if i > 0 {
					tokens = addToken(token[0:i], tokens)
					token = token[i:len(token)]
				}
			} else if idx := middleHyphenReg.FindStringIndex(lower)[0]; idx > -1 && checkIsNeedSplitMiddle(string(token[idx])) && len(lower) > 1 { // bigzhu/hah -> [bigzhu, /,hah]
				tokens = addToken(token[:idx], tokens)
				token = token[idx:]
			} else if idx := hasAnyIndex(lower, notSplitSingleQuote); idx > -1 { // 满足缩写词的拆开 they'll -> [they, 'll].
				tokens = addToken(token[:idx], tokens)
				token = token[idx:]
			} else {
				// 文字中间有非字母字符保持不动 bigzhu.com
				tokens = addToken(token, tokens)
			}
		} else {
			tokens = addToken(token, tokens)
		}
	}

	return append(tokens, suffs...)
}

// tokenize splits a sentence into a slice of words.
func (t *iterTokenizer) tokenize(text string) []*Token {
	tokens := []*Token{}

	clean, white := sanitizer.Replace(text), false
	length := len(clean)

	start, index := 0, 0
	cache := map[string][]*Token{}
	for index <= length {
		uc, size := utf8.DecodeRuneInString(clean[index:])

		if size == 0 {
			break
		} else if index == 0 {
			white = unicode.IsSpace(uc)
		}
		if unicode.IsSpace(uc) != white {
			if start < index {
				span := clean[start:index]
				if toks, found := cache[span]; found {
					tokens = append(tokens, toks...)
				} else {
					toks := doSplit(span)
					cache[span] = toks
					tokens = append(tokens, toks...)
				}
			}
			if uc == ' ' {
				start = index + 1
			} else {
				start = index
			}
			white = !white
		}
		index += size
		// 如果是换行符, 就填充进去, 不要丢
		if uc == rune('\n') {
			tokens = append(tokens, &Token{Text: "\n"})
		}
	}

	if start < index {
		tokens = append(tokens, doSplit(clean[start:index])...)
	}

	return tokens
}

//var internalRE = regexp.MustCompile(`^(?:[A-Za-z]\.){2,}$|^[A-Z][a-z]{1,2}\.$`)
var sanitizer = strings.NewReplacer(
	//"……", "… …", // 两个连续的省略符号会导致程序崩溃, 只有在这里替换分开
	"\u201c", `"`,
	"\u201d", `"`,
	"\u2018", "'",
	"—", "-",
	"…", "...",
	"\u2019", "'",
	"&rsquo;", "'")

//var suffixes = []string{}
//var prefixes = []string{}

var emoticons = map[string]int{
	"(-8":         1,
	"(-;":         1,
	"(-_-)":       1,
	"(._.)":       1,
	"(:":          1,
	"(=":          1,
	"(o:":         1,
	"(¬_¬)":       1,
	"(ಠ_ಠ)":       1,
	"(╯°□°）╯︵┻━┻": 1,
	"-__-":        1,
	"8-)":         1,
	"8-D":         1,
	"8D":          1,
	":(":          1,
	":((":         1,
	":(((":        1,
	":()":         1,
	":)))":        1,
	":-)":         1,
	":-))":        1,
	":-)))":       1,
	":-*":         1,
	":-/":         1,
	":-X":         1,
	":-]":         1,
	":-o":         1,
	":-p":         1,
	":-x":         1,
	":-|":         1,
	":-}":         1,
	":0":          1,
	":3":          1,
	":P":          1,
	":]":          1,
	":`(":         1,
	":`)":         1,
	":`-(":        1,
	":o":          1,
	":o)":         1,
	"=(":          1,
	"=)":          1,
	"=D":          1,
	"=|":          1,
	"@_@":         1,
	"O.o":         1,
	"O_o":         1,
	"V_V":         1,
	"XDD":         1,
	"[-:":         1,
	"^___^":       1,
	"o_0":         1,
	"o_O":         1,
	"o_o":         1,
	"v_v":         1,
	"xD":          1,
	"xDD":         1,
	"¯\\(ツ)/¯":    1,
}
