module api

go 1.18

require (
	app v0.0.0-00010101000000-000000000000
	convert v0.0.0-00010101000000-000000000000
	errcode v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.8.1
	global v0.0.0-00010101000000-000000000000
	service v0.0.0-00010101000000-000000000000
	upload v0.0.0-00010101000000-000000000000
)

require (
	dao v0.0.0-00010101000000-000000000000 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.13.0 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	logger v0.0.0-00010101000000-000000000000 // indirect
	model v0.0.0-00010101000000-000000000000 // indirect
	setting v0.0.0 // indirect
	util v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	app => ../../../pkg/app
	convert => ../../../pkg/convert
	dao => ../../../internal/dao
	errcode => ../../../pkg/errcode
	global => ../../../global
	logger => ../../../pkg/logger
	model => ../../../internal/model
	service => ../../../internal/service
	setting => ../../../pkg/setting
	upload => ../../../pkg/upload
	util => ../../../pkg/util

)
