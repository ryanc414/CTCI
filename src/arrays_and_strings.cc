/**
 * arrays_and_strings.cc
 *
 * Answers to chapter one of CTCI.
 */

#include <cassert>
#include <algorithm>
#include <iostream>
#include <string>
#include <map>

bool is_unique(const std::string input);
bool is_unique_no_data_structs(const std::string input);

/* Check if a string has all unique characters. */
bool is_unique(const std::string input) {
    std::map<char, bool> chars_found;

    for (char c : input) {
        if (chars_found[c]) {
            return false;
        }

        chars_found[c] = true;
    }

    return true;
}

/**
 * Check if a string has all unique characters. No additional data structures
 * used.
 */
bool is_unique_no_data_structs(const std::string input) {
    // Make a mutable local copy of the input string and sort it.
    std::string mut_str = input;
    std::sort(mut_str.begin(), mut_str.end());

    // Now loop through and check each character to its previous one. Since
    // the string characters are sorted, any duplicate characters will be
    // next to each other.
    for (size_t i = 1; i < mut_str.length(); ++i) {
        if (mut_str[i] == mut_str[i - 1]) {
            return false;
        }
    }

    return true;
}

int main() {
    assert(is_unique("abcdefg"));
    assert(!is_unique("abcdefa"));
    assert(is_unique(""));

    assert(is_unique_no_data_structs("abcdefg"));
    assert(!is_unique_no_data_structs("abcdefa"));
    assert(is_unique_no_data_structs(""));

    std::cout << "All assertions passed." << std::endl;

    return 0;
}

