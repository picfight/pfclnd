package lnwallet

import (
	"encoding/json"

	"github.com/picfight/pfcd/blockchain"
	"github.com/picfight/pfcd/rpcclient"
	"github.com/picfight/pfcutil"
)

const (
	// FeePerKwFloor is the lowest fee rate in sat/kw that we should use for
	// determining transaction fees.
	FeePerKwFloor SatPerKWeight = 253
)

// SatPerKVByte represents a fee rate in sat/kb.
type SatPerKVByte pfcutil.Amount

// FeeForVSize calculates the fee resulting from this fee rate and the given
// vsize in vbytes.
func (s SatPerKVByte) FeeForVSize(vbytes int64) pfcutil.Amount {
	return pfcutil.Amount(s) * pfcutil.Amount(vbytes) / 1000
}

// FeePerKWeight converts the current fee rate from sat/kb to sat/kw.
func (s SatPerKVByte) FeePerKWeight() SatPerKWeight {
	return SatPerKWeight(s / blockchain.WitnessScaleFactor)
}

// SatPerKWeight represents a fee rate in sat/kw.
type SatPerKWeight pfcutil.Amount

// FeeForWeight calculates the fee resulting from this fee rate and the given
// weight in weight units (wu).
func (s SatPerKWeight) FeeForWeight(wu int64) pfcutil.Amount {
	// The resulting fee is rounded down, as specified in BOLT#03.
	return pfcutil.Amount(s) * pfcutil.Amount(wu) / 1000
}

// FeePerKVByte converts the current fee rate from sat/kw to sat/kb.
func (s SatPerKWeight) FeePerKVByte() SatPerKVByte {
	return SatPerKVByte(s * blockchain.WitnessScaleFactor)
}

// FeeEstimator provides the ability to estimate on-chain transaction fees for
// various combinations of transaction sizes and desired confirmation time
// (measured by number of blocks).
type FeeEstimator interface {
	// EstimateFeePerKW takes in a target for the number of blocks until an
	// initial confirmation and returns the estimated fee expressed in
	// sat/kw.
	EstimateFeePerKW(numBlocks uint32) (SatPerKWeight, error)

	// Start signals the FeeEstimator to start any processes or goroutines
	// it needs to perform its duty.
	Start() error

	// Stop stops any spawned goroutines and cleans up the resources used
	// by the fee estimator.
	Stop() error

	// RelayFeePerKW returns the minimum fee rate required for transactions
	// to be relayed. This is also the basis for calculation of the dust
	// limit.
	RelayFeePerKW() SatPerKWeight
}

// StaticFeeEstimator will return a static value for all fee calculation
// requests. It is designed to be replaced by a proper fee calculation
// implementation. The fees are not accessible directly, because changing them
// would not be thread safe.
type StaticFeeEstimator struct {
	// feePerKW is the static fee rate in satoshis-per-vbyte that will be
	// returned by this fee estimator.
	feePerKW SatPerKWeight

	// relayFee is the minimum fee rate required for transactions to be
	// relayed.
	relayFee SatPerKWeight
}

// NewStaticFeeEstimator returns a new static fee estimator instance.
func NewStaticFeeEstimator(feePerKW,
	relayFee SatPerKWeight) *StaticFeeEstimator {

	return &StaticFeeEstimator{
		feePerKW: feePerKW,
		relayFee: relayFee,
	}
}

// EstimateFeePerKW will return a static value for fee calculations.
//
// NOTE: This method is part of the FeeEstimator interface.
func (e StaticFeeEstimator) EstimateFeePerKW(numBlocks uint32) (SatPerKWeight, error) {
	return e.feePerKW, nil
}

// RelayFeePerKW returns the minimum fee rate required for transactions to be
// relayed.
//
// NOTE: This method is part of the FeeEstimator interface.
func (e StaticFeeEstimator) RelayFeePerKW() SatPerKWeight {
	return e.relayFee
}

// Start signals the FeeEstimator to start any processes or goroutines
// it needs to perform its duty.
//
// NOTE: This method is part of the FeeEstimator interface.
func (e StaticFeeEstimator) Start() error {
	return nil
}

