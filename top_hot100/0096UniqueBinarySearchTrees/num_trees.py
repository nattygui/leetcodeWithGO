# coding: utf-8

class Solution:
    def numTrees(self, n):
        dp = [1, 1]
        for i in range(1, n+1):
            nums = 0
            for j in range(1, i):
                nums += dp[j-1] * dp[i-j]
            dp.append(nums)
        
        return dp[n]