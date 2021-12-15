package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nill
	}
	return 0, errors.New("Zero length edge not allowed")
}
