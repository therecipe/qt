package interop

import (
	"encoding/json"

	"github.com/dev-drprasad/qt/core"

	"github.com/dev-drprasad/qt"
)

//TODO: this == nil checks

const PseudoQJSValue__UndefinedValue int64 = 0

type PseudoQJSValue struct{ core.QVariant }

type PseudoQJSValue_ITF interface {
	core.QVariant_ITF
	PseudoQJSValue_PTR() *PseudoQJSValue
}

func NewPseudoQJSValue(i int64) *PseudoQJSValue {
	return new(PseudoQJSValue)
}

func NewPseudoQJSValue1(i interface{}) *PseudoQJSValue {
	return NewPseudoQJSValue2(core.NewQVariant1(i))
}

func NewPseudoQJSValue2(d *core.QVariant) *PseudoQJSValue {
	if d != nil {
		qt.SetFinalizer(d, nil) //there is the need to re-assign the finalizer to the PseudoQJSValue pointer to prevent double freeing the QVariant >>>
		ret := &PseudoQJSValue{*d}
		qt.SetFinalizer(ret, (*PseudoQJSValue).DestroyPseudoQJSValue) //<<< by the finalizer in case it was manually destroyed earlier by the destructor
		return ret
	}
	return &PseudoQJSValue{}
}

func NewPseudoQJSValue8(i string) *PseudoQJSValue {
	return NewPseudoQJSValue1(i)
}

func (this *PseudoQJSValue) ToInt() int {
	return this.QVariant.ToInt(nil)
}

func (this *PseudoQJSValue) IsUndefined() bool {
	return this == nil || this.Pointer() == nil
}

//public >>>

func (this *PseudoQJSValue) QPseudoQJSValue_PTR() *PseudoQJSValue {
	return this
}

func (this *PseudoQJSValue) Length() int {
	if this.IsObject() {
		return len(this.ToMap())
	} else if this.IsArray() {
		return len(this.ToList())
	}
	return 0
}

func (this *PseudoQJSValue) IsArray() bool {
	return this.QVariant.Type() == core.QVariant__List
}

func (this *PseudoQJSValue) IsObject() bool {
	return this.QVariant.Type() == core.QVariant__Map
}

func (this *PseudoQJSValue) IsCallable() bool {
	return this.Property("callable").ToBool()
}

func (this *PseudoQJSValue) HasProperty(n string) bool {
	if this.IsObject() {
		if _, ok := this.QVariant.ToMap()[n]; ok {
			return ok
		}
	}
	return false
}

func (this *PseudoQJSValue) Property(n string) *PseudoQJSValue {
	if n == "length" && !this.HasProperty(n) {
		return NewPseudoQJSValue1(this.Length())
	}

	if this.IsObject() {
		return NewPseudoQJSValue2(this.ToMap()[n])
	}
	//TODO: make list usable as well?
	return NewPseudoQJSValue(1)
}

func (this *PseudoQJSValue) Property2(n uint) *PseudoQJSValue {
	if this.IsArray() {
		return NewPseudoQJSValue2(this.ToList()[n])
	}
	//TODO: make map usable as well?
	return NewPseudoQJSValue(1)
}

func (this *PseudoQJSValue) SetProperty(n string, v *PseudoQJSValue) {
	if this.IsObject() {
		data := this.ToMap()
		data[n] = v.ToVariant()
		this.Swap(core.NewQVariant1(data))
	}
	//TODO: make list usable as well?
}

func (this *PseudoQJSValue) SetProperty2(n uint, v *PseudoQJSValue) {
	if this.IsArray() {
		data := this.ToList()
		data[n] = v.ToVariant()
		this.Swap(core.NewQVariant1(data))
	}
	//TODO: make map usable as well?
}

func (this *PseudoQJSValue) DeleteProperty(n string) bool {
	if this.IsObject() {
		data := this.ToMap()
		if _, ok := data[n]; ok {
			delete(data, n)
			this.Swap(core.NewQVariant1(data))
		}
		return true
	}
	return false
}

