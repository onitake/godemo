FROM golang:1.13 AS build

WORKDIR /build
COPY . /build
ENV CGO_ENABLED=0
RUN go build -o godemo *.go

FROM scratch AS runtime

LABEL org.label-schema.name "godemo"
LABEL org.label-schema.description "Just a simple Go container demo"

WORKDIR /
COPY --from=build /build/godemo /godemo

EXPOSE 8080
CMD ["/godemo"]
