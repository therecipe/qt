#include "qwhatsthisclickedevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWhatsThisClickedEvent>
#include "_cgo_export.h"

class MyQWhatsThisClickedEvent: public QWhatsThisClickedEvent {
public:
};

QtObjectPtr QWhatsThisClickedEvent_NewQWhatsThisClickedEvent(char* href){
	return new QWhatsThisClickedEvent(QString(href));
}

char* QWhatsThisClickedEvent_Href(QtObjectPtr ptr){
	return static_cast<QWhatsThisClickedEvent*>(ptr)->href().toUtf8().data();
}

