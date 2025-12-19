package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tiktoken-go/tokenizer"
	tree_sitter_markdown "github.com/tree-sitter-grammars/tree-sitter-markdown/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

var codec tokenizer.Codec

func main() {
	var err error
	codec, err = tokenizer.Get(tokenizer.Cl100kBase)
	if err != nil {
		fmt.Printf("failed to get tokenizer: %v\n", err)
		return
	}
	// 创建解析器
	parser := sitter.NewParser()
	defer parser.Close()

	// 加载 markdown 语法
	markdownLanguage := sitter.NewLanguage(tree_sitter_markdown.Language())
	parser.SetLanguage(markdownLanguage)

	// 读取 md 目录下的所有 markdown 文件
	mdFiles, err := filepath.Glob("md/*.md*")
	if err != nil {
		fmt.Printf("读取 md 目录失败: %v\n", err)
		return
	}

	if len(mdFiles) == 0 {
		fmt.Println("md 目录下没有找到 markdown 文件")
		return
	}

	fmt.Printf("找到 %d 个 markdown 文件:\n", len(mdFiles))
	for _, file := range mdFiles {
		fmt.Printf("  - %s\n", file)
	}
	fmt.Println()

	// 处理每个 markdown 文件
	for _, mdFile := range mdFiles {
		fmt.Printf("正在处理文件: %s\n", mdFile)
		processMarkdownFile(parser, mdFile)
		fmt.Println("----------------------------------------")
	}
}

func processMarkdownFile(parser *sitter.Parser, mdFile string) {
	// 读取 markdown 文件内容
	markdownContent, err := os.ReadFile(mdFile)
	if err != nil {
		fmt.Printf("读取文件 %s 失败: %v\n", mdFile, err)
		return
	}

	// 创建输出文件名
	baseName := filepath.Base(mdFile)
	// 将文件扩展名从 .md 改为 .txt
	outputFileName := strings.TrimSuffix(baseName, filepath.Ext(baseName)) + ".txt"
	outputFile := filepath.Join("md-parse", outputFileName)

	// 记录总开始时间
	totalStartTime := time.Now()
	var output strings.Builder

	// 解析 markdown 内容
	startTime0 := time.Now()
	tree := parser.Parse(markdownContent, nil)
	duration0 := time.Since(startTime0)
	defer tree.Close()

	rootNode := tree.RootNode()

	// 准备输出内容
	output.WriteString(fmt.Sprintf("=== Markdown 解析结果 (%s) ===\n", baseName))
	output.WriteString(fmt.Sprintf("根节点类型: %s\n", rootNode.Kind()))
	output.WriteString(fmt.Sprintf("根节点起始位置: %d\n", rootNode.StartByte()))
	output.WriteString(fmt.Sprintf("根节点结束位置: %d\n", rootNode.EndByte()))
	output.WriteString(fmt.Sprintf("子节点数量: %d\n", rootNode.ChildCount()))
	output.WriteString("\n")

	// 打印解析耗时
	output.WriteString("=== Markdown 解析 ===\n")
	output.WriteString(fmt.Sprintf("Markdown 解析耗时: %v\n", duration0))

	// 递归打印 AST 结构
	// startTime1 := time.Now()
	// output.WriteString("\n=== AST 结构 ===\n")
	// printNodeToString(&output, rootNode, string(markdownContent), 0)
	// duration1 := time.Since(startTime1)
	// output.WriteString(fmt.Sprintf("AST结构打印耗时: %v\n", duration1))

	// 打印节点详细信息
	// startTime2 := time.Now()
	// output.WriteString("\n=== 节点详细信息 ===\n")
	// printNodeDetailsToString(&output, rootNode, string(markdownContent), 0)
	// duration2 := time.Since(startTime2)
	// output.WriteString(fmt.Sprintf("节点详细信息打印耗时: %v\n", duration2))

	// 统计各种节点类型
	// startTime3 := time.Now()
	// output.WriteString("\n=== 节点类型统计 ===\n")
	// nodeStats := make(map[string]int)
	// countNodes(rootNode, nodeStats)
	// for nodeType, count := range nodeStats {
	// 	output.WriteString(fmt.Sprintf("%s: %d\n", nodeType, count))
	// }
	// duration3 := time.Since(startTime3)
	// output.WriteString(fmt.Sprintf("节点类型统计耗时: %v\n", duration3))

	// 按标题分段打印
	startTime4 := time.Now()
	output.WriteString("\n=== 按标题分段打印 ===\n")
	printMarkdownBySections(&output, rootNode, string(markdownContent))
	duration4 := time.Since(startTime4)
	output.WriteString(fmt.Sprintf("按标题分段打印耗时: %v\n", duration4))

	// 计算总耗时
	totalDuration := time.Since(totalStartTime)
	output.WriteString(fmt.Sprintf("\n=== 总执行时间 ===\n"))
	output.WriteString(fmt.Sprintf("总耗时: %v\n", totalDuration))
	output.WriteString(fmt.Sprintf("各方法耗时明细:\n"))
	output.WriteString(fmt.Sprintf("  - Markdown 解析: %v\n", duration0))
	// output.WriteString(fmt.Sprintf("  - AST结构打印: %v\n", duration1))
	// output.WriteString(fmt.Sprintf("  - 节点详细信息打印: %v\n", duration2))
	// output.WriteString(fmt.Sprintf("  - 节点类型统计: %v\n", duration3))
	output.WriteString(fmt.Sprintf("  - 按标题分段打印: %v\n", duration4))

	// 将结果写入文件
	err = os.WriteFile(outputFile, []byte(output.String()), 0644)
	if err != nil {
		fmt.Printf("写入输出文件 %s 失败: %v\n", outputFile, err)
		return
	}

	fmt.Printf("解析结果已保存到: %s\n", outputFile)

	// 同时在控制台输出简要信息
	fmt.Printf("根节点类型: %s, 子节点数量: %d\n", rootNode.Kind(), rootNode.ChildCount())
	fmt.Printf("总耗时: %v\n", totalDuration)
	fmt.Printf("各方法耗时明细:\n")
	fmt.Printf("  - Markdown 解析: %v\n", duration0)
	fmt.Printf("  - 按标题分段打印: %v\n", duration4)
}

