#include "qqmlincubationcontroller.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlIncubationController>
#include "_cgo_export.h"

class MyQQmlIncubationController: public QQmlIncubationController {
public:
};

QtObjectPtr QQmlIncubationController_NewQQmlIncubationController(){
	return new QQmlIncubationController();
}

QtObjectPtr QQmlIncubationController_Engine(QtObjectPtr ptr){
	return static_cast<QQmlIncubationController*>(ptr)->engine();
}

void QQmlIncubationController_IncubateFor(QtObjectPtr ptr, int msecs){
	static_cast<QQmlIncubationController*>(ptr)->incubateFor(msecs);
}

int QQmlIncubationController_IncubatingObjectCount(QtObjectPtr ptr){
	return static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCount();
}

