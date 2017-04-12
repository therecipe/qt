#include <QQmlEngine>
#include "qfmiddleware.h"
#include "quickfluxfunctions.h"


/*!
   \qmltype Middleware
   \inqmlmodule QuickFlux 1.1

The middleware in Quick Flux is similar to the one in Redux and those from server libraries like Express and Koa. It is some code you can put between the Dispatcher and Stores. It could modify/defer/remove received actions and insert new actions that will dispatch to Store components. Users may use it for logging, crash reporting, talking to asynchronous components like FileDialog / Camera etc. So that Store components remain “pure”, it holds application logic only and always return the same result for the same input. It is easier to write test cases.

Example Code

\code
\endcode

    */

QFMiddleware::QFMiddleware(QQuickItem* parent) : QQuickItem(parent),  m_filterFunctionEnabled(false)
{

}

/*! \qmlmethod Middleware::next(string type, object message)

  Pass an action message to next middleware. If it is already the last middleware, the action will be dispatched to Store component.
  */

/*! \qmlproperty bool Store::filterFunctionEnabled
 */

void QFMiddleware::next(QString type, QJSValue message)
{
    QQmlEngine* engine = qmlEngine(this);
    QF_PRECHECK_DISPATCH(engine, type, message);

    if (m_nextCallback.isCallable()) {
        QJSValueList args;
        args << type;
        args << message;
        m_nextCallback.call(args);
    }
}

QJSValue QFMiddleware::nextCallback() const
{
    return m_nextCallback;
}

void QFMiddleware::setNextCallback(const QJSValue &value)
{
    m_nextCallback = value;
    emit _nextCallbackChanged();
}
