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
