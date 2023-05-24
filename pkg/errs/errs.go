package errs

var ErrFmt = "err: [%v]"
var ErrFmtMore = "err: [%v - %v]"

// common
var (
	ErrLoadingTimeZone       = _new("CMN000", "error loading timezone data")
	ErrMarshalingJson        = _new("CMN001", "error marshaling json")
	ErrUnmarshalingJson      = _new("CMN002", "error unmarshaling json")
	ErrParsingTime           = _new("CMN003", "error parsing time")
	ErrNoEntityIdProvided    = _new("CMN004", "entity ID is required but none was provided")
	ErrNoDateProvided        = _new("CMN005", "error no date provided")
	ErrNoPayloadData         = _new("CMN006", "error event contains no payload data")
	ErrRepoMockAction        = _new("CMN007", "error repo mock action")
	ErrUnknownErrorType      = _new("CMN008", "error unknown error type")
	ErrInvalidDate           = _new("CMN009", "error invalid date format")
	ErrConvertingStringToInt = _new("CMN010", "error converting string to int")
	ErrGettingParam          = _new("CMN011", "error getting param")
	ErrConvertingPayload     = _new("CMN012", "error converting payload")
	ErrInvalidEmail          = _new("CMN013", "error invalid email")
	ErrOpenFile              = _new("CMN014", "error open file")
	ErrReadingFile           = _new("CMN015", "error reading file")
	ErrActionNotImplemented  = _new("CMN016", "error action is not implemented")
	ErrParsingAtoi           = _new("CMN017", "error converting string to a int (Atoi)")
)

// pkg/api
var (
	ErrResponseWriter = _new("API000", "error writing to response writer")
)

// pkg/config
var (
	ErrCreatingParamStore    = _new("CFG002", "unable to create param store service")
	ErrUnknownConfigProvider = _new("CFG003", "error unknown config provider")
	ErrGettingEnv            = _new("CFG004", "error unknown get env variables")

	ErrGettingEnvNatsURL  = _new("CFG005", "error unknown get env variable Nats URL")
	ErrGettingEnvNatsPort = _new("CFG006", "error unknown get env variable Nats Port")

	ErrGettingEnvWebSocketURL  = _new("CFG007", "error unknown get env variable Web Socket URL")
	ErrGettingEnvWebSocketPort = _new("CFG008", "error unknown get env variable Web Socket Port")
)

// pkg/stream
var (
	ErrNatsConnection = _new("NAT001", "error starting nats connection")
	ErrNatsEmptyTopic = _new("NAT002", "error getting param topic")
)

// pkg/websocket
var (
	ErrWebSocketConnection   = _new("WBS001", "error starting websocket connection")
	ErrWebSocketWriteMessage = _new("WBS002", "error websocket write message")
	ErrWebSocketReadMessage  = _new("WBS003", "error websocket read message")
	ErrWebSocketUpgrader     = _new("WBS004", "error websocket upgrader connection")
)

// validations
var (
	ErrValidation = _new("VAL000", "error on validation")
)
