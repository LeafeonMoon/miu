package uhttp

// import (
// 	"hp/src/main/utils/ujson"
// 	"hp/src/main/utils/ustring"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// const HTTP string = "http://"
// const HTTPS string = "https://"

// const POST string = "POST"
// const GET string = "GET"

// const CONTENT_TYPE string = "Content-Type"

// const APPLICATION_FORM string = "application/x-www-form-urlencoded"
// const APPLICATION_JSON string = "application/json"
// const APPLICATION_JSON_UTF8 string = "application/json;charset=utf-8"

// func Get(requestUrl string, params map[string]interface{}, headers map[string]interface{}) (string, error) {
// 	uri, err := url.ParseRequestURI(requestUrl)
// 	if err != nil {
// 		return "Request url error", err
// 	}
// 	urlStr := uri.String()
// 	req, err := http.NewRequest(GET, urlStr, nil)
// 	if err != nil {
// 		return "New Request error", err
// 	}
// 	setQuery(req, params)
// 	setHeader(req, headers)
// 	return clientDo(req)
// }

// func Post(requestUrl string, params string, headers map[string]interface{}) (string, error) {
// 	ct := headers[CONTENT_TYPE]
// 	if ct != nil {
// 		if APPLICATION_JSON == ct || APPLICATION_JSON_UTF8 == ct {
// 			// post
// 			paramMap, err := ujson.ToMap(params)
// 			if err != nil {
// 				return "error", err
// 			}
// 			return PostJson(requestUrl, paramMap, headers)
// 		} else {
// 			// form
// 			paramMap, err := ujson.ToMap(params)
// 			if err != nil {
// 				return "error", err
// 			}
// 			return PostForm(requestUrl, paramMap, headers)
// 		}
// 	} else {
// 		paramMap, err := ujson.ToMap(params)
// 		if err != nil {
// 			return "error", err
// 		}
// 		_, _, err = setJson(paramMap)
// 		if err == nil {
// 			// post
// 			return PostJson(requestUrl, paramMap, headers)
// 		} else {
// 			// form
// 			return PostForm(requestUrl, paramMap, headers)
// 		}
// 	}
// }

// func PostForm(requestUrl string, params map[string]interface{}, headers map[string]interface{}) (string, error) {
// 	uri, err := url.ParseRequestURI(requestUrl)
// 	if err != nil {
// 		return "Request url error", err
// 	}
// 	req, err := http.NewRequest(POST, uri.String(), setData(params))
// 	if err != nil {
// 		return "New Request error", err
// 	}
// 	req.Header.Add(CONTENT_TYPE, APPLICATION_FORM)
// 	setHeader(req, headers)
// 	return clientDo(req)
// }

// func PostJson(requestUrl string, params map[string]interface{}, headers map[string]interface{}) (string, error) {
// 	uri, err := url.ParseRequestURI(requestUrl)
// 	if err != nil {
// 		return "Request url error", err
// 	}
// 	_, reader, err := setJson(params)
// 	if err != nil {
// 		return "Format json error", err
// 	}
// 	req, err := http.NewRequest(POST, uri.String(), reader)
// 	if err != nil {
// 		return "New Request error", err
// 	}
// 	req.Header.Add(CONTENT_TYPE, APPLICATION_JSON_UTF8)
// 	setHeader(req, headers)
// 	return clientDo(req)
// }

// func setHeader(request *http.Request, reqHeaders map[string]interface{}) {
// 	if len(reqHeaders) > 0 {
// 		for key, val := range reqHeaders {
// 			request.Header.Add(key, ustring.InterToString(val))
// 		}
// 	}
// }

// func setQuery(request *http.Request, params map[string]interface{}) {
// 	query := request.URL.Query()
// 	if len(params) > 0 {
// 		for key, val := range params {
// 			query.Add(key, ustring.InterToString(val))
// 		}
// 	}
// 	request.URL.RawQuery = query.Encode()
// }

// func setData(params map[string]interface{}) *strings.Reader {
// 	data := url.Values{}
// 	if len(params) > 0 {
// 		for key, val := range params {
// 			data.Add(key, ustring.InterToString(val))
// 		}
// 	}
// 	return strings.NewReader(data.Encode())
// }

// func setJson(params map[string]interface{}) (string, *strings.Reader, error) {
// 	data := url.Values{}
// 	if len(params) > 0 {
// 		for key, val := range params {
// 			data.Add(key, ustring.InterToString(val))
// 		}
// 	}
// 	json, err := ujson.ToJson(params)
// 	if err != nil {
// 		return "error", nil, err
// 	}
// 	return json, strings.NewReader(json), nil
// }

// func clientDo(request *http.Request) (string, error) {
// 	client := &http.Client{}
// 	res, err := client.Do(request)
// 	if err != nil {
// 		return "error", err
// 	}

// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return "error", err
// 	}
// 	return string(body), nil
// }
