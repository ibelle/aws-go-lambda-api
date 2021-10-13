

 aws lambda create-function --function-name menus --runtime go1.x --role arn:aws:iam::562697480070:role/lambda-exec-role --handler main --zip-file fileb:///Users/isaiahbelle/Dev/gourmeals/aws_examples/lambda/go_api/menus/main.zip


# Build and Upload Update
env GOOS=linux GOARCH=amd64 go build -o /tmp/main menus
zip -j main.zip /tmp/main

aws lambda update-function-code --function-name menus --zip-file fileb://main.zip

# Call that shit
aws lambda invoke --function-name menus ./tmp.yaml

 aws dynamodb create-table --table-name Menus \
--attribute-definitions AttributeName=MENUID,AttributeType=S \
--key-schema AttributeName=MENUID,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

aws dynamodb put-item --table-name Menus --item '{"MENUID":{"S": "978-1420931693"}, "restaurant":{"S": "The Republic"}, "cuisine":{"S": "Ameican-Casual"}}'

aws dynamodb put-item --table-name Menus --item '{"MENUID":{"S": "888-123454249"}, "restaurant":{"S": "Roma"}, "cuisine":{"S": "Italian"}}'


# Query: aws dynamodb get-item --table-name Menus --key='{"MENUID": {"S": "978-1420931693"}}'

 env GOOS=linux GOARCH=amd64 go build -o /tmp/main menus
 zip -j /tmp/main.zip /tmp/main



 aws lambda update-function-code --function-name menus --zip-file fileb:///Users/isaiahbelle/Dev/gourmeals/aws_examples/lambda/go_api/menus/main.zip

  aws iam put-role-policy --role-name lambda-exec-role --policy-name dynamodb-item-crud-role --policy-document file://db-policy.json


  # Create API Gateway
  aws apigateway create-rest-api --name foodstore
  - apiKeySource: HEADER
  createdDate: '2021-10-12T16:27:31-04:00'
  disableExecuteApiEndpoint: false
  endpointConfiguration:
    types:
    - EDGE
  id: ooyqsymb9g
  name: foodstore

  aws apigateway get-resources --rest-api-id ooyqsymb9g
  - items:
  - id: gowlm22jlf
    path: /

 aws apigateway create-resource --rest-api-id ooyqsymb9g --parent-id gowlm22jlf --path-part menus
 - id: 8zmr8h
  parentId: gowlm22jlf
  path: /menus
  pathPart: menus


aws apigateway put-method --rest-api-id ooyqsymb9g --resource-id 8zmr8h --http-method ANY --authorization-type AWS_IAM
- apiKeyRequired: false
  authorizationType: AWS_IAM
  httpMethod: ANY


aws apigateway put-integration --rest-api-id ooyqsymb9g --resource-id 8zmr8h --http-method ANY --type AWS_PROXY --integration-http-method POST --uri arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations
 cacheKeyParameters: []
  cacheNamespace: 8zmr8h
  httpMethod: POST
  passthroughBehavior: WHEN_NO_MATCH
  timeoutInMillis: 29000
  type: AWS_PROXY
  uri: arn:aws:apigateway:us-east-1:lambda:path/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations

# Grant Gateway Execute permission on the lambda func
aws lambda add-permission --function-name menus --statement-id HxWzzn8xzoCN5raa --action lambda:InvokeFunction --principal apigateway.amazonaws.com --source-arn 'arn:aws:execute-api:us-east-1:562697480070:ooyqsymb9g/*/*/*'
- Statement: '{"Sid":"HxWzzn8xzoCN5raa","Effect":"Allow","Principal":{"Service":"apigateway.amazonaws.com"},"Action":"lambda:InvokeFunction","Resource":"arn:aws:lambda:us-east-1:562697480070:function:menus","Condition":{"ArnLike":{"AWS:SourceArn":"arn:aws:execute-api:us-east-1:562697480070:ooyqsymb9g/*/*/*"}}}'


