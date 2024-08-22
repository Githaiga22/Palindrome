# Palindrome Checker in Go
### Overview
This Go program reads a string from the user and checks if it is a palindrome. A palindrome is a word, phrase, number, or other sequences of characters that read the same forward and backward (ignoring spaces, punctuation, and capitalization).

## Implementation
- The program uses the following steps to check if a string is a palindrome:

1. Read a string from the user using fmt.Scanln.
2. Convert the input string to lowercase using strings.ToLower to ignore capitalization.
3. Remove spaces from the input string using strings.ReplaceAll to ignore spaces.
4. Call the isPalindrome function to check if the modified string is a palindrome.
5. Print the result to the user.

The isPalindrome function uses a simple iterative approach to check if a string is a palindrome:

1. Iterate over the first half of the string (up to the middle character).
2. Compare each character with the corresponding character from the end of the string.
3. If any pair of characters does not match, return false.
4. If the loop completes without finding any mismatches, return true.

### Usage
To run the program, follow these steps:

1. Install Go on your system if you haven't already.
2. Clone this repository using te following command;
```bash
git clone https://github.com/Githaiga22/Palindrome.git
```
3. Open a terminal and navigate to the directory ;
```bash
cd Palindrome
```
4. Run the program using the command ;
```bash
 go run main.go.
 ```
5. Enter a string when prompted, and the program will print whether it is a palindrome or not.

### Example Use Cases
- Enter a palindrome string: "madam" Output: "madam is a palindrome"
- Enter a non-palindrome string: "hello" Output: "hello is not a palindrome"

### Contributing:
Contributions are welcome! If you'd like to contribute to Groupie Tracker, please follow these steps:

1. Fork the repository:
```bash
 git fork https://github.com/Githaiga22/Palindrome.git
 ```
2. Create a new branch:
```bash
 git branch feature/new-feature
 ```
3. Make your changes and commit them: 
```bash
git commit -m "Added new feature"
```
4. Push your changes: 
```bash
git push origin feature/new-feature
```
5. Create a pull request: 
```bash
git request-pull origin feature/new-feature
```
### License:
Groupie Tracker is licensed under the MIT License. See LICENSE for details.

### Acknowledgments:
[codewars](https://www.codewars.com)

### Author
- Githaiga22


### Github Profile
- Github profile: [Githaiga22](https://github.com/Githaiga22)

