#ifndef QUICKFLUXFUNCTIONS_H
#define QUICKFLUXFUNCTIONS_H

#include <QJSValue>

namespace QuickFlux {

    void printException(QJSValue value);

}

#if (QT_VERSION == QT_VERSION_CHECK(5,7,1)) || (QT_VERSION == QT_VERSION_CHECK(5,8,0) || defined(SAILFISH_OS))

/* A dirty hack to fix QTBUG-58133 and #13 issue.

   Reference:
   1. https://bugreports.qt.io/browse/QTBUG-58133
   2. https://github.com/benlau/quickflux/issues/13
 */
#define QF_PRECHECK_DISPATCH(engine, type, message)  {\
    if (message.isUndefined() && engine) { \
        message = engine->toScriptValue<QVariant>(QVariant()); \
    }\
}

#else
#define QF_PRECHECK_DISPATCH(engine, type, message)
#endif


#endif // QUICKFLUXFUNCTIONS_H