// 递归打印节点信息
func printNode(node *sitter.Node, source string, depth int) {
	indent := strings.Repeat("  ", depth)

	startPoint := node.StartPosition()
	endPoint := node.EndPosition()

	fmt.Printf("%s%s [%d:%d - %d:%d] (%d-%d)\n",
		indent,
		node.Kind(),
		startPoint.Row+1, startPoint.Column,
		endPoint.Row+1, endPoint.Column,
		node.StartByte(), node.EndByte())

	// 如果是文本节点，打印内容
	if node.Kind() == "text" || node.Kind() == "code_span_text" || node.Kind() == "link_text" {
		content := node.Utf8Text([]byte(source))
		if len(content) > 50 {
			content = content[:50] + "..."
		}
		fmt.Printf("%s  内容: %q\n", indent, content)
	}

	// 递归处理子节点
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		printNode(child, source, depth+1)
	}
}

// 统计节点类型
func countNodes(node *sitter.Node, stats map[string]int) {
	nodeType := node.Kind()
	stats[nodeType]++

	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		countNodes(child, stats)
	}
}

// 打印节点详细信息
func printNodeDetails(node *sitter.Node, source string, depth int) {
	// 只打印特定类型的节点块
	nodeType := node.Kind()
	importantNodeTypes := map[string]bool{
		"atx_heading": true, // 标题
		"paragraph":   true, // 段落
		"list":        true, // 列表
		// "list_item":         true, // 列表项
		// "strong":            true, // 粗体
		// "emphasis":          true, // 斜体
		// "link":              true, // 链接
		"fenced_code_block": true, // 代码块
		"block_quote":       true, // 引用块
		"pipe_table":        true, // 表格
	}

	if !importantNodeTypes[nodeType] {
		// 如果不是重要节点类型，递归处理子节点但不打印当前节点
		for i := uint(0); i < node.ChildCount(); i++ {
			child := node.Child(i)
			printNodeDetails(child, source, depth)
		}
		return
	}

	indent := strings.Repeat("  ", depth)

	startPoint := node.StartPosition()
	endPoint := node.EndPosition()

	fmt.Printf("%s=== %s ===\n", indent, getNodeTypeDisplayName(nodeType))
	fmt.Printf("%s位置: [%d:%d - %d:%d]\n", indent,
		startPoint.Row+1, startPoint.Column,
		endPoint.Row+1, endPoint.Column)
	fmt.Printf("%s字节范围: %d-%d\n", indent, node.StartByte(), node.EndByte())

	// 打印节点完整内容，不进行截断
	content := node.Utf8Text([]byte(source))
	if content != "" {
		fmt.Printf("%s内容:\n%s```\n%s\n```\n", indent, indent, content)
	}

	fmt.Printf("%s\n", indent)

	// 对于重要节点类型，不再递归处理子节点，统一在该节点完整打印
	// for i := uint(0); i < node.ChildCount(); i++ {
	// 	child := node.Child(i)
	// 	printNodeDetails(child, source, depth+1)
	// }
}

