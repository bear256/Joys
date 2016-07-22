total = 0
i = 0
File.open('UCloud.txt', 'r').each do |text|
    count = text.split('UCanUup').length - 1
    print "Line #{i} : #{count}\n"
    total += count
    i += 1
end

print "Total: #{total}\n"