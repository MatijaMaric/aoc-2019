package utils

// SafeClose attempts to close a channel
func SafeClose(ch chan int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
		}
	}()
	close(ch)
	return true
}
