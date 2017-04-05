#ifndef QFMIDDLEWARESHOOK_H
#define QFMIDDLEWARESHOOK_H

#include <QObject>
#include <QQmlEngine>
#include <QPointer>
#include "qfhook.h"

class QFMiddlewaresHook : public QFHook
{
    Q_OBJECT
public:
    explicit QFMiddlewaresHook(QObject *parent = 0);

signals:

public:
    void dispatch(QString type, QJSValue message);
    void setup(QQmlEngine* engine, QObject* middlewares);

public slots:
    void next(int senderId, QString type, QJSValue message);
    void resolve(QString type, QJSValue message);


private:
    QJSValue invoke;
    QPointer<QObject> m_middlewares;
};

#endif // QFMIDDLEWARESHOOK_H
