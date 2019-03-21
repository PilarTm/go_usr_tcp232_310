package go_usr_tcp232_485

import (
	"fmt"
	"testing"
)

func TestCreateEmptyConverter(t *testing.T) {

	res, _ := GenPackSettings(115200, 8, 1, "N")
	fmt.Printf("%02X ===", res)

}
