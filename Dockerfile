FROM golang:alpine
LABEL authors="constanna"

RUN apk add --update ffmpeg
RUN wget -qO /usr/local/bin/yt-dlp https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp

WORKDIR /usr/src/funnygomusic

COPY go.mod go.sum main.go ./

COPY ./commands ./commands
COPY ./databaser ./databaser
COPY ./events ./events
COPY ./utils ./utils
COPY ./routes ./routes
COPY ./bot ./bot

RUN go build -v -o /usr/local/bin/funnygomusic

CMD ["funnygomusic"]