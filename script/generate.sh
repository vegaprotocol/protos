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
        sed -i -re 's/this\.Size_/this.Size/' "$pbfile" \
		&& ./script/fix_imports.sh "$pbfile"
	done
	find data-node -name '*.validator.pb.go' | sort | while read -r pbfile
	do
        sed -i -re 's/this\.Size_/this.Size/' "$pbfile" \
		&& ./script/fix_imports.sh "$pbfile"
	done
}

function gen_swagger() {
	buf generate --path=./sources/vega/api --template=./sources/vega/api/v1/buf.gen.yaml # generate swagger
	buf generate --path=./sources/data-node/api/v1 --template=./sources/data-node/api/v1/buf.gen.yaml # generate swagger
}
check
gen_code
gen_swagger
