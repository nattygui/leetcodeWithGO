# coding: utf-8
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isValidBST(self, root):
        pre_value = float('-inf')
        stack = []
        # 中序遍历非递归方法
        while (root != None) or len(stack) > 0:
            # 获取当前的节点 以及左节点
            while root != None:
                stack.append(root)
                root = root.left
            
            root = stack[-1]
            stack = stack[:-1]
            if root.val <= pre_value:
                return False

            pre_value = root.val
            root = root.right
        
        return True
