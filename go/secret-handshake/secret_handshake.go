package secret

/*
1 = wink
10 = double blink
100 = close your eyes
1000 = jump


10000 = Reverse the order of the operations in the secret handshake.
*/
func Handshake(n int) (s []string) {
	if n <= 0 {
		return []string{}
	}
	if n&0x01 == 0x01 {
		s = append(s, "wink")
	}

	if n&0x02 == 0x02 {
		s = append(s, "double blink")
	}

	if n&0x04 == 0x04 {
		s = append(s, "close your eyes")
	}

	if n&0x08 == 0x08 {
		s = append(s, "jump")
	}

	if n&0x10 == 0x10 {
		s = reverse(s)
	}

	return
}

func reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}
