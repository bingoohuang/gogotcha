package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/Masterminds/goutils"
)

func unitHTTPClient(unixPath, urlAddr, postBodyRaw string) {
	fmt.Println("Unix HTTP client")

	hc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unixPath)
			},
		},
	}

	unixHTTPUrl := "http://unix" + urlAddr

	var (
		response *http.Response
		err      error
	)

	if postBodyRaw == "" {
		response, err = hc.Get(unixHTTPUrl) // nolint:bodyclose
	} else {
		response, err = hc.Post(unixHTTPUrl, // nolint:bodyclose
			"application/octet-stream", strings.NewReader(postBodyRaw))
	}

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	dump, _ := httputil.DumpResponse(response, true)
	abbr, _ := goutils.Abbreviate(string(dump), 900)
	fmt.Printf("%s\n", abbr)
}
