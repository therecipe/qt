#ifndef QFAPPLISTENERGROUP_H
#define QFAPPLISTENERGROUP_H

#include <QQuickItem>
#include <QQmlParserStatus>
#include "qflistener.h"

class QFAppListenerGroup : public QQuickItem
{
    Q_OBJECT
    Q_PROPERTY(QList<int> listenerIds READ listenerIds WRITE setListenerIds NOTIFY listenerIdsChanged)
    Q_PROPERTY(QList<int> waitFor READ waitFor WRITE setWaitFor NOTIFY waitForChanged)

public:
    QFAppListenerGroup(QQuickItem* parent = 0);

    QList<int> listenerIds() const;

    void setListenerIds(const QList<int> &listenerIds);

    QList<int> waitFor() const;

    void setWaitFor(const QList<int> &waitFor);

public slots:

signals:

    void listenerIdsChanged();
    void waitForChanged();

private:
    virtual void componentComplete();

    QList<int> search(QQuickItem* item);

    void setListenerWaitFor();

    QList<int> m_waitFor;

    QList<int> m_listenerIds;
    int m_listenerId;
    QFListener* m_listener;
};

#endif // QFAPPLISTENERGROUP_H
