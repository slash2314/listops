VERSION := $$(cat VERSION)
all: build strip
build: listops_linux_amd64 listops_linux_arm64 listops_linux_arm listops_linux_arm64 listops_mac_amd64 listops_mac_arm64 listops_mac_amd64 listops_windows_amd64.exe listops_windows_arm64.exe
templ_gen:
	templ generate
listops_linux_amd64: templ_gen
	CGO_ENABLED=0 go build  -o build/ssltool_linux_amd64_$(VERSION) cmd/listops/*.go
listops_windows_amd64.exe: templ_gen
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o build/listops_win_amd64_$(VERSION).exe cmd/listops/*.go
listops_windows_arm64.exe: templ_gen
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/listops_win_arm64_$(VERSION).exe cmd/listops/*.go
listops_mac_amd64: templ_gen
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/listops_mac_amd64_$(VERSION) cmd/listops/*.go
listops_mac_arm64: templ_gen
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/listops_mac_arm64_$(VERSION) cmd/listops/*.go
listops_linux_arm: templ_gen
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o build/listops_linux_arm_$(VERSION) cmd/listops/*.go
listops_linux_arm64: templ_gen
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/listops_linux_arm64_$(VERSION) cmd/listops/*.go

strip:
	find build -name "build/*" -exec strip {} \;

zip:
	cd build
	zip -r build/ssltool.zip build/*

prod: build strip zip

clean:
	rm build/*
	rm internal/templates/main_templ.go
