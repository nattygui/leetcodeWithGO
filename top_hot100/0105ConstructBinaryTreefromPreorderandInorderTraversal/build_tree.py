# coding: utf-8

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        root = self.newNode(preorder[0])

        rootIndex = self.findIndex(inorder, preorder[0])

        leftPreorder = preorder[1: rootIndex+1]
        rightPreorder = preorder[rootIndex+1:]

        leftInorder = inorder[:rootIndex]
        rightInorder = inorder[rootIndex+1:]

        leftTree = self.buildTree(leftPreorder, leftInorder)
        rightTree = self.buildTree(rightPreorder, rightInorder)

        root.left = leftTree
        root.right = rightTree

        return root

    
    def newNode(self, val):
        return TreeNode(val)

    
    def findIndex(inorder, val):
        for index, num in enumerate(inorder):
            if num == val:
                return index
            
        return -1