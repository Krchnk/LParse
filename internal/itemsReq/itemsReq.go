package itemsReq

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type ItemsRequest struct{}

func (t *ItemsRequest) GetItems(sessionToken string, city string, url string) {
	client := &http.Client{}

	lastDashIndex := strings.LastIndex(url, "-")
	if lastDashIndex == -1 {
		fmt.Println("Дефис не найден")
		return
	}
	categoryId := url[lastDashIndex+1:]

	jsonData := fmt.Sprintf(
		`{"categoryId":%s,"limit":40,"offset":0,"sort":{"type":"popular","order":"desc"},"filters":{"range":[],"checkbox":[],"multicheckbox":[]}}`,
		categoryId,
	)
	var data = strings.NewReader(jsonData)

	req, err := http.NewRequest("POST", "https://lenta.com/api-gateway/v1/catalog/items", data)
	if err != nil {
		log.Fatal(err)
	}

	urlFull := fmt.Sprintf("https://lenta.com/catalog/%s", url)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6,zh;q=0.5,hy;q=0.4,fr;q=0.3")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("DeviceID", "1b642436-ab55-510a-a1c1-942a3cb468f8")
	req.Header.Set("Experiments", "exp_recommendation_cms.true, exp_apigw_purchase.test, exp_lentapay.test, exp_omni_price.test, exp_profile_bell.test, exp_newui_cancel_order.test, exp_newui_history_active_action.test_stars, exp_comment_picker_and_courier.test, exp_general_editing_page.test, exp_cl_omni_support.test, exp_cl_omni_authorization.test, exp_onboarding_sbp.default, exp_fullscreen.test, exp_profile_login.false, exp_new_notifications_show_unauthorized.test, exp_assembly_cost_location.cart, exp_search_bottom.default, exp_onboarding_editing_order.test, exp_cart_new_carousel.default, exp_newui_cart_cancel_editing.test, exp_newui_cart_button.test, exp_new_promov3., exp_sbp_enabled.test, exp_new_my_goods.test, exp_ui_catalog.test, exp_search_out_of_stock.default, exp_profile_settings_email.default, exp_cl_omni_refusalprintreceipts.test, exp_cl_omni_refusalprintcoupons.test, exp_accrual_history.test, exp_personal_recommendations.test_B, exp_newui_chips.test, exp_loyalty_categories.test, exp_growthbooks_aa.OFF, exp_test_ch_web.def, exp_search_suggestions_popular_sku.default, exp_cancel_subscription.test_2, exp_manage_subscription.control, exp_cl_new_csi.default, exp_cl_new_csat.default, exp_delivery_price_info.default, exp_personal_promo_navigation.test, exp_web_feature_test.true, exp_interval_jump.default, exp_cardOne_promo_type.test, exp_qr_cnc.test, exp_popup_about_order.control, exp_apigw_recommendations.test, exp_where_place_cnc.test, exp_editing_cnc_onboarding.default, exp_editing_cnc.default, exp_selection_carousel.test, exp_pickup_in_delivery.false, exp_feature_kpp_test.false, exp_welcome_onboarding.default, exp_cl_new_splash.default, exp_web_referral_program_type.default, exp_where_place_new.default, exp_start_page.default, exp_promocode_bd_coupon.default, exp_personal_promo_swipe_animation.default, exp_default_payment_type.default, exp_main_page_carousel_vs_banner.default, exp_start_page_onboarding.default, exp_newui_cart_check_edit.default, exp_search_new_logic.default, exp_search_ds_pers_similar.default, exp_growthbooks_aa_id_based_feature.test, exp_referral_program_type.default, exp_new_action_pages.default, exp_my_choice_search.default, exp_items_by_rating.default, exp_can_accept_early.default, exp_test_gb_value.false, exp_online_subscription.default, exp_new_nps_keyboard.test, exp_web_b2b_cancel_to_edit.test, exp_main_page_carousel_vs_banner_shop.default, exp_bathcing.default, exp_web_qr_cnc.default, exp_hide_cash_payment_for_cnc_wo_adult_items.default, exp_web_promocode_bd_coupon.default, exp_prices_per_quantum.test, exp_test.default123, exp_web_partner_coupons_separately.default, exp_web_chips_online.default, exp_b2b_web_redesign_reg.default, exp_chips_online.default, exp_promo_without_benefit.default")
	req.Header.Set("Origin", "https://lenta.com")
	req.Header.Set("Referer", urlFull)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("SessionToken", sessionToken)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36")
	req.Header.Set("X-Delivery-Mode", "pickup")
	req.Header.Set("X-Device-Brand", "")
	req.Header.Set("X-Device-ID", "1b642436-ab55-510a-a1c1-942a3cb468f8")
	req.Header.Set("X-Device-Name", "")
	req.Header.Set("X-Device-OS", "Web")
	req.Header.Set("X-Device-OS-Version", "12.4.8")
	req.Header.Set("X-Domain", city)
	req.Header.Set("X-Organization-ID", "")
	req.Header.Set("X-Platform", "omniweb")
	req.Header.Set("X-Retail-Brand", "lo")
	req.Header.Set("baggage", "sentry-environment=production,sentry-release=web-12.0.71,sentry-public_key=b99355c72549498d9e9075cc3d4006a2,sentry-trace_id=7bd91d457494451ab69d69492e439e99,sentry-sample_rate=1,sentry-transaction=%2Fcatalog%2F%3AcategoryId%2F,sentry-sampled=true")
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Google Chrome";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sentry-trace", "7bd91d457494451ab69d69492e439e99-9af093423f864258-1")
	req.Header.Set("traceparent", "00-8ca85812582cda5656d33a26a5de8296-96c93441e6954886-01")
	req.Header.Set("x-span-id", "96c93441e6954886")
	req.Header.Set("x-trace-id", "8ca85812582cda5656d33a26a5de8296")

	cookieString := `OmniPercent=73; OmniGroup=true; cookiesession1=678B2889D184DAF529E8F00FDD12A09C; App_Cache_MPK=mp300-b1de0bac2c257f3257bf5ef2eea4ecbc; App_Cache_LegalEntityId=; App_Cache_PassportId=; GrowthBook_user_id=1f5f81c9-1f79-7dd6-1623-f495be0da38e; oxxfgh=2d1bec6e-7d7b-4353-9211-caf306b8089c%230%235184000000%235000%231800000%2344965; uwyii=f63473c8-9786-7d16-cd0f-5afe98fd11c6; Utk_SessionToken=0000000000; iap.uid=70c720e6a1844eeb837c2ba4f62a967a; agree_with_cookie=true; _ga=GA1.1.800982013.1740987886; tmr_lvid=6f34000af21c40f487a51f4a5847ccc9; tmr_lvidTS=1740987886562; flocktory-uuid=9364f4c1-b109-4e70-8a50-e66dbc779e11-5; _ym_uid=1740987887107533127; _ym_d=1740987887; adrcid=AlW2OVF5X7hIGZMnlw0kF2A; App_Cache_DeliveryMode=%7B%22sessionToken%22%3A%220000000000%22%2C%22type%22%3A%22pickup%22%2C%22storeId%22%3A3258%2C%22addressId%22%3Anull%7D; App_Cache_PickupAddress=%7B%22id%22%3A3258%2C%22alias%22%3A%220279%22%2C%22title%22%3A%22%D0%A2%D0%9A279%22%2C%22marketType%22%3A%22HM%22%2C%22storeClose%22%3A%2223%3A00%22%2C%22storeOpen%22%3A%2208%3A00%22%2C%22pickupWindowOpen%22%3A%2208%3A00%22%2C%22pickupWindowClose%22%3A%2223%3A00%22%2C%22regionId%22%3A115%2C%22addressShort%22%3A%22%D0%A7%D0%B5%D1%80%D0%BA%D0%B5%D1%81%D1%81%D0%BA%2C%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%D0%B0%20%D1%83%D0%BB.%2C%20387%22%2C%22addressFull%22%3A%22%D0%A7%D0%B5%D1%80%D0%BA%D0%B5%D1%81%D1%81%D0%BA%2C%20%D0%9B%D0%B5%D0%BD%D0%B8%D0%BD%D0%B0%20%D1%83%D0%BB.%2C%20387%22%2C%22longitude%22%3A42.03778851%2C%22latitude%22%3A44.19543055%2C%22pickupEnable%22%3Atrue%7D; domain_sid=XqdVevJT5xp6ZrVqYJpnb%3A1741105117715; App_Cache_City=%7B%22id%22%3A1%2C%22name%22%3A%22%D0%9C%D0%BE%D1%81%D0%BA%D0%B2%D0%B0%20%D0%B8%20%D0%9C%D0%9E%22%2C%22slug%22%3A%22moscow%22%2C%22centerLat%22%3A%2255.75322000%22%2C%22centerLng%22%3A%2237.62255200%22%2C%22isDefault%22%3Atrue%2C%22mainDomain%22%3Afalse%7D; App_Cache_CitySlug=moscow; Utk_MrkGrpTkn=DDEA666291ED2E04F126AB6BA1C68247; spses.d58d=*; acs_3=%7B%22hash%22%3A%22be483547539f1e5fb43aa6ae1ea56ef0a5c5be24%22%2C%22nst%22%3A1741272160854%2C%22sl%22%3A%7B%22224%22%3A1741185760854%2C%221228%22%3A1741185760854%7D%7D; adrdel=1741185760883; _ym_isad=2; User_Agent=Mozilla%2F5.0%20(X11%3B%20Linux%20x86_64)%20AppleWebKit%2F537.36%20(KHTML%2C%20like%20Gecko)%20Chrome%2F133.0.0.0%20Safari%2F537.36; Is_Search_Bot=false; Utk_DvcGuid=1b642436-ab55-510a-a1c1-942a3cb468f8; _utm_referrer=; rr-VisitorSegment-splitterGet=exp_apigw_listing_Search%2Cedit_10.02_control_B2B_Checkout%2Cnew_B2B_Reserve; tmr_detect=0%7C1741187184893; qrator_jsid=1741185758.184.ke102z0yb5qgfZQr-2dv4fbkduusbj3t8e1slltuiirc3eg1t; GrowthBook_experiments=exp_personal_recommendations.2%2Cweb_exp.1%2Cexp_qr_cnc.1%2Cexp_where_place_cnc.1%2Cexp_growthbooks_aa_id_based.1%2Cexp_web_b2b_cancel_to_edit.1%2Cexp_prices_per_quantum.1; _ga_QB4J0GGLMG=GS1.1.1741185760.6.1.1741190770.0.0.0; SOURCE_ID_time=2025-03-05%2019%3A06%3A10; uwyiert=705a9a51-c071-ede9-8656-2bd1d7aba432; spid.d58d=1f25636a-7887-48cc-baad-bd5f957f5b69.1740987887.5.1741190770.1741107198.1fa1f9a7-1e4e-4d24-8280-8a2113118f52.a348e4c0-5d62-4df4-8b23-235dc25a26ce.a9c98527-38ac-4762-acf8-e8ca1130e512.1741185760588.153`

	updatedCookieString := fmt.Sprintf(
		strings.Replace(
			strings.Replace(
				cookieString,
				"Utk_SessionToken=0000000000",
				fmt.Sprintf("Utk_SessionToken=%s", sessionToken),
				1,
			),
			"sessionToken%%22%%3A%%220000000000%%",
			fmt.Sprintf("sessionToken%%22%%3A%%22%s%%", sessionToken),
			1,
		),
	)

	updatedCookieString = strings.Replace(
		strings.Replace(
			updatedCookieString,
			`"slug":"moscow"`,
			fmt.Sprintf(`"slug":"%s"`, city),
			1,
		),
		"App_Cache_CitySlug=moscow",
		fmt.Sprintf("App_Cache_CitySlug=%s", city),
		1,
	)

	req.Header.Set("Cookie", updatedCookieString)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)

	var response Response
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}
	/* отладка
	fmt.Println("\nПродукты:")
	for _, product := range response.Items {
		costInRubles := float64(product.Prices.Cost) / 100 // Переводим стоимость в рубли
		fmt.Printf("ID: %d, Name: %s, Slug: %s, Cost: %.2f руб.\n", product.ID, product.Name, product.Slug, costInRubles)

		urlItem := fmt.Sprintf("https://lenta.com/product/%s-%d", product.Slug, product.ID)

		fmt.Printf(urlItem + "\n")
	}
	*/

	filePath := fmt.Sprintf("%s_%s_products.csv", url, city)

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Comma = ';'

	headers := []string{"Name", "Cost", "URL"}
	if err := writer.Write(headers); err != nil {
		fmt.Println("Ошибка при записи заголовков:", err)
		return
	}

	for _, product := range response.Items {
		costInRubles := float64(product.Prices.Cost) / 100 // Переводим стоимость в рубли
		urlItem := fmt.Sprintf("https://lenta.com/product/%s-%d", product.Slug, product.ID)
		record := []string{
			product.Name,
			fmt.Sprintf("%.2f", costInRubles),
			urlItem,
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Ошибка при записи данных:", err)
			return
		}
	}

	fmt.Println("Данные успешно сохранены в файл products.csv")
}

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Prices struct {
		Cost int `json:"cost"`
	} `json:"prices"`
}

type Response struct {
	Items []Product `json:"items"`
}
