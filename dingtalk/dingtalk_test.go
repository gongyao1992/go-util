package dingtalk

import (
	"strconv"
	"testing"
)

func TestDingtalkSend2(t *testing.T) {
	accessToken := "a071c2eaa8538a4a70a953556e6bc9acc9aa1248b6bcd7534c3e5672fa848fc4"
	secret := "SEC52ba5288b2945901f4986f474210f86e091547585d89a9467cf7675f2c204965"

	r := NewRobitText(accessToken, secret, "15010062048", false)

	s := make([]string, 0)
	i := 0
	for true {
		s = append(s, strconv.Itoa(i))
		i ++
		if i >= 1000 {
			break
		}
	}

	t.Log(r.Send(s))
}
