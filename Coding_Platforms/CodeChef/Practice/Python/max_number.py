def MaxNumber(test_data): # create function to check max number

    result = 0
    for k,v in enumerate(test_data): # iterate over list 

        if k == 0: # if value is first

            if (v or len(test_data)) == 1: # if value and size of list is 1 then add 1 to result
                result+=1

        elif test_data[k-1] == 0 and v == 1: # if previous value is 0 and current value is 1 then add 1 in result
            result +=1

        elif k == len(test_data)-1 and v == 0: # if last value is 0 then add 1 to result
            result+=1

    return result

for i in range(int(input())):# iterate over test case

    # take test size
    test_size = int(input())

    # Put list in MaxNumber function then print max number
    print(MaxNumber(list(map(int,input().strip()))))

