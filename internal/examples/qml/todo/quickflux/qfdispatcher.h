#pragma once

#include <QObject>
#include <QVariantMap>
#include <QJSValue>
#include <QQueue>
#include <QPair>
#include <QQmlEngine>
#include <QPointer>
#include "qflistener.h"
#include "qfhook.h"
/// Message Dispatcher

class QFDispatcher : public QObject
{
    Q_OBJECT
public:
    explicit QFDispatcher(QObject *parent = 0);
    ~QFDispatcher();

signals:
    /// Listeners should listen on this signal to get the latest dispatched message from AppDispatcher
    void dispatched(QString type,QJSValue message);

public slots:
    /// Dispatch a message via Dispatcher
    /**
      @param type The type of the message
      @param message The message content
      @reentrant

      Dispatch a message with type via the AppDispatcher. Listeners should listen on the "dispatched"
      signal to be notified.

      Usually, it will emit "dispatched" signal immediately after calling dispatch(). However, if
      AppDispatcher is still dispatching messages, the new messages will be placed on a queue, and
      wait until it is finished. It guarantees the order of messages are arrived in sequence to
      listeners
     */
    Q_INVOKABLE void dispatch(QString type,QJSValue message = QJSValue());

    Q_INVOKABLE void waitFor(QList<int> ids);

    Q_INVOKABLE int addListener(QJSValue callback);

    Q_INVOKABLE void removeListener(int id);

public:

    void dispatch(const QString& type, const QVariant& message);

    int addListener(QFListener* listener);

    QQmlEngine *engine() const;

    void setEngine(QQmlEngine *engine);

    QFHook *hook() const;

    void setHook(QFHook *hook);

private slots:
    /// Invoke listener and emit the dispatched signal
    void send(QString type,QJSValue message);

private:
    void invokeListeners(QList<int> ids);

    bool m_dispatching;

    QPointer<QQmlEngine> m_engine;

    // Queue for dispatching messages
    QQueue<QPair<QString,QJSValue > > m_queue;

    // Next id for listener.
    int nextListenerId;

    // Registered listener
    QMap<int, QPointer<QFListener> > m_listeners;

    // Current dispatching listener id
    int dispatchingListenerId;

    // Current dispatching message
    QJSValue dispatchingMessage;

    // Current dispatching message type
    QString dispatchingMessageType;

    // List of listeners pending to be invoked.
    QMap<int,bool> pendingListeners;

    // List of listeners blocked in waitFor()
    QMap<int,bool> waitingListeners;

    QPointer<QFHook> m_hook;
};

