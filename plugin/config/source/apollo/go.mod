module github.com/stack-labs/stack/plugin/config/source/apollo

go 1.14

replace (
	github.com/stack-labs/stack v1.0.1-rc1 => ../../../../
)

require (
	github.com/magiconair/properties v1.8.4 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stack-labs/stack v1.0.1-rc1
	github.com/tevid/gohamcrest v1.1.1
)
