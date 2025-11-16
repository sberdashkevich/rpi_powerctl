Build:
`set GOOS=linux&&set GOARCH=arm&&set GOARM=6&&go build -x -v -o powerctl main.go`

`sudo crontab -e`
<br>
`@reboot /home/pi/powerctl`
<br>
Example of usage:
```
http://192.168.1.194:8080/shutdown
```
