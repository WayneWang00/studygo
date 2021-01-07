package main

import (
	"fmt"
	"log"
	"strconv"
)

//TODO:最后三个函数没有看明白
func main() {
	//formatBool()
	//parseBool()
	//appendBool()
	//parseFloat()
	//numError()
	//parseUint()
	//parseInt()
	//aToi()
	//formatFloat()
	//appendFloat()
	//formatUint()
	//formatInt()
	//iToa()
	//appendInt()
	//appendUint()
	//quote()
	//appendQuote()
	//quoteToASCII()
	//appendQuoteToASCII()
	//quoteRune()
	//appendQuoteRune()
	//quoteRuneToASCII()
	//appendQuoteRuneToASCII()
	//canBackquote()
	//unquoteChar()
	//unquote()
	isPrint()
}

//func parseInt() {
//	fmt.Println(strconv.ParseInt("FF", 16, 0))
//	// 255 <nil>
//	fmt.Println(strconv.ParseInt("0xFF", 16, 0))
//	// 0 strconv.ParseInt: parsing "0xFF": invalid syntax
//	fmt.Println(strconv.ParseInt("0xFF", 0, 0))
//	// 255 <nil>
//	fmt.Println(strconv.ParseInt("9", 10, 4))
//	// 7 strconv.ParseInt: parsing "9": value out of range
//}
//func parseFloat() {
//	s := "0.12345678901234567890"
//
//	f, err := strconv.ParseFloat(s, 32)
//	fmt.Println(f, err)          // 0.12345679104328156 <nil>
//	fmt.Println(float32(f), err) // 0.12345679 <nil>
//
//	f, err = strconv.ParseFloat(s, 64)
//	fmt.Println(f, err) // 0.12345678901234568 <nil>
//}
//func canBackQuote() {
//	for i := rune(0); i < utf8.MaxRune; i++ {
//		if !strconv.CanBackquote(string(i)) {
//			fmt.Printf("%q, ", i)
//		}
//	}
//	fmt.Println()
//
//	// 结果如下：
//	// '\x00', '\x01', '\x02', '\x03', '\x04', '\x05', '\x06', '\a', '\b', '\n', '\v', '\f', '\r', '\x0e', '\x0f', '\x10', '\x11', '\x12', '\x13', '\x14', '\x15', '\x16', '\x17', '\x18', '\x19', '\x1a', '\x1b', '\x1c', '\x1d', '\x1e', '\x1f', '`', '\u007f', '\ufeff',
//}
//func isPrint() {
//	var rnp, rng, rpng, rgnp []rune
//	const maxLen = 32
//	for i := rune(0); i < utf8.MaxRune; i++ {
//		if !strconv.IsPrint(i) { // 不可打印
//			if len(rnp) < maxLen {
//				rnp = append(rnp, i)
//			}
//			if strconv.IsGraphic(i) && len(rgnp) < maxLen { // 图形字符
//				rgnp = append(rgnp, i)
//			}
//		}
//		if !strconv.IsGraphic(i) { // 非图形字符
//			if len(rng) < maxLen {
//				rng = append(rng, i)
//			}
//			if strconv.IsPrint(i) && len(rpng) < maxLen { // 可打印
//				rpng = append(rpng, i)
//			}
//		}
//	}
//	fmt.Printf("不可打印字符    ：%q\n", rnp)
//	fmt.Printf("非图形字符      ：%q\n", rng)
//	fmt.Printf("不可打印图形字符：%q\n", rgnp)
//	fmt.Printf("可打印非图形字符：%q\n", rpng)
//
//	//结果如下：
//	// 不可打印字符    ：['\x00' '\x01' '\x02' '\x03' '\x04' '\x05' '\x06' '\a' '\b' '\t' '\n' '\v' '\f' '\r' '\x0e' '\x0f' '\x10' '\x11' '\x12' '\x13' '\x14' '\x15' '\x16' '\x17' '\x18' '\x19' '\x1a' '\x1b' '\x1c' '\x1d' '\x1e' '\x1f']
//	// 非图形字符      ：['\x00' '\x01' '\x02' '\x03' '\x04' '\x05' '\x06' '\a' '\b' '\t' '\n' '\v' '\f' '\r' '\x0e' '\x0f' '\x10' '\x11' '\x12' '\x13' '\x14' '\x15' '\x16' '\x17' '\x18' '\x19' '\x1a' '\x1b' '\x1c' '\x1d' '\x1e' '\x1f']
//	// 不可打印图形字符：['\u00a0' '\u1680' '\u2000' '\u2001' '\u2002' '\u2003' '\u2004' '\u2005' '\u2006' '\u2007' '\u2008' '\u2009' '\u200a' '\u202f' '\u205f' '\u3000']
//	// 可打印非图形字符：[]
//}
//func quote() {
//	s := "Hello\t世界！\n"
//	fmt.Println(s)                         // Hello	世界！（换行）
//	fmt.Println(strconv.Quote(s))          // "Hello\t世界！\n"
//	fmt.Println(strconv.QuoteToASCII(s))   // "Hello\t\u4e16\u754c\uff01\n"
//	fmt.Println(strconv.QuoteToGraphic(s)) // "Hello\t世界！\n"
//}
//func unQuote() {
//	s1 := "`Hello	世界！`"                 // 解析反引号字符串
//	s2 := `"Hello\t\u4e16\u754c\uff01"` // 解析双引号字符串
//	fmt.Println(strconv.Unquote(s1))    // Hello	世界！ <nil>
//	fmt.Println(strconv.Unquote(s2))    // Hello	世界！ <nil>
//	fmt.Println()
//	fmt.Println(strconv.UnquoteChar(`\u4e16\u754c\uff01`, 0))
//	// 19990 true \u754c\uff01 <nil>
//	fmt.Println(strconv.UnquoteChar(`\"abc\"`, '"'))
//	// 34 false abc\" <nil>
//}

