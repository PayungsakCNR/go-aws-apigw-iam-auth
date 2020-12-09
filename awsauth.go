package awsauth

import (
	"net/http"
)

//Sign Function
func Sign(request *http.Request, region string, accessKey string, secretKey string) *http.Request {

	prepareRequest(request)

	meta := new(metadata)
	meta.service = "execute-api"
	meta.region = region

	// Task 1
	hashedCanonReq := hashedCanonicalRequest(request, meta)

	// Task 2
	stringToSign := stringToSign(request, hashedCanonReq, meta)

	// Task 3
	signingKey := signingKey(secretKey, meta.date, region, meta.service)

	signature := signature(signingKey, stringToSign)

	// Set up Authorization Header
	request.Header.Set("Authorization", buildAuthHeader(signature, meta, accessKey))

	return request
}

type metadata struct {
	algorithm       string
	credentialScope string
	signedHeaders   string
	date            string
	region          string
	service         string
}
