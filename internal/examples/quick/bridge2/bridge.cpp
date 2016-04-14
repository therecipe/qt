#include "bridge.h"
#include "mysignalhandler.h"
#include <QQmlContext>

void initSignalHandler(void* rootContext) {
    QQmlContext* ptr = static_cast<QQmlContext*>(rootContext);
    ptr->setContextProperty("SignalHandler",new MySignalHandler());
}


