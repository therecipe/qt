package templater

import (
	"fmt"
	"strings"

	"github.com/bluszcz/cutego/internal/binding/parser"
)

func convertToSwift(className string, l string, convertClassMethods bool) string {
	l = strings.TrimSpace(l)

	if strings.HasPrefix(l, "func New"+className+"FromPointer") && !convertClassMethods {
		return fmt.Sprintf("public func New%[1]vFromPointer(ptr:String)->%[1]v { let r = %[1]v(); r.initFrom(p:ptr, n:\"%[2]v.%[1]v\"); return r; }", className, goModule(parser.State.ClassMap[className].Module))
	}

	if (strings.Contains(l, "ptr *"+className+")") && !strings.Contains(l, " __") && convertClassMethods) ||
		(strings.HasPrefix(l, "func  New"+className) && strings.Contains(l, className+"{") && !strings.Contains(l, "FromPointer(") && !convertClassMethods) ||
		(strings.HasPrefix(l, "func  "+className+"_") && !convertClassMethods) {
		l = strings.Replace(l, "  ", " ", -1)
		l = strings.Replace(l, "func (", "function(", -1)
		l = strings.Replace(l, "func New", "function_New", -1)
		l = strings.Replace(l, "func "+className, "function_"+className, -1)
		l = strings.Replace(l, ") )", "))", -1)
		l = strings.Replace(l, ", ", ",", -1)
		l = strings.Replace(l, "*", "", -1)

		if convertClassMethods {
			l = strings.Join(strings.Split(l, ") ")[1:], ") ")
		}

		l = strings.TrimSpace(strings.TrimSuffix(l, "{"))
		l = strings.Replace(l, " ", ":", -1)

		for _, v := range [][]string{{"string", "String"}, {"[]byte", "String"}, {"bool", "Bool"}, {"interface{}", "Any"}, {"...interface{}", "[Any]"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
		}

		for _, v := range []string{"string"} {
			l = strings.Replace(l, ":[]"+v, ":[String]", -1)
			l = strings.Replace(l, "["+v+"]", "[String]", -1)
			l = strings.Replace(l, "]"+v+")", "]String)", -1)
			l = strings.Replace(l, "]"+v, "]String", -1)
		}

		for _, v := range [][]string{{"float32", "Float"}, {"float64", "Float"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
			l = strings.Replace(l, ":[]"+v[0], ":["+v[1]+"]", -1)
			l = strings.Replace(l, "["+v[0]+"]", "["+v[1]+"]", -1)
		}

		for _, v := range []string{"uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "uintptr", "uint", "int", "unsafe.Pointer"} {
			l = strings.Replace(l, ":"+v, ":Int", -1)
			l = strings.Replace(l, ":[]"+v, ":[Int]", -1)
			l = strings.Replace(l, "["+v+"]", "[Int]", -1)
		}

		l = strings.Replace(l, "in:", "i:", -1)

		ls := strings.Split(l, ":")
		for i, lss := range ls {
			for _, sc := range []string{")", ",", ";"} {
				lsss := strings.Split(lss, sc)

				if strings.Contains(lsss[0], "[]") {
					for _, scc := range []string{",", ")", ";"} {
						lssss := strings.Split(lsss[0], scc)
						for j := range lssss {
							if strings.Contains(lssss[j], "[]") {

								if strings.Contains(lssss[j], ".") {
									lssss[j] = strings.Split(lssss[j], ".")[1]
								}

								lssss[j] = "[" + strings.Replace(lssss[j], "[]", "", -1) + "]"
							}
						}
						lsss[0] = strings.Join(lssss, scc)
					}
				}

				if strings.Contains(lsss[0], "map[") {
					lsss[0] = strings.Replace(lsss[0], "map[", "[", -1)
					for i, scc := range []string{",", ")", ";"} {
						lssss := strings.Split(lsss[0], scc)
						for j := range lssss {
							if strings.Contains(lssss[j], "]") {
								if i == 0 {
									lssss[j] = strings.Replace(lssss[j], "]", ",", -1) + "]"
								}

								lsssss := strings.Split(lssss[j], ",")
								for x, k := range lsssss {
									if strings.Contains(k, ".") {
										lsssss[x] = strings.Split(k, ".")[1]
										if x == 0 {
											lsssss[x] = "[" + lsssss[x]
										}
									}
								}
								lssss[j] = strings.Join(lsssss, ":")
							}
						}
						lsss[0] = strings.Join(lssss, scc)
					}
				}

				//TODO: support Enums
				if strings.Contains(lsss[0], "__") {
					index := strings.IndexAny(lsss[0], ",")
					if index != -1 {
						if strings.HasPrefix(lsss[0], "[") {
							lsss[0] = "[" + lsss[0][index:]
						} else {
							lsss[0] = "Int" + lsss[0][index:]
						}
					} else {
						lsss[0] = "Int"
					}
				}

				lss = strings.Join(lsss, sc)
			}
			ls[i] = lss
		}
		l = strings.Join(ls, ":")

		ls = strings.Split(l, "function(")
		for i, lss := range ls {
			if i == 0 {
				continue
			}
			lse := strings.Split(lss, ")")

			for j := range lse {
				if j == 0 || j == len(lse)-1 {
					continue
				}
				if strings.HasPrefix(lse[j], ":") {
					lse[j] = "->" + strings.TrimPrefix(lse[j], ":")
				}
			}

			ls[i] = strings.Join(lse, ")")
		}
		l = strings.Join(ls, "(")

		if strings.HasSuffix(l, ")") && !strings.Contains(l, "):(") {
			l += ":Void"
		}
		l = strings.Replace(l, "))", ")->Void)", -1)

		if strings.Contains(l, "):(") {
			ls = strings.Split(l, "):(")
			ls[1] = strings.TrimSuffix(ls[1], ")")
			if strings.Contains(ls[1], ",") {
				ls[1] = "[" + ls[1] + "]"
			}
			for _, v := range [][]string{{"int", "Int"}, {"bool", "Bool"}, {"error", "String"}} {
				ls[1] = strings.Replace(ls[1], v[0], v[1], -1)
			}
			l = strings.Join(ls, "):")
		}

		l = strings.Replace(l, "function_New", "public func New", -1)
		l = strings.Replace(l, "function_"+className, "public func "+className, -1)
		l = strings.Replace(l, "_ITF", "_ITF?=nil", -1)
		l = strings.Replace(l, "]:", "],", -1)

		///
		///
		///

		var params []string
		if in := getInner(l); in != "" {
			_, params = switchParameters(in)
		}

		var tys []string
		var block bool
		for i, in := range []string{getInner(l), getInner(getInner(l))} {
			_, ty := getTypes(in)
			for _, t := range ty {
				if strings.Contains(t, "[Q") && strings.Contains(t, ":") {
					block = true
				}

				if !strings.Contains(t, ".") {
					if i == 1 {
						tys = append(tys, t)
					}
					continue
				}

				if strings.Contains(t, "->") {
					if rs := strings.Split(t, "->"); strings.Contains(rs[len(rs)-1], ".") {
						l = strings.Replace(l, rs[len(rs)-1], strings.Split(rs[len(rs)-1], ".")[1], -1)

						if i == 1 {
							tys = append(tys, strings.Split(rs[len(rs)-1], ".")[1])
						}
					}
				} else if !strings.ContainsAny(t, "[]") {
					l = strings.Replace(l, t, strings.Split(t, ".")[1], -1)

					if i == 1 {
						tys = append(tys, strings.Split(t, ".")[1])
					}
				}
			}
		}

		ret, name := strings.Split(strings.Split(l, "):")[1], "{")[0], strings.Split(l, "(")[0]
		ret = strings.TrimSpace(ret)
		if ret == "Void" {
			l = strings.Replace(l, ":Void", "", -1)
		}
		l = strings.Replace(l, "):"+ret, ")->"+ret, -1)

		if strings.Contains(ret, ".") {
			l = strings.Replace(l, ret, strings.Split(ret, ".")[1], -1)
			ret = strings.Split(ret, ".")[1]
		}

		name = strings.TrimSpace(name)
		if !convertClassMethods {
			name = strings.TrimPrefix(name, "public func ")
		}

		//TODO: lookup Connect/Disconnect functions to figure out if they are really a signal/slot or virtual function
		if strings.HasPrefix(name, "Connect") && len(params) == 1 && ret == "Void" {
			for _, in := range []string{getInner(getInner(l))} {
				_, p := switchParameters(in)
				for _, pa := range p {
					l = strings.Replace(l, "("+pa, "(_ "+pa, -1)
					l = strings.Replace(l, ","+pa, ",_ "+pa, -1)
				}
			}
		}

		var methodPreamble string

		//TODO: does swift support implicit overriding ?
		if c, ok := parser.State.ClassMap[className]; ok {
			var (
				f       *parser.Function
				classes = append([]string{c.Name}, c.GetAllBases()...)
			)

			for i, bcn := range classes {
				if f == nil {
					f = parser.State.ClassMap[bcn].GetTitledFunction(name)
					if f == nil {
						lookupName := name
						switch {
						case strings.HasPrefix(name, "Connect"):
							lookupName = strings.TrimPrefix(lookupName, "Connect")
						case strings.HasPrefix(name, "Disconnect"):
							lookupName = strings.TrimPrefix(lookupName, "Disconnect")
						case strings.HasSuffix(name, "Default"):
							lookupName = strings.TrimSuffix(lookupName, "Default")
						}
						if lookupName != name {
							f = parser.State.ClassMap[bcn].GetTitledFunction(lookupName)
						}
					}

					if f != nil && i > 0 && strings.Title(f.Name) == strings.Title(name) {
						methodPreamble = "public override func "
					}
					continue
				}

				compareFunc := parser.State.ClassMap[bcn].GetTitledFunction(f.Name)
				if compareFunc == nil {
					compareFunc = parser.State.ClassMap[bcn].GetTitledFunction(name)
					if compareFunc == nil {
						continue
					}
				}

				if strings.HasPrefix(name, "Connect") || strings.HasPrefix(name, "Disconnect") || !strings.HasSuffix(name, "Default") ||
					(strings.HasSuffix(name, "Default") && parser.State.ClassMap[bcn].GetTitledFunction(name) != nil) ||
					(compareFunc.Virtual == parser.PURE && i > 1) || (compareFunc.Meta == parser.SLOT && compareFunc.Virtual == parser.IMPURE) {
					methodPreamble = "public override func "
				}

				if bcf := compareFunc; bcf != nil &&
					((len(bcf.Parameters) != len(f.Parameters)) || (bcf.Output != f.Output)) {
					l = strings.Replace(l, name+"(", name+"_"+classes[i-1]+"(", -1)
					if c.Name == classes[i-1] {
						methodPreamble = ""
					}
					break
				}
				var br bool
				for j, p := range compareFunc.Parameters {
					if p.Value != f.Parameters[j].Value {
						l = strings.Replace(l, name+"(", name+"_"+classes[i-1]+"(", -1)
						br = true
						if c.Name == classes[i-1] {
							methodPreamble = ""
						}
						break
					}
				}
				if br {
					break
				}
			}
		}
		//<<<

		if convertClassMethods {
			if methodPreamble == "" {
				methodPreamble = "public func "
			}
			l = methodPreamble + l
		}

		var input string
		if len(params) > 0 {
			input = "," + strings.Join(params, " as Any,") + " as Any"
		}

		var (
			return_prefix string
			return_suffix string
		)

		if ret != "Void" {
			return_prefix = "return "
			if ret != "Any" {
				return_suffix = " as! " + ret
			}
		} else {
			return_prefix = "_ = "
		}

		if ret == "Int" {
			return_prefix = "return Int("
			return_suffix = " as! Float)"
		}

		//TODO: type Q... does not conform to protocol 'Hashable'; can't use class objects as keys inside dictionaries
		if block || (strings.HasPrefix(ret, "[Q") && strings.Contains(ret, ":")) {
			l = "//TODO: " + l
		}

		/*
			if strings.HasPrefix(ret, "[String:") {
				return_prefix = "final _tmpValue:[String: Any] = "

				key := "k"
				switch {
				case strings.HasPrefix(ret, "[Int"):
					key = "Std.parseInt(k)"
				case strings.HasPrefix(ret, "[Q"):
					key = fmt.Sprintf("New%vFromPointer(k)", strings.TrimPrefix(strings.Split(ret, ",")[0], "["))
				}
				return_suffix = fmt.Sprintf("; return [for (k => v in _tmpValue) %v => v]", key) //Not needed for Map<String ?
			}
		*/

		if convertClassMethods {
			//TODO: lookup Connect/Disconnect functions to figure out if they are really a signal/slot or virtual function
			if strings.HasPrefix(name, "Connect") && len(params) == 1 && ret == "Void" && params[0] == "f" {
				l = strings.Replace(l, "(f:(", "(f: @escaping (", -1)

				fakref := make([]string, len(tys))
				for i := range tys {
					if tys[i] == "Int" {
						fakref[i] = fmt.Sprintf("Int(inp[%v] as! Float)", i)
					} else {
						fakref[i] = fmt.Sprintf("inp[%v] as! %v", i, tys[i])
					}
				}

				input = strings.Replace(input, ",f", fmt.Sprintf(",f:{(inp:[Any])->Any in f(%v);}", strings.Join(fakref, ",")), -1)
				l += fmt.Sprintf("{ _ = callLocalAndRegisterRemoteFunction(l:[\"\",Pointer(),___className,\"%v\",\"___REMOTE_CALLBACK___\"]%v); }", name, input)
			} else if strings.HasPrefix(name, "Disconnect") && len(params) == 0 && ret == "Void" {
				l += fmt.Sprintf("{ _ = callLocalAndDeregisterRemoteFunction(l:[\"\",Pointer(),___className,\"%v\"]); }", name)
			} else {
				l += fmt.Sprintf("{ %vcallLocalFunction(l:[\"\",Pointer(),___className,\"%v\"%v])%v; }", return_prefix, name, input, return_suffix)
			}
		} else {
			//TODO: call enum function
			l += fmt.Sprintf("{ %v.initModule(); %vcallLocalFunction(l:[\"\",\"\",\"%v.%v\",\"\"%v])%v; }", strings.TrimPrefix(parser.State.ClassMap[className].Module, "Qt"), return_prefix, goModule(parser.State.ClassMap[className].Module), name, input, return_suffix)
		}

		if !(name == "Pointer" || name == "SetPointer") {
			return l
		}
	}
	return ""
}
