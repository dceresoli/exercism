package secret

const testVersion = 1

func reverse(r []string) []string {
	var s = make([]string, cap(r))
	for i := 0; i < len(r); i++ {
		s[len(r)-1-i] = r[i]
	}
	return s
}

func Handshake(i uint) []string {
	var response = make([]string, 0)

	if i&1 != 0 {
		response = append(response, "wink")
	}

	if i&2 != 0 {
		response = append(response, "double blink")
	}

	if i&4 != 0 {
		response = append(response, "close your eyes")
	}

	if i&8 != 0 {
		response = append(response, "jump")
	}

	if len(response) == 0 {
		return nil
	}

	if i&16 == 0 {
		return response
	}

	// if bit is set, reverse the slice
	return reverse(response)
}
