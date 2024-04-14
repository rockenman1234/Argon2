![](https://github.com/rockenman1234/Argon2/blob/main/img/icon-128x128.png?raw=true) Welcome to a faster Mac! 
==========


## What's Argon?
Argon is an inert gas that is known for its stability and unreactivity under normal conditions. Similarly, this software is designed to provide a stable and reliable platform for software developers and power users to get started using their Mac's as quickly and reliably as possible - without leaving a trace. 

![](https://github.com/rockenman1234/Argon2/blob/main/img/Main.png?raw=true)
***

## How Does It Work?
Argon is a MacOS app that automates the installation of various software and tools for developers and power users. The script uses Homebrew, a popular package manager for MacOS, and allows users to install developer tools, high-level and low-level programming languages, ease-of-life software, web browsers, and multimedia software all at the click of a button!

The App's installation process is divided into several steps, each of which corresponds to a button in the graphical user interface (GUI) built ontop of Go's Fyne GUI framework. Users can click each of the buttons in the specified order to run the corresponding commands in the Terminal. The script also provides a cleaner to help cut down on junk files!

If you are a macOS user who wants to set up a development environment quickly and easily, or if you are a developer who wants to customize the installation process to your needs, Argon may be a useful tool for you. Give it a try and let us know your feedback!

## Installation Procedure
Head over to the releases tab, and download the MacOS package (.zip) file. Or click here for a direct download: 
<a href="https://github.com/rockenman1234/Argon2/releases/download/0.2.0/Argon2.zip" title="Download Argon">Latest Release: 0.2.0</a>

#### Step 1). Unzip the file:
![](https://github.com/rockenman1234/Argon2/blob/main/img/Unzip.png?raw=true)

#### Step 2). Open it:
![](https://github.com/rockenman1234/Argon2/blob/main/img/run.png?raw=true)

#### Step 2b). If You get an error, just ignore it and paste the command below into the Terminal:
![](https://github.com/rockenman1234/Argon2/blob/main/img/error.png?raw=true)

```
xattr -cr "$HOME/Downloads/Argon2.app/"
```
This is just Apple attempting to quarentine files downloaded from the internet, in order to let Argon run you'll need to disable this functionality on this file. 

#### Step 3). Rock On!🤘🎸
![](https://github.com/rockenman1234/Argon2/blob/main/img/Main.png?raw=true)



#### Congrats! 🎉 You're all done and ready to rock and roll with a dev friendly Mac!

Once you're done, you can either drag and drop the Argon2.app file into your applications folder - or delete it from your system! Because just like Argon - it's gone without a trace!

*** 
Alternatively, you may choose to compile and run the Go binary yourself from source. To do so, just download the source code and run the following commands! Make sure you've got the latest version of Go installed before you start!

```
cd "$HOME/Downloads/"

unzip Argon2-0.2.0.zip

cd Argon2-0.2.0

go mod init Argon2

go get fyne.io/fyne/v2@latest

go install fyne.io/fyne/v2/cmd/fyne@latest

fyne package

open Argon2.app
```
If you get any errors with the code above, please check out the [Fyne setup tool](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/) and the [Fyne troubleshooting guide](https://docs.fyne.io/faq/troubleshoot) for more help on what went wrong!

 
 ***

### Notice:
Please use this application with caution and run it at your own risk - as it is experimental, and may contain bugs that could potentially harm your system!
***
