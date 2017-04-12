#include <QtCore>
#include <QMetaMethod>
#include <QtQml>
#include "qfactioncreator.h"
#include "qfsignalproxy.h"
#include "qfdispatcher.h"

/*!
  \qmltype ActionCreator
  \brief Create message from signal then dispatch via AppDispatcher

ActionCreator is a component that listens on its own signals, convert to message then dispatch via AppDispatcher. The message type will be same as the signal name. There has no limitation on number of arguments and their data type.

For example, you may declare an ActionCreator based component as:

\code
import QtQuick 2.0
import QuickFlux 1.0
pragma singleton

ActionCreator {
   signal open(string url);
}
\endcode

It is equivalent to:

\code
import QtQuick 2.0
import QuickFlux 1.0
pragma singleton

Item {
   function open(url) {
     AppDispatcher.dispatch(“open”, {url: url});
   }
}
\endcode

 */

QFActionCreator::QFActionCreator(QObject *parent) : QObject(parent)
{

}

QString QFActionCreator::genKeyTable()
{
    QStringList imports, header, footer, properties;

    imports << "pragma Singleton"
            << "import QtQuick 2.0"
            << "import QuickFlux 1.0\n";

    header << "KeyTable {\n";

    footer <<  "}";

    const int memberOffset = QObject::staticMetaObject.methodCount();

    const QMetaObject* meta = metaObject();

    int count = meta->methodCount();

    for (int i = memberOffset ; i < count ;i++) {
        QMetaMethod method = meta->method(i);
        if (method.name() == "dispatcherChanged") {
            continue;
        }
        if (method.methodType() == QMetaMethod::Signal) {
            properties << QString("    property string %1;\n").arg(QString(method.name()));
        }
    }

    QStringList content;
    content << imports << header << properties << footer;

    return content.join("\n");
}

void QFActionCreator::dispatch(QString type, QJSValue message)
{
    if (!m_dispatcher.isNull()) {
        m_dispatcher->dispatch(type, message);
    }
}

void QFActionCreator::classBegin()
{

}

void QFActionCreator::componentComplete()
{
    QQmlEngine* engine = qmlEngine(this);

    if (m_dispatcher.isNull()) {
        setDispatcher(qobject_cast<QFDispatcher*>(QFAppDispatcher::instance(engine)));
    }

    QFDispatcher* dispatcher = m_dispatcher.data();

    const int memberOffset = QObject::staticMetaObject.methodCount();

    const QMetaObject* meta = metaObject();

    int count = meta->methodCount();

    for (int i = memberOffset ; i < count ;i++) {
        QMetaMethod method = meta->method(i);
        if (method.name() == "dispatcherChanged") {
            continue;
        }

        if (method.methodType() == QMetaMethod::Signal) {
            QFSignalProxy* proxy = new QFSignalProxy(this);
            proxy->bind(this, i, engine, dispatcher);
            m_proxyList << proxy;
        }
    }
}

QFDispatcher *QFActionCreator::dispatcher() const
{
    return m_dispatcher;
}

void QFActionCreator::setDispatcher(QFDispatcher *value)
{
    m_dispatcher = value;
    for (int i = 0 ; i < m_proxyList.size();i++) {
        m_proxyList[i]->setDispatcher(m_dispatcher);
    }

    emit dispatcherChanged();
}
