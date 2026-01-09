package service

import (
	"AstraFlow-go/internal/model"
	"AstraFlow-go/internal/repository"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AttachmentService interface {
	UploadFile(file *multipart.FileHeader, userID int64, tenantID *int64) (*model.Attachment, error)
	GetAttachmentByID(id int64) (*model.Attachment, error)
	GetAttachmentsByUserID(userID int64) ([]*model.Attachment, error)
	GetAttachmentsByTenantID(tenantID int64) ([]*model.Attachment, error)
	GetAttachmentsByInvoiceID(invoiceID int64) ([]*model.Attachment, error)
	DeleteAttachment(id int64) error
	ValidateFile(file *multipart.FileHeader) error
	UpdateAttachmentInvoiceID(fileID, invoiceID int64) error
}

type attachmentService struct {
	repo repository.AttachmentRepository
}

func NewAttachmentService() AttachmentService {
	return &attachmentService{
		repo: repository.NewAttachmentRepository(),
	}
}

func (s *attachmentService) UploadFile(file *multipart.FileHeader, userID int64, tenantID *int64) (*model.Attachment, error) {
	// 验证文件
	if err := s.ValidateFile(file); err != nil {
		return nil, err
	}

	// 准备存储路径
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create upload directory: %v", err)
	}

	// 生成唯一文件名
	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().UnixNano(), extension)
	filePath := filepath.Join(uploadDir, filename)

	// 保存文件
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// 将源文件内容复制到目标文件
	if _, err := dst.ReadFrom(src); err != nil {
		return nil, err
	}

	// 创建附件记录
	fileType := extension
	attachment := &model.Attachment{
		TenantID: tenantID,
		UserID:   userID,
		FileURL:  "/uploads/" + filename,
		FileType: &fileType,
		FileSize: &file.Size,
	}

	if err := s.repo.Create(attachment); err != nil {
		// 如果数据库保存失败，删除已上传的文件
		os.Remove(filePath)
		return nil, err
	}

	return attachment, nil
}

func (s *attachmentService) GetAttachmentByID(id int64) (*model.Attachment, error) {
	return s.repo.GetByID(id)
}

func (s *attachmentService) GetAttachmentsByUserID(userID int64) ([]*model.Attachment, error) {
	return s.repo.GetByUserID(userID)
}

func (s *attachmentService) GetAttachmentsByTenantID(tenantID int64) ([]*model.Attachment, error) {
	return s.repo.GetByTenantID(tenantID)
}

func (s *attachmentService) GetAttachmentsByInvoiceID(invoiceID int64) ([]*model.Attachment, error) {
	return s.repo.GetByInvoiceID(invoiceID)
}

func (s *attachmentService) DeleteAttachment(id int64) error {
	attachment, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// 删除文件系统中的文件
	if attachment.FileURL != "" {
		filePath := "." + attachment.FileURL
		if err := os.Remove(filePath); err != nil {
			// 即使文件删除失败，也要删除数据库记录
			fmt.Printf("Warning: failed to delete file %s: %v\n", filePath, err)
		}
	}

	return s.repo.Delete(id)
}

func (s *attachmentService) ValidateFile(file *multipart.FileHeader) error {
	// 检查文件类型
	allowedTypes := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".bmp":  true,
		".pdf":  true,
	}

	extension := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedTypes[extension] {
		return errors.New("unsupported file type")
	}

	// 检查文件大小 (最大 10MB)
	const maxFileSize int64 = 10 * 1024 * 1024
	if file.Size > maxFileSize {
		return errors.New("file size exceeds maximum limit of 10MB")
	}

	return nil
}

func (s *attachmentService) UpdateAttachmentInvoiceID(fileID, invoiceID int64) error {
	return s.repo.UpdateInvoiceID(fileID, invoiceID)
}
