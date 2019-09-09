package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	flag "github.com/spf13/pflag"
)

func main() {
	s_num := flag.IntP("start-tpage", "s", 0, "the start page to print [Required]")
	e_num := flag.IntP("end-page", "e", 0, "the end page to print [Required]")
	l_num := flag.IntP("line-number", "l", 72, "specify the number of lines in each page")
	f_sep := flag.BoolP("force-page-break", "f", false, "specify if pages will be forcely seperated by '\\f'")
	d_dest := flag.StringP("destination", "d", "", "specify if the output will be sent to printer")

	elog := log.New(os.Stderr, "", 0)


	flag.Parse()

	if *s_num <= 0 || *e_num <= 0 {
		elog.Println("positive start and end page required!")
		flag.Usage()
		os.Exit(1)
	}

	if *s_num > *e_num {
		elog.Println("start page cannot be larger than end page!")
		flag.Usage()
		os.Exit(1)
	}

	if *l_num != 72 && *f_sep {
		elog.Println("-d and -l cannot be set at the same time!")
		flag.Usage()
		os.Exit(1)
	}

	if *l_num <= 0 {
		elog.Println("line number should be positive!")
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() > 1 {
		elog.Println("too many targets!")
		flag.Usage()
		os.Exit(1)
	}


	buf := make([]byte, 2333)
	var data string
	var reader io.Reader
	if flag.NArg() == 0 {
		reader = os.Stdin
	} else {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			elog.Println("error when read file!")
			os.Exit(1)
		}
		reader = file
	}
	buf_reader := bufio.NewReader(reader)
	size, err := buf_reader.Read(buf)
	for size != 0 && err == nil {
		data = data + string(buf)
		size, err = buf_reader.Read(buf)
	}

	if err != io.EOF {
		elog.Println("error when read", err)
		os.Exit(1)
	}
	var result string

	if *f_sep {
		page_data := strings.Split(data, "\f")

		if len(page_data) < *e_num {
			elog.Printf("page number too large, only %d pages available!\n", len(page_data))
			flag.Usage()
			os.Exit(1)
		}

		result = strings.Join(page_data[*s_num - 1: *e_num], "\n")
	} else {
		lines := strings.Split(data, "\n")
		n_pages := (len(lines) + *l_num - 1) / *l_num

		if *e_num > n_pages {
			elog.Printf("page number too large, only %d pages available!\n", n_pages)
			flag.Usage()
			os.Exit(1)
		}

		if *e_num * *l_num <= len(lines) {
			result = strings.Join(lines[(*s_num-1)*(*l_num):(*e_num)*(*l_num)], "\n")
		} else {
			result = strings.Join(lines[(*s_num-1)*(*l_num):], "\n")
		}
	}


	if *d_dest == "" {
		fmt.Print(result)
	} else {
		cmd := exec.Command("lp", "-d" + *d_dest)
		cmd_in, _ := cmd.StdinPipe()
		go func() {
			defer cmd_in.Close()
			io.WriteString(cmd_in, result)
		}()

		out, err := cmd.CombinedOutput()
		fmt.Print(string(out))
		if err != nil {
			elog.Println("error when execute command lp!")
			os.Exit(1)
		}
		fmt.Println("success!")
	}

}
