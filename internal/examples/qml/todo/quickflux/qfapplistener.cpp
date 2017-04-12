/* Project: Quick Flux
 * Author: Ben Lau
 * Website: https://www.github.com/benlau/quickflux
 *
 */
#include <QtQml>
#include <QtCore>
#include "qfappdispatcher.h"
#include "qfapplistener.h"

/*!
  \qmltype AppListener
  \inqmlmodule QuickFlux

  AppListener is a helper class to listen dispatched signal from AppDispatcher.
  It provides an alternative solution other then using Connections component.
  Moreover, it offers a way to control the order of signal delivery through waitFor property.

  Example:

\code

// Only listen for specific message
AppListener {
    filter: "messageType1";
    onDispatched: {
      // Your code here
    }
}

// Listen for multiple messages.
AppListener {
    onDispatched: {
        switch (type) {
            case "messageType1":
                // ...
                break;
            case "messageType2":
                // ...
                break;
        }
    }
}

// Alternative method to listen for multiple messages

AppListener {

  Component.onComponented: {
    on("messageType1",function() {
       /// Your code here.
    });
    on("messageType2",function() {
       /// Your code here.
    });
  }
}


\endcode

 */

/*!
  \qmlsignal AppListener::dispatched(string type, object message)

  It is a proxy of AppDispatcher.dispatched signal.
  If the enabled property is set to false, this signal will not be emitted.
 */

/*! \qmlproperty bool AppListener::enabled

  This property holds whether the listener receives message.
  If it is false, it won't emit "dispatched" signal and trigger callback registered via "on" function.
  By default this is true.

  The value can be controlled by parent component.
  Setting this property directly affects the enabled value of child items.
  When set to false, the enabled values of all child items also become false.
  When set to true,
  the enabled values of child items are returned to true,
  unless they have explicitly been set to false.
 */

QFAppListener::QFAppListener(QQuickItem *parent) : QQuickItem(parent)
{
    m_alwaysOn = false;
    m_listener = 0;
    m_listenerId = 0;
}

QFAppListener::~QFAppListener()
{
    if (!m_target.isNull()) {
        m_target->removeListener(m_listenerId);
    }
}

QObject *QFAppListener::target() const
{
    return m_target;
}

void QFAppListener::setTarget(QFDispatcher *target)
{
    if (!m_target.isNull()) {
        m_target->removeListener(m_listenerId);
        m_listener->disconnect(this);
        m_listener->deleteLater();
        m_listener = 0;
        setListenerId(0);
    }

    m_target = target;

    if (!m_target.isNull()) {

        m_listener = new QFListener(this);

        setListenerId(m_target->addListener(m_listener));

        setListenerWaitFor();

        connect(m_listener,SIGNAL(dispatched(QString,QJSValue)),
                this,SLOT(onMessageReceived(QString,QJSValue)));
    }
}

/*! \qmlmethod AppListener::on(string type,func callback)

    Add a listener to the end of the listeners array for the specified type of message.
    Multiple calls passing the same combination of event and listener will result in the listener being added multiple times.

 */

QFAppListener *QFAppListener::on(QString type, QJSValue callback)
{
    QList<QJSValue> list;

    if (mapping.contains(type)) {
        list = mapping[type];
    }

    list.append(callback);

    mapping[type] = list;

    return this;
}

/*! \qmlmethod AppListener::removeListener(string type, func callback)

  Remove a listener from the listener array for the specified message.
 */

void QFAppListener::removeListener(QString type, QJSValue callback)
{
    if (!mapping.contains(type)) {
        return;
    };

    QList<QJSValue> list;
    list = mapping[type];

    int index = -1;
    for (int i = 0 ; i < list.size() ;i++) {
        if (list.at(i).equals(callback)) {
            index = i;
            break;
        }
    }

    if (index >=0 ) {
        list.removeAt(index);
        mapping[type] = list;
    }
}

