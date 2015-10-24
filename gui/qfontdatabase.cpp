#include "qfontdatabase.h"
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QByteArray>
#include <QFontInfo>
#include <QString>
#include <QVariant>
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

QtObjectPtr QFontDatabase_NewQFontDatabase(){
	return new QFontDatabase();
}

int QFontDatabase_QFontDatabase_AddApplicationFont(char* fileName){
	return QFontDatabase::addApplicationFont(QString(fileName));
}

int QFontDatabase_QFontDatabase_AddApplicationFontFromData(QtObjectPtr fontData){
	return QFontDatabase::addApplicationFontFromData(*static_cast<QByteArray*>(fontData));
}

char* QFontDatabase_QFontDatabase_ApplicationFontFamilies(int id){
	return QFontDatabase::applicationFontFamilies(id).join("|").toUtf8().data();
}

int QFontDatabase_Bold(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->bold(QString(family), QString(style));
}

char* QFontDatabase_Families(QtObjectPtr ptr, int writingSystem){
	return static_cast<QFontDatabase*>(ptr)->families(static_cast<QFontDatabase::WritingSystem>(writingSystem)).join("|").toUtf8().data();
}

int QFontDatabase_IsBitmapScalable(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isBitmapScalable(QString(family), QString(style));
}

int QFontDatabase_IsFixedPitch(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isFixedPitch(QString(family), QString(style));
}

int QFontDatabase_IsPrivateFamily(QtObjectPtr ptr, char* family){
	return static_cast<QFontDatabase*>(ptr)->isPrivateFamily(QString(family));
}

int QFontDatabase_IsScalable(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isScalable(QString(family), QString(style));
}

int QFontDatabase_IsSmoothlyScalable(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->isSmoothlyScalable(QString(family), QString(style));
}

int QFontDatabase_Italic(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->italic(QString(family), QString(style));
}

char* QFontDatabase_StyleString(QtObjectPtr ptr, QtObjectPtr font){
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFont*>(font)).toUtf8().data();
}

char* QFontDatabase_StyleString2(QtObjectPtr ptr, QtObjectPtr fontInfo){
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFontInfo*>(fontInfo)).toUtf8().data();
}

char* QFontDatabase_Styles(QtObjectPtr ptr, char* family){
	return static_cast<QFontDatabase*>(ptr)->styles(QString(family)).join("|").toUtf8().data();
}

int QFontDatabase_Weight(QtObjectPtr ptr, char* family, char* style){
	return static_cast<QFontDatabase*>(ptr)->weight(QString(family), QString(style));
}

char* QFontDatabase_QFontDatabase_WritingSystemName(int writingSystem){
	return QFontDatabase::writingSystemName(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

char* QFontDatabase_QFontDatabase_WritingSystemSample(int writingSystem){
	return QFontDatabase::writingSystemSample(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

