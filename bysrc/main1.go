package main

import (
	"flag"
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var fname string

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		log.Fatalln("must specify a go source file to tranpiler")
	}
	fname = os.Args[1]
	fio, err := os.Lstat(fname)
	gopp.ErrPrint(err)
	if err != nil {
		return
	}
	if !fio.IsDir() {
		log.Fatalln("Not a dir", fname)
	}
	psctx := NewParserContext(fname, "")
	err = psctx.Init()
	gopp.ErrPrint(err, fname)
	if err != nil {
		os.Exit(-1)
	}
	psctx.walkpass()

	// g2n := g2nim{}
	g2n := g2nc{}
	g2n.basecomp = newbasecomp(psctx)
	g2n.genpkgs()
	code, ext := g2n.code()
	dstfile := psctx.bdpkgs.Name + ".go." + ext
	ioutil.WriteFile("opkgs/"+dstfile, []byte(code), 0644)

	outfile := "opkgs/foo." + ext
	ioutil.WriteFile(outfile, []byte(code), 0644)
	clangfmt(outfile)
}

/*
func maindep() {
	if len(os.Args) < 2 {
		log.Fatalln("must specify a go source file to tranpiler")
	}
	fname = os.Args[1]
	fio, err := os.Lstat(fname)
	gopp.ErrPrint(err)
	if err != nil {
		return
	}
	if !fio.IsDir() {
		log.Fatalln("Not a dir", fname)
	}

	pkgpaths := []string{fname}
	psctxs := []*ParserContext{}
	comps := []*g2nc{}
	pkgrenames := map[string]string{} // path => rename

	gopath := os.Getenv("GOPATH")
	gopaths := strings.Split(gopath, ":")
	gopaths = append(gopaths, runtime.GOROOT())

	for len(pkgpaths) > 0 {
		fname := pkgpaths[0]
		pkgpaths = pkgpaths[1:]
		pkgrename := ""
		segs := strings.Split(fname, ":")
		if len(segs) == 2 {
			fname = segs[0]
			pkgrename = segs[1]
		}

		psctx, g2n := dogen(fname, pkgrename)
		psctxs = append(psctxs, psctx)
		comps = append(comps, g2n)

		imprenames := psctx.getImportNameMap()
		for path, rename := range imprenames {
			pkgrenames[path] = rename
			log.Println(path, rename)
		}

		for _, imppath := range psctx.bdpkgs.Imports {
			log.Println(imppath, psctx.bdpkgs.Dir)
			if imppath == "runtime" ||
				imppath == "atomic" ||
				imppath == "runtime/cgo" ||
				imppath == "syscall" || imppath == "syscall/js" ||
				imppath == "internal/race" {
				continue
			}
			for _, gopath1 := range gopaths {
				impdir := gopath1 + "/src/" + imppath
				if gopp.FileExist(impdir) {
					log.Println("got", impdir)
					pkgpaths = append(pkgpaths, impdir+":"+pkgrenames[imppath])
					break
				}
			}
		}
		log.Println("=================", fname)
	}

	code := ""
	extname := ""
	for i := len(comps) - 1; i >= 0; i-- {
		str, ext := comps[i].code()
		code += str
		extname = ext
	}
	fname := "opkgs/foo." + extname
	ioutil.WriteFile(fname, []byte(code), 0644)
	clangfmt(fname)
}
*/

func clangfmt(fname string) {
	cmdo := exec.Command("clang-format", "-i", fname)
	err := cmdo.Run()
	gopp.ErrPrint(err, fname)
}
func dogen(fname string, pkgrename string) (*ParserContext, *g2nc) {
	psctx := NewParserContext(fname, pkgrename)
	err := psctx.Init()
	if err != nil && !strings.Contains(err.Error(), "declared but not used") {
		gopp.ErrPrint(err)
		println()
		println()
		time.Sleep(1 * time.Second)
	}

	// g2n := g2nim{}
	g2n := g2nc{}
	g2n.basecomp = newbasecomp(psctx)
	g2n.genpkgs()
	code, ext := g2n.code()
	dstfile := psctx.bdpkgs.Name + ".go." + ext
	ioutil.WriteFile("opkgs/"+dstfile, []byte(code), 0644)
	return psctx, &g2n
}
