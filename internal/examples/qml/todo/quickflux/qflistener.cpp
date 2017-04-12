#include <QtCore>
#include "qfdispatcher.h"
#include "qflistener.h"

QFListener::QFListener(QObject *parent) : QObject(parent)
{
    m_listenerId = 0;
}

QFListener::~QFListener()
{
}

QJSValue QFListener::callback() const
{
    return m_callback;
}

void QFListener::setCallback(const QJSValue &callback)
{
    m_callback = callback;
}

void QFListener::dispatch(QFDispatcher *dispatcher,QString type, QJSValue message)
{

    if (m_waitFor.size() > 0) {
        dispatcher->waitFor(m_waitFor);
    }

    if (m_callback.isCallable()) {
        QJSValueList args;
        args << type << message;
        QJSValue ret = m_callback.call(args);

        if (ret.isError()) {
            QString message = QString("%1:%2: %3: %4")
                              .arg(ret.property("fileName").toString())
                              .arg(ret.property("lineNumber").toString())
                              .arg(ret.property("name").toString())
                              .arg(ret.property("message").toString());
            qWarning() << message;
        }
    }

    emit dispatched(type,message);
}

int QFListener::listenerId() const
{
    return m_listenerId;
}

void QFListener::setListenerId(int listenerId)
{
    m_listenerId = listenerId;
}

QList<int> QFListener::waitFor() const
{
    return m_waitFor;
}

void QFListener::setWaitFor(const QList<int> &waitFor)
{
    m_waitFor = waitFor;
}

