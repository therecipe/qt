#ifndef QFAPPDISPATCHER_H
#define QFAPPDISPATCHER_H

#include <QObject>
#include "qfdispatcher.h"

class QFAppDispatcher : public QFDispatcher
{
    Q_OBJECT
public:
    explicit QFAppDispatcher(QObject *parent = 0);

    /// Obtain the singleton instance of AppDispatcher for specific QQmlEngine
    static QFAppDispatcher* instance(QQmlEngine* engine);

    /// Obtain a singleton object from package for specific QQmlEngine
    static QObject* singletonObject(QQmlEngine* engine,QString package,
                                    int versionMajor,
                                    int versionMinor,
                                    QString typeName);
};

#endif // QFAPPDISPATCHER_H
