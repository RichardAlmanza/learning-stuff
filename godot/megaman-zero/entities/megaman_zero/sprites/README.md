# Sprite animations

## Content

- [Sprite animations](#sprite-animations)
  - [Content](#content)
  - [Settings per animation](#settings-per-animation)
    - [Walk](#walk)
    - [Run](#run)
    - [Idle](#idle)
    - [Jump](#jump)

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

- source: [zero_z1standardframes.png](zero_z1standardframes.png)
- size: 40px - 48px
- separation: 0px - 0px
- offset: 86px - 5px
- frames: 5
- the frame 0 and 1 are a transition from floor to jump

Detailed

- frame: 0
- frame: 1
  - size: 34px - 56px
  - offset: 93px - 38px
- frame: 2
  - size: 34px - 56px
  - offset: 93px - 39px
- frame: 3
  - size: 36px - 56px
  - offset: 87px -39px
- frame: 4
  - size: 36px - 56px
  - offset: 88px - 39px
