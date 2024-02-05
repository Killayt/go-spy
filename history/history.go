package history

import (
	"fmt"
	"io/ioutil"
)

func GetHistory() error {
	content, err := ioutil.ReadFile("C:/Users/maxre/AppData/Local/Google/Chrome/User Data/Default/History")
	if err != nil {
		return err
	}
	fmt.Println(string(content))
	return nil
}
