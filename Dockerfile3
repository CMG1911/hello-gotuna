FROM golang:1.20-alpine AS construccion
WORKDIR /opt/gotuna/
COPY . .
EXPOSE 8888
RUN go build -o build/main ./examples/fullapp/cmd/main.go

FROM golang:1.20-alpine AS ejecucion
WORKDIR /opt/gotuna/
EXPOSE 8888
COPY --from=construccion /opt/gotuna/build ./build
CMD ["./build/main"] 
