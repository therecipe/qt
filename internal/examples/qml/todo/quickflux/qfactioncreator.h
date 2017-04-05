#ifndef QFACTIONCREATOR_H
#define QFACTIONCREATOR_H

#include <QObject>
#include <QQmlParserStatus>
#include <QFAppDispatcher>
#include <QPointer>
#include "qfsignalproxy.h"

class QFActionCreator : public QObject, public QQmlParserStatus
{
    Q_OBJECT
    Q_INTERFACES(QQmlParserStatus)
    Q_PROPERTY(QFDispatcher* dispatcher READ dispatcher WRITE setDispatcher NOTIFY dispatcherChanged)

public:
    explicit QFActionCreator(QObject *parent = 0);

    QFDispatcher *dispatcher() const;
    void setDispatcher(QFDispatcher *value);

public slots:
    QString genKeyTable();
    void dispatch(QString type, QJSValue message = QJSValue());

signals:
    void dispatcherChanged();

protected:
    void classBegin();
    void componentComplete();

private:
    QPointer<QFDispatcher> m_dispatcher;
    QList<QFSignalProxy*> m_proxyList;
};

#endif // QFACTIONCREATOR_H
