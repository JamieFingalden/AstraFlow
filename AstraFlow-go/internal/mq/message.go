package mq

// OCRTask 表示一次发票 OCR 处理任务
type OCRTask struct {
	FileID   int64  `json:"file_id"`
	FilePath string `json:"file_path"`
}
