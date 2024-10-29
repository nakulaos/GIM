package dig

import (
	"GIM/apps/apis/upload/internal/service"
)

func provideUpload() {
	Provide(service.NewUploadService)
}
