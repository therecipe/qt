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

void* QCloseEvent_NewQCloseEvent(){
	return new QCloseEvent();
}

