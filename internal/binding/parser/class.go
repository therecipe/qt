package parser

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/StarAurryon/qt/internal/utils"
)

type Class struct {
	Name       string      `xml:"name,attr"`
	Status     string      `xml:"status,attr"`
	Access     string      `xml:"access,attr"`
	Abstract   bool        `xml:"abstract,attr"`
	Bases      string      `xml:"bases,attr"`
	Module     string      `xml:"module,attr"`
	Brief      string      `xml:"brief,attr"`
	Functions  []*Function `xml:"function"`
	Enums      []*Enum     `xml:"enum"`
	Variables  []*Variable `xml:"variable"`
	Properties []*Variable `xml:"property"`
	Classes    []*Class    `xml:"class"`
	Since      string      `xml:"since,attr"`
	Filepath   string      `xml:"filepath,attr"`

	DocModule    string
	Stub         bool
	WeakLink     map[string]struct{}
	Export       bool
	Fullname     string
	Pkg          string
	Path         string
	HasFinalizer bool

	Constructors []string
	Derivations  []string

	ToBeCleanedUp bool

	sync.Mutex
}

func (c *Class) register(m *Module) {
	c.DocModule = c.Module
	c.Module = m.Project
	c.Pkg = m.Pkg
	State.ClassMap[c.Name] = c

	for _, sc := range c.Classes {
		if sc.Name != "PaintContext" { //TODO: remove and support all sub classes
			continue
		}

		sc.Fullname = fmt.Sprintf("%v::%v", c.Name, sc.Name)
		sc.register(m)
	}
}

func (c *Class) derivation() {
	for _, b := range c.GetBases() {
		if bc, e := State.ClassMap[b]; e {
			bc.Derivations = append(bc.Derivations, c.Name)
		}
	}
}

func (c *Class) GetBases() []string {
	if c == nil || c.Bases == "" {
		return make([]string, 0)
	}
	if strings.Contains(c.Bases, ",") {
		return strings.Split(c.Bases, ",")
	}
	return []string{c.Bases}
}

func (c *Class) GetBasesSorted() []string {
	b := c.GetBases()
	sort.SliceStable(b, func(i int, j int) bool {
		return (b[j] == "QObject") || (len(State.ClassMap[b[i]].GetAllBases()) > len(State.ClassMap[b[j]].GetAllBases()))
	})
	return b
}

func (c *Class) GetAllBases() []string {
	var out, _ = c.GetAllBasesRecursiveCheckFailed(0)
	return out
}

func (c *Class) GetAllBasesRecursiveCheckFailed(i int) ([]string, bool) {
	var input = make([]string, 0)

	i++
	if i > 100 {
		return input, true
	}

	for _, b := range c.GetBases() {
		var bc, ok = State.ClassMap[b]
		if !ok {
			continue
		}

		input = append(input, b)
		var bs, isRecursive = bc.GetAllBasesRecursiveCheckFailed(i)
		if isRecursive {
			return input, true
		}
		input = append(input, bs...)
	}

	return input, false
}

func (c *Class) IsSubClassOfQObject() bool {
	return c.IsSubClassOf("QObject")
}

func (c *Class) IsSubClassOf(class string) bool {
	if c == nil {
		return false
	}

	for _, bcn := range append([]string{c.Name}, c.GetAllBases()...) {
		if bcn == class {
			return true
		}
	}

	return false
}

func (c *Class) isSubClass() bool { return c.Fullname != "" }

func (c *Class) GetAllDerivations() []string {

	var input = make([]string, 0)

	for _, b := range c.Derivations {
		var bc, exists = State.ClassMap[b]
		if !exists {
			continue
		}

		input = append(append(input, b), bc.GetAllDerivations()...)
	}

	return input
}

