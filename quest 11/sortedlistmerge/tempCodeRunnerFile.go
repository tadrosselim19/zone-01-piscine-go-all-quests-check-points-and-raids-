	if n1 == nil{
		return nil
	}
	flag := true
	for flag{
		flag = false
		current := n1
		for current.Next != nil{
			if current.Data > current.Next.Data {
				current.Data , current.Next.Data = current.Next.Data , current.Data 
				flag = true
			}
			current =current.Next
		}
	}