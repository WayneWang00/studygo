package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//count()
	//contains()
	//containsAny()
	//containsRune()
	//index()
	//lastIndex()
	//indexRune()
	//indexAny()
	//lastIndexAny()
	//splitN()
	//splitAfterN()
	//split()
	//splitAfter()
	//fields()
	//fieldsFunc()
	//join()
	//hasPreFix()
	//hasSufFix()
	//mapStrings()
	//repeat()
	//toUpLowTitle()
	//toUpLowTitleSpecial()
	//title()
	//trimLeftFunc()
	//trimRightFunc()
	//trimFunc()
	//indexFunc()
	//lastIndexFunc()
	//trim()
	//trimLeft()
	//trimRight()
	//trimSpace()
	//trimPrefix()
	//trimSuffix()
	//replace()
	//equalFold()
	//readerLen()
	//readerRead()
	//readerReadAt()
	//readerReadByte()
	//readerUnreadByte()
	//readerReadRune()
	//readerUnreadRune()
	//readerSeek()
	readerWriteTo()
	//replaceReplace()
	//replaceWriteString()
}

//计算r字符串subst在s中的非重复个数
func count() {
	s := "Hello,世界!!!!!"
	n := strings.Count(s, "!")
	fmt.Println(n) // 5
	n = strings.Count(s, "!!")
	fmt.Println(n) // 2
	n = strings.Count(s, "")
	fmt.Println(n) //14
}

//判断s中是否含有字符串substr
func contains() {
	s := "Hello,世界!!!!!"
	b := strings.Contains(s, "!!")
	fmt.Println(b) // true
	b = strings.Contains(s, "!?")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true
}

//判断s中是否含有字符串chars中的任何一个字符
func containsAny() {
	s := "Hello,世界!"
	b := strings.ContainsAny(s, "abc")
	fmt.Println(b) // false
	b = strings.ContainsAny(s, "def")
	fmt.Println(b) // true
	b = strings.Contains(s, "")
	fmt.Println(b) // true
}

//判断s中是否包含字符r
func containsRune() {
	s := "Hello,世界!"
	b := strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
	b = strings.ContainsRune(s, 0)
	fmt.Println(b) // false
}

//返回字符串substr在s中第一次出现的位置
func index() {
	s := "Hello,世界!"
	i := strings.Index(s, "h")
	fmt.Println(i) // -1
	i = strings.Index(s, "!")
	fmt.Println(i) // 12
	i = strings.Index(s, "")
	fmt.Println(i) // 0
}

//返回字符串substr在s中最后一次出现的位置
func lastIndex() {
	s := "Hello,世界! Hello!"
	i := strings.LastIndex(s, "h")
	fmt.Println(i) // -1
	i = strings.LastIndex(s, "H")
	fmt.Println(i) // 14
	i = strings.LastIndex(s, "")
	fmt.Println(i) // 20
}

//返回字符r在s中第一次出现的位置
func indexRune() {
	s := "Hello,世界! Hello!"
	i := strings.IndexRune(s, '\n')
	fmt.Println(i) // -1
	i = strings.IndexRune(s, '界')
	fmt.Println(i) // 9
	i = strings.IndexRune(s, 0)
	fmt.Println(i) // -1
}

//返回字符串chars的任何一个字符在s中第一次出现的位置
func indexAny() {
	s := "Hello,世界! Hello!"
	i := strings.IndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.IndexAny(s, "dof")
	fmt.Println(i) // 4
	i = strings.IndexAny(s, "")
	fmt.Println(i) // -1
}

//返回字符串chars中的任何一个字符在s中最后一次出现的位置
func lastIndexAny() {
	s := "Hello,世界! Hello!"
	i := strings.LastIndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.LastIndexAny(s, "def")
	fmt.Println(i) // 15
	i = strings.LastIndexAny(s, "")
	fmt.Println(i) // -1
}

//以sep为分隔符，将s切分成多个子串，结果中不包含sep本身
func splitN() {
	s := "Hello, 世界! Hello!"
	ss := strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello," "世界! Hello!"]
	ss = strings.SplitN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.SplitN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]
}

//以sep为分隔符，将s切分成多个子串，结果中包含sep本身
func splitAfterN() {
	s := "Hello, 世界! Hello!"
	ss := strings.SplitAfterN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfterN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfterN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]
}

