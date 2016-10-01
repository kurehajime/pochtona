#  make assets.go by _assets/*
assets.go:_assets/*
	echo "//assets.go" > $@
	echo "package main" >> $@
	echo "var assets = map[string]string{}" >> $@
	echo "func init(){" >> $@
	for file in _assets/* ; do \
	    echo "assets[\"$$file\"]= \`">> $@ ; \
	    cat $$file >> $@ ; \
	    echo "\`">> $@ ; \
	done
	echo "}" >> $@
	gofmt -w $@
	
out:
	go test
	rm -rf ./bin
	gox -output "bin/{{.OS}}_{{.Arch}}/{{.Dir}}" ./...
	-find ./bin \! -name "*.zip" -type d -exec zip -r -j {}.zip {} \; -exec rm -R -d {} \;