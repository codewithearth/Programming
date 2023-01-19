class Solution:
    def canBeIncreasing(self, nums: list[int]) -> bool:

        
        for i in range(len(nums)): # itrerate over length of nums

            new_nums = nums[:i]+nums[i+1:] # store nums in new_nums, of removing 1 element every time.
            if (new_nums == sorted(new_nums)) & (len(new_nums) == len(set(new_nums))): # if after removing 1 element all the list is in increase order then return true
                return True
            
        if (nums == sorted(nums)) & (len(nums) == len(set(nums))): # if all the list is in increase order then return true
            return True

        return False

print(Solution.canBeIncreasing(Solution,[1,1]))