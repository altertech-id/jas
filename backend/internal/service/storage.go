package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"jayaabadisolusi/backend/internal/config"
	"jayaabadisolusi/backend/internal/domain"
	"jayaabadisolusi/backend/internal/repository"

	"github.com/google/uuid"
)

const MaxFileSize = 5 * 1024 * 1024 // 5 MB

type StorageService struct {
	cfg  *config.Config
	repo repository.Repository
}

func NewStorageService(cfg *config.Config, repo repository.Repository) *StorageService {
	return &StorageService{
		cfg:  cfg,
		repo: repo,
	}
}

// UploadImage validates and uploads an image (JPG, PNG, WebP <= 5MB)
func (s *StorageService) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, altText string, uploaderID *string) (*domain.Media, error) {
	if fileHeader.Size > MaxFileSize {
		return nil, errors.New("ukuran file melebihi batas maksimal 5 MB")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		return nil, errors.New("format file tidak didukung. Hanya JPG, PNG, dan WebP yang diperbolehkan")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("gagal membaca file: %w", err)
	}
	defer file.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, fmt.Errorf("gagal buffer file: %w", err)
	}

	mimeType := http.DetectContentType(buf.Bytes())
	if !strings.HasPrefix(mimeType, "image/") {
		return nil, errors.New("konten file bukan gambar valid")
	}

	uniqueFilename := fmt.Sprintf("%d-%s%s", time.Now().UnixNano(), uuid.New().String()[:8], ext)
	storagePath := fmt.Sprintf("uploads/%s", uniqueFilename)

	var publicURL string

	// If Supabase Storage is configured with real key
	if s.cfg.SupabaseServiceKey != "" && !strings.Contains(s.cfg.SupabaseServiceKey, "dummy") {
		uploadURL := fmt.Sprintf("%s/storage/v1/object/media/%s", s.cfg.SupabaseURL, storagePath)
		req, err := http.NewRequestWithContext(ctx, "POST", uploadURL, bytes.NewReader(buf.Bytes()))
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", "Bearer "+s.cfg.SupabaseServiceKey)
		req.Header.Set("Content-Type", mimeType)

		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("gagal upload ke Supabase Storage: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("gagal response storage (%d): %s", resp.StatusCode, string(body))
		}

		publicURL = fmt.Sprintf("%s/storage/v1/object/public/media/%s", s.cfg.SupabaseURL, storagePath)
	} else {
		// Mock / Dev storage URL
		publicURL = fmt.Sprintf("https://images.unsplash.com/photo-1581578731548-c64695cc6952?auto=format&fit=crop&w=800&q=80#dev-%s", uniqueFilename)
	}

	media := &domain.Media{
		StoragePath: storagePath,
		PublicURL:   publicURL,
		AltText:     altText,
		MimeType:    mimeType,
		SizeBytes:   fileHeader.Size,
		UploadedBy:  uploaderID,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.SaveMedia(ctx, media); err != nil {
		return nil, fmt.Errorf("gagal menyimpan metadata media: %w", err)
	}

	return media, nil
}
