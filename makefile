ASSETS =$(ls ./_assets)
assets.go:_assets/*
	echo "//assets.go" > assets.go
	echo "package main" >> assets.go
	echo "var assets = map[string]string{}" >> assets.go
	echo "func init(){" >> assets.go
	for file in $(ASSETS) ; do \
	    echo "assets[\"$$file\"]= \`">> assets.go ; \
	    cat ./_assets/$$file >> assets.go ; \
	    echo "\`">> assets.go ; \
	done
	echo "}" >> assets.go
	gofmt -w assets.go