#include "qlatin1char.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QLatin1Char>
#include "_cgo_export.h"

class MyQLatin1Char: public QLatin1Char {
public:
};

void* QLatin1Char_NewQLatin1Char(char* c){
	return new QLatin1Char(*c);
}

