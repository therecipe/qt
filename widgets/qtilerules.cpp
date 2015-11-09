#include "qtilerules.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTileRules>
#include "_cgo_export.h"

class MyQTileRules: public QTileRules {
public:
};

void* QTileRules_NewQTileRules(int horizontalRule, int verticalRule){
	return new QTileRules(static_cast<Qt::TileRule>(horizontalRule), static_cast<Qt::TileRule>(verticalRule));
}

void* QTileRules_NewQTileRules2(int rule){
	return new QTileRules(static_cast<Qt::TileRule>(rule));
}

