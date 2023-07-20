package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	cnts "github.com/andres06-hub/generate-pdf/constant"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	maroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	// Margin
	maroto.SetPageMargins(15, 10, 15)
	maroto.SetDefaultFontFamily(consts.Arial)

	// Head
	buildHeading(maroto)
	// space
	addSpace(maroto, 10)
	// contend
	buildContent(maroto)
	// space
	addSpace(maroto, 10)
	// tiers
	buildTiers(maroto)
	// space
	addSpace(maroto, 10)
	// Footer
	buildFooter(maroto)

	// Set Date
	timeNow := time.Now()
	maroto.SetCreationDate(timeNow)

	// Output file
	err := maroto.OutputFileAndClose("pdfs/example.pdf")
	if err != nil {
		fmt.Println("⚠️ Could not save PDF:", err)
		os.Exit(1)
	}
}

func fetchImage(url string) string {
	res, resErr := http.Get(url)
	if resErr != nil {
		return ""
	}

	body := res.Body
	bytes, bytesErr := ioutil.ReadAll(body)
	if bytesErr != nil {
		return ""
	}
	base64Image := base64.StdEncoding.EncodeToString(bytes)

	defer res.Body.Close()

	return base64Image
}

func buildHeading(m pdf.Maroto) {
	base64Image := fetchImage("https://quickquote.beyondfloods.com/static/media/BF_Logo_with_Text.d819e341.png")
	timeNow := time.Now()

	m.SetCreationDate(timeNow)
	m.Row(30, func() {
		// TODO: Generate margin
		m.SetBorder(true)
		m.Col(cnts.HALF_WIDTH, func() {
			// m.QrCode("https://github.com/johnfercher/maroto", props.Rect{
			// 	Percent: 50,
			// })
			m.Text(fmt.Sprintf("Quote & Property Flood Outlook Report"), props.Text{
				Style: consts.Bold,
				Top:   10,
				Align: consts.Left,
			})
			m.Text(fmt.Sprintf("Quote Reference: %d       Date: %s", 15475, "19/07/2023"), props.Text{
				Style: consts.Bold,
				Top:   15,
				Align: consts.Left,
			})
			m.Text(fmt.Sprintf("Address: %s", "cra 44 # 26-80"), props.Text{
				Style: consts.Bold,
				Top:   20,
				Align: consts.Left,
			})
		})
		m.Col(cnts.HALF_WIDTH, func() {
			err := m.Base64Image(base64Image, consts.Png, props.Rect{
				Center: true,
			})
			if err != nil {
				return
			}
			// m.Text("Prepared for you by the Div Rhino Fruit Company")
		})
	})
}

func buildContent(m pdf.Maroto) {
	m.Row(50, func() {
		m.Col((8), func() {
			// buildItem(m)
		})
		m.Col(cnts.THIRD_WIDTH, func() {
			m.Text("EXAMPLE 2", props.Text{})
		})
	})
}

func buildTiers(m pdf.Maroto) {
	x := []string{"header1", "header2", "header3"}
	y := [][]string{
		{"content1.1", "content1.2", "content1.3"},
		{"content2.1", "content2.2", "content2.3"},
	}

	m.TableList(x, y, props.TableList{
		VerticalContentPadding: 1,
		HeaderProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
		ContentProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
	})

	m.TableList(x, y, props.TableList{
		VerticalContentPadding: 1,
		HeaderProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
		ContentProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
	})

	m.TableList(x, y, props.TableList{
		VerticalContentPadding: 1,
		HeaderProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
		ContentProp: props.TableListContent{
			GridSizes: []uint{1, 1, 2},
		},
	})

	// m.SetBorder(true)
	// m.Row(100, func() {
	// m.Col(cnts.THIRD_WIDTH, func() {
	// })
	// m.Col(cnts.THIRD_WIDTH, func() {})
	// m.Col(cnts.THIRD_WIDTH, func() {})
	// })
}

func buildItem(m pdf.Maroto) {
	// Fetch image
	base64Image := fetchImage("https://cdn-icons-png.flaticon.com/512/395/395841.png")
	m.Row(10, func() {
		m.Col(cnts.SINGLE_WIDTH, func() {
			m.Base64Image(base64Image, consts.Png, props.Rect{
				Center: true,
			})
		})
		m.Col((cnts.HALF_WIDTH + cnts.SINGLE_WIDTH), func() {
			m.Text("This property is estimated to have a Moderate Risk from flooding.", props.Text{
				Size: 5,
			})
		})
		m.Col(cnts.QUARTER_WIDTH, func() {})
	})
}

func addSpace(m pdf.Maroto, height float64) {
	m.Row(height, func() {
		m.ColSpace(12)
	})
}

func buildFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(30, func() {
			m.Col(6, func() {
				m.Text("Example Example Example Example Example Example Example Example Example Example Example Example Example Example Example Example Example Example Example")
			})
			m.Col(6, func() {
				m.QrCode("https://www.cuemby.com/", props.Rect{
					Center: true,
				})
			})
		})
	})
}
