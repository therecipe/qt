#pragma once
#include <QObject>
#include <QJSValue>

class QFHook : public QObject
{
    Q_OBJECT
public:
    explicit QFHook(QObject *parent = 0);

    virtual void dispatch(QString type, QJSValue message) = 0;

signals:
    void dispatched(QString type, QJSValue message);

};
