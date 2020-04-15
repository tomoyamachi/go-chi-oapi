
gen-server:
	rm -fr pkg/gen
	mkdir -p pkg/gen/store pkg/gen/user
	@oapi-codegen -generate types,chi-server -package store -include-tags store openapi.json > pkg/gen/store/api.go
	@oapi-codegen -generate types,chi-server -package user -include-tags user openapi.json > pkg/gen/user/api.go
