package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//match()
	//glob()
	//clean()
	//toSlash()
	//fromSlash()
	//splitList()
	//split()
	//join()
	//ext()
	//evalSymlinks()
	//abs()
	//rel()
	walk()
}

/*
	match.go
*/

// 判断 name 是否和指定的模式 pattern 完全匹配。只有pattern语法错误时，会返回ErrBadPattern。
// pattern 规则如下：
// Windows下，不能进行转义：'\\'被视为路径分隔符。
// 可以使用 ? 匹配单个任意字符（不匹配路径分隔符）。
// 可以使用 * 匹配 0 个或多个任意字符（不匹配路径分隔符）。
// 可以使用 [] 匹配范围内的任意一个字符（可以包含路径分隔符）。
// 可以使用 [^] 匹配范围外的任意一个字符（无需包含路径分隔符）。
// [] 之内可以使用 - 表示一个区间，比如 [a-z] 表示 a-z 之间的任意一个字符。
// 反斜线用来匹配实际的字符，比如 \* 匹配 *，\[ 匹配 [，\a 匹配 a 等等。
// [] 之内可以直接使用 [ * ?，但不能直接使用 ] -，需要用 \]、\- 进行转义。
func match() {
	fmt.Println(filepath.Match(`???`, `abc`))          // true
	fmt.Println(filepath.Match(`???`, `abcd`))         // false
	fmt.Println(filepath.Match(`*`, `abc`))            // true
	fmt.Println(filepath.Match(`*`, ``))               // true
	fmt.Println(filepath.Match(`a*`, `abc`))           // true
	fmt.Println(filepath.Match(`???\\???`, `abc\def`)) // windows: false linux: true
	fmt.Println(filepath.Match(`???/???`, `abc/def`))  // true
	fmt.Println(filepath.Match(`/*/*/*/`, `/a/b/c/`))  // true
	fmt.Println(filepath.Match(`[aA][bB][cC]`, `aBc`)) // true
	fmt.Println(filepath.Match(`[^aA]*`, `abc`))       // false
	fmt.Println(filepath.Match(`[a-z]*`, `a+b`))       // true
	fmt.Println(filepath.Match(`\[*\]`, `[a+b]`))      // windows: false linux: true
	fmt.Println(filepath.Match(`[[\]]*[[\]]`, `[]`))   // windows: false linux: true
}

// 返回所有匹配模式匹配字符串pattern的文件、目录或nil（如果没有匹配的文件），匹配模式同上面的match一样
func glob() {
	//list, err := filepath.Glob("/usr/*/[Bb][Aa]*")
	list, err := filepath.Glob("D:WayneWang/新赛事/[a-z]*")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	if len(list) < 1 {
		fmt.Println("len(list):", len(list))
		return
	}
	for k, v := range list {
		fmt.Println(k, ":", v)
	}
}

/*
	path.go
*/

// 通过单纯的词法操作返回和path代表同一地址的最短路径
// 规则如下：
// 返回的路径只有其代表一个根地址时才以路径分隔符结尾。Linux的'/'，Windows的'C:\'。如果返回的结果为空字符串，则会返回"."
// 将连续的多个路径分隔符替换为单个路径分隔符
// 剔除每一个.路径名元素（代表当前目录）
// 剔除每一个路径内的..路径名元素（代表父目录）和它前面的非..路径名元素
// 剔除开始一个根路径的..路径名元素，即将路径开始处的"/.."替换为"/"（假设路径分隔符是'/'）
func clean() {
	s := filepath.Clean("/a/./b/..// /c///d///")
	//s := filepath.Clean("")
	//s := filepath.Clean("D:/")
	fmt.Println(s) // \a\ \c\d
}

// 将path中的路径分隔符换成"/"并返回替换结果
func toSlash() {
	//s := filepath.ToSlash("http://www.baidu.com/a/b/c")
	s := filepath.ToSlash(`D:WayneWang\test\a\\b\\\c`)
	fmt.Println(s) // D:WayneWang/test/a//b///c
}

// 将path中的"/"换成路径分隔符并返回替换结果
func fromSlash() {
	s := filepath.FromSlash("http://www.baidu.com/a//b///c")
	fmt.Println(s) // http:\\www.baidu.com\a\\b\\\c
}

