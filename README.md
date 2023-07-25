# Teamgram - Unofficial open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

## Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server implementation written in golang, support private deployment.


### Source code deployment
#### Install [Go environment](https://go.dev/doc/install). Make sure Go version is at least 1.17.


### Docker deployment
#### Install [Docker](https://docs.docker.com/get-docker/)

#### Install [Docker Compose](https://docs.docker.com/compose/install/)

#### Get source code

```
git clone https://github.com/geneva-lake/teamgram-server.git
cd teamgram-server
```

#### Run

At first you need bring into the line versions of components. If you use v0.158.0 version
of [teamgram-proto](https://github.com/teamgram/proto) you should checkout to v.0.90.5 tag
of teamgram server. Teamgram-proto has a huge up to 10 megabyte files which makes build
difficult. Building could consume all memory of your computer and stuck. Setting GOMEMLIMIT 
to restring memory consumption could help. This environment variable was introduced in
go 1.20 so we need to switch to this version of golang. Also you could set smallframes
flag in the compilator arguments.

```  
# run dependency
docker-compose -f ./docker-compose-env.yaml up -d

# run docker-compose
docker-compose up -d
```

please PM the **[author](https://t.me/benqi)**

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
