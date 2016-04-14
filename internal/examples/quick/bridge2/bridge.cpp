#include "bridge.h"
#include "mysignalhandler.h"
#include <QQmlContext>

MySignalHandler *signalHandler = new MySignalHandler();

void initSignalHandler(void* rootContext) {
    QQmlContext* ptr = static_cast<QQmlContext*>(rootContext);
    ptr->setContextProperty("SignalHandler", signalHandler);
}

void sendSignalToQml(char* value) {
    signalHandler->sendToQml(QString(value));
}