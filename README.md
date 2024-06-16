## json2base64

Here is Go tool for decode json file to base64 string

> Work only with valid JSON files

### Usage

`-f /path/to/input.json` Path to the credentials file (default `credentials.json`)

`-o /path/to/out.64` Path to save the base64 encoded output

`-v` Show application version

---
### Examples
Encoded credentials.json file with content `{"test":"test"}` put result to `stdout`
```shell
json2base64
# eyJ0ZXN0IjoidGVzdCJ9
```

Encoded some_file.json file with content `{"test":"test"}` put result to `stdout`
```shell
json2base64 -f ./some_file.json
# eyJ0ZXN0IjoidGVzdCJ9
```

Encoded some_file.json file save result to `some_file.b64` file
```shell
json2base64 -f ./some_file.json -o ./some_file.b64
# Base64 encoded content saved to some_file.b64
```

Print app version
```shell
json2base64 -v
# 1.0.0
```