//以sep为分隔符，将s切分成多个子串，结果中不包含sep本身
func split() {
	s := "Hello, 世界! Hello!"
	ss := strings.Split(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.Split(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello" "世界! Hello!"]
	ss = strings.Split(s, "")
	fmt.Printf("%q\n", ss) // ["H" "e" "l" "l" "o" "," " " "世" "界" "!" " " "H" "e" "l" "l" "o" "!"]
}

//以sep为分隔符，将s切分成多个子串，结果中包含sep本身
func splitAfter() {
	s := "Hello, 世界! Hello!"
	ss := strings.SplitAfter(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfter(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfter(s, "")
	fmt.Printf("%q\n", ss) // ["H" "e" "l" "l" "o" "," " " "世" "界" "!" " " "H" "e" "l" "l" "o" "!"]
}

//以连续的空白字符为分隔符，将s切分成多个子串，结果中不包含空白字符本身
func fields() {
	s := "Hello, 世界! Hello!"
	ss := strings.Fields(s)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
}

//以一个或多个满足f(rune)的字符为分隔符，将s切分成多个子串，结果不包含分隔符本身
func fieldsFunc() {
	s := "C:\\Windows\\System32\\FileName"
	ss := strings.FieldsFunc(s, isSlash)
	fmt.Printf("%q\n", ss) // ["C:" "Windows" "System32" "FileName"]
}

func isSlash(r rune) bool {
	return r == '\\' || r == '/'
}

//将a中的子串连接成一个单独的字符串，子串之间用sep分隔
func join() {
	ss := []string{"Monday", "Tuesday", "Wednesday"}
	s := strings.Join(ss, "|")
	fmt.Println(s) // Monday|Tuesday|Wednesday
}

//判断字符串s是否以prefix开头
func hasPreFix() {
	s := "Hello 世界!"
	b := strings.HasPrefix(s, "hello")
	fmt.Println(b) // false
	b = strings.HasPrefix(s, "Hello")
	fmt.Println(b) // true
}

//判断字符串s是否以suffix结尾
func hasSufFix() {
	s := "Hello 世界!"
	b := strings.HasSuffix(s, "世界")
	fmt.Println(b) // false
	b = strings.HasSuffix(s, "世界!")
	fmt.Println(b) // true
}

//将s中满足mapping（rune）的字符替换成mapping（rune）的返回值。
func mapStrings() {
	s := "C:\\Windows\\System32\\FileName"
	ms := strings.Map(Slash, s)
	fmt.Printf("%q\n", ms) // "C:/Windows/System32/FileName"
}

func Slash(r rune) rune {
	if r == '\\' {
		return '/'
	}
	return r
}

//将count个字符串s连接成一个新的字符串。
func repeat() {
	s := "Hello!"
	rs := strings.Repeat(s, 3)
	fmt.Printf("%q\n", rs) // "Hello!Hello!Hello!"
}

//将s中的所有字符修改成大写、小写、Title格式
func toUpLowTitle() {
	s := "heLLo worLd Ａｂｃ"
	us := strings.ToUpper(s)
	ls := strings.ToLower(s)
	ts := strings.ToTitle(s)
	fmt.Printf("%q\n", us) // "HELLO WORLD ＡＢＣ "
	fmt.Printf("%q\n", ls) // "hello world ａｂｃ"
	fmt.Printf("%q\n", ts) // "HELLO WORLD ＡＢＣ"

	//获取非ASCII字符的Title格式
	for _, cr := range unicode.CaseRanges {
		// u := uint32(cr.Delta[unicode.UpperCase]) // 大写格式
		// l := uint32(cr.Delta[unicode.LowerCase]) // 小写格式
		t := uint32(cr.Delta[unicode.TitleCase]) // Title 格式
		// if t != 0 && t != u {
		if t != 0 {
			for i := cr.Lo; i <= cr.Hi; i++ {
				fmt.Printf("%c -> %c\n", i, i+t)
			}
		}
	}
}

//通过_case中的规则将s中所有的字符修改成其大写、小写、Title格式
func toUpLowTitleSpecial() {
	// 定义转换规则
	var _MyCase = unicode.SpecialCase{
		// 将半角逗号替换为全角逗号，ToTitle 不处理
		unicode.CaseRange{',', ',',
			[unicode.MaxCase]rune{'，' - ',', '，' - ',', 0}},
		// 将半角句号替换为全角句号，ToTitle 不处理
		unicode.CaseRange{'.', '.',
			[unicode.MaxCase]rune{'。' - '.', '。' - '.', 0}},
		// 将 ABC 分别替换为全角的 ＡＢＣ、ａｂｃ，ToTitle 不处理
		unicode.CaseRange{'A', 'C',
			[unicode.MaxCase]rune{'Ａ' - 'A', 'ａ' - 'A', 0}},
	}
	s := "ABCDEF,abcdef."
	us := strings.ToUpperSpecial(_MyCase, s)
	fmt.Printf("%q\n", us) // "ＡＢＣDEF，ABCDEF。"
	ls := strings.ToLowerSpecial(_MyCase, s)
	fmt.Printf("%q\n", ls) // "ａｂｃdef，abcdef。"
	ts := strings.ToTitleSpecial(_MyCase, s)
	fmt.Printf("%q\n", ts) // "ABCDEF,ABCDEF."
}

//将s中的所有单词的首字母修改为其Title格式
func title() {
	s := "heLLo worLd"
	ts := strings.Title(s)
	fmt.Printf("%q\n", ts) // "HeLLo WorLd"
}

//删除s头部连续的满足f（rune）的字符
func trimLeftFunc() {
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimLeftFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "HostName\\C\\Windows\\"
}

//删除s尾部连续的满足f（rune）的字符
func trimRightFunc() {
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimRightFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "\\\\HostName\\C\\Windows"
}

//删除s首尾连续的满足f（rune）的字符
func trimFunc() {
	s := "\\\\HostName\\C\\Windows\\"
	ts := strings.TrimFunc(s, isSlash)
	fmt.Printf("%q\n", ts) // "HostName\\C\\Windows"
	s1 := "abcd"
	ts1 := strings.TrimFunc(s1, isSlash)
	fmt.Printf("%q\n", ts1) // "abcd"
}

//返回s中第一个满足f（rune）的字符的字节位置
func indexFunc() {
	s := "C:\\Windows\\System32"
	i := strings.IndexFunc(s, isSlash)
	fmt.Printf("%v\n", i) // 2
	s1 := "abcd"
	i1 := strings.IndexFunc(s1, isSlash)
	fmt.Printf("%v\n", i1) // -1
}

//返回最后一个满足f（rune）的字符的字节位置
func lastIndexFunc() {
	s := "C:\\Windows\\System32"
	i := strings.LastIndexFunc(s, isSlash)
	fmt.Printf("%v\n", i) // 10
	s1 := "abcd"
	i1 := strings.LastIndexFunc(s1, isSlash)
	fmt.Printf("%v\n", i1) // -1
}

//删除s首尾连续的包含在cutset中的字符
func trim() {
	s := " Hello 世界! "
	ts := strings.Trim(s, " Helo!")
	fmt.Printf("%q\n", ts) // "世界"
}

//删除s头部连续的包含在cutset中的字符
func trimLeft() {
	s := " Hello 世界! "
	ts := strings.TrimLeft(s, " Helo!")
	fmt.Printf("%q\n", ts) // "世界! "
}

//删除s尾部连续的包含在cutset中的字符
func trimRight() {
	s := " Hello 世界! "
	ts := strings.TrimRight(s, " Helo!")
	fmt.Printf("%q\n", ts) // " Hello 世界"
}

//删除s首尾连续的空白字符
func trimSpace() {
	s := " Hello 世界! "
	ts := strings.TrimSpace(s)
	fmt.Printf("%q\n", ts) // "Hello 世界!"
}

//删除s头部的prefix字符串
func trimPrefix() {
	s := "Hello 世界!"
	ts := strings.TrimPrefix(s, "Hello")
	fmt.Printf("%q\n", ts) // " 世界!"
}

//删除s尾部的suffix字符串
func trimSuffix() {
	s := "Hello 世界!!!!!"
	ts := strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts) // "Hello 世界!"
}

//返回s的副本，并将副本中的old字符串换成new字符串
func replace() {
	s := "Hello 世界！"
	s = strings.Replace(s, " ", ",", -1)
	fmt.Println(s) // Hello,世界！
	s = strings.Replace(s, "", "|", -1)
	fmt.Println(s) // |H|e|l|l|o|,|世|界|！|
}

//判断s和t是否相等，忽略大小写，同时会对特殊字符进行转换
func equalFold() {
	s1 := "Hello 世界! ϕ Ǆ"
	s2 := "hello 世界! Φ ǅ"
	b := strings.EqualFold(s1, s2)
	fmt.Printf("%v\n", b) // true
}

//返回r.i之后的所有数据的字节长度
func readerLen() {
	s := "Hello 世界!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 获取字符串的编码长度
	fmt.Println(r.Len()) // 13
}

//将r.i之后的所有数据写入到b中（如果b足够大）
func readerRead() {
	s := "Hello World!"
	//s := "Hello世界!" //"Hello", "世\xe7\x95", "\x8c!",
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)
	// 循环读取 r 中的字符串
	for n, _ := r.Read(b); n > 0; n, _ = r.Read(b) {
		fmt.Printf("%q, ", b[:n]) // "Hello", " Worl", "d!",
	}
	fmt.Println()
}

//将off之后的所有数据写入到b中（如果b足够大）
func readerReadAt() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)
	// 读取 r 中指定位置的字符串
	n, _ := r.ReadAt(b, 0)
	fmt.Printf("%q\n", b[:n]) // "Hello"
	// 读取 r 中指定位置的字符串
	n, _ = r.ReadAt(b, 6)
	fmt.Printf("%q\n", b[:n]) // "World"
}