func (this *PseudoQJSValue) ToVariant() *core.QVariant {
	return &this.QVariant
}

func (this *PseudoQJSValue) DestroyPseudoQJSValue() {
	this.DestroyQVariant()
}

func (this *PseudoQJSValue) Call(in []*PseudoQJSValue) *PseudoQJSValue {

	if !this.IsCallable() {
		return NewPseudoQJSValue1(make(map[string]interface{}))
	}

	if this.Property("callableLocal").ToBool() {
		_, ok := qt.GetFuncMap(this.Property("callableName").ToString())
		if !ok {
			return NewPseudoQJSValue1(make(map[string]interface{}))
		}
		return callLocalStaticFunction(this.Property("callableName").ToString(), in)
	}

	return this.callRemote(in)
}

func (this *PseudoQJSValue) callRemote(in []*PseudoQJSValue) *PseudoQJSValue {

	var pointer interface{}
	if ReturnPointersAsStrings {
		pointer = this.Property("___pointer").ToString()
	} else {
		pointer = this.Property("___pointer").ToInterface()
	}

	i := []interface{}{pointer, this.Property("callableName").ToInterface()}
	for _, v := range in {
		if v.Property("___pointer").ToString() != "" {
			var p interface{}
			if ReturnPointersAsStrings {
				p = v.Property("___pointer").ToString()
			} else {
				p = v.Property("___pointer").ToInterface()
			}

			i = append(i, struct {
				Pointer   interface{} `json:"___pointer"`
				ClassName interface{} `json:"___className"`
			}{
				p,
				v.Property("___className").ToInterface(),
			})
		} else {
			switch v.Type() {
			case core.QVariant__List:
				var o []interface{}
				PseudoQJSEngine_qjsEngine(nil).ToGoType(v, &o)
				i = append(i, o)

			case core.QVariant__Map:
				var o map[string]interface{}
				PseudoQJSEngine_qjsEngine(nil).ToGoType(v, &o)
				i = append(i, o)

			default:
				i = append(i, v.ToInterface())
			}
		}
	}
	ib, _ := json.Marshal(i)

	remoteMainThreadIsBlockedMutex.Lock()
	blocked := remoteMainThreadIsBlocked
	remoteMainThreadIsBlockedMutex.Unlock()

	var (
		callFunc = AsyncCallIntoRemote
		callChan = asyncCallChan
	)

	if blocked {
		if !SupportsSyncCallsIntoRemote {
			//TODO: only for functions that return nothing or if it's requested specifically
			m := make(map[string]interface{})
			m["___earlyReturn"] = "true"
			m["___data"] = string(ib)
			return NewPseudoQJSValue1(m)
		}

		callFunc = SyncCallIntoRemote
		callChan = syncCallChan
	}

	if !localMainThreadIsBlocked {
		localMainThreadIsBlocked = true
		defer func() { localMainThreadIsBlocked = false }()
	}

	callFunc(string(ib))

	var ret string
	for {
		select {
		case ret = <-callChan:
			goto br

		case f := <-mainThreadHelperChan:
			f()
		}
	}
br:

	var o interface{}
	json.Unmarshal([]byte(ret), &o)
	return NewPseudoQJSValue1(o)
}

func (this *PseudoQJSValue) CallMethod(name string, in []*PseudoQJSValue) *PseudoQJSValue {
	//enginePtr, objPtr, className, funcName
	input := []interface{}{"", this.Property("___pointer").ToInterface(), this.Property("___className").ToInterface(), name}
	for _, v := range in {
		input = append(input, v.ToInterface())
	}
	return Z_wrapperFunction(NewPseudoQJSValue1(input))
}

//<<< public

func callLocalStaticFunction(name string, in []*PseudoQJSValue) *PseudoQJSValue {
	//enginePtr, objPtr, funcName, empty
	input := []interface{}{"", "", name, ""}
	for _, v := range in {
		input = append(input, v.ToInterface())
	}
	return Z_wrapperFunction(NewPseudoQJSValue1(input))
}
