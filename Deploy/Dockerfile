FROM go:1.5
WORKDIR /code
COPY . .
RUN go mod tidy
CMD ["go", "run", "main"]
