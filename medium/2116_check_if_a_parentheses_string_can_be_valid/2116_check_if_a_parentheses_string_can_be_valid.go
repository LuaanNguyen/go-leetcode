package check_if_a_parentheses_string_can_be_valid

func canBeValid(s string, locked string) bool {
    if len(s) % 2 == 1 {
        return false 
    }

    openBrackets := make([]int, 0)
    unlocked := make([]int, 0)

    for i, _ := range s {
        if locked[i] == '0' {
            unlocked = append(unlocked, i)
        } else if s[i] == '(' {
            openBrackets = append(openBrackets, i)
        } else if s[i] == ')' {
            if len(openBrackets) > 0 {
                openBrackets = openBrackets[:len(openBrackets) - 1]
            } else if len(unlocked) > 0 {
                unlocked = unlocked[:len(unlocked) - 1]
            } else {
                return false
            }
        }
    }

    for {
        if len(openBrackets) > 0 && len(unlocked) > 0 && openBrackets[len(openBrackets) - 1] < unlocked[len(unlocked) - 1] {
            openBrackets = openBrackets[:len(openBrackets) - 1]
            unlocked = unlocked[:len(unlocked) - 1]
        } else {
            break
        }
    }

    return len(openBrackets) == 0 
    
}