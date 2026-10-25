package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

// saveDiary は既存の日記の後ろに entries を追記して保存する
func saveDiary(entries []Entry) error {
	existing, err := loadDiary()
	if err != nil {
		return err
	}
	existing = append(existing, entries...)

	jsonData, err := json.Marshal(existing)
	if err != nil {
		return fmt.Errorf("JSONの変換に失敗しました：%w", err)
	}

	err = os.WriteFile("diary.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("書き込みに失敗しました：%w", err)
	}

	return nil
}

// loadDiary は保存されている日記をすべて読み込む。ファイルがなければ空のスライスを返す
func loadDiary() ([]Entry, error) {
	jsonData, err := os.ReadFile("diary.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Entry{}, nil // 初回はファイルがない
		}
		return nil, fmt.Errorf("読み込みに失敗しました：%w", err)
	}

	var entries []Entry
	err = json.Unmarshal(jsonData, &entries)
	if err != nil {
		return nil, fmt.Errorf("JSONのパースに失敗しました：%w", err)
	}

	return entries, nil
}

func main() {
	// コマンドライン引数の数をチェック
	if len(os.Args) < 2 {
		fmt.Println("使い方：go run main.go [add|list] ...")
		return
	}

	// 1番目の引数（コマンド）を取得
	command := os.Args[1]

	// コマンドの内容に応じた処理
	switch command {
	case "add": // 日記エントリーの記録
		if len(os.Args) < 4 {
			fmt.Println("使い方：go run main.go add [タイトル] [本文]")
			return
		}

		entry := Entry{
			Timestamp: time.Now(),
			Title:     os.Args[2],
			Content:   os.Args[3],
		}
		if err := saveDiary([]Entry{entry}); err != nil {
			fmt.Println("日記の保存に失敗しました：", err)
			return
		}
		fmt.Println("日記を追加しました")

	case "list": // 日記の一覧表示
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("日記の読み込みに失敗しました：", err)
			return
		}
		if len(entries) == 0 {
			fmt.Println("日記はまだありません。")
			return
		}

		// 日記データの表示
		fmt.Println("--- 日記データ ---")
		for _, entry := range entries {
			fmt.Printf("日時：%s\n", entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("タイトル：%s\n", entry.Title)
			fmt.Printf("本文：\n%s\n", entry.Content)
		}

	default:
		fmt.Println("不明なコマンドです：", command)
	}
}
