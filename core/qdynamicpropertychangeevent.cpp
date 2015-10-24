#include "qdynamicpropertychangeevent.h"
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDynamicPropertyChangeEvent>
#include "_cgo_export.h"

class MyQDynamicPropertyChangeEvent: public QDynamicPropertyChangeEvent {
public:
};

QtObjectPtr QDynamicPropertyChangeEvent_NewQDynamicPropertyChangeEvent(QtObjectPtr name){
	return new QDynamicPropertyChangeEvent(*static_cast<QByteArray*>(name));
}

