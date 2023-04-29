package ticket

import (
	"bytes"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"strings"
	"ticket-backend-gh/contracts/request"
	"ticket-backend-gh/contracts/response"
	"ticket-backend-gh/services/awsService"
	"time"

	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
)

type TicketServiceImpl struct {
}

func (s TicketServiceImpl) GenerateTicket(req request.RequestTicket) (response.Data, error) {

	t, _ := time.Parse("2006-01-02", req.DateBought)

	awsService := awsService.AWSServiceImpl{}

	var objectKey = fmt.Sprintf("TICKET_%v_%s_%v.png",
		strings.ToUpper(t.Format("02012006")), req.Name, time.Now().Nanosecond())

	imgTemplate, err := gg.LoadImage("assets/Tiket GH Yovie.png")
	if err != nil {
		return response.Data{}, err
	}

	// imgWidth := image.Bounds().Dx()
	// imgHeight := image.Bounds().Dy()
	dc := gg.NewContextForImage(imgTemplate)
	if err := dc.LoadFontFace("assets/ARIAL.TTF", 18); err != nil {
		return response.Data{}, err
	}
	dc.DrawImage(imgTemplate, 0, 0)

	qrByte, err := qrcode.Encode(req.TicketNumber, qrcode.Medium, 256)
	if err != nil {
		return response.Data{}, err
	}

	img, err := png.Decode(bytes.NewReader(qrByte))
	if err != nil {
		return response.Data{}, err
	}

	var buf bytes.Buffer

	dc.SetColor(color.Black)
	dc.DrawStringWrapped(req.Name, 370, 750, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(req.Nik, 370, 820, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(t.Format("02 January 2006"), 370, 890, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(req.InvoiceCode, 370, 960, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("Wetalk Chatbot", 370, 1035, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(req.SeatRow, 370, 1105, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(req.TicketNumber, 370, 1175, 0.5, 0.5, 500, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped(req.Category, 730, 700, 0.5, 0.5, 500, 1.5, gg.AlignCenter)
	dc.DrawImage(img, 600, 725)
	dc.EncodePNG(&buf)

	url, err := awsService.UploadFileToS3("wetalk-webservices", objectKey, buf.Bytes())
	if err != nil {
		log.Print(err)
		return response.Data{}, err
	}

	return response.Data{
		TicketUrl: url,
	}, nil
}
