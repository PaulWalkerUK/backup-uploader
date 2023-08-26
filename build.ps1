$scriptPath = (Get-Variable MyInvocation -Scope Script).Value.MyCommand.Path
$dir = Split-Path $scriptpath
Push-Location $dir

$goos = $env:GOOS
$goarch = $env:GOARCH

$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin\windows\backup-uploader.exe
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin\linux\backup-uploader

Copy-Item .\config.json.example bin\windows
Copy-Item .\config.json.example bin\linux

$env:GOOS = $goos
$env:GOARCH = $goarch

Pop-Location