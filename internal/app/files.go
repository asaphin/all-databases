package app

import (
	"context"
	"fmt"
	"github.com/asaphin/all-databases-go/internal/domain"
	"github.com/asaphin/all-databases-go/internal/utils"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var filesList = []string{
	"./testdata/sample.yml",
}

type FilesScenarioService struct {
	filesRepository FilesRepository
	ledger          map[string]struct{}
}

func NewFilesScenarioService(filesRepository FilesRepository) *FilesScenarioService {
	return &FilesScenarioService{
		filesRepository: filesRepository,
		ledger:          make(map[string]struct{}),
	}
}

func (s *FilesScenarioService) Run() {
	log.Debug("files scenario started")

	for _, file := range filesList {
		f, err := s.readFile(file)
		utils.LogAsErrorIfError(err, fmt.Sprintf("unable to load file %s", file))

		err = s.filesRepository.Put(context.Background(), f)
		utils.LogAsErrorIfError(err, "unable to save file")
	}
}

func (s *FilesScenarioService) cleanupFiles() {
	ctx := context.Background()

	for id := range s.ledger {
		err := s.filesRepository.Delete(ctx, id)
		if err != nil {
			log.WithError(err).WithField("fileID", id).Warning("unable to delete file")
		}
	}
}

func (s *FilesScenarioService) readFile(path string) (*domain.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer utils.LogAsWarningIfError(f.Close(), "unable to close file")

	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	file := &domain.File{
		FileListItem: domain.FileListItem{
			ID:        uuid.String(),
			Name:      f.Name(),
			Type:      "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Data: nil,
	}

	return file, nil

}
