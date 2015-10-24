#include "qstatustipevent.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStatusTipEvent>
#include "_cgo_export.h"

class MyQStatusTipEvent: public QStatusTipEvent {
public:
};

QtObjectPtr QStatusTipEvent_NewQStatusTipEvent(char* tip){
	return new QStatusTipEvent(QString(tip));
}

char* QStatusTipEvent_Tip(QtObjectPtr ptr){
	return static_cast<QStatusTipEvent*>(ptr)->tip().toUtf8().data();
}

