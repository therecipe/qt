// +build sailfish sailfish_emulator

#define protected public
#define private public

#include "sailfish_sailfish.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QList>
#include <QString>
#include <QUrl>
#include <sailfishapp.h>

void* SailfishApp_SailfishApp_Application(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return SailfishApp::application(argcs, argvs);
}

int SailfishApp_SailfishApp_Main(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return SailfishApp::main(argcs, argvs);
}

void* SailfishApp_SailfishApp_CreateView()
{
	return SailfishApp::createView();
}

void* SailfishApp_SailfishApp_PathTo(struct QtSailfish_PackedString filename)
{
	return new QUrl(SailfishApp::pathTo(*(new QString(QString::fromUtf8(filename.data, filename.len)))));
}

