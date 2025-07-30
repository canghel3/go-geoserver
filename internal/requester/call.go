package requester

//func call(data internal.GeoserverData, url, method string, headers map[string]string, body ...byte) (*http.Response, error) {
//	request, err := http.NewRequest(method, url, bytes.NewReader(body))
//	if err != nil {
//		return nil, err
//	}
//
//	request.SetBasicAuth(data.Connection.Credentials.Username, data.Connection.Credentials.Password)
//
//	for k, v := range headers {
//		request.Header.Set(k, v)
//	}
//
//	response, err := data.Client.Do(request)
//	if err != nil {
//		return nil, err
//	}
//
//	return response, nil
//}
