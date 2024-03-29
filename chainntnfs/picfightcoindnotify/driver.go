package picfightcoindnotify

import (
	"errors"
	"fmt"

	"github.com/picfight/pfclnd/chainntnfs"
	"github.com/picfight/pfcwallet/chain"
)

// createNewNotifier creates a new instance of the ChainNotifier interface
// implemented by PicfightcoindNotifier.
func createNewNotifier(args ...interface{}) (chainntnfs.ChainNotifier, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("incorrect number of arguments to "+
			".New(...), expected 2, instead passed %v", len(args))
	}

	chainConn, ok := args[0].(*chain.PicfightcoindConn)
	if !ok {
		return nil, errors.New("first argument to picfightcoindnotify.New " +
			"is incorrect, expected a *chain.PicfightcoindConn")
	}

	spendHintCache, ok := args[1].(chainntnfs.SpendHintCache)
	if !ok {
		return nil, errors.New("second argument to picfightcoindnotify.New " +
			"is incorrect, expected a chainntnfs.SpendHintCache")
	}

	confirmHintCache, ok := args[2].(chainntnfs.ConfirmHintCache)
	if !ok {
		return nil, errors.New("third argument to picfightcoindnotify.New " +
			"is incorrect, expected a chainntnfs.ConfirmHintCache")
	}

	return New(chainConn, spendHintCache, confirmHintCache), nil
}

// init registers a driver for the PfcdNotifier concrete implementation of the
// chainntnfs.ChainNotifier interface.
func init() {
	// Register the driver.
	notifier := &chainntnfs.NotifierDriver{
		NotifierType: notifierType,
		New:          createNewNotifier,
	}

	if err := chainntnfs.RegisterNotifier(notifier); err != nil {
		panic(fmt.Sprintf("failed to register notifier driver '%s': %v",
			notifierType, err))
	}
}
