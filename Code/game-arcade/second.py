import pygame, sys
from models import Sprite, Mover, win
import websocket
import _thread as thread
import time

pygame.init()

player = Mover(0, 0, 30, 30, 5, "friend",1)
house = Sprite(250, 5, 170, 100, "house")
door = Sprite(321, 75, 30, 35, "door")
friend = Mover(50, 50, 30, 30, 5, "man",0)

bg = pygame.image.load("images/bg.png")

pygame.display.set_caption("First Game")

def on_message(ws, message):
    _id, x, y, face = [x.split(':')[1] for x in message.split(',')]
    if int(_id) != player.id:
        friend.x = int(x)
        friend.y = int(y)
        if face == 'W':
            friend.sprite = friend.right
        elif face == 'N':
            friend.sprite = friend.up
        elif face == 'S':
            friend.sprite = friend.down
        elif face == 'E':
            friend.sprite = friend.left

def on_error(ws, error):
    print(error)

def on_close(ws):
    print("### closed ###")

def on_open(ws):
    def run(*args):
        while True:
            pygame.time.delay(100)
            for event in pygame.event.get():
                if event.type == pygame.QUIT:
                    pygame.quit()
                    sys.exit(0)

            win.blit(bg, (0, 0))

            player.collide(house)
            friend.moved(player.collide(friend))
            friend.collide(house)
            if friend.collide(door) == "down":
                pygame.quit()
                sys.exit(0)

            player.move(ws)

            house.draw()
            door.draw()
            friend.draw()
            player.draw()

            pygame.display.update()
        ws.close()
    thread.start_new_thread(run, ())


if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://127.0.0.1:8080/ws",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
    ws.on_open = on_open
    ws.run_forever()

