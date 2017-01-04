PEiD (alpha version)
====

Yet another implementation of PEiD with yara 

Features
----
* analyses outputs of yara (see following output)

```
/home/katc/malware/test/PlugX/C116CD083284CC599C024C3479CA9B70_2.tmp_ =>
["Microsoft_Visual_Cpp_v50v60_MFC" "Microsoft_Visual_Cpp_v60_DLL" "IsPE32" "IsDLL" "IsWindowsGUI" "IsPacked" "HasOverlay" "HasDigitalSignature" "HasRichSignature" "contentis_base64" "anti_dbg" "win_mutex" "win_files_operation" "with_urls" "without_images" "without_attachments"]
  PE : 32 bit
  DLL : yes
  Anti-Debug : yes
  Packed : yes
  GUI Program : yes
  mutex : yes
  contains base64
  contains urls
```

TODO
----
- [ ] Fix GOOS env variable in some way
- [ ] Packer support
- [ ] More support for yara rules
- [ ] Add automatic yara rules location finder (implement automatic configurator)
- [ ] Colorize analysis result