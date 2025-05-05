#pragma once

#include <QtWidgets/QMainWindow>
#include "ui_QtSmallWidget.h"

class QtSmallWidget : public QMainWindow
{
    Q_OBJECT

public:
    QtSmallWidget(QWidget *parent = nullptr);
    ~QtSmallWidget();

private:
    Ui::QtSmallWidgetClass ui;
};