//将r.i之后的第一个字节写入到返回值b中
func readerReadByte() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'e', 'l',
	}
	fmt.Println()
}

//撤销上一次的ReadByte操作，即r.i--
func readerUnreadByte() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字节
	for i := 0; i < 3; i++ {
		b, _ := r.ReadByte()
		fmt.Printf("%q, ", b) // 'H', 'H', 'H',
		r.UnreadByte()        // 撤消前一次的字节读取操作
	}
	fmt.Println()
}

//将r.i之后的第一个字符写入到返回值ch中
func readerReadRune() {
	s := "你好 世界！"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, n, _ := r.ReadRune()
		fmt.Printf(`"%c:%v", `, b, n)
		// "你:3", "好:3", " :1", "世:3", "界:3",
	}
	fmt.Println()
}

//撤销上一次的ReadRune操作
func readerUnreadRune() {
	s := "你好 世界！"
	// 创建 Reader
	r := strings.NewReader(s)
	// 读取 r 中的一个字符
	for i := 0; i < 5; i++ {
		b, _, _ := r.ReadRune()
		fmt.Printf("%q, ", b)
		// '你', '你', '你', '你', '你',
		r.UnreadRune() // 撤消前一次的字符读取操作
	}
	fmt.Println()
}

//移动r中的索引位置
func readerSeek() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建读取缓冲区
	b := make([]byte, 5)
	// 读取 r 中指定位置的内容
	i2, err := r.Seek(6, 0)                // 移动索引位置到第 7 个字节
	fmt.Println("i2: ", i2, "err2: ", err) // i2:  6 err2:  <nil>
	r.Read(b)                              // 开始读取
	fmt.Printf("%q\n", b)                  // "World"
	i1, err := r.Seek(-5, 1)               // 将索引位置移回去
	fmt.Println("i1: ", i1, "err1: ", err) // i1:  6 err1:  <nil>
	r.Read(b)                              // 继续读取
	fmt.Printf("%q\n", b)                  // "World"
	i, err := r.Seek(-5, 3)
	fmt.Println("i: ", i, "err: ", err) // i:  0 err:  strings.Reader.Seek: invalid whence
}

