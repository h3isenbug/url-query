package url

import (
	urlRecipient "github.com/h3isenbug/url-query/recipients/url"
	urlRepo "github.com/h3isenbug/url-query/repositories/url"
)

type Service interface {
	GetLongURL(shortPath, etag, userAgent string) (string, error)
}

type ServiceV1 struct {
	urlRepository urlRepo.ReadRepository
	urlRecipient  urlRecipient.ReadRecipient
}

func NewServiceV1(urlRepository urlRepo.ReadRepository, urlRecipient urlRecipient.ReadRecipient) Service {
	return &ServiceV1{
		urlRepository: urlRepository,
		urlRecipient:  urlRecipient,
	}
}

func (service ServiceV1) GetLongURL(shortPath, etag, userAgent string) (string, error) {
	url, err := service.urlRepository.GetLongURL(shortPath)
	if err != nil {
		return "", err
	}
	if err := service.urlRecipient.PublishURLViewed(shortPath, etag, userAgent); err != nil {
		return "", err
	}

	return url, nil
}
