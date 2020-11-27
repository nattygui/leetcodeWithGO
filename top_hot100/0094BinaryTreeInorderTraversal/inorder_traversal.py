# coding: utf-8

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def inorderTraversal(self, root):
        result = []
        self.dfs(root, result)
        return result
    
    def dfs(self, root, result):
        if root == None:
            return
        self.dfs(root.left, result)
        result.append(root.val)
        self.dfs(root.right, result)