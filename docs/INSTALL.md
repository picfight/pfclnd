# Table of Contents
* [Installation](#installation)
    * [Preliminaries](#preliminaries)
    * [Installing lnd](#installing-lnd)
* [Available Backend Operating Modes](#available-backend-operating-modes)
  * [pfcd Options](#pfcd-options)
  * [Neutrino Options](#neutrino-options)
  * [Picfightcoind Options](#bitcoind-options)
  * [Using pfcd](#using-pfcd)
    * [Installing pfcd](#installing-pfcd)
    * [Starting pfcd](#starting-pfcd)
    * [Running lnd using the pfcd backend](#running-lnd-using-the-pfcd-backend)
  * [Using Neutrino](#using-neutrino)
  * [Using picfightcoind or litecoind](#using-bitcoind-or-litecoind)
* [Macaroons](#macaroons)
* [Network Reachability](#network-reachability)
* [Simnet vs. Testnet Development](#simnet-vs-testnet-development)
* [Creating an lnd.conf (Optional)](#creating-an-lndconf-(optional))

# Installation

### Preliminaries
  In order to work with [`lnd`](https://github.com/picfight/pfclnd), the
  following build dependencies are required:

  * **Go:** `lnd` is written in Go. To install, run one of the following commands:


    **Note**: The minimum version of Go supported is Go 1.11. We recommend that
    users use the latest version of Go, which at the time of writing is
    [`1.11`](https://blog.golang.org/go1.11).


    On Linux:
    ```
    sudo apt-get install golang-1.11-go
    ```
    > Note that golang-1.11-go puts binaries in /usr/lib/go-1.11/bin. If you want them on your PATH, you need to make that change yourself. Alternatively, you can run:
    ```
    sudo ln -s /usr/lib/go-1.11/bin/go /usr/local/bin/go
    ```

    On Mac OS X:
    ```
    brew install go
    ```

    On FreeBSD:
    ```
    pkg install go
    ```

    Alternatively, one can download the pre-compiled binaries hosted on the
    [golang download page](https://golang.org/dl/). If one seeks to install
    from source, then more detailed installation instructions can be found
    [here](http://golang.org/doc/install).

    At this point, you should set your `$GOPATH` environment variable, which
    represents the path to your workspace. By default, `$GOPATH` is set to
    `~/go`. You will also need to add `$GOPATH/bin` to your `PATH`. This ensures
    that your shell will be able to detect the binaries you install.

    ```bash
    export GOPATH=~/gocode
    export PATH=$PATH:$GOPATH/bin
    ```

    We recommend placing the above in your .bashrc or in a setup script so that
    you can avoid typing this every time you open a new terminal window.

  * **go modules:** This project uses [go modules](https://github.com/golang/go/wiki/Modules) 
    to manage dependencies as well as to provide *reproducible builds*.

    Usage of go modules (with go 1.11) means that you no longer need to clone
    `lnd` into your `$GOPATH` for development purposes. Instead, your `lnd`
    repo can now live anywhere!

### Installing lnd

With the preliminary steps completed, to install `lnd`, `lncli`, and all
related dependencies run the following commands:
```
go get -d github.com/picfight/pfclnd
cd $GOPATH/src/github.com/picfight/pfclnd
make && make install
```

**NOTE**: Our instructions still use the `$GOPATH` directory from prior
versions of Go, but with go 1.11, it's now possible for `lnd` to live
_anywhere_ on your file system.

For Windows WSL users, make will need to be referenced directly via
/usr/bin/make/, or alternatively by wrapping quotation marks around make,
like so:

```
/usr/bin/make && /usr/bin/make install

"make" && "make" install
```

On FreeBSD, use gmake instead of make.

Alternatively, if one doesn't wish to use `make`, then the `go` commands can be
used directly:
```
GO111MODULE=on go install -v ./...
```

**Updating**

To update your version of `lnd` to the latest version run the following
commands:
```
cd $GOPATH/src/github.com/picfight/pfclnd
git pull
make clean && make && make install
```

On FreeBSD, use gmake instead of make.

Alternatively, if one doesn't wish to use `make`, then the `go` commands can be
used directly:
```
cd $GOPATH/src/github.com/picfight/pfclnd
git pull
GO111MODULE=on go install -v ./...
```

**Tests**

To check that `lnd` was installed properly run the following command:
```
make check
```

# Available Backend Operating Modes

In order to run, `lnd` requires, that the user specify a chain backend. At the
time of writing of this document, there are three available chain backends:
`pfcd`, `neutrino`, `bitcoind`. All but neutrino (atm) can run on mainnet with
an out of the box `lnd` instance. We don't require `--txindex` when running
with `bitcoind` or `pfcd` but activating the `txindex` will generally make
`lnd` run faster.

**NOTE: WE DO NOT FULLY SUPPORT PRUNED OPERATING MODES FOR FULL NODES.** It's
possible to run a node in a pruned mode and have it serve lnd, however one must
take care to ensure that `lnd` has all blocks on disk since the birth of the
wallet, and the age of the earliest channels (which were created around March
2018).

The set of arguments for each of the backend modes is as follows:

## pfcd Options
```
pfcd:
      --pfcd.dir=                                             The base directory that contains the node's data, logs, configuration file, etc. (default: /Users/roasbeef/Library/Application Support/Pfcd)
      --pfcd.rpchost=                                         The daemon's rpc listening address. If a port is omitted, then the default port for the selected chain parameters will be used. (default: localhost)
      --pfcd.rpcuser=                                         Username for RPC connections
      --pfcd.rpcpass=                                         Password for RPC connections
      --pfcd.rpccert=                                         File containing the daemon's certificate file (default: /Users/roasbeef/Library/Application Support/Pfcd/rpc.cert)
      --pfcd.rawrpccert=                                      The raw bytes of the daemon's PEM-encoded certificate chain which will be used to authenticate the RPC connection.
```

## Neutrino Options
```
neutrino:
  -a, --neutrino.addpeer=                                     Add a peer to connect with at startup
      --neutrino.connect=                                     Connect only to the specified peers at startup
      --neutrino.maxpeers=                                    Max number of inbound and outbound peers
      --neutrino.banduration=                                 How long to ban misbehaving peers.  Valid time units are {s, m, h}.  Minimum 1 second
      --neutrino.banthreshold=                                Maximum allowed ban score before disconnecting and banning misbehaving peers.
```

## Picfightcoind Options
```
bitcoind:
      --bitcoind.dir=                                         The base directory that contains the node's data, logs, configuration file, etc. (default: /Users/roasbeef/Library/Application Support/Picfightcoin)
      --bitcoind.rpchost=                                     The daemon's rpc listening address. If a port is omitted, then the default port for the selected chain parameters will be used. (default: localhost)
      --bitcoind.rpcuser=                                     Username for RPC connections
      --bitcoind.rpcpass=                                     Password for RPC connections
      --bitcoind.zmqpubrawblock=                              The address listening for ZMQ connections to deliver raw block notifications
      --bitcoind.zmqpubrawtx=                                 The address listening for ZMQ connections to deliver raw transaction notifications
```

## Using pfcd

### Installing pfcd

On FreeBSD, use gmake instead of make.

To install pfcd, run the following commands:

Install **pfcd**:
```
make pfcd
```

Alternatively, you can install [`pfcd` directly from its
repo](https://github.com/picfight/pfcd).

### Starting pfcd

Running the following command will create `rpc.cert` and default `pfcd.conf`.

```
pfcd --testnet --rpcuser=REPLACEME --rpcpass=REPLACEME
```
If you want to use `lnd` on testnet, `pfcd` needs to first fully sync the
testnet blockchain. Depending on your hardware, this may take up to a few
hours. Note that adding `--txindex` is optional, as it will take longer to sync
the node, but then `lnd` will generally operate faster as it can hit the index
directly, rather than scanning blocks or BIP 158 filters for relevant items.

(NOTE: It may take several minutes to find segwit-enabled peers.)

While `pfcd` is syncing you can check on its progress using pfcd's `getinfo`
RPC command:
```
btcctl --testnet --rpcuser=REPLACEME --rpcpass=REPLACEME getinfo
{
  "version": 120000,
  "protocolversion": 70002,
  "blocks": 1114996,
  "timeoffset": 0,
  "connections": 7,
  "proxy": "",
  "difficulty": 422570.58270815,
  "testnet": true,
  "relayfee": 0.00001,
  "errors": ""
}
```

Additionally, you can monitor pfcd's logs to track its syncing progress in real
time.

You can test your `pfcd` node's connectivity using the `getpeerinfo` command:
```
btcctl --testnet --rpcuser=REPLACEME --rpcpass=REPLACEME getpeerinfo | more
```

### Running lnd using the pfcd backend

If you are on testnet, run this command after `pfcd` has finished syncing.
Otherwise, replace `--bitcoin.testnet` with `--bitcoin.simnet`. If you are
installing `lnd` in preparation for the
[tutorial](http://dev.lightning.community/tutorial), you may skip this step.
```
lnd --bitcoin.active --bitcoin.testnet --debuglevel=debug --pfcd.rpcuser=kek --pfcd.rpcpass=kek --externalip=X.X.X.X
```

## Using Neutrino

In order to run `lnd` in its light client mode, you'll need to locate a
full-node which is capable of serving this new light client mode. `lnd` uses
[BIP 157](https://github.com/bitcoin/bips/tree/master/bip-0157) and [BIP
158](https://github.com/bitcoin/bips/tree/master/bip-0158) for its light client
mode.  A public instance of such a node can be found at
`faucet.lightning.community`.

To run lnd in neutrino mode, run `lnd` with the following arguments, (swapping
in `--bitcoin.simnet` if needed), and also your own `pfcd` node if available:
```
lnd --bitcoin.active --bitcoin.testnet --debuglevel=debug --bitcoin.node=neutrino --neutrino.connect=faucet.lightning.community
```


## Using picfightcoind or litecoind

The configuration for picfightcoind and litecoind are nearly identical, the
following steps can be mirrored with loss of generality to enable a litecoind
backend.  Setup will be described in regards to `bitcoind`, but note that `lnd`
uses a distinct `litecoin.node=litecoind` argument and analogous
subconfigurations prefixed by `litecoind`. Note that adding `--txindex` is
optional, as it will take longer to sync the node, but then `lnd` will
generally operate faster as it can hit the index directly, rather than scanning
blocks or BIP 158 filters for relevant items.

To configure your picfightcoind backend for use with lnd, first complete and verify
the following:

- Since `lnd` uses
  [ZeroMQ](https://github.com/bitcoin/bitcoin/blob/master/doc/zmq.md) to
  interface with `bitcoind`, *your `bitcoind` installation must be compiled with
  ZMQ*. Note that if you installed `bitcoind` from source and ZMQ was not present, 
  then ZMQ support will be disabled, and `lnd` will quit on a `connection refused` error. 
  If you installed `bitcoind` via Homebrew in the past ZMQ may not be included 
  ([this has now been fixed](https://github.com/Homebrew/homebrew-core/pull/23088) 
  in the latest Homebrew recipe for picfightcoin)
- Configure the `bitcoind` instance for ZMQ with `--zmqpubrawblock` and
  `--zmqpubrawtx`. These options must each use their own unique address in order
  to provide a reliable delivery of notifications (e.g.
  `--zmqpubrawblock=tcp://127.0.0.1:28332` and
  `--zmqpubrawtx=tcp://127.0.0.1:28333`).
- Start `bitcoind` running against testnet, and let it complete a full sync with
  the testnet chain (alternatively, use `--bitcoind.regtest` instead).

Here's a sample `bitcoin.conf` for use with lnd:
```
testnet=1
server=1
daemon=1
zmqpubrawblock=tcp://127.0.0.1:28332
zmqpubrawtx=tcp://127.0.0.1:28333
```

Once all of the above is complete, and you've confirmed `bitcoind` is fully
updated with the latest blocks on testnet, run the command below to launch
`lnd` with `bitcoind` as your backend (as with `bitcoind`, you can create an
`lnd.conf` to save these options, more info on that is described further
below):

```
lnd --bitcoin.active --bitcoin.testnet --debuglevel=debug --bitcoin.node=bitcoind --bitcoind.rpcuser=REPLACEME --bitcoind.rpcpass=REPLACEME --bitcoind.zmqpubrawblock=tcp://127.0.0.1:28332 --bitcoind.zmqpubrawtx=tcp://127.0.0.1:28333 --externalip=X.X.X.X
```

*NOTE:*
- The auth parameters `rpcuser` and `rpcpass` parameters can typically be
  determined by `lnd` for a `bitcoind` instance running under the same user,
  including when using cookie auth. In this case, you can exclude them from the
  `lnd` options entirely.
- If you DO choose to explicitly pass the auth parameters in your `lnd.conf` or
  command line options for `lnd` (`bitcoind.rpcuser` and `bitcoind.rpcpass` as
  shown in example command above), you must also specify the
  `bitcoind.zmqpubrawblock` and `bitcoind.zmqpubrawtx` options. Otherwise, `lnd`
  will attempt to get the configuration from your `bitcoin.conf`.
- You must ensure the same addresses are used for the `bitcoind.zmqpubrawblock`
  and `bitcoind.zmqpubrawtx` options passed to `lnd` as for the `zmqpubrawblock`
  and `zmqpubrawtx` passed in the `bitcoind` options respectively.
- When running lnd and picfightcoind on the same Windows machine, ensure you use
  127.0.0.1, not localhost, for all configuration options that require a TCP/IP
  host address.  If you use "localhost" as the host name, you may see extremely
  slow inter-process-communication between lnd and the picfightcoind backend.  If lnd
  is experiencing this issue, you'll see "Waiting for chain backend to finish
  sync, start_height=XXXXXX" as the last entry in the console or log output, and
  lnd will appear to hang.  Normal lnd output will quickly show multiple
  messages like this as lnd consumes blocks from picfightcoind.
- Don't connect more than two or three instances of `lnd` to `bitcoind`. With
  the default `bitcoind` settings, having more than one instance of `lnd`, or
  `lnd` plus any application that consumes the RPC could cause `lnd` to miss
  crucial updates from the backend.

# Macaroons

`lnd`'s authentication system is called **macaroons**, which are decentralized
bearer credentials allowing for delegation, attenuation, and other cool
features. You can learn more about them in Alex Akselrod's [writeup on
Github](https://github.com/picfight/pfclnd/issues/20).

Running `lnd` for the first time will by default generate the `admin.macaroon`,
`read_only.macaroon`, and `macaroons.db` files that are used to authenticate
into `lnd`. They will be stored in the network directory (default:
`lnddir/data/chain/bitcoin/mainnet`) so that it's possible to use a distinct
password for mainnet, testnet, simnet, etc. Note that if you specified an
alternative data directory (via the `--datadir` argument), you will have to
additionally pass the updated location of the `admin.macaroon` file into `lncli`
using the `--macaroonpath` argument.

To disable macaroons for testing, pass the `--no-macaroons` flag into *both*
`lnd` and `lncli`.

# Network Reachability

If you'd like to signal to other nodes on the network that you'll accept
incoming channels (as peers need to connect inbound to initiate a channel
funding workflow), then the `--externalip` flag should be set to your publicly
reachable IP address.

# Simnet vs. Testnet Development

If you are doing local development, such as for the tutorial, you'll want to
start both `pfcd` and `lnd` in the `simnet` mode. Simnet is similar to regtest
in that you'll be able to instantly mine blocks as needed to test `lnd`
locally. In order to start either daemon in the `simnet` mode use `simnet`
instead of `testnet`, adding the `--bitcoin.simnet` flag instead of the
`--bitcoin.testnet` flag.

Another relevant command line flag for local testing of new `lnd` developments
is the `--debughtlc` flag. When starting `lnd` with this flag, it'll be able to
automatically settle a special type of HTLC sent to it. This means that you
won't need to manually insert invoices in order to test payment connectivity.
To send this "special" HTLC type, include the `--debugsend` command at the end
of your `sendpayment` commands.


There are currently two primary ways to run `lnd`: one requires a local `pfcd`
instance with the RPC service exposed, and the other uses a fully integrated
light client powered by [neutrino](https://github.com/picfight/pfcneutrino).

# Creating an lnd.conf (Optional)

Optionally, if you'd like to have a persistent configuration between `lnd`
launches, allowing you to simply type `lnd --bitcoin.testnet --bitcoin.active`
at the command line, you can create an `lnd.conf`.

**On MacOS, located at:**
`/Users/[username]/Library/Application Support/Lnd/lnd.conf`

**On Linux, located at:**
`~/.lnd/lnd.conf`

Here's a sample `lnd.conf` for `pfcd` to get you started:
```
[Application Options]
debuglevel=trace
maxpendingchannels=10

[Picfightcoin]
bitcoin.active=1
```

Notice the `[Picfightcoin]` section. This section houses the parameters for the
Picfightcoin chain. `lnd` also supports Litecoin testnet4 (but not both PFC and LTC
at the same time), so when working with Litecoin be sure to set to parameters
for Litecoin accordingly. See a more detailed sample config file available
[here](https://github.com/picfight/pfclnd/blob/master/sample-lnd.conf)
and explore the other sections for node configuration, including `[Pfcd]`,
`[Picfightcoind]`, `[Neutrino]`, `[Ltcd]`, and `[Litecoind]` depending on which
chain and node type you're using.
