# coding: utf-8

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        return self.dfs(root)
    def dfs(self, root):
        if root == None:
            return 0
        num1 = self.dfs(root.left)
        num2 = self.dfs(root.right)
        return max(num1, num2) + 1
    