#include <QtCore>
#include <QMetaObject>
#include <QtQml>
#include "quickfluxfunctions.h"
#include "qffilter.h"

/*!
   \qmltype Filter
   \brief Add filter rule to AppListener

Filter component listens for the parent's dispatched signal, if a dispatched signal
match with its type, it will emit its own "dispatched" signal. Otherwise, it will
simply ignore the signal.

This component provides an alternative way to filter incoming message which is
suitable for making Store component.

Example:

\code

pragma Singleton
import QtQuick 2.0
import QuickFlux 1.0
import "../actions"

AppListener {
    id: store

    property ListModel model: ListModel { }

    Filter {
        type: ActionTypes.addTask
        onDispatched: {
            model.append({task: message.task});
        }
    }

}

\endcode

It is not suggested to use nested AppListener in a Store component. Because nested AppListener
do not share the same AppListener::listenerId, it will be difficult to control the order
of message reception between store component.

In contrast, Filter share the same listenerId with its parent, and therefore it is a solution
for above problem.

*/

/*!
  \qmlsignal Filter::dispatched(string type, object message)

  It is a proxy of parent's dispatched signal. If the parent emits a signal matched with the Filter::type / Filter::types property,
  it will emit this signal
 */


QFFilter::QFFilter(QObject *parent) : QObject(parent)
{

}

/*! \qmlproperty string Filter::type

  These types determine the filtering rule for incoming message. Only type matched will emit the "dispatched" signal.

\code
  AppListener {
    Filter {
      type: "action1"
      onDispatched: {
        // handle the action
      }
    }
  }
\endcode

\sa Filter::types
 */


QString QFFilter::type() const
{
    if (m_types.size() == 0) {
        return "";
    } else {
        return m_types[0];
    }
}

void QFFilter::setType(const QString &type)
{
    m_types = QStringList() << type;
    emit typeChanged();
    emit typesChanged();
}

void QFFilter::classBegin()
{

}

void QFFilter::componentComplete()
{
    QObject* object = parent();
    m_engine = qmlEngine(this);

    if (!object) {
        qDebug() << "Filter - Disabled due to missing parent.";
        return;
    }

    const QMetaObject* meta = object->metaObject();

    if (meta->indexOfSignal("dispatched(QString,QJSValue)") >= 0) {
        connect(object,SIGNAL(dispatched(QString,QJSValue)),
                this,SLOT(filter(QString,QJSValue)));
    } else if (meta->indexOfSignal("dispatched(QString,QVariant)") >= 0) {
        connect(object,SIGNAL(dispatched(QString,QVariant)),
                this,SLOT(filter(QString,QVariant)));
    } else {
        qDebug() << "Filter - Disabled due to missing dispatched signal in parent object.";
        return;
    }

}

void QFFilter::filter(QString type, QJSValue message)
{
    if (m_types.indexOf(type) >= 0) {
        QF_PRECHECK_DISPATCH(m_engine.data(), type, message);
        emit dispatched(type, message);
    }
}

void QFFilter::filter(QString type, QVariant message)
{
    if (m_types.indexOf(type) >= 0) {
        QJSValue value = message.value<QJSValue>();
        QF_PRECHECK_DISPATCH(m_engine.data(), type, value);

        emit dispatched(type, value);
    }
}

/*! \qmlproperty array Filter::types

  These types determine the filtering rule for incoming message. Only type matched will emit the "dispatched" signal.

\code
  AppListener {
    Filter {
      types: ["action1", "action2"]
      onDispatched: {
        // handle the action
      }
    }
  }
\endcode

\sa Filter::type

 */


QStringList QFFilter::types() const
{
    return m_types;
}

void QFFilter::setTypes(const QStringList &types)
{
    m_types = types;
}

QQmlListProperty<QObject> QFFilter::children()
{
    return QQmlListProperty<QObject>(qobject_cast<QObject*>(this),
                                     m_children);
}
