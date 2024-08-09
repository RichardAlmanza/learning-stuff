# Sprite animations

## Content

- [Sprite animations](#sprite-animations)
  - [Content](#content)
  - [Settings per animation](#settings-per-animation)
    - [Walk](#walk)
    - [Run](#run)
    - [Idle](#idle)
    - [Jump](#jump)
    - [Fall](#fall)
    - [Death](#death)
    - [Z-Buster Stand by (Simple shoot)](#z-buster-stand-by-simple-shoot)
    - [Z-Saber Stand by (Simple slice)](#z-saber-stand-by-simple-slice)

## Settings per animation

### Walk

- source: [zero_z3standardframes.png](zero_z3standardframes.png)
- size: 32px - 48px
- separation: 2px - 0px
- offset: 174px - 2px
- frames: 6

Detailed

- frame: 0
- frame: 1 - 3
  - offset: 172px - 2px
- frame: 4
  - offset: 176px - 2px
- frame: 5
  - offset: 177px - 2px

### Run

- source: [zero_z1standardframes.png](zero_z1standardframes.png)
- size: 42px - 48px
- separation: 0px - 0px
- offset: 91px - 0px
- frames: 11
- the frame 0 is a transition from idle to run

Detailed

- frame: 0 - 2
- frame: 3
  - offset: 95px - 0px
- frame: 4
  - offset: 94px - 0px
- frame: 5
  - size: 40px - 48px
  - offset: 102px - 0px
- frame: 6
  - offset: 84px - 0px
- frame: 7
  - offset: 80px - 0px
- frame: 8
  - offset: 81px - 0px
- frame: 9
  - offset: 88px - 0px
- frame: 10
  - offset: 92px - 0px

### Idle

- source: [zero_z1standardframes.png](zero_z1standardframes.png)
- size: 42px - 48px
- separation: 1px - 0px
- offset: 188px - 0px
- frames: 6

Detailed

- frame: 0 - 1
- frame: 2
  - offset: 190px - 0px
- frame: 3 - 4
  - offset: 193px - 0px
- frame: 5
  - offset: 192px - 0px

### Jump

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 72px - 72px
- separation: 100px - 60px
- offset: 32px - 38px
- frames: 6
- the frame 0 and 1 are a transition from floor to jump

> Note: This is a well made SpriteSheet, there is no need for fixes in this animation.

### Fall

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 72px - 72px
- separation: 100px - 60px
- offset: 32px - 38px
- frames: 4
- the frame 0 and 1 are a transition from jump to fall

> Note: This is a well made SpriteSheet, there is no need for fixes in this animation.

### Death

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 72px - 72px
- separation: 100px - 60px
- offset: 30px - 40px
- frames: 8

> Note: This is a well made SpriteSheet, there is no need for fixes in this animation.

### Z-Buster Stand by (Simple shoot)

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 72px - 72px
- separation: 100px - 60px
- offset: 30px - 40px
- frames: 6

> Note: This is a well made SpriteSheet, there is no need for fixes in this animation.

### Z-Saber Stand by (Simple slice)

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 72px - 72px
- separation: 100px - 60px
- offset: 30px - 40px
- frames: 8

> Note: The afterimage of the Z-Saber I added in a different AnimatedSprite2D

- source: [Game Boy Advance - Mega Man Zero 3 - Zero.png](Game%20Boy%20Advance%20-%20Mega%20Man%20Zero%203%20-%20Zero.png)
- size: 100px - 100px
- separation: 72px - 38px
- offset: 30px - 50px
- frames: 5

> Note: I added two empty Sprites at the start and one at the end to sync with the Zero animation
