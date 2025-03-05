***Города выводятся в консоль при выполнении команды cityReq.GetCity()***

    ulyanovsk
    ufa
    hm
    cheb
    chel
    cher
    cherkessk

***Категории нужно брать из адресной строки***

    moloko-i-slivki-128
	pekarnya-lenta-fresh-818
	uhod-i-gigiena-3509

в случае необходимости работать через прокси
нужно создать HTTP-клиент с настроенным Transport

```
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
```