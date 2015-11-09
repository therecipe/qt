#include "qexposeevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRegion>
#include <QExposeEvent>
#include "_cgo_export.h"

class MyQExposeEvent: public QExposeEvent {
public:
};

void* QExposeEvent_NewQExposeEvent(void* exposeRegion){
	return new QExposeEvent(*static_cast<QRegion*>(exposeRegion));
}

void* QExposeEvent_Region(void* ptr){
	return new QRegion(static_cast<QExposeEvent*>(ptr)->region());
}