//将r.i之后的数据写入到接口w中
func readerWriteTo() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建 bytes.Buffer 对象，它实现了 io.Writer 接口
	buf := bytes.NewBuffer(nil)
	// 将 r 中的数据写入 buf 中
	i, err := r.WriteTo(buf)
	fmt.Println("i: ", i, "err: ", err) // i:  12 err:  <nil>
	fmt.Printf("%q\n", buf)             // "Hello World!"
}

//返回对s进行“查找和替换”后的结果
func replaceReplace() {
	srp := strings.NewReplacer("Hello", "你好", "World", "世界", "!", "！")
	s := "Hello World!Hello World!hello world!"
	rst := srp.Replace(s)
	fmt.Print(rst) // 你好 世界！你好 世界！hello world！
	//方法二：

	/*
			wl := []string{"Hello", "Hi", "Hello", "你好"}
		    srp := strings.NewReplacer(wl...)
		    s := "Hello World! Hello World! hello world!"
		    rst := srp.Replace(s)
		    fmt.Print(rst) // Hi World! Hi World! hello world!
	*/
	fmt.Println()
}

//对s进行“检查和替换”，然后将结果写入w中
func replaceWriteString() {
	wl := []string{"Hello", "你好", "World", "世界", "!", "！"}
	srp := strings.NewReplacer(wl...)
	s := "Hello World!Hello World!hello world!"
	srp.WriteString(os.Stdout, s)
	// 你好 世界！你好 世界！hello world！
}
