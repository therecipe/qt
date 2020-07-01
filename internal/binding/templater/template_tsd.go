package templater

import "strings"

func convertToTypeScriptDefinition(className string, l string, convertClassMethods bool) string {
	l = strings.TrimSpace(l)

	if strings.HasPrefix(l, "func New"+className+"FromPointer") && !convertClassMethods {
		return "function New" + className + "FromPointer(ptr:number):" + className
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

		for _, v := range [][]string{{"[]byte", "string"}, {"bool", "boolean"}, {"interface{}", "any"}, {"...interface{}", "any[]"}} {
			l = strings.Replace(l, ":"+v[0], ":"+v[1], -1)
		}

		for _, v := range []string{"uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "uintptr", "uint", "int", "float32", "float64", "unsafe.Pointer"} {
			l = strings.Replace(l, ":"+v, ":number", -1)
			l = strings.Replace(l, ":[]"+v, ":[]number", -1)
			l = strings.Replace(l, "["+v+"]", "[number]", -1)
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
								lssss[j] = strings.Replace(lssss[j], "[]", "", -1) + "[_BLOCKING_]"
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
							lsss[0] = "Map<number" + lsss[0][index:]
						} else {
							lsss[0] = "number" + lsss[0][index:]
						}
					} else {
						lsss[0] = "number"
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
			for _, v := range [][]string{{"int", "number"}, {"bool", "boolean"}, {"error", "string"}} {
				ls[1] = strings.Replace(ls[1], v[0], v[1], -1)
			}
			l = strings.Join(ls, "):")
		}

		l = strings.Replace(l, "function_New", "function New", -1)
		l = strings.Replace(l, "function_"+className, "function "+className, -1)
		l = strings.Replace(l, "function:", "func:", -1)
		l = strings.Replace(l, "[_BLOCKING_]", "[]", -1)

		return l
	}
	return ""
}
