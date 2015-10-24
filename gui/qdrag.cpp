#include "qdrag.h"
#include <QMimeData>
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPoint>
#include <QDrag>
#include "_cgo_export.h"

class MyQDrag: public QDrag {
public:
void Signal_ActionChanged(Qt::DropAction action){callbackQDragActionChanged(this->objectName().toUtf8().data(), action);};
void Signal_TargetChanged(QObject * newTarget){callbackQDragTargetChanged(this->objectName().toUtf8().data(), newTarget);};
};

QtObjectPtr QDrag_NewQDrag(QtObjectPtr dragSource){
	return new QDrag(static_cast<QObject*>(dragSource));
}

void QDrag_ConnectActionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));;
}

void QDrag_DisconnectActionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));;
}

int QDrag_Exec(QtObjectPtr ptr, int supportedActions){
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions));
}

int QDrag_Exec2(QtObjectPtr ptr, int supportedActions, int defaultDropAction){
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions), static_cast<Qt::DropAction>(defaultDropAction));
}

QtObjectPtr QDrag_MimeData(QtObjectPtr ptr){
	return static_cast<QDrag*>(ptr)->mimeData();
}

void QDrag_SetDragCursor(QtObjectPtr ptr, QtObjectPtr cursor, int action){
	static_cast<QDrag*>(ptr)->setDragCursor(*static_cast<QPixmap*>(cursor), static_cast<Qt::DropAction>(action));
}

void QDrag_SetHotSpot(QtObjectPtr ptr, QtObjectPtr hotspot){
	static_cast<QDrag*>(ptr)->setHotSpot(*static_cast<QPoint*>(hotspot));
}

void QDrag_SetMimeData(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QDrag*>(ptr)->setMimeData(static_cast<QMimeData*>(data));
}

void QDrag_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap){
	static_cast<QDrag*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

QtObjectPtr QDrag_Source(QtObjectPtr ptr){
	return static_cast<QDrag*>(ptr)->source();
}

int QDrag_SupportedActions(QtObjectPtr ptr){
	return static_cast<QDrag*>(ptr)->supportedActions();
}

QtObjectPtr QDrag_Target(QtObjectPtr ptr){
	return static_cast<QDrag*>(ptr)->target();
}

void QDrag_ConnectTargetChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));;
}

void QDrag_DisconnectTargetChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));;
}

void QDrag_DestroyQDrag(QtObjectPtr ptr){
	static_cast<QDrag*>(ptr)->~QDrag();
}

