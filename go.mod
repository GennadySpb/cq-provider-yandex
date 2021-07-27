module github.com/yandex-cloud/cq-provider-yandex

go 1.16

require (
	github.com/GennadySpb/cq-provider-yandex v0.0.6
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/bxcodec/faker v2.0.1+incompatible // indirect
	github.com/cloudquery/cq-provider-sdk v0.2.8
	github.com/cloudquery/faker/v3 v3.7.4 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/go-hclog v0.16.1
	github.com/iancoleman/strcase v0.1.3
	github.com/jhump/protoreflect v1.9.0
	github.com/jinzhu/inflection v1.0.0
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/thoas/go-funk v0.8.1-0.20210502090430-efae847b30ab
	github.com/yandex-cloud/go-genproto v0.0.0-20210517152439-84c9ad4d8b5f
	github.com/yandex-cloud/go-sdk v0.0.0-20210517154707-ca282b96279e
	google.golang.org/grpc v1.39.0
)

replace github.com/cloudquery/faker/v3 v3.7.4 => github.com/daniil-ushkov/faker v1.5.1-0.20210727155430-974b577181cb
