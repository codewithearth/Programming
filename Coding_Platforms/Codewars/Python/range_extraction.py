def solution(args):

    result = ""
    last_value = len(args)-1
    count = 0
    for i in args:
        result+= str(args[i])
        count = i
        while True:
            if (count < last_value) & (args[count]+1 == args[count+1]):
                print(args[count]+1, " ", args[count+1])
                count+=1
            else:
                break
        if count - i > 1:
            result+="-"
        else:
            result+=","
    
    result += str(args[last_value])
    return result

print(solution([-10, -9, -8, -6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]))


# l := len(list) - 1
#   for i, j := 0, 0; i < l; i = j {
#     s += strconv.Itoa(list[i])
#     for j = i; (j < l) && (list[j] + 1 == list[j + 1]); { 
#         j++ 
#     }
#     if j - i > 1 { 
#         s += "-" 
        
#     } else {
#         s += "," 
#     }
#     if i == j {
#         j++ 
#     }
#   }
#   s += strconv.Itoa(list[l])
#   return