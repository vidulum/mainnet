# vidulum

**vidulum** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

Github: https://github.com/vidulum/mainnet
Explorer: https://explorers.vidulum.app/vidulum
Discord for Support: https://discord.gg/QhV2Wx6

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
`vidulumd start --p2p.persistent_peers="209688f5bccb88f6397a97cc11ab545a014aa559@137.184.92.115:26656,d45e9dd8878d7c22d59ded3557f61da37420a4c6@95.217.118.211:26656,cae7d9d21c1752300277eab72d861b0c6638b2e3@164.68.119.151:26656,7a44ea6ecb59b0e4bd01b58a75163ec64b164bb4@63.210.148.24:26656,3bf3d98dfd4000dd5ff8189882a9f96848b99b87@137.220.60.196:26656,057fa262fe2030cc6e9095dc52d15b79ffcb923d@142.115.20.25:26656"`

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
