#!/usr/bin/env python3

import time

n = int(time.time())/10000
xor = []

def random():
    global n
    n = int(((n + 0x4ae7777) * 0x41) % 7123)
    n = (n + 0x325b28043b9b6) ^ 0x586cfc90c5b2
    n = (n * 63) % 0xC1CB7296
    n = n ^ 0x1754F2F
    n = (n*0xFF) % 4394967296
    n = n ^ 0x222F44CB
    n = n | 0x1234167890
    n = ((n + 14351514) * 32) % 7777333
    return n % 100

def XOR():
    key = [random() for i in range(len(flag))]
    for i, v in enumerate(key):
        xor.append(flag[i] ^ v)

with open('flag.txt', 'rb') as f:
    flag = f.read().strip()

if __name__ == '__main__':
    XOR()

    print('XOR: {}\n'.format(xor))
