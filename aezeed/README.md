# aezeed

[In this PR](https://github.com/picfight/pfclnd/pull/773) we add a new package implementing the aezeed cipher
seed scheme (based on [aez](http://web.cs.ucdavis.edu/~rogaway/aez/) ).

This is a new scheme developed that aims to overcome the
two major short comings of BIP39: a lack of a version, and a lack of a
wallet birthday. A lack a version means that wallets may not
necessarily know how to re-derive addresses during the recovery
process. A lack of a birthday means that wallets don’t know how far
back to look in the chain to ensure that they derive all the proper
user addresses. Additionally, BIP39 use a very weak KDF. We use
scrypt with modern parameters (n=32768, r=8, p=1). A set of benchmarks has
been added, on my laptop I get about 100ms per attempt):

```bash
⛰ go test -run=XXX -bench=.

goos: linux
goarch: amd64
pkg: github.com/picfight/pfclnd/aezeed
BenchmarkTomnemonic-4                 20          93280730 ns/op        33559670 B/op         36 allocs/op
BenchmarkToCipherSeed-4               10         102323892 ns/op        36915684 B/op         41 allocs/op
PASS
ok      github.com/picfight/pfclnd/aezeed  4.168s
```

Aside from addressing the shortcomings of BIP 39 a cipher  seed
can: be upgraded, and have it&#39;s password changed,

Sample seed:

```text
ability dance scatter raw fly dentist bar nominee exhaust wine snap super cost case coconut ticket spread funny grain chimney aspect business quiz ginger
```

## Plaintext aezeed encoding

The aezeed scheme addresses these two drawbacks and adds a number of
desirable features. First, we start with the following plaintext seed:

```text
1 byte internal version || 2 byte timestamp || 16 bytes of entropy
```

The version field is for wallets to be able to know how to re-derive
the keys of the wallet.

The 2 byte timestamp is expressed in Picfightcoin Days Genesis, meaning that
the number of days since the timestamp in Picfightcoin’s genesis block. This
allow us to save space, and also avoid using a wasteful level of
granularity. With the currently, this can express time up until 2188.

Finally, the entropy is raw entropy that should be used to derive
wallet’s HD root.

## aezeed enciphering/deciperhing

Next, we’ll take the plaintext seed described above and encipher it to
procure a final cipher text. We’ll then take this cipher text (the
CipherSeed) and encode that using a 24-word mnemonic. The enciphering
process takes a user defined passphrase. If no passphrase is provided,
then the string “aezeed” will be used.

To encipher a plaintext seed (19 bytes) to arrive at an enciphered
cipher seed (33 bytes), we apply the following operations:

* First we take the external version an append it to our buffer. The
external version describes how we encipher. For the first version
(version 0), we’ll use scrypt(n=32768, r=8, p=1) and aezeed.
* Next, we’ll use scrypt (with the version 9 params) to generate a
strong key for encryption. We’ll generate a 32-byte key using 5 bytes
as a salt. The usage of the salt is meant to make the creation of
rainbow tables infeasible.
* Next, the enciphering process. We use aez, modern AEAD with
nonce-misuse resistance properties. The important trait we exploit is
that it’s an arbitrary input length block cipher. Additionally, it
has what’s essentially a configurable MAC size. In our scheme we’ll use
a value of 8, which acts as a 64-bit checksum. We’ll encrypt with our
generated seed, and use an AD of (version || salt).
* Finally, we’ll encode this 33-byte cipher text using the default
world list of BIP 39 to produce 24 english words.

## Properties of the aezeed cipher seed

The aezeed cipher seed scheme has a few cool properties, notably:

* The mnemonic itself is a cipher text, meaning leaving it in
plaintext is advisable if the user also set a passphrase. This is in
contrast to BIP 39 where the mnemonic alone (without a passrphase) may
be sufficient to steal funds.
* A cipherseed can be modified to change the passphrase. This
means that if the users wants a stronger passphrase, they can decipher
(with the old passphrase), then encipher (with a new passphrase).
Compared to BIP 39, where if the users used a passphrase, since the
mapping is one way, they can’t change the passphrase of their existing
HD key chain.
* A cipher seed can be upgraded. Since we have an external version,
offline tools can be provided to decipher using the old params, and
encipher using the new params. In the future if we change ciphers,
change scrypt, or just the parameters of scrypt, then users can easily
upgrade their seed with an offline tool.
