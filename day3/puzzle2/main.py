file1 = open('data.txt', 'r')
Lines = file1.readlines()
alpha = [*'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ']
count = 0
groupCounts = 0
groups = []

def matchStrings(str1, str2, str3):
    for i in str1:
        if i in str2 and i in str3:
            return alpha.index(i) + 1

for backpack in Lines:
    try: groups[groupCounts]
    except IndexError: groups.append([])

    if len(groups[groupCounts]) == 3:
        groupCounts += 1
        groups.append([])

    groups[groupCounts].append(backpack.strip())

for group in groups:
    count += matchStrings(group[0], group[1], group[2])

print('Total count: ' + str(count))