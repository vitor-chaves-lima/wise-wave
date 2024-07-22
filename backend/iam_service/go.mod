module wisewave.tech/iam_service

go 1.22.0

require (
	github.com/aws/aws-lambda-go v1.47.0
	wisewave.tech/common v0.0.0
)

require (
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)

replace wisewave.tech/common => ../common
