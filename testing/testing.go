package testing

import (
	"bytes"
	"errors"
	"io"
)

var errDivisaoInvalida = errors.New("divisão invalida")

func divideInteiros(dividendo, divisor int) (quociente int, resto int, err error) {
	if divisor == 0 {
		err = errDivisaoInvalida
		return
	}
	quociente = dividendo / divisor
	resto = dividendo % divisor
	return
}

func sum(a, b int) (ret int) {
	ret = a + b
	return
}

func leitor(r io.Reader) (ret string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r) // nolint
	ret = buf.String()
	return
}