// 将路径path里的多个路径分隔成多条独立的路径
// 如果是""，返回的是[]string{}
func splitList() {
	sList := filepath.SplitList(`a/b/c;d/e/f;   g/h/i`)
	//sList := filepath.SplitList(``)
	//sList := filepath.SplitList(`/a/b/c:/usr/bin`)
	fmt.Printf("%q", sList) // ["a/b/c" "d/e/f" "   g/h/i"]
}

// 将路径从最后一个路径分隔符后面位置分割成两部分（dir和file）并返回
// 如果路径中没有路径分隔符，返回值中dir为空字符串，file为path
func split() {
	dir, file := filepath.Split(`a/b//c///d`)
	fmt.Printf("%q %q\n", dir, file) // "a/b//c///" "d"
}

// 将任意数量的路径元素放入一个单一的路径，根据需要添加路径分隔符。
// 结果是经过简化的，所有的空字符串元素会被忽略
func join() {
	s := filepath.Join("a", "b", "", ";;;", "  ", "//c///d/")
	fmt.Println(s) // a\b\;;;\  \c\d
}

// 返回路径path的文件扩展名
// 返回的是路径最后一个路径元素的最后一个'.'起始的后缀（包括'.'），如果该元素没有'.'会返回空字符串
func ext() {
	path := filepath.FromSlash(`/a//b///c//d.go`)
	s := filepath.Ext(path)
	fmt.Println(s) // .go
}

// 返回path的实际路径（如果path是个软链接）
func evalSymlinks() {
	s, err := filepath.EvalSymlinks(`D:WayneWang\`)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(s) // D:WayneWang
}

// 返回path代表的绝对路径，如果path不是绝对路径，则加入当前工作目录成为绝对路径
func abs() {
	s, err := filepath.Abs(`a/b/d/d`)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(s)
}

// 返回一个相对路径，将basepath和该路径用路径分隔符连起来的新路径在词法上等价与targpath
// 要求 targpath 和 basepath 必须“都是相对路径”或“都是绝对路径”
func rel() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"
	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

	// 都是绝对路径
	s, err := filepath.Rel(`/a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err) // d/e <nil>

	// 都是相对路径
	s, err = filepath.Rel(`a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err) // d/e <nil>

	// 一个绝对一个相对
	s, err = filepath.Rel(`/a/b/c`, `a/b/c/d/e`)
	fmt.Println(s, err)
	//  Rel: can't make a/b/c/d/e relative to /a/b/c

	// 一个相对一个绝对
	s, err = filepath.Rel(`a/b/c`, `/a/b/c/d/e`)
	fmt.Println(s, err)
	//  Rel: can't make /a/b/c/d/e relative to a/b/c

	// 从 `a/b/c` 进入 `a/b/d/e`，只需要进入 `../d/e` 即可
	s, err = filepath.Rel(`a/b/c`, `a/b/d/e`)
	fmt.Println(s, err) // ../d/e <nil>
}

// 遍历指定目录（包括子目录），对遍历到的项目用walkFunc函数进行处理
func walk() {
	// 列出含有 *.txt 文件的目录（不是全部，因为会跳过一些子目录）
	err := filepath.Walk(`/usr`, findTxtDir)
	fmt.Println(err)

	fmt.Println("==============================")

	// 列出所有以 ab 开头的目录（全部，因为没有跳过任何项目）
	err = filepath.Walk(`/usr`, findabDir)
	fmt.Println(err)
}

// WalkFunc 函数：
// 列出含有 *.txt 文件的目录（不是全部，因为会跳过一些子目录）
func findTxtDir(path string, info os.FileInfo, err error) error {
	ok, err := filepath.Match(`*.txt`, info.Name())
	if ok {
		fmt.Println(filepath.Dir(path), info.Name())
		// 遇到 txt 文件则继续处理所在目录的下一个目录
		// 注意会跳过子目录
		return filepath.SkipDir
	}
	return err
}

// WalkFunc 函数：
// 列出所有以 ab 开头的目录（全部，因为没有跳过任何项目）
func findabDir(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		ok, err := filepath.Match(`[aA][bB]*`, info.Name())
		if err != nil {
			return err
		}
		if ok {
			fmt.Println(path)
		}
	}
	return nil
}
