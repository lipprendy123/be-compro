package storage

import (
	"compro/config"
	"io"

	storage_go "github.com/supabase-community/storage-go"
)

type SupabaseInterface interface {
	UploadFile(path string, file io.Reader) (string error)
}

type supabaseStruct struct {
	cfg *config.Config
}

func (s *supabaseStruct) UploadFile(path string, file io.Reader) (string error) {
	client := storage_go.New
}
