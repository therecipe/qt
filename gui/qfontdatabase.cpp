#include "qfontdatabase.h"
#include <QFontInfo>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QFontDatabase>
#include "_cgo_export.h"

class MyQFontDatabase: public QFontDatabase {
public:
};

int QFontDatabase_Ogham_Type(){
	return QFontDatabase::Ogham;
}

int QFontDatabase_Runic_Type(){
	return QFontDatabase::Runic;
}

int QFontDatabase_Nko_Type(){
	return QFontDatabase::Nko;
}

int QFontDatabase_WritingSystemsCount_Type(){
	return QFontDatabase::WritingSystemsCount;
}

int QFontDatabase_QFontDatabase_RemoveAllApplicationFonts(){
	return QFontDatabase::removeAllApplicationFonts();
}

int QFontDatabase_QFontDatabase_RemoveApplicationFont(int id){
	return QFontDatabase::removeApplicationFont(id);
}

void* QFontDatabase_NewQFontDatabase(){
	return new QFontDatabase();
}

int QFontDatabase_QFontDatabase_AddApplicationFont(char* fileName){
	return QFontDatabase::addApplicationFont(QString(fileName));
}

int QFontDatabase_QFontDatabase_AddApplicationFontFromData(void* fontData){
	return QFontDatabase::addApplicationFontFromData(*static_cast<QByteArray*>(fontData));
}

char* QFontDatabase_QFontDatabase_ApplicationFontFamilies(int id){
	return QFontDatabase::applicationFontFamilies(id).join("|").toUtf8().data();
}

int QFontDatabase_Bold(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->bold(QString(family), QString(style));
}

char* QFontDatabase_Families(void* ptr, int writingSystem){
	return static_cast<QFontDatabase*>(ptr)->families(static_cast<QFontDatabase::WritingSystem>(writingSystem)).join("|").toUtf8().data();
}

int QFontDatabase_IsBitmapScalable(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isBitmapScalable(QString(family), QString(style));
}

int QFontDatabase_IsFixedPitch(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isFixedPitch(QString(family), QString(style));
}

int QFontDatabase_IsPrivateFamily(void* ptr, char* family){
	return static_cast<QFontDatabase*>(ptr)->isPrivateFamily(QString(family));
}

int QFontDatabase_IsScalable(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isScalable(QString(family), QString(style));
}

int QFontDatabase_IsSmoothlyScalable(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isSmoothlyScalable(QString(family), QString(style));
}

int QFontDatabase_Italic(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->italic(QString(family), QString(style));
}

char* QFontDatabase_StyleString(void* ptr, void* font){
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFont*>(font)).toUtf8().data();
}

char* QFontDatabase_StyleString2(void* ptr, void* fontInfo){
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFontInfo*>(fontInfo)).toUtf8().data();
}

char* QFontDatabase_Styles(void* ptr, char* family){
	return static_cast<QFontDatabase*>(ptr)->styles(QString(family)).join("|").toUtf8().data();
}

int QFontDatabase_Weight(void* ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->weight(QString(family), QString(style));
}

char* QFontDatabase_QFontDatabase_WritingSystemName(int writingSystem){
	return QFontDatabase::writingSystemName(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

char* QFontDatabase_QFontDatabase_WritingSystemSample(int writingSystem){
	return QFontDatabase::writingSystemSample(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

