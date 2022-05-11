# Reference: https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324

FROM golang:1.18-alpine3.15
WORKDIR /app
COPY ./main.go /app/main.go
RUN go build main.go

# Reference: https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "10001" \    
    "appuser"

FROM scratch
WORKDIR /app
COPY --from=0 /app/main /app/main

COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /etc/group /etc/group

# Use an underprivileged user.
USER appuser:appuser

ENTRYPOINT ["./main"]