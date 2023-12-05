package gpio

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	gpio = "/sys/class/gpio"
	open []int
)

func WriteFile(file string, args string) error {
	if len(args) < 1 {
		return fmt.Errorf("not enough arguments")
	}
	err := os.WriteFile(file, []byte(args), 0777)
	if os.IsPermission(err) {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(file string)([]byte, error) {
	mem, err := os.OpenFile(file, os.O_RDWR, os.ModePerm)
	if err != nil {
		return []byte{}, err
	}
	defer mem.Close()
	r := bufio.NewReader(mem)
	buf := make([]byte, r.Size())
	r.Read(buf)
	return buf, nil
}

func Open(pin int) error {
	err := WriteFile(gpio + "/export", fmt.Sprintf("%d", pin))
	if err != nil && !strings.HasSuffix(err.Error(), "device or resource busy") {
		fmt.Println("Can't open pin", pin, "err:\n", err.Error())
		return err
	}
	open = append(open, pin)
	return nil
}

func Close() error {
	for _, pin := range open {
		err := WriteFile(gpio + "/unexport", fmt.Sprintf("%d", pin))
		if err != nil {
			fmt.Println("Can't unexport pin", pin, "err:\n", err.Error())
			return err
		}
	}
	return nil
}

func High(pin int) error {
	err := WriteFile(gpio + "/gpio" + strconv.Itoa(pin) + "/value", "1")
	if err != nil {
		fmt.Println("Can't set pin", pin, "err:\n", err.Error())
		return err
	}
	return nil
}

func Low(pin int) error {
	err := WriteFile(gpio + "/gpio" + strconv.Itoa(pin) + "/value", "0")
	if err != nil {
		fmt.Println("Can't set pin", pin, "err:\n", err.Error())
		return err
	}
	return nil
}

func Write(pin int, value string) error {
	err := WriteFile(gpio + "/gpio" + fmt.Sprintf("%d", pin) + "/value", value)
	if err != nil {
		fmt.Println("Can't set pin", pin, "err:\n", err.Error())
		return err
	}
	return nil
}

func Read(pin int) (string, error) {
	b, err := ReadFile(gpio + "/gpio" + strconv.Itoa(pin) + "/value")
	if err != nil {
		fmt.Println("Can't read pin", pin, "err:\n", err.Error())
		return "", err
	}
	return string(b), nil
}