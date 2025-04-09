# Daily LeetCode Solutions in Go ✏️

This repository contains my solutions to LeetCode problems in [Go](https://go.dev/).

<div align="center">
  <img src="https://raw.githubusercontent.com/gist/brudnak/6c21505423e4ff089ab704ec79b5a096/raw/b2d3dec32474b2121b179920734b259323a7c250/go.gif" alt="Go" width="400"/>
</div>

## Folder Structure 🗂️

- `easy/`: Easy level problems.
- `medium/`: Medium level problems.
- `hard/`: Hard level problems.

## Unit Testing 🧪

In your root direction, run:

```bash
go test ./...
```

If you only want to run tests in a specific problem `(e.g., ./easy/0001_two_sum)`, run:

```bash
go test ./easy/0001_two_sum
```

Clean test cache

```bash
go clean -testcache
```

## Solutions (Continue Updating...)

|                                                              Leetcode ID                                                              | Title & Solution                                                                                                            | Coefficient Of Difficulty |                                             Remarks                                              |                                                                                                        Approach                                                                                                        |
| :-----------------------------------------------------------------------------------------------------------------------------------: | :-------------------------------------------------------------------------------------------------------------------------- | :-----------------------: | :----------------------------------------------------------------------------------------------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------: |
|                          [3042](https://leetcode.com/problems/count-prefix-and-suffix-pairs-i/description/)                           | [Count Prefix and Suffix Pairs I](/easy/3042_count_prefix_and_sufix_pairs_I)                                                |           Easy            |                                       _`Array`_ _`String`_                                       |                                                                                                                                                                                                                        |
|                         [2185](https://leetcode.com/problems/counting-words-with-a-given-prefix/description/)                         | [Counting Words With a Given Prefix](/easy/2185_counting_words_with_a_given_string)                                         |           Easy            |                             _`Array`_ _`String`_ _`String Matching`_                             |                                                                                                                                                                                                                        |
|       [1400](https://leetcode.com/problems/construct-k-palindrome-strings/description/?envType=daily-question&envId=2025-01-11)       | [Construct K Palindrome Strings](/medium/1400_construct_k_palindrome_strings)                                               |          Medium           |                         _`HashTable`_ _`String`_ _`Greedy`_ _`Counting`_                         |                                                                                                                                                                                                                        |
| [2116](https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/description/?envType=daily-question&envId=2025-01-12) | [Check if a Parentheses String Can Be Valid](/medium/2116_check_if_a_parentheses_string_can_be_valid)                       |          Medium           |                                 _`Stack`_ _`String`_ _`Greedy`_                                  |                                                                                                                                                                                                                        |
| [3223](https://leetcode.com/problems/minimum-length-of-string-after-operations/description/?envType=daily-question&envId=2025-01-13)  | [Minimum Length of String After Operations](/medium/2116_check_if_a_parentheses_string_can_be_valid)                        |          Medium           |                              _`HashTable`_ _`String`_ _`Counting`_                               |                                                                                                                                                                                                                        |
| [2657](https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/?envType=daily-question&envId=2025-01-14) | [Find the Prefix Common Array of Two Arrays](/medium/2657_find_the_prefix_common_array_of_two_arrays)                       |          Medium           |                                     _`HashTable`_ _`Array`_                                      |                                                                                                                                                                                                                        |
|                [2429](https://leetcode.com/problems/minimize-xor/description/?envType=daily-question&envId=2025-01-15)                | [Minimize XOR](/medium/2429_minimize_XOR)                                                                                   |          Medium           |                                       _`Bit Manipulation`_                                       |                                                                                                                                                                                                                        |
|                              [2425](https://leetcode.com/problems/neighboring-bitwise-xor/description/)                               | [Bitwise XOR of All Pairings](/medium/2429_minimize_XOR)                                                                    |          Medium           |                              _`Bit Manipulation`_ _`Brain Teaser`_                               |                                                                                                                                                                                                                        |
|        [2683](https://leetcode.com/problems/bitwise-xor-of-all-pairings/description/?envType=daily-question&envId=2025-01-16)         | [Neighboring Bitwise XOR](/medium/2683_neighboring_bitwise_XOR)                                                             |          Medium           |                                  _`Bit Manipulation`_ _`Array`_                                  |                                                                                                                                                                                                                        |
| [1368](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/?envType=daily-question&envId=2025-01-18) | [Minimum Cost to Make at Least One Valid Path in a Grid](/hard/1368_minimum_cost_to_make_at_least_one_valid_path_in_a_grid) |           Hard            | _`Breadth First Search`_ _`Graph`_ `Heap(Priority Queue)` _`Matrix`_ _`Shortest Path`_ _`Array`_ |                                                                                                                                                                                                                        |
|                 [407](https://leetcode.com/problems/trapping-rain-water-ii/?envType=daily-question&envId=2025-01-19)                  | [Trapping Rain Water II](/hard/1368_minimum_cost_to_make_at_least_one_valid_path_in_a_grid)                                 |           Hard            | _`Breadth First Search`_ _`Graph`_ `Heap(Priority Queue)` _`Matrix`_ _`Shortest Path`_ _`Array`_ |                                                                                                                                                                                                                        |
|   [2661](https://leetcode.com/problems/first-completely-painted-row-or-column/description/?envType=daily-question&envId=2025-01-20)   | [First Completely Painter Row or Column](/medium/2661_first_completely_painted_row_or_column)                               |          medium           |                                _`Matrix`_ _`HashTable`_ _`Array`_                                |                                                                                                                                                                                                                        |
|                                           [2017](https://leetcode.com/problems/grid-game/)                                            | [Grid Game](/medium/2017_grid_game)                                                                                         |          medium           |                                    _`Matrix`_ _`Prefix Sum`_                                     |                                                                                                                                                                                                                        |
|            [1765](https://leetcode.com/problems/map-of-highest-peak/description/?envType=daily-question&envId=2025-01-22)             | [Map of Highest Peak](/medium/1765_map_of_highest_peak)                                                                     |          medium           |                                _`Graph`_ _`Breadth-First Search`_                                |                                                                                                                                                                                                                        |
|       [1267](https://leetcode.com/problems/count-servers-that-communicate/description/?envType=daily-question&envId=2025-01-23)       | [Count Servers that Communicate](/medium/1267_count_servers_that_communicate)                                               |          medium           |                                      _`Counting`_ _`Array`_                                      |                                                                                                                                                                                                                        |
|          [802](https://leetcode.com/problems/find-eventual-safe-states/description/?envType=daily-question&envId=2025-01-24)          | [Find Eventual Safe States](/medium/802_find_eventual_safe_states)                                                          |          medium           |                            _`Depth-First Search`_ _`Adjacency List`_                             |                                                                                                                                                                                                                        |
|                                           [75](https://leetcode.com/problems/sort-colors/)                                            | [Sort Colors](/medium/75_sort_colors)                                                                                       |          medium           |                              _`Three Pointers`_ _`DNF Algortithm`_                               |                                    Initalize 3 pointers. Left and mid at 0, high at len(nums) - 1. Do a while loop as long as mid <= high and swap elements based on 3 conditions.                                     |
|                                 [229](https://leetcode.com/problems/majority-element-ii/description/)                                 | [Sort Colors](/medium/229_majority_element_II)                                                                              |          medium           |                                     _`Counter`_ _`Hashmap`_                                      | There can be at most 2 such elements with floor(n /3). Do a first pass to find all and eliminate all the non-qualify elements, then the second pass check how many time a each candidate appear in the original array. |
|                          [167](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/)                           | [Two Sum II - Input Array Is Sorted](/medium/167_two_sum_II_input_array_is_sorted)                                          |          medium           |                                         _`Two Pointers`_                                         |                       We know that the array is sorted in asc order. We can track the total sum of 2 pointers left and right and increase or decrease their indices accordingly in a while loop                        |
|                                         [15](https://leetcode.com/problems/3sum/description/)                                         | [3Sum](/medium/15_3sum)                                                                                                     |          medium           |                                         _`Two Pointers`_                                         |                            Interate through the array, initialize 2 other pointers j = i + 1 and k = n - 1 and check the sum and adjust j and k accordingly. Make sure to check duplicates                             |
|                 [3375](https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/description)                  | [3Sum](/easy/3375_minimum_operations_to_make_array_values_equal_to_k)                                                       |          medium           |                                           _`Hashmap`_                                            |                                   There is no solution if a value that smaller than k. If a value is larger than k, we add it to our set and count the length of the set at the end                                    |

## Workflow 🌊

1. `go.yml`: This workflow runs whenver new code is pushed to main or create a pull request

- Runs all Go tests
- Checks code formatting with `go fmt`
- Runs `go vet` for static analysis
- Runs `staticcheck` for additional code quality checks

2. `update-stats.yml`: This runs whenever you push changes to your solution directory (`easy/medium/hard`). It:

- Counts the number of solutions in each level
- Update `README.md` with current stats
- Automaticall commits and pushes the changes



## Statistics 📊

- Easy: 5 solutions
- Medium: 16 solutions
- Hard: 2 solutions
- Total: 23 solutions
## License 🪪

```txt
MIT License
Copyright (c) 2025 Luan Nguyen
```
