package qml

import (
	"math"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
)

var Z_newEngineHelper func(p core.QObject_ITF) *core.QObject

type ptr_itf interface {
	Pointer() unsafe.Pointer
}

func Z_initEngine(engine QJSEngine_ITF) {
	ptr := engine.QJSEngine_PTR()
	if !ptr.GlobalObject().Property("___engineHelper").IsUndefined() {
		return
	}
	ptr.GlobalObject().SetProperty("___engineHelper", ptr.NewQObject(Z_newEngineHelper(engine)))
	ptr.GlobalObject().SetProperty("___factory_single_func", ptr.Evaluate(`
(function(that, cn, fn) {
	return function(...args) {
		return Array.isArray(...args) ? ___engineHelper.wrapperFunction([that, cn, fn].concat(args)) : ___engineHelper.wrapperFunction([that, cn, fn].concat(...args));
	};
});

`, "", 0))
	ptr.GlobalObject().SetProperty("___factory_batch_funcs", ptr.Evaluate(`
(function(that, cn, fns) {
	for (var fn of fns) {
		that[fn] = ___factory_single_func(that, cn, fn, false);
	}
});

`, "", 0))
	ptr.GlobalObject().SetProperty("___factory_batch_global_funcs", ptr.Evaluate(`
(function(that, pfx, fns) {
	for (var fn of fns) {
		that[fn] = ___factory_single_func(___engineHelper, pfx+"."+fn, "", false);
	}
});
`, "", 0))

	ptr.GlobalObject().SetProperty("___factory_batch_global_enums", ptr.Evaluate(`
(function(that, ens) {
	for (var en of ens) {
		that[en[0]] = en[1];
	}
});
`, "", 0))

	prefE := make(map[string][][]interface{})
	for en, ev := range qt.EnumMap {
		pfx := strings.Split(en, ".")
		prefE[pfx[0]] = append(prefE[pfx[0]], []interface{}{pfx[1], ev})
	}
	for pfx, ens := range prefE {
		var jsv *QJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		ptr.GlobalObject().Property("___factory_batch_global_enums").Call([]*QJSValue{jsv, ptr.NewGoType(ens)})
	}

	pref := make(map[string][]string)
	for fn := range qt.FuncMap {
		pfx := strings.Split(fn, ".")
		if len(pfx) == 2 { //TODO: re-register 3rdparty functions?
			pref[pfx[0]] = append(pref[pfx[0]], pfx[1])
		}
	}

	for pfx, fns := range pref {
		var jsv *QJSValue
		if m := ptr.GlobalObject().Property(pfx); m.IsUndefined() {
			jsv = ptr.NewObject()
			ptr.GlobalObject().SetProperty(pfx, jsv)
		} else {
			jsv = m
		}
		ptr.GlobalObject().Property("___factory_batch_global_funcs").Call([]*QJSValue{jsv, NewQJSValue8(pfx), ptr.ToScriptValue(core.NewQVariant1(fns))})
	}
}

func Z_wrapperFunction(jsvals *QJSValue) *QJSValue {

	var (
		m      reflect.Value
		engine *QJSEngine
	)

	if cn, fn := jsvals.Property2(1).ToString(), jsvals.Property2(2).ToString(); fn == "" {
		m = reflect.ValueOf(qt.FuncMap[cn])

		engine = QJSEngine_qjsEngine(jsvals.Property2(0).ToQObject())
	} else {
		var rt reflect.Type
		if t, ok := qt.ItfMap[cn+"_ITF"]; ok {
			rt = reflect.TypeOf(t)
		} else {
			t, _ = qt.ItfMap[cn]
			rt = reflect.TypeOf(t)
		}

		//TODO: use ptr_itf instead ?
		//TODO: use *FromPointer instead
		//TODO: make SetPointer work as backup, for classes that don't support *FromPointer -> needs manual registration like for *FromPointer
		o := reflect.New(rt)
		o.MethodByName("SetPointer").Call([]reflect.Value{reflect.ValueOf(unsafe.Pointer(uintptr(jsvals.Property2(0).Property("___pointer").ToVariant().ToULongLong(nil))))})
		m = o.MethodByName(fn)

		engine = QJSEngine_qjsEngine(jsvals.Property2(0).Property("___engineHelper").ToQObject())
	}

	//

	input := make([]reflect.Value, m.Type().NumIn())
	for i := 0; i < len(input); i++ {
		input[i] = engine.fromJsToRef(m.Type().In(i), jsvals.Property2(uint(3+i)))
	}

	//

	ret := m.Call(input)
	if len(ret) == 0 {
		return NewQJSValue(QJSValue__UndefinedValue)
	}

	return engine.NewGoType(ret[0].Interface())
}

func (ptr *QJSEngine) ToGoType(jsval *QJSValue, dst interface{}) {
	reflect.ValueOf(dst).Elem().Set(ptr.fromJsToRef(reflect.TypeOf(dst), jsval).Elem())
}

func (ptr *QJSEngine) fromJsToRef(tofi reflect.Type, jsval *QJSValue) reflect.Value {
	tofiO := tofi

	if tofi.Kind() == reflect.Ptr {
		tofi = tofi.Elem()
	}

	t, ok := qt.ItfMap[tofi.String()+"_ITF"]
	if ok {
		tofi = reflect.TypeOf(t)
	} else {
		t, ok = qt.ItfMap[tofi.String()]
		if ok {
			tofi = reflect.TypeOf(t)
		}
	}

	switch {
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
	case tofi.ConvertibleTo(reflect.TypeOf(int8(0))):
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
	return reflect.ValueOf(ii.Elem().Interface())
}

//TODO: NewQJSValue1

func (ptr *QJSEngine) NewGoType(i ...interface{}) *QJSValue {
	rv := reflect.ValueOf(i[0])
	if len(i) > 1 {
		rv = reflect.ValueOf(i[1])
	}

	//TODO: merge into core.NewQVariant1 >>>
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

		if _, ok := qt.FuncMap[fn]; !ok {
			qt.FuncMap[fn] = rv.Interface()
		}

		return ptr.makeFuncWrapper(fn)

	case reflect.Slice, reflect.Array:
		jsv := ptr.NewArray(uint(rv.Len()))
		for i := 0; i < rv.Len(); i++ {
			jsv.SetProperty2(uint(i), ptr.NewGoType(rv.Index(i).Interface()))
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

	if v := core.NewQVariant1(i[0]); v.IsValid() {
		if !reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*core.QObject_ITF)(nil)).Elem()) && reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()) {
			jsv = ptr.NewObject()
		} else if reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*core.QObject_ITF)(nil)).Elem()) {
			jsv = ptr.NewQObject(core.NewQObjectFromPointer(i[0].(ptr_itf).Pointer()))
		} else {
			jsv = ptr.ToScriptValue(v)
		}
	}

	if reflect.TypeOf(i[0]).Implements(reflect.TypeOf((*ptr_itf)(nil)).Elem()) {
		ptr.makeObjectWrapper(i[0], jsv)
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

	retV := ptr.GlobalObject().Property("___factory_single_func").Call([]*QJSValue{ptr.GlobalObject().Property("___engineHelper"), NewQJSValue8(fn), NewQJSValue8("")})

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
	jsv.SetProperty("___engineHelper", ptr.GlobalObject().Property("___engineHelper"))

	rv := reflect.ValueOf(in)

	className := rv.Type().Elem().String()
	_, ok1 := qt.ItfMap[className+"_ITF"]
	if _, ok2 := qt.ItfMap[className]; !(ok1 || ok2) {
		qt.ItfMap[className] = rv.Elem().Interface()
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
}
