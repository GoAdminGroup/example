all:
	rm -rf vendor/*
	govendor init
	govendor add +external
	zip -r -q vendor.zip vendor -x "*.DS_Store" -x "__MACOSX"