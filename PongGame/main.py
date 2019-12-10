import pygame
import random

# Define some colors
BLACK = (0,0,0)
WHITE = (255,255,255)

pygame.init()
# Class
class Paddle(pygame.sprite.Sprite):
    def __init__(self, color, width, height):
        #Sprite constructor
        super().__init__()
        self.image = pygame.Surface([width, height])
        self.image.fill(BLACK)
        self.image.set_colorkey(BLACK)

        #draw a rectangle paddle
        pygame.draw.rect(self.image, color, [0, 0, width, height])

        #rectangle object with the dimension of the image
        self.rect = self.image.get_rect()

    def moveUp(self, pixels):
        self.rect.y -= pixels
        #Keep the Paddle in the screen
        if self.rect.y < 0 :
            self.rect.y = 0

    def moveDown(self, pixels):
        self.rect.y += pixels
        #Keep the Paddle in the screen
        if self.rect.y > 900 :
            self.rect.y = 900

class Ball(pygame.sprite.Sprite):
    def __init__(self,color,width,height):
        super().__init__()

        self.image = pygame.Surface([width,height])
        self.image.fill(BLACK)
        self.image.set_colorkey(BLACK)

        #Draw
        pygame.draw.rect(self.image,color,[0,0,width,height])
        #random to prevent bounce back to the same direction
        self.velocity = [random.randint(4,8), random.randint(4,8)]

        self.rect = self.image.get_rect()

    def update(self):
        self.rect.x += self.velocity[0]
        self.rect.y += self.velocity[1]

    def bounce(self):
        self.velocity[0] = -self.velocity[0]
        self.velocity[1] = random.randint(4,8)

#Open a window
size = (1000,1000)
screen = pygame.display.set_mode(size)
pygame.display.set_caption("Pong")

paddleA = Paddle(WHITE,10,100)
paddleA.rect.x = 20
paddleA.rect.y = 400

paddleB = Paddle(WHITE,10,100)
paddleB.rect.x = 970
paddleB.rect.y = 400

ball = Ball(WHITE,15,15)
ball.rect.x = 595
ball.rect.y = 595

all_sprite_list = pygame.sprite.Group()

all_sprite_list.add(paddleA)
all_sprite_list.add(paddleB)
all_sprite_list.add(ball)

#exit flag
carryOn = True

#screen update controller
clock = pygame.time.Clock()

#main program loop 
while carryOn:
    for event in pygame.event.get():
        #QUIT THE GAME IF USER CLICK CLOSE BUTTON OR PRESS X
        if event.type == pygame.QUIT:
            carryOn = False
        elif event.type == pygame.KEYDOWN:
            if event.key == pygame.K_x:
                carryOn = False

    keys = pygame.key.get_pressed()
    if keys[pygame.K_w]:
        paddleA.moveUp(5)
    if keys[pygame.K_s]:
        paddleA.moveDown(5)
    if keys[pygame.K_UP]:
        paddleB.moveUp(5)
    if keys[pygame.K_DOWN]:
        paddleB.moveDown(5)

    #game logic
    all_sprite_list.update()

    #Ball bounce back when hit wall
    if ball.rect.x > 985:
        ball.velocity[0] = -ball.velocity[0]
    if ball.rect.x < 0:
        ball.velocity[0] = -ball.velocity[0]
    if ball.rect.y > 985:
        ball.velocity[1] = -ball.velocity[1]
    if ball.rect.y < 0:
        ball.velocity[1] = -ball.velocity[1]

    #Detect collision
    if pygame.sprite.collide_mask(ball,paddleA) or pygame.sprite.collide_mask(paddleB,ball):
        ball.bounce()

        #--- DRAW CODE ---
    #Draw basic screen
    screen.fill(BLACK)
    #Draw the middle line
    pygame.draw.line(screen, WHITE, [499, 0], [499, 1000], 10)
    #Draw the Sprite on the screen
    all_sprite_list.draw(screen)
    #Update the screen
    pygame.display.flip()
    clock.tick(60)
pygame.quit()
