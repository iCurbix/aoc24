package util

func isDigit(chr uint8) bool {
	if chr >= '0' && chr <= '9' {
		return true
	}
	return false
}

func ScanNumbers(data []byte, atEOF bool) (advance int, token []byte, err error) {
	start := 0
	for ; start < len(data); start += 1 {
		if isDigit(data[start]) {
			break
		}
	}
	for i := start + 1; i < len(data); i += 1 {
		if !isDigit(data[i]) {
			return i, data[start:i], nil
		}
	}
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	return start, nil, nil
}
