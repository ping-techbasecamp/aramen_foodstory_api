package qrcodeHandler

import (
	qrcodeModel "aramen-api/cmd/api/model/qrcode"
	"aramen-api/pkg/errs"
	"aramen-api/srv/aes"
	"aramen-api/srv/foodstory"
	qrcodeService "aramen-api/srv/services/qrcode"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type QrcodeHandlerInterface interface {
	ScanQr(c *gin.Context)
}

type QrcodeHandler struct {
	qrcodeService qrcodeService.QrcodeService
}

func NewQrcodeHandler(qrcodeService qrcodeService.QrcodeService) QrcodeHandler {
	return QrcodeHandler{qrcodeService: qrcodeService}
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func (h QrcodeHandler) ScanQr(c *gin.Context) {
	var req qrcodeModel.QrCode
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrorStructResponse(err))
		return
	}
	// id, _ := c.Get("id")
	// userId := any.ParseInt(id)
	// point, star, err := h.qrcodeService.ScanQr(userId, req.Qrcode)
	// if err == nil && point != nil && star != nil {
	// 	c.JSON(200, map[string]interface{}{
	// 		"point": *point,
	// 		"star":  *star,
	// 	})
	// }
	
	decryptedCode, err := aes.DecryptECB(os.Getenv("FOODSTORY_AES_KEY"), req.Qrcode)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("decryptedCode")
	fmt.Println(decryptedCode)

	foodstoryQrcode, err := qrcodeModel.UnmarshalQrcode([]byte(decryptedCode))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	foodstoryOrder, err := foodstory.ScanQrcodeFoodstory(foodstoryQrcode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, foodstoryOrder)
}
