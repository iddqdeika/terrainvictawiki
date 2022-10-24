FROM golang:1.18-alpine3.15 AS builder

ENV APPDIR $GOPATH/src/remotetest

# путь к исполняемому файлу
ENV MAINPATH ./main.go

# название бинарника
ENV BINNAME myapp

# папка с аутпутом сборки
ENV OUTPUTDIR /out

# путь скопилированного бинарника
ENV BINPATH $OUTPUTDIR/$BINNAME

# создадим папку проекта
RUN mkdir -p ${APPDIR}

# назначим её рабочей
WORKDIR ${APPDIR}

# копирнем проект в эту папку для сборки
COPY . .

# добавим фронт в аутпут
COPY res ${OUTPUTDIR}/res

# добавим сертификат в аутпут
#COPY cert ${OUTPUTDIR}/cert

# посмотрим шо там выходит в папке
RUN echo In project dir:
RUN ls -la

# сбилдим бинарник
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINPATH} ${MAINPATH}

# посмотрим что в аутпуте
RUN echo In output dir:
RUN ls -la ${OUTPUTDIR}

# теперь базовый образ для исполнения
FROM alpine:3.15

# путь с аутпутом в билдере
ENV OUTPUTDIR /out

# целевой путь для прилоги (должен входить в PATH)
ENV USRBINDIR /usr/local/bin

# таймзона
RUN apk --update --no-cache add tzdata git curl
RUN cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime && \
    echo "Europe/Moscow" > /etc/timezone && \
    date

# перенесем аутпут билдера к нам
COPY --from=builder ${OUTPUTDIR} ${USRBINDIR}

# ставим рабочую папку
WORKDIR ${USRBINDIR}

# смотрим что в ней
RUN ls -la

ARG my_arg
ENV my_env_var=$my_arg

ENTRYPOINT echo $my_env_var

# ставим энтрипойнт
ENTRYPOINT myapp $my_env_var