#include "qshowevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QShowEvent>
#include "_cgo_export.h"

class MyQShowEvent: public QShowEvent {
public:
};

QtObjectPtr QShowEvent_NewQShowEvent(){
	return new QShowEvent();
}

