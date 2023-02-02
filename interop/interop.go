package interop

import (
	"encoding/json"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"unsafe"

	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
)

var (
	Z_newEngineHelper func(p core.QObject_ITF) *core.QObject

	finalizerMap      = make(map[unsafe.Pointer][]*core.QObject)
	finalizerMapMutex sync.Mutex

	//needed only for interop  --->
	ReturnPointersAsStrings     bool
	SupportsSyncCallsIntoRemote = true

	syncCallIntoRemoteChan       = make(chan string, 0)
	syncCallIntoRemoteReturnChan = make(chan string, 0)

	fromJsToRefReturnsOnChan bool
	fromJsToRefReturnChan    = make(chan []reflect.Value, 0)

	Z_wrapperFunctionReturnChan = make(chan *PseudoQJSValue, 0)
	//needed only for interop  <---
)

type ptr_itf interface {
	Pointer() unsafe.Pointer
}

func Z_initEngine(engine PseudoQJSEngine_ITF) {
	ptr := engine.PseudoQJSEngine_PTR()
	ptr.QObject.ConnectDestroyed(func(ptr *core.QObject) {
		finalizerMapMutex.Lock()
		pointers := finalizerMap[ptr.Pointer()]
		delete(finalizerMap, ptr.Pointer())
		finalizerMapMutex.Unlock()

		for _, o := range pointers {
			o.DestroyQObject()
		}
		runtime.GC()
	})
	if !ptr.GlobalObject().Property("___engineHelper").IsUndefined() || Z_newEngineHelper == nil {
		return
	}
	ptr.GlobalObject().SetProperty("nil", NewPseudoQJSValue(PseudoQJSValue__UndefinedValue))
	helper := Z_newEngineHelper(engine)
	ptr.GlobalObject().SetProperty("___engineHelper", ptr.NewQObject(helper))
	ptr.GlobalObject().SetProperty("___engineHelperPointer", ptr.NewGoType(helper.Pointer()))
	ptr.GlobalObject().SetProperty("___factory_single_func", ptr.Evaluate(`
(function(cn, fn) {
	return function() {
		var callArgs = [___engineHelperPointer, this.___pointer, cn, fn];
		for (var i = 0; i < arguments.length; i++) { callArgs.push(arguments[i]) }
		return ___engineHelper.wrapperFunction(callArgs);
	};
});
`, "", 0))

	ptr.GlobalObject().SetProperty("___factory_batch_funcs", ptr.Evaluate(`
(function(that, cn, fns) {
	fns.forEach(function(fn) {
		that[fn] = ___factory_single_func(cn, fn);
	})
});
`, "", 0))

	ptr.GlobalObject().SetProperty("___factory_batch_global_funcs", ptr.Evaluate(`
(function(that, pfx, fns) {
	fns.forEach(function(fn) {
		that[fn] = ___factory_single_func(pfx+"."+fn, "");
	})
});
`, "", 0))

	ptr.GlobalObject().SetProperty("___factory_batch_global_enums", ptr.Evaluate(`
(function(that, ens) {
	ens.forEach(function(en) {
		that[en[0]] = en[1];
	})
});
`, "", 0))

	ptr.GlobalObject().SetProperty("___connectDestroyed", ptr.Evaluate(`
(function(that) {
	that.ConnectDestroyed(function(){ that.___pointer = nil })
});
`, "", 0))

	prefE := make(map[string][][]interface{})
	qt.EnumMapMutex.Lock()
	for en, ev := range qt.EnumMap {
		pfx := strings.Split(en, ".")
		prefE[pfx[0]] = append(prefE[pfx[0]], []interface{}{pfx[1], ev})
	}
	qt.EnumMapMutex.Unlock()
	for pfx, ens := range prefE {
		var jsv *PseudoQJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		inp := ptr.NewGoType(ens)
		ptr.GlobalObject().Property("___factory_batch_global_enums").Call([]*PseudoQJSValue{jsv, inp})
		inp.DestroyPseudoQJSValue()
	}

	pref := make(map[string][]string)
	qt.FuncMapMutex.Lock()
	for fn := range qt.FuncMap {
		pfx := strings.Split(fn, ".")
		if len(pfx) == 2 { //TODO: re-register 3rdparty functions?
			pref[pfx[0]] = append(pref[pfx[0]], pfx[1])
		}
	}
	qt.FuncMapMutex.Unlock()
	for pfx, fns := range pref {
		var jsv *PseudoQJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		inp := ptr.NewGoType(fns)
		ptr.GlobalObject().Property("___factory_batch_global_funcs").Call([]*PseudoQJSValue{jsv, NewPseudoQJSValue8(pfx), inp})
		inp.DestroyPseudoQJSValue()
	}
}

