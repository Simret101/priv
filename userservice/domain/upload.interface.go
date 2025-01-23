package domain

type UploadController interface {
	UploadImg()
}

type UploadUsecase interface {
	UploadPicture(path, id string) error
}

type UploadRepository interface {
	sendToRabbitMQ(message interface{}) error
	AddProfile(media Media) error
}
