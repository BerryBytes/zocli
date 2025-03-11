package defaults

// headersAndBodyConv
//
// is responsible for parsing the headers and body values from the,
// map[string]interface{}
func headersAndBodyConv(args ...map[string]interface{}) (map[string]string, []byte) {
	var headers map[string]string
	var body []byte

	for _, single := range args {
		if _, found := single["headers"]; found {
			headers = single["headers"].(map[string]string)
		}
		if _, found := single["body"]; found {
			body = single["body"].([]byte)
		}
	}

	return headers, body
}
