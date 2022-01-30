# vidulum

**vidulum** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

Github: https://github.com/vidulum/mainnet

Explorer: https://explorers.vidulum.app/vidulum

Discord for Support: https://discord.gg/QhV2Wx6

Vidulum App to interact with our blockchains staking, governance and more!

Android: https://play.google.com/store/apps/details?id=com.vidulumwallet.app

iOS: https://apps.apple.com/us/app/id1505859171

Web: https://wallet.vidulum.app


#### - Download Prebuilt Binary
`wget https://github.com/vidulum/mainnet/releases/download/v1.0.0/vidulum_linux_amd64.tar.gz`


#### - Store Binary in local bin
`mv vidulumd /usr/local/bin/vidulumd`


#### - Generate address and build chain data files (Default  .vidulum)
#### - Replace CAPS with your preferred name for the address/key (Example: mainwallet)
`vidulumd keys add KEYNAME --keyring-backend os`


#### - Save the genesis file to Vidulum data folder
`wget -P .vidulum/config/ https://raw.githubusercontent.com/vidulum/mainnet/main/genesis.json`


#### - Start the daemon and save peers to address book
`vidulumd start --p2p.seeds="883ec7d5af7222c206674c20c997ccc5c242b38b@ec2-3-82-120-39.compute-1.amazonaws.com:26656,eed11fff15b1eca8016c6a0194d86e4a60a65f9b@apollo.erialos.me:26656"`


#### - Allow daemon to fully sync



## Release

To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of your blockchain node's binary, execute the following command on your machine:

TODO

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
