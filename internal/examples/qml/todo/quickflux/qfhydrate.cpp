#include <QtCore>
#include <QVariantMap>
#include "qfhydrate.h"
#include "qfobject.h"
#include "qfstore.h"

static QVariantMap dehydrator(QObject* source);

static auto dehydratorFunction = [](const QStringList& ignoreList) {

    return [=](QObject* source) {
        QVariantMap dest;
        const QMetaObject* meta = source->metaObject();

        for (int i = 0 ; i < meta->propertyCount(); i++) {
            const QMetaProperty property = meta->property(i);
            const char* name = property.name();
            QString stringName = name;

            if (ignoreList.indexOf(stringName) >= 0) {
                continue;
            }

            QVariant value = source->property(name);

            if (value.canConvert<QObject*>()) {
                QObject* object = value.value<QObject*>();
                if (!object) {
                    continue;
                }
                value = dehydrator(object);
            }
            dest[stringName] = value;
        }
        return dest;
    };
};

static auto dehydrateQObject = dehydratorFunction(QStringList() << "parent" << "objectName");
static auto dehydrateQFObject = dehydratorFunction(QStringList() << "parent" << "objectName" << "children");
static auto dehydrateQFStore = dehydratorFunction(QStringList() << "parent" << "objectName" << "children" << "bindSource" << "redispatchTargets" << "filterFunctionEnabled");

/// Default dehydrator function
static QVariantMap dehydrator(QObject* source) {
    QVariantMap ret;
    if (qobject_cast<QFStore*>(source)) {
        ret = dehydrateQFStore(source);
    } else if (qobject_cast<QFObject*>(source)) {
        ret = dehydrateQFObject(source);
    } else {
        ret = dehydrateQObject(source);
    }
    return ret;
}

QFHydrate::QFHydrate(QObject *parent) : QObject(parent)
{

}

void QFHydrate::rehydrate(QObject *dest, const QVariantMap &source)
{
    const QMetaObject* meta = dest->metaObject();

    QMap<QString,QVariant>::const_iterator iter = source.begin();
    while (iter != source.end()) {
        QByteArray key = iter.key().toLocal8Bit();

        int index = meta->indexOfProperty(key.constData());
        if (index < 0) {
            qWarning() << QString("Hydrate.rehydrate: %1 property is not existed").arg(iter.key());
            iter++;
            continue;
        }

        QVariant orig = dest->property(key.constData());
        QVariant value = source[iter.key()];

        if (orig.canConvert<QObject*>()) {
            if (value.type() != QVariant::Map) {
                qWarning() << QString("Hydrate.rehydrate: expect a QVariantMap property but it is not: %1");
            } else {
                rehydrate(orig.value<QObject*>(), value.toMap());
            }

        } else if (orig != value) {
            dest->setProperty(key.constData(), value);
        }

        iter++;
    }
}

QVariantMap QFHydrate::dehydrate(QObject *source)
{
    return dehydrator(source);
}