// 获取节点类型的显示名称
func getNodeTypeDisplayName(nodeType string) string {
	displayNames := map[string]string{
		"atx_heading": "标题",
		"paragraph":   "段落",
		"list":        "列表",
		// "list_item":         "列表项",
		// "strong":            "粗体文本",
		// "emphasis":          "斜体文本",
		// "link":              "链接",
		"fenced_code_block": "代码块",
		"block_quote":       "引用块",
		"pipe_table":        "表格",
	}
	if name, exists := displayNames[nodeType]; exists {
		return name
	}
	return nodeType
}

// 递归打印节点信息到字符串
func printNodeToString(output *strings.Builder, node *sitter.Node, source string, depth int) {
	indent := strings.Repeat("  ", depth)

	startPoint := node.StartPosition()
	endPoint := node.EndPosition()

	output.WriteString(fmt.Sprintf("%s%s [%d:%d - %d:%d] (%d-%d)\n",
		indent,
		node.Kind(),
		startPoint.Row+1, startPoint.Column,
		endPoint.Row+1, endPoint.Column,
		node.StartByte(), node.EndByte()))

	// 如果是文本节点，打印内容
	// if node.Kind() == "text" || node.Kind() == "code_span_text" || node.Kind() == "link_text" {
	// 	content := node.Utf8Text([]byte(source))
	// 	if len(content) > 50 {
	// 		content = content[:50] + "..."
	// 	}
	// 	output.WriteString(fmt.Sprintf("%s  内容: %q\n", indent, content))
	// }

	// 递归处理子节点
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		printNodeToString(output, child, source, depth+1)
	}
}

// 打印节点详细信息到字符串
func printNodeDetailsToString(output *strings.Builder, node *sitter.Node, source string, depth int) {
	printNodeDetailsToStringWithHeaders(output, node, source, depth, []string{})
}

// 带标题路径的节点详细信息打印
func printNodeDetailsToStringWithHeaders(output *strings.Builder, node *sitter.Node, source string, depth int, headerPath []string) {
	// 更新标题路径
	currentHeaderPath := make([]string, len(headerPath))
	copy(currentHeaderPath, headerPath)

	// 如果是标题节点，更新标题路径
	if node.Kind() == "atx_heading" {
		titleText := extractTitleText(node, source)
		if titleText != "" {
			currentHeaderPath = append(currentHeaderPath, titleText)
		}
	}

	// 确定传递给子节点的标题路径
	pathForChildren := currentHeaderPath

	// 获取节点类型
	nodeType := node.Kind()

	// 处理标题节点
	if nodeType == "atx_heading" {
		indent := strings.Repeat("  ", depth)

		startPoint := node.StartPosition()
		endPoint := node.EndPosition()

		// 打印标题路径（包含当前标题）
		if len(currentHeaderPath) > 0 {
			output.WriteString(fmt.Sprintf("%s\n", strings.Join(currentHeaderPath, "\n")))
		}

		output.WriteString(fmt.Sprintf("%s=== %s ===\n", indent, getNodeTypeDisplayName(nodeType)))
		output.WriteString(fmt.Sprintf("%s位置: [%d:%d - %d:%d]\n", indent,
			startPoint.Row+1, startPoint.Column,
			endPoint.Row+1, endPoint.Column))
		output.WriteString(fmt.Sprintf("%s字节范围: %d-%d\n", indent, node.StartByte(), node.EndByte()))

		// 打印节点完整内容，不进行截断
		content := source[node.StartByte():node.EndByte()]
		if content != "" {
			output.WriteString(fmt.Sprintf("%s内容:\n%s```\n%s\n```\n", indent, indent, content))
		}

		output.WriteString(fmt.Sprintf("%s\n", indent))
	}

	// 不处理其他节点类型，只处理标题节点

	// 递归处理子节点
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		printNodeDetailsToStringWithHeaders(output, child, source, depth+1, pathForChildren)
	}
}

