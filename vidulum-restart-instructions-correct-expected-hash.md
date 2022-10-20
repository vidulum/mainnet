0.) check your expected hash:

`journalctl -u vidulumd -n 2500 | grep 'wrong Block.Header.AppHash'`

wait for the output to complete, this might take up to 1-2min

if the last lines of your output show:
```
prevote step: ProposalBlock is invalid err="wrong Block.Header.AppHash.  Expected 40AABEE68733114421477087D4C1389BC1737DA38D83989B5067FF3A88261651,  got E6CEB6E49EA23ACB9EFF24D5EE86C64070F67C110D77719B305DC776BC5B8E78" height=5212052...
```

you need to continue with below instructions.

instructions to reset your node:

1.) stop service
```
sudo systemctl stop vidulumd
```

2.) build new binary:
```
cd 
mkdir vidulum && git clone https://github.com/vidulum/mainnet vidulum
cd vidulum
git checkout v1.2.0
make install
```

remember to copy the freshly built binary from $HOME/go/bin/vidulumd to the right location (i.e. /usr/bin, /usr/local/bin or for cosmovisor users: $HOME/.vidulum/cosmovisor/genesis/bin)
```
cp $HOME/go/bin/vidulumd $HOME/.vidulum/cosmovisor/genesis/bin
```

3.) **backup your priv_validator_state.json**
```
mkdir $HOME/validator_state_backup
cp $HOME/.vidulum/data/priv_validator_state.json $HOME/validator_state_backup
```

3.) nuke your data dir and download official snapshot
```
cd $HOME/.vidulum
rm -r data
URL=https://quicksync.ccvalidators.com/SNAPSHOTS/vidulum-1_5212051.tar.lz4 && wget -O - $URL | lz4 -d | tar -xvf -
```

4.) **restore private_validator_state.json**
```
cp $HOME/validator_state_backup/priv_validator_state.json $HOME/.vidulum/data
```

5.) start service and cross fingers (your binary wouldn't start without restoring the validator state - it's **not** included in the snapshot)
```
sudo systemctl start vidulumd && journalctl -ocat -fu vidulumd
```
