#include "QtDialog.h"
#include <QtWidgets/QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QtDialog w;
    w.show();
    return a.exec();
}
