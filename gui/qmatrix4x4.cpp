#include "qmatrix4x4.h"
#include <QVector4D>
#include <QUrl>
#include <QVector3D>
#include <QTransform>
#include <QQuaternion>
#include <QVector>
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMatrix4x4>
#include "_cgo_export.h"

class MyQMatrix4x4: public QMatrix4x4 {
public:
};

QtObjectPtr QMatrix4x4_NewQMatrix4x4(){
	return new QMatrix4x4();
}

QtObjectPtr QMatrix4x4_NewQMatrix4x47(QtObjectPtr transform){
	return new QMatrix4x4(*static_cast<QTransform*>(transform));
}

int QMatrix4x4_IsAffine(QtObjectPtr ptr){
	return static_cast<QMatrix4x4*>(ptr)->isAffine();
}

int QMatrix4x4_IsIdentity(QtObjectPtr ptr){
	return static_cast<QMatrix4x4*>(ptr)->isIdentity();
}

void QMatrix4x4_LookAt(QtObjectPtr ptr, QtObjectPtr eye, QtObjectPtr center, QtObjectPtr up){
	static_cast<QMatrix4x4*>(ptr)->lookAt(*static_cast<QVector3D*>(eye), *static_cast<QVector3D*>(center), *static_cast<QVector3D*>(up));
}

void QMatrix4x4_Optimize(QtObjectPtr ptr){
	static_cast<QMatrix4x4*>(ptr)->optimize();
}

void QMatrix4x4_Ortho2(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRect*>(rect));
}

void QMatrix4x4_Ortho3(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRectF*>(rect));
}

void QMatrix4x4_Rotate2(QtObjectPtr ptr, QtObjectPtr quaternion){
	static_cast<QMatrix4x4*>(ptr)->rotate(*static_cast<QQuaternion*>(quaternion));
}

void QMatrix4x4_Scale(QtObjectPtr ptr, QtObjectPtr vector){
	static_cast<QMatrix4x4*>(ptr)->scale(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_SetColumn(QtObjectPtr ptr, int index, QtObjectPtr value){
	static_cast<QMatrix4x4*>(ptr)->setColumn(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetRow(QtObjectPtr ptr, int index, QtObjectPtr value){
	static_cast<QMatrix4x4*>(ptr)->setRow(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetToIdentity(QtObjectPtr ptr){
	static_cast<QMatrix4x4*>(ptr)->setToIdentity();
}

void QMatrix4x4_Translate(QtObjectPtr ptr, QtObjectPtr vector){
	static_cast<QMatrix4x4*>(ptr)->translate(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Viewport2(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QMatrix4x4*>(ptr)->viewport(*static_cast<QRectF*>(rect));
}

