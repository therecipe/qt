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

QtObjectPtr QMediaBindableInterface_MediaObject(QtObjectPtr ptr){
	return static_cast<QMediaBindableInterface*>(ptr)->mediaObject();
}

void QMediaBindableInterface_DestroyQMediaBindableInterface(QtObjectPtr ptr){
	static_cast<QMediaBindableInterface*>(ptr)->~QMediaBindableInterface();
}

