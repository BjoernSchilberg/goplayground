package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/hybridgroup/mjpeg"
	"gocv.io/x/gocv"
)

var (
	deviceID int
	err      error
	webcam   *gocv.VideoCapture
	stream   *mjpeg.Stream
)

func faceDetect() {
	// https://github.com/opencv/opencv/blob/master/data/haarcascades/haarcascade_frontalface_alt.xml
	xmlFile := "assets/haarcascade_frontalface_alt.xml"

	img := gocv.NewMat()
	defer img.Close()

	blue := color.RGBA{0, 0, 255, 0}

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %s\n", xmlFile)
		return
	}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device skippped: %v\n", deviceID)
		}
		if img.Empty() {
			continue
		}

		rects := classifier.DetectMultiScale(img)
		fmt.Printf("%v faces found\n", len(rects))

		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)
			text := fmt.Sprintf("Human %s", r)
			size := gocv.GetTextSize(text, gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, text, pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		buf, _ := gocv.IMEncode(".jpeg", img)
		stream.UpdateJPEG(buf)
	}

}
func main() {
	deviceID = 0
	host := "0.0.0.0:8083"

	webcam, err = gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return

	}

	defer webcam.Close()

	width := 320
	height := float64(width) * 0.75

	webcam.Set(3, float64(width))
	webcam.Set(4, float64(height))

	//webcam.Set(5, 5) // fps

	stream = mjpeg.NewStream()

	go faceDetect()

	fmt.Println("Capturing, point your browser to", host)

	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(host, nil))

}
