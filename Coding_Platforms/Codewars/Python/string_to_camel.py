def to_camel_case(text):#convert normal string to camel case

    # create two variables, one for result and one for flag
    result = ""
    flag = 0 

    for i in text:#iterate over string

        if i == "_" or i == "-": #if string is underscore or hypen then continue for next iteration and add 1 into flag
            flag = 1
            continue

        elif flag == 1: # if flag is 1 then uppercase the string and continue for nexr iteration and add 0 into flag

            result += i.upper()
            flag = 0
            continue

        result += i

    return result

print(to_camel_case("the-Stealth-Warrior"))