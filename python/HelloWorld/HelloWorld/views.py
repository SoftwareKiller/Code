'''
from django.http import HttpResponse

def hello(request):
    return HttpResponse("Hello world !")
'''
from django.shortcuts import render

# def runoob(request):
#     context = {}
#     context['hello'] = 'Hello World!'
#     return render(request, 'runoob.html', context)

def runoob(request):
    view_name = "View Test"
    return render(request, "runoob.html", {"name": view_name})