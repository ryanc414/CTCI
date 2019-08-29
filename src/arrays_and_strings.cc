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

bool is_unique(const std::string &input);
bool is_unique_no_data_structs(const std::string &input);
bool is_permutation(const std::string &str_a, const std::string &str_b);
std::map<char, int> letter_freq(const std::string &str);
void urlify(char *str, size_t buf_len, size_t true_len);
bool is_palindrome_permutation(const std::string &str);
bool is_one_away(const std::string &str_a, const std::string &str_b);
bool is_one_replace_away(const std::string &str_a, const std::string &str_b);
bool is_one_insert_away(const std::string &str_a, const std::string &str_b);

void test_is_unique();
void test_is_permutation();
void test_urlify();
void test_palindrome_permutation();
void test_one_away();

/* Check if a string has all unique characters. */
bool is_unique(const std::string &input) {
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
bool is_unique_no_data_structs(const std::string &input) {
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
bool is_permutation(const std::string &str_a, const std::string &str_b) {
    if (str_a.length() != str_b.length()) {
        return false;
    }

    return letter_freq(str_a) == letter_freq(str_b);
}

/* Return a mapping of letter to its frequency in a word. */
std::map<char, int> letter_freq(const std::string &str) {
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
    int new_index = buf_len - 1;

    for (int old_index = true_len - 1; old_index >= 0; --old_index) {
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

/**
 * Check if a word is a permutation of a palindrome: a word that reads the
 * same forwards or backwards.
 */
bool is_palindrome_permutation(const std::string &str) {
    // First generate a mapping of the letter frequencies. We can ignore
    // casing and ignore punctuation characters. For simplicity, assume only
    // letters a-z and not any accented unicode letters etc.
    constexpr int kNumLetters = 26;
    std::array<int, kNumLetters> letter_freqs;
    letter_freqs.fill(0);

    for (char c : str) {
        if (c >= 'a' && c <= 'z') {
            ++letter_freqs[c - 'a'];
        } else if (c >= 'A' && c <= 'Z') {
            ++letter_freqs[c -'A'];
        }
    }

    // Now iterate through the frequencies for each letter. We know that a
    // string is a permutation of a palindrome iff it has at most 1 character
    // with an odd frequency - this would be the middle letter in a palindrome.
    bool odd_found = false;
    for (int i = 0; i < kNumLetters; ++i) {
        if (letter_freqs[i] % 2 == 1) {
            if (odd_found) {
                return false;
            } else {
                odd_found = true;
            }
        }
    }

    return true;
}

/**
 * Assuming a string can be edited by replacing, inserting or removing a
 * single character, return if a string is one (or zero) edits away.
 */
bool is_one_away(const std::string &str_a, const std::string &str_b) {
    int length_diff = str_a.length() - str_b.length();

    switch (length_diff) {
        case 0:
            return is_one_replace_away(str_a, str_b);

        case -1:
            return is_one_insert_away(str_a, str_b);

        case 1:
            return is_one_insert_away(str_b, str_a);

        default:
            return false;
    }
}

/* Check if two strings of equal length are at most one character different. */
bool is_one_replace_away(const std::string &str_a, const std::string &str_b) {
    bool replace_found = false;

    for (size_t i = 0; i < str_a.length(); ++i) {
        if (str_a[i] != str_b[i]) {
            if (replace_found) {
                return false;
            } else {
                replace_found = true;
            }
        }
    }

    return true;
}

/* Check if str_b is one character insert away from str_a. */
bool is_one_insert_away(const std::string &str_a, const std::string &str_b) {
    bool insert_found = false;

    size_t a_ix = 0;

    for (size_t b_ix = 0; b_ix < str_b.length(); ++b_ix) {
        if (str_a[a_ix] != str_b[b_ix]) {
            if (insert_found) {
                return false;
            } else {
                insert_found = true;
            }
        } else {
            ++a_ix;
        }
    }

    return true;
}

/**
 * TESTS
 */

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

    size_t bufsize = strlen(expected_output) + 1;
    char *input_buf = new char[bufsize];
    strcpy(input_buf, input);
    assert(!strcmp(input, input_buf));

    urlify(input_buf, bufsize, strlen(input) + 1);

    assert(!strcmp(input_buf, expected_output));
}

/* Test the is_palindrome_permutation function. */
void test_palindrome_permutation() {
    assert(is_palindrome_permutation("Tact Coa"));
    assert(!is_palindrome_permutation("Tact Coat"));
    assert(is_palindrome_permutation(""));
}

/* Test the is_one_away function. */
void test_one_away() {
    assert(is_one_away("pale", "ple"));
    assert(is_one_away("pales", "pale"));
    assert(is_one_away("pale", "bale"));
    assert(!is_one_away("pale", "bae"));
    assert(is_one_away("", ""));
}

/* Test all functions. */
int main() {
    test_is_unique();
    test_is_permutation();
    test_urlify();
    test_palindrome_permutation();
    test_one_away();

    std::cout << "All assertions passed." << std::endl;

    return 0;
}

