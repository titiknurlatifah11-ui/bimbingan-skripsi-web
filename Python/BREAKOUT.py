import pygame
from pygame.locals import *
import os

pygame.init()

# Ukuran layar
screen_width = 1366
screen_height = 768

screen = pygame.display.set_mode((screen_width, screen_height))
pygame.display.set_caption('Breakout')

# Memuat gambar latar belakang dan mengubah ukurannya
background_img = pygame.image.load(os.path.join('assets', 'background.png'))
background_img = pygame.transform.scale(background_img, (screen_width, screen_height))

# Memuat suara
bounce_sound = pygame.mixer.Sound(os.path.join('assets', 'bounce.mp3'))
win_sound = pygame.mixer.Sound(os.path.join('assets', 'win.mp3'))
lose_sound = pygame.mixer.Sound(os.path.join('assets', 'lose.mp3'))

# Mendefinisikan font
font = pygame.font.SysFont('Constantia', 30)

# Mendefinisikan warna
bg = (234, 218, 184)
# Warna blok (brick colors)
brick_red = (150, 40, 27)
brick_brown = (165, 42, 42)
brick_dark = (101, 67, 33)
# Warna paddle dan bola
paddle_col = (0, 0, 255)  # Biru terang
paddle_outline = (255, 255, 255)  # Putih
ball_col = (0, 255, 0)  # Hijau terang
# Warna teks
text_col = (255, 255, 255)  # Putih

# Mendefinisikan variabel game
cols = 10  # 10 kolom
rows = 6   # 6 baris
clock = pygame.time.Clock()
fps = 60
live_ball = False
game_over = 0
score = 0

# Fungsi untuk menampilkan teks di layar
def draw_text(text, font, text_col, x, y):
    img = font.render(text, True, text_col)
    screen.blit(img, (x, y))

# Kelas dinding bata (wall)
class wall():
    def __init__(self):
        self.width = screen_width // cols
        self.height = 50

    def create_wall(self):
        self.blocks = []
        for row in range(rows):
            block_row = []
            for col in range(cols):
                block_x = col * self.width
                block_y = row * self.height
                rect = pygame.Rect(block_x, block_y, self.width, self.height)
                if row < 2:
                    strength = 3
                elif row < 4:
                    strength = 2
                elif row < 6:
                    strength = 1
                block_individual = [rect, strength]
                block_row.append(block_individual)
            self.blocks.append(block_row)

    def draw_wall(self):
        for row in self.blocks:
            for block in row:
                if block[1] == 3:
                    block_col = brick_dark
                elif block[1] == 2:
                    block_col = brick_brown
                elif block[1] == 1:
                    block_col = brick_red
                pygame.draw.rect(screen, block_col, block[0])
                pygame.draw.rect(screen, bg, block[0], 2)

# Kelas paddle
class paddle():
    def __init__(self):
        self.reset()

    def move(self):
        self.direction = 0
        key = pygame.key.get_pressed()
        if key[pygame.K_LEFT] and self.rect.left > 0:
            self.rect.x -= self.speed
            self.direction = -1
        if key[pygame.K_RIGHT] and self.rect.right < screen_width:
            self.rect.x += self.speed
            self.direction = 1

    def draw(self):
        pygame.draw.rect(screen, paddle_col, self.rect)
        pygame.draw.rect(screen, paddle_outline, self.rect, 3)

    def reset(self):
        self.height = 20
        self.width = int(screen_width / cols)
        self.x = int((screen_width / 2) - (self.width / 2))
        self.y = screen_height - (self.height * 2)
        self.speed = 10
        self.rect = Rect(self.x, self.y, self.width, self.height)
        self.direction = 0

