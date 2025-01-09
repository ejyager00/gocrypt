# GoCrypt

A simple message encryption CLA. Currently supports Caesar and Vigenère ciphers.

## Flags

- `-c` boolean: flag to use caesar encryption algorithm
- `-v` boolean: flag to use Vigenère encryption algorithm
- `-e` boolean: flag for encoding
- `-d` boolean: flag for decoding
- `-i` string: the path to the input file (text to encode or decode)
- `-o` string: the path to the desired output file (where to store encoded or decoded message)
- `-k` string: the encryption key (integer for Caesar cipher, string for Vigenère cipher)

## Example usage

### Caesar cipher

```
gocrypt -c -i cleartext.txt -k 5 -e -o ciphertext.txt
gocrypt -c -i ciphertext.txt -k 5 -d -o message.txt
```

### Vigenère cipher

```
gocrypt -v -i cleartext.txt -k forget -e -o ciphertext.txt
gocrypt -v -i ciphertext.txt -k forget -d -o message.txt
```