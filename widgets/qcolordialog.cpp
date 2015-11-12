#include "qcolordialog.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QColor>
#include <QColorDialog>
#include "_cgo_export.h"

class MyQColorDialog: public QColorDialog {
public:
void Signal_ColorSelected(const QColor & color){callbackQColorDialogColorSelected(this->objectName().toUtf8().data(), new QColor(color));};
void Signal_CurrentColorChanged(const QColor & color){callbackQColorDialogCurrentColorChanged(this->objectName().toUtf8().data(), new QColor(color));};
};

void* QColorDialog_CurrentColor(void* ptr){
	return new QColor(static_cast<QColorDialog*>(ptr)->currentColor());
}

int QColorDialog_Options(void* ptr){
	return static_cast<QColorDialog*>(ptr)->options();
}

void QColorDialog_SetCurrentColor(void* ptr, void* color){
	static_cast<QColorDialog*>(ptr)->setCurrentColor(*static_cast<QColor*>(color));
}

void QColorDialog_SetOptions(void* ptr, int options){
	static_cast<QColorDialog*>(ptr)->setOptions(static_cast<QColorDialog::ColorDialogOption>(options));
}

void* QColorDialog_NewQColorDialog(void* parent){
	return new QColorDialog(static_cast<QWidget*>(parent));
}

void* QColorDialog_NewQColorDialog2(void* initial, void* parent){
	return new QColorDialog(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent));
}

void QColorDialog_ConnectColorSelected(void* ptr){
	QObject::connect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::colorSelected), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_ColorSelected));;
}

void QColorDialog_DisconnectColorSelected(void* ptr){
	QObject::disconnect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::colorSelected), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_ColorSelected));;
}

void QColorDialog_ConnectCurrentColorChanged(void* ptr){
	QObject::connect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::currentColorChanged), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_CurrentColorChanged));;
}

void QColorDialog_DisconnectCurrentColorChanged(void* ptr){
	QObject::disconnect(static_cast<QColorDialog*>(ptr), static_cast<void (QColorDialog::*)(const QColor &)>(&QColorDialog::currentColorChanged), static_cast<MyQColorDialog*>(ptr), static_cast<void (MyQColorDialog::*)(const QColor &)>(&MyQColorDialog::Signal_CurrentColorChanged));;
}

void* QColorDialog_QColorDialog_CustomColor(int index){
	return new QColor(QColorDialog::customColor(index));
}

int QColorDialog_QColorDialog_CustomCount(){
	return QColorDialog::customCount();
}

void* QColorDialog_QColorDialog_GetColor(void* initial, void* parent, char* title, int options){
	return new QColor(QColorDialog::getColor(*static_cast<QColor*>(initial), static_cast<QWidget*>(parent), QString(title), static_cast<QColorDialog::ColorDialogOption>(options)));
}

void QColorDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QColorDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QColorDialog_SelectedColor(void* ptr){
	return new QColor(static_cast<QColorDialog*>(ptr)->selectedColor());
}

void QColorDialog_QColorDialog_SetCustomColor(int index, void* color){
	QColorDialog::setCustomColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetOption(void* ptr, int option, int on){
	static_cast<QColorDialog*>(ptr)->setOption(static_cast<QColorDialog::ColorDialogOption>(option), on != 0);
}

void QColorDialog_QColorDialog_SetStandardColor(int index, void* color){
	QColorDialog::setStandardColor(index, *static_cast<QColor*>(color));
}

void QColorDialog_SetVisible(void* ptr, int visible){
	static_cast<QColorDialog*>(ptr)->setVisible(visible != 0);
}

void* QColorDialog_QColorDialog_StandardColor(int index){
	return new QColor(QColorDialog::standardColor(index));
}

int QColorDialog_TestOption(void* ptr, int option){
	return static_cast<QColorDialog*>(ptr)->testOption(static_cast<QColorDialog::ColorDialogOption>(option));
}

void QColorDialog_DestroyQColorDialog(void* ptr){
	static_cast<QColorDialog*>(ptr)->~QColorDialog();
}

