package templater

import (
	"fmt"
	"strings"

	"github.com/StarAurryon/qt/internal/binding/parser"
)

func convertToDart(className string, l string, convertClassMethods bool) string {
	l = strings.TrimSpace(l)

	if strings.HasPrefix(l, "func New"+className+"FromPointer") && !convertClassMethods {
		return fmt.Sprintf("%[1]v New%[1]vFromPointer(int ptr) { final r = new %[1]v(); r.initFrom(ptr, \"%[2]v.%[1]v\"); return r; }", className, goModule(parser.State.ClassMap[className].Module))
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

		for _, v := range [][]string{{"string", "String"}, {"[]byte", "String"}, {"interface{}", "dynamic"}, {"...interface{}", "List<dynamic>"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
		}

		for _, v := range []string{"string"} {
			l = strings.Replace(l, ":[]"+v, ":[]String", -1)
			l = strings.Replace(l, "["+v+"]", "[String]", -1)
			l = strings.Replace(l, "]"+v+")", "]String)", -1)
			l = strings.Replace(l, "]"+v, "]String", -1)
		}

		/* TODO: function signature mismatches in overridden methods + dart does values runtime checks to validate floats/doubles
		for _, v := range [][]string{{"float32", "double"}, {"float64", "double"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
			l = strings.Replace(l, ":[]"+v[0], ":[]"+v[1], -1)
			l = strings.Replace(l, "["+v[0]+"]", "["+v[1]+"]", -1)
		}
		*/

		for _, v := range []string{"uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "uintptr", "uint", "int", "float32", "float64", "unsafe.Pointer"} {
			l = strings.Replace(l, ":"+v, ":num", -1)
			l = strings.Replace(l, ":[]"+v, ":[]num", -1)
			l = strings.Replace(l, "["+v+"]", "[num]", -1)
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
								lssss[j] = "List<" + strings.Replace(lssss[j], "[]", "", -1) + ">"
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
								lssss[j] = strings.Replace(lssss[j], "]", "#", -1) + ">"
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
							lsss[0] = "Map<int" + lsss[0][index:]
						} else {
							lsss[0] = "int" + lsss[0][index:]
						}
					} else {
						lsss[0] = "int"
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
					lse[j] = "=>" + strings.TrimPrefix(lse[j], ":")
				}
			}

			ls[i] = strings.Join(lse, ")")
		}
		l = strings.Join(ls, "(")

		if strings.HasSuffix(l, ")") && !strings.Contains(l, "):(") {
			l += ":void"
		}
		l = strings.Replace(l, "))", ")=>void)", -1)

		if strings.Contains(l, "):(") {
			ls = strings.Split(l, "):(")
			ls[1] = strings.TrimSuffix(ls[1], ")")
			if strings.Contains(ls[1], ",") {
				ls[1] = "[" + ls[1] + "]"
			}
			for _, v := range [][]string{{"error", "String"}} {
				ls[1] = strings.Replace(ls[1], v[0], v[1], -1)
			}
			l = strings.Join(ls, "):")
		}

		l = strings.Replace(l, "function_New", "New", -1)
		l = strings.Replace(l, "function_"+className, className, -1)
		l = strings.Replace(l, "function:", "func:", -1)
		l = strings.Replace(l, "[_BLOCKING_]", "[]", -1)

		///
		///
		///

		l = strings.Replace(l, ":(", ":Function(", -1)

		var params []string
		if in := getInner(l); in != "" {
			n, c := switchParameters(in)
			params = c
			l = strings.Replace(l, in, n, -1)
		}

		l = strings.Replace(l, "#", ",", -1)

		ls = strings.Split(l, ":")
		if len(ls) == 2 {
			ls[0], ls[1] = ls[1], ls[0]
			l = strings.Join(ls, " ")
		}

		ls = strings.Split(strings.Split(l, "(")[0], " ")

		ret, name := ls[0], ls[1]
		ret = strings.TrimSpace(ret)
		name = strings.TrimSpace(name)

		//needed because dart isn't supporting overloading: https://github.com/dart-lang/sdk/issues/26488
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
					continue
				}
				compareFunc := parser.State.ClassMap[bcn].GetTitledFunction(f.Name)
				if compareFunc == nil {
					compareFunc = parser.State.ClassMap[bcn].GetTitledFunction(name)
					if compareFunc == nil {
						continue
					}
				}
				if bcf := compareFunc; bcf != nil &&
					((len(bcf.Parameters) != len(f.Parameters)) || (bcf.Output != f.Output)) {
					l = strings.Replace(l, name+"(", name+"_"+classes[i-1]+"(", -1)
					break
				}
				var br bool
				for j, p := range compareFunc.Parameters {
					if p.Value != f.Parameters[j].Value {
						l = strings.Replace(l, name+"(", name+"_"+classes[i-1]+"(", -1)
						br = true
						break
					}
				}
				if br {
					break
				}
			}
		}
		if name == "String" || name == "Map" { //TODO:
			l = strings.Replace(l, name+"(", name+"_Function(", -1)
		}
		//

		var input string
		if len(params) > 0 {
			input = "," + strings.Join(params, ",")
		}

		var (
			return_prefix string
			return_suffix string
		)

		if ret != "void" {
			return_prefix = "return "
		}
		if strings.HasPrefix(ret, "List<") || strings.HasPrefix(ret, "Map<") {
			return_prefix = "return " + ret + ".from("
			return_suffix = ")"
		}

		if convertClassMethods {
			//TODO: lookup Connect/Disconnect functions to figure out if they are really a signal/slot or virtual function
			if strings.HasPrefix(name, "Connect") && len(params) == 1 && ret == "void" && params[0] == "f" {
				l += fmt.Sprintf("{ callLocalAndRegisterRemoteFunction([\"\",this.Pointer(),this.className,\"%v\",\"___REMOTE_CALLBACK___\"]%v); }", name, input)
			} else if strings.HasPrefix(name, "Disconnect") && len(params) == 0 && ret == "void" {
				l += fmt.Sprintf("{ callLocalAndDeregisterRemoteFunction([\"\",this.Pointer(),this.className,\"%v\"]); }", name)
			} else {
				l += fmt.Sprintf("{ %vcallLocalFunction([\"\",this.Pointer(),this.className,\"%v\"%v])%v; }", return_prefix, name, input, return_suffix)
			}
		} else {
			//TODO: call enum function
			l += fmt.Sprintf("{ initModule(); %vcallLocalFunction([\"\",\"\",\"%v.%v\",\"\"%v])%v; }", return_prefix, goModule(parser.State.ClassMap[className].Module), name, input, return_suffix)
		}

		if !(name == "Pointer" || name == "SetPointer") {
			return l
		}
	}
	return ""
}

func getInner(i string) string {
	pres := strings.Split(i, "(")
	if len(pres) > 1 {
		pres = strings.Split(strings.Join(pres[1:], "("), ")")
		return strings.Join(pres[:len(pres)-1], ")")
	}
	return ""
}

func switchParameters(i string) (string, []string) {
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

			pNames = append(pNames, vs[0])

			vs[0], vs[1] = vs[1], vs[0]

			cs[i] = strings.Join(vs, " ")

			if retval != "" {
				cs[i] = retval + " " + cs[i]
			}
		}
		i = strings.Join(cs, ",")
	}

	if in != "" {
		n, _ := switchParameters(in)
		i = strings.Replace(i, "___PLACEHOLDER___", n, -1)
	}

	return i, pNames
}
