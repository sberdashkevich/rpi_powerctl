Build:
`set GOOS=linux&&set GOARCH=arm&&set GOARM=6&&go build -x -v -o powerctl main.go`

`sudo crontab -e`
`@reboot /usr/local/bin/powerctl`