func Z_wrapperFunction(jsvals *PseudoQJSValue) *PseudoQJSValue {

	var m reflect.Value
	cn, fn := jsvals.Property2(2).ToString(), jsvals.Property2(3).ToString()
	if fn == "" {
		v, _ := qt.GetFuncMap(cn)
		m = reflect.ValueOf(v)

	} else {
		var rt reflect.Type
		if t, ok := qt.GetItfMap(cn + "_ITF"); ok {
			rt = reflect.TypeOf(t)
		} else {
			t, _ = qt.GetItfMap(cn)
			rt = reflect.TypeOf(t)
		}

		//TODO: use ptr_itf instead ?
		//TODO: use *FromPointer instead
		//TODO: make SetPointer work as backup, for classes that don't support *FromPointer -> needs manual registration like for *FromPointer

		o := reflect.New(rt)
		o.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(jsvals.Property2(1).ToVariant().ToULongLong(nil))))})
		m = o.MethodByName(fn)
	}

	engine := PseudoQJSEngine_qjsEngine(core.NewQObjectFromPointer(unsafe.Pointer(uintptr(jsvals.Property2(0).ToVariant().ToULongLong(nil)))))

	if fn == "NewGoType" || (fn == "NewQVariant1" && cn == "core.QVariant") {

		input := make([]interface{}, jsvals.Property("length").ToInt()-4)
		for i := 0; i < len(input); i++ {

			in := jsvals.Property2(uint(4 + i))

			switch in.ToVariant().Type() {
			case core.QVariant__List:
				var o []interface{}
				engine.ToGoType(in, &o)
				input[i] = o

			case core.QVariant__Map:
				var o map[string]interface{}
				engine.ToGoType(in, &o)
				input[i] = o

			default:
				input[i] = in.ToVariant().ToInterface()
			}
		}

		if fn == "NewQVariant1" {
			return engine.NewGoType(core.NewQVariant1(input[0]))
		}
		if cn != "interop.PseudoQJSEngine" {
			refInput := make([]reflect.Value, len(input))
			for i, v := range input {
				if jsv := jsvals.Property2(uint(4 + i)); jsv.IsCallable() {
					refInput[i] = reflect.ValueOf(func(i interface{}) interface{} { //TODO: remote needs to provide infos about the expected arguments (needed for the callbacks for Qml created in the remote)
						return jsv.Call([]*PseudoQJSValue{engine.NewGoType(i)}).ToVariant().ToInterface()
					})
				} else {
					refInput[i] = reflect.ValueOf(v)
				}
			}
			return engine.NewGoType(m.Call(refInput)[0].Interface())
		}
		return engine.NewGoType(engine.NewGoType(input...))

	} else {

		input := make([]reflect.Value, m.Type().NumIn())
		for i := 0; i < len(input); i++ {
			input[i] = engine.fromJsToRef(m.Type().In(i), jsvals.Property2(uint(4+i)))
		}

		ret := m.Call(input)

		//needed only for interop  --->
		if fromJsToRefReturnsOnChan {
			fromJsToRefReturnsOnChan = false
			go func() {
				ret = <-fromJsToRefReturnChan

				if len(ret) == 0 {
					Z_wrapperFunctionReturnChan <- NewPseudoQJSValue(PseudoQJSValue__UndefinedValue)
					return
				}

				rret := engine.NewGoType(ret[0].Interface())
				if reflect.TypeOf(ret[0].Interface()).Implements(reflect.TypeOf((*core.QObject_ITF)(nil)).Elem()) { //TODO: check for destroyed signal instead, or simply override the destructor instead ?
					if qt.ExistsSignal(ret[0].Interface().(ptr_itf).Pointer(), "destroyed") {
						engine.GlobalObject().Property("___connectDestroyed").Call([]*PseudoQJSValue{rret}) //TODO: connect destroyed/destructor from go instead ?
					}
				}

				Z_wrapperFunctionReturnChan <- rret
			}()
			return NewPseudoQJSValue(PseudoQJSValue__UndefinedValue)
		}
		//needed only for interop  <---

		if len(ret) == 0 {
			return NewPseudoQJSValue(PseudoQJSValue__UndefinedValue)
		}

		rret := engine.NewGoType(ret[0].Interface())
		if reflect.TypeOf(ret[0].Interface()).Implements(reflect.TypeOf((*core.QObject_ITF)(nil)).Elem()) { //TODO: check for destroyed signal instead, or simply override the destructor instead ?
			if qt.ExistsSignal(ret[0].Interface().(ptr_itf).Pointer(), "destroyed") {
				engine.GlobalObject().Property("___connectDestroyed").Call([]*PseudoQJSValue{rret}) //TODO: connect destroyed/destructor from go instead ?
			}
		}

		return rret
	}
}

