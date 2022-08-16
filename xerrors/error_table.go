package xerrors

var (
	errorMap map[ErrorType]string
)

func Initialize() error {
	return loadData()
}

// loadData loads data from database and save memcache
func loadData() error {
	errorMap = make(map[ErrorType]string)
	// TODO: should load error table from db or config file and save memcache
	errorMap = map[ErrorType]string{
		Success:              MsgSuccess,
		Unknown:              MsgGeneralError,
		BadRequestErr:        MsgBadRequest,
		AuthenticationFailed: MsgAuthenticateFailed,
		InternalServerError:  MsgGeneralError,
		CallInternalAPIError: MsgGeneralError,
	}
	return nil
}