// 提取标题文本
func extractTitleText(node *sitter.Node, source string) string {
	if node.Kind() != "atx_heading" {
		return ""
	}

	// 直接获取 atx_heading 节点的文本内容，保留 # 符号
	return strings.TrimSpace(source[node.StartByte():node.EndByte()])
}

// 按标题分段打印 Markdown 文档
func printMarkdownBySections(output *strings.Builder, rootNode *sitter.Node, source string) {
	// 收集所有标题节点
	var allHeaders []*sitter.Node
	collectAllHeaders(rootNode, &allHeaders)
	// 按位置排序标题节点
	allHeaders = sortHeadersByPosition(allHeaders)

	var sections []*SectionInfo
	if len(allHeaders) == 0 {
		byteCount := len(source)
		if byteCount > 4000 {
			subChunks := splitTextWithSlidingWindow(source, 0)
			sections = subChunks
		} else {
			chunk := &SectionInfo{
				Content:    source,
				StartPos:   0,
				EndPos:     len(source),
				StartLine:  1,
				EndLine:    1 + countLines(source),
				StartCol:   1,
				EndCol:     calculateColumn(source, len(source)-1),
				TokenCount: byteCount / 4,
			}
			sections = append(sections, chunk)
		}
	} else {
		// 为每个标题节点创建段落
		for i, header := range allHeaders {
			// 获取当前标题的完整路径
			headerPath := getHeaderPath(header, source, allHeaders)

			// 查找下一个标题节点
			var nextHeader *sitter.Node
			if i < len(allHeaders)-1 {
				nextHeader = allHeaders[i+1]
			}

			// 提取内容
			content := extractContentBetweenHeaders(header, nextHeader, source)

			// 创建段落信息
			section := &SectionInfo{
				Headers:   headerPath,
				Content:   content,
				StartLine: int(header.EndPosition().Row + 1),
				StartCol:  int(header.EndPosition().Column),
				StartPos:  int(header.EndByte()),
			}
			if nextHeader != nil {
				section.EndPos = int(nextHeader.StartByte())
				section.EndLine = int(nextHeader.StartPosition().Row + 1)
				section.EndCol = int(nextHeader.StartPosition().Column)
			} else {
				section.EndPos = len(source)
				section.EndLine = int(rootNode.EndPosition().Row + 1)
				section.EndCol = int(rootNode.EndPosition().Column)
			}

			sections = append(sections, section)
		}
	}

	// 按标题分段打印
	for i, section := range sections {
		// 打印段号
		output.WriteString(fmt.Sprintf("\n第%d段: \n", i+1))
		output.WriteString(fmt.Sprintf("位置: [%d:%d - %d:%d]\n",
			section.StartLine, section.StartCol,
			section.EndLine, section.EndCol))
		output.WriteString(fmt.Sprintf("字节范围: %d-%d\n", section.StartPos, section.EndPos))
		output.WriteString("内容: \n```\n")

		// 打印完整的标题路径
		for _, header := range section.Headers {
			output.WriteString(fmt.Sprintf("%s\n", header))
		}

		// 打印内容
		if section.Content != "" {
			output.WriteString(fmt.Sprintf("%s\n", section.Content))
		}

		// 段落分隔符
		output.WriteString("```\n")
	}
}

// 段落信息结构
type SectionInfo struct {
	Headers    []string // 标题路径
	Content    string   // 该段的内容
	StartLine  int      // 开始行
	StartCol   int      // 开始列
	EndLine    int      // 结束行
	EndCol     int      // 结束列
	StartPos   int      // 开始位置
	EndPos     int      // 结束位置
	TokenCount int
}

// 收集所有标题节点
func collectAllHeaders(node *sitter.Node, headers *[]*sitter.Node) {
	// 如果是标题节点，添加到列表中
	if node.Kind() == "atx_heading" {
		*headers = append(*headers, node)
	}

	// 递归处理子节点
	for i := uint(0); i < node.ChildCount(); i++ {
		child := node.Child(i)
		collectAllHeaders(child, headers)
	}
}

