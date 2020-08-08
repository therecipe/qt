package templater

import (
	"fmt"
	"strings"

	"github.com/StarAurryon/qt/internal/binding/parser"

	"github.com/StarAurryon/qt/internal/utils"
)

func convertToHaxe(className string, l string, convertClassMethods bool) string {
	l = strings.TrimSpace(l)

	if strings.HasPrefix(l, "func New"+className+"FromPointer") && !convertClassMethods {
		return fmt.Sprintf("function New%[1]vFromPointer(ptr:String):%[1]v { final r = new %[1]v(); r.initFrom(ptr, \"%[2]v.%[1]v\"); return r; }", className, goModule(parser.State.ClassMap[className].Module))
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

		for _, v := range [][]string{{"string", "String"}, {"[]byte", "String"}, {"bool", "Bool"}, {"interface{}", "Any"}, {"...interface{}", "Array<Any>"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
		}

		for _, v := range []string{"string"} {
			l = strings.Replace(l, ":[]"+v, ":[]String", -1)
			l = strings.Replace(l, "["+v+"]", "[String]", -1)
			l = strings.Replace(l, "]"+v+")", "]String)", -1)
			l = strings.Replace(l, "]"+v, "]String", -1)
		}

		for _, v := range [][]string{{"float32", "Float"}, {"float64", "Float"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
			l = strings.Replace(l, ":[]"+v[0], ":[]"+v[1], -1)
			l = strings.Replace(l, "["+v[0]+"]", "["+v[1]+"]", -1)
		}

		for _, v := range []string{"uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "uintptr", "uint", "int", "unsafe.Pointer"} {
			l = strings.Replace(l, ":"+v, ":Int", -1)
			l = strings.Replace(l, ":[]"+v, ":[]Int", -1)
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

								lssss[j] = "Array<" + strings.Replace(lssss[j], "[]", "", -1) + ">"
							}
						}
						lsss[0] = strings.Join(lssss, scc)
					}
				}

				if strings.Contains(lsss[0], "map[") {
					lsss[0] = strings.Replace(lsss[0], "map[", "Map<", -1)
					for _, scc := range []string{",", ")", ";"} {
						lssss := strings.Split(lsss[0], scc)
						for j := range lssss {
							if strings.Contains(lssss[j], "]") {
								lssss[j] = strings.Replace(lssss[j], "]", ",", -1) + ">"

								lsssss := strings.Split(lssss[j], ",")
								for x, k := range lsssss {
									if strings.Contains(k, ".") {
										lsssss[x] = strings.Split(k, ".")[1]
										if x == 0 {
											lsssss[x] = "Map<" + lsssss[x]
										}
									}
								}
								lssss[j] = strings.Join(lsssss, ",")
							}
						}
						lsss[0] = strings.Join(lssss, scc)
					}
				}

				//TODO: support Enums
				if strings.Contains(lsss[0], "__") {
					index := strings.IndexAny(lsss[0], ",")
					if index != -1 {
						if strings.HasPrefix(lsss[0], "Map<") {
							lsss[0] = "Map<Int" + lsss[0][index:]
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

		l = strings.Replace(l, "function_New", "function New", -1)
		l = strings.Replace(l, "function_"+className, "function "+className, -1)
		l = strings.Replace(l, "function:", "func:", -1)
		l = strings.Replace(l, "[_BLOCKING_]", "[]", -1)

		///
		///
		///

		var params []string
		if in := getInner(l); in != "" {
			_, params = switchParameters(in)
		}

		for _, in := range []string{getInner(l), getInner(getInner(l))} {
			_, ty := getTypes(in)
			for _, t := range ty {
				if !strings.Contains(t, ".") {
					continue
				}

				if !strings.ContainsAny(t, "<>") {
					l = strings.Replace(l, t, strings.Split(t, ".")[1], -1)
				} else if strings.Contains(t, "->") {
					if rs := strings.Split(t, "->"); strings.Contains(rs[len(rs)-1], ".") {
						l = strings.Replace(l, rs[len(rs)-1], strings.Split(rs[len(rs)-1], ".")[1], -1)
					}
				}
			}
		}

		ret, name := strings.Split(strings.Split(l, "):")[1], "{")[0], strings.Split(l, "(")[0]
		ret = strings.TrimSpace(ret)
		if ret == "Void" {
			l = strings.Replace(l, ":Void", "", -1)
		}
		if strings.Contains(ret, ".") {
			l = strings.Replace(l, ret, strings.Split(ret, ".")[1], -1)
		}

		name = strings.TrimSpace(name)
		if !convertClassMethods {
			name = strings.TrimPrefix(name, "function ")
		}

		var methodPreamble string

		//needed because haxe isn't supporting overloading neither implicit overriding
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
						methodPreamble = "public override function "
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
					methodPreamble = "public override function "
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
		//

		if convertClassMethods {
			if methodPreamble == "" {
				methodPreamble = "public function "
			}
			l = methodPreamble + l
		}

		var input string
		if len(params) > 0 {
			input = "," + strings.Join(params, ",")
		}

		var (
			return_prefix string
			return_suffix string
		)

		if ret != "Void" {
			return_prefix = "return "
		}
		if strings.HasPrefix(ret, "Map<") {
			return_prefix = "final _tmpValue:Map<String, Dynamic> = "

			key := "k"
			switch {
			case strings.HasPrefix(ret, "Map<Int"):
				key = "Std.parseInt(k)"
			case strings.HasPrefix(ret, "Map<Q"):
				key = fmt.Sprintf("New%vFromPointer(k)", strings.TrimPrefix(strings.Split(ret, ",")[0], "Map<"))
			}
			return_suffix = fmt.Sprintf("; return [for (k => v in _tmpValue) %v => v]", key) //TODO: not needed for Map<String ?
		}

		if convertClassMethods {
			//TODO: lookup Connect/Disconnect functions to figure out if they are really a signal/slot or virtual function
			if strings.HasPrefix(name, "Connect") && len(params) == 1 && ret == "Void" && params[0] == "f" {
				l += fmt.Sprintf("{ Internal.callLocalAndRegisterRemoteFunction([\"\",Pointer(),___className,\"%v\",\"___REMOTE_CALLBACK___\"]%v); }", name, input)
			} else if strings.HasPrefix(name, "Disconnect") && len(params) == 0 && ret == "Void" {
				l += fmt.Sprintf("{ Internal.callLocalAndDeregisterRemoteFunction([\"\",Pointer(),___className,\"%v\"]); }", name)
			} else {
				l += fmt.Sprintf("{ %vInternal.callLocalFunction([\"\",Pointer(),___className,\"%v\"%v])%v; }", return_prefix, name, input, return_suffix)
			}
		} else {
			//TODO: call enum function
			l += fmt.Sprintf("{ %v.initModule(); %vInternal.callLocalFunction([\"\",\"\",\"%v.%v\",\"\"%v])%v; }", strings.TrimPrefix(parser.State.ClassMap[className].Module, "Qt"), return_prefix, goModule(parser.State.ClassMap[className].Module), name, input, return_suffix)
		}

		if !(name == "Pointer" || name == "SetPointer") {
			return l
		}
	}
	return ""
}

func getTypes(i string) (string, []string) {
	in := getInner(i)
	if in != "" {
		i = strings.Replace(i, in, "___PLACEHOLDER___", -1)
	}

	var pNames []string

	if strings.Contains(i, ":") {
		cs := strings.Split(i, ",")
		for i, v := range cs {
			vs := strings.Split(v, ":")
			if len(vs) == 1 {
				continue
			}

			var retval string
			if strings.Contains(vs[1], "=>") {
				vss := strings.Split(vs[1], "=>")
				vs[1] = vss[0]
				retval = vss[1]
			}

			if utils.QT_GEN_SWIFT() {
				if strings.HasPrefix(vs[1], "[") && !strings.HasSuffix(vs[1], "]") {
					pNames = append(pNames, vs[1]+":"+vs[2])
				} else {
					pNames = append(pNames, vs[1])
				}
			} else {
				pNames = append(pNames, vs[1])
			}

			vs[0], vs[1] = vs[1], vs[0]

			cs[i] = strings.Join(vs, " ")

			if retval != "" {
				cs[i] = retval + " " + cs[i]
			}
		}
		i = strings.Join(cs, ",")
	}

	if in != "" {
		n, _ := getTypes(in)
		i = strings.Replace(i, "___PLACEHOLDER___", n, -1)
	}

	return i, pNames
}
