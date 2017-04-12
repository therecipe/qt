#include <QtCore>
#include "quickfluxfunctions.h"

void QuickFlux::printException(QJSValue value)
{
    if (value.isError()) {
        QString message = QString("%1:%2: %3: %4")
                          .arg(value.property("fileName").toString())
                          .arg(value.property("lineNumber").toString())
                          .arg(value.property("name").toString())
                          .arg(value.property("message").toString());
        qWarning() << message;
    }
}
