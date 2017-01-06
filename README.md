PEiD (alpha version)
====

Yet another implementation of PEiD with yara 

Features
----
* __don't need to install yara and download yara rules__
* support multiple file types: PE, Malicious Documents, etc
* multi platform support: Linux, Windows
* analyze outputs of yara (see following output)

Usage
----
```
% ./PEiD --prepare # if yara and yara rules does not exists 
INFO[0000] prepare successfuly                          
% ./PEiD cmd/anti_dbg_msgbox/anti_dbg_msgbox-upx.exe
INFO[0000] yara = '/home/katc/bin/PEiD/yara'            
INFO[0000] all requirements met                         
RULES_FILE = /home/katc/malware/rules/index.yar
cmd/anti_dbg_msgbox/anti_dbg_msgbox-upx.exe =>
  PE : 32 bit
  DLL : no
  Packed : yes
  Anti-Debug : no (yes)
  GUI Program : no (yes)
  Console Program : yes
  contains base64
  PEiD : ["UPX_wwwupxsourceforgenet_additional" "yodas_Protector_v1033_dllocx_Ashkbiz_Danehkar_h" "UPX_290_LZMA" "UPX_290_LZMA_Markus_Oberhumer_Laszlo_Molnar_John_Reiser" "UPX_290_LZMA_additional" "UPX_wwwupxsourceforgenet"]
```


Requirement
----
### run
there's no requirements!

### build
install

* git
* make
* go
* go-bindata


Build
----

(optional) Download latest following releases to `/data`

* yara
* yara rules: https://github.com/Yara-Rules/rules/

Run following command to `go get` packages

```bash
export GOPATH=`pwd`
make init
```

Finally, 

```bash
make
```


TODO
----
- [ ] version info
- [ ] Colorize analysis result
- [ ] Support Mac
