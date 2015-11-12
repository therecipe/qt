#include "qlatin1string.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QLatin1String>
#include "_cgo_export.h"

class MyQLatin1String: public QLatin1String {
public:
};

void* QLatin1String_NewQLatin1String3(void* str){
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

void* QLatin1String_NewQLatin1String(char* str){
	return new QLatin1String(const_cast<const char*>(str));
}

void* QLatin1String_NewQLatin1String2(char* str, int size){
	return new QLatin1String(const_cast<const char*>(str), size);
}

int QLatin1String_Size(void* ptr){
	return static_cast<QLatin1String*>(ptr)->size();
}