func (c *Class) GetAllDerivationsInSameModule() []string {

	var input = make([]string, 0)

	for _, i := range c.GetAllDerivations() {
		if State.ClassMap[i].Module == c.Module && i != "QWinEventNotifier" && i != "QFutureWatcher" {
			input = append(input, i)
		}
	}

	return input
}

func (c *Class) HasFunction(f *Function) bool {
	for _, cf := range c.Functions {
		if cf.Name == f.Name && cf.Virtual == f.Virtual &&
			cf.Meta == f.Meta &&
			cf.Output == f.Output && len(cf.Parameters) == len(f.Parameters) {

			var similar = true
			for i, cfp := range cf.Parameters {
				if cfp.Value != f.Parameters[i].Value {
					similar = false
				}
			}
			if similar {
				return true
			}
		}
	}

	return false
}

func (c *Class) HasFunctionWithName(n string) bool {
	return c.HasFunctionWithNameAndOverloadNumber(n, "")
}

func (c *Class) HasFunctionWithNameAndOverloadNumber(n string, num string) bool {
	for _, f := range c.Functions {
		if ((strings.ToLower(f.Name) == strings.ToLower(n) && c.Module != MOC) ||
			(f.Name == n && c.Module == MOC)) &&
			f.OverloadNumber == num {
			return true
		}
	}
	return false
}

func (c *Class) IsPolymorphic() bool { return len(c.GetBases()) > 1 }

func (c *Class) HasConstructor() bool {
	for _, f := range c.Functions {
		if f.Meta == CONSTRUCTOR || f.Meta == COPY_CONSTRUCTOR || f.Meta == MOVE_CONSTRUCTOR {
			return true
		}
	}
	return false
}

func (c *Class) HasDestructor() bool {
	for _, f := range c.Functions {
		if f.Meta == DESTRUCTOR {
			return true
		}
	}
	return false
}

func (c *Class) HasCallbackFunctions() bool {
	for _, bcn := range append([]string{c.Name}, c.GetAllBases()...) {
		var bc, ok = State.ClassMap[bcn]
		if !ok {
			continue
		}
		for _, f := range bc.Functions {
			if f.Virtual == IMPURE || f.Virtual == PURE || f.Meta == SIGNAL || f.Meta == SLOT {
				return true
			}
		}
	}
	return false
}

func (c *Class) HasCallbackFunctionsBesideTheDestructor() bool {
	for _, bcn := range append([]string{c.Name}, c.GetAllBases()...) {
		var bc, ok = State.ClassMap[bcn]
		if !ok {
			continue
		}
		for _, f := range bc.Functions {
			if (f.Virtual == IMPURE || f.Virtual == PURE || f.Meta == SIGNAL || f.Meta == SLOT) && f.Meta != DESTRUCTOR {
				return true
			}
		}
	}
	return false
}

