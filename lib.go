package go_usr_tcp232_485

import (
	"errors"
)

const (
	ERR_VAL_DATA_BITS = "Invalid value data bits"
	ERR_VAL_STOP_BITS = "Invalid value stop bits"
)

func GenPackSettings(band_rate, data_bits, stop_bit int, parity string) ([]byte, error) {

	if data_bits < 5 || data_bits > 8 {
		return nil, errors.New(ERR_VAL_DATA_BITS)
	}

	if stop_bit < 1 || stop_bit > 2 {
		return nil, errors.New(ERR_VAL_STOP_BITS)
	}

	// Head
	res := []byte{0x55, 0xAA, 0x55}

	// Band rate
	for i := 0; i < 3; i++ {
		moved := uint8(16 - i*8)
		res = append(res, byte(band_rate>>moved&0xFF))
	}

	// UART settings
	uart := data_bits - 5
	uart += stop_bit << 1

	// Parity
	if parity != "" && parity != "N" {
		uart += (1 << 2)
		switch parity {
		case "E":
			uart += 1 << 3
		case "M":
			uart += 2 << 3
		case "C":
			uart += 3 << 3
		}
	}

	res = append(res, byte(uart))
	res = append(res, crc(res[:len(res)-4]))
	return res, nil
}

func crc(body []byte) byte {
	sum := 0
	for i := 0; i < len(body); i++ {
		sum = sum + int(body[i])
	}
	return byte(sum)
}
