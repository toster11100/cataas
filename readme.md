## Cataas

Cataas is a simple command-line application that downloads cat images from the internet and saves them to your computer.

### Features
- Save the cat images to your computer for use in your own projects
- Сat image size selection
- The ability to write any word on the picture
- Filter selection
- The ability to use different tags
---
## Installation
1. Clone the repository:
```bash
$ git clone https://github.com/toster11100/Cataas.git 
```
2. Change into the newly created directory:
```bash
$ cd cataas
```
3. Compile the application:
```bash
$ go build -o cataas ./cmd/cats/main.go
```
---
## Usage
Once you have compiled the application, you can run it with the following command:
```bash
./cataas -flag
```
---
## Options
The following options are available:
- name:  this is a required option
- `-t` or `--tag`: tag to be added to the image
- `-s` or `--says`: text to be displayed on the image
- `-f` or `--filter`: filter to be applied to the image
- `-h` or `--height`: height of the image in pixels
- `-w` or `--width`: width of the image in pixels

If the flags were not specified, they will be replaced by the flags from the configuration file.

---
## Example
```bash
$ ./cataas cat -t cute --say "Hello World" -f sepia -h 600 -w 600
```
---
## Configuration
Users can also specify the image characteristics in a YAML configuration file, instead of using command-line flags. The path to the configuration file can be specified using the `-c` or `--config` flag.

---
## Contributions
Contributions to this project are welcome and encouraged. If you find a bug or have an idea for a new feature, please open an issue or a pull request.