//将布尔值转换成字符串true或flase
func formatBool() {
	v := true
	s := strconv.FormatBool(v)
	fmt.Printf("%T, %v\n", s, s) // string, true
}

//将字符串转换成布尔值
func parseBool() {
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s) // bool, true
	}
}

//根据b的值将"true"或'false”附加到dst并返回扩展的缓冲区
func appendBool() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b)) // bool:true
}

//将字符串解析为浮点数
func parseFloat() {
	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415927410125732
	}
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // float64, 3.1415926535
	}
}

//记录了转换过程中发生的错误信息
func numError() {
	str := "Not a number"
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		e := err.(*strconv.NumError)
		fmt.Println("Func:", e.Func) // Func: ParseFloat
		fmt.Println("Num:", e.Num)   // Num: Not a number
		fmt.Println("Err:", e.Err)   // Err: invalid syntax
		fmt.Println(err)             // strconv.ParseFloat: parsing "Not a number": invalid syntax
	}
}

//将字符串解析为整数，不支持正负号
func parseUint() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // uint64, 42
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s) // uint64, 42
	}
}

//将字符串解析为整数，支持正负号
func parseInt() {
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, -354634382
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s) // int64, -3546343826724305832
	}
	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}

//将字符串转换为十进制整数,即：ParseInt(s, 10, 0) 的简写
func aToi() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s) // int, 10
	}
}

//将浮点数 f 转换为字符串形式
func formatFloat() {
	v := 3.1415926535
	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32) // string, 3.1415927E+00
	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64) // string, 3.1415926535E+00
}

//等价于append(dst, FormatFloat(f, fmt, prec, bitSize)...)
func appendFloat() {
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32)) // float32:3.1415927E+00
	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64)) // float64:3.1415926535E+00
}

//将无符号整数转换为字符串形式
func formatUint() {
	v := uint64(42)
	s10 := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s10, s10) // string, 42
	s16 := strconv.FormatUint(v, 16)
	fmt.Printf("%T, %v\n", s16, s16) // string, 2a
}

//将整数转换为字符串形式
func formatInt() {
	v := int64(-42)
	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10) // string, -42
	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16) // string, -2a
}

//将整数转换为十进制字符串形式（即：FormatInt(i, 10) 的简写）
func iToa() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s) // string, 10
}

//等价于append(dst, FormatInt(I, base)...)
func appendInt() {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10)) // int (base 10):-42
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16)) // int (base 10):-42
}

//等价于append(dst, FormatUint(I, base)...)
func appendUint() {
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10)) // uint (base 10):42

	b16 := []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16)) // uint (base 16):2a
}

//返回一个双引号的Go字符串字面意思表示s
func quote() {
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s) // "\"Fran & Freddie's Diner\t☺\""
}

//等价于append(dst, Quote(s)...)
func appendQuote() {
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b)) // quote:"\"Fran & Freddie's Diner\""
}

//返回表示s的双引号的Go字符串文字
func quoteToASCII() {
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s) // "\"Fran & Freddie's Diner\t\u263a\""
}

//等价于append(dst, QuoteToASCII(s)...)
func appendQuoteToASCII() {
	b := []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b)) // quote (ascii):"\"Fran & Freddie's Diner\""
}

//返回一个单引号的Go字符字面意思表示r
func quoteRune() {
	s := strconv.QuoteRune('☺')
	fmt.Println(s) // '☺'
}

//等价于append(dst, QuoteRune(r)...)
func appendQuoteRune() {
	b := []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b)) // rune:'☺'
}

//返回一个单引号的Go字符字面意思表示r
func quoteRuneToASCII() {
	s := strconv.QuoteRuneToASCII('☺')
	fmt.Println(s) // '\u263a'
}

//等价于append(dst, QuoteRuneToASCII(r)...)
func appendQuoteRuneToASCII() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b)) // rune (ascii):'\u263a'
}

//返回字符串s是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串
func canBackquote() {
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺")) // true
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))   // false
}

/*
函数假设s是一个表示字符的go语法字符串，解析它并返回四个值：

value，表示一个rune值或者一个byte值
multibyte，表示value是否是一个多字节的utf-8字符
tail，表示字符串剩余的部分
err，表示可能存在的语法错误
quote参数为单引号时，函数认为单引号是语法字符，不接受未转义的单引号；双引号时，函数认为双引号是语法字符，不接受未转义的双引号；如果是零值，函数把单引号和双引号当成普通字符。
*/
func unquoteChar() {
	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v)) // value: "
	fmt.Println("multibyte:", mb)    // multibyte: false
	fmt.Println("tail:", t)          // tail: Fran & Freddie's Diner\"
}

//函数假设s是一个单引号、双引号、反引号包围的go语法字符串，解析它并返回它表示的值。（如果是单引号括起来的，函数会认为s是go字符字面值，返回一个单字符的字符串）
func unquote() {
	s, err := strconv.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err) // "", invalid syntax
	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err) // "The string must be either double-quoted", <nil>
	s, err = strconv.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)       // "or backquoted.", <nil>
	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	fmt.Printf("%q, %v\n", s, err)       // "☺", <nil>
	s, err = strconv.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err) // "", invalid syntax
}

//IsPrint返回一个字符是否是可打印的，和unicode.IsPrint一样，r必须是：字母（广义）、数字、标点、符号、ASCII空格。
func isPrint() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c) // true
	bel := strconv.IsPrint('\007')
	fmt.Println(bel) // false
}
