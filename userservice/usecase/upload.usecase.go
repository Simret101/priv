package usecase

import (
	"time"

	"user/domain"
	"user/repository"
)

type UploadProfileUsecase struct {
	Repo repository.UploadRepo
}

func NewUploadUsecase(repo repository.UploadRepo) *UploadProfileUsecase {
	return &UploadProfileUsecase{
		Repo: repo,
	}
}

// UploadPicture handles the uploading of the profile picture and notifies RabbitMQ for asynchronous processing
func (uploaduc *UploadProfileUsecase) UploadPicture(path string, id string) error {
	profile_picture := domain.Media{
		Path:          path,
		Uplaoded_date: time.Now(),
	}

	// Call the repository to update MongoDB and send the message to RabbitMQ
	err := uploaduc.Repo.AddProfile(profile_picture, id)
	if err != nil {
		return err
	}

	// Return nil to indicate success
	return nil
}
