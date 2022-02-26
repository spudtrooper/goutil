package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	uri := request.CreateRoute("https://www.vividseats.com/hermes/api/v1/productions",
		request.Param{"startDate", `2022-02-26`},
		request.Param{"endDate", `2022-03-05`},
		request.Param{"sortBy", `RANK`},
		request.Param{"pageSize", 4},
		request.Param{"includeIpAddress", true},
		request.Param{"radius", 80450},
		request.Param{"excludeParking", true},
	)
	type cookiePart struct{ key, val string }
	var cookieParts = []cookiePart{
		{"_ga", `GA1.2.2022141717.1645470184`},
		{"__insp_wid", `1561680148`},
		{"__insp_nv", `true`},
		{"__insp_targlpu", `aHR0cHM6Ly9za3lib3gudml2aWRzZWF0cy5jb20vcmVxdWVzdC1hY2NvdW50`},
		{"__insp_targlpt", `UmVxdWVzdCBhbiBBY2NvdW50IC0gU2t5Ym94`},
		{"__insp_norec_sess", `true`},
		{"__insp_slim", `1645470698786`},
		{"userAgent", `%7B%22headerString%22%3A%22Mozilla%2F5.0%20%28Macintosh%3B%20Intel%20Mac%20OS%20X%2010_15_7%29%20AppleWebKit%2F537.36%20%28KHTML%2C%20like%20Gecko%29%20Chrome%2F98.0.4758.102%20Safari%2F537.36%22%2C%22name%22%3A%22Chrome%22%2C%22majorVersion%22%3A98%2C%22minorVersion%22%3A0%2C%22patchVersion%22%3A4758%2C%22deviceType%22%3A%22desktop%22%2C%22deviceName%22%3A%22Mac%22%2C%22osName%22%3A%22Mac%20OS%20X%22%2C%22osMajorVersion%22%3A10%2C%22osMinorVersion%22%3A15%2C%22osPatchVersion%22%3A7%2C%22supported%22%3Atrue%7D`},
		{"clientId", `be71761f-11b2-4269-98a0-1cadb16bdfae`},
		{"lastvisit", `2022-02-21%2015%3A27%3A55`},
		{"vtrk", `v_ref%3Dnull%7Cv_camp%3Dnull%7Cv_cont%3Dnull%7Cv_med%3Dnull%7Cv_src%3Dnull%7Cv_trm%3Dnull%7Cv_kid%3Dnull%7Cgclid%3Dnull%7Cadgroup%3Dnull%7Ctarget%3Dnull%7Cdevice%3Dnull%7CkeywordId%3Dnull`},
		{"optimizely-user-id", `dab2a77d-4113-4bac-8a94-bcea663b99ba`},
		{"VS_SID", `f-1`},
		{"visid_incap_2441468", `8WVhY9yYQ/KcFpel9Hi82coDFGIAAAAAQUIPAAAAAAB4mUh298NsiSFvI4KNbHFd`},
		{"nlbi_2441468", `DSWNcu+YXWa7D2CPIZf6kQAAAAAZf6PmcAX7O1oWbMlS8ukH`},
		{"_pxhd", `LbWY6Ct0S47/74/hVrAzqxIYyyCAAhgoJuSHblkkQNZBLmR0eAbKyx/00QJN1nVsmxxPyLbn2b7V/ms03znaSw==:QQaDznvdz6OYM4nKrpp4sh/KZb3dY-PhWgMHCIBnyXq5iakTfE3woM-kfNJ1FQtQcYYwhRfpO64fSMDQxi/cUGPLtbdzBRpxspEkQDf6/Ec=`},
		{"optimizelyEndUserId", `oeu1645478881821r0.03151458645263472`},
		{"_gcl_au", `1.1.687885081.1645478882`},
		{"IR_gbd", `vividseats.com`},
		{"pxcts", `263172b6-935d-11ec-9c47-437048555047`},
		{"_pxvid", `224bf112-935d-11ec-ab33-4b4441766361`},
		{"_fbp", `fb.1.1645478882373.1981351338`},
		{"notice_behavior", `none`},
		{"optimizely_uuid", `6fc5a4d7-5c2d-43c6-8ba0-a9780ce5e6ff`},
		{"x-device-analytics", `%7B%22customerId%22%3A14723639%2C%22hashedEmailAddress%22%3A%22b07fc737e86096e228db2417435bc844b5809300%22%7D`},
		{"lcId", `b07fc737e86096e228db2417435bc844b5809300`},
		{"ntat", `%7B%22accountId%22%3A14723639%2C%22token%22%3A%228kcih3bf0n9tm9astrg5r8nklud6t4ae2npopmsh5j3k3u5ji1no%22%7D`},
		{"optimizely_target", `true`},
		{"nlbi_2441468_2147483392", `Uv2HdRqPYl0K7gP+IZf6kQAAAAAg3/PWESZQwDfYh9fLyhgv`},
		{"reese84", `3:kAWjsSc3YlioyXoZ8pUgFA==:+U2r2yJV+y9eBhpi6wMW4te0Scv029ewqeAaSpxy4NwkAA6ZHPGHO1laT2hKLMbvDqF7xoRGe7K7+bc9wAKXTTPxPYsAS9umKfnow67k8RBWHEwRGDvHyGlcgmltWsjVNV2iJgSqQBAAr0SZ66j9AwUfE4qiTgMtG/q16h1h16CFor1pYyBZ7/uuqBwaM519vel9v4rG8T6XrRmfvctXrYDEQ3j0BBO+JNq6eGq9S7UPU4anqtQZp0bzjKL2Zao8oM7qSYpyARhmveS7A6U9lhQRvy78KZFZaklumyhSdEZ13LzGjoxlDyqn/uJOPwuPnyBFLwqd5i4QUMP/JymGUK2oEhA9BhqEMlrjYYyEMUqZdO/KI+SkeiSv59WKa5r13pgtm0cxfTYUWc3vXgWgZjatBl28QA7cz/wiiFIJEVDfKu4VJ3oUmBHsDkr5O5UzFHGFtZMGEeM5SKPOKTV2dIYJCpqk+qmQ39m5w66QBYc=:kvl9wv2FXKCM0yxJuttcp349fxvYqwqfn9hGDlLUhyQ=`},
		{"ch", `%5B%7B%22d%22%3A%222022-02-21%2015%3A27%3A55%22%2C%22h%22%3A%22Direct%22%2C%22v%22%3A0%2C%22b%22%3Afalse%2C%22ac%22%3Atrue%7D%2C%7B%22d%22%3A%222022-02-23%2005%3A20%3A49%22%2C%22h%22%3A%22Organic%20Search%22%2C%22a%22%3A%22www.google.com%22%2C%22k%22%3A%22%22%2C%22v%22%3A0%2C%22b%22%3Afalse%2C%22ac%22%3Atrue%7D%5D`},
		{"userData", `%7B%22uuid%22%3A%226fc5a4d7-5c2d-43c6-8ba0-a9780ce5e6ff%22%2C%22regionId%22%3A7%2C%22secondaryRegionId%22%3A0%2C%22tertiaryRegionId%22%3A0%2C%22inboundPhoneNumber%22%3A%22800-501-9983%22%2C%22newSession%22%3Afalse%2C%22orInit%22%3Atrue%2C%22regionName%22%3A%22New%20York%20City%22%7D`},
		{"_abck", `F540770CF239729AA1A9BFC74FA4765F~-1~YAAQvUhyaNV2CyF/AQAAYya6NgfVjalHa1lZgmHfzr+m/yL0XQjd/k9rB+8eicHQ6khx/xaYtzJHAyIAN435SerL04NGLpNet/zDOYqA78BjN4+lQGA2cdTMo5nRSEx+0uIRCRRelV4wZvMmo6KXsti5cKltYvMWQp/OwBBdYcCvKhP4nj4B8FXDlnLiW0T0I631rDkJ+Onc1Lr2rFmcwKBKMNCwMw16edRCu5/BjuRflJBSf71NAP4Kq/tbIYfpUD+J1t4lgeliKZNKo+iK/h65YfBS8N37xTCDJw42lhl2skbSL1kFXA4uOiafq4GmoaHOW+vcjlVAlkZeYgfeuXMv1KcCTHMxyGsDiL9bTjQFo//F2b6LmtbRlDx696mhn7ZGj4lyF9U0LsmliVk=~-1~-1~-1`},
		{"bm_sz", `B01213BCF77D33469B7BEE75CDEB6EE1~YAAQvUhyaNZ2CyF/AQAAYya6Ng6/uE7LSKu9ZOuji1L05nBhpEFa/l2lp77xhZcxyhLx4iZn00av/6YJ5SCwG5kjv6O6RPsqlAO2jaeWlSlFLVEdgSr3B8DXwEXIyL6xSqOJ6cr/5eGcK2Da2yfdjIh30MTmXDt5wAoMVlRdxnElluwqMlPV/ROt2lGn2Ay7qcLKevojMMPdxWrfqkKwG3UMMCQvRlefAwK0xNTKxXnhX3DMcISOUQ1oBJM1SLLCrPlHkL4XpwHzXDTCK4bzYlplu83Wx3Bzi5S5qZf+b5Zk7yrMIpHf~3163193~3688006`},
		{"ak_bmsc", `5E1AC118ED9BEF14CABA835D53AF82E7~000000000000000000000000000000~YAAQVIIsFw4zCQ5/AQAAbii6Ng5eQNVv7GAjVZJ6GHez5694ocDHHFjgJRZXz9eEbpAhUeKvEUFLf1H4TorOWbkTTG2WQl9K5WhO5rN0+3gvjseLkctNLPSD20TdAtVSh4EPF4S6mXCOWHvquiWM4jKmfvckAh+L/X0ia+J0JTIURoAoeeZGJXszueNuTThyQmC4isEl+p1f+VjOjBYkKy2bV9DkufZ63zqo50SYKnDw0jJ0a2tiyojMnmnhS57hR6+9lteyzem+9xW60ov65txBxa/B/OTc+q3Zyz1BTy4H2JWZf/m0TGXSuLYB40sOLEnLt3Z9YBMGOxYNd5TKiNZKPlAvpH9R9BDgWZXjcPisbRUbKDNVRe2fpIvCssz7xPMzKlbtCukuMNidTg==`},
		{"_gid", `GA1.2.1088841078.1645890644`},
		{"_gat", `1`},
		{"ab.storage.deviceId.b8f743b5-775e-4ede-a27c-411f9f57a648", `%7B%22g%22%3A%226e922eb0-b923-851b-319d-8dccddbd1aaa%22%2C%22c%22%3A1645478882300%2C%22l%22%3A1645890644451%7D`},
		{"ab.storage.userId.b8f743b5-775e-4ede-a27c-411f9f57a648", `%7B%22g%22%3A%22B07FC737E86096E228DB2417435BC844B5809300%22%2C%22c%22%3A1645478910943%2C%22l%22%3A1645890644451%7D`},
		{"_clck", `1j1ij0j|1|ezb|0`},
		{"VIVID-CSRFTOKEN", `3edpc56ihf55ci9vg8d0qsj7dulfbnjf12rtv7g6vr98lin1migs`},
		{"page-view-count", `1`},
		{"_pxff_fp", `1`},
		{"ab.storage.sessionId.b8f743b5-775e-4ede-a27c-411f9f57a648", `%7B%22g%22%3A%22653f6894-77d0-1ed0-e041-1be52a31017b%22%2C%22e%22%3A1645892455444%2C%22c%22%3A1645890644450%2C%22l%22%3A1645890655444%7D`},
		{"_clsk", `s741l0|1645890655952|4|0|f.clarity.ms/collect`},
		{"_px2", `eyJ1IjoiZTIwN2RjYzAtOTcxYi0xMWVjLTlmMGUtN2RlNzhiN2M2MjU1IiwidiI6IjIyNGJmMTEyLTkzNWQtMTFlYy1hYjMzLTRiNDQ0MTc2NjM2MSIsInQiOjE2NDU4OTA5NTY0NTUsImgiOiI5OTljZTA2MjE4NmQ4ZWJjNGI0YjM5MzhlZTJjY2NlNzMzYjQ3NzEwYjJhODc4ZTI3NGQzNGIxMmM3ZjA4NzU3In0=`},
		{"JSESSIONID", `BEA0E2B26C9E19B1DF752F0CAC8719F8`},
		{"IR_12730", `1645890657429%7C0%7C1645890657429%7C%7C`},
		{"_uetsid", `db7df7f0971b11eca221e9b795714a01`},
		{"_uetvid", `26220ae0935d11ecb5a6a9491497efe3`},
		{"bm_sv", `40F26A39EE919C1FE54F4D12D1788F1A~ZW4C4Td3C1dm/I2kRvDqMzpHvncG9DVVyhVnQF4k1oG5uWGNcMbM0o6TcqHBHoEkRDzTk+bBZ6HXKoM4FlFPVSIW8P0l52BweEYzWT7JFip3mipipLcW+vvQi8z9LJcaSOP9E7oFqdO/1Oc5L2tJkWf7TmAChFBlo497xqNI3YQ=`},
		{"cto_bundle", `aJN23F9xJTJGQkR2JTJGd1JKekcwYlE5JTJCY0ZONzl0aVIlMkZyVWRIT04yVCUyRnZDN1AlMkJEY2d3ZjZicnJ2ZVF3NHBtU0hYWUo5TnN4Y2F6azhFZXVac2tYRnQ5ankxcjkwY29IV3l5MUp5Z1I3QTVKYjJnY3NRZVhybnl4TXd1cktVWUM0cEFNNFpzYjNVWTdCeTV1UkNsS0d6M1pRS3AlMkZIZyUzRCUzRA`},
	}
	headers := map[string]string{
		"authority":          `www.vividseats.com`,
		"pragma":             `no-cache`,
		"cache-control":      `no-cache`,
		"sec-ch-ua":          `" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"`,
		"accept":             `application/json`,
		"dnt":                `1`,
		"sec-ch-ua-mobile":   `?0`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-site":     `same-origin`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-dest":     `empty`,
		"referer":            `https://www.vividseats.com/`,
		"accept-language":    `en-US,en;q=0.9`,
	}

	serializeCookie := func() string {
		if len(cookieParts) > 0 {
			var cs []string
			for _, c := range cookieParts {
				cs = append(cs, fmt.Sprintf("%s=%s", c.key, c.val))
			}
			return strings.Join(cs, "; ")
		}
		return ""
	}
	if c := serializeCookie(); c != "" {
		headers["cookie"] = c
	}
	var payload interface{}
	res, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers))
	check.Err(err)
	log.Printf("result: %+v", res)
	log.Printf("payload: %s", request.MustFormatString(payload))
}
