#ifndef QFAPPLISTENER_H
#define QFAPPLISTENER_H

#include <QObject>
#include <QJSValue>
#include <QPointer>
#include <QQuickItem>
#include <QQmlParserStatus>
#include <QMap>
#include "qflistener.h"
#include "qfdispatcher.h"

class QFAppListener : public QQuickItem
{
    Q_OBJECT
    Q_PROPERTY(QString filter READ filter WRITE setFilter NOTIFY filterChanged)
    Q_PROPERTY(QStringList filters READ filters WRITE setFilters NOTIFY filtersChanged)
    Q_PROPERTY(bool alwaysOn READ alwaysOn WRITE setAlwaysOn NOTIFY alwaysOnChanged)
    Q_PROPERTY(int listenerId READ listenerId WRITE setListenerId NOTIFY listenerIdChanged)
    Q_PROPERTY(QList<int> waitFor READ waitFor WRITE setWaitFor NOTIFY waitForChanged)

public:
    explicit QFAppListener(QQuickItem *parent = 0);
    ~QFAppListener();

    /// Get the listening target.
    QObject *target() const;

    /// Set the listening target. If the class is constructed by QQmlComponent. It will be set automatically.
    void setTarget(QFDispatcher *target);

    /// Add a listener to the end of the listeners array for the specified message.  Multiple calls passing the same combination of event and listener will result in the listener being added multiple times.
    Q_INVOKABLE QFAppListener* on(QString type,QJSValue callback);

    /// Remove a listener from the listener array for the specified message.
    Q_INVOKABLE void removeListener(QString type,QJSValue callback);

    /// Remove all the listeners for a message with type. If type is empty, it will remove all the listeners.
    Q_INVOKABLE void removeAllListener(QString type = QString());

    /// Get the filter for incoming message
    QString filter() const;

    /// Set a filter to incoming message. Only message with type matched with the filter will emit "dispatched" signal.
    void setFilter(const QString &filter);

    /// Get a list of filter for incoming message
    QStringList filters() const;

    /// Set a list of filter to incoming message. Only message with type matched with the filters will emit "dispatched" signal.
    void setFilters(const QStringList &filters);

    bool alwaysOn() const;
    void setAlwaysOn(bool alwaysOn);

    int listenerId() const;
    void setListenerId(int listenerId);

    QList<int> waitFor() const;

    void setWaitFor(const QList<int> &waitFor);

signals:
    /// It is emitted whatever it has received a dispatched message from AppDispatcher.
    void dispatched(QString type,QJSValue message);

    void filterChanged();

    void filtersChanged();

    void alwaysOnChanged();

    void listenerIdChanged();

    void waitForChanged();

public slots:

private:
    virtual void componentComplete();

    Q_INVOKABLE void onMessageReceived(QString name,QJSValue message);

    void setListenerWaitFor();

    QPointer<QFDispatcher> m_target;

    QMap<QString,QList<QJSValue> >  mapping;

    QString m_filter;
    QStringList m_filters;
    bool m_alwaysOn;

    int m_listenerId;
    QFListener* m_listener;

    QList<int> m_waitFor;
};

#endif // QFAPPLISTENER_H
