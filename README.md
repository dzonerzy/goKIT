# goKIT

golang bindings for KIT (Keep In Touch) which can be found [here](https://github.com/dzonerzy/KIT)

## Installation 

To install you need to go get like this

```bash
go get -u github.com/dzonerzy/goKIT
```

## Ported APIs

 - [x] kit_init → KIT.Init
 - [x] kit_connect → KIT.Connect
 - [x] kit_bind → KIT.Bind
 - [x] kit_listen_and_accept → KIT.ListenAndAccept
 - [x] kit_write → KIT.Write
 - [x] kit_read → KIT.Read
 - [x] kit_is_disconnect → KIT.IsDisconnect
 - [x] kit_notify_disconnect → KIT.NotifyDisconnect
 - [x] kit_disconnect → KIT.Disconnect
 - [x] kit_human_error → KIT.Error
 - [ ] kit_select → KIT.Select
 - [ ] kit_get_error → KIT.ErrorNum

 ## Usage 

 For common usage check [KIT README.md](https://github.com/dzonerzy/KIT/blob/main/README.md) or check example folder

 ## License 
```text
 	MIT License
	Copyright (c) 2021 Daniele Linguaglossa
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
```