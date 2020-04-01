package main

import (
  "fmt"
  "os/exec"
  "github.com/jung-kurt/gofpdf"
  // "image/png"
)

func main () {
  // GenerateTRPpdf("Abc", "", "12345")
  _, err := exec.Command("convert", "TRP872343T.png", "AFF872343T.png", "harhar.pdf").Output()
  if err != nil{
    fmt.Println("err====>", err)
  }
  fmt.Println()
}

func GenerateTRPpdf(TRPname string, quoteId string, vin string) string {
	fmt.Println("TRPname==>", TRPname)
	fmt.Println("vin==>", vin)
	pdf := gofpdf.New("P", "mm", "Letter", "")
	pdf.SetMargins(0, 0, 0)
	pdf.AddPage()
  // folderpath := "/Users/aniket/Desktop/gofpdf"
	path := "/Users/aniket/Desktop/gofpdf/TRP872343T.png"

	pdf.Image(path, 10, 10, 196, 259, false, "", 0, "")
	TRPPdf := "/Users/aniket/Desktop/gofpdf/new.pfd"
	// name := PublicPath + TRPPdf
	// fmt.Println("name-->",name)
	// err := pdf.OutputFileAndClose(name)
	// pdf.AddPage()
	// path = "/Users/aniket/Desktop/gofpdf/AFF872343T.jpeg"
	// pdf.Image(path, 10, 10, 196, 259, false, "", 0, "")
	// name := quoteId + "-AFF.pdf"
	err := pdf.OutputFileAndClose(TRPPdf)
	if err != nil {
		fmt.Println(">>>", err)
		return ""
	} else {
		_, errGS := exec.Command("gs", "-dBATCH", "-dNOPAUSE", "-sDEVICE=pngmono", "-o", path+"-%03d.png", "-r600", TRPPdf).Output()
		if errGS != nil {
			fmt.Println("Error on coverting grayscal:", errGS)
		}
		// if quoteId != "" {
		// 	quoteJSON, err := db.HGet(QuotesHash, quoteId)
		// 	if err == nil {
		// 		var quoteData structs.Quote
		// 		err = json.Unmarshal([]byte(quoteJSON), &quoteData)
		// 		if err == nil {
		// 			quoteData.TRPPdf = TRPPdf
		// 			newJSON, err := json.Marshal(quoteData)
		// 			if err == nil {
		// 				db.HSet(QuotesHash, quoteId, string(newJSON))
		// 			}
		// 		}
		// 	}
		// }
		// SplitTRPPdf(vin, name)
		return TRPPdf
	}
}
// gs \
//  -sDEVICE=pdfwrite \
//  -o foo.pdf \
//   /usr/local/share/ghostscript/9.50/lib/viewjpeg.ps \
//  -c \(Users/aniket/Desktop/gofpdf/TRP872343T.png\) viewJPEG

// gs \
//  -sDEVICE=pdfwrite \
//  -o foo.pdf \
//   /usr/local/share/ghostscript/9.50/lib/viewjpeg.ps \
//  -c "(1st.jpg)  viewJPEG showpage \
//      (2nd.jpg)  viewJPEG showpage \
//      (3rd.jpg)  viewJPEG showpage \
//      (last.jpg) viewJPEG showpage"

// gs \
//  -sDEVICE=pdfwrite \
//  -o foo.pdf \
//   /usr/local/share/ghostscript/9.50/lib/viewpng.ps \
//  -c "(Users/aniket/Desktop/gofpdf/TRP872343T.png\) viewPNG"
