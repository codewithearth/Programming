def rot13(message): # convert rot value

    # create a veriable for store results
    result = ""

    for letter in message: # iterate over the message 

        # if letter is greater then or equal to "a" or "A" and less then
        # or equal to "m" or "M" then add 13 to its ascii code to make it rot value and add that value in result
        if ("a" <= letter <= "m") or ("A" <= letter <= "M"): 
            result+=chr(ord(letter)+13)

        # if letter is greater then or equal to "n" or "N" and less then
        # or equal to "z" or "Z" then subtract 13 to its ascii code to make it rot value and add that value in result
        elif ("n" <= letter <= "z") or ("N" <= letter <= "Z"):
            result+=chr(ord(letter)-13)

        else:# else add letter to result
            result += letter
            
    return result

print(rot13("EBG13 rknzcyr."))