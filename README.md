# gomorse

This project implements a Morse code decoder in Go, encapsulated in the MorseDecoder struct.

The implementation uses a map (map[string]string) that associates each Morse code symbol with its corresponding character (A-Z, 0-9).

Decoding is done by splitting the input code into words using a configurable word separator (default: three spaces " "), then splitting each word into letters using another configurable letter separator (default: single space " ").

Each encoded letter is then translated to its corresponding character, or replaced by a fallback character (default: "?") if the code is invalid or unknown.

The final output is a decoded string with words separated by spaces.

## Running

### With Docker
Build the Docker image:

```bash
docker build -t gomorse .
```

Run the container with your Morse code input:
```bash
docker run --rm gomorse ".... . -.--   .--- ..- -.. ."
```

### Running with golang
You can run the program directly with Go or compile it first.
```bash
go run ./main.go ".... . -.--   .--- ..- -.. ."
```

Or compile and run:  
```
Make build
./releases/gomorse ".... . -.--   .--- ..- -.. ."
```

## Tests

```bash
Make test
```


