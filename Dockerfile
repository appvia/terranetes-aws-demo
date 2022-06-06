FROM golang:1.18

WORKDIR /app

# Create a new user to run the application as.
RUN groupadd --gid 1000 gouser \
    && useradd --uid 1000 --gid 1000 -m gouser

COPY . .
RUN go build -v -o /app/aws-s3

# Give that new user permissions to the app and folders we need.
RUN chown -R gouser:gouser /app \
    && chmod 770 /app \
    && chmod 770 /app/aws-s3

# Instructs Docker to use the user we created to run any further commands.
USER 1000

CMD ["/app/aws-s3"]