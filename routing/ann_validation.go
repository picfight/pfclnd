package routing

import (
	"bytes"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-errors/errors"
	"github.com/picfight/pfcd/chaincfg/chainhash"
	"github.com/picfight/pfcd/pfcec"
	"github.com/picfight/pfclnd/lnwire"
)

// ValidateChannelAnn validates the channel announcement message and checks
// that node signatures covers the announcement message, and that the picfightcoin
// signatures covers the node keys.
func ValidateChannelAnn(a *lnwire.ChannelAnnouncement) error {
	// First, we'll compute the digest (h) which is to be signed by each of
	// the keys included within the node announcement message. This hash
	// digest includes all the keys, so the (up to 4 signatures) will
	// attest to the validity of each of the keys.
	data, err := a.DataToSign()
	if err != nil {
		return err
	}
	dataHash := chainhash.DoubleHashB(data)

	// First we'll verify that the passed picfightcoin key signature is indeed a
	// signature over the computed hash digest.
	bitcoinSig1, err := a.PicfightcoinSig1.ToSignature()
	if err != nil {
		return err
	}
	bitcoinKey1, err := pfcec.ParsePubKey(a.PicfightcoinKey1[:], pfcec.S256())
	if err != nil {
		return err
	}
	if !bitcoinSig1.Verify(dataHash, picfightcoinKey1) {
		return errors.New("can't verify first picfightcoin signature")
	}

	// If that checks out, then we'll verify that the second picfightcoin
	// signature is a valid signature of the picfightcoin public key over hash
	// digest as well.
	bitcoinSig2, err := a.PicfightcoinSig2.ToSignature()
	if err != nil {
		return err
	}
	bitcoinKey2, err := pfcec.ParsePubKey(a.PicfightcoinKey2[:], pfcec.S256())
	if err != nil {
		return err
	}
	if !bitcoinSig2.Verify(dataHash, picfightcoinKey2) {
		return errors.New("can't verify second picfightcoin signature")
	}

	// Both node signatures attached should indeed be a valid signature
	// over the selected digest of the channel announcement signature.
	nodeSig1, err := a.NodeSig1.ToSignature()
	if err != nil {
		return err
	}
	nodeKey1, err := pfcec.ParsePubKey(a.NodeID1[:], pfcec.S256())
	if err != nil {
		return err
	}
	if !nodeSig1.Verify(dataHash, nodeKey1) {
		return errors.New("can't verify data in first node signature")
	}

	nodeSig2, err := a.NodeSig2.ToSignature()
	if err != nil {
		return err
	}
	nodeKey2, err := pfcec.ParsePubKey(a.NodeID2[:], pfcec.S256())
	if err != nil {
		return err
	}
	if !nodeSig2.Verify(dataHash, nodeKey2) {
		return errors.New("can't verify data in second node signature")
	}

	return nil

}

// ValidateNodeAnn validates the node announcement by ensuring that the
// attached signature is needed a signature of the node announcement under the
// specified node public key.
func ValidateNodeAnn(a *lnwire.NodeAnnouncement) error {
	// Reconstruct the data of announcement which should be covered by the
	// signature so we can verify the signature shortly below
	data, err := a.DataToSign()
	if err != nil {
		return err
	}

	nodeSig, err := a.Signature.ToSignature()
	if err != nil {
		return err
	}
	nodeKey, err := pfcec.ParsePubKey(a.NodeID[:], pfcec.S256())
	if err != nil {
		return err
	}

	// Finally ensure that the passed signature is valid, if not we'll
	// return an error so this node announcement can be rejected.
	dataHash := chainhash.DoubleHashB(data)
	if !nodeSig.Verify(dataHash, nodeKey) {
		var msgBuf bytes.Buffer
		if _, err := lnwire.WriteMessage(&msgBuf, a, 0); err != nil {
			return err
		}

		return errors.Errorf("signature on NodeAnnouncement(%x) is "+
			"invalid: %x", nodeKey.SerializeCompressed(),
			msgBuf.Bytes())
	}

	return nil
}

// ValidateChannelUpdateAnn validates the channel update announcement by
// checking that the included signature covers he announcement and has been
// signed by the node's private key.
func ValidateChannelUpdateAnn(pubKey *pfcec.PublicKey,
	a *lnwire.ChannelUpdate) error {

	data, err := a.DataToSign()
	if err != nil {
		return errors.Errorf("unable to reconstruct message: %v", err)
	}
	dataHash := chainhash.DoubleHashB(data)

	nodeSig, err := a.Signature.ToSignature()
	if err != nil {
		return err
	}

	if !nodeSig.Verify(dataHash, pubKey) {
		return errors.Errorf("invalid signature for channel "+
			"update %v", spew.Sdump(a))
	}

	return nil
}
