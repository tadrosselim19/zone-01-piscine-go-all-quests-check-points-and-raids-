func TrimAtoi(s string) int {
	nbyte := []rune(s)
	result :=0 
	for i :=0 ; i < len(nbyte) ; i++{
		if (nbyte[i] > '9' || nbyte[i] < '0' ) && nbyte[i] != '-'{
			continue 
		}else{
			result = result*10 + int(nbyte[i]-'0')
		}
	}
	return result
}