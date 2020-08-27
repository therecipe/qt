package internal

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
	"unsafe"
)

type Internal struct {
	PtrI       uintptr `json:"___pointer"`
	ClassNameI string  `json:"___className"`
}

func (i *Internal) Pointer() uintptr {
	return i.PtrI
}

func (i *Internal) SetPointer(p uintptr) {
	i.PtrI = p
}

func (i *Internal) InitFromInternal(p uintptr, n string) {
	i.PtrI = p
	i.ClassNameI = n
}

func (i *Internal) ClassNameInternalF() string {
	return i.ClassNameI
}

var (
	callbackTable    = make(map[uintptr]map[string]interface{})
	ConstructorTable = make(map[string]interface{})
)

func registerFunctionInCallbackTable(ptr uintptr, fn string, f interface{}) {
	if _, ok := callbackTable[ptr]; !ok {
		callbackTable[ptr] = map[string]interface{}{fn: f}
	} else {
		callbackTable[ptr][fn] = f
	}
}

func deregisterFunctionInCallbackTable(ptr uintptr, fn string) {
	delete(callbackTable[ptr], fn)
}

func handleCallback(message string) string {
	//println("input:", message)
	var msg []interface{}
	json.Unmarshal([]byte(message), &msg)

	meth := reflect.ValueOf(callbackTable[uintptr(msg[0].(float64))][msg[1].(string)])

	rv := make([]reflect.Value, len(msg)-2)
	for i, v := range convertList(msg[2:]) {
		rv[i] = reflect.ValueOf(v).Convert(meth.Type().In(i))
	}

	var output []byte
	if ret := meth.Call(rv); len(ret) > 0 {
		output, _ = json.Marshal(convertToJson(ret[0].Interface()))
	}

	//println("output:", string(output))
	return string(output)
}

func httpRequest(url string, msg string) string {

	resp, err := http.Post("http://127.0.0.1:8000/"+url, "text", strings.NewReader(msg))
	if err != nil {
		fmt.Printf("httpRequest error: %v -> %v\n", msg, err.Error())
		return ""
	}

	bb := new(bytes.Buffer)
	io.Copy(bb, resp.Body)
	resp.Body.Close()

	return bb.String()
}

func syncCallIntoLocal(msg string) string {
	return httpRequest("syncCallIntoLocal", msg)
}

func asyncCallIntoRemoteResponse(msg string) {
	httpRequest("asyncCallIntoRemoteResponse", msg)
}

func syncCallIntoRemoteResponse(msg string) {
	httpRequest("syncCallIntoRemoteResponse", msg)
}

func asyncCallbackHandler(message string) {
	asyncCallIntoRemoteResponse(handleCallback(message))
}

func syncCallbackHandler(message string) {
	syncCallIntoRemoteResponse(handleCallback(message))
}

func convertList(l []interface{}) []interface{} {
	for i := range l {
		l[i] = convert(l[i])
	}
	return l
}

func convertMap(m map[string]interface{}) map[string]interface{} {
	for i := range m {
		m[i] = convert(m[i])
	}
	return m
}

func convert(l interface{}) interface{} {

	switch d := l.(type) {
	case map[string]interface{}:
		___className, okC := d["___className"]
		___pointer, okN := d["___pointer"]
		if okC && okN {
			return reflect.ValueOf(ConstructorTable[___className.(string)]).Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(___pointer.(float64))))})[0].Interface()
		}
		return convertMap(d)
	case []interface{}:
		return convertList(d)
	}

	return l
}

func convertListToJson(l []interface{}) []interface{} {
	for i := range l {
		l[i] = convertToJson(l[i])
	}
	return l
}

func convertMapToJson(l map[string]interface{}) map[string]interface{} {
	for k, v := range l {
		l[k] = convertToJson(v)
	}
	return l
}

func convertToJson(i interface{}) interface{} {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Map:
		return convertMapToJson(i.(map[string]interface{}))

	case reflect.Slice:
		switch i.(type) {
		//TODO:
		case []string:
		case []uint, []uint8, []uint16, []uint32, []uint64:
		case []int, []int8, []int16, []int32, []int64:
		//
		//case []*qml.QQmlError:

		default:
			return convertListToJson(i.([]interface{}))
		}

	case reflect.Ptr:
		if !reflect.ValueOf(i).IsNil() {
			return map[string]interface{}{
				"___pointer":   uintptr(reflect.ValueOf(i).MethodByName("Pointer").Call(nil)[0].Interface().(unsafe.Pointer)),
				"___className": reflect.ValueOf(i).MethodByName("ClassNameInternalF").Call(nil)[0].Interface(),
			}
		}
	}
	return i
}

var inited = false

func init() {
	if inited {
		return
	}
	inited = true
}

func functionName(fnDef string, fnOpt string) string {
	if strings.Contains(fnOpt, ":") {
		return strings.Split(fnOpt, ":")[1]
	}
	return fnDef
}

func CallLocalAndRegisterRemoteFunction(l []interface{}, f interface{}) interface{} {
	registerFunctionInCallbackTable(uintptr(l[1].(uintptr)), functionName(l[3].(string), l[4].(string)), f)
	return CallLocalFunction(l)
}

func CallLocalAndDeregisterRemoteFunction(l []interface{}) {
	CallLocalFunction(l)
	deregisterFunctionInCallbackTable(l[1].(uintptr), l[3].(string))
}

