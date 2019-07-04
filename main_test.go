package main

import (
	"os"
	"strings"
)

func ExampleProcess() {
	input := strings.NewReader(`
diff -u pkg/xerohooks/handler.go.orig pkg/xerohooks/handler.go
--- pkg/xerohooks/handler.go.orig	2019-07-05 09:46:16.000000000 +1000
+++ pkg/xerohooks/handler.go	2019-07-05 09:46:16.000000000 +1000
@@ -15,8 +15,8 @@
 	*mux.Router
 }
 
-func NewHandler    (api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {
-	r    := mux.NewRouter()
+func NewHandler(api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {
+	r := mux.NewRouter()
 	register(api, r, sqser, snser, lambdaer, transport)
 	return &Handler{r}
 }
diff -u pkg/xerohooks/sqs.go.orig pkg/xerohooks/sqs.go
--- pkg/xerohooks/sqs.go.orig	2019-07-05 09:46:16.000000000 +1000
+++ pkg/xerohooks/sqs.go	2019-07-05 09:46:16.000000000 +1000
@@ -15,7 +15,7 @@
 	queueUrl string
 }
 
-func     newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {
+func newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {
 	arnBits := strings.Split(queueArn, ":")
 	region := arnBits[3]
 	accountId := arnBits[4]
`)
	output := os.Stdout
	process(input, output)

	// Output:
	// ##teamcity[inspectionType category='Code style' description='`gofmt` style violation' id='gofmt' name='gofmt']
	// ##teamcity[inspection file='pkg/xerohooks/handler.go' line='18' message=' 	*mux.Router|n }|n |n-func NewHandler    (api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {|n-	r    := mux.NewRouter()|n+func NewHandler(api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {|n+	r := mux.NewRouter()|n 	register(api, r, sqser, snser, lambdaer, transport)|n 	return &Handler{r}|n }|n' typeId='gofmt']
	// ##teamcity[inspection file='pkg/xerohooks/handler.go' line='19' message=' 	*mux.Router|n }|n |n-func NewHandler    (api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {|n-	r    := mux.NewRouter()|n+func NewHandler(api *ApiDefinition, sqser Sqser, snser Snser, lambdaer Lambdaer, transport http.RoundTripper) *Handler {|n+	r := mux.NewRouter()|n 	register(api, r, sqser, snser, lambdaer, transport)|n 	return &Handler{r}|n }|n' typeId='gofmt']
	// ##teamcity[inspection file='pkg/xerohooks/sqs.go' line='18' message=' 	queueUrl string|n }|n |n-func     newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {|n+func newSqsHandler(sqsApi sqsiface.SQSAPI, queueArn string) *sqsHandler {|n 	arnBits := strings.Split(queueArn, ":")|n 	region := arnBits|[3|]|n 	accountId := arnBits|[4|]|n' typeId='gofmt']
}
