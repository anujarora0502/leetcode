package solutions

func WordPattern(pattern string, s string) bool {
    var stringArr []string

    tempString := ""

    for _, element := range s {
        if string(element) == " " {
            if len(tempString) != 0 {
                stringArr = append(stringArr, tempString)
                tempString = ""
            }
        }else {
            tempString += string(element)
        }
    }

	if len(tempString) != 0 {
		stringArr = append(stringArr, tempString)
	}    

    if len(pattern) != len(stringArr){
        return false
    } 


    stringMapSP := make(map[string]rune)
    stringMapPS := make(map[rune]string)
    for index, _ := range pattern {
        value1, ok1 := stringMapPS[rune(pattern[index])]
        value2, ok2 := stringMapSP[stringArr[index]]
        if ok1 != ok2 {
            return false
        }

        if ok1 && (value1 != stringArr[index] || value2 != rune(pattern[index])) {
            return false
        }

        stringMapPS[rune(pattern[index])] = stringArr[index]
        stringMapSP[stringArr[index]] = rune(pattern[index])
    }

    return true
}