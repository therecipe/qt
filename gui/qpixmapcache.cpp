#include "qpixmapcache.h"
#include <QPixmap>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPixmapCache>
#include "_cgo_export.h"

class MyQPixmapCache: public QPixmapCache {
public:
};

int QPixmapCache_QPixmapCache_CacheLimit(){
	return QPixmapCache::cacheLimit();
}

void QPixmapCache_QPixmapCache_Clear(){
	QPixmapCache::clear();
}

void QPixmapCache_QPixmapCache_SetCacheLimit(int n){
	QPixmapCache::setCacheLimit(n);
}

