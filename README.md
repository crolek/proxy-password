#proxy-password
There are times were you might need to update the System Variables HTTP_PROXY and HTTPS_PROXY along with one or more config files. I decided to create a simple updater tool for all of them. Right now its supports .npmrc files, but will soon support git and mercurial config files (along with anything else that makes sense).

##Features
- Creates a new config file if needed (.npmrc, mercurial.ini, etc)
- Updates a config file's password if needed (only supports .npmrc right now)
- Sets/updates the HTTP_PROXY and HTTPS_PROXY System Variables

##Use
Right now it requires admin rights to use it because it updates the System Variables

- Download the proxy-password.exe file from the `/dist` folder.
- Open a command window and navigate to where you downloaded the file
- Run the following command to update a password:

	`proxy-password.exe -password=yourSweetPasswordHere`

- If you need to create a new proxy password configuration use something like this:
	
	`proxy-password.exe -username=yourUsername -password=yourPassword -url=theProxyUrl -port=thePortNumber`


## Questions
Ping me on twitter [@crolek](http://twitter.com/crolek)

##Warning
This is very much an in-development/hacky-weekend project, so please keep that in mind.

## License

The MIT License (MIT)

Copyright (c) 2013 Tiny Factory, LLC

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.