func CallLocalFunction(l []interface{}) interface{} {
	msg, _ := json.Marshal(convertToJson(l))
	var output interface{}
	json.Unmarshal([]byte(syncCallIntoLocal(string(msg))), &output)

	switch d := output.(type) {
	case []string: //TODO: dead code?
		if len(d) == 2 && d[0] == "___block" {
			return CallLocalFunction([]interface{}{"___return", handleCallback(d[1])})
		}

	case []interface{}:
		if s, ok := d[0].(string); ok && len(d) == 2 && s == "___block" {
			return CallLocalFunction([]interface{}{"___return", handleCallback(d[1].(string))})
		}
	}

	return convert(output)
}

type InteropServerConfig struct {
	Download    bool
	Override    bool
	Path        string
	DownloadUrl string
}

var Config = &InteropServerConfig{
	true,
	false,
	"",
	"",
}

var proc *exec.Cmd

func InitProcess() {

	var runPath string

	if Config.Download && Config.Path == "" {

		var ending string
		if runtime.GOOS == "windows" {
			ending = ".exe"
		}

		pwd, _ := os.Getwd()
		arg, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		for i, path := range []string{pwd, arg} {
			if i == 1 && pwd == arg {
				continue
			}

			runPath = filepath.Join(path, fmt.Sprintf("qtbox%v", ending))
			println("looking for qtbox in:", runPath)

			if strings.HasPrefix(runPath, "/private/var/folders/") { //app is quarantined on macOS
				runPath = filepath.Join(os.TempDir(), "qtbox")
				break
			}

			if strings.HasPrefix(runPath, "/var/folders/") ||
				strings.HasPrefix(runPath, "/tmp/go-build") {
				runPath = filepath.Join(pwd, "qtbox")
				break
			}

			if _, err := os.Stat(runPath); err == nil {
				break
			}
		}

		dst := filepath.Dir(runPath)
		_, err := os.Stat(runPath)
		f, errF := os.Stat(filepath.Join(dst, "qtbox"))
		if errF == nil && f.IsDir() {
			runPath = filepath.Join(dst, "qtbox", "qtbox"+ending)
		}

		println("final qtbox location:", runPath)

		if Config.Override || (err != nil && errF != nil) {

			var copyWithProgress = func(w io.Writer, r io.Reader, callback func(off int64)) error {
				tee := io.TeeReader(r, w)
				buf := make([]byte, bytes.MinRead*10)
				off := int64(0)
				for count := 0; ; count++ {
					n, err := tee.Read(buf)
					off += int64(n)
					if count%100 == 0 {
						callback(off)
					}

					if err == io.EOF {
						break
					}
					if err != nil {
						return err
					}
				}
				callback(off)
				return nil
			}

			dlpath := fmt.Sprintf("https://github.com/therecipe/box/releases/download/v0.0.0/%v_%v_%v_%v_%v.zip", runtime.GOOS, runtime.GOARCH, "513", "full", "http")
			if Config.DownloadUrl != "" {
				dlpath = Config.DownloadUrl
			}
			resp, err := http.Get(dlpath)
			if err != nil {
				println("failed to download qtbox:", err.Error())
			} else {

				bb := new(bytes.Buffer)
				copyWithProgress(bb, resp.Body, func(off int64) {
					fmt.Printf("downloading qtbox => %v%%\n", off/(resp.ContentLength/100))
				})
				resp.Body.Close()

				r, _ := zip.NewReader(bytes.NewReader(bb.Bytes()), int64(bb.Len()))
				if len(r.File) == 1 {
					fr, _ := r.File[0].Open()
					fw, _ := os.OpenFile(runPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, r.File[0].Mode())
					io.Copy(fw, fr)
					fr.Close()
					fw.Close()
				} else {

					for _, f := range r.File {

						//TODO: pack runtimes with correct name
						fns := strings.Split(f.Name, "/")
						fns[0] = "qtbox"
						f.Name = strings.Join(fns, "/")
						//

						if f.FileInfo().IsDir() {
							os.MkdirAll(filepath.Join(dst, f.Name), f.Mode())
							continue
						}

						dn, fn := filepath.Split(f.Name)
						if strings.HasPrefix(fn, "full") { //TODO: pack runtimes with correct name
							fn = filepath.Join(dn, "qtbox"+ending)
						} else {
							fn = f.Name
						}

						fw, errC := os.OpenFile(filepath.Join(dst, fn), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
						if errC != nil {
							println("failed to create file", errC.Error())
						}

						fr, errR := f.Open()
						if errR != nil {
							println("failed to read file", errR.Error())
						}

						io.Copy(fw, fr)

						fw.Close()
						fr.Close()
					}

					runPath = filepath.Join(dst, "qtbox", "qtbox"+ending)
				}
			}
		}
	}

	if Config.Path != "" {
		runPath = Config.Path
	}

	process := exec.Command(runPath)
	rc, err := process.StderrPipe()
	if err != nil {
		println(err.Error())
	}
	process.Start()

	time.Sleep(3 * time.Second) //TODO:

	scanner := bufio.NewScanner(rc)

	go func() {
		for scanner.Scan() {
			if l := scanner.Text(); strings.Contains(l, "async:") {
				go asyncCallbackHandler(strings.Split(l, "async:")[1])
			} else if strings.Contains(l, "sync:") {
				go syncCallbackHandler(strings.Split(l, "sync:")[1])
			} else {
				println(l)
			}
		}
		if err := scanner.Err(); err != nil {
			println("reading stderr:", err.Error())
		}
	}()

	proc = process
}

func Exec() {
	proc.Wait()
}
