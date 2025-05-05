#include "QtDialog.h"
#include <QDialog>
#include <QDebug>

QtDialog::QtDialog(QWidget* parent)
	: QMainWindow(parent)
{
	resize(600, 400);

	// 菜单栏创建，最多只能有一个
	QMenuBar* bar = menuBar();
	setMenuBar(bar);

	// 创建菜单
	QMenu* fileMenu = bar->addMenu("文件");
	QMenu* editMenu = bar->addMenu("编辑");

	// 创建菜单项
	QAction* newAction = fileMenu->addAction("新建");
	// 添加分隔符
	fileMenu->addSeparator();

	QAction* openAction = fileMenu->addAction("打开");

	// 点击新建按钮 弹出一个对话框
	connect(newAction, &QAction::triggered, [=]() {
		// 对话框分类
		// 模态对话框(不可以对其他窗口进行操作)
		// 非模态对话框(可以对其他窗口进行操作)

		// 模态创建
		/*QDialog dlg(this);
		dlg.resize(200, 100);
		dlg.exec();

		qDebug() << "模态对话框弹出完毕";*/

		// 非模态对话框
		QDialog* dlg2 = new QDialog(this);
		dlg2->resize(200, 100);

		dlg2->show();
		// 窗口关闭时释放
		dlg2->setAttribute(Qt::WA_DeleteOnClose);
		qDebug() << "非模态对话框弹出";
		});
}

QtDialog::~QtDialog()
{
}
