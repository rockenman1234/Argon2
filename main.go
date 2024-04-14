/*
 *   _______  _____     ________ ________ __    _
 *  |   _   ||    _ |  |       ||       ||  |  | |
 *  |  |_|  ||   | ||  |    ___||   _   ||   |_| |
 *  |       ||   |_||_ |   | __ |  | |  ||       |
 *  |       ||    __  ||   ||  ||  |_|  ||  _    |
 *  |   _   ||   |  | ||   |_| ||       || | |   |
 *  |__| |__||___|  |_||_______||_______||_|  |__|
 *
 * Build: 0.2.0 - ALPHA
 *
 *
 * Mission Statement: Argon is an inert gas that is known for its stability and unreactivity under normal conditions.
 *  Similarly, this software could be designed to provide a stable and reliable platform for software developers and power users
 *  to get started using their Mac's without leaving a trace.
 *
 *
 * Developer: 	Kenneth (Alex) Jenkins - Computer Engineering Student @ Georgia Tech (Go Jackets!)
 *
 * CHANGE CHANGELOG:
 *     0.2.0: Code rebase from java swing to GOLANG using the Fyne.io framework
 *
 */

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"os/exec"
)

func main() {
	a := app.New()
	// this is the main window
	w := a.NewWindow("Argon2")
	title := widget.NewLabelWithStyle("Welcome to Argon2!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true, Italic: false, Monospace: false})
	title.Resize(fyne.NewSize(200, 100))
	// title of the window, and the content of the window

	inst1 := widget.NewLabelWithStyle("Required Steps:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: true})
	// Button 1 - Install HomeBrew
	installButton1 := widget.NewButton("1). Install HomeBrew", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"/bin/bash -c \\\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\\\"\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Close loading window after successful installation
		//loadingWindow.Close()

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})
	// End Button 1

	// Button 2 - Install Xcode Command Line Tools
	installButton2 := widget.NewButton("2). Install Developer Tools", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"/bin/bash -c \\\"$(xcode-select --install; brew install wget coreutils curl)\\\"\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// End Button 2

	// Language Install Buttons:
	label1 := widget.NewLabelWithStyle("Install Programing language support:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: false})

	// Button - Install Python, Node.js, and Ruby
	installButtonUpperProg := widget.NewButton("Install Python, Node.JS, & Ruby", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install python node ruby\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install Rust, Go, Java, and DOTNET (C#)
	installButtonLowerProg := widget.NewButton("Install Rust, Go, Java, & DOTNET (C#)", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install rust google-go openjdk dotnet\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// DevOps Install Buttons:
	label2 := widget.NewLabelWithStyle("Install DevOps software:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: false})

	// Button - Install Docker and Kubernetes
	installButtonDevSoft := widget.NewButton("Install Docker & Kubernetes", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install docker kubernetes-cli\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install VScode
	installButtonIDE := widget.NewButton("Install Visual Studio Code, & Git", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install visual-studio-code git\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Ease-Of-Life Install Buttons:
	label3 := widget.NewLabelWithStyle("Install Ease-Of-Life software:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: false})

	// Button - Install Ease-Of-Life Software (OnyX, Shottr, Alt-Tab, Rectangle, and Bitwarden)
	installButtonEOLsoft := widget.NewButton("Install OnyX, Shottr, Alt-Tab,\nRectangle, & Bitwarden", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install onyx shottr alt-tab rectangle bitwarden\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install Firefox
	installFirefox := widget.NewButton("Install Mozilla Firefox", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install firefox)\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install Chrome
	installButtonChrome := widget.NewButton("Install Google Chrome", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install google-chrome\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	installButtonOpera := widget.NewButton("Install Opera", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install opera\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install TOR Browser
	installButtonTOR := widget.NewButton("Install Tor Browser", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install install --cask tor-browser\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Install Multimedia Software
	label4 := widget.NewLabelWithStyle("Install Multimedia software:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: false})

	// Button - Install VLC and Audacity
	installButtonAudioVideo := widget.NewButton("Install VLC & Audacity", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install --cask vlc --cask audacity\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install OBS and Handbrake
	installButtonMediaConversion := widget.NewButton("Install OBS & Handbrake", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install --cask obs --cask handbrake\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Button - Install GIMP and Inkscape
	installButtonGraphicDesign := widget.NewButton("Install GIMP & Inkscape", func() {
		// Execute command to install Homebrew
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"brew install --cask gimp --cask inkscape\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// System Maintenance Buttons:
	label5 := widget.NewLabelWithStyle("System Maintenance:", fyne.TextAlignCenter, fyne.TextStyle{Bold: false, Italic: false, Monospace: false})

	// Clean temp Files Button
	restartMacButton := widget.NewButton("Clean Up Temporary Files", func() {
		// Execute command to open local script example.sh with osascript
		cmd := exec.Command("osascript", "-e", "tell app \"Terminal\" to do script \"/bin/bash -c \\\"$(sudo curl -fsSL https://raw.githubusercontent.com/rockenman1234/Argon2/main/Scripts/cleaner.sh)\\\"\"")
		err := cmd.Run()
		fmt.Printf("%s\n", cmd)
		if err != nil {
			fmt.Println("Error executing command:", err)
			// Show error dialog
			errorWindow := a.NewWindow("This is awkward...")
			errorWindow.SetContent(widget.NewLabel("Oops an error occurred!\n" +
				"\nPlease try again, and report this issue on GitHub.\n\n\n Error: " +
				err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Starting Install", "Follow instructions on Terminal window.\nWhen done, return here\nand press OK to continue.", w)
		diag.Show()
	})

	// Restart MacOS Button
	cleanMacButton := widget.NewButton("Restart MacOS", func() {
		// Execute command to restart macOS using osascript
		cmd := exec.Command("osascript", "-e", "tell app \"System Events\" to restart")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error restarting macOS:", err)
			// Show error dialog
			errorWindow := a.NewWindow("Error Restarting macOS")
			errorWindow.SetContent(widget.NewLabel("An error occurred while trying to restart macOS.\nPlease try again, and report this issue on GitHub.\n\nError: " + err.Error()))
			errorWindow.Show()
			return
		}

		// Show message dialog
		diag := dialog.NewInformation("Restarting macOS", "Your Mac will restart shortly.", w)
		diag.Show()
	})

	// About Button
	aboutButton := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		popup := a.NewWindow("About")

		urlObj, err := url.Parse("https://github.com/rockenman1234/Argon")
		if err != nil {
			return
		}

		content := container.NewVBox(
			widget.NewIcon(theme.InfoIcon()),
			widget.NewLabel("Argon2 is a continuation of the Argon project,\nrebuilt in GO.\n"+
				"\nArgon2 is designed to allow power users to \nquickly set up their Macs with the tools they need.\n"),
			widget.NewLabel("For more information, visit:"),
			widget.NewHyperlink("Our GitHub", urlObj),
		)

		// Create a container to center the content vertically and horizontally
		centeredContent := container.NewVBox(
			container.NewCenter(content),
		)

		popup.SetContent(centeredContent)
		popup.Resize(fyne.NewSize(300, 200)) // Resizing the popup window
		popup.SetFixedSize(true)
		popup.Show()
	})

	// Adds content to the window
	content := container.NewVBox(
		title,
		aboutButton,
		inst1,
		container.NewHBox(
			installButton1,
			installButton2,
		),
		label1,
		container.NewHBox(
			installButtonUpperProg,
			installButtonLowerProg,
		),
		label2,
		container.NewHBox(
			installButtonDevSoft,
			installButtonIDE,
		),
		label3,
		container.NewHBox(
			installButtonEOLsoft,
			installFirefox,
		),
		container.NewHBox(
			installButtonChrome,
			installButtonOpera,
			installButtonTOR,
			// Add more buttons here
		),
		label4,
		container.NewHBox(
			installButtonAudioVideo,
			installButtonGraphicDesign,
		),
		container.NewHBox(
			installButtonMediaConversion,
		),
		label5,
		container.NewHBox(
			restartMacButton,
			cleanMacButton,
		),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(450, 500)) // resize the window
	w.SetFixedSize(true)

	w.ShowAndRun()
}
