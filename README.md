PEiD (alpha version)
====

Yet another implementation of PEiD with yara 

Features
----
* support multiple file types: PE, Malicious Documents, 
* analyze outputs of yara (see following output)

```
% ./PEiD cmd/anti_dbg_msgbox/anti_dbg_msgbox.exe
INFO[0000] all rewuirements met                         
RULES_FILE = /home/katc/malware/rules/index.yar
cmd/anti_dbg_msgbox/anti_dbg_msgbox.exe =>
["possible_includes_base64_packed_functions" "IsPE32" "IsConsole" "contentis_base64" "DebuggerException__SetConsoleCtrl" "SEH__vectored" "anti_dbg" "network_udp_sock" "network_tcp_listen" "network_tcp_socket" "network_dns" "win_registry" "win_token" "win_files_operation" "Str_Win32_Winsock2_Library" "without_urls" "without_images" "without_attachments"]
  PE : 32 bit
  DLL : no
  Anti-Debug : yes
  Packed : no
  GUI Program : yes
  contains base64

```

TODO
----
- [ ] Fix GOOS env variable in some way
- [ ] Packer support
- [ ] More support for yara rules
- [ ] Add automatic yara rules location finder for Windows/Mac (implement automatic configurator)
- [ ] Colorize analysis result