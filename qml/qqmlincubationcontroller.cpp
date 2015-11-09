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

void* QQmlIncubationController_NewQQmlIncubationController(){
	return new QQmlIncubationController();
}

void* QQmlIncubationController_Engine(void* ptr){
	return static_cast<QQmlIncubationController*>(ptr)->engine();
}

void QQmlIncubationController_IncubateFor(void* ptr, int msecs){
	static_cast<QQmlIncubationController*>(ptr)->incubateFor(msecs);
}

int QQmlIncubationController_IncubatingObjectCount(void* ptr){
	return static_cast<QQmlIncubationController*>(ptr)->incubatingObjectCount();
}

