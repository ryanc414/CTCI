/**
 * arrays_and_strings.cc
 *
 * Answers to chapter one of CTCI.
 */

#include <cassert>
#include <cstring>
#include <algorithm>
#include <iostream>
#include <string>
#include <map>

bool is_unique(const std::string input);
bool is_unique_no_data_structs(const std::string input);
bool is_permutation(const std::string str_a, const std::string str_b);
std::map<char, int> letter_freq(const std::string str);
void urlify(char *str, size_t buf_len, size_t true_len);

void test_is_unique();
void test_is_permutation();
void test_urlify();

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

/* Check if two strings are permuations of each other.*/
bool is_permutation(const std::string str_a, const std::string str_b) {
    if (str_a.length() != str_b.length()) {
        return false;
    }

    return letter_freq(str_a) == letter_freq(str_b);
}

/* Return a mapping of letter to its frequency in a word. */
std::map<char, int> letter_freq(const std::string str) {
    std::map<char, int> freq;

    for (char c : str) {
        ++freq[c];
    }

    return freq;
}

/**
 * Replace all spaces in a string with "%20". The string buffer is modified
 * in-place.
 */
void urlify(char *str, size_t buf_len, size_t true_len) {
    assert(buf_len >= true_len);
    size_t new_index = buf_len - 1;

    for (size_t old_index = true_len - 1; old_index >= 0; --old_index) {
        assert(new_index >= old_index);
        if (str[old_index] == ' ') {
            str[new_index] = '0';
            str[new_index - 1] = '2';
            str[new_index - 2] = '%';
            new_index -= 3;
        } else {
            str[new_index] = str[old_index];
            --new_index;
        }
    }
}

/* Test both is_unique implementations. */
void test_is_unique() {
    assert(is_unique("abcdefg"));
    assert(!is_unique("abcdefa"));
    assert(is_unique(""));

    assert(is_unique_no_data_structs("abcdefg"));
    assert(!is_unique_no_data_structs("abcdefa"));
    assert(is_unique_no_data_structs(""));
}

/* Test the is_permutation function. */
void test_is_permutation() {
    assert(is_permutation("abcdefgh", "hgfedcba"));
    assert(is_permutation("aaabbbccc", "cccbbbaaa"));
    assert(!is_permutation("abcdefgh", "hgfedcbb"));
    assert(!is_permutation("", "abcdefgh"));
    assert(is_permutation("", ""));
}

/* Test the urlify function. */
void test_urlify() {
    const char *input = "Mr John Smith";
    const char *expected_output = "Mr\%20John\%20Smith";

    size_t bufsize = strlen(expected_output);
    char *input_buf = new char[bufsize];
    strcpy(input_buf, input);
    assert(!strcmp(input, input_buf));

    urlify(input_buf, bufsize, strlen(input));

    assert(!strcmp(input_buf, expected_output));
}

/* Test all functions. */
int main() {
    test_is_unique();
    test_is_permutation();
    test_urlify();

    std::cout << "All assertions passed." << std::endl;

    return 0;
}

