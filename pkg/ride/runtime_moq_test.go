// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ride

import (
	"github.com/wavesplatform/gowaves/pkg/proto"
	"github.com/wavesplatform/gowaves/pkg/types"
	"sync"
)

// Ensure, that mockRideEnvironment does implement environment.
// If this is not the case, regenerate this file with moq.
var _ environment = &mockRideEnvironment{}

// mockRideEnvironment is a mock implementation of environment.
//
// 	func TestSomethingThatUsesenvironment(t *testing.T) {
//
// 		// make and configure a mocked environment
// 		mockedenvironment := &mockRideEnvironment{
// 			blockFunc: func() rideObject {
// 				panic("mock out the block method")
// 			},
// 			blockV5ActivatedFunc: func() bool {
// 				panic("mock out the blockV5Activated method")
// 			},
// 			checkMessageLengthFunc: func(n int) bool {
// 				panic("mock out the checkMessageLength method")
// 			},
// 			heightFunc: func() rideInt {
// 				panic("mock out the height method")
// 			},
// 			internalPaymentsValidationHeightFunc: func() uint64 {
// 				panic("mock out the internalPaymentsValidationHeight method")
// 			},
// 			invocationFunc: func() rideObject {
// 				panic("mock out the invocation method")
// 			},
// 			isProtobufTxFunc: func() bool {
// 				panic("mock out the isProtobufTx method")
// 			},
// 			libVersionFunc: func() int {
// 				panic("mock out the libVersion method")
// 			},
// 			maxDataEntriesSizeFunc: func() int {
// 				panic("mock out the maxDataEntriesSize method")
// 			},
// 			rideV6ActivatedFunc: func() bool {
// 				panic("mock out the rideV6Activated method")
// 			},
// 			schemeFunc: func() byte {
// 				panic("mock out the scheme method")
// 			},
// 			setInvocationFunc: func(inv rideObject)  {
// 				panic("mock out the setInvocation method")
// 			},
// 			setNewDAppAddressFunc: func(address proto.WavesAddress)  {
// 				panic("mock out the setNewDAppAddress method")
// 			},
// 			stateFunc: func() types.SmartState {
// 				panic("mock out the state method")
// 			},
// 			takeStringFunc: func(s string, n int) rideString {
// 				panic("mock out the takeString method")
// 			},
// 			thisFunc: func() rideType {
// 				panic("mock out the this method")
// 			},
// 			timestampFunc: func() uint64 {
// 				panic("mock out the timestamp method")
// 			},
// 			transactionFunc: func() rideObject {
// 				panic("mock out the transaction method")
// 			},
// 			txIDFunc: func() rideType {
// 				panic("mock out the txID method")
// 			},
// 			validateInternalPaymentsFunc: func() bool {
// 				panic("mock out the validateInternalPayments method")
// 			},
// 		}
//
// 		// use mockedenvironment in code that requires environment
// 		// and then make assertions.
//
// 	}
type mockRideEnvironment struct {
	// blockFunc mocks the block method.
	blockFunc func() rideObject

	// blockV5ActivatedFunc mocks the blockV5Activated method.
	blockV5ActivatedFunc func() bool

	// checkMessageLengthFunc mocks the checkMessageLength method.
	checkMessageLengthFunc func(n int) bool

	// heightFunc mocks the height method.
	heightFunc func() rideInt

	// internalPaymentsValidationHeightFunc mocks the internalPaymentsValidationHeight method.
	internalPaymentsValidationHeightFunc func() uint64

	// invocationFunc mocks the invocation method.
	invocationFunc func() rideObject

	// isProtobufTxFunc mocks the isProtobufTx method.
	isProtobufTxFunc func() bool

	// libVersionFunc mocks the libVersion method.
	libVersionFunc func() int

	// maxDataEntriesSizeFunc mocks the maxDataEntriesSize method.
	maxDataEntriesSizeFunc func() int

	// rideV6ActivatedFunc mocks the rideV6Activated method.
	rideV6ActivatedFunc func() bool

	// schemeFunc mocks the scheme method.
	schemeFunc func() byte

	// setInvocationFunc mocks the setInvocation method.
	setInvocationFunc func(inv rideObject)

	// setNewDAppAddressFunc mocks the setNewDAppAddress method.
	setNewDAppAddressFunc func(address proto.WavesAddress)

	// stateFunc mocks the state method.
	stateFunc func() types.SmartState

	// takeStringFunc mocks the takeString method.
	takeStringFunc func(s string, n int) rideString

	// thisFunc mocks the this method.
	thisFunc func() rideType

	// timestampFunc mocks the timestamp method.
	timestampFunc func() uint64

	// transactionFunc mocks the transaction method.
	transactionFunc func() rideObject

	// txIDFunc mocks the txID method.
	txIDFunc func() rideType

	// validateInternalPaymentsFunc mocks the validateInternalPayments method.
	validateInternalPaymentsFunc func() bool

	// calls tracks calls to the methods.
	calls struct {
		// block holds details about calls to the block method.
		block []struct {
		}
		// blockV5Activated holds details about calls to the blockV5Activated method.
		blockV5Activated []struct {
		}
		// checkMessageLength holds details about calls to the checkMessageLength method.
		checkMessageLength []struct {
			// N is the n argument value.
			N int
		}
		// height holds details about calls to the height method.
		height []struct {
		}
		// internalPaymentsValidationHeight holds details about calls to the internalPaymentsValidationHeight method.
		internalPaymentsValidationHeight []struct {
		}
		// invocation holds details about calls to the invocation method.
		invocation []struct {
		}
		// isProtobufTx holds details about calls to the isProtobufTx method.
		isProtobufTx []struct {
		}
		// libVersion holds details about calls to the libVersion method.
		libVersion []struct {
		}
		// maxDataEntriesSize holds details about calls to the maxDataEntriesSize method.
		maxDataEntriesSize []struct {
		}
		// rideV6Activated holds details about calls to the rideV6Activated method.
		rideV6Activated []struct {
		}
		// scheme holds details about calls to the scheme method.
		scheme []struct {
		}
		// setInvocation holds details about calls to the setInvocation method.
		setInvocation []struct {
			// Inv is the inv argument value.
			Inv rideObject
		}
		// setNewDAppAddress holds details about calls to the setNewDAppAddress method.
		setNewDAppAddress []struct {
			// Address is the address argument value.
			Address proto.WavesAddress
		}
		// state holds details about calls to the state method.
		state []struct {
		}
		// takeString holds details about calls to the takeString method.
		takeString []struct {
			// S is the s argument value.
			S string
			// N is the n argument value.
			N int
		}
		// this holds details about calls to the this method.
		this []struct {
		}
		// timestamp holds details about calls to the timestamp method.
		timestamp []struct {
		}
		// transaction holds details about calls to the transaction method.
		transaction []struct {
		}
		// txID holds details about calls to the txID method.
		txID []struct {
		}
		// validateInternalPayments holds details about calls to the validateInternalPayments method.
		validateInternalPayments []struct {
		}
	}
	lockblock                            sync.RWMutex
	lockblockV5Activated                 sync.RWMutex
	lockcheckMessageLength               sync.RWMutex
	lockheight                           sync.RWMutex
	lockinternalPaymentsValidationHeight sync.RWMutex
	lockinvocation                       sync.RWMutex
	lockisProtobufTx                     sync.RWMutex
	locklibVersion                       sync.RWMutex
	lockmaxDataEntriesSize               sync.RWMutex
	lockrideV6Activated                  sync.RWMutex
	lockscheme                           sync.RWMutex
	locksetInvocation                    sync.RWMutex
	locksetNewDAppAddress                sync.RWMutex
	lockstate                            sync.RWMutex
	locktakeString                       sync.RWMutex
	lockthis                             sync.RWMutex
	locktimestamp                        sync.RWMutex
	locktransaction                      sync.RWMutex
	locktxID                             sync.RWMutex
	lockvalidateInternalPayments         sync.RWMutex
}

