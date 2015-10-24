#include "qcloseevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QCloseEvent>
#include "_cgo_export.h"

class MyQCloseEvent: public QCloseEvent {
public:
};

QtObjectPtr QCloseEvent_NewQCloseEvent(){
	return new QCloseEvent();
}

