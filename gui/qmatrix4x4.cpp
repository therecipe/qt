#include "qmatrix4x4.h"
#include <QString>
#include <QRect>
#include <QVector>
#include <QVector4D>
#include <QQuaternion>
#include <QRectF>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVector3D>
#include <QTransform>
#include <QMatrix4x4>
#include "_cgo_export.h"

class MyQMatrix4x4: public QMatrix4x4 {
public:
};

void* QMatrix4x4_NewQMatrix4x4(){
	return new QMatrix4x4();
}

void* QMatrix4x4_NewQMatrix4x47(void* transform){
	return new QMatrix4x4(*static_cast<QTransform*>(transform));
}

int QMatrix4x4_IsAffine(void* ptr){
	return static_cast<QMatrix4x4*>(ptr)->isAffine();
}

int QMatrix4x4_IsIdentity(void* ptr){
	return static_cast<QMatrix4x4*>(ptr)->isIdentity();
}

void QMatrix4x4_LookAt(void* ptr, void* eye, void* center, void* up){
	static_cast<QMatrix4x4*>(ptr)->lookAt(*static_cast<QVector3D*>(eye), *static_cast<QVector3D*>(center), *static_cast<QVector3D*>(up));
}

void QMatrix4x4_Optimize(void* ptr){
	static_cast<QMatrix4x4*>(ptr)->optimize();
}

void QMatrix4x4_Ortho2(void* ptr, void* rect){
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRect*>(rect));
}

void QMatrix4x4_Ortho3(void* ptr, void* rect){
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRectF*>(rect));
}

void QMatrix4x4_Rotate2(void* ptr, void* quaternion){
	static_cast<QMatrix4x4*>(ptr)->rotate(*static_cast<QQuaternion*>(quaternion));
}

void QMatrix4x4_Scale(void* ptr, void* vector){
	static_cast<QMatrix4x4*>(ptr)->scale(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_SetColumn(void* ptr, int index, void* value){
	static_cast<QMatrix4x4*>(ptr)->setColumn(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetRow(void* ptr, int index, void* value){
	static_cast<QMatrix4x4*>(ptr)->setRow(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetToIdentity(void* ptr){
	static_cast<QMatrix4x4*>(ptr)->setToIdentity();
}

void QMatrix4x4_Translate(void* ptr, void* vector){
	static_cast<QMatrix4x4*>(ptr)->translate(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Viewport2(void* ptr, void* rect){
	static_cast<QMatrix4x4*>(ptr)->viewport(*static_cast<QRectF*>(rect));
}