// block calls blockFunc.
func (mock *mockRideEnvironment) block() rideObject {
	if mock.blockFunc == nil {
		panic("mockRideEnvironment.blockFunc: method is nil but environment.block was just called")
	}
	callInfo := struct {
	}{}
	mock.lockblock.Lock()
	mock.calls.block = append(mock.calls.block, callInfo)
	mock.lockblock.Unlock()
	return mock.blockFunc()
}

// blockCalls gets all the calls that were made to block.
// Check the length with:
//     len(mockedenvironment.blockCalls())
func (mock *mockRideEnvironment) blockCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockblock.RLock()
	calls = mock.calls.block
	mock.lockblock.RUnlock()
	return calls
}

// blockV5Activated calls blockV5ActivatedFunc.
func (mock *mockRideEnvironment) blockV5Activated() bool {
	if mock.blockV5ActivatedFunc == nil {
		panic("mockRideEnvironment.blockV5ActivatedFunc: method is nil but environment.blockV5Activated was just called")
	}
	callInfo := struct {
	}{}
	mock.lockblockV5Activated.Lock()
	mock.calls.blockV5Activated = append(mock.calls.blockV5Activated, callInfo)
	mock.lockblockV5Activated.Unlock()
	return mock.blockV5ActivatedFunc()
}

