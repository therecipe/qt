#ifndef QFHYDRATE_H
#define QFHYDRATE_H

#include <QObject>
#include <QVariantMap>

class QFHydrate : public QObject
{
    Q_OBJECT
public:
    explicit QFHydrate(QObject *parent = 0);

signals:

public slots:
    void rehydrate(QObject *dest, const QVariantMap & source);

    QVariantMap dehydrate(QObject *source);
};

#endif // QFHYDRATE_H