# Kelas bola (ball)
class game_ball():
    def __init__(self, x, y):
        self.reset(x, y)

    def move(self):
        global score

        collision_thresh = 5
        wall_destroyed = 1
        row_count = 0
        for row in wall.blocks:
            item_count = 0
            for item in row:
                if self.rect.colliderect(item[0]):
                    if abs(self.rect.bottom - item[0].top) < collision_thresh and self.speed_y > 0:
                        self.speed_y *= -1
                    if abs(self.rect.top - item[0].bottom) < collision_thresh and self.speed_y < 0:
                        self.speed_y *= -1
                    if abs(self.rect.right - item[0].left) < collision_thresh and self.speed_x > 0:
                        self.speed_x *= -1
                    if abs(self.rect.left - item[0].right) < collision_thresh and self.speed_x < 0:
                        self.speed_x *= -1
                    if wall.blocks[row_count][item_count][1] > 1:
                        wall.blocks[row_count][item_count][1] -= 1
                    else:
                        wall.blocks[row_count][item_count][0] = (0, 0, 0, 0)
                        score += 10
                        bounce_sound.play()
                if wall.blocks[row_count][item_count][0] != (0, 0, 0, 0):
                    wall_destroyed = 0
                item_count += 1
            row_count += 1
        if wall_destroyed == 1:
            self.game_over = 1
            win_sound.play()
        if self.rect.left < 0 or self.rect.right > screen_width:
            self.speed_x *= -1
        if self.rect.top < 0:
            self.speed_y *= -1
        if self.rect.bottom > screen_height:
            self.game_over = -1
            lose_sound.play()
        if self.rect.colliderect(player_paddle):
            if abs(self.rect.bottom - player_paddle.rect.top) < collision_thresh and self.speed_y > 0:
                self.speed_y *= -1
                self.speed_x += player_paddle.direction
                if self.speed_x > self.speed_max:
                    self.speed_x = self.speed_max
                elif self.speed_x < 0 and self.speed_x < -self.speed_max:
                    self.speed_x = -self.speed_max
            else:
                self.speed_x *= -1
        self.rect.x += self.speed_x
        self.rect.y += self.speed_y
        return self.game_over

    def draw(self):
        pygame.draw.circle(screen, ball_col, (self.rect.x + self.ball_rad, self.rect.y + self.ball_rad), self.ball_rad)
        pygame.draw.circle(screen, paddle_outline, (self.rect.x + self.ball_rad, self.rect.y + self.ball_rad), self.ball_rad, 3)

    def reset(self, x, y):
        self.ball_rad = 10
        self.x = x - self.ball_rad
        self.y = y
        self.rect = Rect(self.x, self.y, self.ball_rad * 2, self.ball_rad * 2)
        self.speed_x = 4
        self.speed_y = -4
        self.speed_max = 5
        self.game_over = 0

# Membuat dinding
wall = wall()
wall.create_wall()

# Membuat paddle
player_paddle = paddle()

# Membuat bola
ball = game_ball(player_paddle.x + (player_paddle.width // 2), player_paddle.y - player_paddle.height)

run = True
while run:
    clock.tick(fps)
    
    screen.blit(background_img, (0, 0))  # Menggambar gambar latar belakang yang telah diubah ukurannya
    
    # Menggambar semua objek
    wall.draw_wall()
    player_paddle.draw()
    ball.draw()

    if live_ball:
        player_paddle.move()
        game_over = ball.move()
        if game_over != 0:
            live_ball = False

    # Menampilkan instruksi pemain atau pesan game over
    if not live_ball:
        if game_over == 0:
            draw_text('KLIK DIMANA SAJA UNTUK MEMULAI', font, text_col, 400, screen_height // 2 + 100)
        elif game_over == 1:
            draw_text('KAMU MENANG!', font, text_col, 550, screen_height // 2 + 50)
            draw_text('KLIK DIMANA SAJA UNTUK MEMULAI', font, text_col, 100, screen_height // 2 + 100)
        elif game_over == -1:
            draw_text('KAMU KALAH!', font, text_col, 550, screen_height // 2 + 50)
            draw_text('KLIK DIMANA SAJA UNTUK MEMULAI', font, text_col, 400, screen_height // 2 + 100)

    # Tampilkan skor di layar
    draw_text(f'Skor: {score}', font, text_col, 5, 5)

    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            run = False
        if event.type == pygame.MOUSEBUTTONDOWN and not live_ball:
            live_ball = True
            ball.reset(player_paddle.x + (player_paddle.width // 2), player_paddle.y - player_paddle.height)
            player_paddle.reset()
            wall.create_wall()
            score = 0  # Reset skor saat memulai permainan baru

    pygame.display.update()

pygame.quit()