func (c *Class) IsSupported() bool {
	if c == nil {
		return false
	}

	if strings.Contains(c.Since, ".") {
		version := strings.TrimPrefix(c.Since, "Qt")
		vmaj, _ := strconv.Atoi(string(version[0]))
		vmin, _ := strconv.Atoi(strings.Replace(version[1:], ".", "", -1))
		if vmaj*1e3+vmin*10 > utils.QT_VERSION_NUM() {
			c.Access = "unsupported_isApiVersionBlockedClass"
			return false
		}
	}

	switch c.Name {
	case "QCborStreamReader", "QCborStreamWriter", "QCborValue", "QScopeGuard", "QTest",
		"QImageReaderWriterHelpers", "QPasswordDigestor", "QDtls", "QDtlsClientVerifier", "QGeoJson",
		"QQuickPopupItem", "QQuickDeferredPointer", "QQuickPopupPositioner", "QQuickVelocityCalculator",
		"QQuickPressHandler", "QQuickStackElement", "QQuickPopupTransitionManager":
		c.Access = "unsupported_isBlockedClass"
		return false
	}

	if utils.QT_VERSION_NUM() >= 5080 {
		switch c.Name {
		case "QSctpServer", "QSctpSocket", "Http2", "QAbstractExtensionFactory":
			{
				c.Access = "unsupported_isBlockedClass"
				return false
			}
		}
	}

	if UseJs() {
		if strings.HasPrefix(c.Name, "QOpenGLFunctions_") || strings.HasPrefix(c.Name, "QSsl") || strings.HasPrefix(c.Name, "QNetwork") {
			c.Access = "unsupported_isBlockedClass"
			return false
		}
		switch c.Name {
		case "QThreadPool", "QSharedMemory", "QPluginLoader", "QSemaphore", "QSemaphoreReleaser",
			"QSystemSemaphore", "QThread", "QWaitCondition", "QUnhandledException", "QFileSystemModel",
			"QLibrary", "QUdpSocket", "QHttpMultiPart", "QHttpPart", "QOpenGLTimeMonitor", "QOpenGLTimerQuery",
			"QRemoteObjectHost": //TODO: only block in 5.11 ?
			{
				c.Access = "unsupported_isBlockedClass"
				return false
			}
		}
	}

	switch c.Name {
	case
		"QString", "QStringList", //mapped to primitive

		"QExplicitlySharedDataPointer", "QFuture", "QDBusPendingReply", "QDBusReply", "QFutureSynchronizer", //needs template
		"QGlobalStatic", "QMultiHash", "QQueue", "QMultiMap", "QScopedPointer", "QSharedDataPointer",
		"QScopedArrayPointer", "QSharedPointer", "QThreadStorage", "QScopedValueRollback", "QVarLengthArray",
		"QWeakPointer", "QWinEventNotifier",

		"QFlags", "QException", "QStandardItemEditorCreator", "QSGSimpleMaterialShader", "QGeoCodeReply", "QFutureWatcher", //other
		"QItemEditorCreator", "QGeoCodingManager", "QGeoCodingManagerEngine", "QQmlListProperty",

		"QPlatformGraphicsBuffer", "QPlatformSystemTrayIcon", "QRasterPaintEngine", "QSupportedWritingSystems", "QGeoLocation", //file not found or QPA API
		"QAbstractOpenGLFunctions",

		"QRemoteObjectPackets",

		"QStaticByteArrayMatcher", "QtDummyFutex", "QtLinuxFutex",
		"QShaderLanguage",

		"AndroidNfc", "OSXBluetooth",

		"QtROClientFactory", "QtROServerFactory",

		"QWebViewFactory", "QGeoServiceProviderFactoryV2",

		"QtDwmApiDll", "QWinMime",

		"QAbstract3DGraph", //TODO: only block for arch with pkg_config

		"QQuickFolderListModel", "QFileDialogOptions", "QMessageDialogOptions",
		"QGeoServiceProviderFactoryV3", "QTextToSpeechProcessorFlite", "QOpenVGMatrix",
		"QSvgIOHandler", "QSvgIconEngine", "QQuickProfilerAdapter",
		"QWavefrontMesh", "QM3uPlaylistPlugin", "QOpenSLESAudioInput",
		"QSGSimpleMaterialComparableMaterial", "QGStreamerAvailabilityControl",
		"QGstreamerV4L2Input", "QSignalMapper",

		"QGeoPositionInfoSourceFactoryV2":
		{
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	if strings.HasPrefix(c.Name, "QProcess") && strings.HasPrefix(State.Target, "ios") {
		c.Access = "unsupported_isBlockedClass"
		return false
	}

	for _, cn := range []string{"QTextToSpeechPlugin", "QTextToSpeechEngine"} {
		if strings.HasPrefix(c.Name, cn) && c.Name != cn {
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	switch {
	case
		c.Name == "QOpenGLFunctions_ES2", strings.HasPrefix(c.Name, "QPlace"), //file not found or QPA API

		strings.HasPrefix(c.Name, "QAtomic"), //other

		strings.HasSuffix(c.Name, "terator") && c.Name != "QJSValueIterator", strings.Contains(c.Brief, "emplate"), //needs template

		strings.HasPrefix(c.Name, "QVulkan"),

		!strings.HasPrefix(c.Name, "Q") && strings.HasPrefix(c.Module, "Qt") && c.Module != "QtSailfish" && c.Module != "QtFelgo" && c.Name != "PaintContext" && c.Name != "RawHeader",
		strings.HasPrefix(c.Name, "Qml") && c.Module == "QtSensors",
		strings.HasPrefix(c.Name, "QQml") && (c.Module == "QtQuick" || c.Module == "QtWebSockets"),
		strings.HasPrefix(c.Name, "QAndroid") && strings.HasPrefix(c.Module, "QtMultimedia"),
		strings.HasPrefix(c.Name, "QDesigner") && !(strings.HasSuffix(c.Name, "Interface") || strings.HasSuffix(c.Name, "Extension")) && c.Module == "QtDesigner",
		strings.HasPrefix(c.Name, "QV4") && c.Module == "QtQuick",
		strings.HasPrefix(c.Name, "QSG") && (strings.HasPrefix(c.Module, "QtMultimedia") || (strings.Contains(c.Name, "Cache") && c.Module == "QtQuick")),
		strings.HasPrefix(c.Name, "QSGOpenVG") && c.Module == "QtQuick",
		strings.HasPrefix(c.Name, "QPlatform") && (c.Module == "QtGui" || c.Module == "QtPrintSupport") && c.Name != "QPlatformSurfaceEvent",
		strings.HasPrefix(c.Name, "QAlsa") && strings.HasPrefix(c.Module, "QtMultimedia"),
		strings.Contains(c.Name, "Qnx") && (strings.HasPrefix(c.Module, "QtMultimedia") || c.Module == "QtRemoteObjects"),
		(strings.HasPrefix(c.Name, "QGstreamer") || strings.HasPrefix(c.Name, "QWin") ||
			strings.HasPrefix(c.Name, "QPulse") || strings.HasPrefix(c.Name, "QOpenSLES") ||
			strings.HasPrefix(c.Name, "QWasapi")) && strings.HasPrefix(c.Module, "QtMultimedia"):
		{
			if c.Module != MOC {
				c.Access = "unsupported_isBlockedClass"
				return false
			}
		}
	}

	if strings.HasPrefix(c.Name, "QOpenGL") && (os.Getenv("DEB_TARGET_ARCH_CPU") == "arm" || (strings.HasPrefix(State.Target, "android") && utils.QT_FAT())) { //TODO: block indiv classes for fat android build instead
		c.Access = "unsupported_isBlockedClass"
		return false
	}

	if utils.QT_VERSION_NUM() <= 5042 {
		if c.Name == "QQmlAbstractProfilerAdapter" ||
			c.Name == "QGraphicsLayout" ||
			c.Name == "QQmlAbstractProfilerAdapter" ||
			c.Name == "QDeclarativeMultimedia" {
			c.Access = "unsupported_isBlockedClass"
			return false
		}
	}

	if strings.HasPrefix(c.Name, "QOpenGLFunctions_") && !utils.QT_GEN_OPENGL() {
		c.Access = "unsupported_isBlockedClass"
		return false
	}

	if State.Minimal {
		return c.Export
	}

	return true
}

func (c *Class) GetFunction(fname string) *Function {
	for _, f := range c.Functions {
		if f.Name == fname {
			return f
		}
	}
	return nil
}

func (c *Class) GetTitledFunction(fname string) *Function {
	for _, f := range c.Functions {
		if strings.Title(f.Name) == strings.Title(fname) {
			return f
		}
	}
	return nil
}

//

func (c *Class) Hash() string {
	h := sha1.New()
	h.Write([]byte(c.Path))
	return hex.EncodeToString(h.Sum(nil)[:3])
}