// blockV5ActivatedCalls gets all the calls that were made to blockV5Activated.
// Check the length with:
//     len(mockedenvironment.blockV5ActivatedCalls())
func (mock *mockRideEnvironment) blockV5ActivatedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockblockV5Activated.RLock()
	calls = mock.calls.blockV5Activated
	mock.lockblockV5Activated.RUnlock()
	return calls
}

// checkMessageLength calls checkMessageLengthFunc.
func (mock *mockRideEnvironment) checkMessageLength(n int) bool {
	if mock.checkMessageLengthFunc == nil {
		panic("mockRideEnvironment.checkMessageLengthFunc: method is nil but environment.checkMessageLength was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockcheckMessageLength.Lock()
	mock.calls.checkMessageLength = append(mock.calls.checkMessageLength, callInfo)
	mock.lockcheckMessageLength.Unlock()
	return mock.checkMessageLengthFunc(n)
}

// checkMessageLengthCalls gets all the calls that were made to checkMessageLength.
// Check the length with:
//     len(mockedenvironment.checkMessageLengthCalls())
func (mock *mockRideEnvironment) checkMessageLengthCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockcheckMessageLength.RLock()
	calls = mock.calls.checkMessageLength
	mock.lockcheckMessageLength.RUnlock()
	return calls
}

// height calls heightFunc.
func (mock *mockRideEnvironment) height() rideInt {
	if mock.heightFunc == nil {
		panic("mockRideEnvironment.heightFunc: method is nil but environment.height was just called")
	}
	callInfo := struct {
	}{}
	mock.lockheight.Lock()
	mock.calls.height = append(mock.calls.height, callInfo)
	mock.lockheight.Unlock()
	return mock.heightFunc()
}

// heightCalls gets all the calls that were made to height.
// Check the length with:
//     len(mockedenvironment.heightCalls())
func (mock *mockRideEnvironment) heightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockheight.RLock()
	calls = mock.calls.height
	mock.lockheight.RUnlock()
	return calls
}

// internalPaymentsValidationHeight calls internalPaymentsValidationHeightFunc.
func (mock *mockRideEnvironment) internalPaymentsValidationHeight() uint64 {
	if mock.internalPaymentsValidationHeightFunc == nil {
		panic("mockRideEnvironment.internalPaymentsValidationHeightFunc: method is nil but environment.internalPaymentsValidationHeight was just called")
	}
	callInfo := struct {
	}{}
	mock.lockinternalPaymentsValidationHeight.Lock()
	mock.calls.internalPaymentsValidationHeight = append(mock.calls.internalPaymentsValidationHeight, callInfo)
	mock.lockinternalPaymentsValidationHeight.Unlock()
	return mock.internalPaymentsValidationHeightFunc()
}

// internalPaymentsValidationHeightCalls gets all the calls that were made to internalPaymentsValidationHeight.
// Check the length with:
//     len(mockedenvironment.internalPaymentsValidationHeightCalls())
func (mock *mockRideEnvironment) internalPaymentsValidationHeightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockinternalPaymentsValidationHeight.RLock()
	calls = mock.calls.internalPaymentsValidationHeight
	mock.lockinternalPaymentsValidationHeight.RUnlock()
	return calls
}

// invocation calls invocationFunc.
func (mock *mockRideEnvironment) invocation() rideObject {
	if mock.invocationFunc == nil {
		panic("mockRideEnvironment.invocationFunc: method is nil but environment.invocation was just called")
	}
	callInfo := struct {
	}{}
	mock.lockinvocation.Lock()
	mock.calls.invocation = append(mock.calls.invocation, callInfo)
	mock.lockinvocation.Unlock()
	return mock.invocationFunc()
}

