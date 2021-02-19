import pygame

DIMENSION = {
    "x": 512,
    "y": 512,
}
win = pygame.display.set_mode((DIMENSION["x"], DIMENSION["y"]))


class Sprite:
    def __init__(self, x, y, width, height, img):
        self.x = x
        self.y = y
        self.width = width
        self.height = height
        self.sprite = pygame.image.load(f'images/{img}.png')

    def draw(self):
        win.blit(self.sprite, self.get_pos())

    def get_pos(self):
        return self.x, self.y


class Mover(Sprite):
    def __init__(self, x, y, width, height, vel, img, _id=0):
        super().__init__(x, y, width, height, img)
        self.id = _id
        self.vel = vel
        self.up = pygame.image.load(f'images/{img}_u.png')
        self.down = pygame.image.load(f'images/{img}_d.png')
        self.left = pygame.image.load(f'images/{img}_l.png')
        self.right = pygame.image.load(f'images/{img}_r.png')

    def move(self, ws):
        def _in_line(direction):
            movement = {
                "left": 0 < self.x,
                "right": self.x < DIMENSION["x"] - self.width,
                "up": 0 < self.y,
                "down": self.y < DIMENSION["y"] - self.height,
            }
            if movement[direction]:
                return True
            else:
                return False

        keys = pygame.key.get_pressed()

        if keys[pygame.K_h] and _in_line("left"):
            self.sprite = self.left
            self.x -= self.vel
            ws.send(f'id:{self.id},x:{self.x},y:{self.y},face:E')
        if keys[pygame.K_j] and _in_line("down"):
            self.sprite = self.down
            self.y += self.vel
            ws.send(f'id:{self.id},x:{self.x},y:{self.y},face:S')
        if keys[pygame.K_k] and _in_line("up"):
            self.sprite = self.up
            self.y -= self.vel
            ws.send(f'id:{self.id},x:{self.x},y:{self.y},face:N')
        if keys[pygame.K_l] and _in_line("right"):
            self.sprite = self.right
            self.x += self.vel
            ws.send(f'id:{self.id},x:{self.x},y:{self.y},face:W')

    def moved(self, direction):
        if direction == "none":
            pass
        elif direction == "left":
            self.sprite = self.right
            self.x += 5
        elif direction == "right":
            self.sprite = self.left
            self.x -= 5
        elif direction == "up":
            self.sprite = self.down
            self.y += 5
        elif direction == "down":
            self.sprite = self.up
            self.y -= 5

    def collide(self, other):
        if (self.x + 5 < other.x - 5
                and self.y < other.y + other.height - 5 and self.y + self.height > other.y + 5):
            if self.x + self.width > other.x + 5:
                self.x -= 5
                return "left"
        if (self.x + 5 > other.x - 5
                and self.y < other.y + other.height - 5 and self.y + self.height > other.y + 5):
            if self.x < other.x + other.width:
                self.x += 5
                return "right"
        if (self.y - 5 > other.y + 5
                and self.x < other.x + other.width - 5 and self.x + self.width > other.x + 5):
            if self.y < other.y + other.height:
                self.y += 5
                return "down"
        if (self.y + 5 < other.y - 5
                and self.x < other.x + other.width - 5 and self.x + self.width > other.x + 5):
            if self.y + self.height > other.y:
                self.y -= 5
                return "up"
        else:
            return "none"
