def duplicate_count(text):
    
    #create two variables, one is for result and another is for store all letters in lower case
    result = 0
    text_lower = text.lower()

    for i in set(text_lower): # iterate over string that converted into set

        if text_lower.count(i) > 1: # if count of letter is greater then 1, then add it into duplicate 
            result +=1
    
    return result


print(duplicate_count("ABBA"))