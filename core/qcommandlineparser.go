package core

//#include "qcommandlineparser.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QCommandLineParser struct {
	ptr unsafe.Pointer
}

type QCommandLineParserITF interface {
	QCommandLineParserPTR() *QCommandLineParser
}

func (p *QCommandLineParser) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCommandLineParser) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCommandLineParser(ptr QCommandLineParserITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLineParserPTR().Pointer()
	}
	return nil
}

func QCommandLineParserFromPointer(ptr unsafe.Pointer) *QCommandLineParser {
	var n = new(QCommandLineParser)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCommandLineParser) QCommandLineParserPTR() *QCommandLineParser {
	return ptr
}

//QCommandLineParser::SingleDashWordOptionMode
type QCommandLineParser__SingleDashWordOptionMode int

var (
	QCommandLineParser__ParseAsCompactedShortOptions = QCommandLineParser__SingleDashWordOptionMode(0)
	QCommandLineParser__ParseAsLongOptions           = QCommandLineParser__SingleDashWordOptionMode(1)
)

func NewQCommandLineParser() *QCommandLineParser {
	return QCommandLineParserFromPointer(unsafe.Pointer(C.QCommandLineParser_NewQCommandLineParser()))
}

func (ptr *QCommandLineParser) AddOption(option QCommandLineOptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QCommandLineParser_AddOption(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCommandLineOption(option))) != 0
	}
	return false
}

func (ptr *QCommandLineParser) AddPositionalArgument(name string, description string, syntax string) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_AddPositionalArgument(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.CString(description), C.CString(syntax))
	}
}

func (ptr *QCommandLineParser) ApplicationDescription() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_ApplicationDescription(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLineParser) ClearPositionalArguments() {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_ClearPositionalArguments(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCommandLineParser) ErrorText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_ErrorText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLineParser) HelpText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_HelpText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLineParser) IsSet2(option QCommandLineOptionITF) bool {
	if ptr.Pointer() != nil {
		return C.QCommandLineParser_IsSet2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCommandLineOption(option))) != 0
	}
	return false
}

func (ptr *QCommandLineParser) IsSet(name string) bool {
	if ptr.Pointer() != nil {
		return C.QCommandLineParser_IsSet(C.QtObjectPtr(ptr.Pointer()), C.CString(name)) != 0
	}
	return false
}

func (ptr *QCommandLineParser) OptionNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_OptionNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Parse(arguments []string) bool {
	if ptr.Pointer() != nil {
		return C.QCommandLineParser_Parse(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(arguments, "|"))) != 0
	}
	return false
}

func (ptr *QCommandLineParser) PositionalArguments() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_PositionalArguments(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Process2(app QCoreApplicationITF) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_Process2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCoreApplication(app)))
	}
}

func (ptr *QCommandLineParser) Process(arguments []string) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_Process(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(arguments, "|")))
	}
}

func (ptr *QCommandLineParser) SetApplicationDescription(description string) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_SetApplicationDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QCommandLineParser) SetSingleDashWordOptionMode(singleDashWordOptionMode QCommandLineParser__SingleDashWordOptionMode) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_SetSingleDashWordOptionMode(C.QtObjectPtr(ptr.Pointer()), C.int(singleDashWordOptionMode))
	}
}

func (ptr *QCommandLineParser) ShowHelp(exitCode int) {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_ShowHelp(C.QtObjectPtr(ptr.Pointer()), C.int(exitCode))
	}
}

func (ptr *QCommandLineParser) ShowVersion() {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_ShowVersion(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QCommandLineParser) UnknownOptionNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_UnknownOptionNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Value2(option QCommandLineOptionITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_Value2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCommandLineOption(option))))
	}
	return ""
}

func (ptr *QCommandLineParser) Value(optionName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineParser_Value(C.QtObjectPtr(ptr.Pointer()), C.CString(optionName)))
	}
	return ""
}

func (ptr *QCommandLineParser) Values2(option QCommandLineOptionITF) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_Values2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCommandLineOption(option)))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) Values(optionName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineParser_Values(C.QtObjectPtr(ptr.Pointer()), C.CString(optionName))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineParser) DestroyQCommandLineParser() {
	if ptr.Pointer() != nil {
		C.QCommandLineParser_DestroyQCommandLineParser(C.QtObjectPtr(ptr.Pointer()))
	}
}
