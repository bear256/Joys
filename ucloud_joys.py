total = 0
i = 0
with open('UCloud.txt') as file:
    for text in file:
        count = len(text.split('UCanUup')) - 1
        print 'Line', i, ':', count
        total += count
        i += 1

print 'Total:', total