// 
// helper.go : Main sources driver. Contains helper functions for other sources.
// Written By : @ice3man (Nizamul Rana)
// 
// Distributed Under MIT License
// Copyrights (C) 2018 Ice3man
//

package helper

import (
	"net/http"
	"crypto/tls"
	"time"
)

// 
// GetHTTPResponse : Returns a HTTP Response object
// @param url : URL To Visit (Note, It needs full url with scheme)
// @param timeout : Seconds to wait for response until timeout
// 
// @return resp : HTTP Response object
// @return err : nil if successfull else error
//
func GetHTTPResponse(url string, timeout int) (resp *http.Response, err error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(timeout) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return resp, err
	}

	// TODO : Figure out a way to handle user agents as per user intention
	// @codingo, I don't think it's correct to spam services by making requests with fake user agent
	// 	What do you think
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.1) Gecko/2008071615 Fedora/3.0.1-1.fc9 Firefox/3.0.1")
	req.Header.Add("Connection", "close")

	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
