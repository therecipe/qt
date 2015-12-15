package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QCommandLineParser struct {
	ptr unsafe.Pointer
}

type QCommandLineParser_ITF interface {
	QCommandLineParser_PTR() *QCommandLineParser
}

func (p *QCommandLineParser) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCommandLineParser) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCommandLineParser(ptr QCommandLineParser_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLineParser_PTR().Pointer()
	}
	return nil
}

func NewQCommandLineParserFromPointer(ptr unsafe.Pointer) *QCommandLineParser {
	var n = new(QCommandLineParser)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCommandLineParser) QCommandLineParser_PTR() *QCommandLineParser {
	return ptr
}

//QCommandLineParser::SingleDashWordOptionMode
type QCommandLineParser__SingleDashWordOptionMode int64

const (
	QCommandLineParser__ParseAsCompactedShortOptions = QCommandLineParser__SingleDashWordOptionMode(0)
	QCommandLineParser__ParseAsLongOptions           = QCommandLineParser__SingleDashWordOptionMode(1)
)

func NewQCommandLineParser() *QCommandLineParser {
	defer qt.Recovering("QCommandLineParser::QCommandLineParser")

	return NewQCommandLineParserFromPointer(C.QCommandLineParser_NewQCommandLineParser())
}

func (ptr *QCommandLineParser) AddHelpOption() *QCommandLineOption {
	defer qt.Recovering("QCommandLineParser::addHelpOption")

	if ptr.Pointer() != nil {
		return NewQCommandLineOptionFromPointer(C.QCommandLineParser_AddHelpOption(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCommandLineParser) AddOption(option QCommandLineOption_ITF) bool {
	defer qt.Recovering("QCommandLineParser::addOption")

	if ptr.Pointer() != nil {
		return C.QCommandLineParser_AddOption(ptr.Pointer(), PointerFromQCommandLineOption(option)) != 0
	}
	return false
}

func (ptr *QCommandLineParser) AddPositionalArgument(name string, description string, syntax string) {
	defer qt.Recovering("QCommandLineParser::addPositionalArgument")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_AddPositionalArgument(ptr.Pointer(), C.CString(name), C.CString(description), C.CString(syntax))
	}
}

func (ptr *QCommandLineParser) AddVersionOption() *QCommandLineOption {
	defer qt.Recovering("QCommandLineParser::addVersionOption")

	if ptr.Pointer() != nil {
		return NewQCommandLineOptionFromPointer(C.QCommandLineParser_AddVersionOption(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCommandLineParser) ApplicationDescription() string {
	defer qt.Recovering("QCommandLineParser::applicationDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_ApplicationDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLineParser) ClearPositionalArguments() {
	defer qt.Recovering("QCommandLineParser::clearPositionalArguments")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_ClearPositionalArguments(ptr.Pointer())
	}
}

func (ptr *QCommandLineParser) ErrorText() string {
	defer qt.Recovering("QCommandLineParser::errorText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_ErrorText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLineParser) HelpText() string {
	defer qt.Recovering("QCommandLineParser::helpText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_HelpText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLineParser) IsSet2(option QCommandLineOption_ITF) bool {
	defer qt.Recovering("QCommandLineParser::isSet")

	if ptr.Pointer() != nil {
		return C.QCommandLineParser_IsSet2(ptr.Pointer(), PointerFromQCommandLineOption(option)) != 0
	}
	return false
}

func (ptr *QCommandLineParser) IsSet(name string) bool {
	defer qt.Recovering("QCommandLineParser::isSet")

	if ptr.Pointer() != nil {
		return C.QCommandLineParser_IsSet(ptr.Pointer(), C.CString(name)) != 0
	}
	return false
}

func (ptr *QCommandLineParser) OptionNames() []string {
	defer qt.Recovering("QCommandLineParser::optionNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_OptionNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Parse(arguments []string) bool {
	defer qt.Recovering("QCommandLineParser::parse")

	if ptr.Pointer() != nil {
		return C.QCommandLineParser_Parse(ptr.Pointer(), C.CString(strings.Join(arguments, ",,,"))) != 0
	}
	return false
}

func (ptr *QCommandLineParser) PositionalArguments() []string {
	defer qt.Recovering("QCommandLineParser::positionalArguments")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_PositionalArguments(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Process2(app QCoreApplication_ITF) {
	defer qt.Recovering("QCommandLineParser::process")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_Process2(ptr.Pointer(), PointerFromQCoreApplication(app))
	}
}

func (ptr *QCommandLineParser) Process(arguments []string) {
	defer qt.Recovering("QCommandLineParser::process")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_Process(ptr.Pointer(), C.CString(strings.Join(arguments, ",,,")))
	}
}

func (ptr *QCommandLineParser) SetApplicationDescription(description string) {
	defer qt.Recovering("QCommandLineParser::setApplicationDescription")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_SetApplicationDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QCommandLineParser) SetSingleDashWordOptionMode(singleDashWordOptionMode QCommandLineParser__SingleDashWordOptionMode) {
	defer qt.Recovering("QCommandLineParser::setSingleDashWordOptionMode")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_SetSingleDashWordOptionMode(ptr.Pointer(), C.int(singleDashWordOptionMode))
	}
}

func (ptr *QCommandLineParser) ShowHelp(exitCode int) {
	defer qt.Recovering("QCommandLineParser::showHelp")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_ShowHelp(ptr.Pointer(), C.int(exitCode))
	}
}

func (ptr *QCommandLineParser) ShowVersion() {
	defer qt.Recovering("QCommandLineParser::showVersion")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_ShowVersion(ptr.Pointer())
	}
}

func (ptr *QCommandLineParser) UnknownOptionNames() []string {
	defer qt.Recovering("QCommandLineParser::unknownOptionNames")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_UnknownOptionNames(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Value2(option QCommandLineOption_ITF) string {
	defer qt.Recovering("QCommandLineParser::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_Value2(ptr.Pointer(), PointerFromQCommandLineOption(option)))
	}
	return ""
}

func (ptr *QCommandLineParser) Value(optionName string) string {
	defer qt.Recovering("QCommandLineParser::value")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_Value(ptr.Pointer(), C.CString(optionName)))
	}
	return ""
}

func (ptr *QCommandLineParser) Values2(option QCommandLineOption_ITF) []string {
	defer qt.Recovering("QCommandLineParser::values")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_Values2(ptr.Pointer(), PointerFromQCommandLineOption(option))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Values(optionName string) []string {
	defer qt.Recovering("QCommandLineParser::values")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_Values(ptr.Pointer(), C.CString(optionName))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) DestroyQCommandLineParser() {
	defer qt.Recovering("QCommandLineParser::~QCommandLineParser")

	if ptr.Pointer() != nil {
		C.QCommandLineParser_DestroyQCommandLineParser(ptr.Pointer())
	}
}
