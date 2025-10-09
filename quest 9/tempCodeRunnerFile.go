func SortWordArr(a []string) {
	for i := 0 ; i < len(a) ; i++{
		for j := 0 ; j <len(a) ; j++{
			if a[i] < a[j]{
				a[i] , a[j] = a[j] , a[i]
			}
		}
	}
}