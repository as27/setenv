# setenv
Simple add some env variables

## Installation

Just use go get:

``` go get github.com/as27/setenv```

## Usage

### Create a file

Create a textfile for example `.setenv`. The file accepts empty lines and comments with `#`. The logic is always `KEY=VALUE`.

```
APP_ID=abc12345
APP_SECRET=abc123456789
```

In that case the env variables APP_ID and APP_SECRET are going to be set.

### Use in code

You can load the file inside your app. The best part would be the init() function to ensure that the enviroment variables are set before you get them over the code.

```go
func init(){
    setenv.ParseFile(".setenv")
}
```
