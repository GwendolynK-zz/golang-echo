// created by: WestleyK
// email: westleyk@nym.hush.com
// https://github.com/WestleyK/golang-echo
// date: Sep 27, 2018
// version-1.0.3
//
// MIT License
//
// Copyright (c) 2018 WestleyK
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//


package main


import (
    "fmt"
    "os"
    "strings"
)

var (
    VERSION string = "version-1.0.3"
    DATE string = "Date: Sep 27, 2018"

    OPTION string = ""
    OPTION2 string = ""
    OPTION3 string = ""
    SCRIPT_NAME string = ""
    noNewLine int = 0
    NO_RESET int = 0
    OPTION_E int = 0
    color string = ""
    VAR string = ""
    MSG string = ""

    colorReset string = "\x1b[0m"
    red string = "\x1b[31m"
    yello string = "\x1b[33m"
    blue string = "\x1b[34m"
    pink string = "\x1b[35m"
    white string = "\x1b[37m"
    teal string = "\x1b[36m"
)

func version_print() {
    fmt.Print(VERSION)
    fmt.Print("\n")
    fmt.Print(DATE)
    fmt.Print("\n")
    os.Exit(0)
}

func info_print() {
    info()
    os.Exit(0)
}

func help_menu() {
    fmt.Print("usage: $ ", SCRIPT_NAME, " [OPTION] [OPTION] [OPTION] [MESSAGE]\n")
    fmt.Print("      -h | -help | --help\n")
    fmt.Print("            print usage menu)\n")
    fmt.Print("      -n [MESSAGE]\n")
    fmt.Print("            no new line\n")
    fmt.Print("      -r [OPTION] [OPTION] [MESSAGE]\n")
    fmt.Print("            dont reset the color.\n")
    fmt.Print("      -s [OPTION] [OPTION] [MESSAGE]\n")
    fmt.Print("            resets the color output.\n")
    fmt.Print("      -e [OPTION] [MESSAGE]\n")
    fmt.Print("            use the \\ options,\n")
    fmt.Print("            eg. $ ", SCRIPT_NAME, " -e \"your message\"\n")
    fmt.Print("      -f | -force [MESSAGE]\n")
    fmt.Print("            force print message, even if its -help\n")
    fmt.Print("      -red [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -yello [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -green [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -blue [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -white [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -teal [MESSAGE]\n")
    fmt.Print("            print message in givin color\n")
    fmt.Print("      -i | -info | --info\n")
    fmt.Print("            print script info\n")
    fmt.Print("      -v | -version | --version\n")
    fmt.Print("            print script version)\n")
    fmt.Print("Source code: https://github.com/WestleyK/golang-echo\n")
    os.Exit(0)
}

func fail(err1 string, err2 string) {
    if err2 == "help" {
        fmt.Print(red, "ERROR: ", colorReset, err1)
        fmt.Print("Try: $ ", SCRIPT_NAME, " -help  (for help menu)\n")
        os.Exit(1)
    }
    fmt.Print(red, "ERROR: ", colorReset, err1, err2)
    
    fmt.Print("\n")
    os.Exit(1)
}

func reset_color() {
    fmt.Print(colorReset)
    os.Exit(0)
}

func check_args(OPTION string) {
    if OPTION == "-h" || OPTION == "-help" || OPTION == "--help" {
        help_menu()
        return
    } else if OPTION == "-n" {
        noNewLine = 1
    } else if OPTION == "-e" {
        OPTION_E = 1
        return
    } else if OPTION == "-r" {
        NO_RESET = 1
        return
    } else if OPTION == "-s" {
        reset_color()
        return
    } else if OPTION == "-f" || OPTION == "-force" {
        return
    } else if OPTION == "-red" {
        color = "\x1b[31m"
        return
    } else if OPTION == "-yello" {
        color = "\x1b[33m"
        return
    } else if OPTION == "-green" {
        color = "\x1b[32m"
        return
    } else if OPTION == "-blue" {
        color = "\x1b[34m"
        return
    } else if OPTION == "-white" {
        color = "\x1b[35m"
        return
    } else if OPTION == "-teal" {
        color = "\x1b[36m"
        return
    } else if OPTION == "-i" || OPTION == "-info" || OPTION == "--info" {
        info_print()
        return
    } else if OPTION == "-v" || OPTION == "-version" || OPTION == "--version" {
        version_print()
        return
    } else if strings.Contains(OPTION, "-") {
        fail("option not found: ", OPTION)
        return
    }
}

func print_message(color string, MSG string) {
    if OPTION_E == 1 {
        fail("-e option is comming soon!\n", "option not avalible.")
    }
    if len(color) == 0 {
        fmt.Print(MSG)
        if noNewLine != 1 {
            fmt.Print("\n")
        }
        os.Exit(0)
    }

    fmt.Print(color, MSG)
    if NO_RESET != 1 {
        fmt.Print(colorReset)
    }
    
    if noNewLine != 1 {
        fmt.Print("\n")
    }
    os.Exit(0)
}

func one_args() {
    OPTION = os.Args[1]
    MSG = os.Args[1]
    check_args(OPTION)
    print_message("", MSG)
}

func two_args() {
    OPTION = os.Args[1]
    MSG = os.Args[2]
    check_args(OPTION)
    print_message(color, MSG)
}

func three_args() {
    OPTION = os.Args[1]
    OPTION2 = os.Args[2]
    MSG = os.Args[3]
    check_args(OPTION)
    check_args(OPTION2)
    print_message(color, MSG)
}

func four_args() {
    OPTION = os.Args[1]
    OPTION2 = os.Args[2]
    OPTION3 = os.Args[3]
    MSG = os.Args[4]
    check_args(OPTION)
    check_args(OPTION2)
    check_args(OPTION3)
    print_message(color, MSG)
}

func main() {
    SCRIPT_NAME = os.Args[0]

    if len(os.Args[1:]) == 0 {
        fmt.Print("\n")
        //fail("no arguments!\n", "need at least one argument")
    }

    arg_len := len(os.Args)

    if arg_len == 2 {
        one_args()
    } else if arg_len == 3 {
        two_args()
    } else if arg_len == 4 {
    	three_args()
    } else if arg_len == 5 {
        four_args()
    } else if arg_len >= 6 {
        fail("too many arguments!\n", "help")
    }

    return
}



//
// End source code
//
