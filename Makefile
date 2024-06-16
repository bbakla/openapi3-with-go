oapi-codegen-model:
	oapi-codegen -package=api --config=model-cfg.yaml  todo-app.yaml

oapi-codegen-server:
	oapi-codegen -package=api --config=server-cfg.yaml  todo-app.yaml

openapi-generator:
	openapi-generator generate --git-host github.com --git-repo-id  openoapi-code-generator --git-user-id bbakla -g go-gin-server -c openapi-generator-cfg.yaml -o openapi-generator/go-gen-gin -i todo-app.yaml

.PHONY: oapi-codegen-model oapi-codegen-server openapi-generator
