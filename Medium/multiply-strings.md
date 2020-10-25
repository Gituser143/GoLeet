Code
====

```go
func reverse(bs []byte) []byte {
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}

	return bs
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	a := len(num1)
	b := len(num2)

	result := make([]byte, a+b)

	i_n1 := 0
	i_n2 := 0

	for i := 0; i < a+b; i++ {
		result[i] = 48
	}

	for i := a - 1; i >= 0; i-- {
		var carry byte = 0
		n1 := num1[i] - 48
		i_n2 = 0

		for j := b - 1; j >= 0; j-- {
			n2 := num2[j] - 48
			sum := int(n1)*int(n2) + int(result[i_n1+i_n2]) - 48 + int(carry)
			carry = byte(sum / 10)
			result[i_n1+i_n2] = byte(48 + (sum % 10))
			i_n2++
		}

		if carry > 0 {
			result[i_n1+i_n2] = byte(int(result[i_n1+i_n2]) + int(carry))
		}
		i_n1++
	}

	solution := string(reverse(result))

	for i := range solution {
		if solution[i] != 48 {
			solution = solution[i:]
			break
		}
	}

	return solution
}
```
