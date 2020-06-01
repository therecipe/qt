package interop

import (
	"encoding/json"
	"strings"

	"github.com/therecipe/qt/core"
)

var (
	syncCallChan  = make(chan string, 0)
	asyncCallChan = make(chan string, 0)

	mainThreadIsBlockedAlready bool

	SyncCallIntoRemote  func(string) string
	AsyncCallIntoRemote func(string)
)

func RunOnMainBlockingWithWorkerQueue(f func()) {
	done := make(chan bool, 0)
	Helper.RunOnMainThread(func() { f(); done <- true })
	for {
		select {
		case s := <-syncCallChan:
			syncCallChan <- SyncCallIntoRemote(s)
		case <-done:
			goto br
		}
	}
br:
}

//
//
//

var Helper = NewHelper(nil)

type helper struct {
	core.QObject
	_ func()         `constructor:"init"`
	_ func(f func()) `slot:"runOnMainThread"`
}

func (h *helper) init() {
	h.ConnectRunOnMainThread(func(f func()) { f() })
}

//
//
//

func _syncCallIntoLocal(s string) string {

	//println("input:", s)
	var i []interface{}
	json.Unmarshal([]byte(s), &i)

	pseudoIn := NewPseudoQJSValue1(i)
	for j := uint(0); j < pseudoIn.Property("length").ToUInt(nil); j++ {
		if v := pseudoIn.Property2(j); strings.HasPrefix(v.ToString(), "___REMOTE_CALLBACK___") {
			f := map[string]interface{}{
				"___pointer":    pseudoIn.Property2(1).ToULongLong(nil),
				"callable":      true,
				"callableLocal": false,
				"callableName":  pseudoIn.Property2(3).ToString(),
			}
			if vs := v.ToString(); strings.Contains(vs, ":") {
				f["callableName"] = strings.Split(vs, ":")[1]
			}
			fun := NewPseudoQJSValue1(f)
			pseudoIn.SetProperty2(j, fun)

			//TODO: allow creation of funcs in arbitrary module depth
			if pseudoIn.Property2(3).ToString() == "NewGoType" {
				ptr := PseudoQJSEngine_qjsEngine(nil)
				fn := f["callableName"].(string)

				if strings.Count(fn, ".") == 0 {
					ptr.GlobalObject().SetProperty(fn, fun)
				} else {

					var jsv *PseudoQJSValue
					if m := ptr.GlobalObject().Property(strings.Split(fn, ".")[0]); m.IsUndefined() {
						jsv = ptr.NewObject()
						//ptr.GlobalObject().SetProperty(strings.Split(fn, ".")[0], jsv) //TODO: look below
					} else {
						jsv = m
					}

					//TODO: PseudoQJSValue should instead automatically update it's parent(s) instead, so it behaves similar to qml.QJSValue
					jsv.SetProperty(strings.Split(fn, ".")[1], fun)
					ptr.GlobalObject().SetProperty(strings.Split(fn, ".")[0], jsv)
				}
			}
		}
	}

	var ret *PseudoQJSValue
	RunOnMainBlockingWithWorkerQueue(func() {
		mainThreadIsBlockedAlready = true
		ret = Z_wrapperFunction(pseudoIn)
		mainThreadIsBlockedAlready = false
	})

	var o []byte
	if ret == nil {
		var someOutput []interface{}
		o, _ = json.Marshal(someOutput)
	} else {
		eg := PseudoQJSEngine_qjsEngine(nil)

		//TODO: use reflection and/or recursion instead
		switch ret.Type() {
		case core.QVariant__Map:

			var typ core.QVariant__Type
			for _, v := range ret.ToMap() {
				typ = v.Type()
				break
			}

			switch typ {
			case core.QVariant__Map:
				var someOutput map[string]map[string]interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)

			case core.QVariant__List:
				var someOutput map[string][]interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)

			default:
				var someOutput map[string]interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)
			}

		case core.QVariant__List:

			var typ core.QVariant__Type
			for _, v := range ret.ToList() {
				typ = v.Type()
				break
			}

			switch typ {
			case core.QVariant__Map:
				var someOutput []map[string]interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)

			case core.QVariant__List:
				var someOutput [][]interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)

			default:
				var someOutput []interface{}
				eg.ToGoType(ret, &someOutput)
				o, _ = json.Marshal(someOutput)
			}

		default:
			var someOutput interface{}
			eg.ToGoType(ret, &someOutput)
			o, _ = json.Marshal(someOutput)
		}
		//<<<
	}

	//println("output:", string(o))
	return string(o)
}

func _asyncCallIntoRemoteResponse(s string) { asyncCallChan <- s }
