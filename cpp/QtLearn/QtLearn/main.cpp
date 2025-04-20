#include "QtLearn.h"
#include <QtWidgets/QApplication>
#include <QtWidgets/qlabel.h>

int main(int argc, char* argv[])
{
	QApplication a(argc, argv);
	QtLearn w;
	w.show();
	/*   QLabel lable(QLabel::tr("Hello Qt!"));
	   lable.show();*/
	return a.exec();
}
