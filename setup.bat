go run ./internal/setup/check.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run ./internal/setup/generate.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run ./internal/setup/install.go %*
if %errorlevel% neq 0 exit /b %errorlevel%
go run ./internal/setup/test.go %*
