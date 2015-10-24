#include "qlatin1string.h"
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include "_cgo_export.h"

class MyQLatin1String: public QLatin1String {
public:
};

QtObjectPtr QLatin1String_NewQLatin1String3(QtObjectPtr str){
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

QtObjectPtr QLatin1String_NewQLatin1String(char* str){
	return new QLatin1String(const_cast<const char*>(str));
}

QtObjectPtr QLatin1String_NewQLatin1String2(char* str, int size){
	return new QLatin1String(const_cast<const char*>(str), size);
}

int QLatin1String_Size(QtObjectPtr ptr){
	return static_cast<QLatin1String*>(ptr)->size();
}

