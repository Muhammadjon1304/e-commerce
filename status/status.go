package status

var (
	NoError                       = 0
	ErrorCodeValidation           = -10
	ErrorCodeValidationDateFormat = -12
	ErrorCodeDB                   = -30
	ErrorCodeRemoteDBO            = -40
	ErrorCodeRemoteESB            = -45
	ErrorCodeRemoteCRM            = -50
	ErrorCodeRemoteOther          = -55
	ErrorCodeRMQ                  = -60
	ErrorAuth                     = -65
	ErrorGenerateExcel            = -70
	ErrorCacheData                = -75
	ErrorImage                    = -80
)

var (
	Success = "Success"
	Failure = "Failure"
)
