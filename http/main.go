package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	_ "net/http/pprof"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com", nil)

	request.Header.Add("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect req: ", req)

			// for i, v := range via {
			// 	fmt.Println("Redirect via: ", i, v)
			// }

			return nil
		},
	}
	resp, err := client.Do(request)
	// resp, err := http.DefaultClient.Do(request)
	// resp, err := http.Get("https://www.glosku.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	// s, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
