# Data Encryptor CLI Application Using Golang

## Overview
Data Encryptor is a command-line tool built with Golang that allows users to encrypt and decrypt files securely. It takes a file as input, encrypts or decrypts the content, and saves the output in a new file.

## Features
- Encrypts text files securely
- Decrypts encrypted files
- Simple CLI interface using `urfave/cli/v2`

## Installation
Ensure you have Go installed on your system. If not, download and install it from [Go official site](https://go.dev/dl/).

### Clone the Repository
```sh
git clone https://github.com/Sushank-ghimire/go.git
cd go/projects/data-encryptor
```

### Install Dependencies
```sh
go mod tidy
```

## Usage
### Encrypt a File
```sh
go run . encrypt --input "data.txt"
```
This will generate an encrypted file `data_encrypted.txt`.

### Decrypt a File
```sh
go run . decrypt --input "data_encrypted.txt"
```
This will generate a decrypted file `data_decrypted.txt`.

## Flags
- `--input, -i` → Specifies the file to be encrypted or decrypted.

## Example
```sh
go run . encrypt --input "message.txt"
# Output: Encrypted data saved to encryptedData.txt

go run . decrypt --input "encryptedData.txt"
# Output: Decrypted data saved to decryptedData.txt
```

## Project Structure
```
/go
│── main.go
│── Utils
│   │── Encryptor.go
│── README.md
│── go.mod
│── go.sum
```

## Dependencies
- [`urfave/cli/v2`](https://github.com/urfave/cli) - Used for building the CLI

## Contributing
Feel free to fork the repository and submit a pull request if you would like to contribute!

## License
This project is licensed under the MIT License.

---
Developed by [Sushank Ghimire](https://github.com/Sushank-ghimire) 🚀

