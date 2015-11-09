#include "qresizeevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QResizeEvent>
#include "_cgo_export.h"

class MyQResizeEvent: public QResizeEvent {
public:
};

void* QResizeEvent_NewQResizeEvent(void* size, void* oldSize){
	return new QResizeEvent(*static_cast<QSize*>(size), *static_cast<QSize*>(oldSize));
}

