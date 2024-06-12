"""
WSGI config for HelloWorld project.

It exposes the WSGI callable as a module-level variable named ``application``.

For more information on this file, see
https://docs.djangoproject.com/en/5.0/howto/deployment/wsgi/
"""

import os

from django.core.wsgi import get_wsgi_application

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'HelloWorld.settings')

application = get_wsgi_application()

from TestModel.task import scheduler

def application(environ, start_response):
    # 启动定时任务调度器
    scheduler.start()