# OneHub
Yet another OneDrive shared file explorer (inspired by [OneIndex](https://github.com/donwa/oneindex)).

The project is still under heavy development. Demo site: https://onehub.boar.ac.cn

![screenshot](https://github.com/hebingchang/onehub/raw/master/docs/screenshot.png)
## Deploy with Docker
(Recommended)
```shell script
docker run --detach \
    --name onehub \
    -v ~/onehub:/go/src/onehub/config \
    -p 80:8080 \
    hebingchang/onehub
```

Then head to http://server-ip/admin/#/init to configure the application.

**Notice that** currently the Microsoft Graph Application preconfigured in default `config.yaml` only supports the redirection back to `onehub.boar.ac.cn`. To support other domains, please create your own application in https://portal.azure.com/.
 
## Deploy
```shell script
git clone --recurse-submodules https://github.com/hebingchang/onehub
cd onehub
cd frontend/public && yarn && yarn build && cd ../../
cd frontend/admin && yarn && yarn build && cd ../../
go run main.go
```

## Build
```shell script
git clone --recurse-submodules https://github.com/hebingchang/onehub
cd onehub
cd frontend/public && yarn && yarn build && cd ../../
cd frontend/admin && yarn && yarn build && cd ../../
docker build . -t hebingchang/onehub
```