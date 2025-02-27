package main

import (
	"data-encryptor/utils"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	app := &cli.App{
		Name:  "data-encryptor",
		Usage: "Used for data encryption and decryption",
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "Encrypts the input data",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Aliases:  []string{"i"},
						Usage:    "File path name to take the data",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					filename := c.String("i")

					basename := filepath.Base(filename)

					file_nameonly := strings.TrimSuffix(basename, ".txt")

					data, err := os.ReadFile(filename)
					if err != nil {
						return fmt.Errorf("failed to read file: %w", err)
					}
					encryptedData := Encryption.EncryptData(string(data), 5) // Convert to string if needed
					outputFile := file_nameonly + "_encrypted.txt"

					// Write the encrypted data to the file
					err = os.WriteFile(outputFile, []byte(encryptedData), 0644)
					if err != nil {
						return fmt.Errorf("failed to write encrypted data: %w", err)
					}

					fmt.Println("Encrypted data saved to:", outputFile)
					return nil
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decryption of data using the filename",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Aliases:  []string{"d"},
						Usage:    "To take the pathname of the encrypted file",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					filename := c.String("d")
					basename := filepath.Base(filename)
					// Extracting the filename only to save the decrypted file
					filename_only := strings.TrimSuffix(basename, ".txt")

					// getting the decrypted file data
					data, err := os.ReadFile(filename)
					if err != nil {
						return fmt.Errorf("failed to read file: %w", err)
					}
					decryptedData := Encryption.DecryptData(string(data), 5)

					saveFileName := filename_only + "_decrypted.txt"

					fileErr := os.WriteFile(saveFileName, []byte(decryptedData), 0644)
					if fileErr != nil {
						return fmt.Errorf("failed to write decrypted data: %w", fileErr)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
