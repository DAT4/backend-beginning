from tkinter import *


class Window(Frame):
    def __init__(self, master=None):
        Frame.__init__(self, master)
        self.master = master


def start():
    root = Tk()
    app = Window(root)
    root.eval('tk::PlaceWindow . center')

    root.wm_title('title')
    root.geometry('300x400')
    root.mainloop()


if __name__ == '__main__':
    start()
