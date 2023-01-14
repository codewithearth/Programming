class Solution:
    def minDominoRotations(self, tops: list[int], bottoms: list[int]) -> int:

        for val in [tops[0],bottoms[0]]: # check the first number from both list 

            if all(val in pair for pair in zip(tops, bottoms)): # if that number present in all the pair of two lists

                # subtract the maximum of that number count(of both list) from length of list.
                return len(tops) - max(tops.count(val), bottoms.count(val))
        
        # if both result is false then return -1
        return -1

print(Solution.minDominoRotations(Solution,[2,1,2,4,2,2],[5,2,6,2,3,2]))