// Stop stops any spawned goroutines and cleans up the resources used
// by the fee estimator.
//
// NOTE: This method is part of the FeeEstimator interface.
func (e StaticFeeEstimator) Stop() error {
	return nil
}

// A compile-time assertion to ensure that StaticFeeEstimator implements the
// FeeEstimator interface.
var _ FeeEstimator = (*StaticFeeEstimator)(nil)

// PfcdFeeEstimator is an implementation of the FeeEstimator interface backed
// by the RPC interface of an active pfcd node. This implementation will proxy
// any fee estimation requests to pfcd's RPC interface.
type PfcdFeeEstimator struct {
	// fallbackFeePerKW is the fall back fee rate in sat/kw that is returned
	// if the fee estimator does not yet have enough data to actually
	// produce fee estimates.
	fallbackFeePerKW SatPerKWeight

	// minFeePerKW is the minimum fee, in sat/kw, that we should enforce.
	// This will be used as the default fee rate for a transaction when the
	// estimated fee rate is too low to allow the transaction to propagate
	// through the network.
	minFeePerKW SatPerKWeight

	pfcdConn *rpcclient.Client
}

// NewPfcdFeeEstimator creates a new PfcdFeeEstimator given a fully populated
// rpc config that is able to successfully connect and authenticate with the
// pfcd node, and also a fall back fee rate. The fallback fee rate is used in
// the occasion that the estimator has insufficient data, or returns zero for a
// fee estimate.
func NewPfcdFeeEstimator(rpcConfig rpcclient.ConnConfig,
	fallBackFeeRate SatPerKWeight) (*PfcdFeeEstimator, error) {

	rpcConfig.DisableConnectOnNew = true
	rpcConfig.DisableAutoReconnect = false
	chainConn, err := rpcclient.New(&rpcConfig, nil)
	if err != nil {
		return nil, err
	}

	return &PfcdFeeEstimator{
		fallbackFeePerKW: fallBackFeeRate,
		pfcdConn:         chainConn,
	}, nil
}

// Start signals the FeeEstimator to start any processes or goroutines
// it needs to perform its duty.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PfcdFeeEstimator) Start() error {
	if err := b.pfcdConn.Connect(20); err != nil {
		return err
	}

	// Once the connection to the backend node has been established, we'll
	// query it for its minimum relay fee.
	info, err := b.pfcdConn.GetInfo()
	if err != nil {
		return err
	}

	relayFee, err := pfcutil.NewAmount(info.RelayFee)
	if err != nil {
		return err
	}

	// The fee rate is expressed in sat/kb, so we'll manually convert it to
	// our desired sat/kw rate.
	minRelayFeePerKw := SatPerKVByte(relayFee).FeePerKWeight()

	// By default, we'll use the backend node's minimum relay fee as the
	// minimum fee rate we'll propose for transacations. However, if this
	// happens to be lower than our fee floor, we'll enforce that instead.
	b.minFeePerKW = minRelayFeePerKw
	if b.minFeePerKW < FeePerKwFloor {
		b.minFeePerKW = FeePerKwFloor
	}

	walletLog.Debugf("Using minimum fee rate of %v sat/kw",
		int64(b.minFeePerKW))

	return nil
}

// Stop stops any spawned goroutines and cleans up the resources used
// by the fee estimator.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PfcdFeeEstimator) Stop() error {
	b.pfcdConn.Shutdown()

	return nil
}

// EstimateFeePerKW takes in a target for the number of blocks until an initial
// confirmation and returns the estimated fee expressed in sat/kw.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PfcdFeeEstimator) EstimateFeePerKW(numBlocks uint32) (SatPerKWeight, error) {
	feeEstimate, err := b.fetchEstimate(numBlocks)
	switch {
	// If the estimator doesn't have enough data, or returns an error, then
	// to return a proper value, then we'll return the default fall back
	// fee rate.
	case err != nil:
		walletLog.Errorf("unable to query estimator: %v", err)
		fallthrough

	case feeEstimate == 0:
		return b.fallbackFeePerKW, nil
	}

	return feeEstimate, nil
}