func (ptr *PseudoQJSEngine) ToGoType(jsval *PseudoQJSValue, dst interface{}) {
	out := ptr.fromJsToRef(reflect.TypeOf(dst), jsval)
	if out.Type().Kind() == reflect.Ptr {
		out = out.Elem()
	}
	reflect.ValueOf(dst).Elem().Set(out)
}

func (ptr *PseudoQJSEngine) fromJsToRef(tofi reflect.Type, jsval *PseudoQJSValue) reflect.Value {
	tofiO := tofi

	if tofi.Kind() == reflect.Ptr {
		tofi = tofi.Elem()
	}

	t, ok := qt.GetItfMap(tofi.String() + "_ITF")
	if ok {
		tofi = reflect.TypeOf(t)
	} else {
		t, ok = qt.GetItfMap(tofi.String())
		if ok {
			tofi = reflect.TypeOf(t)
		}
	}

	switch {
	case tofi.Kind() == reflect.UnsafePointer:
		return reflect.ValueOf(unsafe.Pointer(uintptr(jsval.ToVariant().ToULongLong(nil))))

	case tofi.Kind() == reflect.Uintptr:
		return reflect.ValueOf(uintptr(jsval.ToVariant().ToULongLong(nil)))

	case tofi.Kind() == reflect.Func:
		return reflect.MakeFunc(tofi, func(args []reflect.Value) []reflect.Value {
			input := make([]*PseudoQJSValue, len(args))
			for i, arg := range args {
				input[i] = ptr.NewGoType(arg.Interface())
			}
			ret := jsval.Call(input)

			//the default path to return on this function
			if !(ret.IsObject() && ret.Property("___earlyReturn").ToString() == "true") {
				if tofi.NumOut() != 0 {
					return []reflect.Value{ptr.fromJsToRef(tofi.Out(0), ret)}
				}
				return nil
			}

			//TODO: some functions can't return on a chan such as core.QAbstractListModel::RowCount (when called i.e. by the core.QAbstractListModel::Index method); warn about these methods here if the remote lanuage is not supporting "SupportsSyncCallsIntoRemote" ?

			//needed only for interop  --->
			fromJsToRefReturnsOnChan = true

			syncCallIntoRemoteChan <- string(ret.Property("___data").ToString())
			go func() {
				var o interface{}
				json.Unmarshal([]byte(<-syncCallIntoRemoteReturnChan), &o)
				if tofi.NumOut() != 0 {
					fromJsToRefReturnChan <- []reflect.Value{ptr.fromJsToRef(tofi.Out(0), ptr.ToScriptValue(core.NewQVariant1(o)))}
				} else {
					fromJsToRefReturnChan <- nil
				}
			}()

			if tofi.NumOut() != 0 {
				return []reflect.Value{reflect.Zero(tofi.Out(0))} //TODO: (SupportsSyncCallsIntoRemote related); possible to trigger panic's with something like []reflect.Value{reflect.ValueOf("___earlyReturn")}
			}
			return nil
			//needed only for interop  <---
		})

		//TODO: merge into (*core.QVariant).ToGoType ? >>>
	case tofiO.ConvertibleTo(reflect.TypeOf(int8(0))):
		switch tofi.Kind() {
		case reflect.Float32, reflect.Float64:
			return reflect.ValueOf(jsval.ToVariant().ToDouble(nil)).Convert(tofi)
		}
		return reflect.ValueOf(jsval.ToVariant().ToLongLong(nil)).Convert(tofi)
		//<<<

	case tofi.Kind() == reflect.Slice, tofi.Kind() == reflect.Array:
		var rv reflect.Value
		if ls := jsval.Property("length").ToInt(); tofi.Kind() == reflect.Slice {
			rv = reflect.MakeSlice(tofi, ls, ls)
		} else {
			rv = reflect.New(tofi).Elem()
		}
		for is := 0; is < rv.Len(); is++ {
			rv.Index(is).Set(ptr.fromJsToRef(tofi.Elem(), jsval.Property2(uint(is))))
		}
		return rv

	case tofi.Kind() == reflect.Map:
		rv := reflect.MakeMapWithSize(tofi, jsval.Property("length").ToInt())
		for k, _ := range jsval.ToVariant().ToMap() { //TODO: get keys from js function instead ?
			rv.SetMapIndex(reflect.ValueOf(k).Convert(tofi.Key()), ptr.fromJsToRef(tofi.Elem(), jsval.Property(k)))
		}
		return rv

	case ok:
		if tofi.String() == "interop.PseudoQJSValue" {
			if className := jsval.Property("___className").ToString(); className != "interop.PseudoQJSValue" {
				t, ok := qt.GetItfMap(className + "_ITF")
				if ok {
					tofi = reflect.TypeOf(t)
				} else {
					t, ok = qt.GetItfMap(className)
					if ok {
						tofi = reflect.TypeOf(t)
					}
				}
				if ok {
					rv := reflect.New(tofi)
					rv.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(jsval.Property("___pointer").ToVariant().ToULongLong(nil))))})
					return reflect.ValueOf(ptr.NewGoType(rv.Interface()))
				}
			}
		}

		if tofi.String() == "core.QVariant" && jsval.IsUndefined() {
			return reflect.ValueOf(core.NewQVariant())
		}

		//TODO: use ptr_itf instead ?
		//TODO: use *FromPointer instead
		//TODO: make SetPointer work as backup, for classes that don't support *FromPointer -> needs manual registration like for *FromPointer
		rv := reflect.New(tofi)
		rv.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(jsval.Property("___pointer").ToVariant().ToULongLong(nil))))})
		return rv

	case tofi.Kind() == reflect.Struct && !tofiO.Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()): //TODO: support pure go types (i.e. types that don't support ptr_itf as well)
		rv := reflect.New(tofi).Elem()
		for k, _ := range jsval.ToVariant().ToMap() { //TODO: get keys from js function instead ?
			val := jsval.Property(k)
			realName := k
			toInt := false
			for id := 0; id < rv.NumField(); id++ {
				if tag, ok := rv.Type().Field(id).Tag.Lookup("json"); ok {
					switch {
					case strings.Count(tag, ",") == 0:
						if k == tag {
							realName = rv.Type().Field(id).Name
						}
					default:
						tags := strings.Split(tag, ",")
						if k == tags[0] || (len(tags[0]) == 0 && k == rv.Type().Field(id).Name) {
							realName = rv.Type().Field(id).Name
							if tags[1] == "string" {
								toInt = true
							}
						}
					}
				}
			}
			field := rv.FieldByName(realName)
			if !field.IsValid() {
				continue
			}
			if toInt {
				field.Set(reflect.ValueOf(val.ToVariant().ToLongLong(nil)).Convert(field.Type())) //TODO: pure go type conversion
			} else {
				field.Set(ptr.fromJsToRef(field.Type(), val))
			}
		}
		if tofiO.Kind() == reflect.Ptr {
			return rv.Addr()
		}
		return rv

	case jsval.IsCallable():
		return reflect.ValueOf(jsval)

	case jsval.IsUndefined(), jsval.IsNull():
		return reflect.Zero(tofi)
	}

	ii := reflect.New(tofi)
	jsval.ToVariant().ToGoType(ii.Interface()) //TODO: does ToGoType support QObject or ptr_itf types inside maps/slices/structs ?

	if tofiO.Kind() == reflect.Ptr {
		return reflect.ValueOf(ii.Interface())
	}

	return reflect.ValueOf(ii.Elem().Interface())
}

