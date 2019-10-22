package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// importは3種類
// 1. 標準パッケージ
// 2. サードパーティパッケージ
// 3. 自プロジェクトパッケージ
//
// importのルール
// - 種類毎にまとめて記述し、種類間は1行の空行を挟む
// - 種類内ではパスで昇順ソート
// - 1つしかない場合は丸括弧なし
// - パスはダブルクォートで挟む
//
// import文のパターン
// - パスのみの記述
// - エイリアスあり

func main() {
	// 対象となるファイルのパス
	pathPtr := flag.String("filepath", "", "file path")
	flag.Parse()

	path := *pathPtr
	if path == "" {
		panic("file path not found")
	}

	// ファイルをOpenする
	f, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 行数
	var line, importStart, importEnd uint
	// importの各行
	var importLines []string
	// import開始フラグ
	var importFlag bool
	for scanner.Scan() {
		line++
		// その行の内容
		lineStr := scanner.Text()
		// ここで一行ずつ処理
		//fmt.Printf("L%d: %s\n", line, lineStr)

		// TODO: importがない場合どうするか
		// TODO: importが1つしかない場合は処理をスキップしたい

		if !importFlag && lineStr == "import (" {
			// import部分の読み込み開始
			importFlag = true
			importStart = line
		}

		if importFlag {
			// 対象行の内容を格納
			importLines = append(importLines, lineStr)
		}

		if importFlag && lineStr == ")" {
			// import部分の読み込み終了
			importFlag = false
			importEnd = line
			break
		}
	}

	fmt.Printf("start: %d, end: %d\n", importStart, importEnd)
}
