![](https://github.com/rockenman1234/Argon/blob/main/argon.iconset/icon_128x128.png?raw=true) Argon
==========


## What's Argon?
Argon is an inert gas that is known for its stability and unreactivity under normal conditions. Similarly, this software is designed to provide a stable and reliable platform for software developers and power users to get started using their Mac's as quickly and reliably as possible - without leaving a trace. 

![](https://github.com/rockenman1234/Argon/blob/main/Screenshots/main1.png?raw=true)
***

## How Does It Work?
Argon is a MacOS app that automates the installation of various software and tools for developers and power users. The script uses Homebrew, a popular package manager for MacOS, and allows users to install developer tools, high-level and low-level programming languages, ease-of-life software, web browsers, and multimedia software all at the click of a button!

The App's installation process is divided into several steps, each of which corresponds to a button in the graphical user interface (GUI) built ontop of Java's Swing GUI framework. Users can click each of the buttons in the specified order to run the corresponding commands in the Terminal. The script also provides tooltips for each button to explain what software or tool will be installed.

If you are a macOS user who wants to set up a development environment quickly and easily, or if you are a developer who wants to customize the installation process to your needs, Argon may be a useful tool for you. Give it a try and let us know your feedback!

## Installation Procedure - Easy
Head over to the releases tab, and download the MacOS package (.pkg) installer. Or click here for a direct download: 
<a href="https://github.com/rockenman1234/Argon/releases/download/v0.0.1/Argon-1.0.pkg" title="Download Argon">Latest Release: 0.0.1</a>

#### Step 1). Open the installer package (.pkg) file:
![](https://raw.githubusercontent.com/rockenman1234/Argon/main/Screenshots/install1.png)

#### Step 2). Select your desired installation location:
![](https://raw.githubusercontent.com/rockenman1234/Argon/main/Screenshots/install2.png)

#### Step 3). Enter your TouchID or Password to confirm the install:
![](https://raw.githubusercontent.com/rockenman1234/Argon/main/Screenshots/install3.png)

#### Step 4). Close the Installer, and Run Argon for your Spotlight Search:
![](https://raw.githubusercontent.com/rockenman1234/Argon/main/Screenshots/install4.png)

#### Step 5). Run Argon for your Spotlight Search:
![](https://github.com/rockenman1234/Argon/blob/main/Screenshots/main1.png?raw=true)

#### Step 6). Once you've allowed Argon to install your desired applications, close Argon and delete the Argon.app file from your System's Applications Folder! 

Congratulations! 🎉 
You're done with Argon!

*** 
Alternatively, you may choose to run the precompiled Java executable file (.jar), which is also listed in the releases tab. To run it, follow these instructions, and have the latest JavaJDK installed on your Mac:

```
cd ~/Downloads

java -jar Argon.jar
```


## Installation Procedure - Advanced (Not Recommended)

First, make sure you have the latest JavaJDK installed, along with git.

Step 1). Generate the Java Executable (.jar) file:
```
git clone https://github.com/rockenman1234/Argon.git

cd ~/Argon/src

javac *.java

jar cvfm Argon.jar manifest.txt *.class

java -jar Argon.jar
```

Step 2). OPTIONAL - If you'd like to compile to your own installer package (.pkg) file, follow this steps:
```
jpackage --name Argon --input . --main-jar Argon.jar \
                                                 --resource-dir ~/Argon/package/macos --type pkg
```
 This should create an installer package (.pkg) file inside of the "src" directory. 
 
 ***

### Notice:
Please use this application with caution and run it at your own risk - as it is experimental, and may contain bugs that could potentially harm your system!
***
