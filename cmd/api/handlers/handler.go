package handlers

import (
	authHandler "aramen-api/cmd/api/handlers/auth"
	branchHandler "aramen-api/cmd/api/handlers/branch"
	newsHandler "aramen-api/cmd/api/handlers/news"
	qrcodeHandler "aramen-api/cmd/api/handlers/qrcode"
)

type Handler struct {
	AuthHandler   authHandler.AuthHandler
	BranchHandler branchHandler.BranchHandler
	NewsHandler   newsHandler.NewsHandler
	QrcodeHandler qrcodeHandler.QrcodeHandler
}