// invocationCalls gets all the calls that were made to invocation.
// Check the length with:
//     len(mockedenvironment.invocationCalls())
func (mock *mockRideEnvironment) invocationCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockinvocation.RLock()
	calls = mock.calls.invocation
	mock.lockinvocation.RUnlock()
	return calls
}

// isProtobufTx calls isProtobufTxFunc.
func (mock *mockRideEnvironment) isProtobufTx() bool {
	if mock.isProtobufTxFunc == nil {
		panic("mockRideEnvironment.isProtobufTxFunc: method is nil but environment.isProtobufTx was just called")
	}
	callInfo := struct {
	}{}
	mock.lockisProtobufTx.Lock()
	mock.calls.isProtobufTx = append(mock.calls.isProtobufTx, callInfo)
	mock.lockisProtobufTx.Unlock()
	return mock.isProtobufTxFunc()
}

// isProtobufTxCalls gets all the calls that were made to isProtobufTx.
// Check the length with:
//     len(mockedenvironment.isProtobufTxCalls())
func (mock *mockRideEnvironment) isProtobufTxCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockisProtobufTx.RLock()
	calls = mock.calls.isProtobufTx
	mock.lockisProtobufTx.RUnlock()
	return calls
}

// libVersion calls libVersionFunc.
func (mock *mockRideEnvironment) libVersion() int {
	if mock.libVersionFunc == nil {
		panic("mockRideEnvironment.libVersionFunc: method is nil but environment.libVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.locklibVersion.Lock()
	mock.calls.libVersion = append(mock.calls.libVersion, callInfo)
	mock.locklibVersion.Unlock()
	return mock.libVersionFunc()
}

// libVersionCalls gets all the calls that were made to libVersion.
// Check the length with:
//     len(mockedenvironment.libVersionCalls())
func (mock *mockRideEnvironment) libVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.locklibVersion.RLock()
	calls = mock.calls.libVersion
	mock.locklibVersion.RUnlock()
	return calls
}

// maxDataEntriesSize calls maxDataEntriesSizeFunc.
func (mock *mockRideEnvironment) maxDataEntriesSize() int {
	if mock.maxDataEntriesSizeFunc == nil {
		panic("mockRideEnvironment.maxDataEntriesSizeFunc: method is nil but environment.maxDataEntriesSize was just called")
	}
	callInfo := struct {
	}{}
	mock.lockmaxDataEntriesSize.Lock()
	mock.calls.maxDataEntriesSize = append(mock.calls.maxDataEntriesSize, callInfo)
	mock.lockmaxDataEntriesSize.Unlock()
	return mock.maxDataEntriesSizeFunc()
}

// maxDataEntriesSizeCalls gets all the calls that were made to maxDataEntriesSize.
// Check the length with:
//     len(mockedenvironment.maxDataEntriesSizeCalls())
func (mock *mockRideEnvironment) maxDataEntriesSizeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockmaxDataEntriesSize.RLock()
	calls = mock.calls.maxDataEntriesSize
	mock.lockmaxDataEntriesSize.RUnlock()
	return calls
}

// rideV6Activated calls rideV6ActivatedFunc.
func (mock *mockRideEnvironment) rideV6Activated() bool {
	if mock.rideV6ActivatedFunc == nil {
		panic("mockRideEnvironment.rideV6ActivatedFunc: method is nil but environment.rideV6Activated was just called")
	}
	callInfo := struct {
	}{}
	mock.lockrideV6Activated.Lock()
	mock.calls.rideV6Activated = append(mock.calls.rideV6Activated, callInfo)
	mock.lockrideV6Activated.Unlock()
	return mock.rideV6ActivatedFunc()
}

// rideV6ActivatedCalls gets all the calls that were made to rideV6Activated.
// Check the length with:
//     len(mockedenvironment.rideV6ActivatedCalls())
func (mock *mockRideEnvironment) rideV6ActivatedCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockrideV6Activated.RLock()
	calls = mock.calls.rideV6Activated
	mock.lockrideV6Activated.RUnlock()
	return calls
}

