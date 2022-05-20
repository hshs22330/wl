package main

func RemoveString(sl []interface{}) (aInt []int) {
	aInt = []int{}

	for _, v := range sl {
		if i, ok := v.(int); ok {
			aInt = append(aInt, i)
		}
	}

	return
}

func TransferArray(sl []interface{}) (aTran []interface{}) {
	aTran = []interface{}{}
	if len(sl) <= 1 {
		return
	}

	n, ok := sl[0].(int)
	if !ok {
		return
	}

	if len(sl) > n {
		aTran = sl[1 : n+1]
	} else {
		aTran = sl[1:]
	}

	for i, j := 1, len(aTran)-2; i < j; i, j = i+1, j-1 {
		aTran[i], aTran[j] = aTran[j], aTran[i]
	}

	return
}
