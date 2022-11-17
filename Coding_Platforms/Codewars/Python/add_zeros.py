def move_zeros(lst): # Move all zeros to the end

    # create two lists
    # one for non_zeros and another is for zeros
    non_zeros = []
    zeros = []

    for i in lst: # iterate over the list

        if i == 0: # if integer is zero then add in zeros list

            zeros.append(i)
            continue

        else: # if integer is non zero then add in non_zeros list

            non_zeros.append(i)
    
    # return the combination of non zeros and zeros 
    return non_zeros + zeros


print(move_zeros([1, 0, 1, 2, 0, 1, 3]))