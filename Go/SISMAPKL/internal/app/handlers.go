package app

import "sismapkl/internal/handlers/http/status_check_handler"

type Handlers struct {
	HandleStatus status_check_handler.Contract
}

func RegisterHandlers(s Services) Handlers {
	return Handlers{

		// inisiasi handler
		HandleStatus: status_check_handler.New(),
	}
}
