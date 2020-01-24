GOOS=linux go build 
mv ./nature-remo-create-graph main
zip function.zip ./main
aws lambda update-function-code --function-name nature-remo-create-graph --zip-file fileb://function.zip
