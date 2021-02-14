from tkinter import *
from tkinter import messagebox
from tkinter.ttk import *
import requests,json, uuid, re

def ok():
    user            = {}
    url             = 'https://api.backend.mama.sh/'
    user['ip']      = requests.get('http://httpbin.org/ip').json()['origin']
    user['mac']     = ':'.join(re.findall('..', '%012x' % uuid.getnode()))
    user['name']    = name_var.get()
    res             = requests.post(url,json=user).json()
    message         = f"""
Hi {res['user']['name']}
{res['content']}
your information is:

ip: {res['user']['ip']}
mac: {res['user']['mac']}
"""
    messagebox.showinfo("showinfo", message)

root        = Tk()
name_var    = StringVar()

root.title("First program")

l = Label(root, text = "Your Name:" )
e = Entry(root, textvariable = name_var)
b = Button(root, text = 'OK', command = ok)

l.grid(row=0,column=0, sticky = W, pady = 2, padx = 3)
e.grid(row=0,column=1, pady = 2)
b.grid(row=1,column=1)

root.mainloop()