/*! \qmlmethod AppListener::removeAllListener(string type)
   Remove all the listeners for a message with type. If type is empty, it will remove all the listeners.
 */

void QFAppListener::removeAllListener(QString type)
{
    if (type.isEmpty()) {
        mapping.clear();
    } else {
        mapping.remove(type);
    }
}

void QFAppListener::componentComplete()
{
    QQuickItem::componentComplete();

    QQmlEngine *engine = qmlEngine(this);
    Q_ASSERT(engine);    

    QFAppDispatcher* dispatcher = QFAppDispatcher::instance(engine);
    if (!dispatcher) {
        qWarning() << "Unknown error: Unable to access AppDispatcher";
    } else {
        setTarget(dispatcher);
    }
}

void QFAppListener::onMessageReceived(QString type, QJSValue message)
{
    if (!isEnabled() && !m_alwaysOn)
        return;

    bool dispatch = true;

    QStringList rules = m_filters;
    if (!m_filter.isEmpty()) {
        rules.append(m_filter);
    }

    if (rules.size() > 0) {
        dispatch = false;

        for (int i = 0 ; i < rules.size() ; i++) {
            if (type == rules.at(i)) {
                dispatch = true;
                break;
            }
        }
    }

    if (dispatch) {
        emit dispatched(type,message);
    }

    // Listener registered with on() should not be affected by filter.

    if (!mapping.contains(type))
        return;

    QList<QJSValue> list = mapping[type];

    QList<QJSValue> arguments;
    arguments << message;

    Q_FOREACH(QJSValue value,list)  {
        if (value.isCallable()) {
            value.call(arguments);
        }
    }

}

void QFAppListener::setListenerWaitFor()
{
    if (!m_listener) {
        return;
    }

    m_listener->setWaitFor(m_waitFor);
}

/*! \qmlproperty array AppListener::waitFor

If it is set, it will block the emission of dispatched signal until all the specificed listeners has been invoked.

Example code:

\code
AppListener {
  id: listener1
}

AppListener {
   id: listener2
   waitFor: [listener1.listenerId]
}
\endcode

 */

QList<int> QFAppListener::waitFor() const
{
    return m_waitFor;
}

void QFAppListener::setWaitFor(const QList<int> &waitFor)
{
    m_waitFor = waitFor;
    setListenerWaitFor();
    emit waitForChanged();
}

/*! \qmlproperty int AppListener::listenerId
  The listener ID of this component.
  It could be used with AppListener::waitFor / AppDispatcher::waitFor to control the order of message delivery.
 */

int QFAppListener::listenerId() const
{
    return m_listenerId;
}

void QFAppListener::setListenerId(int listenerId)
{
    m_listenerId = listenerId;
    emit listenerIdChanged();
}

/*! \qmlproperty bool AppListener::alwaysOn

  This property holds a value to indicate if the listener should remain listening message when it is not enabled.

 */

bool QFAppListener::alwaysOn() const
{
    return m_alwaysOn;
}

void QFAppListener::setAlwaysOn(bool alwaysOn)
{
    m_alwaysOn = alwaysOn;
    emit alwaysOnChanged();
}

/*! \qmlproperty array AppListener::filters

  Set a list of filter to incoming messages. Only message with type matched by the filters will emit "dispatched" signal.
  If it is not set, it will dispatch every message.

 */

QStringList QFAppListener::filters() const
{
    return m_filters;
}

void QFAppListener::setFilters(const QStringList &filters)
{
    m_filters = filters;
    emit filtersChanged();
}

/*! \qmlproperty string AppListener::filter

  Set a filter to incoming messages. Only message with type matched with the filter will emit "dispatched" signal.
  If it is not set, it will dispatch every message.
 */

QString QFAppListener::filter() const
{
    return m_filter;
}

void QFAppListener::setFilter(const QString &filter)
{
    m_filter = filter;
    emit filterChanged();
}
