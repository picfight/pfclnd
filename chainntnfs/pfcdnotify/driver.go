package pfcdnotify

import (
	"errors"
	"fmt"

	"github.com/picfight/pfcd/rpcclient"
	"github.com/picfight/pfclnd/chainntnfs"
)

// createNewNotifier creates a new instance of the ChainNotifier interface
// implemented by PfcdNotifier.
func createNewNotifier(args ...interface{}) (chainntnfs.ChainNotifier, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("incorrect number of arguments to "+
			".New(...), expected 2, instead passed %v", len(args))
	}

	config, ok := args[0].(*rpcclient.ConnConfig)
	if !ok {
		return nil, errors.New("first argument to pfcdnotifier.New " +
			"is incorrect, expected a *rpcclient.ConnConfig")
	}

	spendHintCache, ok := args[1].(chainntnfs.SpendHintCache)
	if !ok {
		return nil, errors.New("second argument to pfcdnotifier.New " +
			"is incorrect, expected a chainntnfs.SpendHintCache")
	}

	confirmHintCache, ok := args[2].(chainntnfs.ConfirmHintCache)
	if !ok {
		return nil, errors.New("third argument to pfcdnotifier.New " +
			"is incorrect, expected a chainntnfs.ConfirmHintCache")
	}

	return New(config, spendHintCache, confirmHintCache)
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
