#include "qdynamicpropertychangeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QDynamicPropertyChangeEvent>
#include "_cgo_export.h"

class MyQDynamicPropertyChangeEvent: public QDynamicPropertyChangeEvent {
public:
};

void* QDynamicPropertyChangeEvent_NewQDynamicPropertyChangeEvent(void* name){
	return new QDynamicPropertyChangeEvent(*static_cast<QByteArray*>(name));
}

void* QDynamicPropertyChangeEvent_PropertyName(void* ptr){
	return new QByteArray(static_cast<QDynamicPropertyChangeEvent*>(ptr)->propertyName());
}

