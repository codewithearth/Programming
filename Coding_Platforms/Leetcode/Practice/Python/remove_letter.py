class Solution:
    def equalFrequency(self, word: str) -> bool:

        # create a list of word count
        list1 = [ word.count(i) for i in set(w for w in word)]

        for idx,num in enumerate(list1): # iterate over list

            list2 = list1[:] #create a copy of list1
            if num-1 == 0: # if word count is zero then remove that word
                list2.pop(idx)

            else: # else decrese the count by 1
                list2[idx]=num-1

            if len(set(list2)) <= 1: # if set os copy list is 1 or less then 1 then return True
                return True
       
        return False

print(Solution.equalFrequency(Solution,"abcc"))