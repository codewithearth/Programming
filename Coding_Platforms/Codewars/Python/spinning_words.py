def spin_words(sentence):

    #create list for store words from sentence and
    #one veriable for index 
    word_list = [word for word in sentence.split(" ")]
    idx  = 0

    for word in word_list: # iterate over the list

        if len(word) >4: # if word lenght is greater then four then reverse the word

            word_list[idx] = word[::-1]

        idx +=1

    # join the word of word_list by space
    return " ".join(word_list)


print(spin_words("i love pizza and burger"))