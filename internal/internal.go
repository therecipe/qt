package internal

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os/exec"
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

	rv := make([]reflect.Value, len(msg)-2)
	for i, v := range convertList(msg[2:]) {
		rv[i] = reflect.ValueOf(v)
	}

	var output []byte
	if ret := reflect.ValueOf(callbackTable[uintptr(msg[0].(float64))][msg[1].(string)]).Call(rv); len(ret) > 0 {
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

func asyncCallbackHandler(message string) {
	asyncCallIntoRemoteResponse(handleCallback(message))
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
		return convertListToJson(i.([]interface{}))

	case reflect.Ptr:
		return map[string]interface{}{
			"___pointer":   uintptr(reflect.ValueOf(i).MethodByName("Pointer").Call(nil)[0].Interface().(unsafe.Pointer)),
			"___className": reflect.ValueOf(i).MethodByName("ClassNameInternalF").Call(nil)[0].Interface(),
		}

	default:
		return i
	}
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
	case []string:
		if len(l) == 2 && d[0] == "___block" {
			return CallLocalFunction([]interface{}{"___return", handleCallback(d[1])})
		}
	}

	return convert(output)
}

// TODO: NewQApplication
func InitProcess() (*exec.Cmd, io.ReadCloser) {

	//TODO: download qt server binaries
	switch runtime.GOOS {
	case "windows":
	case "darwin":
	case "linux":
	}

	process := exec.Command("./interop_server")
	rc, err := process.StderrPipe()
	if err != nil {
		println(err.Error())
	}
	process.Start()
	time.Sleep(3 * time.Second)
	return process, rc
}

// TODO: QApplication_Exec
func Exec(p *exec.Cmd, rc io.ReadCloser) {
	scanner := bufio.NewScanner(rc)

	go func() {
		for scanner.Scan() {
			if l := scanner.Text(); strings.Contains(l, "async:") {
				asyncCallbackHandler(strings.Split(l, "async:")[1])
			} else {
				println(l)
			}
		}
		if err := scanner.Err(); err != nil {
			println("reading stderr:", err.Error())
		}
	}()

	p.Wait()
}
