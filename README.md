# gophersay
Gophersay is a small go command line script that will make the Gopher from the Go logo say whatever you want.
The gopher is printed in ASCII Art.

To use it on the command line, use:
```bash
gophersay Hi mom!
```
```
    `.-::::::-.`                
.:-::::::::::::::-:. ===========
`_:::    ::    :::_` = Hi mom! =
 .:( ^   :: ^   ):.  ===========
 `:::   (..)   :::.             
 `:::::::UU:::::::`             
 .::::::::::::::::.             
 O::::::::::::::::O             
 -::::::::::::::::-             
 `::::::::::::::::`             
  .::::::::::::::.              
    oO:::::::Oo                 
```

## Examples
```
    `.-::::::-.`                            
.:-::::::::::::::-:. =======================
`_:::    ::    :::_` = hello gopher world! =
 .:( ^   :: ^   ):.  =======================
 `:::   (..)   :::.                         
 `:::::::UU:::::::`                           
 .::::::::::::::::.                         
 O::::::::::::::::O                         
 -::::::::::::::::-                         
 `::::::::::::::::`                         
  .::::::::::::::.                          
    oO:::::::Oo                             
```

## Installation
- [Make](#install-using-make)
- [Manual](#manual-installation)
### Install using make
```bash
make gophersay
make install
```
### Manual Installation
```bash
go build main.go -o gophersay
```
To install:
```
sudo cp gophersay /usr/local/bin/gophersay
```