//TODO: NewPseudoQJSValue1

func (ptr *PseudoQJSEngine) NewJSType(property *PseudoQJSValue, name string, i *PseudoQJSValue) { //TODO: support for js functions
	ptr.newGoType(property, name, i)
}

func (ptr *PseudoQJSEngine) NewGoType(i ...interface{}) *PseudoQJSValue {
	return ptr.newGoType(nil, "", i...)
}

func (ptr *PseudoQJSEngine) newGoType(property *PseudoQJSValue, name string, i ...interface{}) *PseudoQJSValue {
	rv := reflect.ValueOf(i[0])
	if len(i) > 1 {
		rv = reflect.ValueOf(i[1])
	}

	//TODO: merged into core.NewQVariant1; can probably be removed after testing >>>
	switch rv.Kind() {

	case reflect.UnsafePointer:
		return ptr.ToScriptValue(core.NewQVariant1(uint64(uintptr(rv.Interface().(unsafe.Pointer)))))
	case reflect.Uintptr:
		return ptr.ToScriptValue(core.NewQVariant1(uint64(rv.Interface().(uintptr))))
	}

	if rv.Type().ConvertibleTo(reflect.TypeOf(int8(0))) {
		if rv.Kind() == reflect.Int64 {
			return ptr.ToScriptValue(core.NewQVariant1(rv.Convert(reflect.TypeOf(int64(0))).Interface()))
		}
		return ptr.ToScriptValue(core.NewQVariant1(rv.Interface()))
	}
	//<<<

	switch rv.Kind() {
	case reflect.Func:

		var fn string
		if len(i) > 1 {
			fn = i[0].(string)
		} else {
			fn = runtime.FuncForPC(rv.Pointer()).Name()
			if strings.Contains(fn, "/") {
				fns := strings.Split(fn, "/")
				fn = fns[len(fns)-1]
			}
		}

		if _, ok := qt.GetFuncMap(fn); !ok {
			qt.SetFuncMap(fn, rv.Interface())
		}

		return ptr.makeFuncWrapper(fn)

	case reflect.Slice, reflect.Array:
		jsv := ptr.NewArray(uint(rv.Len()))
		for i := 0; i < rv.Len(); i++ {
			inp := ptr.NewGoType(rv.Index(i).Interface())
			jsv.SetProperty2(uint(i), inp)
			inp.DestroyPseudoQJSValue()
		}
		return jsv

	case reflect.Map:
		jsv := ptr.NewObject()
		for _, k := range rv.MapKeys() {
			if key := core.NewQVariant1(k.Interface()); key.CanConvert(int(core.QMetaType__QString)) { //TODO: replace with pure go (int -> string) conversion
				jsv.SetProperty(key.ToString(), ptr.NewGoType(rv.MapIndex(k).Interface()))
			}
		}
		return jsv

	case reflect.Struct:
		jsv := ptr.NewObject()
		for id := 0; id < rv.NumField(); id++ {
			if key := core.NewQVariant1(rv.Type().Field(id).Name); key.CanConvert(int(core.QMetaType__QString)) { //TODO: replace with pure go (int -> string) conversion
				field := rv.Field(id)
				if !field.CanInterface() {
					continue
				}
				if tag, ok := rv.Type().Field(id).Tag.Lookup("json"); ok {
					switch {
					case (strings.HasSuffix(tag, ",omitempty") && isZero(field)) || tag == "-":
					case strings.Count(tag, ",") == 0:
						jsv.SetProperty(tag, ptr.NewGoType(field.Interface()))
					default:
						tags := strings.Split(tag, ",")
						name := rv.Type().Field(id).Name
						if len(tags[0]) != 0 {
							name = tags[0]
						}
						v := ptr.NewGoType(field.Interface())
						if tags[1] == "string" {
							v = ptr.NewGoType(v.ToString()) //TODO: pure go type conversion
						}
						jsv.SetProperty(name, v)
					}
				} else {
					jsv.SetProperty(rv.Type().Field(id).Name, ptr.NewGoType(field.Interface()))
				}
			}
		}
		return jsv
	}

	var jsv *PseudoQJSValue
	if reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()) {
		jsv = ptr.NewObject()
	} else {
		jsv = ptr.ToScriptValue(core.NewQVariant1(i[0]))
	}

	if rv.Type().String() == "*interop.PseudoQJSValue" {
		className := rv.Interface().(*PseudoQJSValue).Property("___className").ToString()
		if className != "interop.PseudoQJSValue" {
			var tofi reflect.Type
			t, ok := qt.GetItfMap(className + "_ITF")
			if ok {
				tofi = reflect.TypeOf(t)
			} else {
				t, ok = qt.GetItfMap(className)
				if ok {
					tofi = reflect.TypeOf(t)
				}
			}
			if ok {
				rvn := reflect.New(tofi)
				rvn.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(rv.Interface().(*PseudoQJSValue).Property("___pointer").ToVariant().ToULongLong(nil))))})
				ptr.makeObjectWrapper(rvn.Interface(), jsv)
			}
		}
	}

	if property == nil {
		if reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()) {
			ptr.makeObjectWrapper(i[0], jsv)
		}
	} else {
		property.SetProperty(name, jsv)
	}

	return jsv
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{"reflect.Value.IsZero", v.Kind()})
	}
}

