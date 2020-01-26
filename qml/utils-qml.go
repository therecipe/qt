package qml

import (
	"math"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unsafe"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
)

var (
	Z_newEngineHelper func(p core.QObject_ITF) *core.QObject

	finalizerMap      = make(map[unsafe.Pointer][]*core.QObject)
	finalizerMapMutex sync.Mutex
)

type ptr_itf interface {
	Pointer() unsafe.Pointer
}

func Z_initEngine(engine QJSEngine_ITF) {
	ptr := engine.QJSEngine_PTR()
	ptr.ConnectDestroyed(func(ptr *core.QObject) {
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
	ptr.GlobalObject().SetProperty("nil", NewQJSValue(QJSValue__UndefinedValue))
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
		var jsv *QJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		inp := ptr.NewGoType(ens)
		ptr.GlobalObject().Property("___factory_batch_global_enums").Call([]*QJSValue{jsv, inp})
		inp.DestroyQJSValue()
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
		var jsv *QJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		inp := ptr.NewGoType(fns)
		ptr.GlobalObject().Property("___factory_batch_global_funcs").Call([]*QJSValue{jsv, NewQJSValue8(pfx), inp})
		inp.DestroyQJSValue()
	}
}

func Z_wrapperFunction(jsvals *QJSValue) *QJSValue {

	var m reflect.Value
	if cn, fn := jsvals.Property2(2).ToString(), jsvals.Property2(3).ToString(); fn == "" {
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

	//

	engine := QJSEngine_qjsEngine(core.NewQObjectFromPointer(unsafe.Pointer(uintptr(jsvals.Property2(0).ToVariant().ToULongLong(nil)))))

	input := make([]reflect.Value, m.Type().NumIn())
	for i := 0; i < len(input); i++ {
		input[i] = engine.fromJsToRef(m.Type().In(i), jsvals.Property2(uint(4+i)))
	}

	//

	ret := m.Call(input)
	if len(ret) == 0 {
		return NewQJSValue(QJSValue__UndefinedValue)
	}

	rret := engine.NewGoType(ret[0].Interface())
	if reflect.TypeOf(ret[0].Interface()).Implements(reflect.TypeOf((*core.QObject_ITF)(nil)).Elem()) { //TODO: check for destroyed signal instead, or simply override the destructor instead ?
		if qt.ExistsSignal(ret[0].Interface().(ptr_itf).Pointer(), "destroyed") {
			engine.GlobalObject().Property("___connectDestroyed").Call([]*QJSValue{rret}) //TODO: connect destroyed/destructor from go instead ?
		}
	}

	return rret
}

func (ptr *QJSEngine) ToGoType(jsval *QJSValue, dst interface{}) {
	reflect.ValueOf(dst).Elem().Set(ptr.fromJsToRef(reflect.TypeOf(dst), jsval).Elem())
}

func (ptr *QJSEngine) fromJsToRef(tofi reflect.Type, jsval *QJSValue) reflect.Value {
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
			input := make([]*QJSValue, len(args))
			for i, arg := range args {
				input[i] = ptr.NewGoType(arg.Interface())
			}
			if ret := jsval.Call(input); tofi.NumOut() != 0 {
				return []reflect.Value{ptr.fromJsToRef(tofi.Out(0), ret)}
			}
			return nil
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
		if tofi.String() == "qml.QJSValue" {
			if className := jsval.Property("___className").ToString(); className != "qml.QJSValue" {
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

//TODO: NewQJSValue1

func (ptr *QJSEngine) NewJSType(property *QJSValue, name string, i *QJSValue) { //TODO: support for js functions
	ptr.newGoType(property, name, i)
}

func (ptr *QJSEngine) NewGoType(i ...interface{}) *QJSValue {
	return ptr.newGoType(nil, "", i...)
}

func (ptr *QJSEngine) newGoType(property *QJSValue, name string, i ...interface{}) *QJSValue {
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
			inp.DestroyQJSValue()
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

	var jsv *QJSValue
	if reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()) {
		jsv = ptr.NewObject()
	} else {
		jsv = ptr.ToScriptValue(core.NewQVariant1(i[0]))
	}

	if rv.Type().String() == "*qml.QJSValue" {
		className := rv.Interface().(*QJSValue).Property("___className").ToString()
		if className != "qml.QJSValue" {
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
				rvn.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(rv.Interface().(*QJSValue).Property("___pointer").ToVariant().ToULongLong(nil))))})
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

func (ptr *QJSEngine) makeFuncWrapper(fn string) *QJSValue {

	retV := ptr.GlobalObject().Property("___factory_single_func").Call([]*QJSValue{NewQJSValue8(fn), NewQJSValue8("")})

	//TODO: allow creation of funcs in arbitrary module depth
	var jsv *QJSValue
	if m := ptr.GlobalObject().Property(strings.Split(fn, ".")[0]); m.IsUndefined() {
		jsv = ptr.NewObject()
		ptr.GlobalObject().SetProperty(strings.Split(fn, ".")[0], jsv)
	} else {
		jsv = m
	}
	if strings.Count(fn, ".") == 0 {
		ptr.GlobalObject().SetProperty(fn, retV)
	} else {
		jsv.SetProperty(strings.Split(fn, ".")[1], retV)
	}

	return retV
}

func (ptr *QJSEngine) makeObjectWrapper(in interface{}, jsv *QJSValue) {
	jsv.SetProperty("___pointer", ptr.ToScriptValue(core.NewQVariant1(uint64(uintptr(in.(ptr_itf).Pointer()))))) //TODO: can be shortened once NewQVariant1 supports unsafe.Pointer

	rv := reflect.ValueOf(in)

	className := rv.Type().Elem().String()
	jsv.SetProperty("___className", NewQJSValue8(className))
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
	ptr.GlobalObject().Property("___factory_batch_funcs").Call([]*QJSValue{jsv, NewQJSValue8(className), ptr.ToScriptValue(core.NewQVariant1(fns))})

	if qt.HasFinalizer(in) {
		buoy := core.NewQObject(nil)
		buoy.ConnectDestroyed(func(*core.QObject) { runtime.KeepAlive(in) })
		jsv.SetProperty("___buoy", ptr.NewQObject(buoy))

		finalizerMapMutex.Lock()
		finalizerMap[ptr.Pointer()] = append(finalizerMap[ptr.Pointer()], buoy)
		finalizerMapMutex.Unlock()
	}
}
