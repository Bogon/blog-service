module blog-service

go 1.18

replace (
	api => ./internal/routers/api
	app => ./pkg/app
	convert => ./pkg/convert
	dao => ./internal/dao
	docs => ./docs
	email => ./pkg/email
	errcode => ./pkg/errcode
	global => ./global
	limiter => ./pkg/limiter
	logger => ./pkg/logger
	middleware => ./internal/middleware
	model => ./internal/model
	routers => ./internal/routers
	service => ./internal/service
	setting => ./pkg/setting
	tracer => ./pkg/tracer
	upload => ./pkg/upload
	util => ./pkg/util
	v1 => ./internal/routers/api/v1
)

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/natefinch/lumberjack v2.0.0+incompatible
	global v0.0.0
	logger v0.0.0-00010101000000-000000000000
	model v0.0.0-00010101000000-000000000000
	routers v0.0.0-00010101000000-000000000000
	setting v0.0.0
	tracer v0.0.0-00010101000000-000000000000
)

require (
	api v0.0.0-00010101000000-000000000000 // indirect
	app v0.0.0 // indirect
	convert v0.0.0-00010101000000-000000000000 // indirect
	dao v0.0.0-00010101000000-000000000000 // indirect
	docs v0.0.0-00010101000000-000000000000 // indirect
	email v0.0.0-00010101000000-000000000000 // indirect
	errcode v0.0.0-00010101000000-000000000000 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/eddycjy/opentracing-gorm v0.0.0-20200209122056-516a807d2182 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/pprof v1.4.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/juju/ratelimit v1.0.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.13.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a // indirect
	github.com/swaggo/gin-swagger v1.5.3 // indirect
	github.com/swaggo/swag v1.8.6 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	limiter v0.0.0-00010101000000-000000000000 // indirect
	middleware v0.0.0-00010101000000-000000000000 // indirect
	service v0.0.0-00010101000000-000000000000 // indirect
	upload v0.0.0-00010101000000-000000000000 // indirect
	util v0.0.0-00010101000000-000000000000 // indirect
	v1 v0.0.0 // indirect
)
