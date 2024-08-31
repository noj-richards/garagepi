protolint appconfig/appconfig.proto
protoc --go_out=. --go_opt=paths=source_relative appconfig/appconfig.proto
