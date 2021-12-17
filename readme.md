# deweb
**deweb** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Initialize chain
```
starport scaffold chain gitlab.com/deweb-services/deweb-chain
```

Create messages
```
starport scaffold message save_user message chain --module deweb-chain
starport scaffold message delete_key uuid --module deweb-chain
```

Create queries
```
starport scaffold query get_key_record uuid --response message --module deweb-chain
starport scaffold query get_user_key_records address --response uuids --module deweb-chain
starport scaffold query filter_user_key_records address chain deleted --response uuids --module deweb-chain
```

## Test cases from console

**Initial configuration for Cosmos SDK located in folder: cosmos_sdk_config in current repository**

You must scaffold chain and then replace files in config and data home folder (by default located in ~/.deweb-chain).
And then replace generate files of blockchain with files in this repo.
Or you can try to clone this repo, build binary and then set config hole folder via flag --home

First build binary:
```
starport chain build
```
Initialize blockchain
```
starport chain init
```

Test commands
```
Created account "alice" with address "cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t"
Created account "bob" with address "cosmos16vtqmpfd3w08awp2t608s9475747rn4ljk6ja9"

~/go/bin/dewebd tx deweb save-user supermsg cosmos --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-user superuser2 sia --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-user superbob sia --from bob --chain-id deweb

~/go/bin/dewebd q deweb get-user-key-records cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t
uuids:
- 65e0dcea-5cde-11ec-b160-acde48001122
- 81724728-5cde-11ec-b160-acde48001122

~/go/bin/dewebd q deweb get-user-key-records cosmos16vtqmpfd3w08awp2t608s9475747rn4ljk6ja9
uuids:
- 8ab65676-5cde-11ec-b160-acde48001122

~/go/bin/dewebd q deweb filter-user-key-records cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t "" true
~/go/bin/dewebd q deweb filter-user-key-records cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t cosmos true

~/go/bin/dewebd tx deweb delete-key 65e0dcea-5cde-11ec-b160-acde48001122 --from alice --chain-id deweb

~/go/bin/dewebd q deweb filter-user-key-records cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t "" true
~/go/bin/dewebd q deweb filter-user-key-records cosmos1naleh8tvj6ks63270mqdt7wqtee877plrvmk3t "" false
```


## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/deweb-services/deweb@latest! | sudo bash
```
`deweb-services/deweb` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
