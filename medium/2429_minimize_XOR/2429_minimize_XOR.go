package minimize_XOR

import "math/bits"

func minimizeXor(num1, num2 int) int {
	res := num1 

	targetSetBitsCount := bits.OnesCount(uint(num2))
	setBitsCount := bits.OnesCount(uint(num1))

	currentBit := 0 

	for setBitsCount < targetSetBitsCount {
		if !checkBitSet(res, currentBit) {
			res = setBit(res, currentBit)
			setBitsCount++
		}
		currentBit++
	}

	currentBit = 0 

	for setBitsCount > targetSetBitsCount {
		if checkBitSet(res, currentBit) {
			res = unsetBit(res, currentBit)
			setBitsCount--
		}
		currentBit++
	}
	
	return res 
}

func checkBitSet(num int, i int ) bool {
	return (num & (1 << i)) != 0
}

func setBit(num int, i int) int {
	return num | (1 << i)
}

func unsetBit(num int, i int) int {
	return num & ^(1 << i)
}