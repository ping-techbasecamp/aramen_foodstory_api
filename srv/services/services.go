package services

import (
	authService "aramen-api/srv/services/auth"
	branchService "aramen-api/srv/services/branch"
	newsService "aramen-api/srv/services/news"
	qrcodeService "aramen-api/srv/services/qrcode"
)

type Service struct {
	AuthService   authService.AuthService
	BranchService branchService.BranchService
	NewsService   newsService.NewsService
	QrcodeService qrcodeService.QrcodeService
}
