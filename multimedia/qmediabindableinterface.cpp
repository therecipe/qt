#include "qmediabindableinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaBindableInterface>
#include "_cgo_export.h"

class MyQMediaBindableInterface: public QMediaBindableInterface {
public:
};

void* QMediaBindableInterface_MediaObject(void* ptr){
	return static_cast<QMediaBindableInterface*>(ptr)->mediaObject();
}

void QMediaBindableInterface_DestroyQMediaBindableInterface(void* ptr){
	static_cast<QMediaBindableInterface*>(ptr)->~QMediaBindableInterface();
}

