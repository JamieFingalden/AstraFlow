package worker

import (
	"AstraFlow-go/internal/client"
	"AstraFlow-go/internal/database"
	"AstraFlow-go/internal/repository"
	"AstraFlow-go/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// ProcessOCRTask 简单的OCR任务处理函数
// 直接处理OCR任务，避免过度封装
func ProcessOCRTask(fileID int64, filePath string) error {
	log.Printf("开始处理 OCR 任务: fileID=%d filePath=%s", fileID, filePath)

	// 1. 调用 Flask 服务进行 OCR 识别
	resp, err := client.SendFileToFlask(fileID, filePath)
	if err != nil {
		log.Printf("调用 Flask 服务失败: %v", err)
		return err
	}

	// 2. 解析 OCR 结果
	var ocrResult map[string]interface{}
	if err = json.Unmarshal(resp, &ocrResult); err != nil {
		log.Printf("解析 OCR 结果失败: %v", err)
		return err
	}

	if ocrResult["status"] != "success" {
		log.Printf("OCR 处理失败: %v", ocrResult["message"])
		return nil
	}

	fmt.Printf("OCR 结果: %v\n", ocrResult["data"])

	// 3. 从 OCR 结果中提取发票信息
	data, ok := ocrResult["data"].(map[string]interface{})
	if !ok {
		log.Printf("OCR 结果中缺少 data 字段或格式不正确")
		return fmt.Errorf("OCR 结果格式不正确")
	}

	amount, _ := data["amount"].(float64)
	invoiceDateStr, _ := data["invoice_date"].(string)
	vendor, _ := data["vendor"].(string)
	invoiceNumber, _ := data["invoice_number"].(string)
	category, _ := data["category"].(string)
	description, _ := data["description"].(string)

	// 解析发票日期
	var invoiceDate *time.Time
	if invoiceDateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", invoiceDateStr)
		if err != nil {
			log.Printf("解析发票日期失败: %v", err)
			// 如果日期解析失败，使用当前时间
			parsedDate = time.Now()
		}
		invoiceDate = &parsedDate
	} else {
		now := time.Now()
		invoiceDate = &now
	}

	// 4. 根据文件ID获取附件信息，获取用户和租户ID
	attachmentRepo := repository.NewAttachmentRepository()
	attachment, err := attachmentRepo.GetByID(fileID)
	if err != nil {
		log.Printf("获取附件信息失败: %v", err)
		return err
	}

	// 5. 处理可选字段
	var descriptionPtr *string
	if description != "" {
		descriptionPtr = &description
	}

	var categoryPtr *string
	if category != "" {
		categoryPtr = &category
	}

	// 获取租户ID
	var tenantId int64
	if attachment.TenantID != nil {
		tenantId = *attachment.TenantID
	}

	// 确定分类值
	var categoryValue string
	if categoryPtr != nil {
		categoryValue = *categoryPtr
	} else {
		categoryValue = ""
	}

	// 6. 创建发票
	invoiceService := service.NewInvoiceService()
	invoice, err := invoiceService.CreateInvoice(
		tenantId,
		attachment.UserID,
		*invoiceDate,
		amount,
		invoiceNumber,
		vendor,
		"", // taxId 空字符串
		categoryValue,
		"",        // paymentSource 空字符串
		"pending", // status 初始状态
		&filePath,
		descriptionPtr,
	)
	if err != nil {
		log.Printf("创建发票失败: %v", err)
		return err
	}

	// 7. 更新附件信息，关联到创建的发票
	attachment.InvoiceID = &invoice.ID
	if err := attachmentRepo.Update(attachment); err != nil {
		log.Printf("更新附件信息失败: %v", err)
		return err
	}

	// 8. 将OCR结果直接更新到发票记录中（因为结构相似）
	// 更新发票信息
	invoice.Amount = &amount
	invoice.Vendor = vendor
	invoice.InvoiceDate = invoiceDate
	invoice.InvoiceNumber = &invoiceNumber
	if categoryPtr != nil {
		invoice.Category = categoryPtr
	}

	// 设置识别状态为已识别
	invoice.IdentifyStatus = "identified"

	// 更新发票记录
	if err := database.DB.Save(invoice).Error; err != nil {
		log.Printf("更新发票信息失败: %v", err)
		// 这个错误不影响主要流程，记录但不返回错误
	}

	log.Printf("OCR 任务处理完成: fileID=%d, invoiceID=%d", fileID, invoice.ID)
	return nil
}

// ProcessOCRTaskFromQueue 从RabbitMQ队列处理OCR任务
func ProcessOCRTaskFromQueue(task client.OCRQueueTask) error {
	log.Printf("从队列处理 OCR 任务: taskID=%s fileID=%d filePath=%s", task.TaskID, task.FileID, task.FilePath)

	// 直接使用任务中包含的 fileID 进行处理
	err := ProcessOCRTask(task.FileID, task.FilePath)
	if err != nil {
		log.Printf("处理 OCR 任务失败: %v", err)
		return err
	}

	log.Printf("队列 OCR 任务处理完成: taskID=%s", task.TaskID)
	return nil
}

// StartOCRWorker 启动OCR工作进程，监听RabbitMQ队列
func StartOCRWorker() error {
	log.Println("启动 Go OCR 工作进程...")

	for {
		// 创建RabbitMQ客户端
		rabbitmqClient, err := client.NewRabbitMQOCRClient()
		if err != nil {
			log.Printf("初始化 RabbitMQ 客户端失败: %v，5秒后重试...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Println("开始消费 OCR 任务队列...")

		// 开始消费队列任务
		err = rabbitmqClient.ConsumeTasks(ProcessOCRTaskFromQueue)
		if err != nil {
			log.Printf("消费任务队列出错: %v，重新连接...", err)
		} else {
			log.Println("任务队列消费结束，重新连接...")
		}

		rabbitmqClient.Close()

		// 等待一段时间后重新连接
		time.Sleep(2 * time.Second)
	}
}
