package gui

//#include "qopenglshaderprogram.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QOpenGLShaderProgram struct {
	core.QObject
}

type QOpenGLShaderProgramITF interface {
	core.QObjectITF
	QOpenGLShaderProgramPTR() *QOpenGLShaderProgram
}

func PointerFromQOpenGLShaderProgram(ptr QOpenGLShaderProgramITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLShaderProgramPTR().Pointer()
	}
	return nil
}

func QOpenGLShaderProgramFromPointer(ptr unsafe.Pointer) *QOpenGLShaderProgram {
	var n = new(QOpenGLShaderProgram)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QOpenGLShaderProgram_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QOpenGLShaderProgram) QOpenGLShaderProgramPTR() *QOpenGLShaderProgram {
	return ptr
}

func NewQOpenGLShaderProgram(parent core.QObjectITF) *QOpenGLShaderProgram {
	return QOpenGLShaderProgramFromPointer(unsafe.Pointer(C.QOpenGLShaderProgram_NewQOpenGLShaderProgram(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QOpenGLShaderProgram) AddShader(shader QOpenGLShaderITF) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_AddShader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLShader(shader))) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) AddShaderFromSourceCode2(ty QOpenGLShader__ShaderTypeBit, source core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_AddShaderFromSourceCode2(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.QtObjectPtr(core.PointerFromQByteArray(source))) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) AddShaderFromSourceCode3(ty QOpenGLShader__ShaderTypeBit, source string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_AddShaderFromSourceCode3(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(source)) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) AddShaderFromSourceCode(ty QOpenGLShader__ShaderTypeBit, source string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_AddShaderFromSourceCode(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(source)) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) AddShaderFromSourceFile(ty QOpenGLShader__ShaderTypeBit, fileName string) bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_AddShaderFromSourceFile(C.QtObjectPtr(ptr.Pointer()), C.int(ty), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) AttributeLocation2(name core.QByteArrayITF) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_AttributeLocation2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(name))))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) AttributeLocation3(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_AttributeLocation3(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) AttributeLocation(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_AttributeLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) Bind() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_Bind(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) BindAttributeLocation2(name core.QByteArrayITF, location int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_BindAttributeLocation2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(name)), C.int(location))
	}
}

func (ptr *QOpenGLShaderProgram) BindAttributeLocation3(name string, location int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_BindAttributeLocation3(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(location))
	}
}

func (ptr *QOpenGLShaderProgram) BindAttributeLocation(name string, location int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_BindAttributeLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.int(location))
	}
}

func (ptr *QOpenGLShaderProgram) Create() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_Create(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) DisableAttributeArray2(name string) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_DisableAttributeArray2(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QOpenGLShaderProgram) DisableAttributeArray(location int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_DisableAttributeArray(C.QtObjectPtr(ptr.Pointer()), C.int(location))
	}
}

func (ptr *QOpenGLShaderProgram) EnableAttributeArray2(name string) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_EnableAttributeArray2(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QOpenGLShaderProgram) EnableAttributeArray(location int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_EnableAttributeArray(C.QtObjectPtr(ptr.Pointer()), C.int(location))
	}
}

func QOpenGLShaderProgram_HasOpenGLShaderPrograms(context QOpenGLContextITF) bool {
	return C.QOpenGLShaderProgram_QOpenGLShaderProgram_HasOpenGLShaderPrograms(C.QtObjectPtr(PointerFromQOpenGLContext(context))) != 0
}

func (ptr *QOpenGLShaderProgram) IsLinked() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_IsLinked(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) Link() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLShaderProgram_Link(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLShaderProgram) Log() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QOpenGLShaderProgram_Log(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QOpenGLShaderProgram) MaxGeometryOutputVertices() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_MaxGeometryOutputVertices(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) PatchVertexCount() int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_PatchVertexCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) Release() {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_Release(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLShaderProgram) RemoveAllShaders() {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_RemoveAllShaders(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLShaderProgram) RemoveShader(shader QOpenGLShaderITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_RemoveShader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLShader(shader)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray7(name string, values QVector2DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray7(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector2D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray8(name string, values QVector3DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray8(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector3D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray9(name string, values QVector4DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray9(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector4D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray2(location int, values QVector2DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray2(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector2D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray3(location int, values QVector3DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray3(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector3D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeArray4(location int, values QVector4DITF, stride int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeArray4(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector4D(values)), C.int(stride))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue17(name string, value QColorITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue17(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQColor(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue14(name string, value QVector2DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue14(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector2D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue15(name string, value QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue15(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector3D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue16(name string, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue16(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue8(location int, value QColorITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue8(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQColor(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue5(location int, value QVector2DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue5(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector2D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue6(location int, value QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue6(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector3D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetAttributeValue7(location int, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetAttributeValue7(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetPatchVertexCount(count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetPatchVertexCount(C.QtObjectPtr(ptr.Pointer()), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue34(name string, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue34(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue47(name string, value QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue47(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQMatrix4x4(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue35(name string, point core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue35(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(core.PointerFromQPoint(point)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue36(name string, point core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue36(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(core.PointerFromQPointF(point)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue37(name string, size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue37(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue38(name string, size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue38(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue54(name string, value QTransformITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue54(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQTransform(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue31(name string, value QVector2DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue31(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector2D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue32(name string, value QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue32(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector3D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue33(name string, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue33(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue10(location int, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue10(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue23(location int, value QMatrix4x4ITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue23(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQMatrix4x4(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue11(location int, point core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue11(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(core.PointerFromQPoint(point)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue12(location int, point core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue12(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(core.PointerFromQPointF(point)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue13(location int, size core.QSizeITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue13(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(core.PointerFromQSize(size)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue14(location int, size core.QSizeFITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue14(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(core.PointerFromQSizeF(size)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue24(location int, value QTransformITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue24(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQTransform(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue7(location int, value QVector2DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue7(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector2D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue8(location int, value QVector3DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue8(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector3D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValue9(location int, value QVector4DITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValue9(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector4D(value)))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray30(name string, values QMatrix4x4ITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray30(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQMatrix4x4(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray19(name string, values QVector2DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray19(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector2D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray20(name string, values QVector3DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray20(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector3D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray21(name string, values QVector4DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray21(C.QtObjectPtr(ptr.Pointer()), C.CString(name), C.QtObjectPtr(PointerFromQVector4D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray15(location int, values QMatrix4x4ITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray15(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQMatrix4x4(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray4(location int, values QVector2DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray4(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector2D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray5(location int, values QVector3DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray5(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector3D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) SetUniformValueArray6(location int, values QVector4DITF, count int) {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_SetUniformValueArray6(C.QtObjectPtr(ptr.Pointer()), C.int(location), C.QtObjectPtr(PointerFromQVector4D(values)), C.int(count))
	}
}

func (ptr *QOpenGLShaderProgram) UniformLocation2(name core.QByteArrayITF) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_UniformLocation2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(name))))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) UniformLocation3(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_UniformLocation3(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) UniformLocation(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QOpenGLShaderProgram_UniformLocation(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QOpenGLShaderProgram) DestroyQOpenGLShaderProgram() {
	if ptr.Pointer() != nil {
		C.QOpenGLShaderProgram_DestroyQOpenGLShaderProgram(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
