package main

import (
	"lenta/internal/cityReq"
	"lenta/internal/itemsReq"
	"lenta/internal/tokenReq"
)

func main() {
	tokenReq := new(tokenReq.TokenRequest)
	cityReq := new(cityReq.CityRequest)
	itemsReq := new(itemsReq.ItemsRequest)

	token := tokenReq.GetToken()
	cityReq.GetCity()
	itemsReq.GetItems(token, "cheb", "uhod-i-gigiena-3509")
}

/*
	города выводятся в консоль командой cityReq.GetCity()
	ulyanovsk
	ufa
	hm
	cheb
	chel
	cher
	cherkessk

	категории из адресной строки
	moloko-i-slivki-128
	pekarnya-lenta-fresh-818
	uhod-i-gigiena-3509
*/
/*
в случае необходимости работать через прокси
Создается HTTP-клиент с настроенным Transport

proxyURL, err := url.Parse("http://your-proxy-server:port")
if err != nil {
fmt.Println("Ошибка при парсинге URL прокси:", err)
return ""
}
transport := &http.Transport{
Proxy: http.ProxyURL(proxyURL),
}
client := &http.Client{
Transport: transport,
}
*/
