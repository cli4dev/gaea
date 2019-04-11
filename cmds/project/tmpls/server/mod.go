package server

const TPLMod = `
module {{.projectName|lName}}

go 1.12

replace (
	github.com/zkfy/archiver => github.com/mholt/archiver v2.1.0+incompatible
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20180910181607-0e37d006457b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190301231843-5614ed5bae6f
	golang.org/x/net => github.com/golang/net v0.0.0-20180911220305-26e67e76b6c3
	golang.org/x/sync => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180909124046-d0be0721c37e
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20180820150726-fbb02b2291d28
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9
	google.golang.org/grpc => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/asaskevich/govalidator v0.0.0-20180720115003-f9ffefc3facf // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190329173943-551aad21a668 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0 // indirect
	github.com/gmallard/stompngo v1.0.11 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-redis/redis v6.15.2+incompatible // indirect
	github.com/golang/protobuf v1.3.1 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/micro-plat/hydra v0.0.0-20190315085906-11529ac2c330
	github.com/micro-plat/lib4go v0.0.0-20190313063825-3419c90d56a9 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nwaples/rardecode v1.0.0 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/profile v1.3.0 // indirect
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible // indirect
	github.com/stretchr/testify v1.3.0 // indirect
	github.com/ugorji/go/codec v0.0.0-20190320090025-2dc34c0b8780 // indirect
	github.com/urfave/cli v1.20.0 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	github.com/yosssi/gmq v0.0.1 // indirect
	github.com/zkfy/archiver v0.0.0-00010101000000-000000000000 // indirect
	//github.com/zkfy/archiver v1.1.2 // indirect
	github.com/zkfy/cron v0.0.0-20170309132418-df38d32658d8 // indirect
	github.com/zkfy/go-metrics v0.0.0-20161128210544-1f30fe9094a5 // indirect
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0 // indirect
	github.com/zkfy/jwt-go v3.0.0+incompatible // indirect
	github.com/zkfy/log v0.0.0-20180312054228-b2704c3ef896 // indirect
	github.com/zkfy/stompngo v0.0.0-20170803022748-9378e70ca481 // indirect
	google.golang.org/grpc v0.0.0-00010101000000-000000000000 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
`
