#ifndef QFOBJECT_H
#define QFOBJECT_H

#include <QObject>
#include <QQmlListProperty>

/// QFObject provides an QtObject with default children property to hold nested items

class QFObject : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QQmlListProperty<QObject> children READ children)
    Q_CLASSINFO("DefaultProperty", "children")

public:
    explicit QFObject(QObject *parent = 0);

    QQmlListProperty<QObject> children();

signals:

public slots:

private:
     QObjectList m_children;
};

#endif // QFOBJECT_H
