package solutions

import "fmt"

func MinTime(skill []int, mana []int) int64 {
    m := len(skill)
    n := len(mana)
    var res int64
    for i:=1;i<m;i++{
        skill[i] += skill[i-1]
    }
    lastT := timeCal(skill,mana[0])
    
    for i := 1 ; i < n ;i++{
        thisT := timeCal(skill,mana[i])
        thisStep:=lastT[0]
        for j:=1;j<len(thisT);j++{
            tmp := lastT[j]-thisT[j-1]
            if tmp > thisStep{
                thisStep = tmp
            }
        }
        res += int64(thisStep)
        lastT = thisT
    }

    fmt.Println(lastT)
    fmt.Println(res)
    return res + int64(lastT[m-1])
}

func timeCal(skill []int,mana int) []int{
    res := make([]int,len(skill))
    for i:=range skill{
        res[i] = skill[i] * mana
    }
    return res
}