func (ptr *PseudoQJSEngine) makeFuncWrapper(fn string) *PseudoQJSValue {

	retV := ptr.GlobalObject().Property("___factory_single_func").Call([]*PseudoQJSValue{NewPseudoQJSValue8(fn), NewPseudoQJSValue8("")})

	//only needed for generic binding >>>
	retV.SetProperty("callable", ptr.NewGoType(true))
	retV.SetProperty("callableLocal", ptr.NewGoType(true))
	retV.SetProperty("callableName", ptr.NewGoType(fn))
	//<<<

	//TODO: allow creation of funcs in arbitrary module depth
	if strings.Count(fn, ".") == 0 {
		ptr.GlobalObject().SetProperty(fn, retV)
	} else {
		var jsv *PseudoQJSValue
		if m := ptr.GlobalObject().Property(strings.Split(fn, ".")[0]); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(strings.Split(fn, ".")[0], jsv)
		} else {
			jsv = m
		}
		jsv.SetProperty(strings.Split(fn, ".")[1], retV)
	}

	return retV
}

func (ptr *PseudoQJSEngine) makeObjectWrapper(in interface{}, jsv *PseudoQJSValue) {
	if ReturnPointersAsStrings {
		jsv.SetProperty("___pointer", ptr.ToScriptValue(core.NewQVariant1(strconv.FormatUint(uint64(uintptr(in.(ptr_itf).Pointer())), 10)))) //TODO: can be shortened once NewQVariant1 supports unsafe.Pointer
	} else {
		jsv.SetProperty("___pointer", ptr.ToScriptValue(core.NewQVariant1(uint64(uintptr(in.(ptr_itf).Pointer()))))) //TODO: can be shortened once NewQVariant1 supports unsafe.Pointer
	}

	rv := reflect.ValueOf(in)

	className := rv.Type().Elem().String()
	jsv.SetProperty("___className", NewPseudoQJSValue8(className))
	_, ok1 := qt.GetItfMap(className + "_ITF")
	if _, ok2 := qt.GetItfMap(className); !(ok1 || ok2) {
		qt.SetItfMap(className, rv.Elem().Interface())
	}

	/* TODO:
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	*/

	fns := make([]string, rv.NumMethod())
	for i := 0; i < rv.NumMethod(); i++ {
		fns[i] = rv.Type().Method(i).Name
	}
	ptr.GlobalObject().Property("___factory_batch_funcs").Call([]*PseudoQJSValue{jsv, NewPseudoQJSValue8(className), ptr.ToScriptValue(core.NewQVariant1(fns))})

	if qt.HasFinalizer(in) {
		buoy := core.NewQObject(nil)
		buoy.ConnectDestroyed(func(*core.QObject) { runtime.KeepAlive(in) })
		jsv.SetProperty("___buoy", ptr.NewQObject(buoy))

		finalizerMapMutex.Lock()
		finalizerMap[ptr.Pointer()] = append(finalizerMap[ptr.Pointer()], buoy)
		finalizerMapMutex.Unlock()
	}
}
