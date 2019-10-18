FROM registry.svc.ci.openshift.org/openshift/release:golang-1.11 AS builder
WORKDIR /go/src/github.com/mhrivnak/memweb-reconciler
COPY . .
RUN go build cmd/memweb.go


FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/memweb-reconciler \
    USER_UID=1001 \
    USER_NAME=memweb-reconciler

# install operator binary
COPY --from-builder /go/src/github.com/mhrivnak/memweb-reconciler/memweb ${OPERATOR}

RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/memweb-reconciler"]

USER ${USER_UID}
