function is_unique_string(str)
    for i = 1, #str do
        if string.find(str, string.sub(str, i, i), i + 1) then
            return false
        end
    end
    
    return true
end

function find_marker(stream)
    local last_four = ""
    local index = 1
    
    for c in string.gmatch(stream, ".") do
        last_four = last_four .. c
        
        if #last_four > 14 then
            last_four = string.sub(last_four, 2)
        end
        
        if #last_four == 14 then
            if is_unique_string(last_four) then
                return index
            end
        end
        
        index = index + 1
    end
    
    return -1
end

local input = io.open("data.txt", "r"):read("*a")
print(find_marker(input))