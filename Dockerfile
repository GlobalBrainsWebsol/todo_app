FROM golang:1.21.4

# USER gbc
# RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs \
#   github.com/ramya-rao-a/go-outline \
#   github.com/nsf/gocode \
#   github.com/acroca/go-symbols \
#   github.com/fatih/gomodifytags \
#   github.com/josharian/impl \
#   github.com/haya14busa/goplay/cmd/goplay \
RUN go install -v golang.org/x/tools/gopls@latest \
  && go install -v github.com/go-delve/delve/cmd/dlv@latest \
  && go install -v honnef.co/go/tools/cmd/staticcheck@2022.1 \
  && go install -v github.com/cosmtrek/air@latest
RUN apt update && apt install sqlite3


WORKDIR /todo_app
# CMD ["air", "-c", ".air.toml"]
