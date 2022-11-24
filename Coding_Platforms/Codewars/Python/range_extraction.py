def solution(args): # range extraction

    #create three variables 
    # one for result, one for last value and one for count
    result = ""
    last_value = len(args)-1
    count = 0

    for i,_ in enumerate(args): # iterate over list

        if i<count or i == last_value: # skip iteration  until i != count and if i is last value
            continue

        # assign i to the count 
        # add number to the result
        count = i
        result+= str(args[i])

        while True:# a infinite loop until numbers is continuous

            if count == last_value:# if count is equal to last value then break the look
                break

            elif (count < last_value) & (args[count]+1 == args[count+1]):# if count is less then last value and future value is continuous to current value then add one to count
                count+=1

            else: # if value is not continuous then break the loop
                break

        if count - i > 1: # if count is greater then i then use "-" to result
            result+="-"

        else: # else use "," to result
            result+=","
    
    # add last value to the result
    result += str(args[last_value])
    return result

print(solution([-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20,22]))
