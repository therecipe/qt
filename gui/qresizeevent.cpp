#include "qresizeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QResizeEvent>
#include "_cgo_export.h"

class MyQResizeEvent: public QResizeEvent {
public:
};

QtObjectPtr QResizeEvent_NewQResizeEvent(QtObjectPtr size, QtObjectPtr oldSize){
	return new QResizeEvent(*static_cast<QSize*>(size), *static_cast<QSize*>(oldSize));
}

