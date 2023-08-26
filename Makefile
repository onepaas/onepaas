generate-spec:
	swagger generate spec -m \
	| sed 's/v1-/v1./g' \
	| jq '(.definitions[].properties | select(has("metadata")) | .metadata) |= setpath([]; {"readOnly": true,"description": "Metadata Standard object.","allOf": [{"$$ref": "#/definitions/v1.Metadata"}]})' > api/swagger.json

run:
	go run ./cmd/onepaas serve

migrate:
	go run ./cmd/onepaas migration up
