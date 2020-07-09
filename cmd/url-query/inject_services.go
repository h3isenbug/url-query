package main

import (
	log2 "github.com/h3isenbug/url-common/pkg/services/log"
	url2 "github.com/h3isenbug/url-query/recipients/url"
	"github.com/h3isenbug/url-query/repositories/url"
	urlService "github.com/h3isenbug/url-query/services/url"
	"os"
)

func provideURLService(urlRepository url.ReadRepository, urlRecipient url2.ReadRecipient) urlService.Service {
	return urlService.NewServiceV1(urlRepository, urlRecipient)
}

func provideLogService() log2.LogService {
	return log2.NewLogServiceV1(os.Stdout)
}