// scheme calls schemeFunc.
func (mock *mockRideEnvironment) scheme() byte {
	if mock.schemeFunc == nil {
		panic("mockRideEnvironment.schemeFunc: method is nil but environment.scheme was just called")
	}
	callInfo := struct {
	}{}
	mock.lockscheme.Lock()
	mock.calls.scheme = append(mock.calls.scheme, callInfo)
	mock.lockscheme.Unlock()
	return mock.schemeFunc()
}

// schemeCalls gets all the calls that were made to scheme.
// Check the length with:
//     len(mockedenvironment.schemeCalls())
func (mock *mockRideEnvironment) schemeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockscheme.RLock()
	calls = mock.calls.scheme
	mock.lockscheme.RUnlock()
	return calls
}

// setInvocation calls setInvocationFunc.
func (mock *mockRideEnvironment) setInvocation(inv rideObject) {
	if mock.setInvocationFunc == nil {
		panic("mockRideEnvironment.setInvocationFunc: method is nil but environment.setInvocation was just called")
	}
	callInfo := struct {
		Inv rideObject
	}{
		Inv: inv,
	}
	mock.locksetInvocation.Lock()
	mock.calls.setInvocation = append(mock.calls.setInvocation, callInfo)
	mock.locksetInvocation.Unlock()
	mock.setInvocationFunc(inv)
}

// setInvocationCalls gets all the calls that were made to setInvocation.
// Check the length with:
//     len(mockedenvironment.setInvocationCalls())
func (mock *mockRideEnvironment) setInvocationCalls() []struct {
	Inv rideObject
} {
	var calls []struct {
		Inv rideObject
	}
	mock.locksetInvocation.RLock()
	calls = mock.calls.setInvocation
	mock.locksetInvocation.RUnlock()
	return calls
}

// setNewDAppAddress calls setNewDAppAddressFunc.
func (mock *mockRideEnvironment) setNewDAppAddress(address proto.WavesAddress) {
	if mock.setNewDAppAddressFunc == nil {
		panic("mockRideEnvironment.setNewDAppAddressFunc: method is nil but environment.setNewDAppAddress was just called")
	}
	callInfo := struct {
		Address proto.WavesAddress
	}{
		Address: address,
	}
	mock.locksetNewDAppAddress.Lock()
	mock.calls.setNewDAppAddress = append(mock.calls.setNewDAppAddress, callInfo)
	mock.locksetNewDAppAddress.Unlock()
	mock.setNewDAppAddressFunc(address)
}

// setNewDAppAddressCalls gets all the calls that were made to setNewDAppAddress.
// Check the length with:
//     len(mockedenvironment.setNewDAppAddressCalls())
func (mock *mockRideEnvironment) setNewDAppAddressCalls() []struct {
	Address proto.WavesAddress
} {
	var calls []struct {
		Address proto.WavesAddress
	}
	mock.locksetNewDAppAddress.RLock()
	calls = mock.calls.setNewDAppAddress
	mock.locksetNewDAppAddress.RUnlock()
	return calls
}

// state calls stateFunc.
func (mock *mockRideEnvironment) state() types.SmartState {
	if mock.stateFunc == nil {
		panic("mockRideEnvironment.stateFunc: method is nil but environment.state was just called")
	}
	callInfo := struct {
	}{}
	mock.lockstate.Lock()
	mock.calls.state = append(mock.calls.state, callInfo)
	mock.lockstate.Unlock()
	return mock.stateFunc()
}

// stateCalls gets all the calls that were made to state.
// Check the length with:
//     len(mockedenvironment.stateCalls())
func (mock *mockRideEnvironment) stateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockstate.RLock()
	calls = mock.calls.state
	mock.lockstate.RUnlock()
	return calls
}

// takeString calls takeStringFunc.
func (mock *mockRideEnvironment) takeString(s string, n int) rideString {
	if mock.takeStringFunc == nil {
		panic("mockRideEnvironment.takeStringFunc: method is nil but environment.takeString was just called")
	}
	callInfo := struct {
		S string
		N int
	}{
		S: s,
		N: n,
	}
	mock.locktakeString.Lock()
	mock.calls.takeString = append(mock.calls.takeString, callInfo)
	mock.locktakeString.Unlock()
	return mock.takeStringFunc(s, n)
}

// takeStringCalls gets all the calls that were made to takeString.
// Check the length with:
//     len(mockedenvironment.takeStringCalls())
func (mock *mockRideEnvironment) takeStringCalls() []struct {
	S string
	N int
} {
	var calls []struct {
		S string
		N int
	}
	mock.locktakeString.RLock()
	calls = mock.calls.takeString
	mock.locktakeString.RUnlock()
	return calls
}

