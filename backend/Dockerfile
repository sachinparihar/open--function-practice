# Use the buildah builder image
FROM openfunction/buildah:v1.23.1 AS buildah

# Set the working directory
WORKDIR /workspace

# Clone the source code
RUN git clone --single-branch --depth 1 https://github.com/sachinparihar/open--function-practice.git /workspace

# Build the Go application
RUN cd /workspace/backend && \
    CGO_ENABLED=0 go build \
    -tags netgo \
    -ldflags "-s -w -extldflags '-static'" \
    -o /workspace/helloworld

# Final image stage
FROM scratch
COPY --from=buildah /workspace/helloworld /helloworld
ENTRYPOINT [ "/helloworld" ]
EXPOSE 8080