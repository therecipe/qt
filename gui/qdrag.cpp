#include "qdrag.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QPixmap>
#include <QPoint>
#include <QObject>
#include <QDrag>
#include "_cgo_export.h"

class MyQDrag: public QDrag {
public:
void Signal_ActionChanged(Qt::DropAction action){callbackQDragActionChanged(this->objectName().toUtf8().data(), action);};
void Signal_TargetChanged(QObject * newTarget){callbackQDragTargetChanged(this->objectName().toUtf8().data(), newTarget);};
};

void* QDrag_NewQDrag(void* dragSource){
	return new QDrag(static_cast<QObject*>(dragSource));
}

void QDrag_ConnectActionChanged(void* ptr){
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));;
}

void QDrag_DisconnectActionChanged(void* ptr){
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));;
}

int QDrag_Exec(void* ptr, int supportedActions){
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions));
}

int QDrag_Exec2(void* ptr, int supportedActions, int defaultDropAction){
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions), static_cast<Qt::DropAction>(defaultDropAction));
}

void* QDrag_MimeData(void* ptr){
	return static_cast<QDrag*>(ptr)->mimeData();
}

void QDrag_SetDragCursor(void* ptr, void* cursor, int action){
	static_cast<QDrag*>(ptr)->setDragCursor(*static_cast<QPixmap*>(cursor), static_cast<Qt::DropAction>(action));
}

void QDrag_SetHotSpot(void* ptr, void* hotspot){
	static_cast<QDrag*>(ptr)->setHotSpot(*static_cast<QPoint*>(hotspot));
}

void QDrag_SetMimeData(void* ptr, void* data){
	static_cast<QDrag*>(ptr)->setMimeData(static_cast<QMimeData*>(data));
}

void QDrag_SetPixmap(void* ptr, void* pixmap){
	static_cast<QDrag*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QDrag_Source(void* ptr){
	return static_cast<QDrag*>(ptr)->source();
}

int QDrag_SupportedActions(void* ptr){
	return static_cast<QDrag*>(ptr)->supportedActions();
}

void* QDrag_Target(void* ptr){
	return static_cast<QDrag*>(ptr)->target();
}

void QDrag_ConnectTargetChanged(void* ptr){
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));;
}

void QDrag_DisconnectTargetChanged(void* ptr){
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));;
}

void QDrag_DestroyQDrag(void* ptr){
	static_cast<QDrag*>(ptr)->~QDrag();
}

