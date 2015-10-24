#include "qexposeevent.h"
#include <QRegion>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QExposeEvent>
#include "_cgo_export.h"

class MyQExposeEvent: public QExposeEvent {
public:
};

QtObjectPtr QExposeEvent_NewQExposeEvent(QtObjectPtr exposeRegion){
	return new QExposeEvent(*static_cast<QRegion*>(exposeRegion));
}

