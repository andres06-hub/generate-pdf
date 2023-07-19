package main

import(
  "github.com/johnfercher/maroto/pkg/pdf"
  "github.com/johnfercher/maroto/pkg/consts"
  "github.com/andres06-hub/generate-pdf/constants"
)

func main() {
  moroto := pdf.NewMaroto(consts.Portrait, consts.A4)
  // Margin
  moroto.SetPageMargins(20, 10, 20)
  // Head
  buildHeading(moroto)
  // contend
}

func buildHeading(m pdf.Maroto) {
  m.RegisterHeader(func (){
    m.Row(30, func() {
      m.Col()
    })
  })
}
