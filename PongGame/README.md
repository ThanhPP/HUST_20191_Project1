# PONG GAME
_________________________________________
## Môi trường :
**_ Ngôn ngữ sử dụng** : Python 3

**_ Thư viện**: 
1. Pygame
2. random
_________________________________________
## Class 
### Paddle
_ Được thừa kế từ class Sprite
```python
class Paddle(pygame.sprite.Sprite):
    def __init__(self, color, width, height):
    def moveUp(self, pixels):
    def moveDown(self, pixels):
```
### Ball 
_ Được thừa kế từ class Sprite:
```python
class Ball(pygame.sprite.Sprite):
    def __init__(self, color, width, height):
    def update(self):
    def bounce(self):
```
_________________________________________
## main 
### Tạo cửa sổ chính
_ Tạo ra thông số cửa sổ mới
```python
size = (1000,1000)
screen = pygame.display.set_mode(size)
pygame.display.set_caption("Pong")
```
_  Khởi tạo cửa sổ, chờ sự kiện quit :
```python
while carryOn:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            carryOn = False
        elif event.type == pygame.KEYDOWN:
            if event.key == pygame.K_x:
                carryOn = False
pygame.quit()
```

_ Tạo ra cửa sổ, vẽ vạch ở giữa và cập nhật hình 60 lần/giây:
```python
#Draw basic screen
screen.fill(BLACK)
pygame.draw.line(screen, WHITE, [499, 0], [499, 1000], 10)
#Update the screen
pygame.display.flip()
clock.tick(60)

```

### Tạo ra các Paddle(hình vẽ)
_ Tạo ra 2 đối tượng Paddle :
```python
paddleA = Paddle(WHITE,10,100)
paddleA.rect.x = 20
paddleA.rect.y = 400

paddleB = Paddle(WHITE,10,100)
paddleB.rect.x = 970
paddleB.rect.y = 400
```

_ Tạo 1 Sprite group để chứa 2 Paddle :
```python
all_sprite_list = pygame.sprite.Group()
all_sprite_list.add(paddleA)
all_sprite_list.add(paddleB)
```

_ Update các phần tử trong sprite group : 
```python
all_sprite_list.update()
```

_ Vẽ các phần tử trong sprite group :
```python
all_sprite_list.draw(screen)
```

### Điều khiển các Paddle 
_ Bổ sung 2 hàm lên và xuống trong class Paddle : 
```python
def moveUp(self, pixels):
    self.rect.x -= pixels
    #Keep the Paddle in the screen
    if self.rect.x < 0 :
        self.rect.x = 0

def moveDown(self, pixels):
    self.rect.x += pixels
    #Keep the Paddle in the screen
    if self.rect.x > 1000 :
        self.rect = 1000
```
_ Bắt các sự kiện người dùng ấn nút W, S, UP, DOWN:
```python
keys = pygame.key.get_pressed()
if keys[pygame.K_w]:
    paddleA.moveUp(5)
if keys[pygame.K_s]:
    paddleA.moveDown(5)
if keys[pygame.K_UP]:
    paddleB.moveUp(5)
if keys[pygame.K_DOWN]:
    paddleB.moveDown(5)
```

### Thêm bóng nảy (hình vuông)
_ Khởi tạo trái bóng : 
```python
ball = Ball(WHITE,15,15)
ball.rect.x = 595
ball.rect.y = 595
```
_ Thêm bóng vào sprite group :
```python
all_sprite_list.add(ball)
```

_ Bóng nảy lại khi đập vào tường : 
```python
if ball.rect.x > 985:
    ball.velocity[0] = -ball.velocity[0]
if ball.rect.x < 0:
    ball.velocity[0] = -ball.velocity[0]
if ball.rect.y > 985:
    ball.velocity[1] = -ball.velocity[1]
if ball.rect.y < 0:
    ball.velocity[1] = -ball.velocity[1]
```

_ Bóng nảy khi đập vào Paddle : 
>pygame cung cấp hàm để xem 2 sprite có va chạm với nhau không 
```python
if pygame.sprite.collide_mask(ball,paddleA) or pygame.sprite.collide_mask(paddleB,ball):
    ball.bounce()
```
    