// RelayFeePerKW returns the minimum fee rate required for transactions to be
// relayed.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PfcdFeeEstimator) RelayFeePerKW() SatPerKWeight {
	return b.minFeePerKW
}

// fetchEstimate returns a fee estimate for a transaction to be confirmed in
// confTarget blocks. The estimate is returned in sat/kw.
func (b *PfcdFeeEstimator) fetchEstimate(confTarget uint32) (SatPerKWeight, error) {
	// First, we'll fetch the estimate for our confirmation target.
	btcPerKB, err := b.pfcdConn.EstimateFee(int64(confTarget))
	if err != nil {
		return 0, err
	}

	// Next, we'll convert the returned value to satoshis, as it's
	// currently returned in PFC.
	satPerKB, err := pfcutil.NewAmount(btcPerKB)
	if err != nil {
		return 0, err
	}

	// Since we use fee rates in sat/kw internally, we'll convert the
	// estimated fee rate from its sat/kb representation to sat/kw.
	satPerKw := SatPerKVByte(satPerKB).FeePerKWeight()

	// Finally, we'll enforce our fee floor.
	if satPerKw < b.minFeePerKW {
		walletLog.Debugf("Estimated fee rate of %v sat/kw is too low, "+
			"using fee floor of %v sat/kw instead", satPerKw,
			b.minFeePerKW)
		satPerKw = b.minFeePerKW
	}

	walletLog.Debugf("Returning %v sat/kw for conf target of %v",
		int64(satPerKw), confTarget)

	return satPerKw, nil
}

// A compile-time assertion to ensure that PfcdFeeEstimator implements the
// FeeEstimator interface.
var _ FeeEstimator = (*PfcdFeeEstimator)(nil)

// PicfightcoindFeeEstimator is an implementation of the FeeEstimator interface
// backed by the RPC interface of an active picfightcoind node. This implementation
// will proxy any fee estimation requests to picfightcoind's RPC interface.
type PicfightcoindFeeEstimator struct {
	// fallbackFeePerKW is the fallback fee rate in sat/kw that is returned
	// if the fee estimator does not yet have enough data to actually
	// produce fee estimates.
	fallbackFeePerKW SatPerKWeight

	// minFeePerKW is the minimum fee, in sat/kw, that we should enforce.
	// This will be used as the default fee rate for a transaction when the
	// estimated fee rate is too low to allow the transaction to propagate
	// through the network.
	minFeePerKW SatPerKWeight

	bitcoindConn *rpcclient.Client
}

// NewPicfightcoindFeeEstimator creates a new PicfightcoindFeeEstimator given a fully
// populated rpc config that is able to successfully connect and authenticate
// with the picfightcoind node, and also a fall back fee rate. The fallback fee rate
// is used in the occasion that the estimator has insufficient data, or returns
// zero for a fee estimate.
func NewPicfightcoindFeeEstimator(rpcConfig rpcclient.ConnConfig,
	fallBackFeeRate SatPerKWeight) (*PicfightcoindFeeEstimator, error) {

	rpcConfig.DisableConnectOnNew = true
	rpcConfig.DisableAutoReconnect = false
	rpcConfig.DisableTLS = true
	rpcConfig.HTTPPostMode = true
	chainConn, err := rpcclient.New(&rpcConfig, nil)
	if err != nil {
		return nil, err
	}

	return &PicfightcoindFeeEstimator{
		fallbackFeePerKW: fallBackFeeRate,
		bitcoindConn:     chainConn,
	}, nil
}

