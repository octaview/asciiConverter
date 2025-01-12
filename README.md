ASCII Image Converter

This project is a simple Go-based tool that converts images to ASCII art representation.

Features

Supports common image formats: JPEG, PNG, and GIF

Resizes images proportionally to fit the specified output width

Maps image brightness to ASCII characters for visual representation

Requirements

Go 1.18 or higher

golang.org/x/image/draw package

Installation

Clone this repository:

git clone https://github.com/yourusername/ascii-image-converter.git
cd ascii-image-converter

Install dependencies:

go mod tidy

Usage

Place the image you want to convert in the project directory and update the image path in the code:

imgPath := "image.jpg" // Replace with your image file name

Run the program:

go run main.go

The generated ASCII art will be printed in the terminal.

Example Output

Example of the ASCII representation:

@@@@@@@@%%%%%%%%%%%%%%%###############
@@@@@@@@%%%%%%%%%%%%%%%###############
########++++++++++++++++++++++++++++++
--------------------------------------

Customization

Output Width: Adjust the constant outputWidth to change the width of the ASCII art output:

const outputWidth = 100

ASCII Characters: Modify the asciiChars string to use a different set of characters for mapping brightness:

const asciiChars = "@%#*+=-:."

Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue for suggestions and improvements.

License

This project is licensed under the MIT License. See the LICENSE file for more details.

