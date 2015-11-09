#include "qgraphicsobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGraphicsObject>
#include "_cgo_export.h"

class MyQGraphicsObject: public QGraphicsObject {
public:
void Signal_EnabledChanged(){callbackQGraphicsObjectEnabledChanged(this->objectName().toUtf8().data());};
void Signal_OpacityChanged(){callbackQGraphicsObjectOpacityChanged(this->objectName().toUtf8().data());};
void Signal_ParentChanged(){callbackQGraphicsObjectParentChanged(this->objectName().toUtf8().data());};
void Signal_RotationChanged(){callbackQGraphicsObjectRotationChanged(this->objectName().toUtf8().data());};
void Signal_ScaleChanged(){callbackQGraphicsObjectScaleChanged(this->objectName().toUtf8().data());};
void Signal_VisibleChanged(){callbackQGraphicsObjectVisibleChanged(this->objectName().toUtf8().data());};
void Signal_XChanged(){callbackQGraphicsObjectXChanged(this->objectName().toUtf8().data());};
void Signal_YChanged(){callbackQGraphicsObjectYChanged(this->objectName().toUtf8().data());};
void Signal_ZChanged(){callbackQGraphicsObjectZChanged(this->objectName().toUtf8().data());};
};

void QGraphicsObject_ConnectEnabledChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::enabledChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_EnabledChanged));;
}

void QGraphicsObject_DisconnectEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::enabledChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_EnabledChanged));;
}

void QGraphicsObject_GrabGesture(void* ptr, int gesture, int flags){
	static_cast<QGraphicsObject*>(ptr)->grabGesture(static_cast<Qt::GestureType>(gesture), static_cast<Qt::GestureFlag>(flags));
}

void QGraphicsObject_ConnectOpacityChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::opacityChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_OpacityChanged));;
}

void QGraphicsObject_DisconnectOpacityChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::opacityChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_OpacityChanged));;
}

void QGraphicsObject_ConnectParentChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::parentChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ParentChanged));;
}

void QGraphicsObject_DisconnectParentChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::parentChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ParentChanged));;
}

void QGraphicsObject_ConnectRotationChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::rotationChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_RotationChanged));;
}

void QGraphicsObject_DisconnectRotationChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::rotationChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_RotationChanged));;
}

void QGraphicsObject_ConnectScaleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::scaleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ScaleChanged));;
}

void QGraphicsObject_DisconnectScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::scaleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ScaleChanged));;
}

void QGraphicsObject_UngrabGesture(void* ptr, int gesture){
	static_cast<QGraphicsObject*>(ptr)->ungrabGesture(static_cast<Qt::GestureType>(gesture));
}

void QGraphicsObject_ConnectVisibleChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::visibleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_VisibleChanged));;
}

void QGraphicsObject_DisconnectVisibleChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::visibleChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_VisibleChanged));;
}

void QGraphicsObject_ConnectXChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::xChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_XChanged));;
}

void QGraphicsObject_DisconnectXChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::xChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_XChanged));;
}

void QGraphicsObject_ConnectYChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::yChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_YChanged));;
}

void QGraphicsObject_DisconnectYChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::yChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_YChanged));;
}

void QGraphicsObject_ConnectZChanged(void* ptr){
	QObject::connect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::zChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ZChanged));;
}

void QGraphicsObject_DisconnectZChanged(void* ptr){
	QObject::disconnect(static_cast<QGraphicsObject*>(ptr), static_cast<void (QGraphicsObject::*)()>(&QGraphicsObject::zChanged), static_cast<MyQGraphicsObject*>(ptr), static_cast<void (MyQGraphicsObject::*)()>(&MyQGraphicsObject::Signal_ZChanged));;
}

void QGraphicsObject_DestroyQGraphicsObject(void* ptr){
	static_cast<QGraphicsObject*>(ptr)->~QGraphicsObject();
}

