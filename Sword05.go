package main

func replaceSpace(s string) string {
	// -1 代表替换所有匹配的
	//return strings.Replace(s," ","%20",-1)
	var str string = ""
	for _, v := range s{
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}
