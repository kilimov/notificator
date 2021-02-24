#
# Контейнер сборки
#
FROM golang:1.14 as builder

ARG DRONE
ARG DRONE_TAG
ARG DRONE_COMMIT
ARG DRONE_BRANCH

ENV CGO_ENABLED=0

COPY . /go/src/github.com/kilimov/notificator
WORKDIR /go/src/github.com/kilimov/notificator
RUN \
    if [ -z "$DRONE" ] ; then echo "no drone" && version=`git describe --abbrev=6 --always --tag`; \
    else version=${DRONE_TAG}${DRONE_BRANCH}-`echo ${DRONE_COMMIT} | cut -c 1-7` ; fi && \
    echo "version=$version" && \
    cd cmd/notificator && \
    go build -a -tags notificator -installsuffix notificator -ldflags "-X apiserver.version=${version} -s -w" -o /go/bin/notificator

#
# Контейнер для получения актуальных SSL/TLS сертификатов
#
FROM alpine as alpine
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S notificator && adduser -S notificator -G notificator

ENTRYPOINT [ "/bin/notificator" ]

#
# Контейнер рантайма
#
FROM scratch
COPY --from=builder /go/bin/notificator /bin/notificator

# копируем сертификаты из alpine
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs

# копируем документацию
COPY --from=alpine /usr/share/notificator /usr/share/notificator

# копируем пользователя и группу из alpine
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER notificator

ENTRYPOINT ["/bin/notificator"]



