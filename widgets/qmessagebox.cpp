#include "qmessagebox.h"
#include <QCheckBox>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPushButton>
#include <QWidget>
#include <QObject>
#include <QModelIndex>
#include <QMetaObject>
#include <QPixmap>
#include <QAbstractButton>
#include <QMessageBox>
#include "_cgo_export.h"

class MyQMessageBox: public QMessageBox {
public:
void Signal_ButtonClicked(QAbstractButton * button){callbackQMessageBoxButtonClicked(this->objectName().toUtf8().data(), button);};
};

int QMessageBox_ButtonMask_Type(){
	return QMessageBox::ButtonMask;
}

char* QMessageBox_DetailedText(void* ptr){
	return static_cast<QMessageBox*>(ptr)->detailedText().toUtf8().data();
}

int QMessageBox_Icon(void* ptr){
	return static_cast<QMessageBox*>(ptr)->icon();
}

char* QMessageBox_InformativeText(void* ptr){
	return static_cast<QMessageBox*>(ptr)->informativeText().toUtf8().data();
}

void QMessageBox_SetDetailedText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setDetailedText(QString(text));
}

void QMessageBox_SetIcon(void* ptr, int v){
	static_cast<QMessageBox*>(ptr)->setIcon(static_cast<QMessageBox::Icon>(v));
}

void QMessageBox_SetIconPixmap(void* ptr, void* pixmap){
	static_cast<QMessageBox*>(ptr)->setIconPixmap(*static_cast<QPixmap*>(pixmap));
}

void QMessageBox_SetInformativeText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setInformativeText(QString(text));
}

void QMessageBox_SetStandardButtons(void* ptr, int buttons){
	static_cast<QMessageBox*>(ptr)->setStandardButtons(static_cast<QMessageBox::StandardButton>(buttons));
}

void QMessageBox_SetText(void* ptr, char* text){
	static_cast<QMessageBox*>(ptr)->setText(QString(text));
}

void QMessageBox_SetTextFormat(void* ptr, int format){
	static_cast<QMessageBox*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(format));
}

void QMessageBox_SetTextInteractionFlags(void* ptr, int flags){
	static_cast<QMessageBox*>(ptr)->setTextInteractionFlags(static_cast<Qt::TextInteractionFlag>(flags));
}

int QMessageBox_StandardButtons(void* ptr){
	return static_cast<QMessageBox*>(ptr)->standardButtons();
}

char* QMessageBox_Text(void* ptr){
	return static_cast<QMessageBox*>(ptr)->text().toUtf8().data();
}

int QMessageBox_TextFormat(void* ptr){
	return static_cast<QMessageBox*>(ptr)->textFormat();
}

int QMessageBox_TextInteractionFlags(void* ptr){
	return static_cast<QMessageBox*>(ptr)->textInteractionFlags();
}

void* QMessageBox_NewQMessageBox2(int icon, char* title, char* text, int buttons, void* parent, int f){
	return new QMessageBox(static_cast<QMessageBox::Icon>(icon), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QMessageBox_NewQMessageBox(void* parent){
	return new QMessageBox(static_cast<QWidget*>(parent));
}

void QMessageBox_QMessageBox_About(void* parent, char* title, char* text){
	QMessageBox::about(static_cast<QWidget*>(parent), QString(title), QString(text));
}

void QMessageBox_QMessageBox_AboutQt(void* parent, char* title){
	QMessageBox::aboutQt(static_cast<QWidget*>(parent), QString(title));
}

void* QMessageBox_AddButton3(void* ptr, int button){
	return static_cast<QMessageBox*>(ptr)->addButton(static_cast<QMessageBox::StandardButton>(button));
}

void* QMessageBox_AddButton2(void* ptr, char* text, int role){
	return static_cast<QMessageBox*>(ptr)->addButton(QString(text), static_cast<QMessageBox::ButtonRole>(role));
}

void QMessageBox_AddButton(void* ptr, void* button, int role){
	static_cast<QMessageBox*>(ptr)->addButton(static_cast<QAbstractButton*>(button), static_cast<QMessageBox::ButtonRole>(role));
}

void* QMessageBox_Button(void* ptr, int which){
	return static_cast<QMessageBox*>(ptr)->button(static_cast<QMessageBox::StandardButton>(which));
}

void QMessageBox_ConnectButtonClicked(void* ptr){
	QObject::connect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

void QMessageBox_DisconnectButtonClicked(void* ptr){
	QObject::disconnect(static_cast<QMessageBox*>(ptr), static_cast<void (QMessageBox::*)(QAbstractButton *)>(&QMessageBox::buttonClicked), static_cast<MyQMessageBox*>(ptr), static_cast<void (MyQMessageBox::*)(QAbstractButton *)>(&MyQMessageBox::Signal_ButtonClicked));;
}

int QMessageBox_ButtonRole(void* ptr, void* button){
	return static_cast<QMessageBox*>(ptr)->buttonRole(static_cast<QAbstractButton*>(button));
}

void* QMessageBox_CheckBox(void* ptr){
	return static_cast<QMessageBox*>(ptr)->checkBox();
}

void* QMessageBox_ClickedButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->clickedButton();
}

int QMessageBox_QMessageBox_Critical(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::critical(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void* QMessageBox_DefaultButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->defaultButton();
}

void* QMessageBox_EscapeButton(void* ptr){
	return static_cast<QMessageBox*>(ptr)->escapeButton();
}

int QMessageBox_Exec(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QMessageBox*>(ptr), "exec");
}

int QMessageBox_QMessageBox_Information(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::information(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_Open(void* ptr, void* receiver, char* member){
	static_cast<QMessageBox*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QMessageBox_QMessageBox_Question(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::question(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_RemoveButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->removeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetCheckBox(void* ptr, void* cb){
	static_cast<QMessageBox*>(ptr)->setCheckBox(static_cast<QCheckBox*>(cb));
}

void QMessageBox_SetDefaultButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QPushButton*>(button));
}

void QMessageBox_SetDefaultButton2(void* ptr, int button){
	static_cast<QMessageBox*>(ptr)->setDefaultButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetEscapeButton(void* ptr, void* button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QAbstractButton*>(button));
}

void QMessageBox_SetEscapeButton2(void* ptr, int button){
	static_cast<QMessageBox*>(ptr)->setEscapeButton(static_cast<QMessageBox::StandardButton>(button));
}

void QMessageBox_SetVisible(void* ptr, int visible){
	static_cast<QMessageBox*>(ptr)->setVisible(visible != 0);
}

void QMessageBox_SetWindowModality(void* ptr, int windowModality){
	static_cast<QMessageBox*>(ptr)->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

void QMessageBox_SetWindowTitle(void* ptr, char* title){
	static_cast<QMessageBox*>(ptr)->setWindowTitle(QString(title));
}

int QMessageBox_StandardButton(void* ptr, void* button){
	return static_cast<QMessageBox*>(ptr)->standardButton(static_cast<QAbstractButton*>(button));
}

int QMessageBox_QMessageBox_Warning(void* parent, char* title, char* text, int buttons, int defaultButton){
	return QMessageBox::warning(static_cast<QWidget*>(parent), QString(title), QString(text), static_cast<QMessageBox::StandardButton>(buttons), static_cast<QMessageBox::StandardButton>(defaultButton));
}

void QMessageBox_DestroyQMessageBox(void* ptr){
	static_cast<QMessageBox*>(ptr)->~QMessageBox();
}

