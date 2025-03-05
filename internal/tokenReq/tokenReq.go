package tokenReq

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type TokenRequest struct{}

func (t *TokenRequest) GetToken() string {

	client := &http.Client{}
	var data = strings.NewReader(`request={"Head":{"MarketingPartnerKey":"mp300-b1de0bac2c257f3257bf5ef2eea4ecbc","Version":"angular_web_0.0.2","Client":"angular_web_0.0.2","Method":"sessionGet","SessionToken":"","RequestId":"sessionGet_b67edef8e9c93","DeviceId":"63c0d5ab-9e21-b608-14d9-c66038ceb18e","Domain":"moscow"},"Body":{}}`)
	req, err := http.NewRequest("POST", "https://lenta.com/api/rest/sessionGet", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Experiments", "")
	req.Header.Set("Origin", "https://lenta.com")
	req.Header.Set("Referer", "https://lenta.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Set("X-Organization-Id", "")
	req.Header.Set("X-Passport-Id", "")
	req.Header.Set("X-Retail-Brand", "lo")
	req.Header.Set("baggage", "sentry-environment=production,sentry-release=web-12.0.71,sentry-public_key=b99355c72549498d9e9075cc3d4006a2,sentry-trace_id=28cfe2bdb224445a9852557f1f067d76,sentry-sample_rate=1,sentry-sampled=true")
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sentry-trace", "28cfe2bdb224445a9852557f1f067d76-813053c704726652-1")
	req.Header.Set("traceparent", "00-752df802e358228fa201296a2556b1f4-22b05a28c408ff8e-01")
	req.Header.Set("x-span-id", "22b05a28c408ff8e")
	req.Header.Set("x-trace-id", "752df802e358228fa201296a2556b1f4")
	req.Header.Set("Cookie", "qrator_jsr=1741187193.230.YrmFjf6qv5Afunkt-oenlvvi7gpjcjt5u6q8mv61j2ikq361n-00; qrator_jsid=1741187193.230.YrmFjf6qv5Afunkt-ur220ph9pk3eug6m0eo6u0ascul7aro9; OmniPercent=74; OmniGroup=true; cookiesession1=678B2889E6A1BE63B5F954F82947D8DA; User_Agent=Mozilla%2F5.0%20(X11%3B%20Linux%20x86_64)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F133.0.0.0%20Safari%2F537.36; Is_Search_Bot=false; App_Cache_DeliveryMode=%7B%22type%22%3A%22pickup%22%7D; App_Cache_MPK=mp300-b1de0bac2c257f3257bf5ef2eea4ecbc; App_Cache_LegalEntityId=; App_Cache_PassportId=; App_Cache_CitySlug=moscow; Utk_DvcGuid=63c0d5ab-9e21-b608-14d9-c66038ceb18e")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ответ: %s\n", bodyText)

	var response Response

	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return ""
	}

	sessionTokenFromHead := response.Head.SessionToken
	fmt.Println("токен пришел SessionToken из Head:", sessionTokenFromHead)

	sessionTokenFromBody := response.Body.SessionToken
	fmt.Println("токен пришел SessionToken из Body:", sessionTokenFromBody)

	return sessionTokenFromBody
}

type Head struct {
	SessionToken string `json:"SessionToken"`
}

type Body struct {
	SessionToken string `json:"SessionToken"`
}

type Response struct {
	Head Head `json:"Head"`
	Body Body `json:"Body"`
}
