# goencrypter
A CLI app to encrypt binary files with algorithms implemented from scratch - created for the purpose of my university classes

For now it is only using DES (not fully optimized, I know)

Build using go build run goencrypter, specify -input=filepath and -encrypt=false if you want to decrypt
 ```bash
go build
./goencrypter -input=test3.bin
```
 <img src="https://github.com/totoafreeca/goencrypter/blob/master/images/image1.jpg" />
 
 ### Example file test3.bin test with ghex and diff command
 
 Encrypted (left) and original (right)
 
 <img src="https://github.com/totoafreeca/goencrypter/blob/master/images/image2.jpg" />
 
 Decrypted file
 
 <img src="https://github.com/totoafreeca/goencrypter/blob/master/images/image3.jpg" />
 
 Checking differences
 
 <img src="https://github.com/totoafreeca/goencrypter/blob/master/images/image4.jpg" />
