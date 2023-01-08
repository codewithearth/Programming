def AvgNumber(nkv): # check the average number

    # take sum of all data into and
    # subtract the sum of second input list
    data = ((nkv[0]+nkv[1])*nkv[2])-sum(list(map(int, input().split(" ")))) 
    
    # if result is positive integer 
    if ((data)%nkv[1]) == 0 and data > 0: 
        return data//nkv[1]
    
    else: # else retunrn -1
        return -1

for i in range(int(input())): # iterate over testcases
    
    # take list as input and pass into AvgNumber and print result
    print(AvgNumber(list(map(int, input().split(" ")))))