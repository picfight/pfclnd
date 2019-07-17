package lnwallet

import (
	"github.com/picfight/pfcutil"
	"github.com/picfight/pfcwallet/wallet/txrules"
)

// DefaultDustLimit is used to calculate the dust HTLC amount which will be
// send to other node during funding process.
func DefaultDustLimit() pfcutil.Amount {
	return txrules.GetDustThreshold(P2WSHSize, txrules.DefaultRelayFeePerKb)
}
