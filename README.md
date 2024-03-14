# Text Art Encoder and Decoder

This Go program provides functionality to encode and decode text art, allowing users to compress repeated patterns in the text art.

## Features

- Encode text art by compressing repeated patterns within square brackets.
- Decode previously encoded text art to recover the original text.

## Usage

1. Clone the repository:
      git clone https://github.com/your-username/text-art-encoder-decoder.git

2. Navigate to the project directory:
      cd art

3. Build the Go program
      go build

4. Run the program:
  ./art

5. Follow the prompts to choose between encoding or decoding, specify the input type (multiple lines or single line), and provide the text art to encode or decode.

## Examples
1. 
Do you want to encode or decode? (e/d): e
Do you want to enter multiple lines or a single line? (m/s): m
Enter the text art (press Ctrl+D to finish):
[3 A]
[4 I]
Ctrl + D
Encoded text:
AAA
IIII
Do you want to encode or decode another text? (y/n)
n

2. 
Do you want to encode or decode? (e/d): d
Do you want to enter multiple lines or a single line? (m/s): s
Enter the text art or filename:
cats.encoded.txt
Encoded text:
                      /^--^\     /^--^\     /^--^\
                      \____/     \____/     \____/
                     /      \   /      \   /      \
                    |        | |        | |        |
                     \__  __/   \__  __/   \__  __/
|^|^|^|^|^|^|^|^|^|^|^|^\ \^|^|^|^/ /^|^|^|^|^\ \^|^|^|^|^|^|^|^|^|^|^|^|
| | | | | | | | | | | | |\ \| | |/ /| | | | | |\ \| | | | | | | | | | | |
| | | | | | | | | | | | |/ /| | |\ \| | | | | |/ /| | | | | | | | | | | |
| | | | | | | | | | | | |\/ | | | \/| | | | | |\/ | | | | | | | | | | | |
#########################################################################
| | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |
| | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | | |
Do you want to encode or decode another text? (y/n)
n

## Contributing
If you have any suggestions, feature requests, or bug reports, please write to me. (discord: kristjankelk)