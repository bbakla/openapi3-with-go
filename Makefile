oapi-codegen-strict-server-model:
	oapi-codegen  --config=strict-server-model-cfg.yaml  user-app.yaml

oapi-codegen-strict-server:
	oapi-codegen --config=strict-server-cfg.yaml  user-app.yaml

oapi-codegen-non-strict-server-model:
	oapi-codegen --config=non-strict-server-model-cfg.yaml  user-app.yaml

oapi-codegen-non-strict-server:
	oapi-codegen --config=non-strict-server-cfg.yaml  user-app.yaml

openapi-generator:
	openapi-generator generate --git-host github.com --git-repo-id  openoapi-code-generator --git-user-id bbakla -g go-gin-server -c openapi-generator-cfg.yaml -o openapi-generator/generated -i user-app.yaml

.PHONY: oapi-codegen-model oapi-codegen-server openapi-generator