// Start signals the FeeEstimator to start any processes or goroutines
// it needs to perform its duty.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PicfightcoindFeeEstimator) Start() error {
	// Once the connection to the backend node has been established, we'll
	// query it for its minimum relay fee. Since the `getinfo` RPC has been
	// deprecated for `bitcoind`, we'll need to send a `getnetworkinfo`
	// command as a raw request.
	resp, err := b.bitcoindConn.RawRequest("getnetworkinfo", nil)
	if err != nil {
		return err
	}

	// Parse the response to retrieve the relay fee in sat/KB.
	info := struct {
		RelayFee float64 `json:"relayfee"`
	}{}
	if err := json.Unmarshal(resp, &info); err != nil {
		return err
	}

	relayFee, err := pfcutil.NewAmount(info.RelayFee)
	if err != nil {
		return err
	}

	// The fee rate is expressed in sat/kb, so we'll manually convert it to
	// our desired sat/kw rate.
	minRelayFeePerKw := SatPerKVByte(relayFee).FeePerKWeight()

	// By default, we'll use the backend node's minimum relay fee as the
	// minimum fee rate we'll propose for transacations. However, if this
	// happens to be lower than our fee floor, we'll enforce that instead.
	b.minFeePerKW = minRelayFeePerKw
	if b.minFeePerKW < FeePerKwFloor {
		b.minFeePerKW = FeePerKwFloor
	}

	walletLog.Debugf("Using minimum fee rate of %v sat/kw",
		int64(b.minFeePerKW))

	return nil
}

// Stop stops any spawned goroutines and cleans up the resources used
// by the fee estimator.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PicfightcoindFeeEstimator) Stop() error {
	return nil
}

// EstimateFeePerKW takes in a target for the number of blocks until an initial
// confirmation and returns the estimated fee expressed in sat/kw.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PicfightcoindFeeEstimator) EstimateFeePerKW(numBlocks uint32) (SatPerKWeight, error) {
	feeEstimate, err := b.fetchEstimate(numBlocks)
	switch {
	// If the estimator doesn't have enough data, or returns an error, then
	// to return a proper value, then we'll return the default fall back
	// fee rate.
	case err != nil:
		walletLog.Errorf("unable to query estimator: %v", err)
		fallthrough

	case feeEstimate == 0:
		return b.fallbackFeePerKW, nil
	}

	return feeEstimate, nil
}

// RelayFeePerKW returns the minimum fee rate required for transactions to be
// relayed.
//
// NOTE: This method is part of the FeeEstimator interface.
func (b *PicfightcoindFeeEstimator) RelayFeePerKW() SatPerKWeight {
	return b.minFeePerKW
}

// fetchEstimate returns a fee estimate for a transaction to be confirmed in
// confTarget blocks. The estimate is returned in sat/kw.
func (b *PicfightcoindFeeEstimator) fetchEstimate(confTarget uint32) (SatPerKWeight, error) {
	// First, we'll send an "estimatesmartfee" command as a raw request,
	// since it isn't supported by pfcd but is available in picfightcoind.
	target, err := json.Marshal(uint64(confTarget))
	if err != nil {
		return 0, err
	}
	// TODO: Allow selection of economical/conservative modifiers.
	resp, err := b.bitcoindConn.RawRequest(
		"estimatesmartfee", []json.RawMessage{target},
	)
	if err != nil {
		return 0, err
	}

	// Next, we'll parse the response to get the PFC per KB.
	feeEstimate := struct {
		FeeRate float64 `json:"feerate"`
	}{}
	err = json.Unmarshal(resp, &feeEstimate)
	if err != nil {
		return 0, err
	}

	// Next, we'll convert the returned value to satoshis, as it's currently
	// returned in PFC.
	satPerKB, err := pfcutil.NewAmount(feeEstimate.FeeRate)
	if err != nil {
		return 0, err
	}

	// Since we use fee rates in sat/kw internally, we'll convert the
	// estimated fee rate from its sat/kb representation to sat/kw.
	satPerKw := SatPerKVByte(satPerKB).FeePerKWeight()

	// Finally, we'll enforce our fee floor.
	if satPerKw < b.minFeePerKW {
		walletLog.Debugf("Estimated fee rate of %v sat/kw is too low, "+
			"using fee floor of %v sat/kw instead", satPerKw,
			b.minFeePerKW)

		satPerKw = b.minFeePerKW
	}

	walletLog.Debugf("Returning %v sat/kw for conf target of %v",
		int64(satPerKw), confTarget)

	return satPerKw, nil
}

// A compile-time assertion to ensure that PicfightcoindFeeEstimator implements the
// FeeEstimator interface.
var _ FeeEstimator = (*PicfightcoindFeeEstimator)(nil)
