go get golang.org/x/crypto/ssh
if %errorlevel% neq 0 exit /b %errorlevel%
go run ./internal/setup/check.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run %GOPATH%/src/github.com/therecipe/qt/internal/setup/generate.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run %GOPATH%/src/github.com/therecipe/qt/internal/setup/install.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run %GOPATH%/src/github.com/therecipe/qt/internal/setup/test.go %*
