def longest_consec(strarr, k): # check longest consecutive string from slice
    
    # create two variables, one for store result and another is for check maximum length
    result = ""
    max_len = 0

    if len(strarr)<k: # if length of strings in list is smaller then consecutive length
        return result
    else:

        for i in range(len(strarr)-(k-1)): # iterate over the length of list

            # create a variable to store consecutive word
            word = ""

            for j in range(i,k+i): #iterate over consecutive length and store result in word

                word += strarr[j]

            #if length of word is bigger then all previous word then,
			#store that word in result and assign length of that word to max_len
            if len(word)>max_len:
                
                result = word
                max_len = len(word)

    return result

print(longest_consec(["zone", "abigail", "theta", "form", "libe", "zas"], -2))