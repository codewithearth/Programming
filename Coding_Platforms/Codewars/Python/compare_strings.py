def in_array(array1, array2):#compare strings from two arrays

    #create a variable to store result
    result = []

    for data1 in array1: #iterate over first array
        for data2 in array2:#iterate over second array

            if data1 in data2:#if string of 1st array is a subset of string of 2nd array,
                result.append(data1)

    # use set to remove duplicates and sort the result to get lexicographical order result
    data = list(set(result))
    data.sort()
    return data

print(in_array(["live", "arp", "strong"],["lively", "alive", "harp", "sharp", "armstrong"]))