#include "qstatustipevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStatusTipEvent>
#include "_cgo_export.h"

class MyQStatusTipEvent: public QStatusTipEvent {
public:
};

void* QStatusTipEvent_NewQStatusTipEvent(char* tip){
	return new QStatusTipEvent(QString(tip));
}

char* QStatusTipEvent_Tip(void* ptr){
	return static_cast<QStatusTipEvent*>(ptr)->tip().toUtf8().data();
}