# Invoke the API Gateway
aws apigateway test-invoke-method --rest-api-id ooyqsymb9g --resource-id 8zmr8h --http-method "GET"
- body: '{"message": "Internal server error"}'
  headers:
    x-amzn-ErrorType: InternalServerErrorException
  latency: 948
  log: "Execution log for request fb2dedbf-d17c-4328-9c00-6dd42b59893f\nTue Oct 12
    21:09:59 UTC 2021 : Starting execution for request: fb2dedbf-d17c-4328-9c00-6dd42b59893f\nTue
    Oct 12 21:09:59 UTC 2021 : HTTP Method: GET, Resource Path: /menus\nTue Oct 12
    21:09:59 UTC 2021 : Method request path: {}\nTue Oct 12 21:09:59 UTC 2021 : Method
    request query string: {}\nTue Oct 12 21:09:59 UTC 2021 : Method request headers:
    {}\nTue Oct 12 21:09:59 UTC 2021 : Method request body before transformations:
    \nTue Oct 12 21:09:59 UTC 2021 : Endpoint request URI: https://lambda.us-east-1.amazonaws.com/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations\nTue
    Oct 12 21:09:59 UTC 2021 : Endpoint request headers: {X-Amz-Date=20211012T210959Z,
    x-amzn-apigateway-api-id=ooyqsymb9g, Accept=application/json, User-Agent=AmazonAPIGateway_ooyqsymb9g,
    Host=lambda.us-east-1.amazonaws.com, X-Amz-Content-Sha256=215f9ed8c3a67ac149bc802d7a46cd8784e38ee06a4998ca21bef5e6540e6c5d,
    X-Amzn-Trace-Id=Root=1-6165f9a7-a1752c122b4d6e39626f6591, x-amzn-lambda-integration-tag=fb2dedbf-d17c-4328-9c00-6dd42b59893f,
    Authorization=*********************************************************************************************************************************************************************************************************************************************************************************************************************************************72af27,
    X-Amz-Source-Arn=arn:aws:execute-api:us-east-1:562697480070:ooyqsymb9g/test-invoke-stage/GET/menus,
    X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJX//////////wEaCXVzLWVhc3QtMSJHMEUCIQCbcuHBW2jKXIHN8eWObRszoiNLniF/hQzCqm7CvtcHjgIgAlPMER/Q8Xgggfyxclq7pnM0c0I0u0Uyb+onXffs
    [TRUNCATED]\nTue Oct 12 21:09:59 UTC 2021 : Endpoint request body after transformations:
    {\"resource\":\"/menus\",\"path\":\"/menus\",\"httpMethod\":\"GET\",\"headers\":null,\"multiValueHeaders\":null,\"queryStringParameters\":null,\"multiValueQueryStringParameters\":null,\"pathParameters\":null,\"stageVariables\":null,\"requestContext\":{\"resourceId\":\"8zmr8h\",\"resourcePath\":\"/menus\",\"httpMethod\":\"GET\",\"extendedRequestId\":\"HHPyOEcEIAMF7xA=\",\"requestTime\":\"12/Oct/2021:21:09:59
    +0000\",\"path\":\"/menus\",\"accountId\":\"562697480070\",\"protocol\":\"HTTP/1.1\",\"stage\":\"test-invoke-stage\",\"domainPrefix\":\"testPrefix\",\"requestTimeEpoch\":1634072999663,\"requestId\":\"fb2dedbf-d17c-4328-9c00-6dd42b59893f\",\"identity\":{\"cognitoIdentityPoolId\":null,\"cognitoIdentityId\":null,\"apiKey\":\"test-invoke-api-key\",\"principalOrgId\":null,\"cognitoAuthenticationType\":null,\"userArn\":\"arn:aws:iam::562697480070:user/dev_system_account\",\"apiKeyId\":\"test-invoke-api-key-id\",\"userAgent\":\"aws-cli/2.2.43
    Python/3.9.7 Darwin/20.6.0 source/x86_64 prompt/off command/apigateway.test-invoke-method\",\"accountId\":\"562697480070\",\"call
    [TRUNCATED]\nTue Oct 12 21:09:59 UTC 2021 : Sending request to https://lambda.us-east-1.amazonaws.com/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations\nTue
    Oct 12 21:10:00 UTC 2021 : Received response. Status: 200, Integration latency:
    940 ms\nTue Oct 12 21:10:00 UTC 2021 : Endpoint response headers: {Date=Tue, 12
    Oct 2021 21:10:00 GMT, Content-Type=application/json, Content-Length=83, Connection=keep-alive,
    x-amzn-RequestId=88e8f4f4-a7a8-4500-81f5-986d4e912614, x-amzn-Remapped-Content-Length=0,
    X-Amz-Executed-Version=$LATEST, X-Amzn-Trace-Id=root=1-6165f9a7-a1752c122b4d6e39626f6591;sampled=0}\nTue
    Oct 12 21:10:00 UTC 2021 : Endpoint response body before transformations: {\"menuid\":\"978-1420931693\",\"restaurant\":\"The
    Republic\",\"cusisine\":\"Ameican-Casual\"}\nTue Oct 12 21:10:00 UTC 2021 : Execution
    failed due to configuration error: Malformed Lambda proxy response\nTue Oct 12
    21:10:00 UTC 2021 : Method completed with status: 502\n"
  multiValueHeaders:
    x-amzn-ErrorType:
    - InternalServerErrorException
  status: 502
# Error with rendered response, need to modify function to render API-Gateway Friendly response.


# Now call it again
aws apigateway test-invoke-method --rest-api-id ooyqsymb9g --resource-id 8zmr8h --http-method "GET" --path-with-query-string "/menus?menuid=888-123454249"

- body: '{"menuid":"888-123454249","restaurant":"Roma","cuisine":"Italian"}'
  headers:
    X-Amzn-Trace-Id: Root=1-616607b3-b53f9cf80107832ac68ee679;Sampled=0
  latency: 407
  log: "Execution log for request 5f1b41cd-6b45-4c94-8bac-8e9f118420cb\nTue Oct 12
    22:09:55 UTC 2021 : Starting execution for request: 5f1b41cd-6b45-4c94-8bac-8e9f118420cb\nTue
    Oct 12 22:09:55 UTC 2021 : HTTP Method: GET, Resource Path: /menus\nTue Oct 12
    22:09:55 UTC 2021 : Method request path: {}\nTue Oct 12 22:09:55 UTC 2021 : Method
    request query string: {menuid=888-123454249}\nTue Oct 12 22:09:55 UTC 2021 : Method
    request headers: {}\nTue Oct 12 22:09:55 UTC 2021 : Method request body before
    transformations: \nTue Oct 12 22:09:55 UTC 2021 : Endpoint request URI: https://lambda.us-east-1.amazonaws.com/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations\nTue
    Oct 12 22:09:55 UTC 2021 : Endpoint request headers: {X-Amz-Date=20211012T220955Z,
    x-amzn-apigateway-api-id=ooyqsymb9g, Accept=application/json, User-Agent=AmazonAPIGateway_ooyqsymb9g,
    Host=lambda.us-east-1.amazonaws.com, X-Amz-Content-Sha256=9b64c6c11f115e65b768521d207a3ce9c2bac029027ee58fcfae2bdc091d1fc3,
    X-Amzn-Trace-Id=Root=1-616607b3-b53f9cf80107832ac68ee679, x-amzn-lambda-integration-tag=5f1b41cd-6b45-4c94-8bac-8e9f118420cb,
    Authorization=*********************************************************************************************************************************************************************************************************************************************************************************************************************************************faf230,
    X-Amz-Source-Arn=arn:aws:execute-api:us-east-1:562697480070:ooyqsymb9g/test-invoke-stage/GET/menus,
    X-Amz-Security-Token=IQoJb3JpZ2luX2VjEJb//////////wEaCXVzLWVhc3QtMSJFMEMCIH6fzkjGy2aIUEPYpvZKo8SXGRGn/AtXGutWpJ7UHAljAh9Rbcn3SqOad6Qb4/DMxbBLpsdX7juiJTvK2H15d4GD
    [TRUNCATED]\nTue Oct 12 22:09:55 UTC 2021 : Endpoint request body after transformations:
    {\"resource\":\"/menus\",\"path\":\"/menus\",\"httpMethod\":\"GET\",\"headers\":null,\"multiValueHeaders\":null,\"queryStringParameters\":{\"menuid\":\"888-123454249\"},\"multiValueQueryStringParameters\":{\"menuid\":[\"888-123454249\"]},\"pathParameters\":null,\"stageVariables\":null,\"requestContext\":{\"resourceId\":\"8zmr8h\",\"resourcePath\":\"/menus\",\"httpMethod\":\"GET\",\"extendedRequestId\":\"HHYkJGADIAMFjSQ=\",\"requestTime\":\"12/Oct/2021:22:09:55
    +0000\",\"path\":\"/menus\",\"accountId\":\"562697480070\",\"protocol\":\"HTTP/1.1\",\"stage\":\"test-invoke-stage\",\"domainPrefix\":\"testPrefix\",\"requestTimeEpoch\":1634076595941,\"requestId\":\"5f1b41cd-6b45-4c94-8bac-8e9f118420cb\",\"identity\":{\"cognitoIdentityPoolId\":null,\"cognitoIdentityId\":null,\"apiKey\":\"test-invoke-api-key\",\"principalOrgId\":null,\"cognitoAuthenticationType\":null,\"userArn\":\"arn:aws:iam::562697480070:user/dev_system_account\",\"apiKeyId\":\"test-invoke-api-key-id\",\"userAgent\":\"aws-cli/2.2.43
    Python/3.9.7 Darwin/20.6.0 source/x86_64 prompt/off command/apigateway.test-i
    [TRUNCATED]\nTue Oct 12 22:09:55 UTC 2021 : Sending request to https://lambda.us-east-1.amazonaws.com/2015-03-31/functions/arn:aws:lambda:us-east-1:562697480070:function:menus/invocations\nTue
    Oct 12 22:09:56 UTC 2021 : Received response. Status: 200, Integration latency:
    403 ms\nTue Oct 12 22:09:56 UTC 2021 : Endpoint response headers: {Date=Tue, 12
    Oct 2021 22:09:56 GMT, Content-Type=application/json, Content-Length=147, Connection=keep-alive,
    x-amzn-RequestId=71bd26f5-2066-43c9-a3c8-566bb9c114b0, x-amzn-Remapped-Content-Length=0,
    X-Amz-Executed-Version=$LATEST, X-Amzn-Trace-Id=root=1-616607b3-b53f9cf80107832ac68ee679;sampled=0}\nTue
    Oct 12 22:09:56 UTC 2021 : Endpoint response body before transformations: {\"statusCode\":200,\"headers\":null,\"multiValueHeaders\":null,\"body\":\"{\\\"menuid\\\":\\\"888-123454249\\\",\\\"restaurant\\\":\\\"Roma\\\",\\\"cusisine\\\":\\\"Italian\\\"}\"}\nTue
    Oct 12 22:09:56 UTC 2021 : Method response body after transformations: {\"menuid\":\"888-123454249\",\"restaurant\":\"Roma\",\"cusisine\":\"Italian\"}\nTue
    Oct 12 22:09:56 UTC 2021 : Method response headers: {X-Amzn-Trace-Id=Root=1-616607b3-b53f9cf80107832ac68ee679;Sampled=0}\nTue
    Oct 12 22:09:56 UTC 2021 : Successfully completed execution\nTue Oct 12 22:09:56
    UTC 2021 : Method completed with status: 200\n"
  multiValueHeaders:
    X-Amzn-Trace-Id:
    - Root=1-616607b3-b53f9cf80107832ac68ee679;Sampled=0
  status: 200

# Check cloudwatch logs
aws logs filter-log-events --log-group-name /aws/lambda/menus --filter-pattern "ERROR"

# Deploy that bitch
# First to staging
aws apigateway create-deployment --rest-api-id ooyqsymb9g --stage-name staging
- createdDate: '2021-10-12T18:17:27-04:00'
  id: an7amt

# THIS FAILS
curl https://ooyqsymb9g.execute-api.us-east-1.amazonaws.com/staging/menus?menuid=978-1420931693

{"message":"Missing Authentication Token"}

# Install tool to curl with signed requests: https://github.com/okigan/awscurl
brew install awscurl

awscurl    --service execute-api -v https://ooyqsymb9g.execute-api.us-east-1.amazonaws.com/staging/menus?menuid=888-123454249


# Now test Post
 awscurl --service execute-api -X POST -d '{"menuid":"978-0141439587", "restaurant":"New Jamacia", "cuisine": "Carribean"}' https://ooyqsymb9g.execute-api.us-east-1.amazonaws.com/staging/menus

 TODO:
 Export API_Access_Group
 Export  dev_system_account
 Export  lambda-exec-role