from textwrap import wrap

file1 = open('data.txt', 'r')
Lines = file1.readlines()
alpha = [*'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ']
count = 0

def matchStrings(str1, str2):
    for i in str1:
        if i in str2:
            return alpha.index(i)

for backpack in Lines:
    count += 1
    backpackLength = int(len(backpack.strip()) / 2)
    compartments = wrap(backpack.strip(), backpackLength)
    count += matchStrings(compartments[0], compartments[1])

print('Total Count: ' + str(count))