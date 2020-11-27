# coding: utf-8

class Solution:
    def subsets(self, nums):
        result = [[]]
        self.getall(nums, [], result)
        return result
    
    def getall(self, nums, one_set, result):
        if len(nums) == 0:
            return 
        for index, num in enumerate(nums):
            temp = one_set[:] # copy一份新的数组
            temp.append(num)
            result.append(temp)
            self.getall(nums[index+1:], temp, result)