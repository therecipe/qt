#include "mysignalhandler.h"
#include <QDebug>
#include "_cgo_export.h"

MySignalHandler::MySignalHandler(QObject *parent) : QObject(parent)
{

}

void MySignalHandler::callbackFromQml(QString value) {
    qDebug()<<"print from c++:"<<value;
	callbackTestSignal(value.toUtf8().data());
}