// 按位置排序标题节点
func sortHeadersByPosition(headers []*sitter.Node) []*sitter.Node {
	for i := 0; i < len(headers)-1; i++ {
		for j := i + 1; j < len(headers); j++ {
			if headers[i].StartByte() > headers[j].StartByte() {
				headers[i], headers[j] = headers[j], headers[i]
			}
		}
	}
	return headers
}

// 获取标题的完整路径
func getHeaderPath(headerNode *sitter.Node, source string, allHeaders []*sitter.Node) []string {
	var path []string

	// 获取当前标题在文档中的位置
	currentStartByte := headerNode.StartByte()

	// 找到当前标题在排序后列表中的位置
	var currentIndex int
	for i, header := range allHeaders {
		if header.StartByte() == currentStartByte {
			currentIndex = i
			break
		}
	}

	// 构建标题路径
	var currentLevel int
	currentTitle := extractTitleText(headerNode, source)
	if currentTitle != "" {
		// 获取当前标题的级别
		currentLevel = getHeaderLevel(headerNode)
		// 当前标题路径记录在path中，在后续content中不包含
		path = append(path, currentTitle)
	}

	// 向前查找父级标题
	for i := currentIndex - 1; i >= 0; i-- {
		prevHeader := allHeaders[i]
		prevLevel := getHeaderLevel(prevHeader)

		// 如果前一个标题的级别小于当前标题的级别，则是父级标题
		if prevLevel < currentLevel {
			prevTitle := extractTitleText(prevHeader, source)
			if prevTitle != "" {
				path = append([]string{prevTitle}, path...)
				currentLevel = prevLevel
			}
		}
	}

	return path
}

// 获取标题级别
func getHeaderLevel(headerNode *sitter.Node) int {
	if headerNode.Kind() != "atx_heading" {
		return 0
	}

	// 查找 atx_h?_marker 子节点来确定级别
	for i := uint(0); i < headerNode.ChildCount(); i++ {
		child := headerNode.Child(i)
		if strings.HasPrefix(child.Kind(), "atx_h") && strings.HasSuffix(child.Kind(), "_marker") {
			// 从 atx_h1_marker 这样的字符串中提取数字
			levelStr := strings.TrimPrefix(child.Kind(), "atx_h")
			levelStr = strings.TrimSuffix(levelStr, "_marker")
			if level, err := strconv.Atoi(levelStr); err == nil {
				return level
			}
		}
	}

	return 0
}

// 提取两个标题之间的内容
func extractContentBetweenHeaders(currentHeader, nextHeader *sitter.Node, source string) string {
	startPos := int(currentHeader.EndByte())

	var endPos int
	if nextHeader != nil {
		endPos = int(nextHeader.StartByte())
	} else {
		// 如果没有下一个标题，返回文档末尾
		endPos = len(source)
	}

	if startPos >= endPos {
		return ""
	}

	// 提取内容
	content := source[startPos:endPos]

	// 清理内容：去除首尾空白，但保留段落格式
	// content = strings.TrimSpace(content)

	return content
}