// this calls thisFunc.
func (mock *mockRideEnvironment) this() rideType {
	if mock.thisFunc == nil {
		panic("mockRideEnvironment.thisFunc: method is nil but environment.this was just called")
	}
	callInfo := struct {
	}{}
	mock.lockthis.Lock()
	mock.calls.this = append(mock.calls.this, callInfo)
	mock.lockthis.Unlock()
	return mock.thisFunc()
}

// thisCalls gets all the calls that were made to this.
// Check the length with:
//     len(mockedenvironment.thisCalls())
func (mock *mockRideEnvironment) thisCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockthis.RLock()
	calls = mock.calls.this
	mock.lockthis.RUnlock()
	return calls
}

// timestamp calls timestampFunc.
func (mock *mockRideEnvironment) timestamp() uint64 {
	if mock.timestampFunc == nil {
		panic("mockRideEnvironment.timestampFunc: method is nil but environment.timestamp was just called")
	}
	callInfo := struct {
	}{}
	mock.locktimestamp.Lock()
	mock.calls.timestamp = append(mock.calls.timestamp, callInfo)
	mock.locktimestamp.Unlock()
	return mock.timestampFunc()
}

// timestampCalls gets all the calls that were made to timestamp.
// Check the length with:
//     len(mockedenvironment.timestampCalls())
func (mock *mockRideEnvironment) timestampCalls() []struct {
} {
	var calls []struct {
	}
	mock.locktimestamp.RLock()
	calls = mock.calls.timestamp
	mock.locktimestamp.RUnlock()
	return calls
}

// transaction calls transactionFunc.
func (mock *mockRideEnvironment) transaction() rideObject {
	if mock.transactionFunc == nil {
		panic("mockRideEnvironment.transactionFunc: method is nil but environment.transaction was just called")
	}
	callInfo := struct {
	}{}
	mock.locktransaction.Lock()
	mock.calls.transaction = append(mock.calls.transaction, callInfo)
	mock.locktransaction.Unlock()
	return mock.transactionFunc()
}

// transactionCalls gets all the calls that were made to transaction.
// Check the length with:
//     len(mockedenvironment.transactionCalls())
func (mock *mockRideEnvironment) transactionCalls() []struct {
} {
	var calls []struct {
	}
	mock.locktransaction.RLock()
	calls = mock.calls.transaction
	mock.locktransaction.RUnlock()
	return calls
}

// txID calls txIDFunc.
func (mock *mockRideEnvironment) txID() rideType {
	if mock.txIDFunc == nil {
		panic("mockRideEnvironment.txIDFunc: method is nil but environment.txID was just called")
	}
	callInfo := struct {
	}{}
	mock.locktxID.Lock()
	mock.calls.txID = append(mock.calls.txID, callInfo)
	mock.locktxID.Unlock()
	return mock.txIDFunc()
}

// txIDCalls gets all the calls that were made to txID.
// Check the length with:
//     len(mockedenvironment.txIDCalls())
func (mock *mockRideEnvironment) txIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.locktxID.RLock()
	calls = mock.calls.txID
	mock.locktxID.RUnlock()
	return calls
}

// validateInternalPayments calls validateInternalPaymentsFunc.
func (mock *mockRideEnvironment) validateInternalPayments() bool {
	if mock.validateInternalPaymentsFunc == nil {
		panic("mockRideEnvironment.validateInternalPaymentsFunc: method is nil but environment.validateInternalPayments was just called")
	}
	callInfo := struct {
	}{}
	mock.lockvalidateInternalPayments.Lock()
	mock.calls.validateInternalPayments = append(mock.calls.validateInternalPayments, callInfo)
	mock.lockvalidateInternalPayments.Unlock()
	return mock.validateInternalPaymentsFunc()
}

// validateInternalPaymentsCalls gets all the calls that were made to validateInternalPayments.
// Check the length with:
//     len(mockedenvironment.validateInternalPaymentsCalls())
func (mock *mockRideEnvironment) validateInternalPaymentsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockvalidateInternalPayments.RLock()
	calls = mock.calls.validateInternalPayments
	mock.lockvalidateInternalPayments.RUnlock()
	return calls
}
