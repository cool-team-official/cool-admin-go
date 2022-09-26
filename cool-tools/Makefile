pack: pack.template-simple

pack.template-simple:
	@rm -fr temp
	@mkdir temp || exit 0
	@cd temp && git clone -b simple  https://github.com/cool-team-official/cool-admin-go.git cool-admin-go-simple
	@rm -fr temp/cool-admin-go-simple/.git
	@cd temp && gf pack cool-admin-go-simple ../internal/packed/cool-admin-go-simple.go -n=packed -y
	@rm -fr temp

# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://gh.hjmcloud.cn/github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install  && \
	rm ./gf

pack.ssh:
	@rm -fr temp
	@mkdir temp || exit 0
	@cd temp && git clone -b simple  git@github.com:cool-team-official/cool-admin-go.git cool-admin-go-simple
	@rm -fr temp/cool-admin-go-simple/.git
	@cd temp && gf pack cool-admin-go-simple ../internal/packed/cool-admin-go-simple.go -n=packed -y
	@rm -fr temp

# Install/Update to the latest cool-tools.
.PHONY: tools
tools:
	@set -e; \
	curl -L https://download.fastgit.org/cool-team-official/cool-admin-go/releases/latest/download/cool-tools_$(shell go env GOOS)_$(shell go env GOARCH) -o ./cool-tools && \
	chmod +x cool-tools && \
	./cool-tools install && \
	rm ./cool-tools


# Check and install cool-tools.
.PHONY: tools.install
tools.install:
	@set -e; \
	echo "Checking cool-tools..."; \
	cool-tools -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "cool-tools is not installed, start proceeding auto installation..."; \
		make tools; \
	fi;\
	echo "cool-tools is installed.";