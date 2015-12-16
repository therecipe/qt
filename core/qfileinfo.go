package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFileInfo struct {
	ptr unsafe.Pointer
}

type QFileInfo_ITF interface {
	QFileInfo_PTR() *QFileInfo
}

func (p *QFileInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFileInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFileInfo(ptr QFileInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileInfo_PTR().Pointer()
	}
	return nil
}

func NewQFileInfoFromPointer(ptr unsafe.Pointer) *QFileInfo {
	var n = new(QFileInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileInfo) QFileInfo_PTR() *QFileInfo {
	return ptr
}

func NewQFileInfo() *QFileInfo {
	defer qt.Recovering("QFileInfo::QFileInfo")

	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo())
}

func NewQFileInfo5(dir QDir_ITF, file string) *QFileInfo {
	defer qt.Recovering("QFileInfo::QFileInfo")

	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo5(PointerFromQDir(dir), C.CString(file)))
}

func NewQFileInfo4(file QFile_ITF) *QFileInfo {
	defer qt.Recovering("QFileInfo::QFileInfo")

	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo4(PointerFromQFile(file)))
}

func NewQFileInfo6(fileinfo QFileInfo_ITF) *QFileInfo {
	defer qt.Recovering("QFileInfo::QFileInfo")

	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo6(PointerFromQFileInfo(fileinfo)))
}

func NewQFileInfo3(file string) *QFileInfo {
	defer qt.Recovering("QFileInfo::QFileInfo")

	return NewQFileInfoFromPointer(C.QFileInfo_NewQFileInfo3(C.CString(file)))
}

func (ptr *QFileInfo) AbsoluteDir() *QDir {
	defer qt.Recovering("QFileInfo::absoluteDir")

	if ptr.Pointer() != nil {
		return NewQDirFromPointer(C.QFileInfo_AbsoluteDir(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) AbsoluteFilePath() string {
	defer qt.Recovering("QFileInfo::absoluteFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsoluteFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) AbsolutePath() string {
	defer qt.Recovering("QFileInfo::absolutePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_AbsolutePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) BaseName() string {
	defer qt.Recovering("QFileInfo::baseName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) BundleName() string {
	defer qt.Recovering("QFileInfo::bundleName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_BundleName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Caching() bool {
	defer qt.Recovering("QFileInfo::caching")

	if ptr.Pointer() != nil {
		return C.QFileInfo_Caching(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) CanonicalFilePath() string {
	defer qt.Recovering("QFileInfo::canonicalFilePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalFilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CanonicalPath() string {
	defer qt.Recovering("QFileInfo::canonicalPath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CanonicalPath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CompleteBaseName() string {
	defer qt.Recovering("QFileInfo::completeBaseName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteBaseName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) CompleteSuffix() string {
	defer qt.Recovering("QFileInfo::completeSuffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_CompleteSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Created() *QDateTime {
	defer qt.Recovering("QFileInfo::created")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_Created(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) Dir() *QDir {
	defer qt.Recovering("QFileInfo::dir")

	if ptr.Pointer() != nil {
		return NewQDirFromPointer(C.QFileInfo_Dir(ptr.Pointer()))
	}
	return nil
}

func QFileInfo_Exists2(file string) bool {
	defer qt.Recovering("QFileInfo::exists")

	return C.QFileInfo_QFileInfo_Exists2(C.CString(file)) != 0
}

func (ptr *QFileInfo) Exists() bool {
	defer qt.Recovering("QFileInfo::exists")

	if ptr.Pointer() != nil {
		return C.QFileInfo_Exists(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) FileName() string {
	defer qt.Recovering("QFileInfo::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) FilePath() string {
	defer qt.Recovering("QFileInfo::filePath")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_FilePath(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Group() string {
	defer qt.Recovering("QFileInfo::group")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Group(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) IsAbsolute() bool {
	defer qt.Recovering("QFileInfo::isAbsolute")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsBundle() bool {
	defer qt.Recovering("QFileInfo::isBundle")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsBundle(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsDir() bool {
	defer qt.Recovering("QFileInfo::isDir")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsDir(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsExecutable() bool {
	defer qt.Recovering("QFileInfo::isExecutable")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsExecutable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsFile() bool {
	defer qt.Recovering("QFileInfo::isFile")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsFile(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsHidden() bool {
	defer qt.Recovering("QFileInfo::isHidden")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsHidden(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsNativePath() bool {
	defer qt.Recovering("QFileInfo::isNativePath")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsNativePath(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsReadable() bool {
	defer qt.Recovering("QFileInfo::isReadable")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRelative() bool {
	defer qt.Recovering("QFileInfo::isRelative")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRelative(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsRoot() bool {
	defer qt.Recovering("QFileInfo::isRoot")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsRoot(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsSymLink() bool {
	defer qt.Recovering("QFileInfo::isSymLink")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsSymLink(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) IsWritable() bool {
	defer qt.Recovering("QFileInfo::isWritable")

	if ptr.Pointer() != nil {
		return C.QFileInfo_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) LastModified() *QDateTime {
	defer qt.Recovering("QFileInfo::lastModified")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_LastModified(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) LastRead() *QDateTime {
	defer qt.Recovering("QFileInfo::lastRead")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QFileInfo_LastRead(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileInfo) MakeAbsolute() bool {
	defer qt.Recovering("QFileInfo::makeAbsolute")

	if ptr.Pointer() != nil {
		return C.QFileInfo_MakeAbsolute(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileInfo) Owner() string {
	defer qt.Recovering("QFileInfo::owner")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Owner(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Path() string {
	defer qt.Recovering("QFileInfo::path")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Path(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Refresh() {
	defer qt.Recovering("QFileInfo::refresh")

	if ptr.Pointer() != nil {
		C.QFileInfo_Refresh(ptr.Pointer())
	}
}

func (ptr *QFileInfo) SetCaching(enable bool) {
	defer qt.Recovering("QFileInfo::setCaching")

	if ptr.Pointer() != nil {
		C.QFileInfo_SetCaching(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QFileInfo) SetFile3(dir QDir_ITF, file string) {
	defer qt.Recovering("QFileInfo::setFile")

	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile3(ptr.Pointer(), PointerFromQDir(dir), C.CString(file))
	}
}

func (ptr *QFileInfo) SetFile2(file QFile_ITF) {
	defer qt.Recovering("QFileInfo::setFile")

	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile2(ptr.Pointer(), PointerFromQFile(file))
	}
}

func (ptr *QFileInfo) SetFile(file string) {
	defer qt.Recovering("QFileInfo::setFile")

	if ptr.Pointer() != nil {
		C.QFileInfo_SetFile(ptr.Pointer(), C.CString(file))
	}
}

func (ptr *QFileInfo) Size() int64 {
	defer qt.Recovering("QFileInfo::size")

	if ptr.Pointer() != nil {
		return int64(C.QFileInfo_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileInfo) Suffix() string {
	defer qt.Recovering("QFileInfo::suffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_Suffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) Swap(other QFileInfo_ITF) {
	defer qt.Recovering("QFileInfo::swap")

	if ptr.Pointer() != nil {
		C.QFileInfo_Swap(ptr.Pointer(), PointerFromQFileInfo(other))
	}
}

func (ptr *QFileInfo) SymLinkTarget() string {
	defer qt.Recovering("QFileInfo::symLinkTarget")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileInfo_SymLinkTarget(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileInfo) DestroyQFileInfo() {
	defer qt.Recovering("QFileInfo::~QFileInfo")

	if ptr.Pointer() != nil {
		C.QFileInfo_DestroyQFileInfo(ptr.Pointer())
	}
}
