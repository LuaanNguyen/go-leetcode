package minimum_length_of_string_after_operation

func minimumLength(s string) int {
    frequencyTable := make(map[rune]int)

    //Count the frequency of each item in the table 
    for _, c := range s {
        frequencyTable[c]++
    }

    //Calculate the number of character to delete
    count := 0
    for _, val := range frequencyTable {
        if val % 2 == 1 {
            count += val - 1
        } else {
            count += val - 2
        }
    }

    return len(s) - count
}