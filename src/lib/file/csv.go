package file

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVByFields(path string, fields ...string) []map[string]string {
	// 開啟 CSV 檔案
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 讀取 CSV 檔案
	reader := csv.NewReader(file)

	// 讀取 CSV 檔案第一行，即欄位名稱
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("讀取 CSV 檔案欄位錯誤：", err)
		return nil
	}

	// 指定要取得的欄位索引
	columnIndexMap := make(map[int]string, len(fields))
	for _, name := range fields {
		for j, header := range headers {
			if header == name {
				columnIndexMap[j] = name
				break
			}
		}
	}

	records := make([]map[string]string, 0)
	// 讀取 CSV 檔案內容，並取得指定欄位的內容
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		needAppend := false
		// 取得指定欄位的內容
		result := make(map[string]string, len(columnIndexMap))
		for idx, filed := range columnIndexMap {
			result[filed] = record[idx]
			needAppend = true
		}
		if needAppend {
			records = append(records, result)
		}
	}
	return records
}
