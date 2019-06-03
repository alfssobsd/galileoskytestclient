# Test Client Galileosky

Its simple test client for GPS Tracking protocol GalileoSky.
I think this is a tool could help you write your own tracking server.

## Clients
7.0, 7.0 Lite, 5.0, 5.1 - client_v7.go

## Links to docs
* https://7gis.ru/assets/files/docs/manuals_ru/opisanie-protokola-obmena-s-serverom-(487740-v16).pdf
* https://galileosky.com/assets/files/docs/manuals_en/general-information.-server-exchange-protocol-description-(400176-v16).pdf
* Page 103-107 https://7gis.ru/assets/files/docs/manuals_ru/rukovodstvo-polzovatelya-programmiruemyie-kontrolleryi-galileosky-7-(493048-v9).pdf
* Page 45-49 https://dl.dropboxusercontent.com/s/ig2e3224m96g7sh/UserManual_En_Lite_0192.pdf

## HowTo Build
```
dep ensure
go build -o client_v7 client_v7.go
./client_v7 -help
```

## Usage

### Send just one pack
```
./client_v7 -host 127.0.0.1 -port 9998 -usecase onepack -pack 01148001820214033836363731303033373930343936359A15 -response 029a15
```

### Emulate movement
```
./client_v7 -host 127.0.0.1 -port 9998 -usecase movement -movement_config_path client_v7_test_movement.yml
```

### Help
```
./client_v7 --help
```