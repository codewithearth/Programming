def solution(roman):
    
    # create dictionary  for roman number
    roman_data = {
        "I": 1,
        "V": 5,
        "X": 10,
        "L": 50,
        "C": 100,
        "D": 500,
        "M": 1000}

	#create two variables for result and to store previous number
    result = 0
    previous_word = ""

    for i in roman[-1::-1] : #iterate over string for reverse it
        for k, v in roman_data.items(): #iterate over roman data
            if i == k: #check if string is match with roman data

                if result == 0 : #check if result is empty then add first value in it. 
                    result += v

                elif roman_data[previous_word] > v : #check if previous data is bigger-then current data then decrease current data from result
                    result -= v

                else : #if previous data is smaller-then current data then add current data into result
                    result += v
				

                previous_word = k
	
    return result

print(solution("MDCLXVI"))