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

	pwd

	# Make *.validator.pb.go files deterministic.
	find proto -name '*.validator.pb.go' | sort | while read -r pbfile
	do
        sed -i -re 's/this\.Size_/this.Size/' "$pbfile" \
		&& ./script/fix_imports.sh "$pbfile"
	done

	chmod 0644 vega/*.go vega/api/*.go
}

function gen_swagger() {
	buf generate --path=./vega/api --template=./vega/api/buf.gen.yaml # generate swagger
}
check
gen_code
gen_swagger
