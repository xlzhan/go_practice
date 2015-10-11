#!/usr/bin/env python

import random

with open('./in.txt', 'wb') as f:
    numbers = range(10000)
    random.shuffle(numbers)
    for i in numbers:
        f.write(str(i) + "\n")
