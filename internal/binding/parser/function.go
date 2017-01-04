package parser

import (
	"sort"
	"strings"
)

type Function struct {
	Name             string       `xml:"name,attr"`
	Fullname         string       `xml:"fullname,attr"`
	Href             string       `xml:"href,attr"`
	Status           string       `xml:"status,attr"`
	Access           string       `xml:"access,attr"`
	Filepath         string       `xml:"filepath,attr"`
	Virtual          string       `xml:"virtual,attr"`
	Meta             string       `xml:"meta,attr"`
	Static           bool         `xml:"static,attr"`
	Overload         bool         `xml:"overload,attr"`
	OverloadNumber   string       `xml:"overload-number,attr"`
	Output           string       `xml:"type,attr"`
	Signature        string       `xml:"signature,attr"`
	Parameters       []*Parameter `xml:"parameter"`
	Brief            string       `xml:"brief,attr"`
	SignalMode       string
	TemplateModeJNI  string
	Default          bool
	TmpName          string
	Export           bool
	NeedsFinalizer   bool
	Container        string
	TemplateModeGo   string
	Child            *Function
	NonMember        bool
	NoMocDeduce      bool
	PureBaseFunction bool
	AsError          bool
}

type Parameter struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"left,attr"`
}

func (f *Function) Class() (*Class, bool) {
	var class, exist = State.ClassMap[f.ClassName()]
	return class, exist
}

func (f *Function) ClassName() string {
	var s = strings.Split(f.Fullname, "::")
	if len(s) == 3 {
		return s[1]
	}
	return s[0]
}

//TODO: multipoly [][]string
//TODO: connect/disconnect slot functions + add necessary SIGNAL_* functions (check first if really needed)
//TODO: replace self poly deduction with overridden methods ?

func (f *Function) PossiblePolymorphic(self bool) ([]string, string) {
	var out = make([]string, 0)

	var params = func() []*Parameter {
		if self {
			return []*Parameter{{Name: "ptr", Value: f.ClassName()}}
		}
		return f.Parameters
	}()

	var fc, _ = f.Class()

	for _, p := range params {
		var c, exist = State.ClassMap[CleanValue(p.Value)]
		if !exist {
			continue
		}

		for _, class := range State.ClassMap {
			if class.Module != fc.Module {
				continue
			}

			if class.IsPolymorphic() && class.IsSubClassOf(c.Name) {
				out = append(out, class.Name)
			}
		}

		//TODO: multipoly
		if len(out) > 0 {
			sort.Stable(sort.StringSlice(out))
			out = append(out, c.Name)
			return out, CleanName(p.Name, p.Value)
		}
	}

	return out, ""
}

func (f *Function) IsJNIGeneric() bool {

	if f.ClassName() == "QAndroidJniObject" {
		switch f.Name {
		case
			"callMethod",
			"callStaticMethod",

			"getField",
			//"setField", -> uses interface{} if not generic

			"getStaticField",
			//"setStaticField", -> uses interface{} if not generic

			"getObjectField",

			"getStaticObjectField",

			"callObjectMethod",
			"callStaticObjectMethod":
			{
				return true
			}

		case "setStaticField":
			{
				if f.OverloadNumber == "2" || f.OverloadNumber == "4" {
					return true
				}
			}
		}
	}

	return false
}

func (f *Function) IsSupported() bool {

	switch {
	case
		(f.ClassName() == "QAccessibleObject" || f.ClassName() == "QAccessibleInterface" || f.ClassName() == "QAccessibleWidget" || //QAccessible::State -> quint64
			f.ClassName() == "QAccessibleStateChangeEvent") && (f.Name == "state" || f.Name == "changedStates" || f.Name == "m_changedStates" || f.Name == "setM_changedStates" || f.Meta == CONSTRUCTOR),

		f.Fullname == "QPixmapCache::find" && f.OverloadNumber == "4", //Qt::Key -> int
		(f.Fullname == "QPixmapCache::remove" || f.Fullname == "QPixmapCache::insert") && f.OverloadNumber == "2",
		f.Fullname == "QPixmapCache::replace",

		f.Fullname == "QNdefFilter::appendRecord" && !f.Overload, //QNdefRecord::TypeNameFormat -> uint

		f.ClassName() == "QSimpleXmlNodeModel" && f.Meta == CONSTRUCTOR,

		f.Fullname == "QSGMaterialShader::attributeNames",

		f.ClassName() == "QVariant" && (f.Name == "value" || f.Name == "canConvert"), //needs template

		f.Fullname == "QNdefRecord::isRecordType", f.Fullname == "QScriptEngine::scriptValueFromQMetaObject", //needs template
		f.Fullname == "QScriptEngine::fromScriptValue", f.Fullname == "QJSEngine::fromScriptValue",

		f.ClassName() == "QMetaType" && //needs template
			(f.Name == "hasRegisteredComparators" || f.Name == "registerComparators" ||
				f.Name == "hasRegisteredConverterFunction" || f.Name == "registerConverter" ||
				f.Name == "registerEqualsComparator"),

		State.ClassMap[f.ClassName()].Module == MOC && f.Name == "metaObject", //needed for qtmoc

		f.Fullname == "QSignalBlocker::QSignalBlocker" && f.OverloadNumber == "3", //undefined symbol

		(State.ClassMap[f.ClassName()].IsSubClassOf("QCoreApplication") ||
			f.ClassName() == "QAudioInput" || f.ClassName() == "QAudioOutput") && f.Name == "notify", //redeclared (name collision with QObject)

		f.Fullname == "QGraphicsItem::isBlockedByModalPanel", //** problem

		f.Name == "surfaceHandle", //QQuickWindow && QQuickView //unsupported_cppType(QPlatformSurface)

		f.Name == "readData", f.Name == "QNetworkReply", //TODO: char*

		f.Name == "QDesignerFormWindowInterface" || f.Name == "QDesignerFormWindowManagerInterface" || f.Name == "QDesignerWidgetBoxInterface", //unimplemented virtual

		f.Fullname == "QNdefNfcSmartPosterRecord::titleRecords", //T<T> output with unsupported output for *_atList
		f.Fullname == "QHelpEngineCore::filterAttributeSets", f.Fullname == "QHelpSearchEngine::query", f.Fullname == "QHelpSearchQueryWidget::query",
		f.Fullname == "QPluginLoader::staticPlugins", f.Fullname == "QSslConfiguration::ellipticCurves", f.Fullname == "QSslConfiguration::supportedEllipticCurves",
		f.Fullname == "QTextFormat::lengthVectorProperty", f.Fullname == "QTextTableFormat::columnWidthConstraints", f.Fullname == "QHelpContentWidget::selectedIndexes",

		f.Fullname == "QListView::indexesMoved", f.Fullname == "QAudioInputSelectorControl::availableInputs", f.Fullname == "QScxmlStateMachine::initialValuesChanged",
		f.Fullname == "QAudioOutputSelectorControl::availableOutputs", f.Fullname == "QQuickWebEngineProfile::downloadFinished",
		f.Fullname == "QQuickWindow::closing", f.Fullname == "QQuickWebEngineProfile::downloadRequested", f.Fullname == "QWebEnginePage::fullScreenRequested",

		strings.Contains(f.Access, "unsupported"):
		{
			if !strings.Contains(f.Access, "unsupported") {
				f.Access = "unsupported_isBlockedFunction"
			}
			return false
		}
	}

	if strings.ContainsAny(f.Signature, "<>") {
		if IsPackedList(f.Output) {
			for _, p := range f.Parameters {
				if strings.ContainsAny(p.Value, "<>") {
					if !strings.Contains(f.Access, "unsupported") {
						f.Access = "unsupported_isBlockedFunction"
					}
					return false
				}
			}
		} else {
			if !strings.Contains(f.Access, "unsupported") {
				f.Access = "unsupported_isBlockedFunction"
			}
			return false
		}
	}

	if f.Default {
		return f.IsSupportedDefault()
	}

	if State.Minimal {
		return f.Export || f.Meta == DESTRUCTOR || f.Fullname == "QObject::destroyed" || strings.HasPrefix(f.Name, TILDE)
	}

	return true
}

func (f *Function) IsSupportedDefault() bool {
	switch f.Fullname {
	case
		"QAnimationGroup::updateCurrentTime", "QAnimationGroup::duration",
		"QAbstractProxyModel::columnCount", "QAbstractTableModel::columnCount",
		"QAbstractListModel::data", "QAbstractTableModel::data",
		"QAbstractProxyModel::index", "QAbstractProxyModel::parent",
		"QAbstractListModel::rowCount", "QAbstractProxyModel::rowCount", "QAbstractTableModel::rowCount",

		"QPagedPaintDevice::paintEngine", "QAccessibleObject::childCount",
		"QAccessibleObject::indexOfChild", "QAccessibleObject::role",
		"QAccessibleObject::text", "QAccessibleObject::child",
		"QAccessibleObject::parent",

		"QAbstractGraphicsShapeItem::paint", "QGraphicsObject::paint",
		"QLayout::sizeHint", "QAbstractGraphicsShapeItem::boundingRect",
		"QGraphicsObject::boundingRect", "QGraphicsLayout::sizeHint",

		"QSimpleXmlNodeModel::typedValue", "QSimpleXmlNodeModel::documentUri",
		"QSimpleXmlNodeModel::compareOrder", "QSimpleXmlNodeModel::nextFromSimpleAxis",
		"QSimpleXmlNodeModel::kind", "QSimpleXmlNodeModel::root",

		"QAbstractPlanarVideoBuffer::unmap", "QAbstractPlanarVideoBuffer::mapMode",

		"QSGDynamicTexture::bind", "QSGDynamicTexture::hasMipmaps",
		"QSGDynamicTexture::textureSize", "QSGDynamicTexture::hasAlphaChannel",
		"QSGDynamicTexture::textureId",

		"QModbusClient::open", "QModbusClient::close", "QModbusServer::open", "QModbusServer::close",

		"QSimpleXmlNodeModel::name",

		"QSimpleXmlNodeModel::attributes", "QAbstractXmlNodeModel::attributes":

		{
			return false
		}
	}

	//needed for moc
	if State.ClassMap[f.ClassName()].IsSubClassOf("QCoreApplication") && (f.Name == "autoMaximizeThreshold" || f.Name == "setAutoMaximizeThreshold") {
		return false
	}

	if State.Minimal {
		return f.Export
	}

	return true
}

func (f *Function) IsDerivedFromVirtual() bool {
	var class, _ = f.Class()

	for _, bc := range class.GetAllBases() {
		if bclass, exists := State.ClassMap[bc]; exists {
			for _, bcf := range bclass.Functions {
				if f.Name == bcf.Name && bcf.Virtual != "non" {
					return true
				}
			}
		}
	}

	return false
}