func splitFuncWithSlidingWindow(content string, funcStartLine int) []*SectionInfo {
	maxTokens := 1000
	overlapTokens := 100

	if maxTokens <= 0 || overlapTokens < 0 || overlapTokens >= maxTokens {
		return nil
	}

	// 编码内容获取tokens和字节偏移量
	_, tokens, err := codec.Encode(content)
	if err != nil {
		return nil
	}

	totalTokens := len(tokens)
	if totalTokens == 0 {
		return nil
	}

	// 计算每个token的字节偏移量
	byteOffsets := make([]int, len(tokens)+1)
	currentOffset := 0
	for i, token := range tokens {
		byteOffsets[i] = currentOffset
		currentOffset += len(token)
	}
	byteOffsets[len(tokens)] = currentOffset

	// 预分配切片
	estimatedChunks := (totalTokens + maxTokens - 1) / maxTokens
	chunks := make([]*SectionInfo, 0, estimatedChunks)

	startTokenIdx := 0

	for startTokenIdx < totalTokens {
		// 计算当前块的结束位置
		endTokenIdx := startTokenIdx + maxTokens
		if endTokenIdx > totalTokens {
			endTokenIdx = totalTokens
		}

		// 提取代码块
		startByte := byteOffsets[startTokenIdx]
		endByte := byteOffsets[endTokenIdx] - 1
		if endByte >= len(content) {
			endByte = len(content) - 1
		}

		chunkContent := content[startByte : endByte+1]

		// 计算起始行和列
		startLine := funcStartLine + countLines(content[:startByte])
		startColumn := calculateColumn(content, startByte)

		// 计算结束行和列
		endLine := startLine + countLines(chunkContent) - 1
		endColumn := calculateColumn(chunkContent, endByte-startByte)

		chunks = append(chunks, &SectionInfo{
			Content:    chunkContent,
			StartLine:  startLine,
			StartCol:   startColumn,
			EndLine:    endLine,
			EndCol:     endColumn,
			TokenCount: endTokenIdx - startTokenIdx,
		})

		if endTokenIdx >= totalTokens {
			break
		}

		// 计算下一个块的起始位置（应用滑动窗口）
		if remaining := totalTokens - endTokenIdx; remaining < maxTokens {
			// 最后一块，调整重叠量
			startTokenIdx = endTokenIdx - (maxTokens - remaining)
		} else {
			// 正常情况，使用固定重叠
			startTokenIdx = endTokenIdx - overlapTokens
		}

		// 防止索引越界
		if startTokenIdx < 0 {
			startTokenIdx = 0
		}
	}

	return chunks
}

func splitTextWithSlidingWindow(content string, funcStartLine int) []*SectionInfo {
	maxBytes := 4000    // 将token转换为字节数
	overlapBytes := 400 // 将token转换为字节数

	if maxBytes <= 0 || overlapBytes < 0 || overlapBytes >= maxBytes {
		return nil
	}

	totalBytes := len(content)
	if totalBytes == 0 {
		return nil
	}

	// 预分配切片
	estimatedChunks := (totalBytes + maxBytes - 1) / maxBytes
	chunks := make([]*SectionInfo, 0, estimatedChunks)

	startByteIdx := 0

	for startByteIdx < totalBytes {
		// 计算当前块的结束位置
		endByteIdx := startByteIdx + maxBytes
		if endByteIdx > totalBytes {
			endByteIdx = totalBytes
		}

		// 提取代码块
		chunkContent := content[startByteIdx:endByteIdx]

		// 计算起始行和列
		startLine := funcStartLine + countLines(content[:startByteIdx])
		startColumn := calculateColumn(content, startByteIdx)

		// 计算结束行和列
		endLine := startLine + countLines(chunkContent) - 1
		endColumn := calculateColumn(chunkContent, endByteIdx-startByteIdx)

		// 计算token数量（用于兼容性，但不再用于分割逻辑）
		tokenCount := (endByteIdx - startByteIdx) / 4 // 粗略估计：1token≈4字节

		chunks = append(chunks, &SectionInfo{
			Content:    chunkContent,
			StartLine:  startLine,
			StartCol:   startColumn,
			EndLine:    endLine,
			EndCol:     endColumn,
			TokenCount: tokenCount,
		})

		if endByteIdx >= totalBytes {
			break
		}

		// 计算下一个块的起始位置（应用滑动窗口）
		if remaining := totalBytes - endByteIdx; remaining < maxBytes {
			// 最后一块，调整重叠量
			startByteIdx = endByteIdx - (maxBytes - remaining)
		} else {
			// 正常情况，使用固定重叠
			startByteIdx = endByteIdx - overlapBytes
		}

		// 防止索引越界
		if startByteIdx < 0 {
			startByteIdx = 0
		}
	}

	return chunks
}

// calculateColumn 根据字节偏移量计算在当前行的列位置
func calculateColumn(content string, byteOffset int) int {
	if byteOffset >= len(content) {
		byteOffset = len(content) - 1
	}
	if byteOffset < 0 {
		return 0
	}

	// 从字节偏移量向前查找最后一个换行符
	column := 0
	for i := byteOffset; i >= 0; i-- {
		if content[i] == '\n' {
			break
		}
		column++
	}
	return column
}

// countLines 计算字符串中的行数
func countLines(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := 0
	for _, c := range s {
		if c == '\n' {
			count++
		}
	}

	// 最后一行可能没有换行符
	if len(s) > 0 && s[len(s)-1] != '\n' {
		count++
	}

	return count
}
