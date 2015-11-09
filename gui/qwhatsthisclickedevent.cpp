#include "qwhatsthisclickedevent.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWhatsThisClickedEvent>
#include "_cgo_export.h"

class MyQWhatsThisClickedEvent: public QWhatsThisClickedEvent {
public:
};

void* QWhatsThisClickedEvent_NewQWhatsThisClickedEvent(char* href){
	return new QWhatsThisClickedEvent(QString(href));
}

char* QWhatsThisClickedEvent_Href(void* ptr){
	return static_cast<QWhatsThisClickedEvent*>(ptr)->href().toUtf8().data();
}

