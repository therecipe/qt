package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func (f *Function) fix() {
	f.fixGeneral()
	f.fixGeneral_Version()

	f.fixOverload()
	f.fixOverload_Version()

	f.fixGeneric()

	f.fixTemporary()
}

func (f *Function) fixGeneral() {

	//linux fixes

	if f.Fullname == "QThread::start" {
		f.Parameters = make([]*Parameter, 0)
	}

	//virtual fixes

	if f.Virtual == "virtual" {
		f.Virtual = IMPURE
	}

	if f.Virtual == IMPURE || f.Virtual == PURE ||
		f.Meta == SIGNAL || f.Meta == SLOT {

		f.Static = false
	}

	//constructor fixes

	if f.Meta == COPY_CONSTRUCTOR || f.Meta == MOVE_CONSTRUCTOR {
		f.Meta = CONSTRUCTOR
	}

	var class, exists = f.Class()
	if !exists || !class.isSubClass() {
		return
	}

	if f.Meta == CONSTRUCTOR {
		f.Status = "active"
		f.Access = "public"
	}
}

func (f *Function) fixGeneral_AfterClasses() {
	if f.Name != "open" && f.Name != "setGeometry" && f.Name != "setScxmlEvent" && //TODO: solve
		!f.Static && f.Virtual == "non" && f.Meta == PLAIN && f.IsDerivedFromVirtual() {
		f.Virtual = IMPURE
	}
}

func (f *Function) fixGeneral_Version() {
	switch f.Fullname {
	case "QScxmlCppDataModel::setScxmlEvent":
		{
			f.Virtual = "non"
		}

	case "QGraphicsObject::z", "QGraphicsObject::setZ":
		{
			f.Name = func() string {
				if f.Name == "setZ" {
					return "setZValue"
				}
				return "zValue"
			}()
			f.Fullname = fmt.Sprintf("%v::%v", f.ClassName(), f.Name)
		}

	case "QGraphicsObject::effect", "QGraphicsObject::setEffect":
		{
			f.Name = func() string {
				if f.Name == "setEffect" {
					return "setGraphicsEffect"
				}
				return "graphicsEffect"
			}()
			f.Fullname = fmt.Sprintf("%v::%v", f.ClassName(), f.Name)
		}
	}
}

func (f *Function) fixOverload() {

	if strings.Contains(f.Href, "-") {
		var tmp, err = strconv.Atoi(strings.Split(f.Href, "-")[1])
		if err == nil && tmp > 0 {
			f.Overload = true
			tmp++
			f.OverloadNumber = strconv.Itoa(tmp)
		}
	}

	if f.OverloadNumber != "0" {
		return
	}

	f.Overload = false
	f.OverloadNumber = ""
}

func (f *Function) fixOverload_Version() {
	switch f.Fullname {
	case "QGraphicsDropShadowEffect::setOffset", "QGraphicsScene::setSceneRect",
		"QGraphicsView::setSceneRect", "QQuickItem::setFocus",
		"QAccessibleWidget::setText", "QSvgGenerator::setViewBox",
		"QSvgRenderer::setViewBox":
		{
			var class, exist = f.Class()
			if !exist {
				return
			}

			var count int
			for _, sf := range class.Functions {
				if sf.Fullname != f.Fullname {
					continue
				}

				if sf.Signature != f.Signature {
					count++
					continue
				}

				break
			}
			if count == 0 {
				return
			}

			f.Overload = true
			f.OverloadNumber = strconv.Itoa(count + 1)
		}
	}
}

func (f *Function) fixGeneric() {

	switch f.Output {
	case "QModelIndexList":
		{
			f.Output = "QList<QModelIndex>"
		}

	case "QVariantList":
		{
			f.Output = "QList<QVariant>"
		}

	case "QObjectList":
		{
			f.Output = "QList<QObject *>"
		}

	case "QMediaResourceList":
		{
			f.Output = "QList<QMediaResource>"
		}

	case "QFileInfoList":
		{
			f.Output = "QList<QFileInfo>"
		}

	case "QWidgetList":
		{
			f.Output = "QList<QWidget *>"
		}

	case "QCameraFocusZoneList":
		{
			//f.Output = "QList<QCameraFocusZone *>" //TODO: uncomment
		}

	case "QList<T>":
		{
			f.TemplateModeGo = "QObject*"
			f.Output = "QList<QObject*>"
		}

	case "T":
		{
			switch className := f.ClassName(); className {
			case "QObject", "QMediaService":
				{
					f.TemplateModeGo = fmt.Sprintf("%v*", className)
					f.Output = fmt.Sprintf("%v*", className)
				}
			}
		}
	}

	if !IsPackedList(f.Output) {
		return
	}

	var class, exist = f.Class()
	if !exist || class.HasFunctionWithName(fmt.Sprintf("%v_atList", f.Name)) {
		return
	}

	for _, p := range f.Parameters {
		if strings.ContainsAny(p.Value, "<>") {
			return
		}
	}

	class.Functions = append(class.Functions, &Function{
		Name:       fmt.Sprintf("%v_atList", f.Name),
		Fullname:   fmt.Sprintf("%v::%v_atList", class.Name, f.Name),
		Access:     "public",
		Virtual:    "non",
		Meta:       PLAIN,
		Output:     fmt.Sprintf("const %v", strings.Split(strings.Split(f.Output, "<")[1], ">")[0]),
		Parameters: []*Parameter{{Name: "i", Value: "int"}},
		Signature:  "()",
		Container:  strings.Split(f.Output, "<")[0],
	})
}

func (f *Function) fixTemporary() {

	//TODO: needed until input + cgo support for generic Lists/Containers<T>
	if !IsPackedList(f.Output) || f.Virtual == PURE {
		return
	}
	f.Virtual = "non"
	f.Meta = PLAIN
}
