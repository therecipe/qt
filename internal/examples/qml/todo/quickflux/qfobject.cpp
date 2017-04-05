#include "qfobject.h"

QFObject::QFObject(QObject *parent) : QObject(parent)
{

}

QQmlListProperty<QObject> QFObject::children()
{
    return QQmlListProperty<QObject>(qobject_cast<QObject*>(this), m_children);
}
