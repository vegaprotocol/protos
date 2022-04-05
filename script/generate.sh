#!/bin/bash

TARGET="vega"

function check() {
	if [[ ! -d "$TARGET" ]]; then
		echo "Target directory \`$TARGET\` not found, run this script from vega protos's repository root path"
		exit 1
	fi
}

function gen_code() {
	# generate code, grpc and validators code
	buf generate

	# Since ./proto/github/{grpc-ecosystem,mwitkow} are dependencies,
	# buf will generate code for them to
	rm -rf ./github.com

	# Make *.validator.pb.go files deterministic.
	find vega -name '*.validator.pb.go' | sort | while read -r pbfile
	do
        sed -i "" -re 's/this\.Size_/this.Size/' "$pbfile" \
		&& ./script/fix_imports.sh "$pbfile"
	done
	find data-node -name '*.validator.pb.go' | sort | while read -r pbfile
	do
        sed -i "" -re 's/this\.Size_/this.Size/' "$pbfile" \
		&& ./script/fix_imports.sh "$pbfile"
	done

	find vega -name '*.go' | sort | while read -r pbfile
	do
       sed -i "" -e "s/\.String()/.NonDeterministicString()/g" "$pbfile"
       sed -i "" -e "s/ String()/ NonDeterministicString()/g" "$pbfile"
	done
}

function gen_swagger() {
	buf generate --path=./sources/vega/api --template=./sources/vega/api/v1/buf.gen.yaml # generate swagger
	buf generate --path=./sources/data-node/api/v1 --template=./sources/data-node/api/v1/buf.gen.yaml # generate swagger
}

function gen_json() {
	mkdir -p generated/json/vega
	mkdir -p ./generated/json/data-node/api/v1
	mkdir -p ./generated/json/data-node/api/v2

	protoc --jsonschema_out=./generated/json/vega --proto_path=./sources sources/vega/*.proto
	protoc --jsonschema_out=./generated/json/data-node/api/v1 --proto_path=./sources sources/data-node/api/v1/*.proto
	protoc --jsonschema_out=./generated/json/data-node/api/v2 --proto_path=./sources sources/data-node/api/v2/*.proto
}

function gen_docs() {
  mkdir -p generated

  protoc --doc_out=./generated --doc_opt=json,proto.json --proto_path=sources/ \
  sources/vega/*.proto \
  sources/vega/oracles/**/*.proto \
  sources/vega/commands/**/*.proto \
  sources/vega/events/**/*.proto \
  sources/vega/api/**/*.proto \
  sources/vega/checkpoint/**/*.proto \
  sources/vega/snapshot/**/*.proto \
  sources/vega/events/**/*.proto \
  sources/vega/wallet/**/*.proto
}

check
gen_code
gen_swagger
gen_json
gen_docs
