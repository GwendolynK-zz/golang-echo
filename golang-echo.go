// created by: WestleyK
// date: Sep 18, 2018
// email: westleyk@nym.hush.com
//
// MIT Licence
//



package main


import (
    "fmt"
    "os"
    "strings"
)

var (
    VERSION string = "version-1.0.0"
    DATE string = "Date: Sep 18, 2018"

    OPTION string = ""
    OPTION2 string = ""
    SCRIPT_NAME string = ""
    noNewLine int = 0
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
    fmt.Print("usage: $ ", SCRIPT_NAME, " [OPTION] [OPTION]\n")
    fmt.Print("      -h | -help | --help (print usage menu)\n")
    fmt.Print("      -n [MESSAGE] (no new line)\n")
    fmt.Print("      -red [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -yello [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -green [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -blue [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -white [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -teal [MESSAGE] (print message in givin color)\n")
    fmt.Print("      -i | -info | --info (print script info)\n")
    fmt.Print("      -v | -version | --version (print script version)\n")
    fmt.Print("Source code: ????????????????????\n")
    os.Exit(0)
}

func fail(err1 string, err2 string) {
    if err2 == "help" {
        fmt.Print(red, "ERROR: ", colorReset, err1)
        fmt.Print("Try: $ ", SCRIPT_NAME, "-help  (for help menu)\n")
        os.Exit(1)
    }
    fmt.Print(red, "ERROR: ", colorReset, err1, err2)
    
    fmt.Print("\n")
    os.Exit(1)
}

func check_args(OPTION string) {
    if OPTION == "-h" || OPTION == "-help" || OPTION == "--help" {
        help_menu()
        return
    } else if OPTION == "-n" {
        noNewLine = 1
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
        fail("option not found!  ", OPTION)
        return
    }
}

func print_message(color string, MSG string) {
    if len(color) == 0 {
        fmt.Print(MSG)
        if noNewLine != 1 {
            fmt.Print("\n")
        }
        os.Exit(0)
    }

    fmt.Print(color, MSG, colorReset)
    
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

func main() {
    SCRIPT_NAME = os.Args[0]

    if len(os.Args[1:]) == 0 {
        fail("no arguments!\n", "need at least one argument")
    }


    arg_len := len(os.Args)


    if arg_len == 2 {
        one_args()
    } else if arg_len == 3 {
        two_args()
    } else if arg_len == 4 {
    	three_args()
    } else if arg_len >= 5 {
        fail("too many arguments!\n", "help")
    }


    return
}




//
// End source code
//
